FROM golang:alpine as builder

RUN apk add -U --no-cache ca-certificates

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o crawler main.go

FROM scratch

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/crawler /usr/bin/

ENTRYPOINT ["crawler"]