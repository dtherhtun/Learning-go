package main

import "fmt"

func largestPrimeFactor(n int) int {
	i := 2
	k := n
	result := 0
	for i <= k {
		for isPrime(i) && k%i == 0 {
			k = k / i
			result = i
		}
		i++
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
