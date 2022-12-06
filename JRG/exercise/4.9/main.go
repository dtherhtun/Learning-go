package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)

	var num int
	fmt.Print("enter number> ")
	fmt.Scanln(&num)

	fmt.Printf("rand %d number\n", gen.Intn(num))
}
