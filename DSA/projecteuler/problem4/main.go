package main

import "fmt"

func isPalindromicNumber(n int) bool {
	original := n
	reverse := 0

	for 0 < n {
		lastDigit := n % 10
		reverse = reverse*10 + lastDigit
		n /= 10
	}
	return reverse == original
}

func digitGenerator(n int) (int, int) {
	var minDigit int = 1
	var maxDigit int = 0
	for n > 0 {
		minDigit = minDigit * 10
		maxDigit = maxDigit*10 + 9
		n--
	}
	return minDigit / 10, maxDigit
}

func main() {
	fmt.Println(largestPalindromicNumber(3))
}

func largestPalindromicNumber(digit int) int {
	minDigit, maxDigit := digitGenerator(digit)
	maxPalindromic := 0

	for i := maxDigit; i >= minDigit; i-- {
		for j := i; j >= minDigit; j-- {
			product := i * j
			if product <= maxPalindromic {
				break
			}
			if isPalindromicNumber(i * j) {
				maxPalindromic = product
			}
		}
	}

	return maxPalindromic
}
