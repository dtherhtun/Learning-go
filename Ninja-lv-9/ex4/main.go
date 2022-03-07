package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Start of CPU: ", runtime.NumCPU())
	fmt.Println("Start of Goroutine: ", runtime.NumGoroutine())
	var counter int
	const gs int = 100

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			i := counter
			i++
			counter = i
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Mid of Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
	fmt.Println("End of Goroutine: ", runtime.NumGoroutine())
}
