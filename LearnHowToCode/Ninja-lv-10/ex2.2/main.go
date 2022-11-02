package main

import (
	"fmt"
)

func main() {
	cr := make(<-chan int)
	c := make(chan int)

	go func() {
		cr = c
		c <- 42
	}()
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
