# syntax=docker/dockerfile:experimental
FROM golang:1.19 AS builder

ENV GOPATH=/
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ./ ./
USER root
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o book-crud ./cmd/main/app.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/book-crud .
EXPOSE 8080
ENTRYPOINT ["./book-crud"]