#Build stage
FROM golang:1.25 AS builder
WORKDIR /THEPROJECT
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:latest
WORKDIR /THEPROJECT
COPY --from=builder /THEPROJECT/main .




EXPOSE 8082
CMD [ "/THEPROJECT/main" ]