// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zone.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

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

type ZoneServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ZoneServiceProperties) Reset()                    { *m = ZoneServiceProperties{} }
func (m *ZoneServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ZoneServiceProperties) ProtoMessage()               {}
func (*ZoneServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *ZoneServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeZonesInput struct {
	Zones  []string `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	Status []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
}

func (m *DescribeZonesInput) Reset()                    { *m = DescribeZonesInput{} }
func (m *DescribeZonesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeZonesInput) ProtoMessage()               {}
func (*DescribeZonesInput) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

func (m *DescribeZonesInput) GetZones() []string {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *DescribeZonesInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type DescribeZonesOutput struct {
	Action     string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ZoneSet    []*Zone `protobuf:"bytes,4,rep,name=zone_set,json=zoneSet" json:"zone_set,omitempty"`
	TotalCount int32   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeZonesOutput) Reset()                    { *m = DescribeZonesOutput{} }
func (m *DescribeZonesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeZonesOutput) ProtoMessage()               {}
func (*DescribeZonesOutput) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

func (m *DescribeZonesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeZonesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeZonesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeZonesOutput) GetZoneSet() []*Zone {
	if m != nil {
		return m.ZoneSet
	}
	return nil
}

func (m *DescribeZonesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ZoneServiceProperties)(nil), "service.ZoneServiceProperties")
	proto.RegisterType((*DescribeZonesInput)(nil), "service.DescribeZonesInput")
	proto.RegisterType((*DescribeZonesOutput)(nil), "service.DescribeZonesOutput")
}

type ZoneServiceInterface interface {
	DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error)
}

type ZoneService struct {
	Config           *config.Config
	Properties       *ZoneServiceProperties
	LastResponseBody string
}

func NewZoneService(conf *config.Config, zone string) (p *ZoneService) {
	return &ZoneService{
		Config:     conf,
		Properties: &ZoneServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Zone(zone string) (*ZoneService, error) {
	properties := &ZoneServiceProperties{
		Zone: zone,
	}

	return &ZoneService{Config: s.Config, Properties: properties}, nil
}

func (p *ZoneService) DescribeZones(in *DescribeZonesInput) (out *DescribeZonesOutput, err error) {
	if in == nil {
		in = &DescribeZonesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeZonesOutput{}
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

func (p *DescribeZonesInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("zone.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xd1, 0x4a, 0xeb, 0x40,
	0x10, 0x25, 0x6d, 0xda, 0xb4, 0x13, 0xfa, 0x70, 0xe7, 0xde, 0x96, 0xbd, 0x51, 0x31, 0xf4, 0x29,
	0x20, 0x54, 0xa8, 0x7f, 0x60, 0x7d, 0xd1, 0x17, 0x35, 0xbe, 0xf9, 0x12, 0xd2, 0x64, 0x28, 0xc1,
	0xba, 0x1b, 0x77, 0x27, 0x82, 0xfd, 0x04, 0xff, 0xc5, 0x8f, 0xf3, 0x0f, 0x64, 0x37, 0x51, 0x2c,
	0xfa, 0xb2, 0xec, 0x39, 0x87, 0x39, 0x33, 0x73, 0x06, 0x60, 0xa7, 0x24, 0x2d, 0x6a, 0xad, 0x58,
	0x61, 0x60, 0x48, 0x3f, 0x57, 0x05, 0x45, 0x21, 0xbf, 0xd4, 0x64, 0x5a, 0x36, 0x3a, 0x7a, 0xaa,
	0xe4, 0xa6, 0xd8, 0xaa, 0xa6, 0xcc, 0x4c, 0xf9, 0x90, 0xe9, 0x66, 0x4b, 0xa7, 0xf6, 0x69, 0xe5,
	0xf9, 0x09, 0x4c, 0xef, 0x95, 0xa4, 0xbb, 0xb6, 0xf4, 0x46, 0xab, 0x9a, 0x34, 0x57, 0x64, 0x10,
	0xc1, 0xb7, 0xde, 0xc2, 0x8b, 0xbd, 0x64, 0x9c, 0xba, 0xff, 0xfc, 0x1c, 0xf0, 0x82, 0x4c, 0xa1,
	0xab, 0x35, 0xd9, 0x22, 0x73, 0x29, 0xeb, 0x86, 0xf1, 0x1f, 0x0c, 0xac, 0x6a, 0x84, 0x17, 0xf7,
	0x93, 0x71, 0xda, 0x02, 0x9c, 0xc1, 0xd0, 0x70, 0xce, 0x8d, 0x11, 0x3d, 0x47, 0x77, 0x68, 0xfe,
	0xe6, 0xc1, 0xdf, 0x3d, 0x93, 0xeb, 0x86, 0xad, 0xcb, 0x0c, 0x86, 0x79, 0xc1, 0x95, 0x92, 0x5d,
	0xc7, 0x0e, 0xe1, 0x7f, 0x18, 0x69, 0xe2, 0xac, 0x50, 0x25, 0x89, 0x5e, 0xec, 0x25, 0x83, 0x34,
	0xd0, 0xc4, 0x2b, 0x55, 0x12, 0x0a, 0x08, 0x1e, 0xc9, 0x98, 0x7c, 0x43, 0xa2, 0xef, 0x6a, 0x3e,
	0x21, 0x26, 0x30, 0xb2, 0x53, 0x64, 0x86, 0x58, 0xf8, 0x71, 0x3f, 0x09, 0x97, 0x93, 0x45, 0x97,
	0xce, 0xc2, 0x36, 0x4d, 0x83, 0x9d, 0x5b, 0x9a, 0xf1, 0x18, 0x42, 0x56, 0x9c, 0x6f, 0xb3, 0x42,
	0x35, 0x92, 0xc5, 0xc0, 0x75, 0x00, 0x47, 0xad, 0x2c, 0xb3, 0xac, 0x21, 0xfc, 0x16, 0x10, 0x5e,
	0xc1, 0x64, 0x6f, 0x7a, 0x3c, 0xf8, 0x32, 0xfe, 0x19, 0x4d, 0x74, 0xf8, 0xbb, 0xd8, 0xae, 0x1c,
	0x4d, 0x5f, 0xdf, 0xfd, 0x3f, 0x38, 0xbe, 0xad, 0xe4, 0x66, 0x65, 0x0f, 0x14, 0xf9, 0x56, 0x5d,
	0x0f, 0xdd, 0x65, 0xce, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x3e, 0x12, 0xf5, 0xdc, 0x01,
	0x00, 0x00,
}
