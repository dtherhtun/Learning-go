package doublyLinkedList

import (
	"errors"
	"fmt"
)

const showExpectedResult = false

const showHints = false

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (l *DoublyLinkedList[T]) Add(index int, v T) error {
	currentSize := l.size
	if index > currentSize {
		return errors.New("index exceeds list size")
	}
	newElement := &Node[T]{
		value: v,
	}
	l.size = currentSize + 1

	// list is empty
	if l.head == nil {
		l.head, l.tail = newElement, newElement
		return nil
	}

	// change head
	if index == 0 {
		newElement.next = l.head
		l.head.prev, l.head = newElement, newElement
		return nil
	}

	// change tail
	if index == currentSize {
		newElement.prev = l.tail
		l.tail.next, l.tail = newElement, newElement
		return nil
	}

	// find element at index
	current := l.head
	for i := 1; i < index; i++ {
		current = current.next
	}

	newElement.prev = current
	newElement.next = current.next
	current.next.prev, current.next = newElement, newElement

	return nil
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}
