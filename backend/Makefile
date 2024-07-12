BINDIR := ./_bin

.PHONY: generate
generate:
	sqlc generate
	gqlgen generate
	wire gen ./core

.PHOHY: test
test:
	go test -v ./...


.PHONY: build
build:
	go build -v -o $(BINDIR)/app
