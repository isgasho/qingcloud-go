// Code generated by protoc-gen-go. DO NOT EDIT.
// source: key_pair.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"

var _ = config.Config{}
var _ = request.Request{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type KeyPairServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *KeyPairServiceProperties) Reset()                    { *m = KeyPairServiceProperties{} }
func (m *KeyPairServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*KeyPairServiceProperties) ProtoMessage()               {}
func (*KeyPairServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *KeyPairServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyPairServiceProperties)(nil), "spec.KeyPairServiceProperties")
}

type KeyPairServiceInterface interface {
	DescribeKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateKeyPair(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AttachKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DetachKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyKeyPairAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type KeyPairService struct {
	Config     *config.Config
	Properties *KeyPairServiceProperties
}

func NewKeyPairService(conf *config.Config, zone string) (p *KeyPairService, err error) {
	return &KeyPairService{
		Config:     conf,
		Properties: &KeyPairServiceProperties{Zone: zone},
	}, nil
}

func (p *KeyPairService) DescribeKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeKeyPairs",
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
func (p *KeyPairService) CreateKeyPair(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateKeyPair",
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
func (p *KeyPairService) DeleteKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteKeyPairs",
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
func (p *KeyPairService) AttachKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachKeyPairs",
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
func (p *KeyPairService) DetachKeyPairs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachKeyPairs",
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
func (p *KeyPairService) ModifyKeyPairAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyKeyPairAttributes",
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

func init() { proto.RegisterFile("key_pair.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4e, 0xad, 0x8c,
	0x2f, 0x48, 0xcc, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d,
	0x96, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7,
	0xe6, 0x16, 0x94, 0x54, 0x42, 0x94, 0x28, 0xe9, 0x71, 0x49, 0x78, 0xa7, 0x56, 0x06, 0x24, 0x66,
	0x16, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x06, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64,
	0xa6, 0x16, 0x0b, 0x09, 0x71, 0xb1, 0x54, 0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x46, 0xcb, 0x98, 0xb9, 0xf8, 0x50, 0x35, 0x08, 0x39, 0x71, 0x09, 0xb8, 0xa4,
	0x16, 0x27, 0x17, 0x65, 0x26, 0xa5, 0x42, 0x65, 0x8a, 0x85, 0xc4, 0xf4, 0x20, 0x96, 0xea, 0xc1,
	0x2c, 0xd5, 0x73, 0x05, 0x59, 0x2a, 0x85, 0x43, 0x5c, 0xc8, 0x9e, 0x8b, 0xd7, 0xb9, 0x28, 0x35,
	0xb1, 0x04, 0x66, 0x02, 0xc9, 0x06, 0x38, 0x70, 0xf1, 0xb9, 0xa4, 0xe6, 0xa4, 0x96, 0x90, 0xef,
	0x04, 0x07, 0x2e, 0x3e, 0xc7, 0x92, 0x92, 0xc4, 0xe4, 0x0c, 0x4a, 0x4c, 0x70, 0x49, 0xa5, 0xc8,
	0x04, 0x4f, 0x2e, 0x71, 0xdf, 0xfc, 0x94, 0xcc, 0xb4, 0x4a, 0xa8, 0x09, 0x8e, 0x25, 0x25, 0x45,
	0x99, 0x49, 0xa5, 0x25, 0xa9, 0x24, 0x1b, 0x95, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x7b, 0x23, 0x01, 0x14, 0x02, 0x00, 0x00,
}
