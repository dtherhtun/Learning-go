package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("Unable to create request")
	}

	/*
		buffer := &bytes.Buffer{}
		enc := base64.NewEncoder(base64.URLEncoding, buffer)
		enc.Write([]byte("user:passw0rd"))
		enc.Close()

		encodedCreds, err := buffer.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			log.Fatalln("Failed to read encoded buffer")
		}

		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
	*/
	req.SetBasicAuth("user", "passw0rd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Unable to read response")
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}

	fmt.Println(string(content))
	fmt.Println(resp.StatusCode)
}
