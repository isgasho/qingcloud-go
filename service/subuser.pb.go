// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subuser.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = config.Config{}
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SubuserServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SubuserServiceProperties) Reset()                    { *m = SubuserServiceProperties{} }
func (m *SubuserServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SubuserServiceProperties) ProtoMessage()               {}
func (*SubuserServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *SubuserServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeSubUsersInput struct {
}

func (m *DescribeSubUsersInput) Reset()                    { *m = DescribeSubUsersInput{} }
func (m *DescribeSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersInput) ProtoMessage()               {}
func (*DescribeSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

type DescribeSubUsersOutput struct {
}

func (m *DescribeSubUsersOutput) Reset()                    { *m = DescribeSubUsersOutput{} }
func (m *DescribeSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersOutput) ProtoMessage()               {}
func (*DescribeSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

type CreateSubUserInput struct {
}

func (m *CreateSubUserInput) Reset()                    { *m = CreateSubUserInput{} }
func (m *CreateSubUserInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserInput) ProtoMessage()               {}
func (*CreateSubUserInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

type CreateSubUserOutput struct {
}

func (m *CreateSubUserOutput) Reset()                    { *m = CreateSubUserOutput{} }
func (m *CreateSubUserOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserOutput) ProtoMessage()               {}
func (*CreateSubUserOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

type ModifySubUserAttributesInput struct {
}

func (m *ModifySubUserAttributesInput) Reset()                    { *m = ModifySubUserAttributesInput{} }
func (m *ModifySubUserAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesInput) ProtoMessage()               {}
func (*ModifySubUserAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

type ModifySubUserAttributesOutput struct {
}

func (m *ModifySubUserAttributesOutput) Reset()                    { *m = ModifySubUserAttributesOutput{} }
func (m *ModifySubUserAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesOutput) ProtoMessage()               {}
func (*ModifySubUserAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

type DeleteSubUsersInput struct {
}

func (m *DeleteSubUsersInput) Reset()                    { *m = DeleteSubUsersInput{} }
func (m *DeleteSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersInput) ProtoMessage()               {}
func (*DeleteSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

type DeleteSubUsersOutput struct {
}

func (m *DeleteSubUsersOutput) Reset()                    { *m = DeleteSubUsersOutput{} }
func (m *DeleteSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersOutput) ProtoMessage()               {}
func (*DeleteSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

type RestoreSubUsersInput struct {
}

func (m *RestoreSubUsersInput) Reset()                    { *m = RestoreSubUsersInput{} }
func (m *RestoreSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersInput) ProtoMessage()               {}
func (*RestoreSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

type RestoreSubUsersOutput struct {
}

func (m *RestoreSubUsersOutput) Reset()                    { *m = RestoreSubUsersOutput{} }
func (m *RestoreSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersOutput) ProtoMessage()               {}
func (*RestoreSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

func init() {
	proto.RegisterType((*SubuserServiceProperties)(nil), "service.SubuserServiceProperties")
	proto.RegisterType((*DescribeSubUsersInput)(nil), "service.DescribeSubUsersInput")
	proto.RegisterType((*DescribeSubUsersOutput)(nil), "service.DescribeSubUsersOutput")
	proto.RegisterType((*CreateSubUserInput)(nil), "service.CreateSubUserInput")
	proto.RegisterType((*CreateSubUserOutput)(nil), "service.CreateSubUserOutput")
	proto.RegisterType((*ModifySubUserAttributesInput)(nil), "service.ModifySubUserAttributesInput")
	proto.RegisterType((*ModifySubUserAttributesOutput)(nil), "service.ModifySubUserAttributesOutput")
	proto.RegisterType((*DeleteSubUsersInput)(nil), "service.DeleteSubUsersInput")
	proto.RegisterType((*DeleteSubUsersOutput)(nil), "service.DeleteSubUsersOutput")
	proto.RegisterType((*RestoreSubUsersInput)(nil), "service.RestoreSubUsersInput")
	proto.RegisterType((*RestoreSubUsersOutput)(nil), "service.RestoreSubUsersOutput")
}

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	Config           *config.Config
	Properties       *SubuserServiceProperties
	LastResponseBody string
}

func NewSubuserService(conf *config.Config, zone string) (p *SubuserService) {
	return &SubuserService{
		Config:     conf,
		Properties: &SubuserServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Subuser(zone string) (*SubuserService, error) {
	properties := &SubuserServiceProperties{
		Zone: zone,
	}

	return &SubuserService{Config: s.Config, Properties: properties}, nil
}

func (p *SubuserService) DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error) {
	if in == nil {
		in = &DescribeSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeSubUsersInput) Validate() error {
	return nil
}

func (p *SubuserService) CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error) {
	if in == nil {
		in = &CreateSubUserInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSubUser",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateSubUserOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *CreateSubUserInput) Validate() error {
	return nil
}

func (p *SubuserService) ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error) {
	if in == nil {
		in = &ModifySubUserAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySubUserAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifySubUserAttributesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *ModifySubUserAttributesInput) Validate() error {
	return nil
}

func (p *SubuserService) DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error) {
	if in == nil {
		in = &DeleteSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DeleteSubUsersInput) Validate() error {
	return nil
}

func (p *SubuserService) RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error) {
	if in == nil {
		in = &RestoreSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestoreSubUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &RestoreSubUsersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *RestoreSubUsersInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("subuser.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xed, 0x4a, 0x32, 0x41,
	0x14, 0x46, 0x5e, 0x79, 0xc5, 0x03, 0x5a, 0x1c, 0x77, 0x75, 0x99, 0xfc, 0x08, 0xa1, 0xe8, 0xd7,
	0x06, 0x75, 0x05, 0xa1, 0x7f, 0x0a, 0x24, 0x73, 0xe9, 0xb7, 0xb4, 0x7a, 0xb2, 0x21, 0xd9, 0xb1,
	0xf9, 0x08, 0xea, 0x12, 0xba, 0xc0, 0xee, 0xa3, 0x3b, 0x08, 0x9d, 0xc9, 0xda, 0x75, 0xb7, 0xfe,
	0x2c, 0xcb, 0xf3, 0x71, 0xce, 0xcc, 0xf3, 0x30, 0x50, 0x53, 0x26, 0x36, 0x8a, 0x64, 0xb8, 0x92,
	0x42, 0x0b, 0xac, 0x28, 0x92, 0xcf, 0x7c, 0x46, 0xac, 0xf3, 0xc4, 0x93, 0xc5, 0x6c, 0x29, 0xcc,
	0x7c, 0xaa, 0xe6, 0x8f, 0x53, 0x69, 0x96, 0x74, 0xba, 0xfe, 0x58, 0x5d, 0x3f, 0x84, 0x20, 0xb2,
	0xc6, 0xc8, 0x1a, 0xc6, 0x52, 0xac, 0x48, 0x6a, 0x4e, 0x0a, 0x11, 0xca, 0xaf, 0x22, 0xa1, 0xa0,
	0x74, 0x58, 0x3a, 0xa9, 0x4e, 0x36, 0xff, 0xfd, 0x16, 0xf8, 0x43, 0x52, 0x33, 0xc9, 0x63, 0x8a,
	0x4c, 0x7c, 0xab, 0x48, 0xaa, 0xcb, 0x64, 0x65, 0x74, 0x3f, 0x80, 0x66, 0x96, 0xb8, 0x36, 0x7a,
	0xcd, 0x78, 0x80, 0x03, 0x49, 0x77, 0xfa, 0x0b, 0xb7, 0x7a, 0x1f, 0x1a, 0x29, 0xd4, 0x89, 0xbb,
	0xd0, 0x1e, 0x89, 0x39, 0xbf, 0x7f, 0x71, 0xf0, 0x85, 0xd6, 0x92, 0xc7, 0x46, 0x93, 0x5b, 0xd3,
	0x83, 0x4e, 0x01, 0xef, 0x06, 0xf8, 0xd0, 0x18, 0xd2, 0x92, 0x74, 0xe6, 0x78, 0x4d, 0xf0, 0xd2,
	0xb0, 0x93, 0x37, 0xc1, 0x9b, 0x90, 0xd2, 0x42, 0x66, 0xf4, 0x2d, 0xf0, 0x33, 0xb8, 0x35, 0x9c,
	0xbd, 0xff, 0x83, 0x7a, 0x3a, 0x31, 0x8c, 0x60, 0x3f, 0x7b, 0x75, 0xec, 0x86, 0xae, 0x80, 0x30,
	0x37, 0x2e, 0xd6, 0x2b, 0xe4, 0xed, 0x1e, 0xbc, 0x82, 0x5a, 0x2a, 0x1f, 0x3c, 0xd8, 0x3a, 0x76,
	0xd3, 0x64, 0xed, 0x7c, 0xd2, 0xcd, 0x7a, 0x80, 0x56, 0x41, 0x68, 0x78, 0xb4, 0x35, 0xfe, 0x16,
	0x3b, 0x3b, 0xfe, 0x4b, 0xe6, 0x36, 0x8d, 0xa0, 0x9e, 0x8e, 0x19, 0xdb, 0x3f, 0x2e, 0xba, 0x53,
	0x0b, 0xeb, 0x14, 0xb0, 0x6e, 0xdc, 0x18, 0xf6, 0x32, 0x2d, 0xe0, 0xb7, 0x23, 0xaf, 0x37, 0xd6,
	0x2d, 0xa2, 0xed, 0x44, 0x16, 0xbc, 0x7d, 0x94, 0x3d, 0xa8, 0xde, 0xf0, 0x64, 0x31, 0x58, 0x3f,
	0x0a, 0xac, 0xb8, 0x32, 0xe3, 0xff, 0x9b, 0x07, 0x71, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x1d,
	0x85, 0x02, 0x4a, 0x49, 0x03, 0x00, 0x00,
}
