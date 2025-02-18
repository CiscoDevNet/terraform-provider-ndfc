default: lint build install

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	golangci-lint run

generate:
	go generate

fmt:
	gofmt -s -w -e .

testacc:
	$(shell ./run_accept_tests.sh)

.PHONY: fmt lint build install generate