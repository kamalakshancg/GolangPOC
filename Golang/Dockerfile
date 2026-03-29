FROM golang:1.26.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go_poc_app cmd/rest-api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/go_poc_app .
EXPOSE 8080
CMD ["./go_poc_app"]