package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
	"os"
	"strings"
)

func main() {
	//updateCPU := 2
	fname := "nginx-deployment.yaml"
	data, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ReadCPU(data))
	UpdateCPU(data, fname)

}

func ReadCPU(data []byte) string {
	var CpuValue string
	decoder := scheme.Codecs.UniversalDeserializer()

	for _, resourceYAML := range strings.Split(string(data), "---") {

		if len(resourceYAML) == 0 {
			continue
		}

		obj, groupVersionKind, err := decoder.Decode(
			[]byte(resourceYAML),
			nil,
			nil)
		if err != nil {
			log.Print(err)
			continue
		}

		if groupVersionKind.Group == "apps" &&
			groupVersionKind.Version == "v1" &&
			groupVersionKind.Kind == "Deployment" {
			deployment := obj.(*appsv1.Deployment)
			CpuValue = deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu().String()
		}
	}
	return CpuValue
}

func int32Ptr(i int32) *int32 { return &i }

func UpdateCPU(data []byte, fname string) {

	//mydeployment := &appsv1.Deployment{}
	//err := yaml.Unmarshal(data, &mydeployment)
	//if err != nil {
	//	fmt.Println("Yaml Unmarshall", err)
	//}
	decoder := scheme.Codecs.UniversalDeserializer()
	obj, _, err := decoder.Decode(
		data,
		nil,
		&appsv1.Deployment{})

	if err != nil {
		log.Print(err)
	}
	mydeployment := obj.(*appsv1.Deployment)

	mydeployment.Spec.Replicas = int32Ptr(2)

	mydeployment.Spec.Template.Spec.Containers[0].Resources.Requests = make(map[corev1.ResourceName]resource.Quantity)
	mydeployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceCPU] = *resource.NewMilliQuantity(5300, resource.DecimalSI)
	mydeployment.Spec.Template.Spec.Containers[0].Resources.Requests[corev1.ResourceMemory] = *resource.NewQuantity(5*1024*1024*1024, resource.BinarySI)
	newFile, _ := os.Create(fname)
	//manifest, _ := os.Open(fname)
	y := printers.YAMLPrinter{}
	defer newFile.Close()
	y.PrintObj(obj, newFile)
}
