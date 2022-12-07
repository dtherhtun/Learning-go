/*
Just do simple

	exercise
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// a person name
	var name string = "DTher"
	// current date time
	var date time.Time = time.Now()
	fmt.Println(name, " ", date)
}

