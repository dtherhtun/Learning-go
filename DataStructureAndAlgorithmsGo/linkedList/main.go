package main

import (
	"container/list"
	"fmt"
)

func main() {
	xlist := list.New()
	xlist.PushBack(`lu lu`)
	xlist.PushBack(`su su`)
	xlist.PushBack(`nu nu`)

	for element := xlist.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	ylist := list.New()
	ylist.PushBack(`yu yu`)
	ylist.PushBack(`ju ju`)

	xlist.PushFrontList(ylist)
	fmt.Println("After added")

	for element := xlist.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	fmt.Println("After removed")

	for element := xlist.Front(); element != nil; element = element.Next() {
		xlist.Remove(element)
	}
	for element := xlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
