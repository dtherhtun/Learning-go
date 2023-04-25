package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func countLetters(url string, frequency map[string]int, mu *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	mu.Lock()
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	mu.Unlock()
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	mu := sync.Mutex{}
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, frequency, &mu)
	}
	time.Sleep(10 * time.Second)
	mu.Lock()
	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
	mu.Unlock()
}
