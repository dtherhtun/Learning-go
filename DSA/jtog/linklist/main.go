package main

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
	fmt.Println("-----------------")
}

func (l *LinkedList[T]) Head() T {
	return l.head.data
}

func (l *LinkedList[T]) Tail() T {
	return l.tail.data
}

func (l *LinkedList[T]) Length() int {
	return l.size
}

func (l *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{
		data: data,
		next: nil,
	}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	}
	l.tail.next = newNode
	l.tail = newNode
	l.tail.next = nil
	l.size++
}

func (l *LinkedList[T]) RemoveLast() {
	if l.size == 0 {
		fmt.Println("can't remove from empty link list")
		return
	}

	pre, temp := l.head, l.head

	for temp.next != nil {
		pre, temp = temp, temp.next
	}

	l.tail = pre
	l.tail.next = nil
	l.size--
	if l.size == 0 {
		l.head = nil
		l.tail = nil
	}
}

func (l *LinkedList[T]) Prepend(data T) {
	if l.size == 0 {
		l.Append(data)
		return
	}
	newNode := &Node[T]{
		data: data,
		next: l.head,
	}
	l.head = newNode
	l.size++
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.size == 0 {
		return
	}
	temp := l.head
	l.head = l.head.next
	temp.next = nil
	l.size--
	if l.size == 0 {
		l.tail = nil
	}
}

func (l *LinkedList[T]) Get(i int) T {
	var empty T
	if l.size == 0 || i > l.size {
		return empty
	}
	temp := l.head
	for j := 0; j < l.size; j++ {
		if j == i {
			return temp.data
		}
		temp = temp.next
	}
	return empty
}

func main() {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	ll.Print()
	ll.RemoveFirst()
	ll.Prepend(0)

	//fmt.Println("Head ", ll.Head())
	//fmt.Println("Tail ", ll.Tail())
	ll.Print()
	fmt.Println(ll.size)
	fmt.Println(ll.Get(4))
}

//go:generate go build -gcflags "-m" main.go
