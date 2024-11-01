package main

import (
	"fmt"
)

func main() {
	itemOne := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	itemTwo := []int{3, 4, 5, 2, 1}
	itemThree := []int{-1, 10, 3, 4, 5, 2, 1, 7, 8, -1, -3}
	itemfour := []int{22, 10, 6, 2, 1, 5, 8, 3, 4, 7, 9, 100, 30, 70, -1}
	itemFive := []int{7, 2, 1, 6, 8, 5, 3, 4}

	fmt.Println(selectionSort(itemOne))
	fmt.Println(bubbleSort(itemTwo))
	fmt.Println(selectionSortRange(itemThree))
	fmt.Println(mergeSort(itemfour))
	fmt.Println("quickSort->", quickSort(itemFive, 0, len(itemFive)-1))
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

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	left := items[:len(items)/2]
	right := items[len(items)/2:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var results []int
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			results = append(results, left[i])
			i++
		} else {
			results = append(results, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		results = append(results, left[i])
	}

	for ; j < len(right); j++ {
		results = append(results, right[j])
	}

	return results
}

func partation(items []int, low, high int) ([]int, int) {
	pivot := items[high]
	i := low
	for j := low; j < high; j++ {
		if items[j] < pivot {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[high] = items[high], items[i]

	return items, i
}

func quickSort(items []int, low, high int) []int {
	if low < high {
		var p int
		items, p = partation(items, low, high)
		items = quickSort(items, low, p-1)
		items = quickSort(items, p+1, high)
	}
	return items
}
