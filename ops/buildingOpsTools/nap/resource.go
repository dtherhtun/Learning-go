package nap

import (
	"bytes"
	"log"
	"text/template"
)

type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}

func NewResource(endpoint, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}

	template, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalln("Unable to parse endpoint")
	}
	buffer := &bytes.Buffer{}
	template.Execute(buffer, params)

	//endpoint, err := io.ReadAll(buffer)
	//if err != nil {
	//	log.Fatalln("Unable to read endpoint")
	//}
	//
	//return string(endpoint)
	return buffer.String()
}
