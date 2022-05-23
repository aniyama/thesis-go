FROM golang:1.17.2-alpine

# LABEL maintainer ="K"

# WORKDIR /go/src/app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN go build


# RUN apk update && \
#     apk add --no-cache alpine-sdk git

# ENV GO111MODULE=on

# RUN go get -u github.com/cosmtrek/air && \
#     go build -o /go/bin/air github.com/cosmtrek/air

WORKDIR /go/src/app

RUN go get -u github.com/cosmtrek/air

CMD ["air"]



