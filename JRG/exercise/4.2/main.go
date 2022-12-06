package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Enter float number> ")
	fmt.Scanln(&num)

	fmt.Println(math.Trunc(num))
}
