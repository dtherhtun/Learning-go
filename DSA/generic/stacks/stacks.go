package stacks

import "errors"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("can not pop from empty stack")
	}

	// the last element is the one to read
	top := s.elements[len(s.elements)-1]
	// drop read element
	s.elements = s.elements[:len(s.elements)-1]
	return &top, nil
}
