// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

func BuildFileSpec(p *generator.Generator, file *generator.FileDescriptor) *spec_metadata.FileSpec {
	return &spec_metadata.FileSpec{
		FileName:    proto.String(file.GetName()),
		PackageName: proto.String(file.PackageName()),
	}
}

func BuildServiceSpec(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceSpec {
	spec := new(spec_metadata.ServiceSpec)

	spec.ServiceName = proto.String(generator.CamelCase(svc.GetName()))

	if opt := GetFileOption(file.FileDescriptorProto); opt != nil {
		_ = opt
	}

	for _, m := range svc.Method {
		methodSpec := &spec_metadata.ServiceMethodSpec{
			MethodName:     proto.String(generator.CamelCase(m.GetName())),
			InputTypeName:  proto.String(p.TypeName(p.ObjectNamed(m.GetInputType()))),
			OutputTypeName: proto.String(p.TypeName(p.ObjectNamed(m.GetOutputType()))),
		}

		if opt := GetServiceMethodOption(m); opt != nil {
			methodSpec.HttpMethod = proto.String(opt.GetHttpMethod())
		}

		spec.MethodList = append(spec.MethodList, methodSpec)
	}

	return spec
}

func GetFileOption(file *descriptor.FileDescriptorProto) *spec_metadata.FileOption {
	if file.Options != nil && proto.HasExtension(file.Options, spec_metadata.E_FileOption) {
		if ext, _ := proto.GetExtension(file.Options, spec_metadata.E_FileOption); ext != nil {
			if x, _ := ext.(*spec_metadata.FileOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func GetServiceOption(svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceOption {
	if svc.Options != nil && proto.HasExtension(svc.Options, spec_metadata.E_FileOption) {
		if ext, _ := proto.GetExtension(svc.Options, spec_metadata.E_FileOption); ext != nil {
			if x, _ := ext.(*spec_metadata.ServiceOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func GetServiceMethodOption(m *descriptor.MethodDescriptorProto) *spec_metadata.MethodOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_MethodOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_MethodOption); ext != nil {
			if x, _ := ext.(*spec_metadata.MethodOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func GetMessageOption(m *descriptor.DescriptorProto) *spec_metadata.MessageOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_MessageOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_MessageOption); ext != nil {
			if x, _ := ext.(*spec_metadata.MessageOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func GetMessageFieldOption(m *descriptor.FieldDescriptorProto) *spec_metadata.FieldOption {
	if m.Options != nil && proto.HasExtension(m.Options, spec_metadata.E_FieldOption) {
		if ext, _ := proto.GetExtension(m.Options, spec_metadata.E_FieldOption); ext != nil {
			if x, _ := ext.(*spec_metadata.FieldOption); x != nil {
				return x
			}
		}
	}
	return nil
}

func GetMethodInputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetInputType()).(*generator.Descriptor).DescriptorProto
}
func GetMethodOutputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto {
	return p.ObjectNamed(m.GetOutputType()).(*generator.Descriptor).DescriptorProto
}

func IsSupportedRepeated(field *descriptor.FieldDescriptorProto) bool {
	switch field.GetLabel() {
	case descriptor.FieldDescriptorProto_LABEL_REPEATED:
		return true
	}
	return false
}

func IsSupportedBool(field *descriptor.FieldDescriptorProto) bool {
	if IsSupportedRepeated(field) {
		return false
	}
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return true
	}
	return false
}

func IsSupportedInt(field *descriptor.FieldDescriptorProto) bool {
	if IsSupportedRepeated(field) {
		return false
	}
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_INT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_UINT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	}
	return false
}

func IsSupportedFloat(field *descriptor.FieldDescriptorProto) bool {
	if IsSupportedRepeated(field) {
		return false
	}
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_FLOAT, descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return true
	}
	return false
}

func IsSupportedString(field *descriptor.FieldDescriptorProto) bool {
	if IsSupportedRepeated(field) {
		return false
	}
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return true
	}
	return false
}

func _IsSupportedMap(field *descriptor.FieldDescriptorProto) bool {
	if IsSupportedRepeated(field) {
		return false
	}
	return false
}
