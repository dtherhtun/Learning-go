package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here
	date := time.Date(int(year), time.Month(3), 0, 0, 0, 0, 0, time.UTC)
	feb := date.Day()
	d_o_p := 256 - (31 + feb + 31 + 30 + 31 + 30 + 31 + 31)
	if year > 1918 {
		return fmt.Sprintf("%d.09.%d", d_o_p, date.Year())
	} else if year < 1918 {
		if year%4 == 0 {
			return fmt.Sprintf("12.09.%d", year)
		} else {
			return fmt.Sprintf("13.09.%d", year)
		}
	} else {
		return fmt.Sprintf("26.09.%d", year)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
