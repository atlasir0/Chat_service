FROM golang:1.21.9-alpine AS builder


COPY . /github.com/atlasir0/Chat_service/Auth_chat/source/
WORKDIR /github.com/atlasir0/Chat_service/Auth_chat/source/

RUN go mod download
RUN go build -o ./bin/auth_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/atlasir0/Chat_service/Auth_chat/source/bin/auth_service .
ADD .env .

CMD ["./auth_service"]