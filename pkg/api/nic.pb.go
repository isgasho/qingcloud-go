// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nic.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateNicsInput struct {
	Vxnet                *string  `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	NicName              *string  `protobuf:"bytes,2,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Count                *int32   `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	PrivateIps           []string `protobuf:"bytes,4,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNicsInput) Reset()         { *m = CreateNicsInput{} }
func (m *CreateNicsInput) String() string { return proto.CompactTextString(m) }
func (*CreateNicsInput) ProtoMessage()    {}
func (*CreateNicsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{0}
}
func (m *CreateNicsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNicsInput.Unmarshal(m, b)
}
func (m *CreateNicsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNicsInput.Marshal(b, m, deterministic)
}
func (dst *CreateNicsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNicsInput.Merge(dst, src)
}
func (m *CreateNicsInput) XXX_Size() int {
	return xxx_messageInfo_CreateNicsInput.Size(m)
}
func (m *CreateNicsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNicsInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNicsInput proto.InternalMessageInfo

func (m *CreateNicsInput) GetVxnet() string {
	if m != nil && m.Vxnet != nil {
		return *m.Vxnet
	}
	return ""
}

func (m *CreateNicsInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
	}
	return ""
}

func (m *CreateNicsInput) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *CreateNicsInput) GetPrivateIps() []string {
	if m != nil {
		return m.PrivateIps
	}
	return nil
}

type CreateNicsOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Nics                 []*NICIP `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNicsOutput) Reset()         { *m = CreateNicsOutput{} }
func (m *CreateNicsOutput) String() string { return proto.CompactTextString(m) }
func (*CreateNicsOutput) ProtoMessage()    {}
func (*CreateNicsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{1}
}
func (m *CreateNicsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNicsOutput.Unmarshal(m, b)
}
func (m *CreateNicsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNicsOutput.Marshal(b, m, deterministic)
}
func (dst *CreateNicsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNicsOutput.Merge(dst, src)
}
func (m *CreateNicsOutput) XXX_Size() int {
	return xxx_messageInfo_CreateNicsOutput.Size(m)
}
func (m *CreateNicsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNicsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNicsOutput proto.InternalMessageInfo

func (m *CreateNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *CreateNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *CreateNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *CreateNicsOutput) GetNics() []*NICIP {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DescribeNicsInput struct {
	Instances            []string `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	Limit                *int32   `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	NicName              *string  `protobuf:"bytes,3,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	Nics                 []string `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	Offset               *int32   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Status               *string  `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	VxnetType            *int32   `protobuf:"varint,7,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	Vxnets               []string `protobuf:"bytes,8,rep,name=vxnets" json:"vxnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeNicsInput) Reset()         { *m = DescribeNicsInput{} }
func (m *DescribeNicsInput) String() string { return proto.CompactTextString(m) }
func (*DescribeNicsInput) ProtoMessage()    {}
func (*DescribeNicsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{2}
}
func (m *DescribeNicsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeNicsInput.Unmarshal(m, b)
}
func (m *DescribeNicsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeNicsInput.Marshal(b, m, deterministic)
}
func (dst *DescribeNicsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeNicsInput.Merge(dst, src)
}
func (m *DescribeNicsInput) XXX_Size() int {
	return xxx_messageInfo_DescribeNicsInput.Size(m)
}
func (m *DescribeNicsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeNicsInput.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeNicsInput proto.InternalMessageInfo

func (m *DescribeNicsInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *DescribeNicsInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *DescribeNicsInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
	}
	return ""
}

func (m *DescribeNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *DescribeNicsInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeNicsInput) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *DescribeNicsInput) GetVxnetType() int32 {
	if m != nil && m.VxnetType != nil {
		return *m.VxnetType
	}
	return 0
}

func (m *DescribeNicsInput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

type DescribeNicsOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	NicSet               []*NIC   `protobuf:"bytes,4,rep,name=nic_set,json=nicSet" json:"nic_set,omitempty"`
	TotalCount           *int32   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeNicsOutput) Reset()         { *m = DescribeNicsOutput{} }
func (m *DescribeNicsOutput) String() string { return proto.CompactTextString(m) }
func (*DescribeNicsOutput) ProtoMessage()    {}
func (*DescribeNicsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{3}
}
func (m *DescribeNicsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeNicsOutput.Unmarshal(m, b)
}
func (m *DescribeNicsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeNicsOutput.Marshal(b, m, deterministic)
}
func (dst *DescribeNicsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeNicsOutput.Merge(dst, src)
}
func (m *DescribeNicsOutput) XXX_Size() int {
	return xxx_messageInfo_DescribeNicsOutput.Size(m)
}
func (m *DescribeNicsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeNicsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeNicsOutput proto.InternalMessageInfo

func (m *DescribeNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeNicsOutput) GetNicSet() []*NIC {
	if m != nil {
		return m.NicSet
	}
	return nil
}

func (m *DescribeNicsOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type AttachNicsInput struct {
	Nics                 []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	Instance             *string  `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachNicsInput) Reset()         { *m = AttachNicsInput{} }
func (m *AttachNicsInput) String() string { return proto.CompactTextString(m) }
func (*AttachNicsInput) ProtoMessage()    {}
func (*AttachNicsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{4}
}
func (m *AttachNicsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachNicsInput.Unmarshal(m, b)
}
func (m *AttachNicsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachNicsInput.Marshal(b, m, deterministic)
}
func (dst *AttachNicsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachNicsInput.Merge(dst, src)
}
func (m *AttachNicsInput) XXX_Size() int {
	return xxx_messageInfo_AttachNicsInput.Size(m)
}
func (m *AttachNicsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachNicsInput.DiscardUnknown(m)
}

var xxx_messageInfo_AttachNicsInput proto.InternalMessageInfo

func (m *AttachNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *AttachNicsInput) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

type AttachNicsOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId                *string  `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachNicsOutput) Reset()         { *m = AttachNicsOutput{} }
func (m *AttachNicsOutput) String() string { return proto.CompactTextString(m) }
func (*AttachNicsOutput) ProtoMessage()    {}
func (*AttachNicsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{5}
}
func (m *AttachNicsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachNicsOutput.Unmarshal(m, b)
}
func (m *AttachNicsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachNicsOutput.Marshal(b, m, deterministic)
}
func (dst *AttachNicsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachNicsOutput.Merge(dst, src)
}
func (m *AttachNicsOutput) XXX_Size() int {
	return xxx_messageInfo_AttachNicsOutput.Size(m)
}
func (m *AttachNicsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachNicsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_AttachNicsOutput proto.InternalMessageInfo

func (m *AttachNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AttachNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AttachNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AttachNicsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type DetachNicsInput struct {
	Nics                 []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetachNicsInput) Reset()         { *m = DetachNicsInput{} }
func (m *DetachNicsInput) String() string { return proto.CompactTextString(m) }
func (*DetachNicsInput) ProtoMessage()    {}
func (*DetachNicsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{6}
}
func (m *DetachNicsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetachNicsInput.Unmarshal(m, b)
}
func (m *DetachNicsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetachNicsInput.Marshal(b, m, deterministic)
}
func (dst *DetachNicsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetachNicsInput.Merge(dst, src)
}
func (m *DetachNicsInput) XXX_Size() int {
	return xxx_messageInfo_DetachNicsInput.Size(m)
}
func (m *DetachNicsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DetachNicsInput.DiscardUnknown(m)
}

var xxx_messageInfo_DetachNicsInput proto.InternalMessageInfo

func (m *DetachNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DetachNicsOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId                *string  `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetachNicsOutput) Reset()         { *m = DetachNicsOutput{} }
func (m *DetachNicsOutput) String() string { return proto.CompactTextString(m) }
func (*DetachNicsOutput) ProtoMessage()    {}
func (*DetachNicsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{7}
}
func (m *DetachNicsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetachNicsOutput.Unmarshal(m, b)
}
func (m *DetachNicsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetachNicsOutput.Marshal(b, m, deterministic)
}
func (dst *DetachNicsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetachNicsOutput.Merge(dst, src)
}
func (m *DetachNicsOutput) XXX_Size() int {
	return xxx_messageInfo_DetachNicsOutput.Size(m)
}
func (m *DetachNicsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DetachNicsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DetachNicsOutput proto.InternalMessageInfo

func (m *DetachNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DetachNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DetachNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DetachNicsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ModifyNicAttributesInput struct {
	Nic                  *string  `protobuf:"bytes,1,opt,name=nic" json:"nic,omitempty"`
	NicName              *string  `protobuf:"bytes,2,opt,name=nic_name,json=nicName" json:"nic_name,omitempty"`
	PrivateIp            *string  `protobuf:"bytes,3,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyNicAttributesInput) Reset()         { *m = ModifyNicAttributesInput{} }
func (m *ModifyNicAttributesInput) String() string { return proto.CompactTextString(m) }
func (*ModifyNicAttributesInput) ProtoMessage()    {}
func (*ModifyNicAttributesInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{8}
}
func (m *ModifyNicAttributesInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyNicAttributesInput.Unmarshal(m, b)
}
func (m *ModifyNicAttributesInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyNicAttributesInput.Marshal(b, m, deterministic)
}
func (dst *ModifyNicAttributesInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyNicAttributesInput.Merge(dst, src)
}
func (m *ModifyNicAttributesInput) XXX_Size() int {
	return xxx_messageInfo_ModifyNicAttributesInput.Size(m)
}
func (m *ModifyNicAttributesInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyNicAttributesInput.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyNicAttributesInput proto.InternalMessageInfo

func (m *ModifyNicAttributesInput) GetNic() string {
	if m != nil && m.Nic != nil {
		return *m.Nic
	}
	return ""
}

func (m *ModifyNicAttributesInput) GetNicName() string {
	if m != nil && m.NicName != nil {
		return *m.NicName
	}
	return ""
}

func (m *ModifyNicAttributesInput) GetPrivateIp() string {
	if m != nil && m.PrivateIp != nil {
		return *m.PrivateIp
	}
	return ""
}

type ModifyNicAttributesOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyNicAttributesOutput) Reset()         { *m = ModifyNicAttributesOutput{} }
func (m *ModifyNicAttributesOutput) String() string { return proto.CompactTextString(m) }
func (*ModifyNicAttributesOutput) ProtoMessage()    {}
func (*ModifyNicAttributesOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{9}
}
func (m *ModifyNicAttributesOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyNicAttributesOutput.Unmarshal(m, b)
}
func (m *ModifyNicAttributesOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyNicAttributesOutput.Marshal(b, m, deterministic)
}
func (dst *ModifyNicAttributesOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyNicAttributesOutput.Merge(dst, src)
}
func (m *ModifyNicAttributesOutput) XXX_Size() int {
	return xxx_messageInfo_ModifyNicAttributesOutput.Size(m)
}
func (m *ModifyNicAttributesOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyNicAttributesOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyNicAttributesOutput proto.InternalMessageInfo

func (m *ModifyNicAttributesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ModifyNicAttributesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ModifyNicAttributesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type DeleteNicsInput struct {
	Nics                 []string `protobuf:"bytes,1,rep,name=nics" json:"nics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNicsInput) Reset()         { *m = DeleteNicsInput{} }
func (m *DeleteNicsInput) String() string { return proto.CompactTextString(m) }
func (*DeleteNicsInput) ProtoMessage()    {}
func (*DeleteNicsInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{10}
}
func (m *DeleteNicsInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNicsInput.Unmarshal(m, b)
}
func (m *DeleteNicsInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNicsInput.Marshal(b, m, deterministic)
}
func (dst *DeleteNicsInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNicsInput.Merge(dst, src)
}
func (m *DeleteNicsInput) XXX_Size() int {
	return xxx_messageInfo_DeleteNicsInput.Size(m)
}
func (m *DeleteNicsInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNicsInput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNicsInput proto.InternalMessageInfo

func (m *DeleteNicsInput) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

type DeleteNicsOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNicsOutput) Reset()         { *m = DeleteNicsOutput{} }
func (m *DeleteNicsOutput) String() string { return proto.CompactTextString(m) }
func (*DeleteNicsOutput) ProtoMessage()    {}
func (*DeleteNicsOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_nic_c508102ef86e4d6a, []int{11}
}
func (m *DeleteNicsOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNicsOutput.Unmarshal(m, b)
}
func (m *DeleteNicsOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNicsOutput.Marshal(b, m, deterministic)
}
func (dst *DeleteNicsOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNicsOutput.Merge(dst, src)
}
func (m *DeleteNicsOutput) XXX_Size() int {
	return xxx_messageInfo_DeleteNicsOutput.Size(m)
}
func (m *DeleteNicsOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNicsOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNicsOutput proto.InternalMessageInfo

func (m *DeleteNicsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteNicsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteNicsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateNicsInput)(nil), "service.CreateNicsInput")
	proto.RegisterType((*CreateNicsOutput)(nil), "service.CreateNicsOutput")
	proto.RegisterType((*DescribeNicsInput)(nil), "service.DescribeNicsInput")
	proto.RegisterType((*DescribeNicsOutput)(nil), "service.DescribeNicsOutput")
	proto.RegisterType((*AttachNicsInput)(nil), "service.AttachNicsInput")
	proto.RegisterType((*AttachNicsOutput)(nil), "service.AttachNicsOutput")
	proto.RegisterType((*DetachNicsInput)(nil), "service.DetachNicsInput")
	proto.RegisterType((*DetachNicsOutput)(nil), "service.DetachNicsOutput")
	proto.RegisterType((*ModifyNicAttributesInput)(nil), "service.ModifyNicAttributesInput")
	proto.RegisterType((*ModifyNicAttributesOutput)(nil), "service.ModifyNicAttributesOutput")
	proto.RegisterType((*DeleteNicsInput)(nil), "service.DeleteNicsInput")
	proto.RegisterType((*DeleteNicsOutput)(nil), "service.DeleteNicsOutput")
}

func init() { proto.RegisterFile("nic.proto", fileDescriptor_nic_c508102ef86e4d6a) }

var fileDescriptor_nic_c508102ef86e4d6a = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xbf, 0x34, 0x49, 0x7d, 0x5b, 0x7d, 0x2d, 0xe6, 0x47, 0x53, 0x43, 0x45, 0xb1, 0x54,
	0xa9, 0x1b, 0x12, 0xd1, 0x05, 0x0b, 0x76, 0x51, 0x2a, 0xa1, 0x48, 0x10, 0x90, 0x61, 0x89, 0x64,
	0x4d, 0xc6, 0x93, 0x66, 0xaa, 0x78, 0xc6, 0x78, 0x6e, 0xa2, 0x66, 0xcd, 0x3b, 0xf0, 0x0e, 0xbc,
	0x02, 0xef, 0xc2, 0x73, 0x20, 0x76, 0x68, 0x3c, 0x93, 0xd8, 0x69, 0x43, 0x59, 0x85, 0x9d, 0xcf,
	0x1d, 0xcf, 0x9c, 0x73, 0xcf, 0x9c, 0x3b, 0xe0, 0x4b, 0xc1, 0x3a, 0x79, 0xa1, 0x50, 0x05, 0x6d,
	0xcd, 0x8b, 0xb9, 0x60, 0x3c, 0xdc, 0xc3, 0x45, 0xce, 0xb5, 0xad, 0x86, 0x44, 0xe7, 0x9c, 0x25,
	0x19, 0x47, 0x9a, 0x52, 0xa4, 0x5d, 0x83, 0xec, 0x4a, 0xb4, 0x80, 0x83, 0x7e, 0xc1, 0x29, 0xf2,
	0xa1, 0x60, 0x7a, 0x20, 0xf3, 0x19, 0x06, 0x0f, 0xa0, 0x39, 0xbf, 0x96, 0x1c, 0x89, 0x77, 0xe2,
	0x9d, 0xf9, 0xb1, 0x05, 0xc1, 0x11, 0xec, 0x4a, 0xc1, 0x12, 0x49, 0x33, 0x4e, 0xfe, 0x2b, 0x17,
	0xda, 0x52, 0xb0, 0x21, 0xcd, 0xb8, 0xd9, 0xc0, 0xd4, 0x4c, 0x22, 0x69, 0x9c, 0x78, 0x67, 0xcd,
	0xd8, 0x82, 0xe0, 0x29, 0xec, 0xe5, 0x85, 0x98, 0x53, 0xe4, 0x89, 0xc8, 0x35, 0xd9, 0x39, 0x69,
	0x9c, 0xf9, 0x31, 0xb8, 0xd2, 0x20, 0xd7, 0xd1, 0x17, 0x0f, 0x0e, 0x2b, 0xee, 0x77, 0x33, 0x34,
	0xe4, 0x8f, 0xa0, 0x45, 0x19, 0x0a, 0x25, 0x1d, 0xbb, 0x43, 0x86, 0xbe, 0xe0, 0x98, 0x30, 0x95,
	0x5a, 0xfa, 0x66, 0xdc, 0x2e, 0x38, 0xf6, 0x55, 0xca, 0x03, 0x02, 0xed, 0x8c, 0x6b, 0x4d, 0x2f,
	0x79, 0x29, 0xc0, 0x8f, 0x97, 0x30, 0x88, 0x60, 0x47, 0x0a, 0x66, 0xb9, 0xf7, 0xce, 0xff, 0xef,
	0x38, 0x6f, 0x3a, 0xc3, 0x41, 0x7f, 0xf0, 0x3e, 0x2e, 0xd7, 0xa2, 0x1f, 0x1e, 0xdc, 0xbb, 0xe0,
	0x9a, 0x15, 0x62, 0x54, 0xf3, 0xe0, 0x09, 0xf8, 0x42, 0x6a, 0xa4, 0x92, 0x71, 0x4d, 0xbc, 0x52,
	0x7a, 0x55, 0x30, 0x0d, 0x4f, 0x45, 0x26, 0xd0, 0x29, 0xb1, 0x60, 0xcd, 0xa1, 0xc6, 0xba, 0x43,
	0x41, 0x4d, 0x88, 0x6f, 0x89, 0x4d, 0xa7, 0x6a, 0x3c, 0xd6, 0x1c, 0x49, 0xb3, 0x3c, 0xc5, 0x21,
	0x53, 0xd7, 0x48, 0x71, 0xa6, 0x49, 0xcb, 0x3a, 0x60, 0x51, 0x70, 0x0c, 0x50, 0xde, 0x44, 0x62,
	0x2e, 0x96, 0xb4, 0xcb, 0x3d, 0x7e, 0x59, 0xf9, 0xb8, 0xc8, 0xb9, 0xd9, 0x56, 0x02, 0x4d, 0x76,
	0x4b, 0x12, 0x87, 0xa2, 0x6f, 0x1e, 0x04, 0xf5, 0xfe, 0xb6, 0xe1, 0xf3, 0x29, 0x98, 0x4e, 0x13,
	0xd3, 0x8b, 0xb5, 0x7a, 0xbf, 0x6e, 0x75, 0xdc, 0x92, 0x82, 0x7d, 0xe0, 0x65, 0x22, 0x50, 0x21,
	0x9d, 0x26, 0x36, 0x2d, 0xb6, 0x6d, 0x28, 0x4b, 0x7d, 0x53, 0x89, 0x7a, 0x70, 0xd0, 0x43, 0xa4,
	0x6c, 0x52, 0x5d, 0xc4, 0xd2, 0x39, 0xaf, 0xe6, 0x5c, 0x08, 0xbb, 0xcb, 0xbb, 0x70, 0x51, 0x5c,
	0xe1, 0x68, 0x0e, 0x87, 0xd5, 0x11, 0xdb, 0xe8, 0xf5, 0x21, 0xb4, 0xae, 0xd4, 0x28, 0x11, 0x29,
	0xd9, 0xb1, 0xe3, 0x71, 0xa5, 0x46, 0x83, 0x34, 0x3a, 0x85, 0x83, 0x0b, 0xfe, 0x57, 0xe9, 0x46,
	0x5e, 0xf5, 0xdb, 0x3f, 0x94, 0x37, 0x06, 0xf2, 0x56, 0xa5, 0x62, 0xbc, 0x18, 0x0a, 0xd6, 0x43,
	0x2c, 0xc4, 0x68, 0x86, 0xdc, 0xe9, 0x3c, 0x84, 0x86, 0x14, 0xcc, 0x91, 0x9b, 0xcf, 0xbb, 0x66,
	0xfd, 0x18, 0xa0, 0x9a, 0x6a, 0x47, 0xee, 0xaf, 0x86, 0x3a, 0x9a, 0xc0, 0xd1, 0x06, 0x9e, 0x2d,
	0x34, 0x6a, 0x0d, 0x9f, 0xf2, 0xfa, 0xc3, 0xb5, 0xc9, 0xf0, 0xc4, 0x18, 0xbe, 0xfc, 0x6d, 0x0b,
	0x3a, 0xce, 0x7f, 0x36, 0x00, 0x86, 0x26, 0xdf, 0x65, 0xde, 0x83, 0x1e, 0x40, 0xf5, 0xa6, 0x05,
	0x64, 0x35, 0x07, 0x37, 0x1e, 0xd9, 0xf0, 0x68, 0xc3, 0x8a, 0x93, 0xf7, 0x1a, 0xf6, 0xeb, 0x03,
	0x1b, 0x84, 0xab, 0x5f, 0x6f, 0xbd, 0x53, 0xe1, 0xe3, 0x8d, 0x6b, 0xee, 0xa0, 0x1e, 0x40, 0x35,
	0x0b, 0x35, 0x2d, 0x37, 0x66, 0xac, 0xa6, 0xe5, 0xd6, 0xe8, 0xf4, 0x00, 0xaa, 0xbc, 0xd6, 0x8e,
	0xb8, 0x91, 0xf5, 0xda, 0x11, 0xb7, 0xe2, 0xfd, 0x09, 0xee, 0x6f, 0x88, 0x44, 0xf0, 0x6c, 0xb5,
	0xe3, 0x4f, 0xc1, 0x0c, 0xa3, 0xbb, 0x7e, 0xa9, 0x0b, 0x5c, 0xde, 0xef, 0x9a, 0xc0, 0xb5, 0x6c,
	0xac, 0x09, 0x5c, 0x8f, 0x43, 0xf8, 0xf2, 0xfb, 0xd7, 0x5f, 0x6f, 0x5e, 0x40, 0x37, 0x7c, 0x3e,
	0x41, 0xcc, 0xf5, 0xab, 0x6e, 0x37, 0x55, 0x4c, 0x77, 0x3e, 0x0b, 0x79, 0xc9, 0xa6, 0x6a, 0x96,
	0x76, 0x98, 0xca, 0xba, 0x34, 0x17, 0x5d, 0x29, 0x58, 0x57, 0xc8, 0x94, 0x5f, 0x77, 0x26, 0x98,
	0x4d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xee, 0xfe, 0x3f, 0x08, 0x76, 0x07, 0x00, 0x00,
}
