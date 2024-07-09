BINDIR := ./_bin

.PHONY: generate
generate:
	sqlc generate
	wire gen ./core

.PHOHY: test
test:
	go test -v ./...


.PHONY: build
build:
	go build -v -o $(BINDIR)/app
