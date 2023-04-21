package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func countWord(url string, frequency map[string]int) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWord(url, frequency)
	}
	time.Sleep(10 * time.Second)

	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
}
