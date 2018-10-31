# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

# docker run --rm -it -v `pwd`:/root -w /root chai2010/qingcloud-go qcli

FROM golang:1.11 as builder

WORKDIR /build-dir
COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOBIN=/build-dir

RUN echo module qingcloud_sdk > /build-dir/go.mod
RUN git describe --tags --always > /build-dir/version
RUN git describe --exact-match 2>/dev/null || git log -1 --format="%H" > /build-dir/version

RUN go get github.com/chai2010/qingcloud-go/cmd/qcli@$(cat /build-dir/version)
RUN go get github.com/chai2010/qingcloud-go/cmd/qlua@$(cat /build-dir/version)

RUN echo version: $(cat /build-dir/version)

# -----------------------------------------------------------------------------
# for image
# -----------------------------------------------------------------------------

FROM alpine:3.7

COPY --from=builder /go/bin/qcli /usr/local/bin/qcli
COPY --from=builder /go/bin/qlua /usr/local/bin/qlua

ENTRYPOINT []
CMD ["/usr/local/bin/qcli"]

# -----------------------------------------------------------------------------
# END
# -----------------------------------------------------------------------------
