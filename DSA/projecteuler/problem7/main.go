package main

import "fmt"

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

func thPrime(n int) int {

	i := 1
	count := 0
	for count <= n {

		if isPrime(i) {
			count++

		}
		if count == n {
			return i
		}
		i++
	}

	return 0

}

func main() {
	fmt.Println(thPrime(10001)) // 104743
}
