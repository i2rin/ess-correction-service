FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ess-server ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/ess-server .

EXPOSE 8080
CMD ["./ess-server"]
