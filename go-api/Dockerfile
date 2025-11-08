FROM golang:1.21-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]