// hackerrank.com/challenges/breaking-best-and-worst-records/problem
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
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	scoreL := scores[0]
	scoreH := scores[0]
	var scoresL []int32
	var scoresH []int32
	var scoresHL []int32
	for _, v := range scores {
		if scoreL > v {
			scoreL = v
			scoresL = append(scoresL, scoreL)
		}
		if scoreH < v {
			scoreH = v
			scoresH = append(scoresH, scoreH)
		}
	}
	fmt.Println("score Low->", scoreL)
	fmt.Println("score High->", scoreH)
	fmt.Println("scores Low-->", scoresL)
	fmt.Println("scores Hight-->", scoresH)
	scoresHL = append(scoresHL, int32(len(scoresH)), int32(len(scoresL)))
	return scoresHL
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	scoresTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var scores []int32

	for i := 0; i < int(n); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	result := breakingRecords(scores)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
