package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func stingy(money *int, mu *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mu.Lock()
		*money += 10
		mu.Unlock()
	}
	fmt.Println("Stingy Done")
}

func spendy(money *int, mu *sync.Mutex) {
	for i := 0; i < 200000; i++ {
		mu.Lock()
		for *money < 50 {
			mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		mu.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mu := sync.Mutex{}
	go stingy(&money, &mu)
	go spendy(&money, &mu)
	time.Sleep(2 * time.Second)
	mu.Lock()
	fmt.Println("Money in bank account:", money)
	mu.Unlock()
}
