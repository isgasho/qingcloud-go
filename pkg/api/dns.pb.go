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
func (*DescribeDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*DescribeDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
	return fileDescriptor3, []int{1, 0}
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
func (*AssociateDNSAliasInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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
func (*AssociateDNSAliasOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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
func (*DissociateDNSAliasesInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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
func (*DissociateDNSAliasesOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

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
func (*GetDNSLabelInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

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
func (*GetDNSLabelOutput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

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
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x71, 0x93, 0xb6, 0x93, 0xf2, 0xd3, 0x55, 0x49, 0x5d, 0x73, 0x68, 0x30, 0x1c, 0x7a,
	0xb2, 0xa5, 0x22, 0x84, 0x04, 0xbd, 0x54, 0x8d, 0x84, 0x22, 0x55, 0x45, 0x72, 0x10, 0x88, 0x93,
	0xb5, 0xf1, 0x4e, 0xd2, 0xad, 0x6c, 0xaf, 0xf1, 0xae, 0xa1, 0x47, 0x9e, 0x82, 0x23, 0x77, 0xe0,
	0x0d, 0x38, 0xf1, 0x00, 0xbc, 0x10, 0x37, 0xe4, 0x5d, 0x3b, 0x89, 0x9a, 0xb4, 0xe5, 0x92, 0xe3,
	0x8c, 0xbf, 0x99, 0xf9, 0xbe, 0xf9, 0x66, 0x0d, 0x9b, 0x2c, 0x93, 0x7e, 0x5e, 0x08, 0x25, 0xc8,
	0xba, 0xc4, 0xe2, 0x13, 0x8f, 0xd1, 0x75, 0x64, 0x8e, 0x71, 0x94, 0xa2, 0xa2, 0x8c, 0x2a, 0x1a,
	0x54, 0x91, 0x81, 0xb8, 0xfb, 0x13, 0x21, 0x26, 0x09, 0x06, 0x3a, 0x1a, 0x95, 0xe3, 0x40, 0xf1,
	0x14, 0xa5, 0xa2, 0x69, 0x6e, 0x00, 0xde, 0x0f, 0x0b, 0x76, 0xfb, 0x28, 0xe3, 0x82, 0x8f, 0xb0,
	0x7f, 0x36, 0x3c, 0x4e, 0x38, 0x95, 0x28, 0x07, 0x59, 0x5e, 0x2a, 0xb2, 0x0f, 0x1d, 0x96, 0xc9,
	0x88, 0x9a, 0x9c, 0x63, 0xf5, 0xec, 0x83, 0xcd, 0x10, 0x58, 0x26, 0x6b, 0x54, 0x05, 0x28, 0x50,
	0x8a, 0xb2, 0x88, 0x31, 0xe2, 0xcc, 0xb9, 0xd3, 0xb3, 0x2a, 0x40, 0x93, 0x1a, 0xb0, 0x0a, 0x20,
	0x91, 0x16, 0xf1, 0x79, 0xf4, 0x59, 0x14, 0xcc, 0xb1, 0x0d, 0xc0, 0xa4, 0xde, 0x8b, 0x82, 0x91,
	0x2e, 0xb4, 0xc5, 0x78, 0x2c, 0x51, 0x39, 0x6b, 0x3d, 0xeb, 0xa0, 0x15, 0xd6, 0x11, 0xd9, 0x81,
	0x56, 0xc2, 0x53, 0xae, 0x9c, 0x96, 0x4e, 0x9b, 0xc0, 0xfb, 0x69, 0x83, 0xb3, 0x48, 0xf6, 0x4d,
	0xa9, 0x2a, 0xb6, 0x5d, 0x68, 0xd3, 0x58, 0x71, 0x91, 0x39, 0x96, 0x1e, 0x53, 0x47, 0x64, 0x0f,
	0x36, 0x0a, 0x54, 0x51, 0x2c, 0x18, 0x6a, 0x86, 0xad, 0x70, 0xbd, 0x40, 0x75, 0x22, 0x18, 0x12,
	0x07, 0xd6, 0x53, 0x94, 0x92, 0x4e, 0xb0, 0xa6, 0xd6, 0x84, 0x24, 0x84, 0xbb, 0x53, 0xe9, 0x91,
	0xa1, 0x67, 0x1f, 0x74, 0x0e, 0x7d, 0xbf, 0x5e, 0xb9, 0x7f, 0x1d, 0x0d, 0x3f, 0x44, 0x99, 0x8b,
	0x4c, 0xe2, 0x40, 0x61, 0x1a, 0x76, 0x9a, 0x65, 0x0d, 0x51, 0xaf, 0x53, 0x09, 0x45, 0x93, 0x28,
	0x16, 0x65, 0xd6, 0x28, 0x03, 0x9d, 0x3a, 0xa9, 0x32, 0xee, 0x1f, 0x0b, 0xb6, 0xe6, 0xcb, 0x2b,
	0x49, 0x52, 0x51, 0x55, 0xca, 0x46, 0x92, 0x89, 0x48, 0x0f, 0xb6, 0x66, 0xec, 0x66, 0x8b, 0x6f,
	0x86, 0x99, 0xc5, 0xcf, 0x3b, 0x63, 0x2f, 0x38, 0xf3, 0x14, 0xee, 0xcd, 0x5a, 0x64, 0x34, 0x45,
	0x6d, 0xc0, 0x66, 0xb8, 0xd5, 0x34, 0x39, 0xa3, 0x29, 0x92, 0x57, 0xd0, 0x89, 0x0b, 0xa4, 0x0a,
	0xa3, 0xea, 0x6e, 0x34, 0xe5, 0xce, 0xa1, 0xeb, 0x9b, 0xa3, 0xf2, 0x9b, 0xa3, 0xf2, 0xdf, 0x36,
	0x47, 0x15, 0x82, 0x81, 0x57, 0x09, 0xaf, 0x80, 0xee, 0xb1, 0x94, 0x22, 0xe6, 0x54, 0x4d, 0xd7,
	0x64, 0x0e, 0xab, 0x0b, 0xed, 0xbc, 0xc0, 0x31, 0xbf, 0x6c, 0x74, 0x99, 0x88, 0xb8, 0x95, 0x55,
	0x86, 0x62, 0xad, 0x69, 0x1a, 0x2f, 0x21, 0x6c, 0x2f, 0x12, 0xf6, 0x7e, 0x5b, 0xb0, 0xbb, 0x30,
	0x74, 0x15, 0x07, 0x72, 0xd5, 0x82, 0xb5, 0x65, 0x16, 0x30, 0x91, 0x52, 0x9e, 0x19, 0xb6, 0xad,
	0x1a, 0xa0, 0x53, 0x7a, 0xb9, 0x0f, 0xa1, 0x7d, 0x21, 0x46, 0x55, 0x71, 0x5b, 0x7f, 0x6b, 0x5d,
	0x88, 0xd1, 0x80, 0x79, 0x47, 0xb0, 0xd7, 0xe7, 0x57, 0x25, 0xfc, 0xef, 0x93, 0xf4, 0xbe, 0x58,
	0xe0, 0x2e, 0x2b, 0x5f, 0xc5, 0x0e, 0x66, 0x02, 0xd6, 0xe6, 0x05, 0x10, 0x78, 0xf0, 0x1a, 0x55,
	0xff, 0x6c, 0x78, 0x4a, 0x47, 0x98, 0x68, 0xde, 0xde, 0x37, 0x0b, 0xb6, 0xe7, 0x92, 0xab, 0x60,
	0xf3, 0x48, 0xff, 0x1a, 0xa3, 0xa4, 0xea, 0x5f, 0x13, 0xda, 0x60, 0x99, 0xd4, 0xf3, 0x6e, 0x35,
	0xe3, 0xf0, 0xbb, 0x0d, 0xf7, 0x9b, 0x6d, 0x0d, 0xcd, 0x1b, 0x27, 0x1f, 0x80, 0x2c, 0x3e, 0x73,
	0xd2, 0xbb, 0xe1, 0x1f, 0xa0, 0xc5, 0xba, 0x8f, 0x6f, 0xfd, 0x4b, 0x90, 0x77, 0xb0, 0xbd, 0x70,
	0xa6, 0x64, 0x7f, 0x5a, 0xb7, 0xfc, 0xdd, 0xb8, 0xbd, 0xeb, 0x01, 0x75, 0xdf, 0x08, 0x76, 0x96,
	0xb9, 0x4f, 0xbc, 0x19, 0xa5, 0xeb, 0x6e, 0xcb, 0x7d, 0x72, 0x23, 0xa6, 0x1e, 0xd0, 0x87, 0xce,
	0x9c, 0x8f, 0x64, 0x6f, 0x5a, 0x73, 0xd5, 0x72, 0xd7, 0x5d, 0xf6, 0xc9, 0x74, 0x71, 0x8f, 0x7e,
	0x7d, 0xfd, 0x7b, 0xfa, 0x02, 0x9e, 0xbb, 0xcf, 0xce, 0x95, 0xca, 0xe5, 0xcb, 0x20, 0x60, 0x22,
	0x96, 0xfe, 0x47, 0x9e, 0x4d, 0xe2, 0x44, 0x94, 0xcc, 0x8f, 0x45, 0x1a, 0xd0, 0x9c, 0x07, 0xd3,
	0x63, 0x0f, 0x78, 0xc6, 0xf0, 0xd2, 0x3f, 0x57, 0x69, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0x9d, 0x8b, 0x62, 0x03, 0x07, 0x00, 0x00,
}
