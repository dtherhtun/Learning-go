package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.FileServer(http.Dir("assets"))); err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
