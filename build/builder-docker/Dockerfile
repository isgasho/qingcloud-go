# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

FROM golang:1.9.2-alpine3.6 as builder

RUN apk add --no-cache git

RUN go get -d   github.com/chai2010/qingcloud-go/pkg/version
RUN go generate github.com/chai2010/qingcloud-go/pkg/version
RUN go install  github.com/chai2010/qingcloud-go/cmd/protoc-gen-qingcloud

RUN go get github.com/golang/protobuf/protoc-gen-go

FROM golang:1.9.2-alpine3.6

RUN apk add --no-cache git protobuf

COPY --from=builder /go/bin /go/bin
