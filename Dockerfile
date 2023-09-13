FROM golang:latest as builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . ./

RUN go mod download && go mod verify
RUN go build -v -o atsast-apply-module-server

FROM debian as runner
MAINTAINER GuGu Bai <0xfaner@gmail.com>

ENV GIN_MODE=release
WORKDIR /app

# Optionally update mirror site
# RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources

# apt-get best practice <https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#apt-get>
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/atsast-apply-module-server ./
COPY config/prod.yml ./config/

CMD ["./atsast-apply-module-server", "-conf=config/prod.yml"]
