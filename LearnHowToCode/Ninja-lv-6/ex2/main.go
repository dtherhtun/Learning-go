package main

import "fmt"

func foo(x ...int) int {
	var y int
	for _, v := range x {
		y += v
	}
	return y
}

func bar(z []int) int {
	w := 0
	for _, v := range z {
		w += v
	}

	return w
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	total := foo(nums...)
	fmt.Println(total)

	num2 := []int{2, 4, 6, 8, 9}
	bbar := bar(num2)
	fmt.Println(bbar)
}
