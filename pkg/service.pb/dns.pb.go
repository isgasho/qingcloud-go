// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/service.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DNSAliasServiceProperties struct {
	Zone             *string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DNSAliasServiceProperties) Reset()                    { *m = DNSAliasServiceProperties{} }
func (m *DNSAliasServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*DNSAliasServiceProperties) ProtoMessage()               {}
func (*DNSAliasServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *DNSAliasServiceProperties) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type DescribeDNSAliasesInput struct {
	DnsAliases       []string `protobuf:"bytes,1,rep,name=dns_aliases,json=dnsAliases" json:"dns_aliases,omitempty"`
	ResourceId       *string  `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	SearchWord       *string  `protobuf:"bytes,3,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Offset           *int32   `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	Limit            *int32   `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DescribeDNSAliasesInput) Reset()                    { *m = DescribeDNSAliasesInput{} }
func (m *DescribeDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesInput) ProtoMessage()               {}
func (*DescribeDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *DescribeDNSAliasesInput) GetDnsAliases() []string {
	if m != nil {
		return m.DnsAliases
	}
	return nil
}

func (m *DescribeDNSAliasesInput) GetResourceId() string {
	if m != nil && m.ResourceId != nil {
		return *m.ResourceId
	}
	return ""
}

func (m *DescribeDNSAliasesInput) GetSearchWord() string {
	if m != nil && m.SearchWord != nil {
		return *m.SearchWord
	}
	return ""
}

func (m *DescribeDNSAliasesInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeDNSAliasesInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type DescribeDNSAliasesOutput struct {
	Action           *string                                  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32                                   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string                                  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	DnsAliasSet      []*DescribeDNSAliasesOutput_ResponseItem `protobuf:"bytes,4,rep,name=dns_alias_set,json=dnsAliasSet" json:"dns_alias_set,omitempty"`
	TotalCount       *int32                                   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *DescribeDNSAliasesOutput) Reset()                    { *m = DescribeDNSAliasesOutput{} }
func (m *DescribeDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesOutput) ProtoMessage()               {}
func (*DescribeDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *DescribeDNSAliasesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeDNSAliasesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeDNSAliasesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeDNSAliasesOutput) GetDnsAliasSet() []*DescribeDNSAliasesOutput_ResponseItem {
	if m != nil {
		return m.DnsAliasSet
	}
	return nil
}

func (m *DescribeDNSAliasesOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type DescribeDNSAliasesOutput_ResponseItem struct {
	Status           *string                     `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	DnsAliasId       *string                     `protobuf:"bytes,2,opt,name=dns_alias_id,json=dnsAliasId" json:"dns_alias_id,omitempty"`
	ResourceId       *string                     `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	DnsAliasName     *string                     `protobuf:"bytes,4,opt,name=dns_alias_name,json=dnsAliasName" json:"dns_alias_name,omitempty"`
	CreateTime       *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *DescribeDNSAliasesOutput_ResponseItem) Reset()         { *m = DescribeDNSAliasesOutput_ResponseItem{} }
func (m *DescribeDNSAliasesOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeDNSAliasesOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeDNSAliasesOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2, 0}
}

func (m *DescribeDNSAliasesOutput_ResponseItem) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *DescribeDNSAliasesOutput_ResponseItem) GetDnsAliasId() string {
	if m != nil && m.DnsAliasId != nil {
		return *m.DnsAliasId
	}
	return ""
}

func (m *DescribeDNSAliasesOutput_ResponseItem) GetResourceId() string {
	if m != nil && m.ResourceId != nil {
		return *m.ResourceId
	}
	return ""
}

func (m *DescribeDNSAliasesOutput_ResponseItem) GetDnsAliasName() string {
	if m != nil && m.DnsAliasName != nil {
		return *m.DnsAliasName
	}
	return ""
}

func (m *DescribeDNSAliasesOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type AssociateDNSAliasInput struct {
	Prefix           *string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	Resource         *string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	DnsAliasName     *string `protobuf:"bytes,3,opt,name=dns_alias_name,json=dnsAliasName" json:"dns_alias_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AssociateDNSAliasInput) Reset()                    { *m = AssociateDNSAliasInput{} }
func (m *AssociateDNSAliasInput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasInput) ProtoMessage()               {}
func (*AssociateDNSAliasInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *AssociateDNSAliasInput) GetPrefix() string {
	if m != nil && m.Prefix != nil {
		return *m.Prefix
	}
	return ""
}

func (m *AssociateDNSAliasInput) GetResource() string {
	if m != nil && m.Resource != nil {
		return *m.Resource
	}
	return ""
}

func (m *AssociateDNSAliasInput) GetDnsAliasName() string {
	if m != nil && m.DnsAliasName != nil {
		return *m.DnsAliasName
	}
	return ""
}

type AssociateDNSAliasOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	DnsAliasId       *string `protobuf:"bytes,4,opt,name=dns_alias_id,json=dnsAliasId" json:"dns_alias_id,omitempty"`
	DomainName       *string `protobuf:"bytes,5,opt,name=domain_name,json=domainName" json:"domain_name,omitempty"`
	JobId            *string `protobuf:"bytes,6,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AssociateDNSAliasOutput) Reset()                    { *m = AssociateDNSAliasOutput{} }
func (m *AssociateDNSAliasOutput) String() string            { return proto.CompactTextString(m) }
func (*AssociateDNSAliasOutput) ProtoMessage()               {}
func (*AssociateDNSAliasOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *AssociateDNSAliasOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AssociateDNSAliasOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AssociateDNSAliasOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AssociateDNSAliasOutput) GetDnsAliasId() string {
	if m != nil && m.DnsAliasId != nil {
		return *m.DnsAliasId
	}
	return ""
}

func (m *AssociateDNSAliasOutput) GetDomainName() string {
	if m != nil && m.DomainName != nil {
		return *m.DomainName
	}
	return ""
}

func (m *AssociateDNSAliasOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type DissociateDNSAliasesInput struct {
	DnsAliases       []string `protobuf:"bytes,1,rep,name=dns_aliases,json=dnsAliases" json:"dns_aliases,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DissociateDNSAliasesInput) Reset()                    { *m = DissociateDNSAliasesInput{} }
func (m *DissociateDNSAliasesInput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesInput) ProtoMessage()               {}
func (*DissociateDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *DissociateDNSAliasesInput) GetDnsAliases() []string {
	if m != nil {
		return m.DnsAliases
	}
	return nil
}

type DissociateDNSAliasesOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DissociateDNSAliasesOutput) Reset()                    { *m = DissociateDNSAliasesOutput{} }
func (m *DissociateDNSAliasesOutput) String() string            { return proto.CompactTextString(m) }
func (*DissociateDNSAliasesOutput) ProtoMessage()               {}
func (*DissociateDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *DissociateDNSAliasesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DissociateDNSAliasesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DissociateDNSAliasesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DissociateDNSAliasesOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type GetDNSLabelInput struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetDNSLabelInput) Reset()                    { *m = GetDNSLabelInput{} }
func (m *GetDNSLabelInput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelInput) ProtoMessage()               {}
func (*GetDNSLabelInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type GetDNSLabelOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	DnsLabel         *string `protobuf:"bytes,4,opt,name=dns_label,json=dnsLabel" json:"dns_label,omitempty"`
	DomainName       *string `protobuf:"bytes,5,opt,name=domain_name,json=domainName" json:"domain_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDNSLabelOutput) Reset()                    { *m = GetDNSLabelOutput{} }
func (m *GetDNSLabelOutput) String() string            { return proto.CompactTextString(m) }
func (*GetDNSLabelOutput) ProtoMessage()               {}
func (*GetDNSLabelOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *GetDNSLabelOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetDNSLabelOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetDNSLabelOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GetDNSLabelOutput) GetDnsLabel() string {
	if m != nil && m.DnsLabel != nil {
		return *m.DnsLabel
	}
	return ""
}

func (m *GetDNSLabelOutput) GetDomainName() string {
	if m != nil && m.DomainName != nil {
		return *m.DomainName
	}
	return ""
}

func init() {
	proto.RegisterType((*DNSAliasServiceProperties)(nil), "service.DNSAliasServiceProperties")
	proto.RegisterType((*DescribeDNSAliasesInput)(nil), "service.DescribeDNSAliasesInput")
	proto.RegisterType((*DescribeDNSAliasesOutput)(nil), "service.DescribeDNSAliasesOutput")
	proto.RegisterType((*DescribeDNSAliasesOutput_ResponseItem)(nil), "service.DescribeDNSAliasesOutput.ResponseItem")
	proto.RegisterType((*AssociateDNSAliasInput)(nil), "service.AssociateDNSAliasInput")
	proto.RegisterType((*AssociateDNSAliasOutput)(nil), "service.AssociateDNSAliasOutput")
	proto.RegisterType((*DissociateDNSAliasesInput)(nil), "service.DissociateDNSAliasesInput")
	proto.RegisterType((*DissociateDNSAliasesOutput)(nil), "service.DissociateDNSAliasesOutput")
	proto.RegisterType((*GetDNSLabelInput)(nil), "service.GetDNSLabelInput")
	proto.RegisterType((*GetDNSLabelOutput)(nil), "service.GetDNSLabelOutput")
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x6f, 0x7e, 0xda, 0x9c, 0xf4, 0xfe, 0x74, 0xd4, 0x9b, 0x3a, 0xbe, 0xba, 0x6a, 0xae,
	0x2f, 0x8b, 0xac, 0x6c, 0xa9, 0xec, 0x80, 0x4d, 0xd5, 0x48, 0x28, 0x12, 0x14, 0x70, 0x10, 0x88,
	0x95, 0xe5, 0x78, 0x4e, 0xd2, 0x29, 0xb6, 0xc7, 0xcc, 0x8c, 0xa1, 0x62, 0xc5, 0x9a, 0x87, 0xe0,
	0x01, 0xe0, 0x0d, 0x58, 0xf1, 0x00, 0xbc, 0x0c, 0x6f, 0x80, 0x3c, 0x63, 0x27, 0x51, 0x92, 0xb6,
	0x6c, 0xba, 0xa9, 0x7a, 0x8e, 0xbf, 0x73, 0xe6, 0xfb, 0xce, 0x77, 0x66, 0x02, 0x1d, 0x9a, 0x49,
	0x2f, 0x17, 0x5c, 0x71, 0xb2, 0x23, 0x51, 0xbc, 0x65, 0x31, 0x3a, 0xff, 0xbe, 0x61, 0xd9, 0x3c,
	0x4e, 0x78, 0x41, 0x43, 0x49, 0x5f, 0x87, 0xa2, 0x48, 0xd0, 0x2f, 0xff, 0x18, 0x9c, 0x73, 0x34,
	0xe7, 0x7c, 0x9e, 0xa0, 0xaf, 0xa3, 0x69, 0x31, 0xf3, 0x15, 0x4b, 0x51, 0xaa, 0x28, 0xcd, 0x0d,
	0xc0, 0xf5, 0xa1, 0x3f, 0x3a, 0x9b, 0x9c, 0x24, 0x2c, 0x92, 0x13, 0xd3, 0xf2, 0xa9, 0xe0, 0x39,
	0x0a, 0xc5, 0x50, 0x12, 0x02, 0xcd, 0xf7, 0x3c, 0x43, 0xdb, 0x1a, 0x58, 0xc3, 0x4e, 0xa0, 0xff,
	0x77, 0x3f, 0x5b, 0x70, 0x38, 0x42, 0x19, 0x0b, 0x36, 0xc5, 0xba, 0x12, 0xe5, 0x38, 0xcb, 0x0b,
	0x45, 0x8e, 0xa0, 0x4b, 0x33, 0x19, 0x46, 0x26, 0x67, 0x5b, 0x83, 0xc6, 0xb0, 0x13, 0x00, 0xcd,
	0x64, 0x85, 0x2a, 0x01, 0x02, 0x25, 0x2f, 0x44, 0x8c, 0x21, 0xa3, 0xf6, 0x6f, 0xba, 0x2f, 0xd4,
	0xa9, 0x31, 0x2d, 0x01, 0x12, 0x23, 0x11, 0x9f, 0x87, 0xef, 0xb8, 0xa0, 0x76, 0xc3, 0x00, 0x4c,
	0xea, 0x25, 0x17, 0x94, 0xf4, 0xa0, 0xcd, 0x67, 0x33, 0x89, 0xca, 0x6e, 0x0e, 0xac, 0x61, 0x2b,
	0xa8, 0x22, 0x72, 0x00, 0xad, 0x84, 0xa5, 0x4c, 0xd9, 0x2d, 0x9d, 0x36, 0x81, 0xfb, 0xa5, 0x01,
	0xf6, 0x26, 0xd9, 0x27, 0x85, 0x2a, 0xd9, 0xf6, 0xa0, 0x1d, 0xc5, 0x8a, 0xf1, 0xac, 0xd2, 0x57,
	0x45, 0xa4, 0x0f, 0xbb, 0x02, 0x55, 0x18, 0x73, 0x8a, 0x9a, 0x61, 0x2b, 0xd8, 0x11, 0xa8, 0x4e,
	0x39, 0x45, 0x62, 0xc3, 0x4e, 0x8a, 0x52, 0x46, 0x73, 0xac, 0xa8, 0xd5, 0x21, 0x09, 0xe0, 0xf7,
	0x85, 0xf4, 0xd0, 0xd0, 0x6b, 0x0c, 0xbb, 0xc7, 0x9e, 0x57, 0x19, 0xe5, 0x5d, 0x45, 0xc3, 0x0b,
	0x50, 0xe6, 0x3c, 0x93, 0x38, 0x56, 0x98, 0x06, 0xdd, 0x7a, 0x58, 0x13, 0xd4, 0xe3, 0x54, 0x5c,
	0x45, 0x49, 0x18, 0xf3, 0x22, 0xab, 0x95, 0x81, 0x4e, 0x9d, 0x96, 0x19, 0xe7, 0xbb, 0x05, 0x7b,
	0xab, 0xe5, 0xa5, 0x24, 0xa9, 0x22, 0x55, 0xc8, 0x5a, 0x92, 0x89, 0xc8, 0x00, 0xf6, 0x96, 0xec,
	0x96, 0x83, 0xaf, 0x0f, 0x33, 0x83, 0x5f, 0x75, 0xa6, 0xb1, 0xe1, 0xcc, 0x1d, 0xf8, 0x63, 0xd9,
	0x22, 0x8b, 0x52, 0xd4, 0x06, 0x74, 0x82, 0xbd, 0xba, 0xc9, 0x59, 0x94, 0x22, 0xb9, 0x0f, 0xdd,
	0x58, 0x60, 0xa4, 0x30, 0x2c, 0x17, 0x4d, 0x53, 0xee, 0x1e, 0x3b, 0x9e, 0xd9, 0x42, 0xaf, 0xde,
	0x42, 0xef, 0x79, 0xbd, 0x85, 0x01, 0x18, 0x78, 0x99, 0x70, 0x05, 0xf4, 0x4e, 0xa4, 0xe4, 0x31,
	0x8b, 0xd4, 0x62, 0x4c, 0x66, 0xb1, 0x7a, 0xd0, 0xce, 0x05, 0xce, 0xd8, 0x65, 0xad, 0xcb, 0x44,
	0xc4, 0x29, 0xad, 0x32, 0x14, 0x2b, 0x4d, 0x8b, 0x78, 0x0b, 0xe1, 0xc6, 0x26, 0x61, 0xf7, 0x9b,
	0x05, 0x87, 0x1b, 0x87, 0xde, 0xc6, 0x82, 0xac, 0x5b, 0xd0, 0xdc, 0x66, 0x01, 0xe5, 0x69, 0xc4,
	0x32, 0xc3, 0xb6, 0x55, 0x01, 0x74, 0x4a, 0x0f, 0xf7, 0x6f, 0x68, 0x5f, 0xf0, 0x69, 0x59, 0xdc,
	0xd6, 0xdf, 0x5a, 0x17, 0x7c, 0x3a, 0xa6, 0xee, 0x03, 0xe8, 0x8f, 0xd8, 0xba, 0x84, 0x5f, 0xbd,
	0x92, 0xee, 0x07, 0x0b, 0x9c, 0x6d, 0xe5, 0xb7, 0x31, 0x83, 0xa5, 0x80, 0xe6, 0xaa, 0x00, 0x02,
	0x7f, 0x3d, 0x44, 0x35, 0x3a, 0x9b, 0x3c, 0x8a, 0xa6, 0x98, 0x68, 0xde, 0xee, 0x27, 0x0b, 0xf6,
	0x57, 0x92, 0xb7, 0xc1, 0xe6, 0x1f, 0xfd, 0xa0, 0x86, 0x49, 0xd9, 0xbf, 0x22, 0xb4, 0x4b, 0x33,
	0xa9, 0xcf, 0xbb, 0xd1, 0x8c, 0xe3, 0xaf, 0x0d, 0xf8, 0x73, 0xed, 0xe5, 0x24, 0xaf, 0x80, 0x6c,
	0x5e, 0x73, 0x32, 0xb8, 0xe6, 0x0d, 0xd0, 0x62, 0x9d, 0xff, 0x6e, 0x7c, 0x25, 0xc8, 0x0b, 0xd8,
	0xdf, 0x58, 0x53, 0x72, 0xb4, 0xa8, 0xdb, 0x7e, 0x6f, 0x9c, 0xc1, 0xd5, 0x80, 0xaa, 0x6f, 0x08,
	0x07, 0xdb, 0xdc, 0x27, 0xee, 0x92, 0xd2, 0x55, 0xbb, 0xe5, 0xfc, 0x7f, 0x2d, 0xa6, 0x3a, 0x60,
	0x04, 0xdd, 0x15, 0x1f, 0x49, 0x7f, 0x51, 0xb3, 0x6e, 0xb9, 0xe3, 0x6c, 0xfb, 0x64, 0xba, 0x38,
	0x8f, 0x3f, 0xfe, 0x68, 0x8e, 0xe1, 0xee, 0xb9, 0x52, 0xb9, 0xbc, 0xe7, 0xfb, 0x94, 0xc7, 0xd2,
	0x5b, 0xfc, 0xf2, 0x79, 0x31, 0x4f, 0xfd, 0x28, 0x67, 0xfe, 0x62, 0xd3, 0x7d, 0x96, 0x51, 0xbc,
	0xf4, 0xce, 0x55, 0x9a, 0x10, 0xf2, 0x8c, 0x65, 0xf3, 0x53, 0x8d, 0xab, 0xc9, 0xfd, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0x1a, 0xb6, 0xb1, 0x4a, 0x07, 0x00, 0x00,
}
