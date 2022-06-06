package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "UDDDUDUU"
	steps := 8
	var current_level int32
	var valleys int32
	paths := strings.Split(path, "")

	for i := 0; i < steps; i++ {
		if paths[i] == "U" {
			current_level++
			if current_level == 0 {
				valleys++
			}
		} else if paths[i] == "D" {
			current_level--
		}
	}
	fmt.Println(valleys)
}
