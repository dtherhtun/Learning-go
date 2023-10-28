package main

import "fmt"

func filter(input []float64, f func(float642 float64) bool) []float64 {
	var result []float64
	for _, value := range input {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func genericFilter[T any](intput []T, f func(T) bool) []T {
	var result []T
	for _, val := range intput {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	res := filter(input, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})
	fmt.Println(res)
	genericRes := genericFilter(input, func(num float64) bool {
		if num >= 10.0 {
			return true
		}
		return false
	})
	fmt.Println(genericRes)

	oddNumbers := genericFilter([]int{4, 6, 5, 2, 20, 1, 7}, func(i int) bool {
		return i%2 == 1
	})
	fmt.Println(oddNumbers)
}
