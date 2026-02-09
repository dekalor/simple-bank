# Build stage
FROM golang:1.25-alpine3.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env.example app.env

EXPOSE 3000
CMD [ "/app/main" ]