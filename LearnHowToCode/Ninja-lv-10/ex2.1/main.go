package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)
	c := make(chan int)

	go func() {
		cs <- 42
	}()
	cs = c
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
