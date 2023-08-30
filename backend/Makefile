.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: generate
generate:
	buf generate _schema/proto
	sqlc generate
