// hackerrank.com/challenges/30-data-types/problem
// input
/*
12
4.0
is the best place to learn and practice coding!
*/
// output
/*
16
8.0
HackerRank is the best place to learn and practice coding!
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var arr []string
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, text)
		} else {
			break
		}
	}
	u, _ := strconv.ParseUint(arr[0], 0, 64)
	f, _ := strconv.ParseFloat(arr[1], 64)

	fmt.Println(i + u)
	fmt.Printf("%.1f\n", d+f)
	fmt.Println(s + arr[2])
}
