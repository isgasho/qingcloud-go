# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/google/protobuf/releases
# go install github.com/golang/protobuf/protoc-gen-go
# go install github.com/chai2010/qingcloud-go/cmd/protoc-gen-qingcloud
# go install github.com/gogo/protobuf/protoc-gen-gogo

PROTO_API_FILES=$(sort $(wildcard ./*.proto))
PROTO_API_METADATA_FILES=$(sort $(wildcard ./spec_metadata/*.proto))
PROTO_API_VALIDATE_FILES=$(sort $(wildcard ./validate/*.proto))

.PHONY: default
default: $(SOURCES) Makefile spec_metadata validate api

.PHONY: api
api: $(PROTO_API_FILES)
	go install ../cmd/protoc-gen-qingcloud
	go install ../cmd/protoc-gen-qingcloud-validate
	protoc --go_out=../pkg/api $(PROTO_API_FILES)
	protoc --qingcloud_out=plugin=golang:../pkg/api $(PROTO_API_FILES)

.PHONY: spec_metadata
spec_metadata: $(PROTO_API_METADATA_FILES)
	protoc --go_out=../../../.. $(PROTO_API_METADATA_FILES)

.PHONY: validate
validate: $(PROTO_API_VALIDATE_FILES)
	protoc --gogo_out=../../../.. $(PROTO_API_VALIDATE_FILES)


clean:
