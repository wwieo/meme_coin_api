FROM        golang:1.21.5-alpine AS builder
WORKDIR     /app
COPY        . .
RUN         apk add --no-cache make
RUN         go install github.com/swaggo/swag/cmd/swag@latest
RUN         make clean
RUN         make format
RUN         make docs cmd=memeCoin
RUN         CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o memeCoin -v ./cmd/memeCoin/main.go

FROM        alpine:latest
COPY        --from=builder /app/memeCoin /app/memeCoin
COPY        --from=builder /app/config/config.json /app/memeCoin.json
COPY        --from=builder /app/config/healthCheckOcean.json /app/healthCheck.json
EXPOSE      8000
CMD         ["/app/memeCoin", "-config=/app/memeCoin.json"]