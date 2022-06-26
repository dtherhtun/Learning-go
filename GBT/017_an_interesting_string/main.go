// The string type supports type conversion from int,
// it’ll treat this integer as a rune.
// The rune 169 is the copyright sign (©)
package main

import (
	"fmt"
)

func main() {
	i := 169
	s := string(i)
	fmt.Println(s)
}
