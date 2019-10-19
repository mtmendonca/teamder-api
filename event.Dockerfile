ARG GO_VERSION

FROM golang:${GO_VERSION}
WORKDIR /app
ADD ./event/build/main .
CMD ./main
