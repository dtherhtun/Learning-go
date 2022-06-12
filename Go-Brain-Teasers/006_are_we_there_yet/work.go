package main

import (
	"fmt"
	"time"
)

func main() {
	var timeout time.Duration = 3
	fmt.Printf("before ")
	time.Sleep(timeout * time.Millisecond)
	fmt.Println("after ")
}
