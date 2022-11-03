package main

import (
	"fmt"
)

func binarySearch(arr []int, first, last, key int) int {
	if last >= first {
		mid := (last + first) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[mid] >= key {
			return binarySearch(arr, first, mid-1, key)
		} else {
			return binarySearch(arr, mid+1, last, key)
		}
	}
	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	key := 70
	last := len(items) - 1
	result := binarySearch(items, 0, last, key)
	if result == -1 {
		fmt.Println("not found")
	} else {
		fmt.Println("Found at position - ", result)
	}
}
