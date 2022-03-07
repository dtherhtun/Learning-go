package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("foo DEFFER ran")
	}()
	fmt.Println("foo ran")
}

func main() {
	defer foo()
	fmt.Println("main ran")
}
