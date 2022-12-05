package main

import "fmt"

func main() {
	fmt.Println(a())
}

func a() int {
	a := b()
	return *a
}

func b() *int {
	b := 1
	return &b
}

