// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subuser.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"

var _ = config.Config{}
var _ = request.Request{}

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
func (*SubuserServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

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
func (*DescribeSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

type DescribeSubUsersOutput struct {
}

func (m *DescribeSubUsersOutput) Reset()                    { *m = DescribeSubUsersOutput{} }
func (m *DescribeSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSubUsersOutput) ProtoMessage()               {}
func (*DescribeSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

type CreateSubUserInput struct {
}

func (m *CreateSubUserInput) Reset()                    { *m = CreateSubUserInput{} }
func (m *CreateSubUserInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserInput) ProtoMessage()               {}
func (*CreateSubUserInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

type CreateSubUserOutput struct {
}

func (m *CreateSubUserOutput) Reset()                    { *m = CreateSubUserOutput{} }
func (m *CreateSubUserOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSubUserOutput) ProtoMessage()               {}
func (*CreateSubUserOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{4} }

type ModifySubUserAttributesInput struct {
}

func (m *ModifySubUserAttributesInput) Reset()                    { *m = ModifySubUserAttributesInput{} }
func (m *ModifySubUserAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesInput) ProtoMessage()               {}
func (*ModifySubUserAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{5} }

type ModifySubUserAttributesOutput struct {
}

func (m *ModifySubUserAttributesOutput) Reset()                    { *m = ModifySubUserAttributesOutput{} }
func (m *ModifySubUserAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifySubUserAttributesOutput) ProtoMessage()               {}
func (*ModifySubUserAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{6} }

type DeleteSubUsersInput struct {
}

func (m *DeleteSubUsersInput) Reset()                    { *m = DeleteSubUsersInput{} }
func (m *DeleteSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersInput) ProtoMessage()               {}
func (*DeleteSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{7} }

type DeleteSubUsersOutput struct {
}

func (m *DeleteSubUsersOutput) Reset()                    { *m = DeleteSubUsersOutput{} }
func (m *DeleteSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubUsersOutput) ProtoMessage()               {}
func (*DeleteSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{8} }

type RestoreSubUsersInput struct {
}

func (m *RestoreSubUsersInput) Reset()                    { *m = RestoreSubUsersInput{} }
func (m *RestoreSubUsersInput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersInput) ProtoMessage()               {}
func (*RestoreSubUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{9} }

type RestoreSubUsersOutput struct {
}

func (m *RestoreSubUsersOutput) Reset()                    { *m = RestoreSubUsersOutput{} }
func (m *RestoreSubUsersOutput) String() string            { return proto.CompactTextString(m) }
func (*RestoreSubUsersOutput) ProtoMessage()               {}
func (*RestoreSubUsersOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{10} }

func init() {
	proto.RegisterType((*SubuserServiceProperties)(nil), "spec.SubuserServiceProperties")
	proto.RegisterType((*DescribeSubUsersInput)(nil), "spec.DescribeSubUsersInput")
	proto.RegisterType((*DescribeSubUsersOutput)(nil), "spec.DescribeSubUsersOutput")
	proto.RegisterType((*CreateSubUserInput)(nil), "spec.CreateSubUserInput")
	proto.RegisterType((*CreateSubUserOutput)(nil), "spec.CreateSubUserOutput")
	proto.RegisterType((*ModifySubUserAttributesInput)(nil), "spec.ModifySubUserAttributesInput")
	proto.RegisterType((*ModifySubUserAttributesOutput)(nil), "spec.ModifySubUserAttributesOutput")
	proto.RegisterType((*DeleteSubUsersInput)(nil), "spec.DeleteSubUsersInput")
	proto.RegisterType((*DeleteSubUsersOutput)(nil), "spec.DeleteSubUsersOutput")
	proto.RegisterType((*RestoreSubUsersInput)(nil), "spec.RestoreSubUsersInput")
	proto.RegisterType((*RestoreSubUsersOutput)(nil), "spec.RestoreSubUsersOutput")
}

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	Config     *config.Config
	Properties *SubuserServiceProperties
}

func NewSubuserService(conf *config.Config, zone string) (p *SubuserService, err error) {
	return &SubuserService{
		Config:     conf,
		Properties: &SubuserServiceProperties{Zone: zone},
	}, nil
}

func (p *SubuserService) DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error) {
	if in == nil {
		in = &DescribeSubUsersInput{}
	}
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *RestoreSubUsersInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("subuser.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x29, 0x16, 0xc1, 0x0b, 0xad, 0x72, 0x3b, 0x7f, 0x4d, 0xeb, 0x0f, 0xe3, 0xc6, 0xd5,
	0x2c, 0xf4, 0x09, 0xc4, 0x01, 0x51, 0x28, 0x4a, 0x07, 0x1f, 0xc0, 0x8c, 0x57, 0x18, 0x90, 0x66,
	0xc8, 0x8f, 0xa0, 0x0f, 0xe3, 0xb3, 0x4a, 0x9b, 0x74, 0x91, 0x34, 0xed, 0x6e, 0x38, 0xe7, 0xe4,
	0xdc, 0xc9, 0x77, 0x03, 0x23, 0x65, 0xb8, 0x51, 0x24, 0xab, 0x5e, 0x0a, 0x2d, 0x70, 0xa8, 0x7a,
	0x6a, 0xcb, 0x0a, 0x8a, 0xc6, 0xca, 0x0d, 0xc9, 0xef, 0xae, 0xa5, 0x57, 0x29, 0x7a, 0x92, 0xba,
	0x23, 0x85, 0x08, 0xc3, 0x5f, 0xb1, 0xa2, 0x62, 0x70, 0x35, 0xb8, 0x39, 0x59, 0x6e, 0xbe, 0xcb,
	0x1c, 0xd2, 0x9a, 0x54, 0x2b, 0x3b, 0x4e, 0x8d, 0xe1, 0x6f, 0x8a, 0xa4, 0x7a, 0x5a, 0xf5, 0x46,
	0x97, 0x05, 0x64, 0xa1, 0xf1, 0x62, 0xf4, 0xda, 0x49, 0x00, 0x1f, 0x24, 0xbd, 0xeb, 0xad, 0x6e,
	0xf3, 0x29, 0x4c, 0x3c, 0xd5, 0x85, 0x2f, 0x60, 0xbe, 0x10, 0x1f, 0xdd, 0xe7, 0x8f, 0x93, 0xef,
	0xb5, 0x96, 0x1d, 0x37, 0x9a, 0xdc, 0x98, 0x4b, 0x38, 0xdf, 0xe3, 0xbb, 0x82, 0x14, 0x26, 0x35,
	0x7d, 0x91, 0x0e, 0x7e, 0x2f, 0x83, 0xc4, 0x97, 0x5d, 0x3c, 0x83, 0x64, 0x49, 0x4a, 0x0b, 0x19,
	0xe4, 0x73, 0x48, 0x03, 0xdd, 0x1e, 0xb8, 0xfd, 0x3b, 0x82, 0xb1, 0x4f, 0x0c, 0x17, 0x70, 0x16,
	0x5e, 0x1d, 0x67, 0xd5, 0x1a, 0x6f, 0x15, 0x65, 0xc5, 0xe6, 0x71, 0xd3, 0x4e, 0xc0, 0x1a, 0x46,
	0x1e, 0x19, 0x2c, 0x6c, 0x7c, 0x17, 0x22, 0x9b, 0x46, 0x1c, 0xd7, 0xc2, 0x21, 0xdf, 0x03, 0x0a,
	0x4b, 0x7b, 0xea, 0x10, 0x67, 0x76, 0x7d, 0x30, 0xe3, 0x66, 0x3c, 0xc2, 0xd8, 0x87, 0x8a, 0xd3,
	0xed, 0xcd, 0x76, 0x36, 0xc0, 0x58, 0xcc, 0x72, 0x45, 0xcf, 0x70, 0x1a, 0xd0, 0x46, 0x17, 0x8f,
	0x2d, 0x87, 0xcd, 0xa2, 0x9e, 0xed, 0xe2, 0xc7, 0x9b, 0xe7, 0x7d, 0xf7, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0x1c, 0xf0, 0x4f, 0xef, 0x02, 0x00, 0x00,
}
