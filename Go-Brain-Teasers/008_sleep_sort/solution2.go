package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for _, n := range []int{3, 1, 2} {
		n := n 				// n now is a new variable that lives only in the for loop scope and more importantly in the goroutine closure. It "shadows" the outer n in line 12.
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n)
		}()
	}
	wg.Wait()
	fmt.Println()
}
