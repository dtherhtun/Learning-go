package main

import "errors"

var (
	ErrNotNumber        = errors.New("data is not numeric")
	ErrInvaildColumn    = errors.New("invalid column number")
	ErrNoFiles          = errors.New("no input files")
	ErrInvalidOperation = errors.New("invalid operation")
)
