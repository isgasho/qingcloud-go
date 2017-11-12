# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

FROM golang:alpine as builder

RUN apk add --no-cache git

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/chai2010/qingcloud-go/cmd/...

FROM alpine

RUN apk add --no-cache protobuf

COPY --from=builder /go/bin/protoc-gen-go /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-qingcloud /usr/local/bin/

CMD ["sh"]
