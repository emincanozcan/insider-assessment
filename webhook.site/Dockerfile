FROM golang:1.22.2 AS builder

RUN mkdir /app-build
COPY .. /app-build/
WORKDIR /app-build/
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o /root/app /app-build/main.go

FROM alpine:3.17


WORKDIR /root
COPY --from=builder /root/app app

ENTRYPOINT ["./app"]
