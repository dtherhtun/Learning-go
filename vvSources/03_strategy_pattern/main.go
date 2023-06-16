package main

import (
	"flag"
	"fmt"
)

type printer struct {
	strategy outputStrategy
}

func (rcv *printer) setStrategy(s outputStrategy) {
	rcv.strategy = s
}

func (rcv *printer) print(input string) {
	output := rcv.strategy.createOutput(input)
	fmt.Println(output)
}

func main() {
	input := flag.String("i", "", "the input")
	strat := flag.String("s", "", "the strategy")

	flag.Parse()

	p := printer{}

	switch {
	case *strat == "string":
		p.setStrategy(stringStrategy{})
	case *strat == "byte":
		p.setStrategy(byteStrategy{})
	case *strat == "rune":
		p.setStrategy(runeStrategy{})
	default:
		fmt.Println("no strategy specified")
	}

	p.print(*input)
}
