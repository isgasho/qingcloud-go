# Copyright 2017 ChaiShushan<chaishushan@gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/google/protobuf/releases
# go install github.com/chai2010/qingcloud-go/protoc-gen-go

PROTO_FILES=$(wildcard ./*.proto)

default: $(SOURCES) Makefile
	protoc --go_out=../vendor ./qingcloud_sdk_rule/rule.proto
	go install ../protoc-gen-go
	protoc --go_out=plugins=qingcloud,aa=11,bb=22:../service --plugin=aa=11 ${PROTO_FILES}

debug:
	go install ../protoc-gen-qingcloud-go
	protoc --qingcloud-go_out=plugins=qingcloud,aa=11,bb=22:../service-debug ${PROTO_FILES}

rule: ./qingcloud_sdk_rule/rule.proto
	protoc --go_out=../vendor ./qingcloud_sdk_rule/rule.proto

clean:
