// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
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
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */
 func rNum(num int32) int32 {
    var res int32
    for num>0 {
        remainder := num % 10
        res = (res * 10) + remainder
        num /= 10
    }
    return res
}

func beautifulDays(i int32, j int32, k int32) int32 {
    // Write your code here
    var beautiful_day int32
    for i := i; i <= j ; i++ {
        if (i - rNum(i))%k == 0 {
            beautiful_day++
        }
    }
    return beautiful_day
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    i := int32(iTemp)

    jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    j := int32(jTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := beautifulDays(i, j, k)

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

