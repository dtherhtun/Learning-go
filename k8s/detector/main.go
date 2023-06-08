package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	labelSelector := "app=asc-api"
	threshold := resource.MustParse("0.1G")

	interval := 3 * time.Minute

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	ctx := context.Background()
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	for range ticker.C {
		runCode(*kubeconfig, labelSelector, threshold, ctx)
	}
}

func runCode(kubeconfig, labelSelector string, threshold resource.Quantity, ctx context.Context) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("applications").List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, pod := range pods.Items {
		podName := pod.ObjectMeta.Name
		containerName := pod.Spec.Containers[0].Name // Assuming only one container per pod

		metrics, err := metricsClientset.MetricsV1beta1().PodMetricses("applications").List(ctx, metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(labels.Set(pod.ObjectMeta.Labels)).String(),
		})
		if err != nil {
			log.Fatal(err)
		}
		var memoryUsage resource.Quantity
		for _, podMetrics := range metrics.Items {
			if podMetrics.Name == podName {
				for _, container := range podMetrics.Containers {
					if container.Name == containerName {
						memoryUsage = container.Usage[v1.ResourceMemory]
						break
					}
				}
			}
		}
		result := memoryUsage.Cmp(threshold)
		if result == -1 {
			fmt.Println("quantity1 is less than quantity2")
		} else if result == 0 {
			fmt.Println("quantity1 is equal to quantity2")
		} else if result == 1 {
			fmt.Println("quantity1 is greater than quantity2")
			err = clientset.CoreV1().Pods("applications").Delete(context.TODO(), podName, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalf("Failed to delete pod %s: %v", podName, err)
			}
		} else {
			log.Fatal("Invalid comparison result")
		}
	}
}
