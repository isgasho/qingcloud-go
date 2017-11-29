// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validate/validate.proto

/*
Package validator is a generated protocol buffer package.

It is generated from these files:
	validate/validate.proto

It has these top-level messages:
	FieldValidator
*/
package validator

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of integer strictly equal this value.
	LengthEq         *int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FieldValidator) Reset()                    { *m = FieldValidator{} }
func (m *FieldValidator) String() string            { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()               {}
func (*FieldValidator) Descriptor() ([]byte, []int) { return fileDescriptorValidate, []int{0} }

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         20171130,
	Name:          "qingcloud.api.validator.field",
	Tag:           "bytes,20171130,opt,name=field",
	Filename:      "validate/validate.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "qingcloud.api.validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validate/validate.proto", fileDescriptorValidate) }

var fileDescriptorValidate = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x15, 0xb6, 0x6e, 0xa9, 0xbb, 0x95, 0xca, 0x02, 0xcd, 0x80, 0x26, 0x22, 0x38, 0x90,
	0x03, 0x24, 0x63, 0x47, 0x40, 0x1c, 0x40, 0xa1, 0x97, 0x02, 0x52, 0x0e, 0x1c, 0xb8, 0x44, 0x5e,
	0xf2, 0xd5, 0xb5, 0x70, 0x6c, 0xd7, 0xf9, 0x82, 0xca, 0x0b, 0xf0, 0x00, 0xbc, 0x18, 0xef, 0xc3,
	0x09, 0xc5, 0x59, 0xba, 0x20, 0x8d, 0x9b, 0xfd, 0xfb, 0x7d, 0xfe, 0x3b, 0xb1, 0xfe, 0xe4, 0xec,
	0x3b, 0x57, 0xb2, 0xe2, 0x08, 0xe9, 0xb0, 0x48, 0xac, 0x33, 0x68, 0xe8, 0xd9, 0x56, 0x6a, 0x51,
	0x2a, 0xd3, 0x56, 0x09, 0xb7, 0x32, 0xb9, 0xb6, 0xc6, 0x3d, 0x8c, 0x84, 0x31, 0x42, 0x41, 0xea,
	0xc7, 0xae, 0xda, 0x75, 0x5a, 0x41, 0x53, 0x3a, 0x69, 0xd1, 0xb8, 0xfe, 0xe8, 0x93, 0x9f, 0x87,
	0x64, 0xfe, 0x41, 0x82, 0xaa, 0xbe, 0x0c, 0x87, 0xe8, 0x3d, 0x32, 0x71, 0x20, 0x60, 0xc7, 0x82,
	0x28, 0x88, 0xa7, 0x79, 0xbf, 0xa1, 0xf7, 0xc9, 0x91, 0xd4, 0x58, 0x08, 0x64, 0x77, 0xa2, 0x20,
	0x3e, 0xc8, 0x27, 0x52, 0xe3, 0x12, 0x07, 0xac, 0x90, 0x1d, 0xec, 0xf1, 0x0a, 0xe9, 0x39, 0x21,
	0x75, 0x23, 0x0a, 0xd8, 0xc9, 0x06, 0x1b, 0x76, 0x18, 0x05, 0x71, 0x98, 0x4f, 0xeb, 0x46, 0x64,
	0x1e, 0xd0, 0xc7, 0x64, 0xb6, 0x69, 0x6b, 0xae, 0x0b, 0x70, 0xce, 0x38, 0x36, 0xf1, 0x17, 0x11,
	0x8f, 0xb2, 0x8e, 0xd0, 0x07, 0x24, 0x5c, 0x2b, 0xc3, 0xfd, 0x7d, 0x47, 0x51, 0x10, 0x07, 0xf9,
	0xb1, 0xdf, 0x2f, 0xf1, 0x46, 0x29, 0x64, 0xc7, 0x23, 0xb5, 0x42, 0xfa, 0x94, 0x9c, 0xf6, 0x0a,
	0x6c, 0x23, 0x95, 0xd1, 0x2c, 0xf4, 0xfe, 0xc4, 0xc3, 0xac, 0x67, 0xf4, 0x11, 0x99, 0x0e, 0xd1,
	0xc0, 0xa6, 0x7e, 0x20, 0xbc, 0xce, 0x86, 0x1b, 0xa9, 0x10, 0x18, 0x19, 0xc9, 0x15, 0x02, 0x8d,
	0xc9, 0xa2, 0x41, 0x27, 0xb5, 0x28, 0xb4, 0xc1, 0x02, 0x6a, 0x8b, 0x3f, 0xd8, 0xcc, 0xff, 0xda,
	0xbc, 0xe7, 0x9f, 0x0c, 0x66, 0x1d, 0xa5, 0xcf, 0x09, 0x75, 0x60, 0x81, 0x23, 0x54, 0x45, 0x69,
	0x5a, 0x8d, 0x45, 0x2d, 0x35, 0x3b, 0xf1, 0x2f, 0xb4, 0x18, 0xcc, 0xfb, 0x4e, 0x7c, 0x94, 0xfa,
	0xb6, 0x69, 0xbe, 0x63, 0xa7, 0xb7, 0x4d, 0xf3, 0x5d, 0xf7, 0x89, 0x0a, 0xb4, 0xc0, 0x4d, 0xf7,
	0x36, 0x73, 0x3f, 0x14, 0xf6, 0x60, 0x89, 0x23, 0xa9, 0x90, 0xdd, 0x1d, 0xcb, 0xd5, 0x58, 0xc2,
	0x96, 0x2d, 0xc6, 0x32, 0xdb, 0xbe, 0x2a, 0xc8, 0x64, 0xdd, 0xf5, 0x80, 0x9e, 0x27, 0x7d, 0x69,
	0x92, 0xa1, 0x34, 0x89, 0xef, 0xc7, 0x67, 0x8b, 0xd2, 0xe8, 0x86, 0xfd, 0xf9, 0xf5, 0xbb, 0x7b,
	0xb4, 0xd9, 0xe5, 0xb3, 0xe4, 0x3f, 0xad, 0x4b, 0xfe, 0xed, 0x53, 0xde, 0xe7, 0xbe, 0x7b, 0xfb,
	0xf5, 0x8d, 0x90, 0xb8, 0x69, 0xaf, 0x92, 0xd2, 0xd4, 0x69, 0xb9, 0xe1, 0xf2, 0xf2, 0xe2, 0xe5,
	0x45, 0xba, 0x0f, 0x79, 0x21, 0x4c, 0x6a, 0xbf, 0x89, 0x94, 0x5b, 0x99, 0xee, 0xb3, 0x5e, 0xef,
	0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa1, 0x26, 0xd0, 0xfe, 0x02, 0x00, 0x00,
}
