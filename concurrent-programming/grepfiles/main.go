package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	searchString := os.Args[1]
	filenames := os.Args[2:]
	for _, filename := range filenames {
		go grepFile(filename, searchString)
	}
	time.Sleep(2 * time.Second)
}

func grepFile(filename string, searchString string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(content), searchString) {
		fmt.Println(filename, "contains a match with", searchString)
	} else {
		fmt.Println(filename, "does NOT contain a match with", searchString)
	}
}
