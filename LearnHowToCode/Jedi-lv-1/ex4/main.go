package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y := int(x)
	y = 42
	fmt.Println(y)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
}
