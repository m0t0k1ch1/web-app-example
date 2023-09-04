FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /usr/src/app/cmd/app

RUN go build -v -o /usr/local/bin/app

CMD ["app"]
