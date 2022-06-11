// the code will compile since now n has a type.
// Here are some places where nil is used:
// The zero value for map, slice and chan is nil.
// You can’t compare slices or maps using ==, you can only compare them to nil.
// Sending to or receiving from a nil channel will block forever. You can use this to avoid a busy wait.

package main

import (
	"fmt"
)

func main() {
	var n *int = nil
	fmt.Println(n)
}
