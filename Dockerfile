FROM golang:1.26

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]


