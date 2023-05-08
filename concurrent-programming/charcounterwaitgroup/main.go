package main

import (
	"fmt"
	"sync"

	"github.com/dtherhtun/Learning-go/concurrent-programming/charcountermutex"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(201)
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go func() {
			charcountermutex.CountLetters(url, frequency, &mutex)
			wg.Done()
		}()
	}
	wg.Wait()
	mutex.Lock()
	fmt.Println(frequency)
	mutex.Unlock()
}
