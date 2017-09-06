// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongo.proto

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

type MongoServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *MongoServiceProperties) Reset()                    { *m = MongoServiceProperties{} }
func (m *MongoServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*MongoServiceProperties) ProtoMessage()               {}
func (*MongoServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *MongoServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*MongoServiceProperties)(nil), "spec.MongoServiceProperties")
}

type MongoServiceInterface interface {
	DescribeMongoNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeMongoParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateMongo(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateMongoFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ChangeMongoVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RemoveMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyMongoAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GetMongoMonitor(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type MongoService struct {
	Config     *config.Config
	Properties *MongoServiceProperties
}

func NewMongoService(conf *config.Config, zone string) (p *MongoService, err error) {
	return &MongoService{
		Config:     conf,
		Properties: &MongoServiceProperties{Zone: zone},
	}, nil
}

func (p *MongoService) DescribeMongoNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) DescribeMongoParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) ResizeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) CreateMongo(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) StopMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) StartMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) DescribeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) DeleteMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) CreateMongoFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) ChangeMongoVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) AddMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) RemoveMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) ModifyMongoAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) ModifyMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *MongoService) GetMongoMonitor(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("mongo.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xdf, 0x4b, 0x3a, 0x41,
	0x14, 0xc5, 0x11, 0xe4, 0xcb, 0xb7, 0xab, 0xf4, 0x63, 0x29, 0x2d, 0x83, 0x8a, 0x9e, 0x7a, 0x88,
	0x15, 0xea, 0xb5, 0x22, 0xd3, 0x14, 0x1f, 0x0c, 0x71, 0xa1, 0xf7, 0xd9, 0xdd, 0xeb, 0x3a, 0xe0,
	0xce, 0x1d, 0x66, 0xae, 0x96, 0xfe, 0x0f, 0xfd, 0xcf, 0xe1, 0x6e, 0x82, 0x0a, 0x3d, 0xec, 0xf4,
	0xb6, 0x7b, 0xcf, 0x39, 0x9f, 0x7b, 0xb8, 0x0c, 0x54, 0x52, 0x52, 0x09, 0xf9, 0xda, 0x10, 0x93,
	0x57, 0xb6, 0x1a, 0xa3, 0xc6, 0x59, 0x42, 0x94, 0x4c, 0xb1, 0x99, 0xcd, 0xc2, 0xd9, 0xb8, 0x29,
	0xd4, 0x22, 0x37, 0x34, 0xce, 0x77, 0x25, 0x4c, 0x35, 0xaf, 0xc5, 0xcb, 0x5d, 0x91, 0x65, 0x8a,
	0x96, 0x45, 0xaa, 0x7f, 0x0c, 0x17, 0xbb, 0x86, 0x0f, 0x23, 0xb4, 0x46, 0x63, 0x73, 0xfd, 0xfa,
	0x16, 0x6a, 0x83, 0x55, 0x9b, 0x00, 0xcd, 0x5c, 0x46, 0x38, 0x34, 0xa4, 0xd1, 0xb0, 0x44, 0xeb,
	0x79, 0x50, 0x5e, 0x92, 0xc2, 0xd3, 0xd2, 0x55, 0xe9, 0x66, 0x6f, 0x94, 0x7d, 0xdf, 0x7d, 0xfd,
	0x87, 0xea, 0xa6, 0xdd, 0xeb, 0x80, 0xd7, 0x41, 0x1b, 0x19, 0x19, 0x62, 0x36, 0x7f, 0xa3, 0x18,
	0xad, 0x57, 0xf3, 0xf3, 0xad, 0xfe, 0x7a, 0xab, 0xff, 0xba, 0xea, 0xdc, 0xf8, 0x65, 0xee, 0xf5,
	0xa1, 0xbe, 0x45, 0x19, 0x0a, 0x23, 0x52, 0x64, 0x34, 0xc5, 0x51, 0x4f, 0x50, 0x1d, 0xa1, 0x95,
	0xcb, 0x1c, 0x54, 0x3c, 0xff, 0x08, 0x95, 0xb6, 0x41, 0xc1, 0x79, 0xbe, 0x70, 0xfc, 0x01, 0x20,
	0x60, 0xd2, 0xee, 0xcb, 0x03, 0x16, 0x86, 0x1d, 0xe3, 0xcf, 0xb0, 0xbf, 0x75, 0x46, 0xa7, 0xeb,
	0x75, 0x70, 0x8a, 0xec, 0x9a, 0xef, 0x43, 0x7d, 0xe3, 0x7a, 0x5d, 0x43, 0x69, 0xa0, 0x84, 0xb6,
	0x13, 0xe2, 0xc2, 0xa8, 0x17, 0x38, 0x6c, 0x4f, 0x84, 0x4a, 0x72, 0xd4, 0xfb, 0xa7, 0xc2, 0xe2,
	0x8c, 0x36, 0x1c, 0xb5, 0xe2, 0x38, 0x03, 0xf4, 0x95, 0x65, 0xa1, 0x22, 0x87, 0xc7, 0xd9, 0x85,
	0xe3, 0x11, 0xa6, 0x34, 0xc7, 0x3f, 0x72, 0x7a, 0x70, 0x32, 0xa0, 0x58, 0x8e, 0x17, 0x19, 0xa7,
	0xc5, 0x6c, 0x64, 0x38, 0x63, 0xb7, 0x42, 0x1b, 0x20, 0xf7, 0x42, 0x2d, 0x38, 0xe8, 0x61, 0xfe,
	0xd6, 0x06, 0xa4, 0x24, 0x93, 0x29, 0x8a, 0x08, 0xff, 0x65, 0xff, 0xf7, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x81, 0x0b, 0x2f, 0xd2, 0x04, 0x00, 0x00,
}
