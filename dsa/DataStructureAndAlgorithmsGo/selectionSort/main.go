package main

import "fmt"

func selectionSort(arr []int) {
	for out := 0; out < len(arr)-1; out++ {
		min := out
		for in := out + 1; in < len(arr); in++ {
			if arr[in] < arr[min] {
				min = in
			}
			//fmt.Println(arr)
		}
		arr[out], arr[min] = arr[min], arr[out]
		//fmt.Println("\t", arr)
	}
}

// it is dirrerent between -1/+1 and without -1/+1

func main() {
	arr := []int{10, 100, 20, 40, 60, 5, 3, 70}
	selectionSort(arr)
	fmt.Println(arr)
}
