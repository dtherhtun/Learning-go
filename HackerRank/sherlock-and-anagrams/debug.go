package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "abba"
	var result int
	for i := 1; i < len(s); i++ {
		fmt.Println("outer-i=", i)
		m := make(map[string]int)
		fmt.Println("m=", m)
		for j := 0; j < len(s)-i+1; j++ {
			fmt.Println("j=", j)
			fmt.Println("s=", s)
			sa := strings.Split(s, "")
			fmt.Println("split=", sa)
			b := sa[j : j+i]
			fmt.Println("before-sort=", b, "j:j+i=", j, ":", j+i)
			sort.Strings(b)
			fmt.Println("after-sort=", b)
			sub := strings.Join(b, "")
			fmt.Println("substring=", sub)

			if _, ok := m[sub]; ok {
				m[sub] += 1
				fmt.Println("m-exist=", m[sub])
			} else {
				fmt.Println("m-not-exist=", m[sub])
				m[sub] = 1
			}
			result += m[sub] - 1
			fmt.Println("m-inner=", m, " m-sub=", m[sub])
			fmt.Println("inner-result=", result)

		}

	}
	fmt.Println("result=", result)
}
