package main

import (
	"fmt"
)

func sumOfPrimeBelow(n int) int {
	i := 5
	sum := 5
	for i < n {
		if isPrime(i) {
			sum = sum + i
		}
		i = i + 2
		if i <= n && isPrime(i) {
			sum = sum + i
		}
		i += 4
	}
	return sum
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

func sieveOfEratosthenes(limit int) []bool {
	prime := make([]bool, limit)
	for i := 2; i < limit; i++ {
		prime[i] = true
	}
	for p := 2; p*p < limit; p++ {
		if prime[p] {
			for i := p * p; i < limit; i += p {
				prime[i] = false
			}
		}
	}
	return prime
}

func sumOfPrimes(limit int) int {
	prime := sieveOfEratosthenes(limit)
	sum := 0
	for i := 2; i < limit; i++ {
		if prime[i] {
			sum += i
		}
	}
	return sum
}

func main() {
	//fmt.Println(sumOfPrimeBelow(2000000))
	fmt.Println(sumOfPrimes(2000000))
}
