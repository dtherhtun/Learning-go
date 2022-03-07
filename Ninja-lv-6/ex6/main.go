package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello from anonymous function")
	}()
}
