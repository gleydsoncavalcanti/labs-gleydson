# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar arquivos de go.mod e go.sum primeiro (para cache de dependências)
COPY go.mod go.sum ./

RUN go mod download

# Copiar todo o código fonte
COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# frontend
COPY ./static ./static

EXPOSE 8080

CMD ["./main"]
