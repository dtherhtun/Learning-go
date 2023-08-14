package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var router = NewRouter()

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln("unable to read response")
	}
	defer resp.Body.Close()
	router.Process(resp)
}

func init() {
	router.Register(200, func(resp *http.Response) {
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("unable to read content")
		}
		fmt.Println(string(content))
	})

	router.Register(404, func(resp *http.Response) {
		log.Fatalln("not found (404): ", resp.Request.URL.String())
	})
}
