package main

import (
	"errors"
	"fmt"
)

var (
	ErrorValidation = errors.New("validation failed")
	ErrSignal       = errors.New("Receivd signal")
)

type stepErr struct {
	step  string
	msg   string
	cause error
}

func (s *stepErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}

func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}
	return t.step == s.step
}

func (s *stepErr) UnWarp() error {
	return s.cause
}
