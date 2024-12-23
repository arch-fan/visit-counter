FROM golang:1.23.4 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o visit-counter .

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /build/visit-counter .

EXPOSE 8080
CMD ["./visit-counter"]
