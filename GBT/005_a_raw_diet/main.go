// One of the most common uses for raw strings is to create multi-line strings. 

package main

import "fmt"

func main() {
	s := `a\tb`
	fmt.Println(s)
	fmt.Println("\u2122")
	fmt.Println("\u2E2E\u1F441")
}
