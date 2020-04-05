build: deps generate
	go build ./cmd/daybookr

deps:
	go get ./...

generate:
	go generate ./cmd/daybookr

.PHONY: build generate deps