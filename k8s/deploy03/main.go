package main

import (
	"log"
	"os"

	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Load the Kubernetes configuration from the default location (~/.kube/config)
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Kubernetes client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Get the deployment with the specified name
	deployment, err := client.AppsV1().Deployments("default").Get("my-deployment", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new YAML printer
	yamlPrinter := printers.NewYAMLPrinter(printers.PrintOptions{
		// Use the jsonpath-file option to specify a file containing a jsonpath template
		// that filters out the creationTimestamp and status fields
		JSONPathFile: "jsonpath-template.txt",
	})

	// Print the deployment using the YAML printer
	err = yamlPrinter.PrintObj(deployment, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

