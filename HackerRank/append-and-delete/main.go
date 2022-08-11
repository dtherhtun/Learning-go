// https://www.hackerrank.com/challenges/append-and-delete/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	var cLength, min, tLength int32
	var result string
	tLength = int32(len(s) + len(t))
	if len(s) < len(t) {
		min = int32(len(s))
	} else {
		min = int32(len(t))
	}

	for i := int32(0); i < min; i++ {
		if s[i] == t[i] {
			cLength++
		} else {
			break
		}
	}
	if tLength-2*cLength > k {
		result = "No"
	} else if (tLength-2*cLength)%2 == k%2 {
		result = "Yes"
	} else if tLength-k < 0 {
		result = "Yes"
	} else {
		result = "No"
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
