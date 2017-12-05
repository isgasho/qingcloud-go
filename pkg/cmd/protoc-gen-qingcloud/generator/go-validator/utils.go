// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package golang_validator

import (
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

func hasFieldOptions(message *protobuf.DescriptorProto) bool {
	for _, field := range message.GetField() {
		if getMessageFieldOption(field) != nil {
			return true
		}
	}
	return false
}

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

func pkgIsSupportedBool(field *protobuf.FieldDescriptorProto) bool {
	switch field.GetType() {
	case protobuf.FieldDescriptorProto_TYPE_BOOL:
		return true
	}
	return false
}

func pkgIsSupportedInt(field *protobuf.FieldDescriptorProto) bool {
	switch field.GetType() {
	case protobuf.FieldDescriptorProto_TYPE_INT32, protobuf.FieldDescriptorProto_TYPE_INT64:
		return true
	case protobuf.FieldDescriptorProto_TYPE_UINT32, protobuf.FieldDescriptorProto_TYPE_UINT64:
		return true
	case protobuf.FieldDescriptorProto_TYPE_SINT32, protobuf.FieldDescriptorProto_TYPE_SINT64:
		return true
	}
	return false
}

func pkgIsSupportedFloat(field *protobuf.FieldDescriptorProto) bool {
	switch field.GetType() {
	case protobuf.FieldDescriptorProto_TYPE_FLOAT, protobuf.FieldDescriptorProto_TYPE_DOUBLE:
		return true
	case protobuf.FieldDescriptorProto_TYPE_FIXED32, protobuf.FieldDescriptorProto_TYPE_FIXED64:
		return true
	case protobuf.FieldDescriptorProto_TYPE_SFIXED32, protobuf.FieldDescriptorProto_TYPE_SFIXED64:
		return true
	}
	return false
}

func pkgIsSupportedString(field *protobuf.FieldDescriptorProto) bool {
	switch field.GetType() {
	case protobuf.FieldDescriptorProto_TYPE_STRING:
		return true
	}
	return false
}
