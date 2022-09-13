package main

import (
	"log"
	"os"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {

	// Load the file into a buffer
	fname := "deployment.yaml"
	data, err := os.ReadFile(fname) //ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Create a runtime.Decoder from the Codecs field within
	// k8s.io/client-go that's preloaded with the schemas for all
	// the standard Kubernetes resource types.
	decoder := scheme.Codecs.UniversalDeserializer()

	for _, resourceYAML := range strings.Split(string(data), "---") {

		// skip empty documents, `Decode` will fail on them
		if len(resourceYAML) == 0 {
			continue
		}

		// - obj is the API object (e.g., Deployment)
		// - groupVersionKind is a generic object that allows
		//   detecting the API type we are dealing with, for
		//   accurate type casting later.
		obj, groupVersionKind, err := decoder.Decode(
			[]byte(resourceYAML),
			nil,
			nil)
		if err != nil {
			log.Print(err)
			continue
		}

		// Figure out from `Kind` the resource type, and attempt
		// to cast appropriately.
		if groupVersionKind.Group == "apps" &&
			groupVersionKind.Version == "v1" &&
			groupVersionKind.Kind == "Deployment" {
			deployment := obj.(*appsv1.Deployment)
			log.Print(deployment.Spec.Template.Spec.Containers[0].Resources.Requests.Cpu())
		}
	}
}
