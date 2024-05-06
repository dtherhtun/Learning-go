package main

import "fmt"

func pythagoreanTriplet(n int) int {
	for a := 1; a < (n-3)/3; a++ {
		for b := a + 1; b < (n-1)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				product := a * b * c
				return product
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(pythagoreanTriplet(1000))
}
