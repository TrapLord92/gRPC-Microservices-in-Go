FROM golang:1.23.1-alpine AS builder

COPY . /grpc-Microservice-BigTech-view/gRPC-Microservices-in-Go/ClientandServerSide/source/
WORKDIR /grpc-Microservice-BigTech-view/gRPC-Microservices-in-Go/ClientandServerSide/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /grpc-Microservice-BigTech-view/gRPC-Microservices-in-Go/ClientandServerSide/source/bin/crud_server .

CMD ["./crud_server"]