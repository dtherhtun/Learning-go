package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binaryAddition("11011", "1111"))
}

func binaryAddition(x, y string) string {

	var sum string
	var c int = 0

	if len(x) > len(y) {
		y = padStart(y, "0", len(x))
	}
	if len(y) > len(x) {
		x = padStart(x, "0", len(y))
	}

	for i := len(x); i > 0; i-- {
		a, _ := strconv.Atoi(x[i-1 : i])
		b, _ := strconv.Atoi(y[i-1 : i])
		left, right := fullAdder(a, b, c)

		c = left
		sum = strconv.Itoa(right) + sum
	}

	if c == 1 {
		return strconv.Itoa(c) + sum
	}
	return sum
}

func padStart(s string, pad string, plength int) string {

	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}
func and(x, y int) int { return x & y }
func xor(x, y int) int { return x ^ y }
func or(x, y int) int  { return x | y }

func equiv(x, y int) int {

	if x == y {
		return 1
	}
	return 0
}

func halfAdder(x, y int) (int, int) {
	return and(x, y), xor(x, y)
}

func halfNadder(x, y int) (int, int) {
	return or(x, y), equiv(x, y)
}

func fullAdder(x, y, c int) (int, int) {

	if c == 1 {
		return halfNadder(x, y)
	}
	return halfAdder(x, y)
}
