# Build stage
FROM golang:1.24.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /main

EXPOSE 8080
ENTRYPOINT ["/main"]
