# Build stage
FROM golang:1.25 AS builder
WORKDIR /THEPROJECT

# Copy source code
COPY . .

# Build Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
RUN apk add curl 

# Download and extract migrate binary properly
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.0/migrate.linux-amd64.tar.gz \
  -o migrate.tar.gz && \
  tar -xvzf migrate.tar.gz && \
  mv migrate.linux-amd64 migrate && \
  rm migrate.tar.gz

# Run stage
FROM alpine:latest
WORKDIR /THEPROJECT

# Copy the compiled binary and migrate tool
COPY --from=builder /THEPROJECT/main .
COPY --from=builder /THEPROJECT/migrate ./migrate

# Copy env and migration files
COPY app.env .
COPY db/migration ./migration

# Expose API port
EXPOSE 8082

# Command to start app
CMD ["./main"]

