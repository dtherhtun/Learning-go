package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	w io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	return c.w.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
