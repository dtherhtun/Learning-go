package main

import "fmt"

func main() {
	ctr := 0
	for ctr < 20 {
		fmt.Println("ctr: ", ctr)
		ctr += 2
	}
}
