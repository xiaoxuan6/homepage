FROM golang:1.20.4-alpine3.18 AS build-dev

WORKDIR /go/src/app

COPY --link go.mod go.sum ./

RUN apk add --no-cache upx || \
        go version && \
        go env -w GO111MODULE=on && \
        go env -w GOPROXY=https://goproxy.cn,direct && \
        go mod tidy

COPY --link . .

RUN go build -o homepage . && \
    [ -e /usr/bin/upx ] && upx homepage || echo

FROM scratch

WORKDIR /src

COPY --link --from=build-dev /go/src/app/homepage ./homepage
COPY .env ./.env
COPY templates ./templates
COPY static ./static

EXPOSE 8080

CMD ["./homepage"]


