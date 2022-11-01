// https://www.hackerrank.com/challenges/30-review-loop/problem
// Input
// 2
// Hacker
// Rank
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []string
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			data = append(data, text)
		} else {
			break
		}
	}

	words := data[1:]

	for _, v := range words {
		fmt.Println(solution(v))
	}

}

func solution(word string) string {
	wordData := strings.Split(word, "")
	var even, odd string

	for i, v := range wordData {
		if i%2 == 0 {
			even += v
		} else {
			odd += v
		}

	}
	return even + " " + odd
}
