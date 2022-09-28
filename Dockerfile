FROM golang:1.17.2-alpine as dev

WORKDIR /go/src/app

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get -u github.com/cosmtrek/air \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/nsf/gocode@latest \
    # Alpine の cgo は musl-dev必要
    && apk add --no-cache musl-dev

CMD ["air"]
