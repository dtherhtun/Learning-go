package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	cbyte := flag.Bool("b", false, "Count Bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *cbyte))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		if countBytes {
			wc = len(scanner.Bytes())
			break
		}
		wc++
	}

	return wc
}
