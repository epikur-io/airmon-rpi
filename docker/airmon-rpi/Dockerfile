# build stage
FROM golang:1.25.1 AS builder

WORKDIR /app
COPY . .
RUN DBG_GOPROXY="direct" go install github.com/go-task/task/v3/cmd/task@v3.45.4
RUN go env
RUN go mod tidy
RUN BUILD_OUTPUT="build/airmon-rpi" CGO_ENABLED=0 task build

# image stage
FROM alpine:latest

USER root
WORKDIR /root/

#RUN apk add --no-cache i2c-tools
# Use a mirror!
# RUN apk --no-cache -X https://mirror1.hs-esslingen.de/pub/Mirrors/alpine/latest-stable/main -X https://mirror1.hs-esslingen.de/pub/Mirrors/alpine/latest-stable/community \
#     add i2c-tools

COPY --from=builder --chown=root:root /app/docker/airmon-rpi/custom-mirror-repo /etc/apk/repositories 

RUN apk update
RUN apk add htop



COPY --from=builder --chown=root:root /app/build/airmon-rpi .
RUN mkdir -p /app/data
RUN touch /app/data/airmon_config.yaml

EXPOSE 8042
ENTRYPOINT ["./airmon-rpi", "-config", "/app/data/airmon_config.yaml"]
