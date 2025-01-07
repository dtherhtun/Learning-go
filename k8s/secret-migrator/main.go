package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/homedir"

	"k8s.io/client-go/kubernetes/scheme"

	sealsecret "github.com/bitnami-labs/sealed-secrets/pkg/apis/sealedsecrets/v1alpha1"
	sealsecretclient "github.com/bitnami-labs/sealed-secrets/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

type ExportMode string

const (
	ExportModeYAML   ExportMode = "yaml"
	ExportModeDirect ExportMode = "direct"
	maxWorkers                  = 10
)

type ProcessResult struct {
	secretName string
	success    bool
	skipped    bool
	err        error
}

type SecretProcessor struct {
	newNamespace        string
	publicKey           *rsa.PublicKey
	exportMode          ExportMode
	outputDir           string
	sealedSecretsClient *sealsecretclient.Clientset
	results             chan ProcessResult
	wg                  sync.WaitGroup
}

func main() {
	start := time.Now()
	var kubeconfig, oldContext, newContext, oldNamespace, newNamespace, exportMode, outputDir, ssns string

	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "path to kubeconfig file")
	}
	flag.StringVar(&oldContext, "old-context", "", "Source cluster context")
	flag.StringVar(&newContext, "new-context", "", "Destination cluster context")
	flag.StringVar(&oldNamespace, "old-namespace", "", "Source namespace")
	flag.StringVar(&newNamespace, "new-namespace", "", "Destination namespace")
	flag.StringVar(&exportMode, "export-mode", string(ExportModeYAML), "Export mode: 'yaml' or 'direct'")
	flag.StringVar(&outputDir, "output-dir", "sealed-secrets", "Output directory for YAML files")
	flag.StringVar(&ssns, "sealed-secret-ns", "kube-system", "Sealed secrets namespace")

	flag.Parse()

	validateFlags(oldContext, newContext, oldNamespace, newNamespace)

	clients := setupClients(kubeconfig, oldContext, newContext, ssns)

	if ExportMode(exportMode) == ExportModeYAML {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			log.Fatalf("Error creating output directory: %v", err)
		}
	}

	secrets, err := clients.oldClient.CoreV1().Secrets(oldNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error getting secrets from old cluster: %v", err)
	}

	fmt.Printf("Found %d secrets in namespace %s (context: %s)\n", len(secrets.Items), oldNamespace, oldContext)

	processor := &SecretProcessor{
		newNamespace:        newNamespace,
		publicKey:           clients.publicKey,
		exportMode:          ExportMode(exportMode),
		outputDir:           outputDir,
		sealedSecretsClient: clients.sealedSecretsClient,
		results:             make(chan ProcessResult, len(secrets.Items)),
	}

	semaphore := make(chan struct{}, maxWorkers)

	for i := range secrets.Items {
		secret := secrets.Items[i]
		processor.wg.Add(1)
		go func(s corev1.Secret) {
			defer processor.wg.Done()
			semaphore <- struct{}{}
			processor.processSecret(s)
			<-semaphore
		}(secret)
	}

	processor.wg.Wait()
	close(processor.results)

	var successCount, skippedCount, failedCount int
	for result := range processor.results {
		switch {
		case result.skipped:
			skippedCount++
		case result.success:
			successCount++
		default:
			failedCount++
			log.Printf("Error processing secret %s: %v", result.secretName, result.err)
		}
	}

	printSummary(len(secrets.Items), successCount, skippedCount, failedCount, outputDir, ExportMode(exportMode), time.Since(start))
}

func (p *SecretProcessor) processSecret(secret corev1.Secret) {
	if secret.Type == corev1.SecretTypeServiceAccountToken || secret.Type == "helm.sh/release.v1" {
		p.results <- ProcessResult{
			secretName: secret.Name,
			skipped:    true,
		}
		return
	}

	cleanSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        secret.Name,
			Namespace:   p.newNamespace,
			Labels:      secret.Labels,
			Annotations: secret.Annotations,
		},
		Type: secret.Type,
		Data: secret.Data,
	}

	ss, err := sealsecret.NewSealedSecret(scheme.Codecs, p.publicKey, cleanSecret)
	if err != nil {
		p.results <- ProcessResult{
			secretName: secret.Name,
			err:        fmt.Errorf("error sealing secret: %v", err),
		}
		return
	}

	var processErr error
	switch p.exportMode {
	case ExportModeYAML:
		processErr = p.exportSealedSecretToYAML(ss)
	case ExportModeDirect:
		processErr = p.createSealedSecret(ss)
	}

	p.results <- ProcessResult{
		secretName: secret.Name,
		success:    processErr == nil,
		err:        processErr,
	}
}

func (p *SecretProcessor) exportSealedSecretToYAML(sealedSecret *sealsecret.SealedSecret) error {
	sealedSecret.TypeMeta = metav1.TypeMeta{
		APIVersion: "bitnami.com/v1alpha1",
		Kind:       "SealedSecret",
	}

	yamlData, err := yaml.Marshal(sealedSecret)
	if err != nil {
		return fmt.Errorf("error marshaling to YAML: %v", err)
	}

	filename := filepath.Join(p.outputDir, fmt.Sprintf("%s.yaml", sealedSecret.Name))
	return os.WriteFile(filename, yamlData, 0644)
}

func (p *SecretProcessor) createSealedSecret(ss *sealsecret.SealedSecret) error {
	_, err := p.sealedSecretsClient.BitnamiV1alpha1().SealedSecrets(p.newNamespace).Create(
		context.TODO(),
		ss,
		metav1.CreateOptions{},
	)
	return err
}

type clients struct {
	oldClient           *kubernetes.Clientset
	newClient           *kubernetes.Clientset
	sealedSecretsClient *sealsecretclient.Clientset
	publicKey           *rsa.PublicKey
}

func setupClients(kubeconfig, oldContext, newContext, ssns string) *clients {
	oldClient, err := createKubernetesClientWithContext(kubeconfig, oldContext)
	if err != nil {
		log.Fatalf("Error creating client for old context: %v", err)
	}

	newClient, err := createKubernetesClientWithContext(kubeconfig, newContext)
	if err != nil {
		log.Fatalf("Error creating client for new context: %v", err)
	}

	sealedSecretsClient, err := createSealedSecretsClient(kubeconfig, newContext)
	if err != nil {
		log.Fatalf("Error creating sealed secrets client: %v", err)
	}

	publicKey, err := fetchSealedSecretsPublicKey(newClient, ssns)
	if err != nil {
		log.Fatalf("Error fetching public key: %v", err)
	}

	return &clients{
		oldClient:           oldClient,
		newClient:           newClient,
		sealedSecretsClient: sealedSecretsClient,
		publicKey:           publicKey,
	}
}

func validateFlags(oldContext, newContext, oldNamespace, newNamespace string) {
	if oldContext == "" || newContext == "" {
		log.Fatal("Both old and new contexts must be specified")
	}
	if oldNamespace == "" || newNamespace == "" {
		log.Fatal("Both old and new namespaces must be specified")
	}
}

func printSummary(total, success, skipped, failed int, outputDir string, exportMode ExportMode, duration time.Duration) {
	fmt.Printf("\nMigration complete in %v:\n", duration)
	fmt.Printf("- Total secrets found: %d\n", total)
	fmt.Printf("- Successfully processed: %d\n", success)
	fmt.Printf("- Skipped: %d\n", skipped)
	fmt.Printf("- Failed: %d\n", failed)
	if exportMode == ExportModeYAML {
		fmt.Printf("- Output directory: %s\n", outputDir)
	}
}

func createKubernetesClientWithContext(kubeConfig string, context string) (*kubernetes.Clientset, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfig}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading kubeconfig for context %s: %v", context, err)
	}

	return kubernetes.NewForConfig(config)
}

func createSealedSecretsClient(kubeconfig string, context string) (*sealsecretclient.Clientset, error) {
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading kubeconfig: %v", err)
	}

	return sealsecretclient.NewForConfig(config)
}

func fetchSealedSecretsPublicKey(client *kubernetes.Clientset, namespace string) (*rsa.PublicKey, error) {
	secrets, err := client.CoreV1().Secrets(namespace).List(
		context.TODO(),
		metav1.ListOptions{
			LabelSelector: "sealedsecrets.bitnami.com/sealed-secrets-key=active",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching sealed secrets key: %v", err)
	}

	var keySecret *corev1.Secret
	mostRecent := secrets.Items[0]
	for _, secret := range secrets.Items[1:] {
		if secret.CreationTimestamp.After(mostRecent.CreationTimestamp.Time) {
			mostRecent = secret
		}
	}
	keySecret = &mostRecent
	block, _ := pem.Decode(keySecret.Data["tls.crt"])
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block from certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %v", err)
	}

	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("certificate public key is not RSA")
	}

	return publicKey, nil
}
