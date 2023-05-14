package main

import (
	"fmt"

	"github.com/dtherhtun/Learning-go/concurrent-programming/waitgrpsema"
)

func main() {
	wg := waitgrpsema.NewWaitGrp(4)
	for i := 1; i <= 4; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
	fmt.Println("All completed")
}

func doWork(id int, wg *waitgrpsema.WaitGrp) {
	fmt.Println(id, "Done Working")
	wg.Done()
}
