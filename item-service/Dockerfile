# Build stage
# Use a multi-stage build to reduce the final image size
FROM golang:1.24.1-alpine3.21 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg go mod download && go mod verify
COPY . .
RUN go build -o /app ./cmd/main.go

# Final stage
FROM alpine:3.21
COPY --from=builder /app /app
EXPOSE 8080
CMD ["/app"]