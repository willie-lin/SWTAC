FROM golang:1.20-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest
RUN apk add curl
CMD ["air", "-c", ".air.toml"]