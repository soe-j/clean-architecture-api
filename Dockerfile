FROM golang:1.14 as base
WORKDIR /go/src/app
COPY . .
RUN go mod download

FROM base as run
ENTRYPOINT [ "go" ]
CMD [ "run main.go" ]

FROM base as lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0
ENTRYPOINT [ "golangci-lint" ]
CMD [ "run" ]

FROM base as air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
ENTRYPOINT [ "air" ]
