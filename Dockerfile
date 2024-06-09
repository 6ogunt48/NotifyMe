# Build stage
FROM  golang:1.22.4-alpine3.19 AS builder
WORKDIR /sourcecode
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o checkemailbot . && chmod +x ./checkemailbot

# Final stage
FROM alpine:3.20.0 AS final
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot
WORKDIR /app
COPY --chown=nonroot:nonroot --from=builder /sourcecode/checkemailbot /app/checkemailbot
COPY --chown=nonroot:nonroot --from=builder /sourcecode/config.toml /app/config.toml
CMD ["./checkemailbot", "-config", "config.toml"]

# FROM --platform=linux/amd64 golang:1.22.4-alpine3.19 AS builder
# FROM --platform=linux/amd64 alpine:3.20.0 AS final



