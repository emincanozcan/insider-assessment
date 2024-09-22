FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app main.go

FROM alpine:3.17

WORKDIR /root

COPY --from=builder /app/app .
COPY --from=builder /app/internal/database/migrations /root/internal/database/migrations

ENTRYPOINT ["./app"]
