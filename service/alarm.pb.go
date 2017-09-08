// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	alarm.proto
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
	qingcloud.proto
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
	types.proto
	user_data.proto
	volume.proto
	vxnet.proto
	zone.proto

It has these top-level messages:
	AlarmServiceProperties
	CacheServiceProperties
	ClusterServiceProperties
	CreateClusterInput
	CreateClusterOutput
	DescribeClustersInput
	DescribeClustersOutput
	DescribeClusterNodesInput
	DescribeClusterNodesOutput
	StopClustersInput
	StopClustersOutput
	StartClustersInput
	StartClustersOutput
	DeleteClustersInput
	DeleteClustersOutput
	LeaseInput
	LeaseOutput
	AddClusterNodesInput
	AddClusterNodesOutput
	DeleteClusterNodesInput
	DeleteClusterNodesOutput
	ResizeClusterInput
	ResizeClusterOutput
	ChangeClusterVxnetInput
	ChangeClusterVxnetOutput
	SuspendClustersInput
	SuspendClustersOutput
	UpdateClusterEnvironmentInput
	UpdateClusterEnvironmentOutput
	ModifyClusterAttributesInput
	ModifyClusterAttributesOutput
	ModifyClusterNodeAttributesInput
	ModifyClusterNodeAttributesOutput
	GetClustersStatsInput
	GetClustersStatsOutput
	DescribeClusterUsersInput
	DescribeClusterUsersOutput
	RestartClusterServiceInput
	RestartClusterServiceOutput
	UpgradeClustersInput
	UpgradeClustersOutput
	AuthorizeClustersBrokerToDeveloperInput
	AuthorizeClustersBrokerToDeveloperOutput
	RevokeClustersBrokerFromDeveloperInput
	RevokeClustersBrokerFromDeveloperOutput
	DNSAliasServiceProperties
	DescribeDNSAliasesInput
	DescribeDNSAliasesOutput
	AssociateDNSAliasInput
	AssociateDNSAliasOutput
	DissociateDNSAliasesInput
	DissociateDNSAliasesOutput
	GetDNSLabelInput
	GetDNSLabelOutput
	EipServiceProperties
	HadoopServiceProperties
	ImageServiceProperties
	InstanceServiceProperties
	DescribeInstanceTypesInput
	DescribeInstanceTypesOutput
	JobServiceProperties
	DescribeJobsInput
	DescribeJobsOutput
	KeyPairServiceProperties
	LoadBalancerServiceProperties
	MiscServiceProperties
	GrantQuotaIndepInput
	GrantQuotaIndepOutput
	RevokeQuotaIndepInput
	RevokeQuotaIndepOutput
	GetQuotaLeftInput
	GetQuotaLeftOutput
	MongoServiceProperties
	DescribeMongosInput
	DescribeMongosOutput
	MonitorServiceProperties
	NicServiceProperties
	CreateNicsInput
	CreateNicsOutput
	DescribeNicsInput
	DescribeNicsOutput
	AttachNicsInput
	AttachNicsOutput
	DetachNicsInput
	DetachNicsOutput
	ModifyNicAttributesInput
	ModifyNicAttributesOutput
	DeleteNicsInput
	DeleteNicsOutput
	NotificationCenterServiceProperties
	DescribeNotificationCenterUserPostsInput
	DescribeNotificationCenterUserPostsOutput
	QingCloudServiceProperties
	RDBServiceProperties
	ResourceACLServiceProperties
	RouterServiceProperties
	S2ServiceProperties
	SecurityGroupServiceProperties
	SnapshotServiceProperties
	SpanServiceProperties
	SparkServiceProperties
	SubuserServiceProperties
	DescribeSubUsersInput
	DescribeSubUsersOutput
	CreateSubUserInput
	CreateSubUserOutput
	ModifySubUserAttributesInput
	ModifySubUserAttributesOutput
	DeleteSubUsersInput
	DeleteSubUsersOutput
	RestoreSubUsersInput
	RestoreSubUsersOutput
	TagServiceProperties
	DescribeTagsInput
	DescribeTagsOutput
	CreateTagInput
	CreateTagOutput
	DeleteTagsInput
	DeleteTagsOutput
	ModifyTagAttributesInput
	ModifyTagAttributesOutput
	AttachTagsInput
	AttachTagsOutput
	DetachTagsInput
	DetachTagsOutput
	Tag
	EIP
	Job
	Volume
	NIC
	KeyPair
	VxNet
	Router
	Instance
	Image
	Mongo
	MongoNode
	UserDataServiceProperties
	UploadUserDataAttachmentInput
	UploadUserDataAttachmentOutput
	VolumesServiceProperties
	DescribeVolumesInput
	DescribeVolumesOutput
	CreateVolumesInput
	CreateVolumesOutput
	DeleteVolumesInput
	DeleteVolumesOutput
	AttachVolumesInput
	AttachVolumesOutput
	DetachVolumesInput
	DetachVolumesOutput
	ResizeVolumesInput
	ResizeVolumesOutput
	ModifyVolumeAttributesInput
	ModifyVolumeAttributesOutput
	VxnetServiceProperties
	ZoneServiceProperties
	DescribeZonesInput
	DescribeZonesOutput
*/
package service

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
	proto.RegisterType((*AlarmServiceProperties)(nil), "service.AlarmServiceProperties")
}

type AlarmServiceInterface interface {
	DescribeAlarmPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyAlarmPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteAlarmPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyAlarmPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyAlarmPolicyActionAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AssociateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DissociateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplyAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeAlarms(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeAlarmHistory(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type AlarmService struct {
	Config     *config.Config
	Properties *AlarmServiceProperties
}

func NewAlarmService(conf *config.Config, zone string) (p *AlarmService) {
	return &AlarmService{
		Config:     conf,
		Properties: &AlarmServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Alarm(zone string) (*AlarmService, error) {
	properties := &AlarmServiceProperties{
		Zone: zone,
	}

	return &AlarmService{Config: s.Config, Properties: properties}, nil
}

func (p *AlarmService) DescribeAlarmPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicies",
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

func (p *AlarmService) CreateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateAlarmPolicy",
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

func (p *AlarmService) ModifyAlarmPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyAttributes",
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

func (p *AlarmService) DeleteAlarmPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicies",
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

func (p *AlarmService) DescribeAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyRules",
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

func (p *AlarmService) AddAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyRules",
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

func (p *AlarmService) ModifyAlarmPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyRuleAttributes",
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

func (p *AlarmService) DeleteAlarmPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyRules",
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

func (p *AlarmService) DescribeAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyActions",
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

func (p *AlarmService) AddAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyActions",
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

func (p *AlarmService) ModifyAlarmPolicyActionAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyActionAttributes",
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

func (p *AlarmService) DeleteAlarmPolicyActions(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyActions",
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

func (p *AlarmService) AssociateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateAlarmPolicy",
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

func (p *AlarmService) DissociateAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateAlarmPolicy",
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

func (p *AlarmService) ApplyAlarmPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyAlarmPolicy",
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

func (p *AlarmService) DescribeAlarms(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarms",
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

func (p *AlarmService) DescribeAlarmHistory(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmHistory",
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

func init() { proto.RegisterFile("alarm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x29, 0x88, 0xd2, 0x51, 0x44, 0xa3, 0x2d, 0xa5, 0x3d, 0xf8, 0x71, 0xf2, 0x20, 0x29,
	0xe8, 0x0b, 0xb8, 0xb6, 0xb5, 0x45, 0x2c, 0xd4, 0xf6, 0x09, 0xf2, 0x31, 0x2d, 0x0b, 0xdb, 0x4e,
	0xd8, 0xdd, 0x08, 0xf1, 0xcd, 0xbd, 0x49, 0x36, 0x08, 0x49, 0xa3, 0x87, 0xec, 0x7a, 0x4b, 0x86,
	0xd9, 0xdf, 0xfc, 0xf7, 0xc7, 0x2c, 0x1c, 0x07, 0x22, 0x90, 0x5b, 0x3f, 0x91, 0xa4, 0xc9, 0x3b,
	0x52, 0x28, 0x3f, 0x78, 0x84, 0xfd, 0xc1, 0x86, 0x68, 0x23, 0x70, 0x68, 0xca, 0x61, 0xba, 0x1e,
	0xe2, 0x36, 0xd1, 0x59, 0xd1, 0x75, 0x7b, 0x0f, 0x5d, 0x96, 0x1f, 0x5a, 0x15, 0xcd, 0x0b, 0x49,
	0x09, 0x4a, 0xcd, 0x51, 0x79, 0x1e, 0x1c, 0x7c, 0xd2, 0x0e, 0x7b, 0xad, 0xeb, 0xd6, 0x5d, 0x7b,
	0x69, 0xbe, 0x1f, 0xbe, 0xda, 0x70, 0x52, 0x6e, 0xf7, 0xa6, 0xd0, 0x19, 0xa3, 0x8a, 0x24, 0x0f,
	0xd1, 0xd4, 0x17, 0x24, 0x78, 0x94, 0x9f, 0xee, 0xfa, 0xc5, 0x54, 0xff, 0x67, 0xaa, 0x3f, 0xc9,
	0xa7, 0xf6, 0xff, 0xa8, 0x7b, 0x23, 0x38, 0x1f, 0x49, 0x0c, 0x74, 0x09, 0x93, 0x35, 0x86, 0xcc,
	0x61, 0x30, 0xa7, 0x98, 0xaf, 0xb3, 0x12, 0x84, 0x69, 0x2d, 0x79, 0x98, 0x6a, 0x8b, 0x4c, 0x13,
	0xb8, 0x18, 0xa3, 0x40, 0xed, 0x78, 0xb5, 0x57, 0xe8, 0xd5, 0x1d, 0x65, 0xcb, 0x54, 0xd8, 0x45,
	0x62, 0x71, 0xec, 0x8c, 0x79, 0x87, 0xab, 0x9a, 0xa8, 0x9c, 0xe4, 0x20, 0x6b, 0x06, 0xdd, 0x7d,
	0x59, 0x96, 0xe1, 0xde, 0xa0, 0xff, 0x8b, 0x2f, 0x16, 0x69, 0x4e, 0xbb, 0xe6, 0xb4, 0x29, 0x74,
	0xaa, 0xc6, 0x6c, 0x41, 0x2b, 0xb8, 0xa9, 0x2f, 0x97, 0x61, 0x39, 0x58, 0x33, 0xbb, 0xb1, 0x67,
	0xcd, 0x36, 0xe0, 0x0b, 0x5c, 0x32, 0xa5, 0x28, 0xe2, 0x8e, 0xaf, 0x28, 0x7f, 0xd3, 0xfc, 0x3f,
	0x40, 0xcf, 0x70, 0xc6, 0x92, 0x44, 0x64, 0x2e, 0x8c, 0x27, 0x38, 0xad, 0x2c, 0x83, 0x95, 0x96,
	0x0a, 0x61, 0xc6, 0x95, 0x26, 0xd9, 0x38, 0x49, 0x78, 0x68, 0xfe, 0x1f, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xba, 0xf8, 0x26, 0xf8, 0x65, 0x05, 0x00, 0x00,
}
