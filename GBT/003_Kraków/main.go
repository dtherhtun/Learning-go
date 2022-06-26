// Go strings are UTF-8 encoded, this means that a character (called rune) can be from one to four bytes long.
// In this example the rune ó is taking 2 bytes

package main

import (
	"fmt"
)

func main() {
	city := "Kraków"
	fmt.Println(len(city))
}
