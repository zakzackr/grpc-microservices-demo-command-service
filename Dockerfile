FROM golang:1.23-alpine
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/command
WORKDIR /go/src/command
COPY . /go/src/command

## Ginkgo CLI/SQL BoilerのCLIツール
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest && \
    go install github.com/aarondl/sqlboiler/v4@latest && \
    go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-mysql@latest