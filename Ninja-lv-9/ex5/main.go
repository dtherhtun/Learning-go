package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Start of CPU: ", runtime.NumCPU())
	fmt.Println("Start of Goroutine: ", runtime.NumGoroutine())
	var counter int64
	const gs int = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Mid of Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
	fmt.Println("End of Goroutine: ", runtime.NumGoroutine())
}
