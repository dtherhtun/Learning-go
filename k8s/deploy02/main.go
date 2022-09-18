package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	updateCPU := 2
	fname := "deployment.yaml"
	data, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ReadCPU(data))
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
