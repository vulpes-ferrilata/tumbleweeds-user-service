FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./user-service ./cmd

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/user-service .

COPY ./locales ./locales
COPY ./config.yaml .

EXPOSE 8080

ENTRYPOINT ["/user-service"]