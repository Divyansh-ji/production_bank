# ===== Build Stage =====
FROM golang:1.25-alpine AS builder
WORKDIR /THEPROJECT

# Install dependencies (curl, tar, gzip, ca-certificates)
RUN apk add --no-cache curl tar gzip ca-certificates

# Copy source code
COPY . .

# Build Go binary (use platform defaults)
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Download and extract migrate binary safely for the current arch
ARG TARGETOS=linux
ARG TARGETARCH
RUN curl -fsSL https://github.com/golang-migrate/migrate/releases/download/v4.19.0/migrate.${TARGETOS}-${TARGETARCH}.tar.gz \
  -o migrate.tar.gz && \
  tar -xzf migrate.tar.gz && \
  if [ -f migrate.${TARGETOS}-${TARGETARCH} ]; then mv migrate.${TARGETOS}-${TARGETARCH} migrate; else mv migrate migrate; fi && \
  rm -f migrate.tar.gz


# ===== Run Stage =====
FROM alpine:latest
WORKDIR /THEPROJECT

# Install SSL certs and minimal shell tools
RUN apk add --no-cache ca-certificates

# Copy binaries from builder
COPY --from=builder /THEPROJECT/main .
COPY --from=builder /THEPROJECT/migrate ./migrate

# Copy environment and migration files
COPY app.env .
COPY db/migration ./migration
COPY start.sh .

# Give execution permission
RUN sed -i 's/\r$//' start.sh && chmod +x start.sh

# Expose app port
EXPOSE 8082

# Entry point and default command
ENTRYPOINT ["/THEPROJECT/start.sh"]
CMD ["./main"]
