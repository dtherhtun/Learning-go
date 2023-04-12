package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filenames := os.Args[1:]
	for _, filename := range filenames {
		go printFile(filename)
	}
	time.Sleep(2 * time.Second)
}

func printFile(file string) {
	data, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
