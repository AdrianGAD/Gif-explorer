# ──── Stage 1: Build ────
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gif-backend ./main.go

# ──── Stage 2: Runtime ────
FROM alpine:latest

RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /app/gif-backend /app/gif-backend
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app
EXPOSE 5050

HEALTHCHECK --interval=30s --timeout=5s --start-period=15s \
  CMD wget -qO- http://localhost:5050/health | grep -q '"status":"ok"' || exit 1

ENTRYPOINT ["./gif-backend"]
