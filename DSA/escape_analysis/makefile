# https://appliedgo.com/blog/how-to-do-escape-analysis
SHELL := /bin/bash

ans01:
	go tool compile -m analysis01/main.go
	rm -rf main.o

ans02:
	go tool compile -m analysis02/main.go
	rm -rf main.o

ans01a:
	go run -gcflags="-m -l" analysis01/main.go

ans02a:
	go run -gcflags="-m -l" analysis02/main.go
