package main

import "fmt"

func main() {
	// integer
	printNum(10)    // 10 of type int
	printNum(010)   // 8 of type int
	printNum(0x10)  // 16 of typeint
	printNum(0b10)  // 2 of type int
	printNum(1_000) // 1000 of type int

	// Float
	printNum(3.14)
	printNum(.2)
	printNum(1e3)
	printNum(0x1p-2)

	// Complex
	printNum(1i)
	printNum(3 + 7i)
	printNum(1 + 0i)
}

func printNum(n interface{}) {
	fmt.Printf("%v of type %T\n", n, n)
}
