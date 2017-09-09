// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdb.proto

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

type RDBServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *RDBServiceProperties) Reset()                    { *m = RDBServiceProperties{} }
func (m *RDBServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*RDBServiceProperties) ProtoMessage()               {}
func (*RDBServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *RDBServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*RDBServiceProperties)(nil), "service.RDBServiceProperties")
}

type RDBServiceInterface interface {
	CreateRDB(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RDBsLeaveVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RDBsJoinVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateRDBFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateTempRDBInstanceFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GetRDBInstanceFiles(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CopyRDBInstanceFilesToFTP(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	PurgeRDBLogs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CeaseRDBInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GetRDBMonitor(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyRDBParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ApplyRDBParameterGroup(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeRDBParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type RDBService struct {
	Config           *config.Config
	Properties       *RDBServiceProperties
	LastResponseBody string
}

func NewRDBService(conf *config.Config, zone string) (p *RDBService) {
	return &RDBService{
		Config:     conf,
		Properties: &RDBServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) RDB(zone string) (*RDBService, error) {
	properties := &RDBServiceProperties{
		Zone: zone,
	}

	return &RDBService{Config: s.Config, Properties: properties}, nil
}

func (p *RDBService) CreateRDB(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDB",
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

func (p *RDBService) DescribeRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBs",
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

func (p *RDBService) DeleteRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRDBs",
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

func (p *RDBService) StartRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartRDBs",
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

func (p *RDBService) StopRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopRDBs",
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

func (p *RDBService) ResizeRDBs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeRDBs",
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

func (p *RDBService) RDBsLeaveVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsLeaveVxnet",
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

func (p *RDBService) RDBsJoinVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsJoinVxnet",
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

func (p *RDBService) CreateRDBFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDBFromSnapshot",
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

func (p *RDBService) CreateTempRDBInstanceFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateTempRDBInstanceFromSnapshot",
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

func (p *RDBService) GetRDBInstanceFiles(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBInstanceFiles",
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

func (p *RDBService) CopyRDBInstanceFilesToFTP(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopyRDBInstanceFilesToFTP",
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

func (p *RDBService) PurgeRDBLogs(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PurgeRDBLogs",
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

func (p *RDBService) CeaseRDBInstance(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CeaseRDBInstance",
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

func (p *RDBService) GetRDBMonitor(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBMonitor",
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

func (p *RDBService) ModifyRDBParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRDBParameters",
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

func (p *RDBService) ApplyRDBParameterGroup(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyRDBParameterGroup",
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

func (p *RDBService) DescribeRDBParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBParameters",
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

func init() { proto.RegisterFile("rdb.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x86, 0xb9, 0x78, 0x51, 0x73, 0x50, 0xb9, 0xc4, 0x6b, 0xd1, 0x88, 0xa0, 0xae, 0xc4, 0x45,
	0x0a, 0xba, 0xb3, 0xe2, 0x47, 0x92, 0xb6, 0x7e, 0xb4, 0x10, 0x93, 0xe2, 0xb6, 0x4c, 0x92, 0xd3,
	0x38, 0x98, 0xcc, 0x19, 0x67, 0x26, 0xc5, 0xf6, 0x27, 0xf8, 0x4b, 0xdd, 0xfa, 0x0f, 0x64, 0x1a,
	0xd4, 0x2a, 0xb8, 0xc8, 0xdc, 0x4d, 0xc8, 0xc7, 0x79, 0x9f, 0xf3, 0xe6, 0x9d, 0xc3, 0x01, 0x4f,
	0x55, 0x45, 0x28, 0x15, 0x19, 0xf2, 0xaf, 0x68, 0x54, 0x5b, 0x5e, 0x62, 0x70, 0xef, 0x0b, 0x17,
	0x75, 0xd9, 0x50, 0x57, 0xad, 0x75, 0xf5, 0x79, 0xad, 0xba, 0x06, 0xc7, 0xf6, 0xd2, 0xd7, 0x05,
	0x77, 0x6b, 0xa2, 0xba, 0xc1, 0xf1, 0xe1, 0xa9, 0xe8, 0x36, 0x63, 0x6c, 0xa5, 0xd9, 0xf5, 0x1f,
	0x1f, 0x3e, 0x86, 0xf3, 0x2c, 0x89, 0xf2, 0x9e, 0x94, 0x2a, 0x92, 0xa8, 0x0c, 0x47, 0xed, 0xfb,
	0x70, 0xba, 0x27, 0x81, 0xb7, 0x4f, 0xee, 0x9f, 0x3c, 0xf2, 0xb2, 0xc3, 0xfd, 0x93, 0xef, 0x1e,
	0xc0, 0x9f, 0x62, 0x7f, 0x02, 0x5e, 0xac, 0x90, 0x19, 0xcc, 0x92, 0xc8, 0x1f, 0x85, 0x7d, 0x97,
	0xf0, 0x57, 0x97, 0x70, 0x6a, 0xbb, 0x04, 0xff, 0x79, 0xef, 0xbf, 0x80, 0x6b, 0x09, 0xea, 0x52,
	0xf1, 0xc2, 0xca, 0xf5, 0x60, 0xfd, 0x73, 0x80, 0x04, 0x1b, 0x34, 0x6e, 0xea, 0x09, 0x78, 0xb9,
	0x61, 0xca, 0x38, 0x89, 0x9f, 0xc1, 0xd5, 0xdc, 0x90, 0x74, 0xb5, 0x9d, 0xa1, 0xe6, 0x7b, 0x37,
	0xdb, 0xaf, 0xe0, 0x86, 0xd5, 0x2d, 0x90, 0x6d, 0xf1, 0xe3, 0x57, 0x81, 0x66, 0x30, 0xe1, 0x25,
	0x5c, 0xb7, 0x84, 0x77, 0xc4, 0x85, 0x1b, 0x60, 0x0e, 0xb7, 0x7e, 0x1f, 0xfa, 0x4c, 0x51, 0x9b,
	0x0b, 0x26, 0xf5, 0x27, 0x1a, 0x0e, 0xca, 0xe1, 0x41, 0x0f, 0x5a, 0x61, 0x6b, 0xb3, 0x7c, 0x2b,
	0xb4, 0x61, 0xa2, 0xc4, 0x0b, 0x41, 0xa7, 0x70, 0x73, 0x8e, 0xe6, 0x98, 0xc6, 0x1b, 0x1c, 0x9e,
	0xf3, 0x7b, 0xb8, 0x13, 0x93, 0xdc, 0xfd, 0xcb, 0x59, 0xd1, 0x6c, 0x95, 0xba, 0x4c, 0x7a, 0xda,
	0xa9, 0xda, 0x06, 0xb6, 0xa0, 0x7a, 0xb8, 0x99, 0x08, 0xce, 0x62, 0x64, 0x1a, 0x8f, 0xdc, 0xb8,
	0x1c, 0x7b, 0x9f, 0xcb, 0x92, 0x04, 0x37, 0xa4, 0x5c, 0x82, 0x5d, 0x52, 0xc5, 0x37, 0x36, 0x93,
	0x94, 0x29, 0xd6, 0xa2, 0x41, 0x35, 0xfc, 0x5f, 0xde, 0xc0, 0xe8, 0xb5, 0x94, 0xcd, 0x5f, 0x94,
	0xb9, 0xa2, 0x4e, 0xba, 0xcc, 0xe1, 0xd1, 0xfe, 0x70, 0xb7, 0x14, 0x9c, 0x7f, 0xfb, 0x71, 0x7a,
	0x06, 0xde, 0x07, 0x2e, 0xea, 0xd8, 0xae, 0x50, 0xff, 0x52, 0x96, 0x44, 0xc5, 0xe5, 0x43, 0xd5,
	0xd3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xdc, 0x34, 0x63, 0x6f, 0x05, 0x00, 0x00,
}
