FROM golang:1.21.1-alpine3.18 AS builder


WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o plutusapi ./cmd/api/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/plutusapi .
EXPOSE 8080
CMD ["./plutusapi"]
