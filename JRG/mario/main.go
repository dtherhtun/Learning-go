package main

import "fmt"

func main() {
	input := 1
	for i := 1; i <= input; i++ {
		for j := i; j < input; j++ {
			fmt.Print("j")
		}
		for k := 0; k < i; k++ {
			fmt.Print("k")
		}

		fmt.Print("  ")
		for k := 0; k < i; k++ {
			fmt.Print("k")
		}
		fmt.Println()
	}
}
