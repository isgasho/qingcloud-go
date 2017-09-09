// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hadoop.proto

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

type HadoopServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *HadoopServiceProperties) Reset()                    { *m = HadoopServiceProperties{} }
func (m *HadoopServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*HadoopServiceProperties) ProtoMessage()               {}
func (*HadoopServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *HadoopServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*HadoopServiceProperties)(nil), "service.HadoopServiceProperties")
}

type HadoopServiceInterface interface {
	AddHadoopNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteHadoopNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartHadoops(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopHadoops(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type HadoopService struct {
	Config           *config.Config
	Properties       *HadoopServiceProperties
	LastResponseBody string
}

func NewHadoopService(conf *config.Config, zone string) (p *HadoopService) {
	return &HadoopService{
		Config:     conf,
		Properties: &HadoopServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Hadoop(zone string) (*HadoopService, error) {
	properties := &HadoopServiceProperties{
		Zone: zone,
	}

	return &HadoopService{Config: s.Config, Properties: properties}, nil
}

func (p *HadoopService) AddHadoopNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddHadoopNodes",
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

func (p *HadoopService) DeleteHadoopNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteHadoopNodes",
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

func (p *HadoopService) StartHadoops(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartHadoops",
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

func (p *HadoopService) StopHadoops(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopHadoops",
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

func init() { proto.RegisterFile("hadoop.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4c, 0xc9,
	0xcf, 0x2f, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x95, 0x92, 0x2d, 0xcc, 0xcc, 0x4b, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0x89, 0x2f, 0x4e, 0xc9,
	0x8e, 0x2f, 0x2a, 0xcd, 0x49, 0xd5, 0x07, 0x11, 0x10, 0x75, 0x52, 0xd2, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x6a, 0x6e, 0x41, 0x49, 0x25, 0x44, 0x52,
	0x49, 0x97, 0x4b, 0xdc, 0x03, 0x6c, 0x68, 0x30, 0xc4, 0xb0, 0x80, 0xa2, 0xfc, 0x82, 0xd4, 0xa2,
	0x92, 0xcc, 0xd4, 0x62, 0x21, 0x21, 0x2e, 0x96, 0xaa, 0xfc, 0xbc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0xdb, 0x68, 0x1d, 0x13, 0x17, 0x2f, 0x8a, 0x7a, 0x21, 0x07, 0x2e, 0x3e,
	0xc7, 0x94, 0x14, 0x88, 0x98, 0x5f, 0x7e, 0x4a, 0x6a, 0xb1, 0x90, 0x98, 0x1e, 0xc4, 0x42, 0x3d,
	0x98, 0x85, 0x7a, 0xae, 0x20, 0x0b, 0xa5, 0x70, 0x88, 0x0b, 0x39, 0x73, 0x09, 0xba, 0xa4, 0xe6,
	0xa4, 0x96, 0xa4, 0x52, 0x62, 0x88, 0x1d, 0x17, 0x4f, 0x70, 0x49, 0x62, 0x51, 0x09, 0xc4, 0x0c,
	0xd2, 0xf5, 0xdb, 0x72, 0x71, 0x07, 0x97, 0xe4, 0x17, 0x90, 0xa9, 0x5d, 0x4a, 0xbc, 0xeb, 0x23,
	0x8b, 0x30, 0x17, 0x67, 0x60, 0x66, 0x5e, 0xba, 0x33, 0x28, 0x22, 0x84, 0xd8, 0x20, 0x26, 0x25,
	0xb1, 0x81, 0x15, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0x12, 0x87, 0x30, 0xbb, 0x01,
	0x00, 0x00,
}
