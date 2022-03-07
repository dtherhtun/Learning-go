package main

import "fmt"

func foo(f func(xi []int) int, yi []int) int {
	n := f(yi)
	n++
	return n
}

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		} else if len(xi) == 1 {
			return xi[0]
		} else {
			return xi[0] + xi[len(xi)-1]
		}
	}
	result := foo(g, []int{1, 3, 4})
	fmt.Println(result)
}
