#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.15-alpine as builder

# Force the go compiler to use modules
ENV GO111MODULE=on

# Update OS package and install Git
RUN apk update && apk add git openssh && apk add build-base

# Set working directory
WORKDIR /go/src/gitlab.com/rabbit-test

# Install wait-for
ADD ./resources/docker/libs/wait-for /usr/local/bin/wait-for
RUN chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
ADD Makefile Makefile

RUN go mod download

# Install Fresh for local development
RUN go get github.com/pilu/fresh
#
# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.36.0
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/AlekSi/gocov-xml
