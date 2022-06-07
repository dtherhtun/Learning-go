package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var x uint32
		fmt.Scan(&x)
		fmt.Println(^x)
	}
}

/*
input -
3
0
802743475
35601423
----------
output - 
4294967295
3492223820
4259365872
*/
