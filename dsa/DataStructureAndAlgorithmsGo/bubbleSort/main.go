package main

import "fmt"

func bubbleSort(arr []int) []int {
	for out := len(arr) - 1; out >= 1; out-- {
		for in := 0; in < out; in++ {
			if arr[in] > arr[in+1] {
				temp := arr[in]
				arr[in] = arr[in+1]
				arr[in+1] = temp
			}
			//fmt.Println(arr)
		}
		//fmt.Println("\t", arr)
	}
	return arr
}

func main() {
	arr := []int{10, 100, 20, 40, 60, 5, 3, 70}
	result := bubbleSort(arr)
	fmt.Println(result)
}
