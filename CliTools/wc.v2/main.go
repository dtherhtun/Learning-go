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
	fname := flag.String("f", "", "name of the file")
	flag.Parse()

	if *fname != "" {
		fInfo, err := os.Stat(*fname)
		if err != nil {
			fmt.Printf("Can't open file: %q", err)
			os.Exit(1)
		}

		if fInfo.Size() == 0 {
			fmt.Println("empty file")
			os.Exit(1)
		}

		f, err := os.Open(*fname)
		if err != nil {
			fmt.Printf("Can't open file: %q", err)
			os.Exit(1)
		}
		defer f.Close()
		fmt.Println(count(f, *lines, *cbyte))
	}

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
