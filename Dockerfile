# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

FROM golang:alpine as builder

WORKDIR /go/src/github.com/chai2010/qingcloud-go/
COPY . .
RUN go install ./cmd/...

FROM alpine

COPY --from=builder /go/bin/qingcloud-cli /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-qingcloud /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/qingcloud-cli"]
CMD []
