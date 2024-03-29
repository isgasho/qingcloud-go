// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spark.proto

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

type CreateSparkInput struct {
	SparkVersion         *string            `protobuf:"bytes,1,opt,name=spark_version,json=sparkVersion" json:"spark_version,omitempty"`
	NodeCount            *int32             `protobuf:"varint,2,opt,name=node_count,json=nodeCount" json:"node_count,omitempty"`
	EnableHdfs           *int32             `protobuf:"varint,3,opt,name=enable_hdfs,json=enableHdfs" json:"enable_hdfs,omitempty"`
	SparkType            *int32             `protobuf:"varint,4,opt,name=spark_type,json=sparkType" json:"spark_type,omitempty"`
	StorageSize          *int32             `protobuf:"varint,5,opt,name=storage_size,json=storageSize" json:"storage_size,omitempty"`
	Vxnet                *string            `protobuf:"bytes,6,opt,name=vxnet" json:"vxnet,omitempty"`
	Description          *string            `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	SparkName            *string            `protobuf:"bytes,8,opt,name=spark_name,json=sparkName" json:"spark_name,omitempty"`
	SparkClass           *int32             `protobuf:"varint,9,opt,name=spark_class,json=sparkClass" json:"spark_class,omitempty"`
	PrivateIps           []*SparkPrivateIps `protobuf:"bytes,10,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateSparkInput) Reset()         { *m = CreateSparkInput{} }
func (m *CreateSparkInput) String() string { return proto.CompactTextString(m) }
func (*CreateSparkInput) ProtoMessage()    {}
func (*CreateSparkInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{0}
}
func (m *CreateSparkInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSparkInput.Unmarshal(m, b)
}
func (m *CreateSparkInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSparkInput.Marshal(b, m, deterministic)
}
func (dst *CreateSparkInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSparkInput.Merge(dst, src)
}
func (m *CreateSparkInput) XXX_Size() int {
	return xxx_messageInfo_CreateSparkInput.Size(m)
}
func (m *CreateSparkInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSparkInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSparkInput proto.InternalMessageInfo

func (m *CreateSparkInput) GetSparkVersion() string {
	if m != nil && m.SparkVersion != nil {
		return *m.SparkVersion
	}
	return ""
}

func (m *CreateSparkInput) GetNodeCount() int32 {
	if m != nil && m.NodeCount != nil {
		return *m.NodeCount
	}
	return 0
}

func (m *CreateSparkInput) GetEnableHdfs() int32 {
	if m != nil && m.EnableHdfs != nil {
		return *m.EnableHdfs
	}
	return 0
}

func (m *CreateSparkInput) GetSparkType() int32 {
	if m != nil && m.SparkType != nil {
		return *m.SparkType
	}
	return 0
}

func (m *CreateSparkInput) GetStorageSize() int32 {
	if m != nil && m.StorageSize != nil {
		return *m.StorageSize
	}
	return 0
}

func (m *CreateSparkInput) GetVxnet() string {
	if m != nil && m.Vxnet != nil {
		return *m.Vxnet
	}
	return ""
}

func (m *CreateSparkInput) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *CreateSparkInput) GetSparkName() string {
	if m != nil && m.SparkName != nil {
		return *m.SparkName
	}
	return ""
}

func (m *CreateSparkInput) GetSparkClass() int32 {
	if m != nil && m.SparkClass != nil {
		return *m.SparkClass
	}
	return 0
}

func (m *CreateSparkInput) GetPrivateIps() []*SparkPrivateIps {
	if m != nil {
		return m.PrivateIps
	}
	return nil
}

type CreateSparkOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId              *string  `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
	SparkName            *string  `protobuf:"bytes,5,opt,name=spark_name,json=sparkName" json:"spark_name,omitempty"`
	SparkVersion         *string  `protobuf:"bytes,6,opt,name=spark_version,json=sparkVersion" json:"spark_version,omitempty"`
	VxnetId              *string  `protobuf:"bytes,7,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	SparkNodeIds         []string `protobuf:"bytes,8,rep,name=spark_node_ids,json=sparkNodeIds" json:"spark_node_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSparkOutput) Reset()         { *m = CreateSparkOutput{} }
func (m *CreateSparkOutput) String() string { return proto.CompactTextString(m) }
func (*CreateSparkOutput) ProtoMessage()    {}
func (*CreateSparkOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{1}
}
func (m *CreateSparkOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSparkOutput.Unmarshal(m, b)
}
func (m *CreateSparkOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSparkOutput.Marshal(b, m, deterministic)
}
func (dst *CreateSparkOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSparkOutput.Merge(dst, src)
}
func (m *CreateSparkOutput) XXX_Size() int {
	return xxx_messageInfo_CreateSparkOutput.Size(m)
}
func (m *CreateSparkOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSparkOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSparkOutput proto.InternalMessageInfo

func (m *CreateSparkOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *CreateSparkOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *CreateSparkOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkId() string {
	if m != nil && m.SparkId != nil {
		return *m.SparkId
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkName() string {
	if m != nil && m.SparkName != nil {
		return *m.SparkName
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkVersion() string {
	if m != nil && m.SparkVersion != nil {
		return *m.SparkVersion
	}
	return ""
}

func (m *CreateSparkOutput) GetVxnetId() string {
	if m != nil && m.VxnetId != nil {
		return *m.VxnetId
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkNodeIds() []string {
	if m != nil {
		return m.SparkNodeIds
	}
	return nil
}

type DescribeSparksInput struct {
	Sparks               []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
	Status               []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
	SearchWord           *string  `protobuf:"bytes,3,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Verbose              *int32   `protobuf:"varint,5,opt,name=verbose" json:"verbose,omitempty"`
	Offset               *int32   `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit                *int32   `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeSparksInput) Reset()         { *m = DescribeSparksInput{} }
func (m *DescribeSparksInput) String() string { return proto.CompactTextString(m) }
func (*DescribeSparksInput) ProtoMessage()    {}
func (*DescribeSparksInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{2}
}
func (m *DescribeSparksInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSparksInput.Unmarshal(m, b)
}
func (m *DescribeSparksInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSparksInput.Marshal(b, m, deterministic)
}
func (dst *DescribeSparksInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSparksInput.Merge(dst, src)
}
func (m *DescribeSparksInput) XXX_Size() int {
	return xxx_messageInfo_DescribeSparksInput.Size(m)
}
func (m *DescribeSparksInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSparksInput.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSparksInput proto.InternalMessageInfo

func (m *DescribeSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

func (m *DescribeSparksInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeSparksInput) GetSearchWord() string {
	if m != nil && m.SearchWord != nil {
		return *m.SearchWord
	}
	return ""
}

func (m *DescribeSparksInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeSparksInput) GetVerbose() int32 {
	if m != nil && m.Verbose != nil {
		return *m.Verbose
	}
	return 0
}

func (m *DescribeSparksInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeSparksInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type DescribeSparksOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkSet             []string `protobuf:"bytes,4,rep,name=spark_set,json=sparkSet" json:"spark_set,omitempty"`
	TotalCount           *int32   `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeSparksOutput) Reset()         { *m = DescribeSparksOutput{} }
func (m *DescribeSparksOutput) String() string { return proto.CompactTextString(m) }
func (*DescribeSparksOutput) ProtoMessage()    {}
func (*DescribeSparksOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{3}
}
func (m *DescribeSparksOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeSparksOutput.Unmarshal(m, b)
}
func (m *DescribeSparksOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeSparksOutput.Marshal(b, m, deterministic)
}
func (dst *DescribeSparksOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeSparksOutput.Merge(dst, src)
}
func (m *DescribeSparksOutput) XXX_Size() int {
	return xxx_messageInfo_DescribeSparksOutput.Size(m)
}
func (m *DescribeSparksOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeSparksOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeSparksOutput proto.InternalMessageInfo

func (m *DescribeSparksOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeSparksOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeSparksOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeSparksOutput) GetSparkSet() []string {
	if m != nil {
		return m.SparkSet
	}
	return nil
}

func (m *DescribeSparksOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type AddSparkNodesInput struct {
	Spark                *string            `protobuf:"bytes,1,opt,name=spark" json:"spark,omitempty"`
	NodeCount            *int32             `protobuf:"varint,2,opt,name=node_count,json=nodeCount" json:"node_count,omitempty"`
	SparkNodeName        *string            `protobuf:"bytes,3,opt,name=spark_node_name,json=sparkNodeName" json:"spark_node_name,omitempty"`
	PrivateIps           []*SparkPrivateIps `protobuf:"bytes,4,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddSparkNodesInput) Reset()         { *m = AddSparkNodesInput{} }
func (m *AddSparkNodesInput) String() string { return proto.CompactTextString(m) }
func (*AddSparkNodesInput) ProtoMessage()    {}
func (*AddSparkNodesInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{4}
}
func (m *AddSparkNodesInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSparkNodesInput.Unmarshal(m, b)
}
func (m *AddSparkNodesInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSparkNodesInput.Marshal(b, m, deterministic)
}
func (dst *AddSparkNodesInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSparkNodesInput.Merge(dst, src)
}
func (m *AddSparkNodesInput) XXX_Size() int {
	return xxx_messageInfo_AddSparkNodesInput.Size(m)
}
func (m *AddSparkNodesInput) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSparkNodesInput.DiscardUnknown(m)
}

var xxx_messageInfo_AddSparkNodesInput proto.InternalMessageInfo

func (m *AddSparkNodesInput) GetSpark() string {
	if m != nil && m.Spark != nil {
		return *m.Spark
	}
	return ""
}

func (m *AddSparkNodesInput) GetNodeCount() int32 {
	if m != nil && m.NodeCount != nil {
		return *m.NodeCount
	}
	return 0
}

func (m *AddSparkNodesInput) GetSparkNodeName() string {
	if m != nil && m.SparkNodeName != nil {
		return *m.SparkNodeName
	}
	return ""
}

func (m *AddSparkNodesInput) GetPrivateIps() []*SparkPrivateIps {
	if m != nil {
		return m.PrivateIps
	}
	return nil
}

type AddSparkNodesOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId              *string  `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
	SparkNewNodeIds      []string `protobuf:"bytes,5,rep,name=spark_new_node_ids,json=sparkNewNodeIds" json:"spark_new_node_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSparkNodesOutput) Reset()         { *m = AddSparkNodesOutput{} }
func (m *AddSparkNodesOutput) String() string { return proto.CompactTextString(m) }
func (*AddSparkNodesOutput) ProtoMessage()    {}
func (*AddSparkNodesOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{5}
}
func (m *AddSparkNodesOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSparkNodesOutput.Unmarshal(m, b)
}
func (m *AddSparkNodesOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSparkNodesOutput.Marshal(b, m, deterministic)
}
func (dst *AddSparkNodesOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSparkNodesOutput.Merge(dst, src)
}
func (m *AddSparkNodesOutput) XXX_Size() int {
	return xxx_messageInfo_AddSparkNodesOutput.Size(m)
}
func (m *AddSparkNodesOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSparkNodesOutput.DiscardUnknown(m)
}

var xxx_messageInfo_AddSparkNodesOutput proto.InternalMessageInfo

func (m *AddSparkNodesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AddSparkNodesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AddSparkNodesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AddSparkNodesOutput) GetSparkId() string {
	if m != nil && m.SparkId != nil {
		return *m.SparkId
	}
	return ""
}

func (m *AddSparkNodesOutput) GetSparkNewNodeIds() []string {
	if m != nil {
		return m.SparkNewNodeIds
	}
	return nil
}

type DeleteSparkNodesInput struct {
	Spark                *string  `protobuf:"bytes,1,opt,name=spark" json:"spark,omitempty"`
	SparkNodes           []string `protobuf:"bytes,2,rep,name=spark_nodes,json=sparkNodes" json:"spark_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSparkNodesInput) Reset()         { *m = DeleteSparkNodesInput{} }
func (m *DeleteSparkNodesInput) String() string { return proto.CompactTextString(m) }
func (*DeleteSparkNodesInput) ProtoMessage()    {}
func (*DeleteSparkNodesInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{6}
}
func (m *DeleteSparkNodesInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSparkNodesInput.Unmarshal(m, b)
}
func (m *DeleteSparkNodesInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSparkNodesInput.Marshal(b, m, deterministic)
}
func (dst *DeleteSparkNodesInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSparkNodesInput.Merge(dst, src)
}
func (m *DeleteSparkNodesInput) XXX_Size() int {
	return xxx_messageInfo_DeleteSparkNodesInput.Size(m)
}
func (m *DeleteSparkNodesInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSparkNodesInput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSparkNodesInput proto.InternalMessageInfo

func (m *DeleteSparkNodesInput) GetSpark() string {
	if m != nil && m.Spark != nil {
		return *m.Spark
	}
	return ""
}

func (m *DeleteSparkNodesInput) GetSparkNodes() []string {
	if m != nil {
		return m.SparkNodes
	}
	return nil
}

type DeleteSparkNodesOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId              *string  `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSparkNodesOutput) Reset()         { *m = DeleteSparkNodesOutput{} }
func (m *DeleteSparkNodesOutput) String() string { return proto.CompactTextString(m) }
func (*DeleteSparkNodesOutput) ProtoMessage()    {}
func (*DeleteSparkNodesOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{7}
}
func (m *DeleteSparkNodesOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSparkNodesOutput.Unmarshal(m, b)
}
func (m *DeleteSparkNodesOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSparkNodesOutput.Marshal(b, m, deterministic)
}
func (dst *DeleteSparkNodesOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSparkNodesOutput.Merge(dst, src)
}
func (m *DeleteSparkNodesOutput) XXX_Size() int {
	return xxx_messageInfo_DeleteSparkNodesOutput.Size(m)
}
func (m *DeleteSparkNodesOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSparkNodesOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSparkNodesOutput proto.InternalMessageInfo

func (m *DeleteSparkNodesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteSparkNodesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteSparkNodesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DeleteSparkNodesOutput) GetSparkId() string {
	if m != nil && m.SparkId != nil {
		return *m.SparkId
	}
	return ""
}

type StartSparksInput struct {
	Sparks               []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartSparksInput) Reset()         { *m = StartSparksInput{} }
func (m *StartSparksInput) String() string { return proto.CompactTextString(m) }
func (*StartSparksInput) ProtoMessage()    {}
func (*StartSparksInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{8}
}
func (m *StartSparksInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartSparksInput.Unmarshal(m, b)
}
func (m *StartSparksInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartSparksInput.Marshal(b, m, deterministic)
}
func (dst *StartSparksInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartSparksInput.Merge(dst, src)
}
func (m *StartSparksInput) XXX_Size() int {
	return xxx_messageInfo_StartSparksInput.Size(m)
}
func (m *StartSparksInput) XXX_DiscardUnknown() {
	xxx_messageInfo_StartSparksInput.DiscardUnknown(m)
}

var xxx_messageInfo_StartSparksInput proto.InternalMessageInfo

func (m *StartSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type StartSparksOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId                *string  `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartSparksOutput) Reset()         { *m = StartSparksOutput{} }
func (m *StartSparksOutput) String() string { return proto.CompactTextString(m) }
func (*StartSparksOutput) ProtoMessage()    {}
func (*StartSparksOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{9}
}
func (m *StartSparksOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartSparksOutput.Unmarshal(m, b)
}
func (m *StartSparksOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartSparksOutput.Marshal(b, m, deterministic)
}
func (dst *StartSparksOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartSparksOutput.Merge(dst, src)
}
func (m *StartSparksOutput) XXX_Size() int {
	return xxx_messageInfo_StartSparksOutput.Size(m)
}
func (m *StartSparksOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_StartSparksOutput.DiscardUnknown(m)
}

var xxx_messageInfo_StartSparksOutput proto.InternalMessageInfo

func (m *StartSparksOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *StartSparksOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *StartSparksOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *StartSparksOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type StopSparksInput struct {
	Sparks               []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopSparksInput) Reset()         { *m = StopSparksInput{} }
func (m *StopSparksInput) String() string { return proto.CompactTextString(m) }
func (*StopSparksInput) ProtoMessage()    {}
func (*StopSparksInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{10}
}
func (m *StopSparksInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopSparksInput.Unmarshal(m, b)
}
func (m *StopSparksInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopSparksInput.Marshal(b, m, deterministic)
}
func (dst *StopSparksInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopSparksInput.Merge(dst, src)
}
func (m *StopSparksInput) XXX_Size() int {
	return xxx_messageInfo_StopSparksInput.Size(m)
}
func (m *StopSparksInput) XXX_DiscardUnknown() {
	xxx_messageInfo_StopSparksInput.DiscardUnknown(m)
}

var xxx_messageInfo_StopSparksInput proto.InternalMessageInfo

func (m *StopSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type StopSparksOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId                *string  `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopSparksOutput) Reset()         { *m = StopSparksOutput{} }
func (m *StopSparksOutput) String() string { return proto.CompactTextString(m) }
func (*StopSparksOutput) ProtoMessage()    {}
func (*StopSparksOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{11}
}
func (m *StopSparksOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopSparksOutput.Unmarshal(m, b)
}
func (m *StopSparksOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopSparksOutput.Marshal(b, m, deterministic)
}
func (dst *StopSparksOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopSparksOutput.Merge(dst, src)
}
func (m *StopSparksOutput) XXX_Size() int {
	return xxx_messageInfo_StopSparksOutput.Size(m)
}
func (m *StopSparksOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_StopSparksOutput.DiscardUnknown(m)
}

var xxx_messageInfo_StopSparksOutput proto.InternalMessageInfo

func (m *StopSparksOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *StopSparksOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *StopSparksOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *StopSparksOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type DeleteSparksInput struct {
	Sparks               []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSparksInput) Reset()         { *m = DeleteSparksInput{} }
func (m *DeleteSparksInput) String() string { return proto.CompactTextString(m) }
func (*DeleteSparksInput) ProtoMessage()    {}
func (*DeleteSparksInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{12}
}
func (m *DeleteSparksInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSparksInput.Unmarshal(m, b)
}
func (m *DeleteSparksInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSparksInput.Marshal(b, m, deterministic)
}
func (dst *DeleteSparksInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSparksInput.Merge(dst, src)
}
func (m *DeleteSparksInput) XXX_Size() int {
	return xxx_messageInfo_DeleteSparksInput.Size(m)
}
func (m *DeleteSparksInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSparksInput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSparksInput proto.InternalMessageInfo

func (m *DeleteSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type DeleteSparksOutput struct {
	Action               *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode              *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkIds             []string `protobuf:"bytes,4,rep,name=spark_ids,json=sparkIds" json:"spark_ids,omitempty"`
	JobId                *string  `protobuf:"bytes,5,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSparksOutput) Reset()         { *m = DeleteSparksOutput{} }
func (m *DeleteSparksOutput) String() string { return proto.CompactTextString(m) }
func (*DeleteSparksOutput) ProtoMessage()    {}
func (*DeleteSparksOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_spark_aa790e75cbcaf5ce, []int{13}
}
func (m *DeleteSparksOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSparksOutput.Unmarshal(m, b)
}
func (m *DeleteSparksOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSparksOutput.Marshal(b, m, deterministic)
}
func (dst *DeleteSparksOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSparksOutput.Merge(dst, src)
}
func (m *DeleteSparksOutput) XXX_Size() int {
	return xxx_messageInfo_DeleteSparksOutput.Size(m)
}
func (m *DeleteSparksOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSparksOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSparksOutput proto.InternalMessageInfo

func (m *DeleteSparksOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DeleteSparksOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DeleteSparksOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DeleteSparksOutput) GetSparkIds() []string {
	if m != nil {
		return m.SparkIds
	}
	return nil
}

func (m *DeleteSparksOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSparkInput)(nil), "service.CreateSparkInput")
	proto.RegisterType((*CreateSparkOutput)(nil), "service.CreateSparkOutput")
	proto.RegisterType((*DescribeSparksInput)(nil), "service.DescribeSparksInput")
	proto.RegisterType((*DescribeSparksOutput)(nil), "service.DescribeSparksOutput")
	proto.RegisterType((*AddSparkNodesInput)(nil), "service.AddSparkNodesInput")
	proto.RegisterType((*AddSparkNodesOutput)(nil), "service.AddSparkNodesOutput")
	proto.RegisterType((*DeleteSparkNodesInput)(nil), "service.DeleteSparkNodesInput")
	proto.RegisterType((*DeleteSparkNodesOutput)(nil), "service.DeleteSparkNodesOutput")
	proto.RegisterType((*StartSparksInput)(nil), "service.StartSparksInput")
	proto.RegisterType((*StartSparksOutput)(nil), "service.StartSparksOutput")
	proto.RegisterType((*StopSparksInput)(nil), "service.StopSparksInput")
	proto.RegisterType((*StopSparksOutput)(nil), "service.StopSparksOutput")
	proto.RegisterType((*DeleteSparksInput)(nil), "service.DeleteSparksInput")
	proto.RegisterType((*DeleteSparksOutput)(nil), "service.DeleteSparksOutput")
}

func init() { proto.RegisterFile("spark.proto", fileDescriptor_spark_aa790e75cbcaf5ce) }

var fileDescriptor_spark_aa790e75cbcaf5ce = []byte{
	// 906 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0x24, 0x1e, 0x3b, 0xae, 0x49, 0x76, 0x93, 0xde, 0x6c, 0x34, 0x99, 0xec, 0xb2, 0x66,
	0x40, 0x28, 0xb0, 0x92, 0x2d, 0x96, 0x0b, 0x70, 0x5b, 0x25, 0x12, 0x18, 0x41, 0x40, 0x36, 0x82,
	0xe3, 0xa8, 0x3d, 0x5d, 0x49, 0x66, 0xb1, 0xa7, 0x87, 0xe9, 0xb6, 0xbd, 0xbb, 0xa7, 0x7d, 0x09,
	0xb8, 0x73, 0x45, 0x3c, 0x01, 0x47, 0x9e, 0x80, 0x07, 0xe1, 0x05, 0xb8, 0xa1, 0xae, 0x6e, 0xdb,
	0x63, 0xc7, 0xab, 0x84, 0x43, 0x72, 0x73, 0xfd, 0x74, 0xf5, 0x57, 0x5f, 0xd5, 0xd7, 0x63, 0x08,
	0x54, 0xc1, 0xcb, 0x9f, 0xda, 0x45, 0x29, 0xb5, 0x64, 0x0d, 0x85, 0xe5, 0x24, 0x4b, 0x31, 0x0a,
	0xf4, 0xab, 0x02, 0x95, 0xf5, 0x46, 0xa1, 0x2a, 0x30, 0x4d, 0x46, 0xa8, 0xb9, 0xe0, 0x9a, 0x77,
	0x8c, 0x65, 0x23, 0xf1, 0x3f, 0x1b, 0xb0, 0x7b, 0x52, 0x22, 0xd7, 0xd8, 0x37, 0x55, 0xba, 0x79,
	0x31, 0xd6, 0xec, 0x3d, 0xd8, 0xa1, 0x9a, 0xc9, 0x04, 0x4b, 0x95, 0xc9, 0x3c, 0xf4, 0x5a, 0xde,
	0x71, 0xb3, 0xb7, 0x4d, 0xce, 0x1f, 0xac, 0x8f, 0x3d, 0x06, 0xc8, 0xa5, 0xc0, 0x24, 0x95, 0xe3,
	0x5c, 0x87, 0x1b, 0x2d, 0xef, 0xd8, 0xef, 0x35, 0x8d, 0xe7, 0xc4, 0x38, 0xd8, 0x13, 0x08, 0x30,
	0xe7, 0x83, 0x21, 0x26, 0x97, 0xe2, 0x5c, 0x85, 0x9b, 0x14, 0x07, 0xeb, 0xfa, 0x52, 0x9c, 0x2b,
	0x73, 0xde, 0x5e, 0x62, 0x80, 0x86, 0x35, 0x7b, 0x9e, 0x3c, 0xdf, 0xbf, 0x2a, 0x90, 0xbd, 0x0b,
	0xdb, 0x4a, 0xcb, 0x92, 0x5f, 0x60, 0xa2, 0xb2, 0xd7, 0x18, 0xfa, 0x94, 0x10, 0x38, 0x5f, 0x3f,
	0x7b, 0x8d, 0x6c, 0x1f, 0xfc, 0xc9, 0xcb, 0x1c, 0x75, 0x58, 0x27, 0x78, 0xd6, 0x60, 0x2d, 0x08,
	0x04, 0xaa, 0xb4, 0xcc, 0x0a, 0x6d, 0xa0, 0x37, 0x28, 0x56, 0x75, 0x2d, 0x6e, 0xce, 0xf9, 0x08,
	0xc3, 0x2d, 0x4a, 0xb0, 0x37, 0x9f, 0xf1, 0x11, 0x1a, 0xe4, 0x36, 0x9c, 0x0e, 0xb9, 0x52, 0x61,
	0xd3, 0x22, 0x27, 0xd7, 0x89, 0xf1, 0xb0, 0xcf, 0x20, 0x28, 0xca, 0x6c, 0xc2, 0x35, 0x26, 0x59,
	0xa1, 0x42, 0x68, 0x6d, 0x1e, 0x07, 0xcf, 0xc2, 0xb6, 0x63, 0xbe, 0x4d, 0x44, 0x7e, 0x67, 0x13,
	0xba, 0x85, 0xea, 0x41, 0x31, 0xff, 0x1d, 0xbf, 0xd9, 0x80, 0xbd, 0x0a, 0xdd, 0xdf, 0x8e, 0xb5,
	0xe1, 0xfb, 0x00, 0xea, 0x3c, 0xd5, 0x0b, 0xa2, 0x9d, 0xc5, 0x0e, 0x61, 0xab, 0x44, 0x9d, 0xa4,
	0x52, 0xa0, 0x23, 0xb8, 0x51, 0xa2, 0x3e, 0x91, 0x02, 0x59, 0x08, 0x8d, 0x11, 0x2a, 0xc5, 0x2f,
	0x90, 0xa8, 0x6d, 0xf6, 0x66, 0xa6, 0x39, 0x64, 0xe1, 0x67, 0x82, 0x58, 0x6d, 0xf6, 0x1a, 0x64,
	0x77, 0xc5, 0x4a, 0xe3, 0xfe, 0x6a, 0xe3, 0x57, 0xc6, 0x5e, 0x5f, 0x33, 0xf6, 0x43, 0xd8, 0x22,
	0x9e, 0x4d, 0x79, 0xcb, 0x6d, 0x83, 0xec, 0xae, 0x60, 0xef, 0xc3, 0x3d, 0x57, 0xde, 0xec, 0x45,
	0x26, 0x54, 0xb8, 0xd5, 0xda, 0x9c, 0x17, 0x38, 0x93, 0x02, 0xbb, 0x42, 0xc5, 0x7f, 0x79, 0xf0,
	0xe0, 0x94, 0xa6, 0x31, 0xb0, 0x24, 0x28, 0xbb, 0x74, 0x07, 0x50, 0xa7, 0x3c, 0x15, 0x7a, 0x74,
	0xca, 0x59, 0xe4, 0xd7, 0x5c, 0x8f, 0x55, 0xb8, 0xe1, 0xfc, 0x64, 0xd1, 0x98, 0x90, 0x97, 0xe9,
	0x65, 0x32, 0x95, 0xa5, 0x70, 0x2c, 0x80, 0x75, 0xfd, 0x28, 0x4b, 0xc1, 0x18, 0xd4, 0x34, 0xbf,
	0x50, 0x61, 0x8d, 0x8e, 0xd1, 0x6f, 0x43, 0xdb, 0x04, 0xcb, 0x81, 0x54, 0xb3, 0x85, 0x9a, 0x99,
	0xe6, 0x1a, 0x79, 0x7e, 0xae, 0xdc, 0x36, 0xf9, 0x3d, 0x67, 0x99, 0x25, 0x1b, 0x66, 0xa3, 0x4c,
	0x53, 0xb3, 0x7e, 0xcf, 0x1a, 0xf1, 0x6f, 0x1e, 0xec, 0x2f, 0x37, 0x71, 0x1b, 0xa3, 0x3c, 0x02,
	0x3b, 0x9d, 0xc4, 0xc0, 0xb2, 0x6d, 0xd8, 0xd9, 0xf6, 0x91, 0x04, 0xa6, 0xa5, 0xe6, 0x43, 0x27,
	0x40, 0xdb, 0x0e, 0x90, 0x8b, 0x14, 0x18, 0xff, 0xe1, 0x01, 0x7b, 0x2e, 0x44, 0x7f, 0x46, 0xbe,
	0xe3, 0x79, 0x1f, 0x7c, 0xaa, 0xe1, 0x00, 0x5a, 0xe3, 0x3a, 0x35, 0x7f, 0x00, 0xf7, 0x2b, 0xa3,
	0xa5, 0xf5, 0xb1, 0x58, 0x77, 0xe6, 0xb3, 0xa5, 0x15, 0x5a, 0x91, 0x46, 0xed, 0x7f, 0x48, 0xe3,
	0x77, 0x0f, 0x1e, 0x2c, 0xc1, 0xbd, 0x63, 0x71, 0x3c, 0x05, 0xe6, 0x5a, 0xc4, 0xe9, 0x62, 0x83,
	0x7d, 0x62, 0xdd, 0x36, 0x7f, 0x86, 0xd3, 0xd9, 0x12, 0x9f, 0xc1, 0xc3, 0x53, 0x1c, 0xa2, 0x93,
	0xf1, 0xb5, 0xec, 0xce, 0x9f, 0x14, 0x53, 0x77, 0xb6, 0xc8, 0x30, 0xa7, 0x4e, 0xc5, 0x6f, 0x3c,
	0x38, 0x58, 0x2d, 0x78, 0xb7, 0xfd, 0xc7, 0x1f, 0xc1, 0x6e, 0x5f, 0xf3, 0x52, 0xdf, 0x40, 0x93,
	0xf1, 0x14, 0xf6, 0x2a, 0xb9, 0xb7, 0x01, 0xf4, 0x21, 0xd4, 0x5f, 0xc8, 0xc1, 0x02, 0xa6, 0xff,
	0x42, 0x0e, 0xba, 0x22, 0xfe, 0x10, 0xee, 0xf7, 0xb5, 0x2c, 0x6e, 0x82, 0x71, 0x62, 0xfa, 0x99,
	0xa5, 0xde, 0x21, 0xc4, 0xa7, 0xb0, 0x57, 0x99, 0xe4, 0x35, 0x20, 0x7f, 0xf1, 0x80, 0x55, 0xb3,
	0x6f, 0xf5, 0x15, 0x31, 0xfb, 0x5c, 0x7d, 0x45, 0xba, 0x42, 0x55, 0x9a, 0xf0, 0x2b, 0x4d, 0x3c,
	0xfb, 0xbb, 0x06, 0xdb, 0x7d, 0xfb, 0xd2, 0x90, 0x72, 0xd9, 0x29, 0x04, 0x95, 0xef, 0x16, 0x3b,
	0x9c, 0x4b, 0x7a, 0xf5, 0xcf, 0x43, 0x14, 0xad, 0x0b, 0xb9, 0xbe, 0xbe, 0x81, 0x7b, 0xcb, 0xaf,
	0x26, 0x7b, 0x34, 0xcf, 0x5e, 0xf3, 0x4d, 0x88, 0x1e, 0xbf, 0x25, 0xea, 0xca, 0x7d, 0x05, 0x3b,
	0x4b, 0x2f, 0x06, 0x3b, 0x9a, 0xe7, 0x5f, 0x7d, 0xf8, 0xa2, 0x47, 0xeb, 0x83, 0xae, 0x56, 0x1f,
	0x76, 0x57, 0x05, 0xc8, 0xde, 0xa9, 0x5c, 0xbf, 0x46, 0xec, 0xd1, 0x93, 0xb7, 0xc6, 0x5d, 0xd1,
	0x53, 0x08, 0x2a, 0x3a, 0xa9, 0xb0, 0xb6, 0xaa, 0xb4, 0x0a, 0x6b, 0x57, 0x85, 0xf5, 0x1c, 0x60,
	0xb1, 0xc9, 0xac, 0xf2, 0x9a, 0x2e, 0x2b, 0x21, 0x3a, 0x5c, 0x13, 0x71, 0x25, 0xbe, 0x80, 0xed,
	0xea, 0x9a, 0xb1, 0x68, 0x1d, 0x72, 0x57, 0xe6, 0x68, 0x6d, 0xcc, 0x16, 0x8a, 0x3e, 0xfd, 0xf3,
	0xd7, 0x7f, 0xbf, 0xfe, 0x04, 0x3e, 0x8e, 0x3a, 0x97, 0x5a, 0x17, 0xea, 0xf3, 0x4e, 0x47, 0xc8,
	0x54, 0xb5, 0x7f, 0xce, 0xf2, 0x8b, 0x74, 0x28, 0xc7, 0xa2, 0x9d, 0xca, 0x51, 0x87, 0x17, 0x59,
	0x87, 0x36, 0xac, 0x93, 0xe5, 0x02, 0x5f, 0xb6, 0x2f, 0xf5, 0x68, 0xf8, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0xd4, 0x1a, 0x5e, 0xa7, 0x0a, 0x00, 0x00,
}
