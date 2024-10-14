FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o csv-generator ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/csv-generator .
CMD ["./csv-generator"]
EXPOSE 8000