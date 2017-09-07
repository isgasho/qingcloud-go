// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_data.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import request_data_pkg "github.com/chai2010/qingcloud-go/request/data"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = request_data_pkg.Operation{}
var _ = errors.ParameterRequiredError{}

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
func (*UserDataServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

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
func (*UploadUserDataAttachmentInput) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

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
func (*UploadUserDataAttachmentOutput) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{2} }

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
	proto.RegisterType((*UserDataServiceProperties)(nil), "spec.UserDataServiceProperties")
	proto.RegisterType((*UploadUserDataAttachmentInput)(nil), "spec.UploadUserDataAttachmentInput")
	proto.RegisterType((*UploadUserDataAttachmentOutput)(nil), "spec.UploadUserDataAttachmentOutput")
}

type UserDataServiceInterface interface {
	UploadUserDataAttachment(in *UploadUserDataAttachmentInput) (out *UploadUserDataAttachmentOutput, err error)
}

type UserDataService struct {
	Config     *config.Config
	Properties *UserDataServiceProperties
}

func NewUserDataService(conf *config.Config, zone string) (p *UserDataService, err error) {
	return &UserDataService{
		Config:     conf,
		Properties: &UserDataServiceProperties{Zone: zone},
	}, nil
}

func (p *UserDataService) UploadUserDataAttachment(in *UploadUserDataAttachmentInput) (out *UploadUserDataAttachmentOutput, err error) {
	if in == nil {
		in = &UploadUserDataAttachmentInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UploadUserDataAttachment",
		RequestMethod: "POST", // GET or POST
	}

	x := &UploadUserDataAttachmentOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func init() { proto.RegisterFile("user_data.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x4d, 0xf9, 0xfa, 0x01, 0x4e, 0x10, 0xe2, 0x1e, 0x4c, 0x21, 0x6a, 0x48, 0x31, 0x91, 0x8b,
	0x90, 0xe8, 0x2f, 0x50, 0xbc, 0x70, 0x11, 0x02, 0x72, 0x26, 0x63, 0x77, 0x82, 0x4d, 0xe8, 0xee,
	0x66, 0x77, 0xaa, 0x89, 0x47, 0x4f, 0x9e, 0xf5, 0xa7, 0xf9, 0x17, 0xfc, 0x21, 0xa6, 0x05, 0x2c,
	0x31, 0x81, 0xdb, 0xce, 0x7b, 0x6f, 0xf6, 0xbd, 0x99, 0x81, 0x46, 0xea, 0xc8, 0xce, 0x25, 0x32,
	0xf6, 0x8c, 0xd5, 0xac, 0x85, 0xef, 0x0c, 0x45, 0xad, 0x93, 0x85, 0xd6, 0x8b, 0x25, 0xf5, 0xd1,
	0xc4, 0x7d, 0x54, 0x4a, 0x33, 0x72, 0xac, 0x95, 0x5b, 0x69, 0xc2, 0x3e, 0x34, 0x67, 0x8e, 0xec,
	0x1d, 0x32, 0x4e, 0xc9, 0x3e, 0xc7, 0x11, 0x8d, 0xad, 0x36, 0x64, 0x39, 0x26, 0x27, 0x04, 0xf8,
	0xaf, 0x5a, 0x51, 0xe0, 0xb5, 0xbd, 0xee, 0xc1, 0x24, 0x7f, 0x87, 0x2f, 0x70, 0x3a, 0x33, 0x4b,
	0x8d, 0x72, 0xd3, 0x76, 0xc3, 0x8c, 0xd1, 0x53, 0x42, 0x8a, 0x87, 0xca, 0xa4, 0x2c, 0x2e, 0x41,
	0xe0, 0x2f, 0x34, 0x8f, 0xb4, 0x62, 0x52, 0x1c, 0x94, 0xda, 0x5e, 0xb7, 0x36, 0x39, 0x2a, 0x98,
	0xc1, 0x8a, 0x10, 0x17, 0xd0, 0xd8, 0x92, 0x2b, 0x4c, 0x36, 0x76, 0xf5, 0x02, 0xbe, 0xc7, 0x84,
	0xc2, 0x0f, 0x0f, 0xce, 0x76, 0x39, 0x8f, 0x52, 0xce, 0xac, 0x8f, 0xa1, 0x8c, 0x51, 0x36, 0xdd,
	0xfa, 0x8b, 0x75, 0x25, 0x9a, 0x50, 0xb5, 0x94, 0x65, 0x91, 0x94, 0x07, 0xf9, 0x3f, 0xa9, 0x58,
	0xe2, 0x81, 0x96, 0x24, 0x02, 0xa8, 0x24, 0xe4, 0x1c, 0x2e, 0x28, 0xf8, 0x97, 0xf7, 0x6c, 0x4a,
	0xd1, 0x81, 0xc3, 0xad, 0x60, 0xb1, 0x0c, 0xfc, 0x9c, 0xaf, 0x15, 0xe0, 0x50, 0x5e, 0xbd, 0x7b,
	0xd0, 0xf8, 0xb3, 0x3f, 0x91, 0x42, 0xb0, 0x2b, 0xa7, 0xe8, 0xf4, 0xb2, 0x9b, 0xf4, 0xf6, 0x6e,
	0xb0, 0x75, 0xbe, 0x5f, 0xb4, 0x1a, 0x36, 0xac, 0xbf, 0x7d, 0x7d, 0x7f, 0x96, 0xaa, 0xb7, 0x65,
	0xf0, 0xc7, 0xa3, 0xe9, 0xc3, 0x63, 0x39, 0x3f, 0xe8, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x38, 0xdc, 0xe4, 0x07, 0x02, 0x00, 0x00,
}
