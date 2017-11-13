# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/google/protobuf/releases
# go install github.com/golang/protobuf/protoc-gen-go
# go install github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go

PROTO_API_FILES=$(wildcard ./*.proto)
PROTO_API_METADATA_FILES=$(wildcard ./*.proto)

.PHONY: default
default: $(SOURCES) Makefile spec_metadata api

.PHONY: api
api: $(PROTO_API_FILES)
	go install ../cmd/protoc-gen-qingcloud
	protoc --go_out=../pkg/api $(PROTO_API_FILES)
	protoc --qingcloud_out=plugin=golang:../pkg/api $(PROTO_API_FILES)

.PHONY: spec_metadata
spec_metadata: $(PROTO_API_METADATA_FILES)
	protoc --go_out=../../../.. $(PROTO_API_METADATA_FILES)

clean:
