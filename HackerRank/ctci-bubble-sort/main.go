// https://www.hackerrank.com/challenges/ctci-bubble-sort
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
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(a []int32) {
    // Write your code here
    var count int
    n := len(a)
    for i := 0; i < n ; i++ {
        for j := 0; j < n -1 ; j++ {
            if a[j] > a[j+1] {
                a[j], a[j+1] = a[j+1], a[j]
                count++
            }
        }
    }
    fmt.Printf("Array is sorted in %d swaps.\n", count)
    fmt.Println("First Element:", a[0])
    fmt.Println("Last Element:", a[n-1])

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    countSwaps(a)
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

