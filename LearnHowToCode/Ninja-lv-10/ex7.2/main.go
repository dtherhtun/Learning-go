package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := 10
	b := 10
	c := gen(a, b)

	for i := 0; i < a*b; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(a, b int) <-chan int {
	c := make(chan int)

	for i := 0; i < a; i++ {
		go func() {
			for j := 0; j < b; j++ {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}
