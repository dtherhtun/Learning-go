package main

import (
	"fmt"
	"math/rand"
)

func main() {
	one := []string{"pig", "thunder", "spin", "coin", "coin-bag", "shield", "horn"}

	var randResult []int

	for i := 0; i < 3; i++ {
		a := rand.Perm(len(one))
		randResult = append(randResult, a[len(one)/2])
	}
	fmt.Println(one[randResult[0]], one[randResult[1]], one[randResult[2]])
}
