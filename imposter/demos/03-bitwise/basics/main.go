package main

import (
	"fmt"
)

func main() {
	a := 0b11011
	b := 0b01111
	fmt.Println(a ^ b)
	fmt.Println(a & b)
}

