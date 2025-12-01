SHELL := /bin/bash

.PHONY: test lint fmt tidy gen run submit test-aoc test-quiet

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	gofumpt -w .
	goimports -w .

tidy:
	go mod tidy

gen:
	go run . -g -n

run:
	@read -p "Year: " Y; \
	read -p "Day: " D; \
	read -p "Part: " P; \
	go run . -y $$Y -d $$D -p $$P

submit:
	@read -p "Year: " Y; \
	read -p "Day: " D; \
	read -p "Part: " P; \
	go run . -s -y $$Y -d $$D -p $$P

test-aoc:
	go run . -t

test-quiet:
	go run . -q
