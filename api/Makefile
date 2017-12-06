# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/google/protobuf/releases
# go install github.com/golang/protobuf/protoc-gen-go
# go install github.com/chai2010/qingcloud-go/cmd/protoc-gen-qingcloud
# go install github.com/gogo/protobuf/protoc-gen-gogo

PROTO_API_FILES=$(sort $(wildcard ./*.proto))
PROTO_API_METADATA_FILES=$(sort $(wildcard ./spec_metadata/*.proto))

.PHONY: default
default: $(SOURCES) Makefile spec_metadata api

.PHONY: api
api: $(PROTO_API_FILES)
	go install ../cmd/protoc-gen-qingcloud
	protoc -I. --go_out=../pkg/api $(PROTO_API_FILES)
	protoc -I. --qingcloud_out=plugin=go-sdk:../pkg/api $(PROTO_API_FILES)
	protoc -I. --qingcloud_out=plugin=go-validator:../pkg/api $(PROTO_API_FILES)
	protoc -I. --qingcloud_out=plugin=go-server:../pkg/api $(PROTO_API_FILES)
	protoc -I. --qingcloud_out=plugin=qcli:../pkg/cmd/qcli/api $(PROTO_API_FILES)

.PHONY: spec_metadata
spec_metadata: $(PROTO_API_METADATA_FILES)
	protoc -I. --go_out=../../../.. $(PROTO_API_METADATA_FILES)

clean:
