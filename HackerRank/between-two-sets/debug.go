// https://go.dev/play/p/NbirGGcbeT0

// A is A factor of x
// B is A multiple of x
package main

import "fmt"

func main() {
	a := []int32{2, 4}
	b := []int32{16, 32, 96}
	maxA := a[0]
	minB := b[0]
	var isFactorMultiple bool
	var count int32
	for _, v := range a {
		if maxA < v {
			maxA = v
		}
	}
	for _, v := range b {
		if minB > v {
			minB = v
		}
	}
	fmt.Println("maxB->", maxA, " minA->", minB+1)
	for i := maxA; i < minB+1; i++ {
		isFactorMultiple = true
		fmt.Println("value of i->", i)
		fmt.Println("Before condition->", isFactorMultiple)
		for _, v := range a {
			fmt.Println(i, "%", v, "=", i%v, "-> inside a")
			if i%v != 0 {
				isFactorMultiple = false
				fmt.Println(isFactorMultiple)
				break
			}
		}
		for _, v := range b {
			fmt.Println(i, "%", v, "=", i%v, "-> inside b")
			if v%i != 0 {
				isFactorMultiple = false
				fmt.Println(isFactorMultiple)
				break
			}
		}
		fmt.Println("After condition->", isFactorMultiple)
		if isFactorMultiple == true {
			count++
		}
	}
	fmt.Println("Total count->", count)
}
