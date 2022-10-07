FROM golang:latest as builder

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . ./

RUN go mod download && go mod verify
RUN go build -v -o atsast-apply-module-server github.com/njupt-sast/atsast-apply-module-server

FROM debian as runner
MAINTAINER GuGu Bai <0xfaner@gmail.com>

ENV GIN_MODE=release
WORKDIR /app

RUN apt update && apt install ca-certificates -y
COPY --from=builder /app/atsast-apply-module-server ./
COPY config/prod.yml ./config/

CMD ["./atsast-apply-module-server", "-conf=config/prod.yml"]
