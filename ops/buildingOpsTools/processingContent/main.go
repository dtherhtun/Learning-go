package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GetResponse struct {
	Origin  string            `json:"origin"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (r *GetResponse) toString() string {
	s := fmt.Sprintf("GET Response\nOrigin IP: %s\nRequest URL: %s\n", r.Origin, r.URL)
	for k, v := range r.Headers {
		s += fmt.Sprintf("Header: %s = %s\n", k, v)
	}
	return s
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	respContent := &GetResponse{}
	json.Unmarshal(content, respContent)

	fmt.Println(respContent.toString())
}
