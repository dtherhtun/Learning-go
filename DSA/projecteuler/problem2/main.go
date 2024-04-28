package main

import "fmt"

func main() {
	i := 1
	sum := 0
	prev := 0
	for i < 4000000 {
		prev = prev + i
		i = i + prev
		if prev%2 == 0 {
			sum += prev
		}
		if i%2 == 0 {
			sum += i
		}
	}
	result := sumEvenFibonacci(4000000)
	fmt.Println(sum)
	fmt.Println(result)
}

func sumEvenFibonacci(limit int) int {
	sum := 0
	a, b := 0, 1
	for b <= limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}
