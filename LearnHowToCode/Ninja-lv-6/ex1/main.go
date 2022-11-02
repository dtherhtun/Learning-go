package main

import "fmt"

func foo() int {
	return 4
}

func bar() (int, string) {
	return 46, "bar"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
