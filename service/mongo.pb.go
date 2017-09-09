// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongo.proto

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

type MongoServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *MongoServiceProperties) Reset()                    { *m = MongoServiceProperties{} }
func (m *MongoServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*MongoServiceProperties) ProtoMessage()               {}
func (*MongoServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

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
func (*DescribeMongosInput) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

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
func (*DescribeMongosOutput) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

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
func (*DescribeMongosOutput_Mongo) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2, 0} }

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
	proto.RegisterType((*MongoServiceProperties)(nil), "service.MongoServiceProperties")
	proto.RegisterType((*DescribeMongosInput)(nil), "service.DescribeMongosInput")
	proto.RegisterType((*DescribeMongosOutput)(nil), "service.DescribeMongosOutput")
	proto.RegisterType((*DescribeMongosOutput_Mongo)(nil), "service.DescribeMongosOutput.Mongo")
}

type MongoServiceInterface interface {
	DescribeMongoNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeMongoParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateMongo(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error)
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
	Config           *config.Config
	Properties       *MongoServiceProperties
	LastResponseBody string
}

func NewMongoService(conf *config.Config, zone string) (p *MongoService) {
	return &MongoService{
		Config:     conf,
		Properties: &MongoServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Mongo(zone string) (*MongoService, error) {
	properties := &MongoServiceProperties{
		Zone: zone,
	}

	return &MongoService{Config: s.Config, Properties: properties}, nil
}

func (p *MongoService) DescribeMongoNodes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoNodes",
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

func (p *MongoService) DescribeMongoParameters(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoParameters",
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

func (p *MongoService) ResizeMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeMongos",
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

func (p *MongoService) CreateMongo(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongo",
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

func (p *MongoService) StopMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopMongos",
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

func (p *MongoService) StartMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartMongos",
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

func (p *MongoService) DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error) {
	if in == nil {
		in = &DescribeMongosInput{}
	}
	o := &data.Operation{
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
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeMongosInput) Validate() error {
	return nil
}

func (p *MongoService) DeleteMongos(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteMongos",
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

func (p *MongoService) CreateMongoFromSnapshot(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongoFromSnapshot",
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

func (p *MongoService) ChangeMongoVxnet(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeMongoVxnet",
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

func (p *MongoService) AddMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddMongoInstances",
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

func (p *MongoService) RemoveMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RemoveMongoInstances",
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

func (p *MongoService) ModifyMongoAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoAttributes",
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

func (p *MongoService) ModifyMongoInstances(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoInstances",
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

func (p *MongoService) GetMongoMonitor(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetMongoMonitor",
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

func init() { proto.RegisterFile("mongo.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x4e, 0x13, 0x4d,
	0x14, 0x4e, 0x81, 0xb6, 0xf4, 0x94, 0xf0, 0xf3, 0x0f, 0x58, 0xd6, 0x2a, 0x11, 0xf1, 0x86, 0x0b,
	0x53, 0x12, 0xb8, 0x55, 0x23, 0x14, 0x21, 0xbd, 0x28, 0xe2, 0x56, 0xb9, 0xdd, 0x4c, 0x77, 0x4f,
	0x97, 0x09, 0xdd, 0x99, 0x75, 0xe6, 0x6c, 0x23, 0xbc, 0x81, 0x3e, 0x8e, 0x4f, 0xe0, 0x33, 0xf8,
	0x26, 0xbe, 0x81, 0x99, 0x99, 0x6d, 0x42, 0x13, 0x31, 0x69, 0xbd, 0x69, 0xf6, 0xfb, 0xe6, 0x9c,
	0x6f, 0xcf, 0x77, 0xe6, 0x6b, 0x0b, 0xcd, 0x4c, 0xc9, 0x54, 0x75, 0x72, 0xad, 0x48, 0xb1, 0xba,
	0x41, 0x3d, 0x11, 0x31, 0xb6, 0x77, 0x3e, 0x0b, 0x99, 0xc6, 0x63, 0x55, 0x24, 0x91, 0x49, 0x6e,
	0x22, 0x5d, 0x8c, 0xf1, 0xc0, 0x7e, 0xf8, 0xba, 0xf6, 0x93, 0x54, 0xa9, 0x74, 0x8c, 0x07, 0x0e,
	0x0d, 0x8b, 0xd1, 0x01, 0x66, 0x39, 0xdd, 0xfa, 0xc3, 0xbd, 0x97, 0xd0, 0xea, 0x5b, 0xcd, 0x81,
	0xd7, 0xba, 0xd4, 0x2a, 0x47, 0x4d, 0x02, 0x0d, 0x63, 0xb0, 0x72, 0xa7, 0x24, 0x06, 0x95, 0xdd,
	0xca, 0x7e, 0x23, 0x74, 0xcf, 0x7b, 0x3f, 0x2a, 0xb0, 0x79, 0x8a, 0x26, 0xd6, 0x62, 0x88, 0xae,
	0xcd, 0xf4, 0x64, 0x5e, 0x10, 0xdb, 0x82, 0xea, 0x58, 0x64, 0x82, 0x5c, 0x71, 0x35, 0xf4, 0x80,
	0xed, 0x00, 0xb8, 0x79, 0x23, 0xc9, 0x33, 0x0c, 0x96, 0x9c, 0x4e, 0xc3, 0x31, 0x17, 0x3c, 0x43,
	0xd6, 0x82, 0x9a, 0x03, 0x26, 0x58, 0xde, 0x5d, 0xde, 0x6f, 0x84, 0x25, 0xb2, 0xbc, 0x1a, 0x8d,
	0x0c, 0x52, 0xb0, 0xe2, 0xd4, 0x4a, 0x64, 0x79, 0x43, 0x9c, 0x0a, 0x13, 0x54, 0x7d, 0xbd, 0x47,
	0x76, 0x50, 0xe2, 0xa9, 0x09, 0x6a, 0x8e, 0x75, 0xcf, 0x2c, 0x80, 0xfa, 0x04, 0xf5, 0x50, 0x19,
	0x0c, 0xea, 0x4e, 0x64, 0x0a, 0xf7, 0x7e, 0x2e, 0xc1, 0xd6, 0xac, 0x85, 0xf7, 0x05, 0x59, 0x0f,
	0x2d, 0xa8, 0xf1, 0x98, 0x84, 0x92, 0xa5, 0xe3, 0x12, 0xb1, 0xc7, 0xb0, 0xaa, 0x91, 0xa2, 0x58,
	0x25, 0xde, 0x43, 0x35, 0xac, 0x6b, 0xa4, 0xae, 0x4a, 0xd0, 0xbe, 0x25, 0x43, 0x63, 0x78, 0x8a,
	0xc1, 0xb2, 0xeb, 0x99, 0x42, 0xf6, 0x0c, 0x9a, 0xa4, 0x88, 0x8f, 0xa3, 0x58, 0x15, 0x72, 0x6a,
	0x04, 0x1c, 0xd5, 0xb5, 0x0c, 0x7b, 0x0b, 0x7e, 0x13, 0x91, 0xf5, 0x69, 0xfd, 0x34, 0x0f, 0x5f,
	0x74, 0xca, 0x0b, 0xed, 0xfc, 0x69, 0xbe, 0x8e, 0x03, 0xe1, 0x6a, 0xe6, 0x6f, 0x8b, 0xda, 0x5f,
	0x2b, 0x50, 0x75, 0x1c, 0x7b, 0x0e, 0x6b, 0x7c, 0xcc, 0x75, 0x16, 0x95, 0xeb, 0xf1, 0xf3, 0x37,
	0x1d, 0x37, 0xf0, 0x3b, 0xda, 0x87, 0x0d, 0x5e, 0x90, 0x8a, 0x86, 0x3c, 0xbe, 0x29, 0xf2, 0x88,
	0x44, 0x36, 0x35, 0xb3, 0x6e, 0xf9, 0x13, 0x47, 0x7f, 0x14, 0x19, 0xb2, 0x23, 0x68, 0xb9, 0xca,
	0x4c, 0x48, 0xa5, 0xa3, 0x09, 0xea, 0xa8, 0xc8, 0x53, 0xcd, 0x13, 0x6f, 0xb1, 0x1a, 0x6e, 0xda,
	0xd3, 0xbe, 0x3d, 0xbc, 0x42, 0xfd, 0xc9, 0x1f, 0x1d, 0x7e, 0x5f, 0x85, 0xb5, 0xfb, 0x31, 0x62,
	0xa7, 0xc0, 0x66, 0x4c, 0x5c, 0xa8, 0x04, 0x0d, 0x6b, 0x75, 0x7c, 0x14, 0x3b, 0xd3, 0x28, 0x76,
	0xde, 0xd9, 0x28, 0xb6, 0x1f, 0xe0, 0x59, 0x0f, 0xb6, 0x67, 0x54, 0x2e, 0xb9, 0xe6, 0x19, 0x12,
	0xea, 0xf9, 0xa5, 0xde, 0xc0, 0x5a, 0x88, 0x46, 0xdc, 0x95, 0x3b, 0x9d, 0xbb, 0xff, 0x35, 0x34,
	0xbb, 0x1a, 0x39, 0xf9, 0xfe, 0xb9, 0xdb, 0x5f, 0x01, 0x0c, 0x48, 0xe5, 0x8b, 0xbf, 0x7c, 0x40,
	0x5c, 0xd3, 0x82, 0xed, 0x7d, 0x58, 0x9f, 0x4d, 0x14, 0x7b, 0xfa, 0x40, 0xd4, 0xdc, 0xb7, 0xb9,
	0xbd, 0xf3, 0xd7, 0x20, 0xda, 0x55, 0x9e, 0xe2, 0x18, 0x69, 0xd1, 0x55, 0xf6, 0x60, 0xfb, 0xde,
	0x2a, 0xcf, 0xb4, 0xca, 0x06, 0x92, 0xe7, 0xe6, 0x5a, 0xd1, 0xdc, 0x52, 0x27, 0xb0, 0xd1, 0xbd,
	0xe6, 0x32, 0xf5, 0x52, 0x57, 0x5f, 0x24, 0xce, 0xaf, 0xd1, 0x85, 0xff, 0x8f, 0x93, 0xc4, 0x09,
	0xf4, 0xa4, 0x21, 0x2e, 0xe3, 0x05, 0x92, 0x7a, 0x06, 0x5b, 0x21, 0x66, 0x6a, 0x82, 0xff, 0xa8,
	0x73, 0x0e, 0x8f, 0xfa, 0x2a, 0x11, 0xa3, 0x5b, 0xa7, 0x73, 0x4c, 0xa4, 0xc5, 0xb0, 0xa0, 0xc5,
	0x06, 0xba, 0x27, 0xb4, 0xf8, 0x40, 0xc7, 0xf0, 0xdf, 0x39, 0xfa, 0xe0, 0xf5, 0x95, 0x14, 0xa4,
	0xf4, 0xbc, 0x12, 0xed, 0xd6, 0xb7, 0x5f, 0x2b, 0x0c, 0x1a, 0x1f, 0x84, 0x4c, 0xbb, 0xf6, 0x4f,
	0x8a, 0xf9, 0x9f, 0xad, 0x61, 0xcd, 0xd5, 0x1d, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x94, 0x62,
	0x83, 0xdd, 0xd5, 0x06, 0x00, 0x00,
}
