package main

import (
	"fmt"
)

func main() {
	x := "James Bond"
	if x == "D ther" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BOND BOND ", x)
	} else {
		fmt.Println("neither")
	}
}
