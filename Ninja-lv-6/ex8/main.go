package main

import "fmt"

func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	a := foo()
	fmt.Println(a())
}
