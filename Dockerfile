FROM golang:1.18.0-stretch

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", "main.go"]