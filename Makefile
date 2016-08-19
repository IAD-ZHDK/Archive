PKGS=$(shell glide nv)

all: fmt vet lint

vet:
	go vet $(PKGS)

fmt:
	go fmt $(PKGS)

lint:
	golint $(PKGS)
