package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dtherhtun/Learning-go/concurrent-programming/barrier"
)

const matrixSize = 1200

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, barrier *barrier.Barrier) {
	for {
		barrier.Wait()
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func main() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	br := barrier.NewBarrier(matrixSize + 1)
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, br)
	}
	start := time.Now()
	for i := 0; i < 4; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		br.Wait()
		br.Wait()
	}
	fmt.Println("Complete in", time.Since(start))
}
