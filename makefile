.PHONY: build
build:
	go get ./...
	go build cmd/daybookr/daybookr.go