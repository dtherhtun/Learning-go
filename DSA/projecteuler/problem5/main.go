package main

import "fmt"

func main() {
	fmt.Println(leastCommonMultiple(20))
	fmt.Println(simpleSmallestMultiple(1, 20))
}

func leastCommonMultiple(n int) int {
	lcm := 1
	for i := 2; i <= n; i++ {
		lcm = calculateLCMofTwo(lcm, i)
	}
	return lcm
}

func calculateLCMofTwo(a, b int) int {
	return a * b / greatestCommonMultiple(a, b)
}

func greatestCommonMultiple(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simpleSmallestMultiple(a uint64, b uint64) uint64 {
	var i uint64

	i = b * a

	for {
		var deviseCount uint64
		for j := a; j <= b; j++ {
			if i%j == 0 {
				deviseCount++
			}
		}

		if deviseCount > b-a {
			return i
		}

		i += b
	}
}
