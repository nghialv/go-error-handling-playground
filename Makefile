.PHONY: build
build:
	go build main.go

.PHONY: deps
deps:
	go mod tidy
