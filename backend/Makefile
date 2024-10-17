BINDIR := $(shell pwd)/_bin

.PHONY: install
install:
	GOBIN=$(BINDIR) go install github.com/99designs/gqlgen@$(shell go list -m -f "{{.Version}}" github.com/99designs/gqlgen)
	GOBIN=$(BINDIR) go install github.com/google/wire/cmd/wire@$(shell go list -m -f "{{.Version}}" github.com/google/wire)
	GOBIN=$(BINDIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@$(shell go list -m -f "{{.Version}}" github.com/sqlc-dev/sqlc)

.PHONY: generate
generate:
	$(BINDIR)/sqlc generate
	$(BINDIR)/gqlgen generate
	$(BINDIR)/wire gen ./core

.PHOHY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -v -o $(BINDIR)/app
