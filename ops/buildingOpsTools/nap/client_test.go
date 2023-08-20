package nap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	client := NewClient()
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	resource := NewResource("/get", "GET", router)
	if err := client.ProcessRequest("https://httpbin.org", resource, nil); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}