// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vxnet.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

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

type VxnetServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *VxnetServiceProperties) Reset()                    { *m = VxnetServiceProperties{} }
func (m *VxnetServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*VxnetServiceProperties) ProtoMessage()               {}
func (*VxnetServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *VxnetServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*VxnetServiceProperties)(nil), "service.VxnetServiceProperties")
}

type VxnetServiceInterface interface {
	DescribeVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	JoinVxnet(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	LeaveVxnet(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyVxnetAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeVxnetInstances(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
}

type VxnetService struct {
	Config           *config.Config
	Properties       *VxnetServiceProperties
	LastResponseBody string
}

func NewVxnetService(conf *config.Config, zone string) (p *VxnetService) {
	return &VxnetService{
		Config:     conf,
		Properties: &VxnetServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Vxnet(zone string) (*VxnetService, error) {
	properties := &VxnetServiceProperties{
		Zone: zone,
	}

	return &VxnetService{Config: s.Config, Properties: properties}, nil
}

func (p *VxnetService) DescribeVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) CreateVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) DeleteVxnets(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVxnets",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) JoinVxnet(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "JoinVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) LeaveVxnet(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "LeaveVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) ModifyVxnetAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVxnetAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *VxnetService) DescribeVxnetInstances(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnetInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func init() { proto.RegisterFile("vxnet.proto", fileDescriptor30) }

var fileDescriptor30 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xab, 0xc8, 0x4b,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x95, 0x92, 0x2d, 0xcc, 0xcc, 0x4b, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0x89, 0x2f, 0x4e, 0xc9, 0x8e,
	0x2f, 0x2a, 0xcd, 0x49, 0xd5, 0x07, 0x11, 0x10, 0x75, 0x52, 0xd2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0x52, 0x49,
	0x87, 0x4b, 0x2c, 0x0c, 0x64, 0x66, 0x30, 0xc4, 0xac, 0x80, 0xa2, 0xfc, 0x82, 0xd4, 0xa2, 0x92,
	0xcc, 0xd4, 0x62, 0x21, 0x21, 0x2e, 0x96, 0xaa, 0xfc, 0xbc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x30, 0xdb, 0xe8, 0x1b, 0x33, 0x17, 0x0f, 0xb2, 0x72, 0x21, 0x07, 0x2e, 0x3e, 0x97,
	0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0xa4, 0x54, 0xb0, 0x78, 0xb1, 0x90, 0x98, 0x1e, 0xc4, 0x3a, 0x3d,
	0x98, 0x75, 0x7a, 0xae, 0x20, 0xeb, 0xa4, 0x70, 0x88, 0x0b, 0xd9, 0x71, 0xf1, 0x38, 0x17, 0xa5,
	0x26, 0x96, 0x50, 0xa0, 0xdf, 0x25, 0x35, 0x27, 0x95, 0x6c, 0xfd, 0xd6, 0x5c, 0x9c, 0x5e, 0xf9,
	0x99, 0x79, 0x60, 0xdd, 0x24, 0x6b, 0xb6, 0xe1, 0xe2, 0xf2, 0x49, 0x4d, 0x2c, 0x4b, 0x25, 0x4f,
	0xb7, 0x3b, 0x97, 0xa8, 0x6f, 0x7e, 0x4a, 0x66, 0x5a, 0x25, 0x58, 0xbb, 0x63, 0x49, 0x49, 0x51,
	0x66, 0x52, 0x69, 0x49, 0x2a, 0xe9, 0x7e, 0xf0, 0xe0, 0x12, 0x43, 0x89, 0x05, 0xcf, 0xbc, 0xe2,
	0x92, 0xc4, 0xbc, 0x64, 0xd2, 0x4d, 0x92, 0x12, 0xeb, 0xfa, 0xc8, 0x22, 0xc4, 0xc5, 0x19, 0x98,
	0x99, 0x97, 0xee, 0x0c, 0x4a, 0x50, 0x42, 0xac, 0x60, 0x03, 0x93, 0xd8, 0xc0, 0xea, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x5c, 0xf0, 0x2f, 0x81, 0x02, 0x00, 0x00,
}
