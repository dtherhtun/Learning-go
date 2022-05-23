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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var Sum []int64
	var first, second, third, fourth, fifth int64
	for _, v := range arr[1:len(arr)] {
		first += int64(v)
	}
	for _, v := range append(arr[2:len(arr)], arr[0]) {
		second += int64(v)
	}
	for _, v := range append(arr[3:len(arr)], arr[:2]...) {
		third += int64(v)
	}
	for _, v := range append(arr[4:len(arr)], arr[:3]...) {
		fourth += int64(v)
	}
	for _, v := range arr[:len(arr)-1] {
		fifth += int64(v)
	}
	Sum = append(Sum, first, second, third, fourth, fifth)
	max := Sum[0]
	min := Sum[0]
	for _, v := range Sum {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	fmt.Println(min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
