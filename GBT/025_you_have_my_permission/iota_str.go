package main

import (
	"fmt"
)

type FilePerm uint64

const (
	Read FilePerm = 1 << iota
	Write
	Execute
)

func (p FilePerm) String() string {
	switch p {
	case Read:
		return "read"
	case Write:
		return "write"
	case Execute:
		return "execute"
	}
	return fmt.Sprintf("unknown FilePerm: %d", p) // don't use %s here
}

func main() {
	fmt.Println(Execute)
	fmt.Printf("%d\n", Execute)
	fmt.Println(Write)
	fmt.Printf("%d\n", Write)
}
