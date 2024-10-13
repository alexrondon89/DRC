FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o collector ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl unzip bash openjdk11

RUN mkdir -p /flyway
RUN wget https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/9.16.0/flyway-commandline-9.16.0.tar.gz \
  && gzip -d flyway-commandline-9.16.0.tar.gz \
  && tar -xf flyway-commandline-9.16.0.tar --strip-components=1 -C /flyway \
  && rm flyway-commandline-9.16.0.tar
ENV PATH="/flyway:${PATH}"

WORKDIR /app
COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/flyway/ ./flyway/
COPY --from=builder /app/collector .
CMD ["./collector"]
EXPOSE 8000
