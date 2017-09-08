// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdb.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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
func (*RDBServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *RDBServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*RDBServiceProperties)(nil), "spec.RDBServiceProperties")
}

type RDBServiceInterface interface {
	CreateRDB(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RDBsLeaveVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RDBsJoinVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateRDBFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateTempRDBInstanceFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	GetRDBInstanceFiles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CopyRDBInstanceFilesToFTP(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	PurgeRDBLogs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CeaseRDBInstance(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	GetRDBMonitor(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyRDBParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplyRDBParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeRDBParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type RDBService struct {
	Config     *config.Config
	Properties *RDBServiceProperties
}

func NewRDBService(conf *config.Config, zone string) (p *RDBService, err error) {
	return &RDBService{
		Config:     conf,
		Properties: &RDBServiceProperties{Zone: zone},
	}, nil
}

func (p *RDBService) CreateRDB(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDB",
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

func (p *RDBService) DescribeRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBs",
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

func (p *RDBService) DeleteRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRDBs",
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

func (p *RDBService) StartRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartRDBs",
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

func (p *RDBService) StopRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopRDBs",
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

func (p *RDBService) ResizeRDBs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeRDBs",
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

func (p *RDBService) RDBsLeaveVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsLeaveVxnet",
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

func (p *RDBService) RDBsJoinVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsJoinVxnet",
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

func (p *RDBService) CreateRDBFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDBFromSnapshot",
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

func (p *RDBService) CreateTempRDBInstanceFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateTempRDBInstanceFromSnapshot",
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

func (p *RDBService) GetRDBInstanceFiles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBInstanceFiles",
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

func (p *RDBService) CopyRDBInstanceFilesToFTP(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopyRDBInstanceFilesToFTP",
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

func (p *RDBService) PurgeRDBLogs(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PurgeRDBLogs",
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

func (p *RDBService) CeaseRDBInstance(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CeaseRDBInstance",
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

func (p *RDBService) GetRDBMonitor(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBMonitor",
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

func (p *RDBService) ModifyRDBParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRDBParameters",
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

func (p *RDBService) ApplyRDBParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyRDBParameterGroup",
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

func (p *RDBService) DescribeRDBParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBParameters",
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

func init() { proto.RegisterFile("rdb.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x29, 0x94, 0x3f, 0xff, 0x0c, 0x2a, 0x12, 0xb5, 0x68, 0xbd, 0xa8, 0x27, 0xf1, 0x90,
	0x82, 0xde, 0xac, 0xf8, 0xd2, 0xa6, 0xad, 0x2f, 0x2d, 0x84, 0xa4, 0x78, 0xdf, 0xa4, 0xd3, 0xb8,
	0x90, 0x64, 0x96, 0xdd, 0x6d, 0xb1, 0xfd, 0x98, 0x7e, 0x22, 0xd9, 0x06, 0xb5, 0x0a, 0x1e, 0xb2,
	0xde, 0x92, 0x61, 0x9e, 0xdf, 0x3c, 0x79, 0x66, 0x08, 0x38, 0x72, 0x12, 0x7b, 0x42, 0x92, 0x26,
	0xb7, 0xae, 0x04, 0x26, 0xcd, 0xc3, 0x94, 0x28, 0xcd, 0xb0, 0xb5, 0xaa, 0xc5, 0xb3, 0x69, 0x0b,
	0x73, 0xa1, 0x17, 0x65, 0xcb, 0xc9, 0x19, 0xec, 0x86, 0x7e, 0x27, 0x42, 0x39, 0xe7, 0x09, 0x06,
	0x92, 0x04, 0x4a, 0xcd, 0x51, 0xb9, 0x2e, 0xd4, 0x97, 0x54, 0xe0, 0x7e, 0xed, 0xa8, 0x76, 0xea,
	0x84, 0xab, 0xe7, 0xf3, 0x37, 0x07, 0xe0, 0xab, 0xd9, 0x6d, 0x83, 0xd3, 0x95, 0xc8, 0x34, 0x86,
	0x7e, 0xc7, 0x6d, 0x78, 0xe5, 0x14, 0xef, 0x63, 0x8a, 0xd7, 0x33, 0x53, 0x9a, 0xbf, 0xd4, 0xdd,
	0x6b, 0xd8, 0xf0, 0x51, 0x25, 0x92, 0xc7, 0x46, 0xae, 0x2a, 0xeb, 0xaf, 0x00, 0x7c, 0xcc, 0x50,
	0xdb, 0xa9, 0xdb, 0xe0, 0x44, 0x9a, 0x49, 0x6d, 0x25, 0xbe, 0x84, 0xff, 0x91, 0x26, 0x61, 0x6b,
	0x3b, 0x44, 0xc5, 0x97, 0x76, 0xb6, 0x6f, 0x61, 0xcb, 0xe8, 0x86, 0xc8, 0xe6, 0xf8, 0xfc, 0x5a,
	0xa0, 0xae, 0x4c, 0xb8, 0x81, 0x4d, 0x43, 0x78, 0x24, 0x5e, 0xd8, 0x01, 0x06, 0xb0, 0xf7, 0xb9,
	0xf4, 0xbe, 0xa4, 0x3c, 0x2a, 0x98, 0x50, 0x2f, 0x54, 0x1d, 0x14, 0xc1, 0x71, 0x09, 0x1a, 0x63,
	0x6e, 0xb2, 0x7c, 0x28, 0x94, 0x66, 0x45, 0x82, 0x7f, 0x82, 0xf6, 0x60, 0x67, 0x80, 0x7a, 0x9d,
	0xc6, 0x33, 0xac, 0x9e, 0xf3, 0x13, 0x1c, 0x74, 0x49, 0x2c, 0x7e, 0x72, 0xc6, 0xd4, 0x1f, 0x07,
	0x36, 0x97, 0x1e, 0xcc, 0x64, 0x6a, 0x02, 0x1b, 0x52, 0x5a, 0xdd, 0x4c, 0x07, 0xb6, 0xbb, 0xc8,
	0x14, 0xae, 0xb9, 0xb1, 0x59, 0x7b, 0x99, 0xcb, 0x88, 0x0a, 0xae, 0x49, 0xda, 0x04, 0x3b, 0xa2,
	0x09, 0x9f, 0x9a, 0x4c, 0x02, 0x26, 0x59, 0x8e, 0x1a, 0x65, 0xf5, 0x6f, 0xb9, 0x87, 0xc6, 0x9d,
	0x10, 0xd9, 0x37, 0xca, 0x40, 0xd2, 0x4c, 0xd8, 0xdc, 0xe1, 0xda, 0xff, 0xc3, 0xde, 0x52, 0xfc,
	0x6f, 0xf5, 0x7e, 0xf1, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd2, 0xbd, 0x99, 0x37, 0x05, 0x00,
	0x00,
}