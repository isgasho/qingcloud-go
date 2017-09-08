// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type TagServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *TagServiceProperties) Reset()                    { *m = TagServiceProperties{} }
func (m *TagServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*TagServiceProperties) ProtoMessage()               {}
func (*TagServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *TagServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeTagsInput struct {
}

func (m *DescribeTagsInput) Reset()                    { *m = DescribeTagsInput{} }
func (m *DescribeTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeTagsInput) ProtoMessage()               {}
func (*DescribeTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

type DescribeTagsOutput struct {
}

func (m *DescribeTagsOutput) Reset()                    { *m = DescribeTagsOutput{} }
func (m *DescribeTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeTagsOutput) ProtoMessage()               {}
func (*DescribeTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

type CreateTagInput struct {
}

func (m *CreateTagInput) Reset()                    { *m = CreateTagInput{} }
func (m *CreateTagInput) String() string            { return proto.CompactTextString(m) }
func (*CreateTagInput) ProtoMessage()               {}
func (*CreateTagInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

type CreateTagOutput struct {
}

func (m *CreateTagOutput) Reset()                    { *m = CreateTagOutput{} }
func (m *CreateTagOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateTagOutput) ProtoMessage()               {}
func (*CreateTagOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

type DeleteTagsInput struct {
}

func (m *DeleteTagsInput) Reset()                    { *m = DeleteTagsInput{} }
func (m *DeleteTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteTagsInput) ProtoMessage()               {}
func (*DeleteTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

type DeleteTagsOutput struct {
}

func (m *DeleteTagsOutput) Reset()                    { *m = DeleteTagsOutput{} }
func (m *DeleteTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteTagsOutput) ProtoMessage()               {}
func (*DeleteTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

type ModifyTagAttributesInput struct {
}

func (m *ModifyTagAttributesInput) Reset()                    { *m = ModifyTagAttributesInput{} }
func (m *ModifyTagAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyTagAttributesInput) ProtoMessage()               {}
func (*ModifyTagAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

type ModifyTagAttributesOutput struct {
}

func (m *ModifyTagAttributesOutput) Reset()                    { *m = ModifyTagAttributesOutput{} }
func (m *ModifyTagAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyTagAttributesOutput) ProtoMessage()               {}
func (*ModifyTagAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

type AttachTagsInput struct {
}

func (m *AttachTagsInput) Reset()                    { *m = AttachTagsInput{} }
func (m *AttachTagsInput) String() string            { return proto.CompactTextString(m) }
func (*AttachTagsInput) ProtoMessage()               {}
func (*AttachTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

type AttachTagsOutput struct {
}

func (m *AttachTagsOutput) Reset()                    { *m = AttachTagsOutput{} }
func (m *AttachTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*AttachTagsOutput) ProtoMessage()               {}
func (*AttachTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

type DetachTagsInput struct {
}

func (m *DetachTagsInput) Reset()                    { *m = DetachTagsInput{} }
func (m *DetachTagsInput) String() string            { return proto.CompactTextString(m) }
func (*DetachTagsInput) ProtoMessage()               {}
func (*DetachTagsInput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{11} }

type DetachTagsOutput struct {
}

func (m *DetachTagsOutput) Reset()                    { *m = DetachTagsOutput{} }
func (m *DetachTagsOutput) String() string            { return proto.CompactTextString(m) }
func (*DetachTagsOutput) ProtoMessage()               {}
func (*DetachTagsOutput) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{12} }

func init() {
	proto.RegisterType((*TagServiceProperties)(nil), "spec.TagServiceProperties")
	proto.RegisterType((*DescribeTagsInput)(nil), "spec.DescribeTagsInput")
	proto.RegisterType((*DescribeTagsOutput)(nil), "spec.DescribeTagsOutput")
	proto.RegisterType((*CreateTagInput)(nil), "spec.CreateTagInput")
	proto.RegisterType((*CreateTagOutput)(nil), "spec.CreateTagOutput")
	proto.RegisterType((*DeleteTagsInput)(nil), "spec.DeleteTagsInput")
	proto.RegisterType((*DeleteTagsOutput)(nil), "spec.DeleteTagsOutput")
	proto.RegisterType((*ModifyTagAttributesInput)(nil), "spec.ModifyTagAttributesInput")
	proto.RegisterType((*ModifyTagAttributesOutput)(nil), "spec.ModifyTagAttributesOutput")
	proto.RegisterType((*AttachTagsInput)(nil), "spec.AttachTagsInput")
	proto.RegisterType((*AttachTagsOutput)(nil), "spec.AttachTagsOutput")
	proto.RegisterType((*DetachTagsInput)(nil), "spec.DetachTagsInput")
	proto.RegisterType((*DetachTagsOutput)(nil), "spec.DetachTagsOutput")
}

type TagServiceInterface interface {
	DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error)
	CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error)
	DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error)
	ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error)
	AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error)
	DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error)
}

type TagService struct {
	Config     *config.Config
	Properties *TagServiceProperties
}

func NewTagService(conf *config.Config, zone string) (p *TagService, err error) {
	return &TagService{
		Config:     conf,
		Properties: &TagServiceProperties{Zone: zone},
	}, nil
}

func (p *TagService) DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error) {
	if in == nil {
		in = &DescribeTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeTags",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeTagsOutput{}
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

func (p *DescribeTagsInput) Validate() error {
	return nil
}

func (p *TagService) CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error) {
	if in == nil {
		in = &CreateTagInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateTag",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateTagOutput{}
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

func (p *CreateTagInput) Validate() error {
	return nil
}

func (p *TagService) DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error) {
	if in == nil {
		in = &DeleteTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteTags",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteTagsOutput{}
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

func (p *DeleteTagsInput) Validate() error {
	return nil
}

func (p *TagService) ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error) {
	if in == nil {
		in = &ModifyTagAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyTagAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyTagAttributesOutput{}
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

func (p *ModifyTagAttributesInput) Validate() error {
	return nil
}

func (p *TagService) AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error) {
	if in == nil {
		in = &AttachTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachTags",
		RequestMethod: "GET", // GET or POST
	}

	x := &AttachTagsOutput{}
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

func (p *AttachTagsInput) Validate() error {
	return nil
}

func (p *TagService) DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error) {
	if in == nil {
		in = &DetachTagsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachTags",
		RequestMethod: "GET", // GET or POST
	}

	x := &DetachTagsOutput{}
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

func (p *DetachTagsInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("tag.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x29, 0x16, 0xa1, 0x83, 0xd8, 0x66, 0x9a, 0x6a, 0x8c, 0xa0, 0x92, 0x93, 0x78, 0xc8,
	0x41, 0x2f, 0x82, 0xa7, 0x60, 0x2f, 0x1e, 0x44, 0xd1, 0xe0, 0x7d, 0x13, 0xc7, 0x18, 0x90, 0x26,
	0x6c, 0x26, 0x82, 0xfe, 0x12, 0x7f, 0xae, 0xa4, 0xbb, 0xe9, 0x26, 0x9b, 0xf6, 0x16, 0xde, 0xbc,
	0x99, 0xbc, 0xf9, 0x66, 0x61, 0xc2, 0x22, 0x0b, 0x4b, 0x59, 0x70, 0x81, 0xe3, 0xaa, 0xa4, 0x34,
	0xb8, 0x02, 0x37, 0x16, 0xd9, 0x2b, 0xc9, 0xef, 0x3c, 0xa5, 0x67, 0x59, 0x94, 0x24, 0x39, 0xa7,
	0x0a, 0x11, 0xc6, 0xbf, 0xc5, 0x8a, 0xbc, 0xd1, 0xc5, 0xe8, 0x72, 0xf2, 0xb2, 0xfe, 0x0e, 0xe6,
	0xe0, 0x2c, 0xa9, 0x4a, 0x65, 0x9e, 0x50, 0x2c, 0xb2, 0xea, 0x61, 0x55, 0xd6, 0x1c, 0xb8, 0x80,
	0x5d, 0xf1, 0xa9, 0xe6, 0x46, 0x9d, 0xc1, 0xe1, 0xbd, 0x24, 0xc1, 0x8d, 0xa6, 0x7c, 0x0e, 0x4c,
	0x37, 0x8a, 0x36, 0x39, 0x30, 0x5d, 0xd2, 0x17, 0x71, 0x67, 0x1a, 0xc2, 0xcc, 0x48, 0xda, 0xe6,
	0x83, 0xf7, 0x58, 0xbc, 0xe7, 0x1f, 0x3f, 0xb1, 0xc8, 0x22, 0x66, 0x99, 0x27, 0x35, 0x93, 0xf6,
	0x9f, 0xc2, 0xc9, 0x96, 0x9a, 0x99, 0x1f, 0x31, 0x8b, 0xf4, 0xb3, 0x37, 0xdf, 0x48, 0xdd, 0x18,
	0x03, 0x9b, 0x91, 0x94, 0xed, 0xfa, 0x6f, 0x0f, 0xc0, 0xa0, 0xc2, 0x08, 0x0e, 0xba, 0x7b, 0xe3,
	0x71, 0xd8, 0xf0, 0x0c, 0x07, 0x80, 0x7c, 0x6f, 0x58, 0x50, 0x13, 0xf1, 0x16, 0x26, 0x1b, 0x24,
	0xe8, 0x2a, 0x5b, 0x9f, 0x9a, 0xbf, 0xb0, 0x54, 0xdd, 0x79, 0x07, 0x60, 0x30, 0xe1, 0xa2, 0xfd,
	0x43, 0x8f, 0xa5, 0x7f, 0x64, 0xcb, 0xba, 0xf9, 0x0d, 0xe6, 0x5b, 0x98, 0xe1, 0x99, 0xb2, 0xef,
	0x42, 0xed, 0x9f, 0xef, 0xac, 0x9b, 0x50, 0x86, 0x6d, 0x1b, 0xca, 0x3a, 0x40, 0x1b, 0xca, 0x3e,
	0x82, 0xda, 0xc8, 0x6e, 0xb6, 0xce, 0x62, 0x36, 0xea, 0x37, 0x27, 0xfb, 0xeb, 0x17, 0x7d, 0xf3,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x61, 0x0c, 0xed, 0xde, 0x02, 0x00, 0x00,
}
