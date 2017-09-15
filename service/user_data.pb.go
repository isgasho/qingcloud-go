// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_data.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserDataServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *UserDataServiceProperties) Reset()                    { *m = UserDataServiceProperties{} }
func (m *UserDataServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*UserDataServiceProperties) ProtoMessage()               {}
func (*UserDataServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *UserDataServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type UploadUserDataAttachmentInput struct {
	AttachmentContent []byte `protobuf:"bytes,2,opt,name=attachment_content,json=attachmentContent,proto3" json:"attachment_content,omitempty"`
	AttachmentName    string `protobuf:"bytes,1,opt,name=attachment_name,json=attachmentName" json:"attachment_name,omitempty"`
}

func (m *UploadUserDataAttachmentInput) Reset()                    { *m = UploadUserDataAttachmentInput{} }
func (m *UploadUserDataAttachmentInput) String() string            { return proto.CompactTextString(m) }
func (*UploadUserDataAttachmentInput) ProtoMessage()               {}
func (*UploadUserDataAttachmentInput) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{1} }

func (m *UploadUserDataAttachmentInput) GetAttachmentContent() []byte {
	if m != nil {
		return m.AttachmentContent
	}
	return nil
}

func (m *UploadUserDataAttachmentInput) GetAttachmentName() string {
	if m != nil {
		return m.AttachmentName
	}
	return ""
}

type UploadUserDataAttachmentOutput struct {
	Action       string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode      int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	AttachmentId string `protobuf:"bytes,4,opt,name=attachment_id,json=attachmentId" json:"attachment_id,omitempty"`
}

func (m *UploadUserDataAttachmentOutput) Reset()                    { *m = UploadUserDataAttachmentOutput{} }
func (m *UploadUserDataAttachmentOutput) String() string            { return proto.CompactTextString(m) }
func (*UploadUserDataAttachmentOutput) ProtoMessage()               {}
func (*UploadUserDataAttachmentOutput) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{2} }

func (m *UploadUserDataAttachmentOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *UploadUserDataAttachmentOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *UploadUserDataAttachmentOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadUserDataAttachmentOutput) GetAttachmentId() string {
	if m != nil {
		return m.AttachmentId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDataServiceProperties)(nil), "service.UserDataServiceProperties")
	proto.RegisterType((*UploadUserDataAttachmentInput)(nil), "service.UploadUserDataAttachmentInput")
	proto.RegisterType((*UploadUserDataAttachmentOutput)(nil), "service.UploadUserDataAttachmentOutput")
}

func init() { proto.RegisterFile("user_data.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x55, 0x76, 0xb3, 0xb0, 0x6b, 0xb1, 0x8b, 0xd6, 0x5a, 0xad, 0x02, 0x12, 0x15, 0x4a, 0xa5,
	0xc2, 0xa5, 0x89, 0x44, 0x6f, 0xdc, 0x2a, 0x7a, 0x28, 0x52, 0x55, 0x28, 0x94, 0x73, 0xe4, 0xc6,
	0x23, 0xb0, 0x9a, 0xd8, 0xa9, 0x3d, 0x69, 0xab, 0x7e, 0x02, 0xb7, 0xf6, 0xeb, 0xfa, 0x1b, 0xfc,
	0x41, 0x95, 0x84, 0x00, 0xaa, 0x4a, 0x7b, 0xb1, 0x3c, 0x33, 0x6f, 0xde, 0x3c, 0x8f, 0x1f, 0xa9,
	0xa7, 0x06, 0x74, 0xc0, 0x19, 0x32, 0x2f, 0xd1, 0x0a, 0x15, 0xad, 0x1a, 0xd0, 0xf7, 0x22, 0x84,
	0x66, 0xeb, 0x4e, 0xc8, 0x79, 0x18, 0xa9, 0x94, 0x07, 0x86, 0xdf, 0x06, 0x3a, 0x8d, 0xc0, 0xcf,
	0x8e, 0x02, 0xe7, 0xfa, 0xa4, 0x31, 0x33, 0xa0, 0xcf, 0x18, 0xb2, 0x69, 0xd1, 0x31, 0xd6, 0x2a,
	0x01, 0x8d, 0x02, 0x0c, 0xa5, 0xc4, 0x7e, 0x52, 0x12, 0x1c, 0xab, 0x6d, 0x75, 0x7f, 0x4d, 0xf2,
	0xbb, 0xfb, 0x6c, 0x91, 0xd6, 0x2c, 0x89, 0x14, 0xe3, 0x65, 0xdf, 0x29, 0x22, 0x0b, 0x17, 0x31,
	0x48, 0x1c, 0xca, 0x24, 0x45, 0x7a, 0x4c, 0x28, 0xdb, 0xa4, 0x82, 0x50, 0x49, 0x04, 0x89, 0xce,
	0xb7, 0xb6, 0xd5, 0xad, 0x4d, 0xfe, 0x6e, 0x2b, 0x83, 0xa2, 0x40, 0x3b, 0xa4, 0xbe, 0x03, 0x97,
	0x2c, 0x2e, 0xe7, 0xfd, 0xd9, 0xa6, 0x2f, 0x59, 0x0c, 0x7d, 0x67, 0xb9, 0xb2, 0xff, 0x7d, 0xc4,
	0xed, 0xbe, 0x58, 0xe4, 0x60, 0x9f, 0xa6, 0x51, 0x8a, 0x99, 0xa8, 0xff, 0xa4, 0xc2, 0x42, 0x14,
	0x4a, 0xae, 0xc9, 0xd7, 0x11, 0x6d, 0x90, 0x9f, 0x1a, 0x32, 0x26, 0x0e, 0xb9, 0xc4, 0x1f, 0x93,
	0xaa, 0x06, 0x1c, 0x28, 0x0e, 0xd4, 0x21, 0xd5, 0x18, 0x8c, 0x61, 0x73, 0x70, 0xbe, 0xe7, 0x3d,
	0x65, 0x48, 0x0f, 0xc9, 0xef, 0x1d, 0x15, 0x82, 0x3b, 0x76, 0x5e, 0xaf, 0x6d, 0x93, 0x43, 0xde,
	0x7b, 0xb5, 0x48, 0xfd, 0xdd, 0x6a, 0xe9, 0x03, 0x71, 0xf6, 0xe9, 0xa4, 0x47, 0xde, 0xfa, 0xcb,
	0xbc, 0x4f, 0xd7, 0xdb, 0xec, 0x7c, 0x89, 0x2b, 0x9e, 0xec, 0x92, 0xe5, 0xca, 0xae, 0x50, 0x7b,
	0x3c, 0x9a, 0x5e, 0x37, 0x2f, 0x96, 0x2b, 0xfb, 0x9c, 0xf4, 0x16, 0x88, 0x89, 0xe9, 0xfb, 0x3e,
	0x57, 0xa1, 0xf1, 0x36, 0xc6, 0xf0, 0x42, 0x15, 0xfb, 0x2c, 0x11, 0x7e, 0x66, 0xa2, 0xcc, 0x43,
	0xbe, 0x90, 0x1c, 0x1e, 0xbd, 0x05, 0xc6, 0x11, 0xa5, 0x57, 0x42, 0xce, 0x07, 0x39, 0xac, 0x9c,
	0x74, 0x53, 0xc9, 0xbd, 0x73, 0xf2, 0x16, 0x00, 0x00, 0xff, 0xff, 0x55, 0xd4, 0x9e, 0xe1, 0x76,
	0x02, 0x00, 0x00,
}
