package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"chocolate", "martini", "run and coke"},
	}
	p2 := person{
		first:      "Miss",
		last:       "Moneypenny",
		favFlavors: []string{"strawberry", "vanilla", "Capuccino"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for k, v := range p1.favFlavors {
		fmt.Println(k, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for k, v := range p2.favFlavors {
		fmt.Println(k, v)
	}
}
