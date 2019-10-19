ARG GO_VERSION

FROM golang:${GO_VERSION}
WORKDIR /app
ADD ./gateway/build/main .
CMD ./main
