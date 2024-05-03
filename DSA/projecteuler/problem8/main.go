package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readInt(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	var result [][]int
	var sequence []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		sequence = append(sequence, x)

		if len(sequence) > 13 {
			sequence = sequence[1:]
		}

		if len(sequence) == 13 && !containsZero(sequence) {
			seqCopy := make([]int, len(sequence))
			copy(seqCopy, sequence)
			result = append(result, seqCopy)
		}
	}
	return result
}

func containsZero(sequence []int) bool {
	for _, digit := range sequence {
		if digit == 0 {
			return true
		}
	}
	return false
}

func greatestProduct(a [][]int) int {
	result := 0
	for i, v := range a {
		product := 1
		for j := 0; j < len(v); j++ {
			product = product * a[i][j]
		}
		if result < product {
			result = product
		}
	}

	return result
}

func main() {
	fd, _ := os.Open("./numbers.txt")
	a := readInt(fd)
	fmt.Println(greatestProduct(a))
	fmt.Println(a[55])
}
