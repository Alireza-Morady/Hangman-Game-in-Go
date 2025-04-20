.DEFAULT_GOAL := execute
.PHONY: fmt vet build clean execute

BINARY_NAME := HangmanProject

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o $(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)

execute: build
	clear
	@sleep 0.05
	@./$(BINARY_NAME)