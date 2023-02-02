package main

import "fmt"

func main() {
	var headerWords uint8 = 5
	headerLen := headerWords * 32 / 8
	b := make([]byte, headerLen)

	s := headerWords << 1
	fmt.Printf("%08b\n", headerWords)
	fmt.Printf("%08b\n", s)
	b[13] = b[13] | s

	fmt.Printf("%08b\n", b[13])
}
