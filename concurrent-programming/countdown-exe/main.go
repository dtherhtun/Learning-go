package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mu *sync.Mutex) {
	mu.Lock()
	remaining := *seconds
	mu.Unlock()
	for remaining > 0 {
		time.Sleep(1 * time.Second)
		mu.Lock()
		*seconds -= 1
		remaining = *seconds
		mu.Unlock()
	}

}

func main() {
	count := 5
	mu := sync.Mutex{}
	go countdown(&count, &mu)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		mu.Lock()
		fmt.Println(count)
		mu.Unlock()
	}
}
