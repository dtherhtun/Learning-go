package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "DE THAR"
	//(*p).first = "DE THAR"  // same
}

func main() {
	p1 := person{"DTher", "HTUN", 27}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
