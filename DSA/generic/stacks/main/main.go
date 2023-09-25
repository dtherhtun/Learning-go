package main

import (
	"fmt"

	"github.com/dtherhtun/Learning-go/DSA/generic/stacks"
)

func main() {
	a := &stacks.Stack[string]{}

	a.Push("hello")
	a.Push("world")
	a.Push("gopher")

	for !a.IsEmpty() {
		b, err := a.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*b)
	}

	c := &stacks.Stack[int]{}

	c.Push(3)
	c.Push(2)
	c.Push(1)
	for !c.IsEmpty() {
		b, err := c.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*b)
	}
}
