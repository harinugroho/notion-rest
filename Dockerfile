FROM golang:1.18.0-stretch

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

CMD air