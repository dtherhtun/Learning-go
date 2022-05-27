package main

import (
	"fmt"
	"strings"
)

func main() {
	var max int32
	ls := "abcz"
	ar := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	m := make(map[string]int32)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	sa := strings.Split(alphabet, "")
	for i := 0; i < len(ar); i++ {
		m[sa[i]] = ar[i]
	}
	la := strings.Split(ls, "")
	for _, v := range la {
		if max < m[v] {
			max = m[v]
		}
	}
	fmt.Println(max)
	fmt.Println(m)
}
