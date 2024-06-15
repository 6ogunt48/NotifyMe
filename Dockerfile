FROM golang:1.22.4-alpine3.19 AS builder
WORKDIR /sourcecode
COPY . .
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 go build -o checkemailbot . && chmod +x ./checkemailbot

FROM alpine:3.20.0 AS final
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot
WORKDIR /app
COPY --chown=nonroot:nonroot --from=builder /sourcecode/checkemailbot /app/checkemailbot
CMD ["./checkemailbot", "-config", "/app/config.toml"]