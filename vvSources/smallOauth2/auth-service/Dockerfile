# Dockerfile for auth-service
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o auth-service

EXPOSE 9096

CMD ["./auth-service"]
