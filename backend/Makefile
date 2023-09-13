.PHONY: generate
generate:
	buf generate _schema/proto
	sqlc generate
	wire gen ./core

.PHOHY: test
test:
	go test -v ./...
