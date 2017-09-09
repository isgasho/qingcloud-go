// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type ImageServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ImageServiceProperties) Reset()                    { *m = ImageServiceProperties{} }
func (m *ImageServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ImageServiceProperties) ProtoMessage()               {}
func (*ImageServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *ImageServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageServiceProperties)(nil), "service.ImageServiceProperties")
}

type ImageServiceInterface interface {
	DescribeImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CaptureInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyImageAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GrantImageToUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RevokeImageFromUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeImageUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CloneImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type ImageService struct {
	Config           *config.Config
	Properties       *ImageServiceProperties
	LastResponseBody string
}

func NewImageService(conf *config.Config, zone string) (p *ImageService) {
	return &ImageService{
		Config:     conf,
		Properties: &ImageServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Image(zone string) (*ImageService, error) {
	properties := &ImageServiceProperties{
		Zone: zone,
	}

	return &ImageService{Config: s.Config, Properties: properties}, nil
}

func (p *ImageService) DescribeImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) CaptureInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CaptureInstance",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) DeleteImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) ModifyImageAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyImageAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) GrantImageToUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GrantImageToUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) RevokeImageFromUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeImageFromUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) DescribeImageUsers(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeImageUsers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func (p *ImageService) CloneImages(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CloneImages",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func init() { proto.RegisterFile("image.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x19, 0x54, 0x65, 0x67, 0x43, 0x31, 0x68, 0x91, 0x8a, 0x20, 0x5e, 0x79, 0x21, 0x2d,
	0xe8, 0xb5, 0xe2, 0x68, 0xdd, 0xd8, 0x85, 0xa0, 0x53, 0xaf, 0x47, 0xff, 0x9c, 0x95, 0xb0, 0x36,
	0xa7, 0x9e, 0xa4, 0x83, 0xf9, 0x08, 0x3e, 0x98, 0xef, 0xe3, 0x1b, 0x48, 0x53, 0x04, 0xbd, 0xd8,
	0xc5, 0x7a, 0x13, 0x92, 0x2f, 0x5f, 0x7e, 0x5f, 0x92, 0x73, 0x60, 0x20, 0xcb, 0x38, 0x47, 0xbf,
	0x62, 0x32, 0x24, 0xf6, 0x34, 0xf2, 0x4a, 0xa6, 0xe8, 0x9d, 0xbd, 0x4b, 0x95, 0xa7, 0x05, 0xd5,
	0xd9, 0x5c, 0x67, 0xcb, 0x39, 0xd7, 0x05, 0x06, 0xcd, 0xd0, 0xfa, 0xbc, 0xd3, 0x9c, 0x28, 0x2f,
	0x30, 0xb0, 0xab, 0xa4, 0x5e, 0x04, 0x58, 0x56, 0x66, 0xdd, 0x6e, 0x5e, 0x5c, 0x81, 0x3b, 0x6d,
	0x98, 0x2f, 0x2d, 0xeb, 0x89, 0xa9, 0x42, 0x36, 0x12, 0xb5, 0x10, 0xe0, 0x7c, 0x90, 0xc2, 0x93,
	0xde, 0x79, 0xef, 0xb2, 0x3f, 0xb3, 0xf3, 0xeb, 0x2f, 0x07, 0x86, 0x7f, 0xed, 0xe2, 0x1e, 0xf6,
	0x23, 0xd4, 0x29, 0xcb, 0x04, 0xad, 0xae, 0x85, 0xeb, 0xb7, 0x71, 0xfe, 0x6f, 0x9c, 0xff, 0xd0,
	0xc4, 0x79, 0x1b, 0x74, 0x31, 0x82, 0x83, 0x30, 0xae, 0x4c, 0xcd, 0x38, 0x55, 0xda, 0xc4, 0x2a,
	0xc5, 0xad, 0x11, 0x77, 0x30, 0x8c, 0xb0, 0x40, 0xd3, 0xf5, 0x0a, 0x13, 0x38, 0x7e, 0xa4, 0x4c,
	0x2e, 0xd6, 0xf6, 0xfc, 0xc8, 0x18, 0x96, 0x49, 0x6d, 0x3a, 0x80, 0x42, 0x38, 0x9c, 0x70, 0xac,
	0x8c, 0xe5, 0xbc, 0xd2, 0x9b, 0x46, 0xde, 0x1e, 0x32, 0x86, 0xa3, 0x19, 0xae, 0x68, 0xd9, 0xbe,
	0x66, 0xcc, 0x54, 0x76, 0xe3, 0x44, 0x20, 0xfe, 0x95, 0xa6, 0x1b, 0xe5, 0x16, 0x06, 0x61, 0x41,
	0xaa, 0xe3, 0xd7, 0x7a, 0xee, 0xe7, 0xb7, 0x23, 0xa0, 0xff, 0x2c, 0x55, 0x1e, 0x36, 0x0d, 0x2a,
	0x76, 0x2c, 0x28, 0xd9, 0xb5, 0xbe, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xcf, 0xbf,
	0x9c, 0xd1, 0x02, 0x00, 0x00,
}
