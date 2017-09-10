// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qingcloud_sdk_rule/rule.proto

/*
Package qingcloud_sdk_rule is a generated protocol buffer package.

It is generated from these files:
	qingcloud_sdk_rule/rule.proto

It has these top-level messages:
	ServiceRule
	MethodRule
	MethodInputRule
*/
package qingcloud_sdk_rule

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 服务规则
// 有主服务和子服务之分, 子服务隶属于某个主服务
type ServiceRule struct {
	DocUrl         string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	ServiceName    string `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	SubServiceName string `protobuf:"bytes,3,opt,name=sub_service_name,json=subServiceName" json:"sub_service_name,omitempty"`
}

func (m *ServiceRule) Reset()                    { *m = ServiceRule{} }
func (m *ServiceRule) String() string            { return proto.CompactTextString(m) }
func (*ServiceRule) ProtoMessage()               {}
func (*ServiceRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *ServiceRule) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceRule) GetSubServiceName() string {
	if m != nil {
		return m.SubServiceName
	}
	return ""
}

// 方法规则
type MethodRule struct {
	DocUrl     string `protobuf:"bytes,1,opt,name=doc_url,json=docUrl" json:"doc_url,omitempty"`
	HttpMethod string `protobuf:"bytes,2,opt,name=http_method,json=httpMethod" json:"http_method,omitempty"`
}

func (m *MethodRule) Reset()                    { *m = MethodRule{} }
func (m *MethodRule) String() string            { return proto.CompactTextString(m) }
func (*MethodRule) ProtoMessage()               {}
func (*MethodRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MethodRule) GetDocUrl() string {
	if m != nil {
		return m.DocUrl
	}
	return ""
}

func (m *MethodRule) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

// 输入参数规则
// 输入参数成员只有数值类型和字符串, 以及对应的数组, 不含嵌套结构
type MethodInputRule struct {
	RequiredFileds  string `protobuf:"bytes,1,opt,name=required_fileds,json=requiredFileds" json:"required_fileds,omitempty"`
	DefaultValue    string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	EnumValue       string `protobuf:"bytes,3,opt,name=enum_value,json=enumValue" json:"enum_value,omitempty"`
	MinValue        string `protobuf:"bytes,4,opt,name=min_value,json=minValue" json:"min_value,omitempty"`
	MaxValue        string `protobuf:"bytes,5,opt,name=max_value,json=maxValue" json:"max_value,omitempty"`
	MultipleOfValue string `protobuf:"bytes,6,opt,name=multiple_of_value,json=multipleOfValue" json:"multiple_of_value,omitempty"`
}

func (m *MethodInputRule) Reset()                    { *m = MethodInputRule{} }
func (m *MethodInputRule) String() string            { return proto.CompactTextString(m) }
func (*MethodInputRule) ProtoMessage()               {}
func (*MethodInputRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MethodInputRule) GetRequiredFileds() string {
	if m != nil {
		return m.RequiredFileds
	}
	return ""
}

func (m *MethodInputRule) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func (m *MethodInputRule) GetEnumValue() string {
	if m != nil {
		return m.EnumValue
	}
	return ""
}

func (m *MethodInputRule) GetMinValue() string {
	if m != nil {
		return m.MinValue
	}
	return ""
}

func (m *MethodInputRule) GetMaxValue() string {
	if m != nil {
		return m.MaxValue
	}
	return ""
}

func (m *MethodInputRule) GetMultipleOfValue() string {
	if m != nil {
		return m.MultipleOfValue
	}
	return ""
}

var E_ServiceRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
	ExtensionType: (*ServiceRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.service_rule",
	Tag:           "bytes,10001,opt,name=service_rule,json=serviceRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MethodRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.method_rule",
	Tag:           "bytes,10001,opt,name=method_rule,json=methodRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

var E_MethodInputRule = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MethodInputRule)(nil),
	Field:         10001,
	Name:          "qingcloud.sdk.rule.method_input_rule",
	Tag:           "bytes,10001,opt,name=method_input_rule,json=methodInputRule",
	Filename:      "qingcloud_sdk_rule/rule.proto",
}

func init() {
	proto.RegisterType((*ServiceRule)(nil), "qingcloud.sdk.rule.ServiceRule")
	proto.RegisterType((*MethodRule)(nil), "qingcloud.sdk.rule.MethodRule")
	proto.RegisterType((*MethodInputRule)(nil), "qingcloud.sdk.rule.MethodInputRule")
	proto.RegisterExtension(E_ServiceRule)
	proto.RegisterExtension(E_MethodRule)
	proto.RegisterExtension(E_MethodInputRule)
}

func init() { proto.RegisterFile("qingcloud_sdk_rule/rule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xf8, 0x08, 0x64, 0x5c, 0x1a, 0xea, 0x0b, 0x11, 0xa8, 0x4d, 0x48, 0x0f, 0x44, 0x48,
	0xac, 0x4b, 0xb9, 0x85, 0x1b, 0x87, 0x4a, 0x1c, 0x9a, 0xa2, 0x54, 0x70, 0xe0, 0x62, 0x6c, 0xef,
	0xd8, 0x59, 0x75, 0xd7, 0xeb, 0xec, 0x47, 0xd5, 0xbf, 0xc1, 0x7f, 0xe4, 0x87, 0x20, 0xef, 0x6e,
	0x1c, 0xda, 0xd0, 0x5e, 0xac, 0xf5, 0x7b, 0x6f, 0xe6, 0x79, 0xe6, 0x79, 0xe1, 0x70, 0xcd, 0xea,
	0xaa, 0xe0, 0xd2, 0xd2, 0x54, 0xd3, 0xab, 0x54, 0x59, 0x8e, 0x49, 0xfb, 0x20, 0x8d, 0x92, 0x46,
	0xc6, 0x71, 0x47, 0x13, 0x4d, 0xaf, 0x48, 0xcb, 0xbc, 0x9e, 0x54, 0x52, 0x56, 0x1c, 0x13, 0xa7,
	0xc8, 0x6d, 0x99, 0x50, 0xd4, 0x85, 0x62, 0x8d, 0x91, 0xca, 0x57, 0x4d, 0x35, 0x44, 0x97, 0xa8,
	0xae, 0x59, 0x81, 0x4b, 0xcb, 0x31, 0x7e, 0x05, 0xcf, 0xa8, 0x2c, 0x52, 0xab, 0xf8, 0xa8, 0x37,
	0xe9, 0xcd, 0x06, 0xcb, 0x3e, 0x95, 0xc5, 0x77, 0xc5, 0xe3, 0xb7, 0xb0, 0xa7, 0xbd, 0x2e, 0xad,
	0x33, 0x81, 0xa3, 0x47, 0x8e, 0x8d, 0x02, 0xb6, 0xc8, 0x04, 0xc6, 0x33, 0x78, 0xa9, 0x6d, 0x9e,
	0xde, 0x92, 0x3d, 0x76, 0xb2, 0x7d, 0x6d, 0xf3, 0xcb, 0xad, 0x72, 0x7a, 0x06, 0x70, 0x8e, 0x66,
	0x25, 0xe9, 0xc3, 0x9e, 0x63, 0x88, 0x56, 0xc6, 0x34, 0xa9, 0x70, 0xda, 0x60, 0x09, 0x2d, 0xe4,
	0xab, 0xa7, 0x7f, 0x7a, 0x30, 0xf4, 0xc7, 0xaf, 0x75, 0x63, 0x8d, 0xeb, 0xf6, 0x0e, 0x86, 0x0a,
	0xd7, 0x96, 0x29, 0xa4, 0x69, 0xc9, 0x38, 0x52, 0x1d, 0xba, 0xee, 0x6f, 0xe0, 0x33, 0x87, 0xc6,
	0xc7, 0xf0, 0x82, 0x62, 0x99, 0x59, 0x6e, 0xd2, 0xeb, 0x8c, 0xdb, 0xcd, 0x48, 0x7b, 0x01, 0xfc,
	0xd1, 0x62, 0xf1, 0x21, 0x00, 0xd6, 0x56, 0x04, 0x85, 0x9f, 0x66, 0xd0, 0x22, 0x9e, 0x7e, 0x03,
	0x03, 0xc1, 0xea, 0xc0, 0x3e, 0x71, 0xec, 0x73, 0xc1, 0xea, 0x2d, 0x99, 0xdd, 0x04, 0xf2, 0x69,
	0x20, 0xb3, 0x1b, 0x4f, 0xbe, 0x87, 0x03, 0x61, 0xb9, 0x61, 0x0d, 0xc7, 0x54, 0x96, 0x41, 0xd4,
	0x77, 0xa2, 0xe1, 0x86, 0xb8, 0x28, 0x9d, 0x76, 0x4e, 0xb7, 0xbb, 0x6f, 0x53, 0x8d, 0xc7, 0xc4,
	0xc7, 0x4a, 0x36, 0xb1, 0x92, 0xb0, 0xdc, 0x8b, 0xc6, 0x30, 0x59, 0xeb, 0xd1, 0xef, 0xc5, 0xa4,
	0x37, 0x8b, 0x4e, 0xc7, 0x64, 0xf7, 0x97, 0x20, 0xff, 0xa4, 0xdd, 0xc5, 0xd7, 0xbe, 0xcc, 0x7f,
	0x41, 0xe4, 0x17, 0xed, 0x4d, 0x8e, 0x76, 0x4c, 0xfc, 0xa6, 0xef, 0x78, 0x1c, 0xfd, 0xcf, 0x63,
	0x1b, 0xee, 0x12, 0x44, 0x77, 0x9e, 0xaf, 0xe1, 0x20, 0x38, 0xb0, 0x36, 0xae, 0xfb, 0x86, 0x39,
	0x47, 0xad, 0xb3, 0xea, 0xee, 0x30, 0xc7, 0xf7, 0x1b, 0x75, 0xe1, 0x2f, 0x87, 0xe2, 0x36, 0xf0,
	0xe5, 0xdb, 0xcf, 0x45, 0xc5, 0xcc, 0xca, 0xe6, 0xa4, 0x90, 0x22, 0x29, 0x56, 0x19, 0x3b, 0x3d,
	0xf9, 0x78, 0x92, 0x74, 0xad, 0x3e, 0x54, 0x32, 0xd1, 0x0d, 0x16, 0xa4, 0xc9, 0x93, 0xdd, 0xeb,
	0xf5, 0x79, 0x17, 0xca, 0xfb, 0xee, 0x3b, 0x3f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x82, 0x97,
	0xe4, 0x5c, 0x8e, 0x03, 0x00, 0x00,
}
