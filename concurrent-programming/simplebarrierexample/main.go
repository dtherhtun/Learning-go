package main

import (
	"fmt"
	"time"

	"github.com/dtherhtun/Learning-go/concurrent-programming/barrier"
)

func WorkAndWait(name string, WorkStart int, br *barrier.Barrier) {
	start := time.Now()
	fmt.Println(time.Since(start), name, "is running")
	time.Sleep(time.Duration(WorkStart) * time.Second)
	fmt.Println(time.Since(start), name, "is waiting on barrier")
	br.Wait()
}

func main() {
	br := barrier.NewBarrier(2)
	go WorkAndWait("Red", 4, br)
	go WorkAndWait("Blue", 10, br)
	time.Sleep(time.Duration(100) * time.Second)
}
