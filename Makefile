all: fmt vet lint test

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint $(shell glide nv)

test:
	go test ./...
