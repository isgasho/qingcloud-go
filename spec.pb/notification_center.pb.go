// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification_center.proto

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

type NotificationCenterServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *NotificationCenterServiceProperties) Reset()         { *m = NotificationCenterServiceProperties{} }
func (m *NotificationCenterServiceProperties) String() string { return proto.CompactTextString(m) }
func (*NotificationCenterServiceProperties) ProtoMessage()    {}
func (*NotificationCenterServiceProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor16, []int{0}
}

func (m *NotificationCenterServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*NotificationCenterServiceProperties)(nil), "spec.NotificationCenterServiceProperties")
}

type NotificationCenterServiceInterface interface {
	DescribeNotificationCenterUserPosts(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type NotificationCenterService struct {
	Config     *config.Config
	Properties *NotificationCenterServiceProperties
}

func NewNotificationCenterService(conf *config.Config, zone string) (p *NotificationCenterService, err error) {
	return &NotificationCenterService{
		Config:     conf,
		Properties: &NotificationCenterServiceProperties{Zone: zone},
	}, nil
}

func (p *NotificationCenterService) DescribeNotificationCenterUserPosts(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("notification_center.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x46, 0x29, 0x14, 0xc1, 0x1c, 0x73, 0x10, 0x5b, 0x41, 0xc5, 0x5e, 0x3c, 0xa5, 0xa0, 0x27,
	0xcf, 0xea, 0x55, 0x8a, 0xd2, 0xb3, 0xa4, 0x61, 0x5a, 0x02, 0x36, 0x13, 0x66, 0xc6, 0x5d, 0x76,
	0x7f, 0xfd, 0xb2, 0xe9, 0x16, 0x96, 0x2e, 0xbd, 0x25, 0xdf, 0x7b, 0x81, 0x17, 0x55, 0x04, 0x14,
	0xdf, 0x7b, 0x67, 0xc5, 0x63, 0xf8, 0x75, 0x10, 0x04, 0xc8, 0x44, 0x42, 0x41, 0x9d, 0x73, 0x04,
	0x57, 0x16, 0x03, 0xe2, 0xf0, 0x07, 0x75, 0xda, 0xba, 0xff, 0xbe, 0xb6, 0x61, 0x37, 0x09, 0xe5,
	0xdd, 0x12, 0xc1, 0x18, 0x65, 0x86, 0x0f, 0x4b, 0x28, 0x7e, 0x04, 0x16, 0x3b, 0xc6, 0x93, 0x70,
	0xbf, 0x14, 0xb6, 0x64, 0x63, 0x04, 0xe2, 0x89, 0x3f, 0xbd, 0xa9, 0xea, 0xeb, 0xac, 0xed, 0x3d,
	0xa5, 0xfd, 0x00, 0x6d, 0xbc, 0x83, 0x86, 0x30, 0x02, 0x89, 0x07, 0xd6, 0x5a, 0xe5, 0x7b, 0x0c,
	0x70, 0x9b, 0x3d, 0x66, 0xcf, 0xd7, 0xdf, 0xe9, 0xfc, 0x42, 0xaa, 0x58, 0x7d, 0xaa, 0x5b, 0x55,
	0x7d, 0x00, 0x3b, 0xf2, 0x1d, 0x5c, 0x4a, 0x2d, 0x03, 0x35, 0xc8, 0xc2, 0xfa, 0xc6, 0x4c, 0x7d,
	0x66, 0xee, 0x33, 0x9f, 0xc7, 0xdf, 0x95, 0x2b, 0x7b, 0x77, 0x95, 0xee, 0xaf, 0x87, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0x02, 0x4b, 0x5d, 0x51, 0x01, 0x00, 0x00,
}
