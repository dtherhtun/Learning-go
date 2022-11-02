package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("hello")
}

type human interface {
	speak()
}

func saySometing(h human) {
	h.speak()
}

func main() {
	var p1 person
	p1.name = "dther"

	saySometing(&p1)
	p1.speak()
}
