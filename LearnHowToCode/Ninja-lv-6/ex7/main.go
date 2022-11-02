package main

import "fmt"

var g func() = func() {
	fmt.Println("g from outside")
}

func main() {
	a := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}
	a()
	g()
	g = a
	g()
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", g)
	fmt.Println("Done.")
}
