package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("my name is ", p.first, p.last, " and age is ", p.age)
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	p1.speak()
}
