FROM golang:1.23.3 as builder

ENV CGO_ENABLED=0

ADD . /go/src

WORKDIR /go/src

RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler
RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go generate -v ./...
RUN go build -v -o echo-grpc-server ./cmd/...

FROM scratch

COPY --from=builder /go/src/echo-grpc-server /echo-grpc-server

ENTRYPOINT [ "/echo-grpc-server" ]
