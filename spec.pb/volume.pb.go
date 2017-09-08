// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volume.proto

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

type VolumesServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *VolumesServiceProperties) Reset()                    { *m = VolumesServiceProperties{} }
func (m *VolumesServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*VolumesServiceProperties) ProtoMessage()               {}
func (*VolumesServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *VolumesServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*VolumesServiceProperties)(nil), "spec.VolumesServiceProperties")
}

type VolumesServiceInterface interface {
	DescribeVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AttachVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DetachVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyVolumeAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type VolumesService struct {
	Config     *config.Config
	Properties *VolumesServiceProperties
}

func NewVolumesService(conf *config.Config, zone string) (p *VolumesService, err error) {
	return &VolumesService{
		Config:     conf,
		Properties: &VolumesServiceProperties{Zone: zone},
	}, nil
}

func (p *VolumesService) DescribeVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVolumes",
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
func (p *VolumesService) CreateVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumes",
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
func (p *VolumesService) DeleteVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVolumes",
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
func (p *VolumesService) AttachVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachVolumes",
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
func (p *VolumesService) DetachVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachVolumes",
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
func (p *VolumesService) ResizeVolumes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeVolumes",
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
func (p *VolumesService) ModifyVolumeAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVolumeAttributes",
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

func init() { proto.RegisterFile("volume.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xcb, 0xcf, 0x29,
	0xcd, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x48, 0x4d, 0x96, 0x92,
	0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16,
	0x94, 0x54, 0x42, 0x94, 0x28, 0xe9, 0x71, 0x49, 0x84, 0x81, 0xb5, 0x14, 0x07, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x06, 0x14, 0xe5, 0x17, 0xa4, 0x16, 0x95, 0x64, 0xa6, 0x16, 0x0b, 0x09, 0x71,
	0xb1, 0x54, 0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x8f,
	0x98, 0xb9, 0xf8, 0x50, 0x35, 0x08, 0x39, 0x72, 0xf1, 0xbb, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x26,
	0xa5, 0x42, 0x65, 0x84, 0xc4, 0xf4, 0x20, 0x76, 0xea, 0xc1, 0xec, 0xd4, 0x73, 0x05, 0xd9, 0x29,
	0x85, 0x43, 0x5c, 0xc8, 0x9e, 0x8b, 0xd7, 0xb9, 0x28, 0x35, 0xb1, 0x84, 0x12, 0x03, 0x5c, 0x52,
	0x73, 0x52, 0x29, 0x32, 0xc0, 0xb1, 0xa4, 0x24, 0x31, 0x39, 0x83, 0x22, 0x17, 0x50, 0x68, 0x40,
	0x50, 0x6a, 0x71, 0x66, 0x15, 0xd9, 0x5e, 0xf0, 0xe0, 0x12, 0xf3, 0xcd, 0x4f, 0xc9, 0x4c, 0xab,
	0x84, 0x18, 0xe0, 0x58, 0x52, 0x52, 0x94, 0x99, 0x54, 0x5a, 0x42, 0xba, 0x49, 0x49, 0x6c, 0x60,
	0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x56, 0xa6, 0x0b, 0x7b, 0x4e, 0x02, 0x00, 0x00,
}
