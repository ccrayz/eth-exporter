FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o eth_exporter ./cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/eth_exporter .
CMD ["./eth_exporter"]
