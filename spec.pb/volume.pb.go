// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volume.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

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
	DescribeVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AttachVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DetachVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyVolumeAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
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

func (p *VolumesService) DescribeVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) CreateVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) DeleteVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) AttachVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) DetachVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) ResizeVolumes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *VolumesService) ModifyVolumeAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x29, 0x16, 0xc1, 0xe0, 0x0f, 0x64, 0x51, 0xea, 0x08, 0x2a, 0xae, 0x5c, 0xa5, 0xa0,
	0x0f, 0x20, 0x83, 0x15, 0xdc, 0x08, 0x52, 0xc1, 0x7d, 0x66, 0x3c, 0xad, 0x81, 0x99, 0xde, 0x70,
	0x73, 0xa7, 0xd2, 0x3e, 0xae, 0x4f, 0x22, 0x26, 0xed, 0xc2, 0x01, 0x17, 0x9a, 0x5d, 0x72, 0xbf,
	0x93, 0xef, 0x9e, 0xa8, 0xc3, 0x15, 0x35, 0x5d, 0x0b, 0xe3, 0x99, 0x84, 0xf4, 0x30, 0x78, 0xd4,
	0xc5, 0xe9, 0x82, 0x68, 0xd1, 0x60, 0x12, 0x67, 0x55, 0x37, 0x9f, 0xd8, 0xe5, 0x3a, 0x05, 0x8a,
	0xb3, 0x3e, 0x42, 0xeb, 0x65, 0x07, 0x2f, 0xfa, 0x50, 0x5c, 0x8b, 0x20, 0xb6, 0xf5, 0xdb, 0xc0,
	0x79, 0x3f, 0xf0, 0xc1, 0xd6, 0x7b, 0x70, 0x48, 0xfc, 0xca, 0xa8, 0xf1, 0x6b, 0xac, 0x13, 0x5e,
	0xc0, 0x2b, 0x57, 0xe3, 0x99, 0xc9, 0x83, 0xc5, 0x21, 0x68, 0xad, 0x86, 0x1b, 0x5a, 0x62, 0x3c,
	0xb8, 0x1c, 0x5c, 0x1f, 0xcc, 0xe2, 0xf9, 0xe6, 0x73, 0x4f, 0x1d, 0xff, 0x7c, 0xa0, 0x4b, 0x75,
	0x32, 0x45, 0xa8, 0xd9, 0x55, 0xd8, 0x12, 0x3d, 0x32, 0x69, 0xad, 0xd9, 0xad, 0x35, 0x0f, 0xdf,
	0xa5, 0x8b, 0x5f, 0xe6, 0xfa, 0x4e, 0x1d, 0xdd, 0x33, 0xac, 0xe4, 0x08, 0xa6, 0x68, 0x90, 0x25,
	0x28, 0x45, 0x6c, 0xfd, 0x9e, 0xd5, 0x20, 0x53, 0x30, 0x43, 0x70, 0x9b, 0x7f, 0x7f, 0xe1, 0x51,
	0x8d, 0x9e, 0xe8, 0xcd, 0xcd, 0xd7, 0x49, 0x50, 0x8a, 0xb0, 0xab, 0x3a, 0xf9, 0xbb, 0xa9, 0xda,
	0x8f, 0xf7, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xf4, 0x68, 0x68, 0xaa, 0x02, 0x00,
	0x00,
}
