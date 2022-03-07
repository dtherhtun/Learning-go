package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("start CPUs:\t", runtime.NumCPU())
	fmt.Println("start Goroutines:\t", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("foo")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("bar")
		}
		wg.Done()
	}()

	fmt.Println("mid CPUs:\t", runtime.NumCPU())
	fmt.Println("mid Goroutines:\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPUs:\t", runtime.NumCPU())
	fmt.Println("End Goroutines:\t", runtime.NumGoroutine())

}
