// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alarm.proto

/*
Package spec is a generated protocol buffer package.

It is generated from these files:
	alarm.proto
	base.proto
	cache.proto
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
	shared_storage.proto
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
	Options
	CommonRequest
	Error
	DescribeInstanceTypes_Request
	DescribeInstanceTypes_Reponse
	InstanceTypeSetElem
*/
package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "context"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AlarmService interface {
	DescribeAlarmPolicies(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	CreateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicies(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyRuleAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AddAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyAlarmPolicyActionAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AssociateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DissociateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ApplyAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarms(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeAlarmHistory(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type AlarmServiceClient struct{}

// NewAlarmServiceClient returns a AlarmService stub to handle
// requests to the set of AlarmService at the other end of the connection.
func NewAlarmServiceClient(opt *Options) *AlarmServiceClient {
	return &AlarmServiceClient{}
}

func (c *AlarmServiceClient) DescribeAlarmPolicies(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) CreateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) ModifyAlarmPolicyAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DeleteAlarmPolicies(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DescribeAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) AddAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) ModifyAlarmPolicyRuleAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DeleteAlarmPolicyRules(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DescribeAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) AddAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) ModifyAlarmPolicyActionAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DeleteAlarmPolicyActions(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) AssociateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DissociateAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) ApplyAlarmPolicy(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DescribeAlarms(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *AlarmServiceClient) DescribeAlarmHistory(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("alarm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x6f, 0x44, 0xf0, 0x28, 0xa2, 0xd5, 0x0d, 0xed, 0xc0, 0xe1, 0x0b, 0x74, 0xa0, 0x2f,
	0x60, 0xdc, 0xe6, 0x86, 0x38, 0x50, 0xf7, 0x04, 0x69, 0x76, 0x36, 0x02, 0xe9, 0x12, 0x92, 0x33,
	0xa5, 0x8f, 0xee, 0x9d, 0xb4, 0x73, 0xb0, 0xb6, 0x7a, 0xd1, 0xc4, 0xcb, 0x9e, 0xff, 0xf0, 0xf1,
	0xf7, 0xe3, 0x04, 0x8e, 0xb9, 0xe2, 0x36, 0x4b, 0x8c, 0xd5, 0xa4, 0xa3, 0x03, 0x67, 0x50, 0xc4,
	0xd7, 0x2b, 0xad, 0x57, 0x0a, 0x07, 0xe5, 0x2c, 0xdd, 0x2c, 0x07, 0x7c, 0x9d, 0x6f, 0x17, 0xe2,
	0x5e, 0x3d, 0xc2, 0xcc, 0xd0, 0x2e, 0xec, 0xd7, 0x43, 0x92, 0x19, 0x3a, 0xe2, 0x99, 0xf9, 0x59,
	0xb8, 0xa9, 0x2f, 0x7c, 0x5a, 0x6e, 0x0c, 0x5a, 0xb7, 0xcd, 0xef, 0xbe, 0x8e, 0xe0, 0x84, 0x15,
	0x75, 0xe6, 0x68, 0x3f, 0xa4, 0xc0, 0x68, 0x02, 0x9d, 0x11, 0x3a, 0x61, 0x65, 0x8a, 0xe5, 0xfc,
	0x55, 0x2b, 0x29, 0x24, 0xba, 0xa8, 0x9b, 0x6c, 0x51, 0xc9, 0x0e, 0x95, 0x8c, 0x8b, 0x22, 0xf1,
	0x1f, 0xf3, 0x68, 0x08, 0xe7, 0x43, 0x8b, 0x9c, 0xf6, 0x30, 0x79, 0x6b, 0xc8, 0x0c, 0x7a, 0x33,
	0xbd, 0x90, 0xcb, 0x7c, 0x0f, 0xc2, 0x88, 0xac, 0x4c, 0x37, 0xe4, 0xd1, 0x69, 0x0c, 0x17, 0x23,
	0x54, 0x48, 0x81, 0xbf, 0xf6, 0x0c, 0x57, 0x4d, 0x47, 0xf9, 0xfb, 0x46, 0xf9, 0x55, 0x62, 0x8b,
	0x45, 0x30, 0xe6, 0x0d, 0xfa, 0x0d, 0x51, 0x05, 0x29, 0x40, 0xd6, 0x14, 0xba, 0x75, 0x59, 0x9e,
	0xe5, 0x5e, 0x20, 0xfe, 0xc5, 0x17, 0x13, 0x24, 0xf5, 0xba, 0x3d, 0x6d, 0x02, 0x9d, 0xaa, 0x31,
	0x5f, 0xd0, 0x1c, 0x6e, 0x9b, 0xc7, 0x55, 0xb2, 0x02, 0xac, 0x95, 0xb7, 0x51, 0xb3, 0xe6, 0x5b,
	0xf0, 0x09, 0x2e, 0x99, 0x73, 0x5a, 0xc8, 0xc0, 0x57, 0x54, 0xbc, 0x69, 0xf9, 0x1f, 0xa0, 0x47,
	0x38, 0x63, 0xc6, 0xa8, 0x3c, 0x84, 0xf1, 0x00, 0xa7, 0x95, 0x63, 0xf0, 0xd2, 0x52, 0x21, 0x4c,
	0xa5, 0x23, 0x6d, 0x5b, 0x37, 0x49, 0x0f, 0xcb, 0xef, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9b, 0x5d, 0xb5, 0x7f, 0x90, 0x05, 0x00, 0x00,
}