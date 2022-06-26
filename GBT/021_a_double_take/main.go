// init is the final stage in package initialization after imported
// modules are initialized and package-level variables are initialized.

package main

import (
	"fmt"
)

func init() {
	fmt.Printf("A ")
}

func init() {
	fmt.Print("B ")
}

func main() {
	fmt.Println()
}
