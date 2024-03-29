// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetMonitorInput struct {
	Resource             *string              `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Meters               []string             `protobuf:"bytes,2,rep,name=meters" json:"meters,omitempty"`
	Step                 *string              `protobuf:"bytes,3,opt,name=step" json:"step,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetMonitorInput) Reset()         { *m = GetMonitorInput{} }
func (m *GetMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetMonitorInput) ProtoMessage()    {}
func (*GetMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{0}
}
func (m *GetMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMonitorInput.Unmarshal(m, b)
}
func (m *GetMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMonitorInput.Merge(dst, src)
}
func (m *GetMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetMonitorInput.Size(m)
}
func (m *GetMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetMonitorInput proto.InternalMessageInfo

func (m *GetMonitorInput) GetResource() string {
	if m != nil && m.Resource != nil {
		return *m.Resource
	}
	return ""
}

func (m *GetMonitorInput) GetMeters() []string {
	if m != nil {
		return m.Meters
	}
	return nil
}

func (m *GetMonitorInput) GetStep() string {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return ""
}

func (m *GetMonitorInput) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetMonitorInput) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetMonitorOutput struct {
	Action               *string                          `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32                           `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string                          `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ResourceId           *string                          `protobuf:"bytes,4,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	MeterSet             []*GetMonitorOutput_ResponseItem `protobuf:"bytes,5,rep,name=meter_set,json=meterSet" json:"meter_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetMonitorOutput) Reset()         { *m = GetMonitorOutput{} }
func (m *GetMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetMonitorOutput) ProtoMessage()    {}
func (*GetMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{1}
}
func (m *GetMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMonitorOutput.Unmarshal(m, b)
}
func (m *GetMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMonitorOutput.Merge(dst, src)
}
func (m *GetMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetMonitorOutput.Size(m)
}
func (m *GetMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetMonitorOutput proto.InternalMessageInfo

func (m *GetMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GetMonitorOutput) GetResourceId() string {
	if m != nil && m.ResourceId != nil {
		return *m.ResourceId
	}
	return ""
}

func (m *GetMonitorOutput) GetMeterSet() []*GetMonitorOutput_ResponseItem {
	if m != nil {
		return m.MeterSet
	}
	return nil
}

type GetMonitorOutput_ResponseItem struct {
	// data: [[1392072000,[12,12]],[15,29],[11,12]]
	VxnetId              *string  `protobuf:"bytes,2,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	MeterId              *string  `protobuf:"bytes,3,opt,name=meter_id,json=meterId" json:"meter_id,omitempty"`
	Sequence             *int32   `protobuf:"varint,4,opt,name=sequence" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMonitorOutput_ResponseItem) Reset()         { *m = GetMonitorOutput_ResponseItem{} }
func (m *GetMonitorOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*GetMonitorOutput_ResponseItem) ProtoMessage()    {}
func (*GetMonitorOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{1, 0}
}
func (m *GetMonitorOutput_ResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMonitorOutput_ResponseItem.Unmarshal(m, b)
}
func (m *GetMonitorOutput_ResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMonitorOutput_ResponseItem.Marshal(b, m, deterministic)
}
func (dst *GetMonitorOutput_ResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMonitorOutput_ResponseItem.Merge(dst, src)
}
func (m *GetMonitorOutput_ResponseItem) XXX_Size() int {
	return xxx_messageInfo_GetMonitorOutput_ResponseItem.Size(m)
}
func (m *GetMonitorOutput_ResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMonitorOutput_ResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_GetMonitorOutput_ResponseItem proto.InternalMessageInfo

func (m *GetMonitorOutput_ResponseItem) GetVxnetId() string {
	if m != nil && m.VxnetId != nil {
		return *m.VxnetId
	}
	return ""
}

func (m *GetMonitorOutput_ResponseItem) GetMeterId() string {
	if m != nil && m.MeterId != nil {
		return *m.MeterId
	}
	return ""
}

func (m *GetMonitorOutput_ResponseItem) GetSequence() int32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return 0
}

type GetLoadBalancerMonitorInput struct {
	Resource             *string              `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Meters               []string             `protobuf:"bytes,2,rep,name=meters" json:"meters,omitempty"`
	Step                 *string              `protobuf:"bytes,3,opt,name=step" json:"step,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetLoadBalancerMonitorInput) Reset()         { *m = GetLoadBalancerMonitorInput{} }
func (m *GetLoadBalancerMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorInput) ProtoMessage()    {}
func (*GetLoadBalancerMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{2}
}
func (m *GetLoadBalancerMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLoadBalancerMonitorInput.Unmarshal(m, b)
}
func (m *GetLoadBalancerMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLoadBalancerMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetLoadBalancerMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLoadBalancerMonitorInput.Merge(dst, src)
}
func (m *GetLoadBalancerMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetLoadBalancerMonitorInput.Size(m)
}
func (m *GetLoadBalancerMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLoadBalancerMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetLoadBalancerMonitorInput proto.InternalMessageInfo

func (m *GetLoadBalancerMonitorInput) GetResource() string {
	if m != nil && m.Resource != nil {
		return *m.Resource
	}
	return ""
}

func (m *GetLoadBalancerMonitorInput) GetMeters() []string {
	if m != nil {
		return m.Meters
	}
	return nil
}

func (m *GetLoadBalancerMonitorInput) GetStep() string {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return ""
}

func (m *GetLoadBalancerMonitorInput) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetLoadBalancerMonitorInput) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetLoadBalancerMonitorOutput struct {
	Action               *string                                      `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32                                       `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string                                      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ResourceId           *string                                      `protobuf:"bytes,4,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	MeterSet             []*GetLoadBalancerMonitorOutput_ResponseItem `protobuf:"bytes,5,rep,name=meter_set,json=meterSet" json:"meter_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *GetLoadBalancerMonitorOutput) Reset()         { *m = GetLoadBalancerMonitorOutput{} }
func (m *GetLoadBalancerMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorOutput) ProtoMessage()    {}
func (*GetLoadBalancerMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{3}
}
func (m *GetLoadBalancerMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput.Unmarshal(m, b)
}
func (m *GetLoadBalancerMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetLoadBalancerMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLoadBalancerMonitorOutput.Merge(dst, src)
}
func (m *GetLoadBalancerMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput.Size(m)
}
func (m *GetLoadBalancerMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLoadBalancerMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetLoadBalancerMonitorOutput proto.InternalMessageInfo

func (m *GetLoadBalancerMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetLoadBalancerMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput) GetResourceId() string {
	if m != nil && m.ResourceId != nil {
		return *m.ResourceId
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput) GetMeterSet() []*GetLoadBalancerMonitorOutput_ResponseItem {
	if m != nil {
		return m.MeterSet
	}
	return nil
}

type GetLoadBalancerMonitorOutput_ResponseItem struct {
	EipId *string `protobuf:"bytes,1,opt,name=eip_id,json=eipId" json:"eip_id,omitempty"`
	// data: ...
	MeterId              *string  `protobuf:"bytes,3,opt,name=meter_id,json=meterId" json:"meter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLoadBalancerMonitorOutput_ResponseItem) Reset() {
	*m = GetLoadBalancerMonitorOutput_ResponseItem{}
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorOutput_ResponseItem) ProtoMessage()    {}
func (*GetLoadBalancerMonitorOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{3, 0}
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem.Unmarshal(m, b)
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem.Marshal(b, m, deterministic)
}
func (dst *GetLoadBalancerMonitorOutput_ResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem.Merge(dst, src)
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) XXX_Size() int {
	return xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem.Size(m)
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_GetLoadBalancerMonitorOutput_ResponseItem proto.InternalMessageInfo

func (m *GetLoadBalancerMonitorOutput_ResponseItem) GetEipId() string {
	if m != nil && m.EipId != nil {
		return *m.EipId
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput_ResponseItem) GetMeterId() string {
	if m != nil && m.MeterId != nil {
		return *m.MeterId
	}
	return ""
}

type GetRDBMonitorInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRDBMonitorInput) Reset()         { *m = GetRDBMonitorInput{} }
func (m *GetRDBMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetRDBMonitorInput) ProtoMessage()    {}
func (*GetRDBMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{4}
}
func (m *GetRDBMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRDBMonitorInput.Unmarshal(m, b)
}
func (m *GetRDBMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRDBMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetRDBMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRDBMonitorInput.Merge(dst, src)
}
func (m *GetRDBMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetRDBMonitorInput.Size(m)
}
func (m *GetRDBMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRDBMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetRDBMonitorInput proto.InternalMessageInfo

type GetRDBMonitorOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRDBMonitorOutput) Reset()         { *m = GetRDBMonitorOutput{} }
func (m *GetRDBMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetRDBMonitorOutput) ProtoMessage()    {}
func (*GetRDBMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{5}
}
func (m *GetRDBMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRDBMonitorOutput.Unmarshal(m, b)
}
func (m *GetRDBMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRDBMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetRDBMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRDBMonitorOutput.Merge(dst, src)
}
func (m *GetRDBMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetRDBMonitorOutput.Size(m)
}
func (m *GetRDBMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRDBMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetRDBMonitorOutput proto.InternalMessageInfo

func (m *GetRDBMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetRDBMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetRDBMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type GetCacheMonitorInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCacheMonitorInput) Reset()         { *m = GetCacheMonitorInput{} }
func (m *GetCacheMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetCacheMonitorInput) ProtoMessage()    {}
func (*GetCacheMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{6}
}
func (m *GetCacheMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCacheMonitorInput.Unmarshal(m, b)
}
func (m *GetCacheMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCacheMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetCacheMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheMonitorInput.Merge(dst, src)
}
func (m *GetCacheMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetCacheMonitorInput.Size(m)
}
func (m *GetCacheMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheMonitorInput proto.InternalMessageInfo

type GetCacheMonitorOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCacheMonitorOutput) Reset()         { *m = GetCacheMonitorOutput{} }
func (m *GetCacheMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetCacheMonitorOutput) ProtoMessage()    {}
func (*GetCacheMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{7}
}
func (m *GetCacheMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCacheMonitorOutput.Unmarshal(m, b)
}
func (m *GetCacheMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCacheMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetCacheMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCacheMonitorOutput.Merge(dst, src)
}
func (m *GetCacheMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetCacheMonitorOutput.Size(m)
}
func (m *GetCacheMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCacheMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetCacheMonitorOutput proto.InternalMessageInfo

func (m *GetCacheMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetCacheMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetCacheMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type GetZooKeeperMonitorInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetZooKeeperMonitorInput) Reset()         { *m = GetZooKeeperMonitorInput{} }
func (m *GetZooKeeperMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetZooKeeperMonitorInput) ProtoMessage()    {}
func (*GetZooKeeperMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{8}
}
func (m *GetZooKeeperMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetZooKeeperMonitorInput.Unmarshal(m, b)
}
func (m *GetZooKeeperMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetZooKeeperMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetZooKeeperMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetZooKeeperMonitorInput.Merge(dst, src)
}
func (m *GetZooKeeperMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetZooKeeperMonitorInput.Size(m)
}
func (m *GetZooKeeperMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetZooKeeperMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetZooKeeperMonitorInput proto.InternalMessageInfo

type GetZooKeeperMonitorOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetZooKeeperMonitorOutput) Reset()         { *m = GetZooKeeperMonitorOutput{} }
func (m *GetZooKeeperMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetZooKeeperMonitorOutput) ProtoMessage()    {}
func (*GetZooKeeperMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{9}
}
func (m *GetZooKeeperMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetZooKeeperMonitorOutput.Unmarshal(m, b)
}
func (m *GetZooKeeperMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetZooKeeperMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetZooKeeperMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetZooKeeperMonitorOutput.Merge(dst, src)
}
func (m *GetZooKeeperMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetZooKeeperMonitorOutput.Size(m)
}
func (m *GetZooKeeperMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetZooKeeperMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetZooKeeperMonitorOutput proto.InternalMessageInfo

func (m *GetZooKeeperMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetZooKeeperMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetZooKeeperMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type GetQueueMonitorInput struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQueueMonitorInput) Reset()         { *m = GetQueueMonitorInput{} }
func (m *GetQueueMonitorInput) String() string { return proto.CompactTextString(m) }
func (*GetQueueMonitorInput) ProtoMessage()    {}
func (*GetQueueMonitorInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{10}
}
func (m *GetQueueMonitorInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueueMonitorInput.Unmarshal(m, b)
}
func (m *GetQueueMonitorInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueueMonitorInput.Marshal(b, m, deterministic)
}
func (dst *GetQueueMonitorInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueueMonitorInput.Merge(dst, src)
}
func (m *GetQueueMonitorInput) XXX_Size() int {
	return xxx_messageInfo_GetQueueMonitorInput.Size(m)
}
func (m *GetQueueMonitorInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueueMonitorInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueueMonitorInput proto.InternalMessageInfo

type GetQueueMonitorOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQueueMonitorOutput) Reset()         { *m = GetQueueMonitorOutput{} }
func (m *GetQueueMonitorOutput) String() string { return proto.CompactTextString(m) }
func (*GetQueueMonitorOutput) ProtoMessage()    {}
func (*GetQueueMonitorOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_33a9499a8c480c79, []int{11}
}
func (m *GetQueueMonitorOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQueueMonitorOutput.Unmarshal(m, b)
}
func (m *GetQueueMonitorOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQueueMonitorOutput.Marshal(b, m, deterministic)
}
func (dst *GetQueueMonitorOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueueMonitorOutput.Merge(dst, src)
}
func (m *GetQueueMonitorOutput) XXX_Size() int {
	return xxx_messageInfo_GetQueueMonitorOutput.Size(m)
}
func (m *GetQueueMonitorOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueueMonitorOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueueMonitorOutput proto.InternalMessageInfo

func (m *GetQueueMonitorOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetQueueMonitorOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetQueueMonitorOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMonitorInput)(nil), "service.GetMonitorInput")
	proto.RegisterType((*GetMonitorOutput)(nil), "service.GetMonitorOutput")
	proto.RegisterType((*GetMonitorOutput_ResponseItem)(nil), "service.GetMonitorOutput.ResponseItem")
	proto.RegisterType((*GetLoadBalancerMonitorInput)(nil), "service.GetLoadBalancerMonitorInput")
	proto.RegisterType((*GetLoadBalancerMonitorOutput)(nil), "service.GetLoadBalancerMonitorOutput")
	proto.RegisterType((*GetLoadBalancerMonitorOutput_ResponseItem)(nil), "service.GetLoadBalancerMonitorOutput.ResponseItem")
	proto.RegisterType((*GetRDBMonitorInput)(nil), "service.GetRDBMonitorInput")
	proto.RegisterType((*GetRDBMonitorOutput)(nil), "service.GetRDBMonitorOutput")
	proto.RegisterType((*GetCacheMonitorInput)(nil), "service.GetCacheMonitorInput")
	proto.RegisterType((*GetCacheMonitorOutput)(nil), "service.GetCacheMonitorOutput")
	proto.RegisterType((*GetZooKeeperMonitorInput)(nil), "service.GetZooKeeperMonitorInput")
	proto.RegisterType((*GetZooKeeperMonitorOutput)(nil), "service.GetZooKeeperMonitorOutput")
	proto.RegisterType((*GetQueueMonitorInput)(nil), "service.GetQueueMonitorInput")
	proto.RegisterType((*GetQueueMonitorOutput)(nil), "service.GetQueueMonitorOutput")
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor_monitor_33a9499a8c480c79) }

var fileDescriptor_monitor_33a9499a8c480c79 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x96, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0xb4, 0x69, 0x9a, 0xe9, 0xaf, 0x3f, 0xd0, 0xd2, 0x46, 0xae, 0x5b, 0x68, 0xb0,
	0x00, 0xe5, 0x64, 0x8b, 0xa0, 0x1e, 0xe8, 0x09, 0x5a, 0xa4, 0x28, 0x50, 0x54, 0x70, 0x39, 0x21,
	0xa4, 0xb0, 0xf5, 0x0e, 0x89, 0xa5, 0xd8, 0xeb, 0x7a, 0xc7, 0x55, 0x9f, 0x84, 0x2b, 0x47, 0xde,
	0x81, 0x33, 0x37, 0x1e, 0x80, 0x67, 0xe1, 0x86, 0xfc, 0x27, 0xad, 0x9d, 0x38, 0x2d, 0x97, 0x08,
	0x89, 0x5b, 0xbf, 0xbb, 0x3b, 0x33, 0xdf, 0x99, 0xfd, 0x78, 0x1b, 0x58, 0xf7, 0x65, 0xe0, 0x91,
	0x8c, 0xac, 0x30, 0x92, 0x24, 0x59, 0x43, 0x61, 0x74, 0xee, 0xb9, 0x68, 0xe8, 0x2a, 0x44, 0x77,
	0xe0, 0x23, 0x71, 0xc1, 0x89, 0xdb, 0x89, 0xca, 0x8e, 0x18, 0xbb, 0x43, 0x29, 0x87, 0x63, 0xb4,
	0x53, 0x75, 0x1a, 0x7f, 0xb2, 0xc9, 0xf3, 0x51, 0x11, 0xf7, 0xc3, 0xec, 0x80, 0xf9, 0x43, 0x83,
	0x5b, 0x3d, 0xa4, 0xd7, 0x59, 0xe2, 0x7e, 0x10, 0xc6, 0xc4, 0x0c, 0x58, 0x8d, 0x50, 0xc9, 0x38,
	0x72, 0x51, 0xd7, 0xda, 0x5a, 0xa7, 0xe9, 0x5c, 0x6a, 0xd6, 0x82, 0x15, 0x1f, 0x09, 0x23, 0xa5,
	0xd7, 0xda, 0x4b, 0x9d, 0xa6, 0x93, 0x2b, 0xc6, 0x60, 0x59, 0x11, 0x86, 0xfa, 0x52, 0x7a, 0x3e,
	0xfd, 0x9b, 0x3d, 0x05, 0x50, 0xc4, 0x23, 0x1a, 0x24, 0x45, 0xf5, 0xe5, 0xb6, 0xd6, 0x59, 0xeb,
	0x1a, 0x56, 0xe6, 0xc8, 0x9a, 0x38, 0xb2, 0xde, 0x4d, 0x1c, 0x39, 0xcd, 0xf4, 0x74, 0xa2, 0xd9,
	0x1e, 0xac, 0x62, 0x20, 0xb2, 0xc0, 0xfa, 0x8d, 0x81, 0x0d, 0x0c, 0x44, 0xa2, 0xcc, 0xaf, 0x35,
	0xb8, 0x7d, 0xd5, 0xcd, 0x71, 0x4c, 0x49, 0x3b, 0x2d, 0x58, 0xe1, 0x2e, 0x79, 0x32, 0xc8, 0x9b,
	0xc9, 0x15, 0xdb, 0x4a, 0xda, 0xa4, 0x81, 0x2b, 0x05, 0xea, 0xb5, 0xb6, 0xd6, 0xa9, 0x3b, 0x8d,
	0x08, 0xe9, 0x50, 0x0a, 0x64, 0x3a, 0x34, 0x7c, 0x54, 0x8a, 0x0f, 0x31, 0x6f, 0x68, 0x22, 0xd9,
	0x2e, 0xac, 0x4d, 0x66, 0x31, 0xf0, 0x44, 0xda, 0x54, 0xd3, 0x81, 0xc9, 0x52, 0x5f, 0xb0, 0x43,
	0x68, 0xa6, 0x23, 0x19, 0x28, 0x24, 0xbd, 0xde, 0x5e, 0xea, 0xac, 0x75, 0x1f, 0x59, 0xf9, 0x45,
	0x59, 0xd3, 0xde, 0x2c, 0x07, 0x55, 0x28, 0x03, 0x85, 0x7d, 0x42, 0xdf, 0x59, 0x4d, 0x03, 0x4f,
	0x90, 0x8c, 0x8f, 0xf0, 0x5f, 0x71, 0x27, 0xb1, 0x7a, 0x7e, 0x11, 0x20, 0x25, 0x25, 0x6b, 0x99,
	0xa1, 0x54, 0xf7, 0x45, 0xb2, 0x95, 0xd5, 0xf3, 0xc4, 0x95, 0x57, 0xc2, 0xa8, 0x2f, 0x92, 0x7b,
	0x54, 0x78, 0x16, 0x63, 0xe0, 0x66, 0xd3, 0xaf, 0x3b, 0x97, 0xda, 0xfc, 0xa9, 0xc1, 0x76, 0x0f,
	0xe9, 0x48, 0x72, 0x71, 0xc0, 0xc7, 0x3c, 0x70, 0x31, 0xfa, 0x47, 0x18, 0xf8, 0x52, 0x83, 0x9d,
	0xea, 0xce, 0xfe, 0x0a, 0x0f, 0xc7, 0xb3, 0x3c, 0x74, 0x8b, 0x3c, 0xcc, 0xf5, 0x39, 0x8f, 0x8d,
	0x67, 0x53, 0x6c, 0x6c, 0xc2, 0x0a, 0x7a, 0x61, 0x52, 0x3c, 0x6b, 0xa7, 0x8e, 0x5e, 0x78, 0x2d,
	0x17, 0xe6, 0x06, 0xb0, 0x1e, 0x92, 0xf3, 0xe2, 0xa0, 0x78, 0xe3, 0xe6, 0x29, 0xdc, 0x29, 0xad,
	0x2e, 0x60, 0x5a, 0x66, 0x0b, 0x36, 0x7a, 0x48, 0x87, 0xdc, 0x1d, 0x61, 0xa9, 0xb6, 0x80, 0xcd,
	0xa9, 0xf5, 0x45, 0x54, 0x37, 0x40, 0xef, 0x21, 0xbd, 0x97, 0xf2, 0x15, 0x62, 0x58, 0xe6, 0xdd,
	0x1c, 0xc1, 0x56, 0xc5, 0xde, 0xe2, 0x66, 0xf0, 0x36, 0xc6, 0xb8, 0x6a, 0x06, 0xc5, 0xf5, 0x05,
	0x54, 0xef, 0x7e, 0x5f, 0x86, 0xff, 0xf3, 0xf4, 0x27, 0x19, 0x84, 0xec, 0x39, 0xc0, 0xd5, 0xbb,
	0xc4, 0xf4, 0x8a, 0xc7, 0x2a, 0x35, 0x68, 0x6c, 0xcd, 0x7d, 0xc6, 0x18, 0x42, 0xab, 0x1a, 0x65,
	0xf6, 0xe0, 0x06, 0xd6, 0xb3, 0xd4, 0x0f, 0xff, 0xe8, 0x8b, 0x60, 0x2f, 0x61, 0xbd, 0x84, 0x28,
	0xdb, 0x2e, 0xc6, 0x4d, 0x01, 0x6d, 0xec, 0x54, 0x6f, 0xe6, 0xb9, 0xde, 0xa4, 0xff, 0xf7, 0x8a,
	0xc8, 0xb1, 0xbb, 0xc5, 0x80, 0x19, 0x48, 0x8d, 0x7b, 0xf3, 0xb6, 0xf3, 0x8c, 0x1f, 0xd2, 0x0f,
	0x68, 0x1a, 0x21, 0x76, 0xbf, 0x18, 0x56, 0x09, 0x9f, 0x61, 0x5e, 0x77, 0xa4, 0xe4, 0xb7, 0x88,
	0x47, 0xd9, 0xef, 0x0c, 0x50, 0x65, 0xbf, 0xb3, 0x5c, 0x19, 0xfb, 0xdf, 0x3e, 0xff, 0x3a, 0xda,
	0x83, 0x27, 0xc6, 0xe3, 0x11, 0x51, 0xa8, 0xf6, 0x6d, 0x5b, 0x48, 0x57, 0x59, 0x67, 0x5e, 0x30,
	0x74, 0xc7, 0x32, 0x16, 0x96, 0x2b, 0x7d, 0x9b, 0x87, 0x9e, 0x9d, 0xff, 0xea, 0xb0, 0xbd, 0x40,
	0xe0, 0x85, 0x35, 0x22, 0x7f, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9c, 0x50, 0x0d, 0x8a,
	0x08, 0x00, 0x00,
}
