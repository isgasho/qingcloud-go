// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm.proto

/*
Package spec is a generated protocol buffer package.

It is generated from these files:
	alarm.proto
	base.proto
	cache.proto
	cluster.proto
	dns.proto
	eip.proto
	hadoop.proto
	image.proto
	instances.proto
	job.proto
	key_pair.proto
	load_balancer.proto
	misc.proto
	mongo.proto
	monitor.proto
	nic.proto
	notification_center.proto
	rdb.proto
	resource_acl.proto
	router.proto
	s2.proto
	security_group.proto
	snapshot.proto
	span.proto
	spark.proto
	subuser.proto
	tag.proto
	user_data.proto
	volume.proto
	vxnet.proto
	zone.proto

It has these top-level messages:
	AlarmServiceProperties
	Options
	CommonRequest
	Error
	CacheServiceProperties
	ClusterServiceProperties
	DNSAliasServiceProperties
	EipServiceProperties
	HadoopServiceProperties
	ImageServiceProperties
	InstanceServiceProperties
	DescribeInstanceTypes_Request
	DescribeInstanceTypes_Reponse
	InstanceTypeSetElem
	JobServiceProperties
	KeyPairServiceProperties
	LoadBalancerServiceProperties
	MiscServiceProperties
	MongoServiceProperties
	MonitorServiceProperties
	NicServiceProperties
	CreateNicsInput
	CreateNicsOutput
	DescribeNicsInput
	DescribeNicsOutput
	NIC
	Eip
	Tag
	ResourceTagPair
	ResourceTypeCount
	AttachNicsInput
	AttachNicsOutput
	DetachNicsInput
	DetachNicsOutput
	ModifyNicAttributesInput
	ModifyNicAttributesOutput
	DeleteNicsInput
	DeleteNicsOutput
	NotificationCenterServiceProperties
	RDBServiceProperties
	ResourceACLServiceProperties
	RouterServiceProperties
	S2ServiceProperties
	SecurityGroupServiceProperties
	SnapshotServiceProperties
	SpanServiceProperties
	SparkServiceProperties
	SubuserServiceProperties
	TagServiceProperties
	UserDataServiceProperties
	VolumesServiceProperties
	VxnetServiceProperties
	ZoneServiceProperties
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AlarmServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *AlarmServiceProperties) Reset()                    { *m = AlarmServiceProperties{} }
func (m *AlarmServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*AlarmServiceProperties) ProtoMessage()               {}
func (*AlarmServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlarmServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*AlarmServiceProperties)(nil), "spec.AlarmServiceProperties")
}

type AlarmServiceInterface interface {
	DescribeAlarmPolicies(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicies(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyRuleAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyActionAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AssociateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DissociateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ApplyAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarms(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmHistory(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type AlarmService struct {
	Config     *config.Config
	Properties *AlarmServiceProperties
}

func NewAlarmService(conf *config.Config, zone string) (p *AlarmService, err error) {
	return &AlarmService{
		Config:     conf,
		Properties: &AlarmServiceProperties{Zone: zone},
	}, nil
}

func (p *AlarmService) DescribeAlarmPolicies(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) CreateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) ModifyAlarmPolicyAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DeleteAlarmPolicies(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DescribeAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) AddAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) ModifyAlarmPolicyRuleAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyRuleAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DeleteAlarmPolicyRules(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DescribeAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) AddAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) ModifyAlarmPolicyActionAttributes(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyActionAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DeleteAlarmPolicyActions(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) AssociateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DissociateAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) ApplyAlarmPolicy(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DescribeAlarms(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarms",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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
func (p *AlarmService) DescribeAlarmHistory(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	if in == nil {
		in = &google_protobuf1.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmHistory",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf1.Empty{}
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

func init() { proto.RegisterFile("alarm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x4b, 0x02, 0x41,
	0x10, 0xc7, 0x11, 0x24, 0x70, 0x8a, 0xa8, 0x2b, 0xc5, 0x4e, 0xc8, 0xea, 0xa9, 0x87, 0x38, 0xa1,
	0xbe, 0x40, 0x97, 0x9a, 0x12, 0x09, 0xa6, 0x9f, 0xe0, 0xee, 0x1c, 0x65, 0x61, 0xcf, 0x5d, 0x76,
	0xc7, 0xe2, 0xfa, 0xe6, 0xbd, 0xc5, 0xad, 0x09, 0x7a, 0xd6, 0xc3, 0xed, 0xf6, 0x76, 0x37, 0x33,
	0xf7, 0xdb, 0xff, 0xfd, 0x98, 0x85, 0xc3, 0x88, 0x47, 0x2a, 0x0d, 0xa4, 0x12, 0x24, 0xbc, 0xaa,
	0x96, 0x98, 0xf8, 0x17, 0x0b, 0x21, 0x16, 0x1c, 0x3b, 0xa6, 0x16, 0xaf, 0xe6, 0x9d, 0x68, 0x99,
	0xad, 0x07, 0xfc, 0x56, 0xb1, 0x85, 0xa9, 0xa4, 0x4d, 0xb3, 0x5d, 0x6c, 0x12, 0x4b, 0x51, 0x53,
	0x94, 0xca, 0x9f, 0x81, 0xcb, 0xe2, 0xc0, 0x87, 0x8a, 0xa4, 0x44, 0xa5, 0xd7, 0xfd, 0x9b, 0x3b,
	0x68, 0x84, 0x79, 0x9a, 0x29, 0xaa, 0x77, 0x96, 0xe0, 0x58, 0x09, 0x89, 0x8a, 0x18, 0x6a, 0xcf,
	0x83, 0xea, 0xa7, 0x58, 0x62, 0xb3, 0x72, 0x55, 0xb9, 0xad, 0x4d, 0xcc, 0xf3, 0xfd, 0x57, 0x0d,
	0x8e, 0xb6, 0xc7, 0xbd, 0x01, 0xd4, 0x7b, 0xa8, 0x13, 0xc5, 0x62, 0x34, 0xf5, 0xb1, 0xe0, 0x2c,
	0xc9, 0xbf, 0x6e, 0x04, 0xeb, 0x83, 0x83, 0xcd, 0xc1, 0x41, 0x3f, 0x8f, 0xed, 0xff, 0x51, 0xf7,
	0xba, 0x70, 0xda, 0x55, 0x18, 0xd1, 0x16, 0x26, 0x2b, 0x0d, 0x19, 0x41, 0x6b, 0x24, 0x66, 0x6c,
	0x9e, 0x6d, 0x41, 0x42, 0x22, 0xc5, 0xe2, 0x15, 0x59, 0x64, 0xea, 0xc3, 0x59, 0x0f, 0x39, 0x92,
	0xe3, 0xaf, 0xbd, 0x40, 0x73, 0xdf, 0x51, 0x36, 0x59, 0x71, 0xbb, 0x48, 0xe1, 0x6c, 0xe6, 0x8c,
	0x79, 0x83, 0xf6, 0x9e, 0xa8, 0x9c, 0xe4, 0x20, 0x6b, 0x08, 0x8d, 0xa2, 0x2c, 0xcb, 0x70, 0xaf,
	0xe0, 0xff, 0xe2, 0x2b, 0x4c, 0x88, 0x89, 0x65, 0x79, 0xda, 0x00, 0xea, 0xbb, 0xc6, 0x6c, 0x41,
	0x53, 0xb8, 0xde, 0x5f, 0x2e, 0xc3, 0x72, 0xb0, 0x66, 0x76, 0xa3, 0x60, 0xcd, 0x36, 0xe0, 0x33,
	0x9c, 0x87, 0x5a, 0x8b, 0x84, 0x39, 0xde, 0xa2, 0xfc, 0x4e, 0xb3, 0xff, 0x00, 0x3d, 0xc1, 0x49,
	0x28, 0x25, 0xcf, 0x5c, 0x18, 0x8f, 0x70, 0xbc, 0xb3, 0x0c, 0x56, 0x5a, 0x76, 0x08, 0x43, 0xa6,
	0x49, 0xa8, 0xd2, 0x49, 0xe2, 0x03, 0xf3, 0xfe, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x33, 0x18,
	0x7c, 0x3c, 0xbe, 0x05, 0x00, 0x00,
}
