package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for _, n := range []int{3, 1, 2} {
		fmt.Println("before go routine n ->", n)
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("inside go routine n ->", n)
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n)
		}()
		fmt.Println("after go routine n ->", n)
	}
	wg.Wait()
	fmt.Println()
}
