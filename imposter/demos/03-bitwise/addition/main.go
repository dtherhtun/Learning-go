package main

import "fmt"

func main() {
	fmt.Println(add(27, 15))
}

func add(x, y int) int {
	var shouldCarry int = 1
	for shouldCarry != 0 {
		shouldCarry = x & y
		fmt.Printf("shouldCarry(%d) = x(%d) & y(%d)\n", shouldCarry, x, y)
		x = x ^ y
		fmt.Println("x = x ^ y ->", x)
		y = shouldCarry << 1
		fmt.Println("y = shouldCarry << 1 ->", y)
	}
	return x
}
