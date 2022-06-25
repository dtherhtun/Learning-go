// hackerrank.com/challenges/sherlock-and-squares

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
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */

func squares(a int32, b int32) int32 {
    // Write your code here
    var count int32
    /*for i := a; i <= b ; i++ {
        k := i-(i/2)
        for j := k; j >= 0 ; j-- {
            fmt.Println(j ,"*", j , "=", i)
            if j * j == i {
                count++
            }
        }
    }*/
    // return int32(math.Floor(math.Sqrt(float64(b))) - math.Ceil(math.Sqrt(float64(a)))) + 1
    for i := int32(1); i <= b; i++ {
        sqr := i * i
        if sqr >= a && sqr <= b {
            count++
        }
        if sqr > b {
            break
        }
    }
    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        aTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        a := int32(aTemp)

        bTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        b := int32(bTemp)

        result := squares(a, b)

        fmt.Fprintf(writer, "%d\n", result)
    }

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

