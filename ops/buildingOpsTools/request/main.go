package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader("hello, this is hope!"))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))

	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp2, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp2.Body.Close()

	content2, err := io.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content2))
}
