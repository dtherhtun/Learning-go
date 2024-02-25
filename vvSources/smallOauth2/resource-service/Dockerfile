# Dockerfile for protected-service
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o protected-service

EXPOSE 9097

CMD ["./protected-service"]
