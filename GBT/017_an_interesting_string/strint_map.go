// When you convert from byte slice to string,
// Go will copy the byte slice - which does a memory allocation.
// In maps, where you can’t use a []byte as a key, there is a compiler optimization:
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"hello": 3,
	}
	key := []byte{'h', 'e', 'l', 'l', 'o'}
	val := m[string(key)] // no memory allocation
	fmt.Println(val)
}
