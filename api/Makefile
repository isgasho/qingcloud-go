# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/google/protobuf/releases
# go install github.com/golang/protobuf/protoc-gen-go
# go install github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go

PROTO_FILES=$(wildcard ./*.proto)

default: $(SOURCES) Makefile rule spec_metadata
	go install ../cmd/protoc-gen-qingcloud-go
	protoc --go_out=../pkg/service.pb ${PROTO_FILES}
	protoc --qingcloud-go_out=../pkg/service.pb ${PROTO_FILES}

rule: ./qingcloud_sdk_rule/rule.proto
	protoc --go_out=../../../.. ./qingcloud_sdk_rule/rule.proto

.PHONY: spec_metadata
spec_metadata: $(wildcard ./spec_metadata/*.proto)
	protoc --go_out=../../../.. $<

clean:
