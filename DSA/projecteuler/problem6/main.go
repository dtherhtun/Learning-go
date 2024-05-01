package main

import "fmt"

func power(base, pow int) int {
	result := 1
	for pow > 0 {
		result = result * base
		pow--
	}
	return result
}

func sumSquareDiff(n int) int {
	a := 0
	b := 0
	i := 1
	for i <= n {
		a = a + power(i, 2)
		b = b + i
		i++
	}

	return power(b, 2) - a
}
func efficientSumSquareDiff(n int) int {
	sum := n * (n + 1) / 2
	sumSq := (2*n + 1) * (n + 1) * n / 6
	fmt.Println(sum, " ", sumSq)
	return power(sum, 2) - sumSq
}
func main() {
	fmt.Println(sumSquareDiff(100))
	fmt.Println(efficientSumSquareDiff(20))
}
