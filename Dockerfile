FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
RUN go build -o app ./cmd/app


FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./app"]
