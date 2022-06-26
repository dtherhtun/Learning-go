// Go strings are UTF-8 encoded. When you access a string with [] or with len,
// Go will access the underlying []byte. byte is a type in Go that’s aliased to uint8.
// When you iterate over a string in Go, you will get the Unicode character -
// called rune which is an alias to int32. This is due to the fact the characters in UTF-8 can be up to 4 bytes.

package main

import (
	"fmt"
)

func main() {
	msg := "π = 3.14159265358..."
	fmt.Printf("%T ", msg[0])
	for _, c := range msg {
		fmt.Printf("%T\n", c)
		break
	}
}
