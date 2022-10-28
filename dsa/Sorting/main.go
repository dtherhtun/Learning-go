package main

import "fmt"

func main() {
	itemOne := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	itemTwo := []int{3, 4, 5, 2, 1}
	itemThree := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}

	fmt.Println(selectionSort(itemOne))
	fmt.Println(bubbleSort(itemTwo))
	fmt.Println(selectionSortRange(itemThree))
}

func selectionSort(items []int) []int {
	n := len(items)

	for i := 0; i < n; i++ {
		min := i
		for j := min; j < n; j++ {
			if items[j] < items[min] {
				items[min], items[j] = items[j], items[min]
			}
		}
	}

	return items
}

func selectionSortRange(items []int) []int {

	for i, _ := range items {
		min := i
		for j, _ := range items {
			if items[j] > items[min] {
				items[j], items[min] = items[min], items[j]
			}
		}
	}

	return items
}

func bubbleSort(items []int) []int {

	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	return items
}
