FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/app ./cmd/app
COPY ./internal ./internal

RUN go build -o main ./cmd/app

EXPOSE 8081

CMD ["./main"]
