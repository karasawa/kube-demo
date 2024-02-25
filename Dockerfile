FROM golang:1.21-bookworm as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM debian:bookworm-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main /app/.env ./

CMD ["/app/main"]
