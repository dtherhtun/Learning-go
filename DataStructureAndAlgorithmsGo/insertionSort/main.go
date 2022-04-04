package main

import "fmt"

func insertionSort(arr []int) []int {
	for out := 1; out < len(arr); out++ {
		temp := arr[out]
		in := out

		for in > 0 && arr[in-1] >= temp {
			arr[in] = arr[in-1]
			in--
			//fmt.Println(arr)
		}
		arr[in] = temp
		//fmt.Println("\t", arr)
	}
	return arr
}

func main() {
	arr := []int{10, 100, 20, 40, 60, 5, 3, 70}
	result := insertionSort(arr)
	fmt.Println(result)
}
