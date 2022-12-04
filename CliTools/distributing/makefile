SHELL := /bin/bash

build-alpine:
	docker build -t dther/pomo:alpine -f containers/Dockerfile.builder .

build:
	docker build -t dther/pomo:latest -f containers/Dockerfile.scratch .

run:
	docker run --rm -it dther/pomo:latest

push-alpine:
	docker push dther/pomo:alpine

push:
	docker push dther/pomo:latest
