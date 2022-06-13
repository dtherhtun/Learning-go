package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}

	fmt.Println(a[:1])
	b := append(a[:1], 10)
	fmt.Println(a[:1])
	fmt.Printf("a=%v, b=%v\n", a, b)
}
