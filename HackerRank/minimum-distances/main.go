//hackerrank.com/challenges/minimum-distances/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readInts("testcase.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := minimumDistances(data)
	fmt.Println("resoult->", result)
}

func readInts(path string) ([]int32, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int32
	var x string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x = scanner.Text()
		if err != nil {
			return nil, err
		}
	}
	sr := strings.Split(x, " ")
	for i := 0; i < len(sr); i++ {
		temp, _ := strconv.Atoi(sr[i])
		data = append(data, int32(temp))
	}
	return data, scanner.Err()
}

func minimumDistances(a []int32) int32 {
	// Write your code here
	m := make(map[int32]int32)
	for _, v := range a {
		m[v]++
	}
	var nd []int32
	var d []int32
	for k, v := range m {
		if v == 2 {
			fmt.Println(v, k)
			for i, j := range a {
				if j == k {
					nd = append(nd, int32(i))
				}
			}
		}
	}
	if len(nd) == 0 {
		return -1
	}
	for i := 1; i <= len(nd); i = i + 2 {
		temp := nd[i] - nd[i-1]
		d = append(d, temp)
	}
	min := d[0]
	for _, v := range d {
		if min > v {
			min = v
		}
	}
	return min
}
