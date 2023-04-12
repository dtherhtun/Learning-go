package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	searchString := os.Args[1]
	path := os.Args[2]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range files {
		go grepPath(path, fileInfo, searchString)
	}
	time.Sleep(2 * time.Second)
}

func grepPath(path string, dirEntry fs.DirEntry, searchString string) {
	fullPath := filepath.Join(path, dirEntry.Name())
	if dirEntry.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			go grepPath(fullPath, file, searchString)
		}
	} else {
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(content), searchString) {
			fmt.Println(fullPath, "contains a match with", searchString)
		} else {
			fmt.Println(fullPath, "does NOT contain a match with", searchString)
		}
	}
}
