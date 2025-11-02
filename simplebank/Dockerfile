# Build stage
FROM golang:1.25 AS builder
WORKDIR /THEPROJECT
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Run stage
FROM alpine:latest
WORKDIR /THEPROJECT
COPY --from=builder /THEPROJECT/main .
COPY app.env .

EXPOSE 8082
CMD ["./main"]
