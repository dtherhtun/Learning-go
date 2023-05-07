package main

import (
	"fmt"

	"github.com/dtherhtun/Learning-go/concurrent-programming/semaphore"
)

func main() {
	smp := semaphore.NewSemaphore(0)
	for i := 0; i < 50000; i++ {
		go doWork(smp)
		fmt.Println("Waiting for child goroutine")
		smp.Acquire()
		fmt.Println("Child goroutine finished")
	}
}

func doWork(smp *semaphore.Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	smp.Release()
}
