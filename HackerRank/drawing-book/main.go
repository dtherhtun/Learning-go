// https://www.hackerrank.com/challenges/drawing-book/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
    // Write your code here
    var from_start float64
    var from_end float64
    for i := 1; i < int(n);i++ {
        if i < int(p) {
            from_start++
        }
    }
    for i := int(n-1); i > 0; i-- {
        if i > int(p) {
            from_end++
        }
    }
    if from_start < from_end {
        return int32(math.Round(from_start / 2))
    } else if n%2 == 0 && n-p == 1 && p > 1 {
        return 1
    } else {
        return int32(math.Round(from_end+1)) / 2
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    p := int32(pTemp)

    result := pageCount(n, p)

    fmt.Fprintf(writer, "%d\n", result)

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

