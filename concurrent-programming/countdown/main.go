package main

import (
	"fmt"
	"time"
)

func main() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}

func countdown(second *int) {
	for *second > 0 {
		time.Sleep(1 * time.Second)
		*second -= 1
	}
}
