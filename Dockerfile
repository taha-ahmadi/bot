# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/deep-bot .

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/deep-bot /app/deep-bot
ENV TICKETS_FILE=/data/tickets.json
USER nonroot:nonroot
ENTRYPOINT ["/app/deep-bot"]
