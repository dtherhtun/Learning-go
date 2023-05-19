package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

func fileSearch(dir, filename string, wg *sync.WaitGroup, mu *sync.Mutex, matches *[]string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Directory read failed:", err)
	}
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mu.Lock()
			*matches = append(*matches, fpath)
			mu.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mu, matches)
		}
	}
	wg.Done()
}

func main() {
	results := make([]string, 0)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg, &mu, &results)
	wg.Wait()
	mu.Lock()
	sort.Strings(results)
	fmt.Println(strings.Join(results, "\n"))
	mu.Unlock()
}
