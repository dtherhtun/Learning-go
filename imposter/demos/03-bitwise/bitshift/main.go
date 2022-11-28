package main

import "fmt"

func main() {
	a := 0b1

	// left shift
	for i := 1; i < 10; i++ {
		a = a << 1
		fmt.Printf("%[1]b = %[1]d\n", a)
	}

	// right shift
	for i := 1; i < 10; i++ {
		a = a >> 1
		fmt.Printf("%[1]b = %[1]d\n", a)
	}
}
