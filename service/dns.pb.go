// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

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

type DNSAliasServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *DNSAliasServiceProperties) Reset()                    { *m = DNSAliasServiceProperties{} }
func (m *DNSAliasServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*DNSAliasServiceProperties) ProtoMessage()               {}
func (*DNSAliasServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DNSAliasServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeDNSAliasesInput struct {
}

func (m *DescribeDNSAliasesInput) Reset()                    { *m = DescribeDNSAliasesInput{} }
func (m *DescribeDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesInput) ProtoMessage()               {}
func (*DescribeDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type DescribeDNSAliasesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeDNSAliasesOutput) Reset()                    { *m = DescribeDNSAliasesOutput{} }
func (m *DescribeDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesOutput) ProtoMessage()               {}
func (*DescribeDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *DescribeDNSAliasesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeDNSAliasesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeDNSAliasesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AssociateDNSAliasInput struct {
}

func (m *AssociateDNSAliasInput) Reset()                    { *m = AssociateDNSAliasInput{} }
func (m *AssociateDNSAliasInput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasInput) ProtoMessage()               {}
func (*AssociateDNSAliasInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type AssociateDNSAliasOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AssociateDNSAliasOutput) Reset()                    { *m = AssociateDNSAliasOutput{} }
func (m *AssociateDNSAliasOutput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasOutput) ProtoMessage()               {}
func (*AssociateDNSAliasOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *AssociateDNSAliasOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AssociateDNSAliasOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AssociateDNSAliasOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DissociateDNSAliasesInput struct {
}

func (m *DissociateDNSAliasesInput) Reset()                    { *m = DissociateDNSAliasesInput{} }
func (m *DissociateDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesInput) ProtoMessage()               {}
func (*DissociateDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type DissociateDNSAliasesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DissociateDNSAliasesOutput) Reset()                    { *m = DissociateDNSAliasesOutput{} }
func (m *DissociateDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesOutput) ProtoMessage()               {}
func (*DissociateDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *DissociateDNSAliasesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DissociateDNSAliasesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DissociateDNSAliasesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetDNSLabelInput struct {
}

func (m *GetDNSLabelInput) Reset()                    { *m = GetDNSLabelInput{} }
func (m *GetDNSLabelInput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelInput) ProtoMessage()               {}
func (*GetDNSLabelInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type GetDNSLabelOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GetDNSLabelOutput) Reset()                    { *m = GetDNSLabelOutput{} }
func (m *GetDNSLabelOutput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelOutput) ProtoMessage()               {}
func (*GetDNSLabelOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *GetDNSLabelOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetDNSLabelOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetDNSLabelOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DNSAliasServiceProperties)(nil), "service.DNSAliasServiceProperties")
	proto.RegisterType((*DescribeDNSAliasesInput)(nil), "service.DescribeDNSAliasesInput")
	proto.RegisterType((*DescribeDNSAliasesOutput)(nil), "service.DescribeDNSAliasesOutput")
	proto.RegisterType((*AssociateDNSAliasInput)(nil), "service.AssociateDNSAliasInput")
	proto.RegisterType((*AssociateDNSAliasOutput)(nil), "service.AssociateDNSAliasOutput")
	proto.RegisterType((*DissociateDNSAliasesInput)(nil), "service.DissociateDNSAliasesInput")
	proto.RegisterType((*DissociateDNSAliasesOutput)(nil), "service.DissociateDNSAliasesOutput")
	proto.RegisterType((*GetDNSLabelInput)(nil), "service.GetDNSLabelInput")
	proto.RegisterType((*GetDNSLabelOutput)(nil), "service.GetDNSLabelOutput")
}

type DNSAliasServiceInterface interface {
	DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error)
	AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error)
	DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error)
	GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error)
}

type DNSAliasService struct {
	Config           *config.Config
	Properties       *DNSAliasServiceProperties
	LastResponseBody string
}

func NewDNSAliasService(conf *config.Config, zone string) (p *DNSAliasService) {
	return &DNSAliasService{
		Config:     conf,
		Properties: &DNSAliasServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) DNSAlias(zone string) (*DNSAliasService, error) {
	properties := &DNSAliasServiceProperties{
		Zone: zone,
	}

	return &DNSAliasService{Config: s.Config, Properties: properties}, nil
}

func (p *DNSAliasService) DescribeDNSAliases(in *DescribeDNSAliasesInput) (out *DescribeDNSAliasesOutput, err error) {
	if in == nil {
		in = &DescribeDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeDNSAliases",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeDNSAliasesOutput{}
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

func (p *DescribeDNSAliasesInput) Validate() error {
	return nil
}

func (p *DNSAliasService) AssociateDNSAlias(in *AssociateDNSAliasInput) (out *AssociateDNSAliasOutput, err error) {
	if in == nil {
		in = &AssociateDNSAliasInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateDNSAlias",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateDNSAliasOutput{}
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

func (p *AssociateDNSAliasInput) Validate() error {
	return nil
}

func (p *DNSAliasService) DissociateDNSAliases(in *DissociateDNSAliasesInput) (out *DissociateDNSAliasesOutput, err error) {
	if in == nil {
		in = &DissociateDNSAliasesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateDNSAliases",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateDNSAliasesOutput{}
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

func (p *DissociateDNSAliasesInput) Validate() error {
	return nil
}

func (p *DNSAliasService) GetDNSLabel(in *GetDNSLabelInput) (out *GetDNSLabelOutput, err error) {
	if in == nil {
		in = &GetDNSLabelInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetDNSLabel",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetDNSLabelOutput{}
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

func (p *GetDNSLabelInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x4f, 0xe2, 0x50,
	0x14, 0x0d, 0x03, 0xc3, 0xc7, 0x9d, 0xc5, 0x0c, 0x37, 0x33, 0xd0, 0xbe, 0x89, 0xb1, 0xd6, 0x0d,
	0x2b, 0x48, 0xf4, 0x17, 0x10, 0x9a, 0x18, 0x13, 0x83, 0x0a, 0x89, 0x89, 0xab, 0x5a, 0xda, 0x6b,
	0xf3, 0x62, 0xed, 0xab, 0xef, 0xbd, 0xba, 0xf0, 0x07, 0xb8, 0xf0, 0xdf, 0xfa, 0x0f, 0x0c, 0xb4,
	0x45, 0x42, 0x0b, 0x3b, 0x36, 0x4d, 0x5f, 0xef, 0x39, 0xf7, 0x9c, 0xf4, 0x9c, 0x07, 0x9d, 0x20,
	0x56, 0xc3, 0x44, 0x0a, 0x2d, 0xb0, 0xa5, 0x48, 0xbe, 0x72, 0x9f, 0xd8, 0xd1, 0x0b, 0x8f, 0x43,
	0x3f, 0x12, 0x69, 0xe0, 0xaa, 0xe0, 0xc9, 0x95, 0x69, 0x44, 0xa3, 0xe5, 0x23, 0xc3, 0xd9, 0x23,
	0x30, 0x9d, 0xe9, 0x7c, 0x1c, 0x71, 0x4f, 0xcd, 0x33, 0xc6, 0x8d, 0x14, 0x09, 0x49, 0xcd, 0x49,
	0x21, 0x42, 0xe3, 0x4d, 0xc4, 0x64, 0xd4, 0xac, 0xda, 0xa0, 0x33, 0x5b, 0xbd, 0xdb, 0x26, 0xf4,
	0x1d, 0x52, 0xbe, 0xe4, 0x0b, 0x2a, 0x88, 0xa4, 0x2e, 0xe3, 0x24, 0xd5, 0x76, 0x08, 0x46, 0x79,
	0x74, 0x9d, 0xea, 0x24, 0xd5, 0xd8, 0x83, 0xa6, 0xe7, 0x6b, 0x2e, 0xe2, 0x7c, 0x59, 0x7e, 0x42,
	0x13, 0xda, 0x92, 0xb4, 0xeb, 0x8b, 0x80, 0x8c, 0x1f, 0x56, 0x6d, 0xf0, 0x73, 0xd6, 0x92, 0xa4,
	0x27, 0x22, 0x20, 0x34, 0xa0, 0xf5, 0x4c, 0x4a, 0x79, 0x21, 0x19, 0xf5, 0x15, 0xa7, 0x38, 0xda,
	0x06, 0xf4, 0xc6, 0x4a, 0x09, 0x9f, 0x7b, 0x7a, 0xad, 0x94, 0x59, 0x78, 0x84, 0x7e, 0x69, 0x72,
	0x08, 0x07, 0xff, 0xc1, 0x74, 0xf8, 0xb6, 0x50, 0xf1, 0x1f, 0x38, 0xb0, 0xaa, 0xe1, 0x21, 0x7c,
	0x20, 0xfc, 0xb9, 0x20, 0xed, 0x4c, 0xe7, 0x57, 0xde, 0x82, 0xa2, 0x4c, 0xfe, 0x01, 0xba, 0x1b,
	0xdf, 0x0e, 0xa0, 0x7a, 0xf6, 0x5e, 0x87, 0xdf, 0x5b, 0xad, 0xc1, 0x7b, 0xc0, 0x72, 0xf8, 0x68,
	0x0d, 0xf3, 0x1e, 0x0e, 0x77, 0x94, 0x86, 0x9d, 0xec, 0x41, 0xe4, 0xde, 0xef, 0xa0, 0x5b, 0x0a,
	0x15, 0x8f, 0xd7, 0xbc, 0xea, 0x2a, 0x30, 0x6b, 0x37, 0x20, 0xdf, 0xeb, 0xc2, 0xdf, 0xaa, 0x9c,
	0xd0, 0xfe, 0xb6, 0xb4, 0x2b, 0x63, 0x76, 0xba, 0x17, 0x93, 0x0b, 0x38, 0xf0, 0x6b, 0x23, 0x09,
	0x34, 0xd7, 0x9c, 0xed, 0xcc, 0x18, 0xab, 0x1a, 0x65, 0x5b, 0x98, 0xf9, 0xf1, 0xd9, 0xf8, 0x07,
	0x9d, 0x5b, 0x1e, 0x87, 0x93, 0xe5, 0x3d, 0xc6, 0x76, 0x21, 0xb4, 0x68, 0xae, 0x2e, 0xf1, 0xf9,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x69, 0x01, 0xc0, 0xf9, 0x03, 0x00, 0x00,
}
