FROM golang:1.26.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o Go_POC_App cmd/rest-api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/Go_POC_App .
EXPOSE 8080
CMD ["./Go_POC_App"]