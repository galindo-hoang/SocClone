FROM golang:latest

COPY . .
RUN go mod download


RUN go build -o App ./cmd/api

CMD "./App"

