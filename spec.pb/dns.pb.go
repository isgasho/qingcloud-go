// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

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

type DNSAliasServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *DNSAliasServiceProperties) Reset()                    { *m = DNSAliasServiceProperties{} }
func (m *DNSAliasServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*DNSAliasServiceProperties) ProtoMessage()               {}
func (*DNSAliasServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DNSAliasServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*DNSAliasServiceProperties)(nil), "spec.DNSAliasServiceProperties")
}

type DNSAliasServiceInterface interface {
	DescribeDNSAliases(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AssociateDNSAlias(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DissociateDNSAliases(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	GetDNSLabel(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type DNSAliasService struct {
	Config     *config.Config
	Properties *DNSAliasServiceProperties
}

func NewDNSAliasService(conf *config.Config, zone string) (p *DNSAliasService, err error) {
	return &DNSAliasService{
		Config:     conf,
		Properties: &DNSAliasServiceProperties{Zone: zone},
	}, nil
}

func (p *DNSAliasService) DescribeDNSAliases(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *DNSAliasService) AssociateDNSAlias(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *DNSAliasService) DissociateDNSAliases(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (p *DNSAliasService) GetDNSLabel(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8e, 0xbd, 0x4a, 0xc5, 0x30,
	0x14, 0xc7, 0xa9, 0x14, 0xa1, 0xc7, 0x41, 0x0c, 0x22, 0xb6, 0x82, 0x8a, 0x93, 0x53, 0x0a, 0x3a,
	0x3b, 0x14, 0xa3, 0x2e, 0x52, 0xc4, 0x3e, 0x41, 0x12, 0x8f, 0x25, 0xd0, 0x36, 0x21, 0x27, 0x2a,
	0xfa, 0x3a, 0xbe, 0xa8, 0xd8, 0x8f, 0x3b, 0x04, 0xee, 0xd0, 0xbb, 0x25, 0xe7, 0xf7, 0xff, 0x82,
	0xec, 0x6d, 0x20, 0xee, 0xbc, 0x0d, 0x96, 0xa5, 0xe4, 0x50, 0x17, 0x79, 0x6b, 0x6d, 0xdb, 0x61,
	0x39, 0xde, 0xd4, 0xc7, 0x7b, 0x29, 0x87, 0xef, 0x49, 0x50, 0x9c, 0xc5, 0x08, 0x7b, 0x17, 0x16,
	0x78, 0x11, 0xc3, 0x60, 0x7a, 0xa4, 0x20, 0x7b, 0x37, 0x0b, 0xce, 0x63, 0xc1, 0x97, 0x97, 0xce,
	0xa1, 0x9f, 0xeb, 0xaf, 0x4a, 0xc8, 0x45, 0xdd, 0x54, 0x9d, 0x91, 0xd4, 0xa0, 0xff, 0x34, 0x1a,
	0x5f, 0xbc, 0x75, 0xe8, 0x83, 0x41, 0x62, 0x0c, 0xd2, 0x1f, 0x3b, 0xe0, 0x69, 0x72, 0x99, 0x5c,
	0x67, 0xaf, 0xe3, 0xfb, 0xe6, 0x77, 0x0f, 0x0e, 0x23, 0x07, 0x13, 0xc0, 0x04, 0x92, 0xf6, 0x46,
	0xe1, 0x82, 0x90, 0xd8, 0x09, 0x9f, 0xba, 0xf9, 0xd2, 0xcd, 0x1f, 0xfe, 0x97, 0x17, 0x5b, 0xee,
	0xec, 0x1e, 0x8e, 0x2a, 0x22, 0xab, 0x8d, 0x0c, 0x9b, 0x98, 0xd5, 0x21, 0x8f, 0x70, 0x2c, 0x4c,
	0x9c, 0xb2, 0xc3, 0x98, 0x3b, 0x38, 0x78, 0xc2, 0x20, 0xea, 0xe6, 0x59, 0x2a, 0xec, 0xd6, 0xda,
	0xd5, 0xfe, 0xf8, 0xbf, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x91, 0x09, 0x41, 0xe9, 0x01,
	0x00, 0x00,
}
