package main

import (
	"fmt"
)

func main() {
	str := "abbba"
	isPalindrome(str)
	if isPalindrome(str) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func isPalindrome(str string) bool {
	result := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			result = false
			return result
		}
	}
	return result
}
