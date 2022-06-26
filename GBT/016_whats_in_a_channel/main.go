// The question is - what happens when you receive from a closed channel?
// The answer is that if there are values in the buffer - you will get them,
// otherwise, you will get the zero value for the channel type.
// This is why a gets the value of 2 which was in the buffer, and b gets 0 which is the zero value for int.
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	<-ch
	close(ch)
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}
