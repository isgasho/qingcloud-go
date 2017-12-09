// Code generated by protoc-gen-go. DO NOT EDIT.
// source: span.proto

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

type CreateSpanInput struct {
	SpanName         *string `protobuf:"bytes,1,opt,name=span_name,json=spanName" json:"span_name,omitempty"`
	Flag             *int32  `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
	IpAddr           *string `protobuf:"bytes,3,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
	TunnelType       *string `protobuf:"bytes,4,opt,name=tunnel_type,json=tunnelType" json:"tunnel_type,omitempty"`
	TunnelKey        *int32  `protobuf:"varint,5,opt,name=tunnel_key,json=tunnelKey" json:"tunnel_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateSpanInput) Reset()                    { *m = CreateSpanInput{} }
func (m *CreateSpanInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSpanInput) ProtoMessage()               {}
func (*CreateSpanInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *CreateSpanInput) GetSpanName() string {
	if m != nil && m.SpanName != nil {
		return *m.SpanName
	}
	return ""
}

func (m *CreateSpanInput) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *CreateSpanInput) GetIpAddr() string {
	if m != nil && m.IpAddr != nil {
		return *m.IpAddr
	}
	return ""
}

func (m *CreateSpanInput) GetTunnelType() string {
	if m != nil && m.TunnelType != nil {
		return *m.TunnelType
	}
	return ""
}

func (m *CreateSpanInput) GetTunnelKey() int32 {
	if m != nil && m.TunnelKey != nil {
		return *m.TunnelKey
	}
	return 0
}

type CreateSpanOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SpanId           *string `protobuf:"bytes,4,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateSpanOutput) Reset()                    { *m = CreateSpanOutput{} }
func (m *CreateSpanOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSpanOutput) ProtoMessage()               {}
func (*CreateSpanOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *CreateSpanOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *CreateSpanOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *CreateSpanOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *CreateSpanOutput) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

type DescribeSpansInput struct {
	Spans            []string `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
	SpanName         *string  `protobuf:"bytes,2,opt,name=span_name,json=spanName" json:"span_name,omitempty"`
	IpAddr           *string  `protobuf:"bytes,3,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
	Tags             []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Offset           *int32   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Limit            *int32   `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DescribeSpansInput) Reset()                    { *m = DescribeSpansInput{} }
func (m *DescribeSpansInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSpansInput) ProtoMessage()               {}
func (*DescribeSpansInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func (m *DescribeSpansInput) GetSpans() []string {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *DescribeSpansInput) GetSpanName() string {
	if m != nil && m.SpanName != nil {
		return *m.SpanName
	}
	return ""
}

func (m *DescribeSpansInput) GetIpAddr() string {
	if m != nil && m.IpAddr != nil {
		return *m.IpAddr
	}
	return ""
}

func (m *DescribeSpansInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeSpansInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeSpansInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type DescribeSpansOutput struct {
	Action           *string                             `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32                              `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string                             `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SpanSet          []*DescribeSpansOutput_ResponseItem `protobuf:"bytes,4,rep,name=span_set,json=spanSet" json:"span_set,omitempty"`
	TotalCount       *int32                              `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *DescribeSpansOutput) Reset()                    { *m = DescribeSpansOutput{} }
func (m *DescribeSpansOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSpansOutput) ProtoMessage()               {}
func (*DescribeSpansOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{3} }

func (m *DescribeSpansOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeSpansOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeSpansOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeSpansOutput) GetSpanSet() []*DescribeSpansOutput_ResponseItem {
	if m != nil {
		return m.SpanSet
	}
	return nil
}

func (m *DescribeSpansOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type DescribeSpansOutput_ResponseItem struct {
	IsApplied        *int32                                         `protobuf:"varint,1,opt,name=is_applied,json=isApplied" json:"is_applied,omitempty"`
	SpanName         *string                                        `protobuf:"bytes,2,opt,name=span_name,json=spanName" json:"span_name,omitempty"`
	IpAddr           *string                                        `protobuf:"bytes,3,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
	TunnelType       *string                                        `protobuf:"bytes,4,opt,name=tunnel_type,json=tunnelType" json:"tunnel_type,omitempty"`
	Tags             []string                                       `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	Flag             *int32                                         `protobuf:"varint,6,opt,name=flag" json:"flag,omitempty"`
	TunnelKey        *int32                                         `protobuf:"varint,7,opt,name=tunnel_key,json=tunnelKey" json:"tunnel_key,omitempty"`
	CreateTime       *google_protobuf1.Timestamp                    `protobuf:"bytes,8,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime       *google_protobuf1.Timestamp                    `protobuf:"bytes,9,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	SpanId           *string                                        `protobuf:"bytes,10,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Members          []*DescribeSpansOutput_ResponseItem_MemberItem `protobuf:"bytes,11,rep,name=members" json:"members,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *DescribeSpansOutput_ResponseItem) Reset()         { *m = DescribeSpansOutput_ResponseItem{} }
func (m *DescribeSpansOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeSpansOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeSpansOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{3, 0}
}

func (m *DescribeSpansOutput_ResponseItem) GetIsApplied() int32 {
	if m != nil && m.IsApplied != nil {
		return *m.IsApplied
	}
	return 0
}

func (m *DescribeSpansOutput_ResponseItem) GetSpanName() string {
	if m != nil && m.SpanName != nil {
		return *m.SpanName
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem) GetIpAddr() string {
	if m != nil && m.IpAddr != nil {
		return *m.IpAddr
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem) GetTunnelType() string {
	if m != nil && m.TunnelType != nil {
		return *m.TunnelType
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeSpansOutput_ResponseItem) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *DescribeSpansOutput_ResponseItem) GetTunnelKey() int32 {
	if m != nil && m.TunnelKey != nil {
		return *m.TunnelKey
	}
	return 0
}

func (m *DescribeSpansOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeSpansOutput_ResponseItem) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *DescribeSpansOutput_ResponseItem) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem) GetMembers() []*DescribeSpansOutput_ResponseItem_MemberItem {
	if m != nil {
		return m.Members
	}
	return nil
}

type DescribeSpansOutput_ResponseItem_MemberItem struct {
	SpanId           *string                     `protobuf:"bytes,1,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	ResourceId       *string                     `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	Status           *string                     `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	CreateTime       *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime       *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) Reset() {
	*m = DescribeSpansOutput_ResponseItem_MemberItem{}
}
func (m *DescribeSpansOutput_ResponseItem_MemberItem) String() string {
	return proto.CompactTextString(m)
}
func (*DescribeSpansOutput_ResponseItem_MemberItem) ProtoMessage() {}
func (*DescribeSpansOutput_ResponseItem_MemberItem) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{3, 0, 0}
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) GetResourceId() string {
	if m != nil && m.ResourceId != nil {
		return *m.ResourceId
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeSpansOutput_ResponseItem_MemberItem) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DeleteSpansInput struct {
	Spans            []string `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeleteSpansInput) Reset()                    { *m = DeleteSpansInput{} }
func (m *DeleteSpansInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSpansInput) ProtoMessage()               {}
func (*DeleteSpansInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{4} }

func (m *DeleteSpansInput) GetSpans() []string {
	if m != nil {
		return m.Spans
	}
	return nil
}

type DeleteSpansOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteSpansOutput) Reset()                    { *m = DeleteSpansOutput{} }
func (m *DeleteSpansOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSpansOutput) ProtoMessage()               {}
func (*DeleteSpansOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{5} }

func (m *DeleteSpansOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteSpansOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteSpansOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type AddSpanMembersInput struct {
	Span             *string  `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
	Resources        []string `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AddSpanMembersInput) Reset()                    { *m = AddSpanMembersInput{} }
func (m *AddSpanMembersInput) String() string            { return proto.CompactTextString(m) }
func (*AddSpanMembersInput) ProtoMessage()               {}
func (*AddSpanMembersInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{6} }

func (m *AddSpanMembersInput) GetSpan() string {
	if m != nil && m.Span != nil {
		return *m.Span
	}
	return ""
}

func (m *AddSpanMembersInput) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

type AddSpanMembersOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AddSpanMembersOutput) Reset()                    { *m = AddSpanMembersOutput{} }
func (m *AddSpanMembersOutput) String() string            { return proto.CompactTextString(m) }
func (*AddSpanMembersOutput) ProtoMessage()               {}
func (*AddSpanMembersOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{7} }

func (m *AddSpanMembersOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AddSpanMembersOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AddSpanMembersOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AddSpanMembersOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type RemoveSpanMembersInput struct {
	Span             *string  `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
	Resources        []string `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RemoveSpanMembersInput) Reset()                    { *m = RemoveSpanMembersInput{} }
func (m *RemoveSpanMembersInput) String() string            { return proto.CompactTextString(m) }
func (*RemoveSpanMembersInput) ProtoMessage()               {}
func (*RemoveSpanMembersInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{8} }

func (m *RemoveSpanMembersInput) GetSpan() string {
	if m != nil && m.Span != nil {
		return *m.Span
	}
	return ""
}

func (m *RemoveSpanMembersInput) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

type RemoveSpanMembersOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoveSpanMembersOutput) Reset()                    { *m = RemoveSpanMembersOutput{} }
func (m *RemoveSpanMembersOutput) String() string            { return proto.CompactTextString(m) }
func (*RemoveSpanMembersOutput) ProtoMessage()               {}
func (*RemoveSpanMembersOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{9} }

func (m *RemoveSpanMembersOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *RemoveSpanMembersOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *RemoveSpanMembersOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *RemoveSpanMembersOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ModifySpanAttributesInput struct {
	SpanId           *string `protobuf:"bytes,1,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	SpanName         *string `protobuf:"bytes,2,opt,name=span_name,json=spanName" json:"span_name,omitempty"`
	Flag             *int32  `protobuf:"varint,3,opt,name=flag" json:"flag,omitempty"`
	IpAddr           *string `protobuf:"bytes,4,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
	TunnelType       *string `protobuf:"bytes,5,opt,name=tunnel_type,json=tunnelType" json:"tunnel_type,omitempty"`
	TunnelKey        *int32  `protobuf:"varint,6,opt,name=tunnel_key,json=tunnelKey" json:"tunnel_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifySpanAttributesInput) Reset()                    { *m = ModifySpanAttributesInput{} }
func (m *ModifySpanAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySpanAttributesInput) ProtoMessage()               {}
func (*ModifySpanAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{10} }

func (m *ModifySpanAttributesInput) GetSpanId() string {
	if m != nil && m.SpanId != nil {
		return *m.SpanId
	}
	return ""
}

func (m *ModifySpanAttributesInput) GetSpanName() string {
	if m != nil && m.SpanName != nil {
		return *m.SpanName
	}
	return ""
}

func (m *ModifySpanAttributesInput) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *ModifySpanAttributesInput) GetIpAddr() string {
	if m != nil && m.IpAddr != nil {
		return *m.IpAddr
	}
	return ""
}

func (m *ModifySpanAttributesInput) GetTunnelType() string {
	if m != nil && m.TunnelType != nil {
		return *m.TunnelType
	}
	return ""
}

func (m *ModifySpanAttributesInput) GetTunnelKey() int32 {
	if m != nil && m.TunnelKey != nil {
		return *m.TunnelKey
	}
	return 0
}

type ModifySpanAttributesOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifySpanAttributesOutput) Reset()                    { *m = ModifySpanAttributesOutput{} }
func (m *ModifySpanAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifySpanAttributesOutput) ProtoMessage()               {}
func (*ModifySpanAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{11} }

func (m *ModifySpanAttributesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ModifySpanAttributesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ModifySpanAttributesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type UpdateSpanInput struct {
	Span             *string `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateSpanInput) Reset()                    { *m = UpdateSpanInput{} }
func (m *UpdateSpanInput) String() string            { return proto.CompactTextString(m) }
func (*UpdateSpanInput) ProtoMessage()               {}
func (*UpdateSpanInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{12} }

func (m *UpdateSpanInput) GetSpan() string {
	if m != nil && m.Span != nil {
		return *m.Span
	}
	return ""
}

type UpdateSpanOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateSpanOutput) Reset()                    { *m = UpdateSpanOutput{} }
func (m *UpdateSpanOutput) String() string            { return proto.CompactTextString(m) }
func (*UpdateSpanOutput) ProtoMessage()               {}
func (*UpdateSpanOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{13} }

func (m *UpdateSpanOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *UpdateSpanOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *UpdateSpanOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *UpdateSpanOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSpanInput)(nil), "service.CreateSpanInput")
	proto.RegisterType((*CreateSpanOutput)(nil), "service.CreateSpanOutput")
	proto.RegisterType((*DescribeSpansInput)(nil), "service.DescribeSpansInput")
	proto.RegisterType((*DescribeSpansOutput)(nil), "service.DescribeSpansOutput")
	proto.RegisterType((*DescribeSpansOutput_ResponseItem)(nil), "service.DescribeSpansOutput.ResponseItem")
	proto.RegisterType((*DescribeSpansOutput_ResponseItem_MemberItem)(nil), "service.DescribeSpansOutput.ResponseItem.MemberItem")
	proto.RegisterType((*DeleteSpansInput)(nil), "service.DeleteSpansInput")
	proto.RegisterType((*DeleteSpansOutput)(nil), "service.DeleteSpansOutput")
	proto.RegisterType((*AddSpanMembersInput)(nil), "service.AddSpanMembersInput")
	proto.RegisterType((*AddSpanMembersOutput)(nil), "service.AddSpanMembersOutput")
	proto.RegisterType((*RemoveSpanMembersInput)(nil), "service.RemoveSpanMembersInput")
	proto.RegisterType((*RemoveSpanMembersOutput)(nil), "service.RemoveSpanMembersOutput")
	proto.RegisterType((*ModifySpanAttributesInput)(nil), "service.ModifySpanAttributesInput")
	proto.RegisterType((*ModifySpanAttributesOutput)(nil), "service.ModifySpanAttributesOutput")
	proto.RegisterType((*UpdateSpanInput)(nil), "service.UpdateSpanInput")
	proto.RegisterType((*UpdateSpanOutput)(nil), "service.UpdateSpanOutput")
}

func init() { proto.RegisterFile("span.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0x24, 0xb6, 0x13, 0x97, 0x81, 0xcd, 0xf6, 0x86, 0xdd, 0xc9, 0xec, 0xae, 0x62, 0x0d,
	0x42, 0x32, 0x97, 0x31, 0x8a, 0x90, 0x90, 0xe0, 0x64, 0x25, 0x12, 0xf2, 0x42, 0x16, 0xc9, 0x59,
	0xb8, 0x0e, 0xed, 0xe9, 0xb2, 0xb7, 0x17, 0xcf, 0xf4, 0x30, 0xdd, 0x8e, 0xd6, 0x48, 0x3c, 0x00,
	0x2f, 0xc0, 0x85, 0x1b, 0x8f, 0xc1, 0x81, 0xb7, 0x40, 0xe2, 0x09, 0x78, 0x07, 0x6e, 0xa8, 0x7f,
	0x26, 0x9e, 0xf1, 0x8e, 0x13, 0x7e, 0x36, 0x7b, 0x9b, 0xaa, 0xae, 0xae, 0xfa, 0xaa, 0xea, 0xab,
	0xea, 0x01, 0x90, 0x39, 0xcd, 0xa2, 0xbc, 0x10, 0x4a, 0x90, 0x3d, 0x89, 0xc5, 0x25, 0x4f, 0x30,
	0xf0, 0x65, 0x8e, 0x49, 0x9c, 0xa2, 0xa2, 0x8c, 0x2a, 0x3a, 0xd4, 0x92, 0x35, 0x09, 0x8e, 0xe7,
	0x42, 0xcc, 0x17, 0x38, 0x34, 0xd2, 0x74, 0x39, 0x1b, 0x2a, 0x9e, 0xa2, 0x54, 0x34, 0xcd, 0xad,
	0x41, 0xf8, 0xb3, 0x07, 0x77, 0x4e, 0x0b, 0xa4, 0x0a, 0x2f, 0x72, 0x9a, 0x8d, 0xb3, 0x7c, 0xa9,
	0xc8, 0x43, 0xe8, 0xea, 0x28, 0x71, 0x46, 0x53, 0xf4, 0xbd, 0xbe, 0x37, 0xe8, 0x4e, 0xf6, 0xb5,
	0xe2, 0x29, 0x4d, 0x91, 0x10, 0x68, 0xcd, 0x16, 0x74, 0xee, 0xef, 0xf4, 0xbd, 0x41, 0x7b, 0x62,
	0xbe, 0xc9, 0x03, 0xd8, 0xe3, 0x79, 0x4c, 0x19, 0x2b, 0xfc, 0x5d, 0x63, 0xde, 0xe1, 0xf9, 0x88,
	0xb1, 0x82, 0x1c, 0x43, 0x4f, 0x2d, 0xb3, 0x0c, 0x17, 0xb1, 0x5a, 0xe5, 0xe8, 0xb7, 0xcc, 0x21,
	0x58, 0xd5, 0xb3, 0x55, 0x8e, 0xe4, 0x31, 0x38, 0x29, 0xfe, 0x16, 0x57, 0x7e, 0xdb, 0xf8, 0xec,
	0x5a, 0xcd, 0xe7, 0xb8, 0x0a, 0x5f, 0xc2, 0xc1, 0x1a, 0xdc, 0x97, 0x4b, 0xa5, 0xd1, 0xdd, 0x87,
	0x0e, 0x4d, 0x14, 0x17, 0x99, 0x83, 0xe6, 0x24, 0x72, 0x04, 0xfb, 0x05, 0xaa, 0x38, 0x11, 0x0c,
	0x1d, 0xb8, 0xbd, 0x02, 0xd5, 0xa9, 0x60, 0x48, 0x7c, 0xd8, 0x4b, 0x51, 0x4a, 0x3a, 0x47, 0x87,
	0xaf, 0x14, 0x35, 0x72, 0x93, 0x2a, 0x67, 0x0e, 0x5c, 0x47, 0x8b, 0x63, 0x16, 0xfe, 0xe2, 0x01,
	0x39, 0x43, 0x99, 0x14, 0x7c, 0x6a, 0x82, 0x4b, 0x5b, 0x9a, 0x43, 0x68, 0x6b, 0x03, 0xe9, 0x7b,
	0xfd, 0xdd, 0x41, 0x77, 0x62, 0x85, 0x7a, 0xc1, 0x76, 0x36, 0x0a, 0xb6, 0xb5, 0x38, 0x04, 0x5a,
	0x8a, 0xce, 0xa5, 0xdf, 0x32, 0xae, 0xcc, 0xb7, 0x4e, 0x4e, 0xcc, 0x66, 0x12, 0x95, 0xab, 0x85,
	0x93, 0x74, 0xdc, 0x05, 0x4f, 0xb9, 0xf2, 0x3b, 0x46, 0x6d, 0x85, 0xf0, 0xf7, 0x0e, 0xdc, 0xab,
	0x81, 0xbc, 0x8d, 0x12, 0x9d, 0x81, 0xc9, 0x25, 0xd6, 0xa0, 0x34, 0xd4, 0xde, 0xc9, 0x07, 0x91,
	0x23, 0x5e, 0xd4, 0x10, 0x3c, 0x9a, 0xa0, 0xcc, 0x45, 0x26, 0x71, 0xac, 0x30, 0x9d, 0x98, 0xea,
	0x5e, 0xa0, 0x32, 0x4c, 0x10, 0x8a, 0x2e, 0xe2, 0x44, 0x2c, 0xb3, 0x32, 0x3b, 0x30, 0xaa, 0x53,
	0xad, 0x09, 0x7e, 0x6c, 0xc3, 0x5b, 0xd5, 0xab, 0x9a, 0x1a, 0x5c, 0xc6, 0x34, 0xcf, 0x17, 0x1c,
	0x99, 0x49, 0xa4, 0x3d, 0xe9, 0x72, 0x39, 0xb2, 0x8a, 0xff, 0x58, 0xf3, 0x1b, 0x09, 0x59, 0x36,
	0xa5, 0x5d, 0x69, 0x4a, 0x49, 0xf9, 0x4e, 0x85, 0xf2, 0x75, 0xe2, 0xee, 0x6d, 0x10, 0x97, 0x7c,
	0x0a, 0xbd, 0xc4, 0x10, 0x37, 0xd6, 0x03, 0xe7, 0xef, 0xf7, 0xbd, 0x41, 0xef, 0x24, 0x88, 0xec,
	0x34, 0x46, 0xe5, 0x34, 0x46, 0xcf, 0xca, 0x69, 0x9c, 0x80, 0x35, 0xd7, 0x0a, 0x7d, 0x59, 0x2a,
	0xaa, 0x96, 0xd2, 0x5e, 0xee, 0xde, 0x7c, 0xd9, 0x9a, 0x9b, 0xcb, 0x15, 0x46, 0x43, 0x95, 0xd1,
	0xe4, 0xa9, 0xee, 0x70, 0x3a, 0xc5, 0x42, 0xfa, 0x3d, 0xd3, 0xc6, 0x8f, 0xfe, 0x71, 0x1b, 0xa3,
	0x73, 0x73, 0xd1, 0x76, 0xd4, 0x39, 0x09, 0xfe, 0xf0, 0x00, 0xd6, 0xfa, 0x6a, 0x5c, 0xaf, 0x16,
	0xf7, 0x18, 0x7a, 0x05, 0x4a, 0xb1, 0x2c, 0x12, 0xd4, 0x87, 0xb6, 0x55, 0x50, 0xaa, 0xc6, 0x4c,
	0xb3, 0xd5, 0xe2, 0x2f, 0x7b, 0x65, 0xa5, 0xcd, 0x1a, 0xb6, 0xfe, 0x4f, 0x0d, 0xdb, 0xff, 0xa6,
	0x86, 0xe1, 0x00, 0x0e, 0xce, 0x70, 0x81, 0xea, 0xc6, 0xc9, 0x0f, 0xbf, 0x81, 0xbb, 0x15, 0xcb,
	0x5b, 0x18, 0xbf, 0xf0, 0x33, 0xb8, 0x37, 0x62, 0x4c, 0xbb, 0xb7, 0xc5, 0x76, 0x70, 0x08, 0xb4,
	0x34, 0x02, 0x17, 0xc1, 0x7c, 0x93, 0x47, 0xd0, 0x2d, 0xcb, 0x2a, 0xfd, 0x1d, 0x03, 0x73, 0xad,
	0x08, 0xbf, 0x87, 0xc3, 0xba, 0xa3, 0xdb, 0x58, 0x16, 0xef, 0x42, 0xe7, 0x85, 0x98, 0xae, 0xd7,
	0x69, 0xfb, 0x85, 0x98, 0x8e, 0x59, 0xf8, 0x04, 0xee, 0x4f, 0x30, 0x15, 0x97, 0xf8, 0x1a, 0xf2,
	0xf8, 0x01, 0x1e, 0xbc, 0xe2, 0xeb, 0x0d, 0xa6, 0xf2, 0x9b, 0x07, 0x47, 0xe7, 0x82, 0xf1, 0xd9,
	0x4a, 0xc7, 0x1f, 0x29, 0x55, 0xf0, 0xe9, 0x52, 0xa1, 0x4b, 0x67, 0xeb, 0x14, 0x5c, 0xbb, 0xae,
	0xca, 0x05, 0xb3, 0xdb, 0xfc, 0xa6, 0xb6, 0xae, 0x5b, 0x61, 0xed, 0x1b, 0xde, 0xd4, 0xce, 0xe6,
	0x9b, 0xca, 0x21, 0x68, 0xc2, 0x7f, 0x1b, 0xdc, 0x7d, 0x1f, 0xee, 0x7c, 0x95, 0xb3, 0xda, 0xbf,
	0x45, 0x43, 0xbf, 0xc3, 0x4b, 0x38, 0x58, 0x9b, 0xbd, 0xb9, 0x56, 0x9e, 0xfc, 0xd9, 0x82, 0xde,
	0x85, 0x79, 0x9f, 0xcc, 0x1a, 0x24, 0x23, 0x80, 0xf5, 0xdf, 0x06, 0xf1, 0xaf, 0xd6, 0xe3, 0xc6,
	0xff, 0x51, 0x70, 0xd4, 0x70, 0xe2, 0x60, 0x3f, 0x81, 0xb7, 0x6b, 0xcb, 0x94, 0x3c, 0x6c, 0x5e,
	0xb2, 0xd6, 0xd1, 0xa3, 0xeb, 0x36, 0x30, 0x39, 0x83, 0x5e, 0x65, 0xb7, 0x90, 0xa3, 0x8a, 0x71,
	0x7d, 0x37, 0x05, 0x41, 0xd3, 0x91, 0xf3, 0x72, 0x0e, 0xef, 0xd4, 0xc7, 0x9e, 0xac, 0xa3, 0x36,
	0x2c, 0x96, 0xe0, 0xf1, 0x96, 0x53, 0xe7, 0xee, 0x6b, 0xb8, 0xfb, 0xca, 0xf4, 0x91, 0xe3, 0xab,
	0x3b, 0xcd, 0x53, 0x1e, 0xf4, 0xb7, 0x1b, 0x38, 0xbf, 0x31, 0x1c, 0x36, 0xb1, 0x92, 0x84, 0x57,
	0x37, 0xb7, 0x0e, 0x5d, 0xf0, 0xde, 0xb5, 0x36, 0x2e, 0xc0, 0x08, 0x60, 0x4d, 0xb2, 0x4a, 0x73,
	0x37, 0x08, 0x5a, 0x69, 0xee, 0x26, 0x27, 0x83, 0x8f, 0x7f, 0xfd, 0xe9, 0xaf, 0x2f, 0x4e, 0xe0,
	0xc3, 0x20, 0x7a, 0xae, 0x54, 0x2e, 0x3f, 0x19, 0x0e, 0x99, 0x48, 0x64, 0xf4, 0x1d, 0xcf, 0xe6,
	0xc9, 0x42, 0x2c, 0x59, 0x94, 0x88, 0x74, 0x48, 0x73, 0x3e, 0xd4, 0xac, 0x1e, 0xf2, 0x8c, 0xe1,
	0xcb, 0xe8, 0xb9, 0x4a, 0x17, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x1a, 0x98, 0x5c, 0xb5,
	0x0b, 0x00, 0x00,
}
