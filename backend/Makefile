.PHONY: generate
generate:
	buf generate _schema/proto
	sqlc generate

.PHOHY: test
test:
	go test -v ./...
