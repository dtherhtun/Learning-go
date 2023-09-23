package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	// TODO
}


func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}

// Output should be all capital latter
