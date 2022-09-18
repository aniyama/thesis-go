FROM golang:1.17.2-alpine as dev

# RUN go install golang.org/x/tools/cmd/goimports@latest \
#     && go install golang.org/x/tools/gopls@latest \
#     && go install golang.org/x/tools/cmd/godoc@latest \
#     && go install golang.org/x/lint/golint@latest \
#     && go install github.com/rogpeppe/godef@latest \
#     && go install github.com/nsf/gocode@latest \
#     # hot relord
#     && go install github.com/cosmtrek/air@latest \
#     # debug
#     && go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /go/src/app

# COPY go.mod ./ \
#     && go.sum ./

# RUN go mod download

RUN go get -u github.com/cosmtrek/air

CMD ["air"]



