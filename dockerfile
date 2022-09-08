FROM golang:latest

COPY ./ ./

RUN go build -o app .
CMD [ "./main" ]

EXPOSE 8080

