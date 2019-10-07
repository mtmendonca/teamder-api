ARG GO_VERSION
ARG SERVICE
ARG ACCOUNT_API_PORT

FROM golang:${GO_VERSION}
RUN apk update && apk add git
WORKDIR /go/src/github.com/mtmendonca/teamder-api
ADD . .
CMD go build -o build/main cmd/main.go && ./build/main