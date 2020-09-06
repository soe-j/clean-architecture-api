FROM golang:1.14

WORKDIR /go/src/app

CMD go mod download && go run main.go
