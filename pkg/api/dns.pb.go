// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
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
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0xdf, 0x7c, 0xb4, 0x99, 0xf4, 0x7d, 0x5f, 0xba, 0x2a, 0xa9, 0x63, 0x0e, 0x0d, 0x86,
	0x43, 0x4e, 0xb6, 0x54, 0x84, 0x90, 0xa0, 0x97, 0xaa, 0x91, 0x50, 0xa4, 0xaa, 0x20, 0x07, 0x81,
	0x38, 0x59, 0x1b, 0xef, 0x24, 0xdd, 0xca, 0xf6, 0x1a, 0xef, 0x06, 0x2a, 0x4e, 0xfc, 0x0a, 0x8e,
	0xdc, 0x81, 0x7f, 0xc0, 0x89, 0x1f, 0xc0, 0x1f, 0xe2, 0x86, 0xbc, 0x6b, 0x27, 0x51, 0x92, 0xb6,
	0x5c, 0x7a, 0xf3, 0x8c, 0x9f, 0x99, 0x7d, 0x9e, 0x79, 0x66, 0x17, 0x5a, 0x2c, 0x95, 0x5e, 0x96,
	0x0b, 0x25, 0xc8, 0x96, 0xc4, 0xfc, 0x3d, 0x8f, 0xd0, 0xb1, 0x65, 0x86, 0x51, 0x98, 0xa0, 0xa2,
	0x8c, 0x2a, 0xea, 0x17, 0x91, 0x81, 0x38, 0x07, 0x53, 0x21, 0xa6, 0x31, 0xfa, 0x3a, 0x1a, 0xcf,
	0x26, 0xbe, 0xe2, 0x09, 0x4a, 0x45, 0x93, 0xcc, 0x00, 0x5c, 0x1f, 0xba, 0x83, 0xb3, 0xd1, 0x71,
	0xcc, 0xa9, 0x1c, 0x99, 0x6e, 0x2f, 0x73, 0x91, 0x61, 0xae, 0x38, 0x4a, 0x42, 0xa0, 0xfe, 0x51,
	0xa4, 0x68, 0x5b, 0x3d, 0xab, 0xdf, 0x0a, 0xf4, 0xb7, 0xfb, 0xcd, 0x82, 0xfd, 0x01, 0xca, 0x28,
	0xe7, 0x63, 0xac, 0x2a, 0x51, 0x0e, 0xd3, 0x6c, 0xa6, 0xc8, 0x01, 0xb4, 0x59, 0x2a, 0x43, 0x6a,
	0x72, 0xb6, 0xd5, 0xab, 0xf5, 0x5b, 0x01, 0xb0, 0x54, 0x96, 0xa8, 0x02, 0x90, 0xa3, 0x14, 0xb3,
	0x3c, 0xc2, 0x90, 0x33, 0xfb, 0x1f, 0xdd, 0x17, 0xaa, 0xd4, 0x90, 0x15, 0x00, 0x89, 0x34, 0x8f,
	0xce, 0xc3, 0x0f, 0x22, 0x67, 0x76, 0xcd, 0x00, 0x4c, 0xea, 0x8d, 0xc8, 0x19, 0xe9, 0x40, 0x53,
	0x4c, 0x26, 0x12, 0x95, 0x5d, 0xef, 0x59, 0xfd, 0x46, 0x50, 0x46, 0x64, 0x0f, 0x1a, 0x31, 0x4f,
	0xb8, 0xb2, 0x1b, 0x3a, 0x6d, 0x02, 0xf7, 0x7b, 0x0d, 0xec, 0x75, 0xb2, 0x2f, 0x66, 0xaa, 0x60,
	0xdb, 0x81, 0x26, 0x8d, 0x14, 0x17, 0x69, 0xa9, 0xaf, 0x8c, 0x48, 0x17, 0xb6, 0x73, 0x54, 0x61,
	0x24, 0x18, 0x6a, 0x86, 0x8d, 0x60, 0x2b, 0x47, 0x75, 0x22, 0x18, 0x12, 0x1b, 0xb6, 0x12, 0x94,
	0x92, 0x4e, 0xb1, 0xa4, 0x56, 0x85, 0x24, 0x80, 0x7f, 0xe7, 0xd2, 0x43, 0x43, 0xaf, 0xd6, 0x6f,
	0x1f, 0x7a, 0x5e, 0xe9, 0x91, 0x77, 0x15, 0x0d, 0x2f, 0x40, 0x99, 0x89, 0x54, 0xe2, 0x50, 0x61,
	0x12, 0xb4, 0xab, 0x61, 0x8d, 0x50, 0x8f, 0x53, 0x09, 0x45, 0xe3, 0x30, 0x12, 0xb3, 0xb4, 0x52,
	0x06, 0x3a, 0x75, 0x52, 0x64, 0x9c, 0x5f, 0x16, 0xec, 0x2c, 0x97, 0x17, 0x92, 0xa4, 0xa2, 0x6a,
	0x26, 0x2b, 0x49, 0x26, 0x22, 0x3d, 0xd8, 0x59, 0xb0, 0x5b, 0x0c, 0xbe, 0x3a, 0xcc, 0x0c, 0x7e,
	0xd9, 0x99, 0xda, 0x9a, 0x33, 0x0f, 0xe1, 0xbf, 0x45, 0x8b, 0x94, 0x26, 0xa8, 0x0d, 0x68, 0x05,
	0x3b, 0x55, 0x93, 0x33, 0x9a, 0x20, 0x79, 0x06, 0xed, 0x28, 0x47, 0xaa, 0x30, 0x2c, 0x16, 0x4d,
	0x53, 0x6e, 0x1f, 0x3a, 0x9e, 0xd9, 0x42, 0xaf, 0xda, 0x42, 0xef, 0x55, 0xb5, 0x85, 0x01, 0x18,
	0x78, 0x91, 0x70, 0x73, 0xe8, 0x1c, 0x4b, 0x29, 0x22, 0x4e, 0xd5, 0x7c, 0x4c, 0x66, 0xb1, 0x3a,
	0xd0, 0xcc, 0x72, 0x9c, 0xf0, 0xcb, 0x4a, 0x97, 0x89, 0x88, 0x53, 0x58, 0x65, 0x28, 0x96, 0x9a,
	0xe6, 0xf1, 0x06, 0xc2, 0xb5, 0x75, 0xc2, 0xee, 0x4f, 0x0b, 0xf6, 0xd7, 0x0e, 0xbd, 0x8d, 0x05,
	0x59, 0xb5, 0xa0, 0xbe, 0xc9, 0x02, 0x26, 0x12, 0xca, 0x53, 0xc3, 0xb6, 0x51, 0x02, 0x74, 0x4a,
	0x0f, 0xf7, 0x2e, 0x34, 0x2f, 0xc4, 0xb8, 0x28, 0x6e, 0xea, 0x7f, 0x8d, 0x0b, 0x31, 0x1e, 0x32,
	0xf7, 0x08, 0xba, 0x03, 0xbe, 0x2a, 0xe1, 0x6f, 0xaf, 0xa4, 0xfb, 0xc9, 0x02, 0x67, 0x53, 0xf9,
	0x6d, 0xcc, 0x60, 0x21, 0xa0, 0xbe, 0x2c, 0x80, 0xc0, 0x9d, 0xe7, 0xa8, 0x06, 0x67, 0xa3, 0x53,
	0x3a, 0xc6, 0x58, 0xf3, 0x76, 0xbf, 0x58, 0xb0, 0xbb, 0x94, 0xbc, 0x0d, 0x36, 0xf7, 0xf4, 0x5b,
	0x1a, 0xc6, 0x45, 0xff, 0x92, 0xd0, 0x36, 0x4b, 0xa5, 0x3e, 0xef, 0x46, 0x33, 0x0e, 0xbf, 0xd6,
	0xe0, 0xff, 0x95, 0x97, 0x93, 0xbc, 0x05, 0xb2, 0x7e, 0xcd, 0x49, 0xef, 0x9a, 0x37, 0x40, 0x8b,
	0x75, 0xee, 0xdf, 0xf8, 0x4a, 0x90, 0xd7, 0xb0, 0xbb, 0xb6, 0xa6, 0xe4, 0x60, 0x5e, 0xb7, 0xf9,
	0xde, 0x38, 0xbd, 0xab, 0x01, 0x65, 0xdf, 0x10, 0xf6, 0x36, 0xb9, 0x4f, 0xdc, 0x05, 0xa5, 0xab,
	0x76, 0xcb, 0x79, 0x70, 0x2d, 0xa6, 0x3c, 0x60, 0x00, 0xed, 0x25, 0x1f, 0x49, 0x77, 0x5e, 0xb3,
	0x6a, 0xb9, 0xe3, 0x6c, 0xfa, 0x65, 0xba, 0x38, 0x47, 0x3f, 0x3e, 0xff, 0x3e, 0x7d, 0x02, 0x8f,
	0x9d, 0x47, 0xe7, 0x4a, 0x65, 0xf2, 0xa9, 0xef, 0x33, 0x11, 0x49, 0xef, 0x1d, 0x4f, 0xa7, 0x51,
	0x2c, 0x66, 0xcc, 0x8b, 0x44, 0xe2, 0xd3, 0x8c, 0xfb, 0xf3, 0x65, 0xf7, 0x79, 0xca, 0xf0, 0xd2,
	0x3b, 0x57, 0x49, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd3, 0x0b, 0x68, 0x34, 0x07, 0x00,
	0x00,
}
