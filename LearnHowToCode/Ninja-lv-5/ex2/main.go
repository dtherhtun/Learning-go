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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, j := range v.favFlavors {
			fmt.Println(i, j)
		}
		fmt.Println("--------------")
	}
}
