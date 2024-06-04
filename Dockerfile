FROM --platform=linux/amd64 golang:1.22.4-alpine3.19 as builder
ENV GOARCH=amd64
WORKDIR /sourcecode
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM --platform=linux/amd64 alpine:3.20.0 as final
USER root
RUN apk update && apk upgrade && apk --no-cache add dcron libcap  # cron needs root so we need dcron and lib cap to run it without root
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot
WORKDIR /app
RUN chown nonroot:nonroot /usr/sbin/crond && setcap cap_setgid=ep /usr/sbin/crond
COPY --chown=nonroot:nonroot --from=builder /sourcecode/main /app/main
COPY --chown=nonroot:nonroot --from=builder /sourcecode/config.toml /app/config.toml