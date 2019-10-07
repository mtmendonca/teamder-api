ARG GO_VERSION

FROM golang:${GO_VERSION}
RUN apk update && apk add dep git
RUN go get github.com/codegangsta/gin
WORKDIR /go/src/github.com/mtmendonca/backend
ADD . .
WORKDIR /go/src/github.com/mtmendonca/backend/account
RUN dep ensure
CMD gin -p 80 -a 8080 run ./main.go