FROM golang:1.19

ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o book-crud ./cmd/main/app.go

CMD ["./books-crud"]