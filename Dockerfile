FROM golang:1.23.4-alpine AS builder

WORKDIR /build

RUN apk add --no-cache build-base

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o visit-counter .

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /build/visit-counter .

EXPOSE 8080
CMD ["./visit-counter"]
