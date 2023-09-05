.PHONY: generate
generate:
	buf generate _schema/proto
	sqlc generate
	wire gen ./...

.PHOHY: test
test:
	go test -v ./...
