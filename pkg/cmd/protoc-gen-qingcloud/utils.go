// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"fmt"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

func getMessageDescriptor(msg proto.Message) (fd *protobuf.FileDescriptorProto, md *protobuf.DescriptorProto) {
	if msg, ok := msg.(descriptor.Message); ok {
		return descriptor.ForMessage(msg)
	}
	return
}

func getMessageOption(msg proto.Message) *spec_metadata.MessageOption {
	_, m := getMessageDescriptor(msg)

	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_MessageOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_MessageOption); ext != nil {
			if x, _ := ext.(*spec_metadata.MessageOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func getMessageFieldOption(m *protobuf.FieldDescriptorProto) *spec_metadata.FieldOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_FieldOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_FieldOption); ext != nil {
			if x, _ := ext.(*spec_metadata.FieldOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func validateMessage(msg proto.Message) error {
	if _, err := proto.Marshal(msg); err != nil {
		if reqNotSet, ok := err.(*proto.RequiredNotSetError); ok {
			return reqNotSet
		}
	}

	msgOption := getMessageOption(msg)
	if msgOption == nil {
		return nil
	}

	// check fields

	return fmt.Errorf("TODO")
}
