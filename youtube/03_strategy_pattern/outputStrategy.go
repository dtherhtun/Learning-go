package main

import "fmt"

type outputStrategy interface {
	createOutput(s string) string
}

type stringStrategy struct{}

func (rcv stringStrategy) createOutput(s string) string {
	return s
}

type byteStrategy struct{}

func (rcv byteStrategy) createOutput(s string) string {
	return fmt.Sprintf("%v", []byte(s))
}

type runeStrategy struct{}

func (rcv runeStrategy) createOutput(s string) string {
	return fmt.Sprintf("%v", []rune(s))
}
