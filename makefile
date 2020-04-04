.PHONY: build
build:
	go get ./...
	go build cmd/daybookr/daybookr.go

.PHONY: generate
generate:
	go run cmd/daybookr/daybookr.go --input test_dir