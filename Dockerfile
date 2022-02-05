FROM golang:1.13-alpine AS build_base

RUN apk update && apk add ca-certificates git

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /app

COPY --from=server_builder /app/main .

CMD ["./main"]

EXPOSE 8080
