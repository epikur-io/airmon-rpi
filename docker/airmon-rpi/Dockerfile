# build stage
FROM golang:1.25.1 AS builder

WORKDIR /app
COPY . .
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go env
RUN go mod tidy
RUN OUTPUT="build/airmon-rpi" CGO_ENABLED=0 task build

# image stage
FROM alpine:latest

USER root
WORKDIR /root/
RUN apk add --no-cache i2c-tools

COPY --from=builder --chown=root:root /app/build/airmon-rpi .
RUN mkdir -p /app/data
RUN touch /app/data/airmon_config.yaml

EXPOSE 8042
ENTRYPOINT ["./airmon-rpi", "-config", "/app/data/airmon_config.yaml"]
