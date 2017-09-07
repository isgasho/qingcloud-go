// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongo.proto

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

type DescribeMongosInput struct {
	Limit     int32    `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	MongoName string   `protobuf:"bytes,2,opt,name=mongo_name,json=mongoName" json:"mongo_name,omitempty"`
	Mongos    []string `protobuf:"bytes,3,rep,name=mongos" json:"mongos,omitempty"`
	Offset    int32    `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Status    []string `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	Tags      []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Verbose   int32    `protobuf:"varint,7,opt,name=verbose" json:"verbose,omitempty"`
}

func (m *DescribeMongosInput) Reset()                    { *m = DescribeMongosInput{} }
func (m *DescribeMongosInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeMongosInput) ProtoMessage()               {}
func (*DescribeMongosInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *DescribeMongosInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeMongosInput) GetMongoName() string {
	if m != nil {
		return m.MongoName
	}
	return ""
}

func (m *DescribeMongosInput) GetMongos() []string {
	if m != nil {
		return m.Mongos
	}
	return nil
}

func (m *DescribeMongosInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeMongosInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeMongosInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeMongosInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

type DescribeMongosOutput struct {
	Action     string                        `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                         `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                        `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	TotalCount int32                         `protobuf:"varint,4,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	MongoSet   []*DescribeMongosOutput_Mongo `protobuf:"bytes,5,rep,name=mongo_set,json=mongoSet" json:"mongo_set,omitempty"`
}

func (m *DescribeMongosOutput) Reset()                    { *m = DescribeMongosOutput{} }
func (m *DescribeMongosOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeMongosOutput) ProtoMessage()               {}
func (*DescribeMongosOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *DescribeMongosOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeMongosOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeMongosOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeMongosOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeMongosOutput) GetMongoSet() []*DescribeMongosOutput_Mongo {
	if m != nil {
		return m.MongoSet
	}
	return nil
}

type DescribeMongosOutput_Mongo struct {
	AlarmStatus         string `protobuf:"bytes,1,opt,name=alarm_status,json=alarmStatus" json:"alarm_status,omitempty"`
	AutoBackupTime      int32  `protobuf:"varint,2,opt,name=auto_backup_time,json=autoBackupTime" json:"auto_backup_time,omitempty"`
	AutoMinorVerUpgrade int32  `protobuf:"varint,3,opt,name=auto_minor_ver_upgrade,json=autoMinorVerUpgrade" json:"auto_minor_ver_upgrade,omitempty"`
}

func (m *DescribeMongosOutput_Mongo) Reset()                    { *m = DescribeMongosOutput_Mongo{} }
func (m *DescribeMongosOutput_Mongo) String() string            { return proto.CompactTextString(m) }
func (*DescribeMongosOutput_Mongo) ProtoMessage()               {}
func (*DescribeMongosOutput_Mongo) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2, 0} }

func (m *DescribeMongosOutput_Mongo) GetAlarmStatus() string {
	if m != nil {
		return m.AlarmStatus
	}
	return ""
}

func (m *DescribeMongosOutput_Mongo) GetAutoBackupTime() int32 {
	if m != nil {
		return m.AutoBackupTime
	}
	return 0
}

func (m *DescribeMongosOutput_Mongo) GetAutoMinorVerUpgrade() int32 {
	if m != nil {
		return m.AutoMinorVerUpgrade
	}
	return 0
}

func init() {
	proto.RegisterType((*MongoServiceProperties)(nil), "spec.MongoServiceProperties")
	proto.RegisterType((*DescribeMongosInput)(nil), "spec.DescribeMongosInput")
	proto.RegisterType((*DescribeMongosOutput)(nil), "spec.DescribeMongosOutput")
	proto.RegisterType((*DescribeMongosOutput_Mongo)(nil), "spec.DescribeMongosOutput.Mongo")
}

type MongoServiceInterface interface {
	DescribeMongoNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeMongoParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateMongo(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error)
	DeleteMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateMongoFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ChangeMongoVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RemoveMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyMongoAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	GetMongoMonitor(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
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

func (p *MongoService) DescribeMongoNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoNodes",
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
func (p *MongoService) DescribeMongoParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoParameters",
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
func (p *MongoService) ResizeMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeMongos",
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
func (p *MongoService) CreateMongo(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongo",
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
func (p *MongoService) StopMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopMongos",
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
func (p *MongoService) StartMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartMongos",
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
func (p *MongoService) DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error) {
	if in == nil {
		in = &DescribeMongosInput{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongos",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeMongosOutput{}
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
func (p *MongoService) DeleteMongos(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteMongos",
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
func (p *MongoService) CreateMongoFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongoFromSnapshot",
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
func (p *MongoService) ChangeMongoVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeMongoVxnet",
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
func (p *MongoService) AddMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddMongoInstances",
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
func (p *MongoService) RemoveMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RemoveMongoInstances",
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
func (p *MongoService) ModifyMongoAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoAttributes",
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
func (p *MongoService) ModifyMongoInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoInstances",
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
func (p *MongoService) GetMongoMonitor(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetMongoMonitor",
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

func init() { proto.RegisterFile("mongo.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x4e, 0x1b, 0x31,
	0x10, 0x55, 0x80, 0x04, 0x98, 0x20, 0x4a, 0x0d, 0x0d, 0x4b, 0xaa, 0xaa, 0x94, 0x13, 0x87, 0x6a,
	0x91, 0xe0, 0x5a, 0x2a, 0x41, 0x28, 0x28, 0x87, 0x50, 0xb4, 0x69, 0xb9, 0xae, 0x9c, 0xdd, 0xc9,
	0x62, 0x35, 0xb6, 0x57, 0xf6, 0x6c, 0x54, 0xf8, 0x83, 0xde, 0xfb, 0x41, 0xfd, 0x83, 0xfe, 0x52,
	0x65, 0x3b, 0x91, 0x12, 0x09, 0x0e, 0x49, 0x6f, 0xfb, 0x9e, 0x67, 0x9e, 0xe7, 0x8d, 0xdf, 0x42,
	0x53, 0x6a, 0x55, 0xe8, 0xb8, 0x34, 0x9a, 0x34, 0x5b, 0xb3, 0x25, 0x66, 0xed, 0xb7, 0x85, 0xd6,
	0xc5, 0x08, 0x4f, 0x3c, 0x37, 0xa8, 0x86, 0x27, 0x28, 0x4b, 0x7a, 0x0c, 0x25, 0x47, 0x1f, 0xa1,
	0xd5, 0x73, 0x1d, 0x7d, 0x34, 0x63, 0x91, 0xe1, 0x9d, 0xd1, 0x25, 0x1a, 0x12, 0x68, 0x19, 0x83,
	0xb5, 0x27, 0xad, 0x30, 0xaa, 0x1d, 0xd6, 0x8e, 0x37, 0x13, 0xff, 0x7d, 0xf4, 0xa7, 0x06, 0xbb,
	0x57, 0x68, 0x33, 0x23, 0x06, 0xe8, 0xdb, 0x6c, 0x57, 0x95, 0x15, 0xb1, 0x3d, 0xa8, 0x8f, 0x84,
	0x14, 0xe4, 0x8b, 0xeb, 0x49, 0x00, 0xec, 0x1d, 0x80, 0x9f, 0x26, 0x55, 0x5c, 0x62, 0xb4, 0xe2,
	0x75, 0x36, 0x3d, 0x73, 0xcb, 0x25, 0xb2, 0x16, 0x34, 0x3c, 0xb0, 0xd1, 0xea, 0xe1, 0xea, 0xf1,
	0x66, 0x32, 0x41, 0x8e, 0xd7, 0xc3, 0xa1, 0x45, 0x8a, 0xd6, 0xbc, 0xda, 0x04, 0x39, 0xde, 0x12,
	0xa7, 0xca, 0x46, 0xf5, 0x50, 0x1f, 0x90, 0x1b, 0x94, 0x78, 0x61, 0xa3, 0x86, 0x67, 0xfd, 0x37,
	0x8b, 0x60, 0x7d, 0x8c, 0x66, 0xa0, 0x2d, 0x46, 0xeb, 0x5e, 0x64, 0x0a, 0x8f, 0xfe, 0xae, 0xc0,
	0xde, 0xbc, 0x85, 0xaf, 0x15, 0x39, 0x0f, 0x2d, 0x68, 0xf0, 0x8c, 0x84, 0x56, 0x13, 0xc7, 0x13,
	0xc4, 0x0e, 0x60, 0xc3, 0x20, 0xa5, 0x99, 0xce, 0x83, 0x87, 0x7a, 0xb2, 0x6e, 0x90, 0x3a, 0x3a,
	0x47, 0x77, 0x8b, 0x44, 0x6b, 0x79, 0x81, 0xd1, 0xaa, 0xef, 0x99, 0x42, 0xf6, 0x1e, 0x9a, 0xa4,
	0x89, 0x8f, 0xd2, 0x4c, 0x57, 0x6a, 0x6a, 0x04, 0x3c, 0xd5, 0x71, 0x0c, 0x3b, 0x87, 0xb0, 0x89,
	0xd4, 0xf9, 0x74, 0x7e, 0x9a, 0xa7, 0x87, 0xb1, 0x7b, 0xae, 0xf8, 0xb9, 0xe1, 0x62, 0x0f, 0x92,
	0x0d, 0x19, 0x9e, 0x8a, 0xda, 0xbf, 0x6a, 0x50, 0xf7, 0x1c, 0xfb, 0x00, 0x5b, 0x7c, 0xc4, 0x8d,
	0x4c, 0x27, 0xbb, 0x09, 0xc3, 0x37, 0x3d, 0xd7, 0x0f, 0x0b, 0x3a, 0x86, 0x1d, 0x5e, 0x91, 0x4e,
	0x07, 0x3c, 0xfb, 0x51, 0x95, 0x29, 0x09, 0x39, 0x75, 0xb2, 0xed, 0xf8, 0x4b, 0x4f, 0x7f, 0x13,
	0x12, 0xd9, 0x19, 0xb4, 0x7c, 0xa5, 0x14, 0x4a, 0x9b, 0x74, 0x8c, 0x26, 0xad, 0xca, 0xc2, 0xf0,
	0x3c, 0xf8, 0xab, 0x27, 0xbb, 0xee, 0xb4, 0xe7, 0x0e, 0xef, 0xd1, 0x7c, 0x0f, 0x47, 0xa7, 0xbf,
	0x37, 0x60, 0x6b, 0x36, 0x43, 0xec, 0x0a, 0xd8, 0x9c, 0x89, 0x5b, 0x9d, 0xa3, 0x65, 0xad, 0x38,
	0xe4, 0x30, 0x9e, 0xe6, 0x30, 0xfe, 0xe2, 0x72, 0xd8, 0x7e, 0x81, 0x67, 0x5d, 0xd8, 0x9f, 0x53,
	0xb9, 0xe3, 0x86, 0x4b, 0x24, 0x34, 0x8b, 0x4b, 0x7d, 0x86, 0xad, 0x04, 0xad, 0x78, 0x9a, 0xec,
	0x74, 0xe1, 0xfe, 0x73, 0x68, 0x76, 0x0c, 0x72, 0x0a, 0xfd, 0x0b, 0xb7, 0x7f, 0x02, 0xe8, 0x93,
	0x2e, 0x97, 0xbf, 0xbc, 0x4f, 0xdc, 0xd0, 0x92, 0xed, 0x37, 0xb0, 0x3d, 0x9f, 0x28, 0x76, 0xf0,
	0x5c, 0xce, 0xfc, 0x7f, 0xdc, 0x6e, 0xbf, 0x1c, 0x41, 0xb7, 0xc4, 0x2b, 0x1c, 0x21, 0x2d, 0xbb,
	0xc4, 0x2e, 0xec, 0xcf, 0x2c, 0xf1, 0xda, 0x68, 0xd9, 0x57, 0xbc, 0xb4, 0x0f, 0x9a, 0x16, 0x96,
	0xba, 0x84, 0x9d, 0xce, 0x03, 0x57, 0x45, 0x90, 0xba, 0xff, 0xa9, 0x70, 0x71, 0x8d, 0x0e, 0xbc,
	0xbe, 0xc8, 0x73, 0x2f, 0xd0, 0x55, 0x96, 0xb8, 0xca, 0x96, 0xc8, 0xe8, 0x35, 0xec, 0x25, 0x28,
	0xf5, 0x18, 0xff, 0x53, 0xe7, 0x06, 0xde, 0xf4, 0x74, 0x2e, 0x86, 0x8f, 0x5e, 0xe7, 0x82, 0xc8,
	0x88, 0x41, 0x45, 0xcb, 0x0d, 0x34, 0x23, 0xb4, 0xfc, 0x40, 0x17, 0xf0, 0xea, 0x06, 0x43, 0xe4,
	0x7a, 0x5a, 0x09, 0xd2, 0x66, 0x51, 0x89, 0x41, 0xc3, 0xe3, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0xf8, 0x33, 0x08, 0x92, 0x06, 0x00, 0x00,
}
