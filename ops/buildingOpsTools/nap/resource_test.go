package nap

import "testing"

func TestEndpointTemplate(t *testing.T) {
	res := &RestResource{
		Endpoint: "/user/{{.user}}",
		Method:   "GET",
		Router:   NewRouter(),
	}
	renderedEndpoint := res.RenderEndpoint(map[string]string{
		"user": "dther",
	})
	if renderedEndpoint != "/user/dther" {
		t.Fail()
	}
}
