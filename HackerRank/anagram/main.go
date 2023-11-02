package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := "anagram"
	b := "nagaram"

	if isAnagram(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	if isAnagramByte(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	if isAnagramCounter(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	if isAnagramMap(a, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isAnagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	if len(a) != len(b) {
		return false
	}
	aSlice := []rune(a)
	sort.Slice(aSlice, func(i, j int) bool {
		return aSlice[i] < aSlice[j]
	})
	bSlice := []rune(b)
	sort.Slice(bSlice, func(i, j int) bool {
		return bSlice[i] < bSlice[j]
	})

	for i := 0; i < len(aSlice); i++ {
		if aSlice[i] != bSlice[i] {
			return false
		}
	}
	return true
}

func isAnagramByte(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aSlice := []byte(a)
	sort.Slice(aSlice, func(i, j int) bool {
		return aSlice[i] < aSlice[j]
	})
	bSlice := []byte(b)
	sort.Slice(bSlice, func(i, j int) bool {
		return bSlice[i] < bSlice[j]
	})

	return bytes.Equal(aSlice, bSlice)
}

func isAnagramCounter(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var alphabetCounter [26]int

	for i := 0; i < len(a); i++ {
		alphabetCounter[a[i]-'a']++
		alphabetCounter[b[i]-'a']--
	}

	for i := 0; i < len(alphabetCounter); i++ {
		if alphabetCounter[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagramMap(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aSlice := make(map[rune]int)
	bSlice := make(map[rune]int)

	for _, i := range a {
		aSlice[i]++
	}
	for _, i := range b {
		bSlice[i]++
	}

	for i, j := range aSlice {

		if bCount, ok := bSlice[i]; !ok || j != bCount {
			return false
		}
	}

	return true
}
