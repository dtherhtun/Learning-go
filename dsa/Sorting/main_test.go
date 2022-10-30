package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSelectionSort(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := make([]int, size)
			for i := 0; i < size; i++ {
				testList[i] = rand.Intn(size)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				selectionSort(testList)
			}
		})
	}
}

func BenchmarkSelectionSortRange(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := make([]int, size)
			for i := 0; i < size; i++ {
				testList[i] = rand.Intn(size)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				selectionSortRange(testList)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := make([]int, size)
			for i := 0; i < size; i++ {
				testList[i] = rand.Intn(size)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bubbleSort(testList)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	inputSize := []int{10, 100, 1000, 10000, 100000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			testList := make([]int, size)
			for i := 0; i < size; i++ {
				testList[i] = rand.Intn(size)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				mergeSort(testList)
			}
		})
	}
}
