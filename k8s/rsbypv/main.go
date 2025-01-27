package main

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type WorkloadReference struct {
	Kind      string
	Name      string
	Namespace string
}

func FindWorkloadsUsingPV(clientset *kubernetes.Clientset, pvName string) (*WorkloadReference, error) {
	fmt.Println(pvName)
	pv, err := clientset.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get PV %s: %v", pvName, err)
	}

	if pv.Spec.ClaimRef == nil {
		return nil, fmt.Errorf("PV %s is not bound to any PVC", pvName)
	}

	pvcNamespace := pv.Spec.ClaimRef.Namespace
	pvcName := pv.Spec.ClaimRef.Name

	fmt.Println(pvcNamespace, pvcName)

	//pvc, err := clientset.CoreV1().PersistentVolumeClaims(pvcNamespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get PVC %s/%s: %v", pvcNamespace, pvcName, err)
	//}

	pods, err := clientset.CoreV1().Pods(pvcNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods in namespace %s: %v", pvcNamespace, err)
	}
	fmt.Println(len(pods.Items))

	for _, pod := range pods.Items {
		for _, volume := range pod.Spec.Volumes {
			if volume.PersistentVolumeClaim != nil && volume.PersistentVolumeClaim.ClaimName == pvcName {
				fmt.Println(pod.Name)
				for _, owner := range pod.OwnerReferences {
					fmt.Println(owner.Kind, owner.Name)
					if owner.Kind == "ReplicaSet" {
						rs, err := clientset.AppsV1().ReplicaSets(pvcNamespace).Get(context.TODO(), owner.Name, metav1.GetOptions{})
						if err != nil {
							continue
						}

						for _, rsOwner := range rs.OwnerReferences {
							if rsOwner.Kind == "Deployment" {
								return &WorkloadReference{
									Kind:      "Deployment",
									Name:      rsOwner.Name,
									Namespace: pvcNamespace,
								}, nil
							}
						}
					} else if owner.Kind == "StatefulSet" {
						return &WorkloadReference{
							Kind:      "StatefulSet",
							Name:      owner.Name,
							Namespace: pvcNamespace,
						}, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("no Deployment or StatefulSet found using PV %s", pvName)
}

func main() {

	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err)
	//}
	//

	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(fmt.Errorf("error building kubeconfig: %v", err))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("error creating clientset: %v", err))
	}

	workload, err := FindWorkloadsUsingPV(clientset, "pvc-efda7a42-eb01-4135-a5ee-109640ac750a")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %s %s/%s using the PV\n", workload.Kind, workload.Namespace, workload.Name)
}
