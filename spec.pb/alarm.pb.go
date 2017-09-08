// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm.proto

/*
Package spec is a generated protocol buffer package.

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

func NewAlarmService(conf *config.Config, zone string) (p *AlarmService, err error) {
	return &AlarmService{
		Config:     conf,
		Properties: &AlarmServiceProperties{Zone: zone},
	}, nil
}

func (p *AlarmService) DescribeAlarmPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	o := &request.Operation{
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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0x87, 0x19, 0x0c, 0x61, 0x47, 0x11, 0xad, 0x6e, 0x8c, 0xee, 0xc2, 0x3f, 0x57, 0x5e, 0x48,
	0x07, 0xfa, 0x02, 0xc6, 0x6d, 0x6e, 0x88, 0x83, 0xb9, 0x3d, 0x41, 0x9b, 0x9e, 0x8d, 0x40, 0xb6,
	0x84, 0x24, 0x15, 0xea, 0x9b, 0x7b, 0x27, 0xcd, 0x10, 0xda, 0x55, 0x2f, 0x9a, 0x78, 0xd7, 0x1e,
	0x92, 0xef, 0xfc, 0xf2, 0x71, 0x0e, 0x1c, 0xc7, 0x3c, 0x56, 0xdb, 0x48, 0x2a, 0x61, 0x44, 0xd0,
	0xd6, 0x12, 0x69, 0x38, 0xd8, 0x08, 0xb1, 0xe1, 0x38, 0xb4, 0xb5, 0x24, 0x5b, 0x0f, 0x71, 0x2b,
	0x4d, 0xbe, 0x3f, 0x72, 0x7b, 0x0f, 0x3d, 0x52, 0xdc, 0x58, 0xa1, 0xfa, 0x60, 0x14, 0x17, 0x4a,
	0x48, 0x54, 0x86, 0xa1, 0x0e, 0x02, 0x68, 0x7f, 0x8a, 0x1d, 0xf6, 0x5b, 0xd7, 0xad, 0xbb, 0xce,
	0xd2, 0x7e, 0x3f, 0x7c, 0x75, 0xe0, 0xa4, 0x7c, 0x3c, 0x98, 0x42, 0x77, 0x8c, 0x9a, 0x2a, 0x96,
	0xa0, 0xad, 0x2f, 0x04, 0x67, 0xb4, 0xb8, 0xdd, 0x8b, 0xf6, 0x5d, 0xa3, 0x9f, 0xae, 0xd1, 0xa4,
	0xe8, 0x1a, 0xfe, 0x51, 0x0f, 0x46, 0x70, 0x3e, 0x52, 0x18, 0x9b, 0x12, 0x26, 0x6f, 0x0c, 0x99,
	0xc3, 0x60, 0x2e, 0x52, 0xb6, 0xce, 0x4b, 0x10, 0x62, 0x8c, 0x62, 0x49, 0x66, 0x1c, 0x32, 0x4d,
	0xe0, 0x62, 0x8c, 0x1c, 0x8d, 0xe7, 0xd3, 0x5e, 0xa1, 0x5f, 0x77, 0x94, 0x2f, 0x33, 0xee, 0x16,
	0x89, 0xa4, 0xa9, 0x37, 0xe6, 0x1d, 0xae, 0x6a, 0xa2, 0x0a, 0x92, 0x87, 0xac, 0x19, 0xf4, 0x0e,
	0x65, 0x39, 0x86, 0x7b, 0x83, 0xf0, 0x17, 0x5f, 0x84, 0x1a, 0x26, 0x76, 0xcd, 0x69, 0x53, 0xe8,
	0x56, 0x8d, 0xb9, 0x82, 0x56, 0x70, 0x53, 0x1f, 0x2e, 0xcb, 0xf2, 0xb0, 0x66, 0x67, 0xe3, 0xc0,
	0x9a, 0x6b, 0xc0, 0x17, 0xb8, 0x24, 0x5a, 0x0b, 0xca, 0x3c, 0xb7, 0xa8, 0xd8, 0x69, 0xf6, 0x1f,
	0xa0, 0x67, 0x38, 0x23, 0x52, 0xf2, 0xdc, 0x87, 0xf1, 0x04, 0xa7, 0x95, 0x61, 0x70, 0xd2, 0x52,
	0x21, 0xcc, 0x98, 0x36, 0x42, 0x35, 0x4e, 0x92, 0x1c, 0xd9, 0xff, 0xc7, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x0c, 0xe7, 0xbe, 0x62, 0x05, 0x00, 0x00,
}
