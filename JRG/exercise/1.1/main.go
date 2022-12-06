package main

import "fmt"

func main() {
	ctr := 0
	for ctr < 10 {
		fmt.Println("ctr: ", ctr)
		ctr += 1
	}
}

