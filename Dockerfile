FROM golang:1.8 as builder

WORKDIR /go/src/app
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o app

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /go/src/app/app .

EXPOSE 8080

CMD ["./app"]
