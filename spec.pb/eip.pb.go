// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eip.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type EipServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *EipServiceProperties) Reset()                    { *m = EipServiceProperties{} }
func (m *EipServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*EipServiceProperties) ProtoMessage()               {}
func (*EipServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *EipServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*EipServiceProperties)(nil), "spec.EipServiceProperties")
}

type EipServiceInterface interface {
	DescribeEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AllocateEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ReleaseEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AssociateEip(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DissociateEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ChangeEipsBandwidth(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ChangeEipsBillingMode(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyEipAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type EipService struct {
	Config     *config.Config
	Properties *EipServiceProperties
}

func NewEipService(conf *config.Config, zone string) (p *EipService, err error) {
	return &EipService{
		Config:     conf,
		Properties: &EipServiceProperties{Zone: zone},
	}, nil
}

func (p *EipService) DescribeEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) AllocateEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AllocateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) ReleaseEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ReleaseEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) AssociateEip(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEip",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) DissociateEips(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) ChangeEipsBandwidth(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBandwidth",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) ChangeEipsBillingMode(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBillingMode",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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
func (p *EipService) ModifyEipAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyEipAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
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

func init() { proto.RegisterFile("eip.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x29, 0x04, 0xa1, 0xa3, 0x78, 0x58, 0xff, 0x20, 0xf5, 0x22, 0x9e, 0xc4, 0x43, 0x0a,
	0x7a, 0x56, 0xac, 0x76, 0xf1, 0x54, 0x90, 0xfa, 0x04, 0xc9, 0x66, 0x9a, 0x0e, 0xac, 0x99, 0x61,
	0x77, 0xaa, 0xd4, 0xe7, 0xf1, 0x41, 0xa5, 0x09, 0x62, 0x2e, 0x3d, 0x6c, 0x6e, 0xbb, 0xc3, 0xfc,
	0xbe, 0xdf, 0x37, 0x30, 0x46, 0x92, 0x5c, 0x02, 0x2b, 0x9b, 0x2c, 0x0a, 0xba, 0xc9, 0x65, 0xcd,
	0x5c, 0x7b, 0x9c, 0xb6, 0xb3, 0x72, 0xb3, 0x9a, 0xe2, 0x87, 0xe8, 0xb6, 0x5b, 0xb9, 0xbe, 0x85,
	0x53, 0x4b, 0xf2, 0x8e, 0xe1, 0x93, 0x1c, 0xbe, 0x05, 0x16, 0x0c, 0x4a, 0x18, 0x8d, 0x81, 0xec,
	0x9b, 0x1b, 0xbc, 0x18, 0x5d, 0x8d, 0x6e, 0xc6, 0xcb, 0xf6, 0x7d, 0xf7, 0x93, 0x01, 0xfc, 0x2f,
	0x9b, 0x47, 0x38, 0x9a, 0x63, 0x74, 0x81, 0x4a, 0xb4, 0x24, 0xd1, 0x9c, 0xe7, 0x9d, 0x28, 0xff,
	0x13, 0xe5, 0x76, 0x27, 0x9a, 0xec, 0x99, 0xef, 0xf8, 0x99, 0xf7, 0xec, 0x0a, 0x1d, 0xc6, 0x3f,
	0xc0, 0xe1, 0x12, 0x3d, 0x16, 0x71, 0xb8, 0x3e, 0x46, 0x76, 0xd4, 0xf9, 0x93, 0xf9, 0x27, 0x38,
	0x9e, 0x53, 0x2f, 0x20, 0xbd, 0x81, 0x85, 0x93, 0x97, 0x75, 0xd1, 0xd4, 0x2d, 0xfd, 0x5c, 0x34,
	0xd5, 0x17, 0x55, 0xba, 0x4e, 0x8e, 0x79, 0x85, 0xb3, 0x5e, 0x0c, 0x79, 0x4f, 0x4d, 0xbd, 0xe0,
	0x0a, 0x87, 0xf4, 0x59, 0x70, 0x45, 0xab, 0xad, 0x25, 0x99, 0xa9, 0x06, 0x2a, 0x37, 0x8a, 0xc9,
	0x67, 0x95, 0x07, 0xed, 0xff, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x66, 0x17, 0x2f, 0x8f, 0x89,
	0x02, 0x00, 0x00,
}
