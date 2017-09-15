// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spark.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SparkServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SparkServiceProperties) Reset()                    { *m = SparkServiceProperties{} }
func (m *SparkServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SparkServiceProperties) ProtoMessage()               {}
func (*SparkServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *SparkServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type CreateSparkInput struct {
	SparkVersion string             `protobuf:"bytes,1,opt,name=spark_version,json=sparkVersion" json:"spark_version,omitempty"`
	NodeCount    int32              `protobuf:"varint,2,opt,name=node_count,json=nodeCount" json:"node_count,omitempty"`
	EnableHdfs   int32              `protobuf:"varint,3,opt,name=enable_hdfs,json=enableHdfs" json:"enable_hdfs,omitempty"`
	SparkType    int32              `protobuf:"varint,4,opt,name=spark_type,json=sparkType" json:"spark_type,omitempty"`
	StorageSize  int32              `protobuf:"varint,5,opt,name=storage_size,json=storageSize" json:"storage_size,omitempty"`
	Vxnet        string             `protobuf:"bytes,6,opt,name=vxnet" json:"vxnet,omitempty"`
	Description  string             `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	SparkName    string             `protobuf:"bytes,8,opt,name=spark_name,json=sparkName" json:"spark_name,omitempty"`
	SparkClass   int32              `protobuf:"varint,9,opt,name=spark_class,json=sparkClass" json:"spark_class,omitempty"`
	PrivateIps   []*SparkPrivateIps `protobuf:"bytes,10,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
}

func (m *CreateSparkInput) Reset()                    { *m = CreateSparkInput{} }
func (m *CreateSparkInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSparkInput) ProtoMessage()               {}
func (*CreateSparkInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *CreateSparkInput) GetSparkVersion() string {
	if m != nil {
		return m.SparkVersion
	}
	return ""
}

func (m *CreateSparkInput) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *CreateSparkInput) GetEnableHdfs() int32 {
	if m != nil {
		return m.EnableHdfs
	}
	return 0
}

func (m *CreateSparkInput) GetSparkType() int32 {
	if m != nil {
		return m.SparkType
	}
	return 0
}

func (m *CreateSparkInput) GetStorageSize() int32 {
	if m != nil {
		return m.StorageSize
	}
	return 0
}

func (m *CreateSparkInput) GetVxnet() string {
	if m != nil {
		return m.Vxnet
	}
	return ""
}

func (m *CreateSparkInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateSparkInput) GetSparkName() string {
	if m != nil {
		return m.SparkName
	}
	return ""
}

func (m *CreateSparkInput) GetSparkClass() int32 {
	if m != nil {
		return m.SparkClass
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
	Action       string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode      int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message      string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId      string   `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
	SparkName    string   `protobuf:"bytes,5,opt,name=spark_name,json=sparkName" json:"spark_name,omitempty"`
	SparkVersion string   `protobuf:"bytes,6,opt,name=spark_version,json=sparkVersion" json:"spark_version,omitempty"`
	VxnetId      string   `protobuf:"bytes,7,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	SparkNodeIds []string `protobuf:"bytes,8,rep,name=spark_node_ids,json=sparkNodeIds" json:"spark_node_ids,omitempty"`
}

func (m *CreateSparkOutput) Reset()                    { *m = CreateSparkOutput{} }
func (m *CreateSparkOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSparkOutput) ProtoMessage()               {}
func (*CreateSparkOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

func (m *CreateSparkOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateSparkOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateSparkOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkId() string {
	if m != nil {
		return m.SparkId
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkName() string {
	if m != nil {
		return m.SparkName
	}
	return ""
}

func (m *CreateSparkOutput) GetSparkVersion() string {
	if m != nil {
		return m.SparkVersion
	}
	return ""
}

func (m *CreateSparkOutput) GetVxnetId() string {
	if m != nil {
		return m.VxnetId
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
	Sparks     []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
	Status     []string `protobuf:"bytes,2,rep,name=status" json:"status,omitempty"`
	SearchWord string   `protobuf:"bytes,3,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags       []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Verbose    int32    `protobuf:"varint,5,opt,name=verbose" json:"verbose,omitempty"`
	Offset     int32    `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit      int32    `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeSparksInput) Reset()                    { *m = DescribeSparksInput{} }
func (m *DescribeSparksInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSparksInput) ProtoMessage()               {}
func (*DescribeSparksInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

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
	if m != nil {
		return m.SearchWord
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
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeSparksInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeSparksInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeSparksOutput struct {
	Action     string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkSet   []string `protobuf:"bytes,4,rep,name=spark_set,json=sparkSet" json:"spark_set,omitempty"`
	TotalCount int32    `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeSparksOutput) Reset()                    { *m = DescribeSparksOutput{} }
func (m *DescribeSparksOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSparksOutput) ProtoMessage()               {}
func (*DescribeSparksOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{4} }

func (m *DescribeSparksOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeSparksOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeSparksOutput) GetMessage() string {
	if m != nil {
		return m.Message
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
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type AddSparkNodesInput struct {
	Spark         string             `protobuf:"bytes,1,opt,name=spark" json:"spark,omitempty"`
	NodeCount     int32              `protobuf:"varint,2,opt,name=node_count,json=nodeCount" json:"node_count,omitempty"`
	SparkNodeName string             `protobuf:"bytes,3,opt,name=spark_node_name,json=sparkNodeName" json:"spark_node_name,omitempty"`
	PrivateIps    []*SparkPrivateIps `protobuf:"bytes,4,rep,name=private_ips,json=privateIps" json:"private_ips,omitempty"`
}

func (m *AddSparkNodesInput) Reset()                    { *m = AddSparkNodesInput{} }
func (m *AddSparkNodesInput) String() string            { return proto.CompactTextString(m) }
func (*AddSparkNodesInput) ProtoMessage()               {}
func (*AddSparkNodesInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{5} }

func (m *AddSparkNodesInput) GetSpark() string {
	if m != nil {
		return m.Spark
	}
	return ""
}

func (m *AddSparkNodesInput) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *AddSparkNodesInput) GetSparkNodeName() string {
	if m != nil {
		return m.SparkNodeName
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
	Action          string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode         int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message         string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId         string   `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
	SparkNewNodeIds []string `protobuf:"bytes,5,rep,name=spark_new_node_ids,json=sparkNewNodeIds" json:"spark_new_node_ids,omitempty"`
}

func (m *AddSparkNodesOutput) Reset()                    { *m = AddSparkNodesOutput{} }
func (m *AddSparkNodesOutput) String() string            { return proto.CompactTextString(m) }
func (*AddSparkNodesOutput) ProtoMessage()               {}
func (*AddSparkNodesOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{6} }

func (m *AddSparkNodesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AddSparkNodesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AddSparkNodesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddSparkNodesOutput) GetSparkId() string {
	if m != nil {
		return m.SparkId
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
	Spark      string   `protobuf:"bytes,1,opt,name=spark" json:"spark,omitempty"`
	SparkNodes []string `protobuf:"bytes,2,rep,name=spark_nodes,json=sparkNodes" json:"spark_nodes,omitempty"`
}

func (m *DeleteSparkNodesInput) Reset()                    { *m = DeleteSparkNodesInput{} }
func (m *DeleteSparkNodesInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSparkNodesInput) ProtoMessage()               {}
func (*DeleteSparkNodesInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{7} }

func (m *DeleteSparkNodesInput) GetSpark() string {
	if m != nil {
		return m.Spark
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
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkId string `protobuf:"bytes,4,opt,name=spark_id,json=sparkId" json:"spark_id,omitempty"`
}

func (m *DeleteSparkNodesOutput) Reset()                    { *m = DeleteSparkNodesOutput{} }
func (m *DeleteSparkNodesOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSparkNodesOutput) ProtoMessage()               {}
func (*DeleteSparkNodesOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{8} }

func (m *DeleteSparkNodesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSparkNodesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSparkNodesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteSparkNodesOutput) GetSparkId() string {
	if m != nil {
		return m.SparkId
	}
	return ""
}

type StartSparksInput struct {
	Sparks []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
}

func (m *StartSparksInput) Reset()                    { *m = StartSparksInput{} }
func (m *StartSparksInput) String() string            { return proto.CompactTextString(m) }
func (*StartSparksInput) ProtoMessage()               {}
func (*StartSparksInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{9} }

func (m *StartSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type StartSparksOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *StartSparksOutput) Reset()                    { *m = StartSparksOutput{} }
func (m *StartSparksOutput) String() string            { return proto.CompactTextString(m) }
func (*StartSparksOutput) ProtoMessage()               {}
func (*StartSparksOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{10} }

func (m *StartSparksOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *StartSparksOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *StartSparksOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StartSparksOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type StopSparksInput struct {
	Sparks []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
}

func (m *StopSparksInput) Reset()                    { *m = StopSparksInput{} }
func (m *StopSparksInput) String() string            { return proto.CompactTextString(m) }
func (*StopSparksInput) ProtoMessage()               {}
func (*StopSparksInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{11} }

func (m *StopSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type StopSparksOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *StopSparksOutput) Reset()                    { *m = StopSparksOutput{} }
func (m *StopSparksOutput) String() string            { return proto.CompactTextString(m) }
func (*StopSparksOutput) ProtoMessage()               {}
func (*StopSparksOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{12} }

func (m *StopSparksOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *StopSparksOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *StopSparksOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StopSparksOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type DeleteSparksInput struct {
	Sparks []string `protobuf:"bytes,1,rep,name=sparks" json:"sparks,omitempty"`
}

func (m *DeleteSparksInput) Reset()                    { *m = DeleteSparksInput{} }
func (m *DeleteSparksInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSparksInput) ProtoMessage()               {}
func (*DeleteSparksInput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{13} }

func (m *DeleteSparksInput) GetSparks() []string {
	if m != nil {
		return m.Sparks
	}
	return nil
}

type DeleteSparksOutput struct {
	Action   string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode  int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message  string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	SparkIds []string `protobuf:"bytes,4,rep,name=spark_ids,json=sparkIds" json:"spark_ids,omitempty"`
	JobId    string   `protobuf:"bytes,5,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *DeleteSparksOutput) Reset()                    { *m = DeleteSparksOutput{} }
func (m *DeleteSparksOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSparksOutput) ProtoMessage()               {}
func (*DeleteSparksOutput) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{14} }

func (m *DeleteSparksOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSparksOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSparksOutput) GetMessage() string {
	if m != nil {
		return m.Message
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
	if m != nil {
		return m.JobId
	}
	return ""
}

func init() {
	proto.RegisterType((*SparkServiceProperties)(nil), "service.SparkServiceProperties")
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

func init() { proto.RegisterFile("spark.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 939 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xa6, 0x5e, 0x3b, 0x7e, 0x4e, 0x9a, 0x64, 0x9a, 0x46, 0x9b, 0x4d, 0x43, 0xcd, 0x82,
	0x50, 0xa0, 0xc8, 0x96, 0xca, 0x09, 0x6e, 0xc5, 0x11, 0x60, 0x24, 0x42, 0xb1, 0x11, 0x1c, 0x57,
	0xeb, 0x9d, 0x17, 0x67, 0x5b, 0x7b, 0x67, 0x99, 0x19, 0x3b, 0x6d, 0x4e, 0x3d, 0x73, 0xe7, 0x03,
	0x70, 0x45, 0x7c, 0x0a, 0x3e, 0x08, 0x9f, 0x80, 0x0b, 0xdf, 0x00, 0xcd, 0x9b, 0xb1, 0xbd, 0x76,
	0x5c, 0xa5, 0x1c, 0x92, 0x8b, 0xb5, 0xef, 0xef, 0xbc, 0xf7, 0x7b, 0xef, 0x37, 0x1e, 0x68, 0xa8,
	0x22, 0x91, 0x2f, 0x5b, 0x85, 0x14, 0x5a, 0xb0, 0x9a, 0x42, 0x39, 0xcd, 0x52, 0x0c, 0x1b, 0xfa,
	0x75, 0x81, 0xca, 0x6a, 0xc3, 0xe3, 0x5f, 0xb2, 0x7c, 0x98, 0x8e, 0xc4, 0x84, 0xc7, 0x8a, 0xbf,
	0x8c, 0xe5, 0x64, 0x84, 0x6d, 0xf3, 0x63, 0xcd, 0xd1, 0xa7, 0x70, 0xd0, 0x37, 0x39, 0xfa, 0x36,
	0xf6, 0xb9, 0x14, 0x05, 0x4a, 0x9d, 0xa1, 0x62, 0x0c, 0x2a, 0x57, 0x22, 0xc7, 0xc0, 0x6b, 0x7a,
	0x27, 0xf5, 0x1e, 0x7d, 0x47, 0xff, 0x6c, 0xc0, 0x6e, 0x47, 0x62, 0xa2, 0x91, 0x82, 0xba, 0x79,
	0x31, 0xd1, 0xec, 0x03, 0xd8, 0xa6, 0x32, 0xe2, 0x29, 0x4a, 0x95, 0x89, 0xdc, 0x45, 0x6c, 0x91,
	0xf2, 0x27, 0xab, 0x63, 0xc7, 0x00, 0xb9, 0xe0, 0x18, 0xa7, 0x62, 0x92, 0xeb, 0x60, 0xa3, 0xe9,
	0x9d, 0xf8, 0xbd, 0xba, 0xd1, 0x74, 0x8c, 0x82, 0x3d, 0x86, 0x06, 0xe6, 0xc9, 0x60, 0x84, 0xf1,
	0x05, 0x3f, 0x57, 0xc1, 0x3d, 0xb2, 0x83, 0x55, 0x7d, 0xc3, 0xcf, 0x95, 0x89, 0xb7, 0x87, 0x98,
	0xde, 0x82, 0x8a, 0x8d, 0x27, 0xcd, 0x8f, 0xaf, 0x0b, 0x64, 0xef, 0xc3, 0x96, 0xd2, 0x42, 0x26,
	0x43, 0x8c, 0x55, 0x76, 0x85, 0x81, 0x4f, 0x0e, 0x0d, 0xa7, 0xeb, 0x67, 0x57, 0xc8, 0xf6, 0xc1,
	0x9f, 0xbe, 0xca, 0x51, 0x07, 0x55, 0x2a, 0xcf, 0x0a, 0xac, 0x09, 0x0d, 0x8e, 0x2a, 0x95, 0x59,
	0xa1, 0x4d, 0xe9, 0x35, 0xb2, 0x95, 0x55, 0x8b, 0x93, 0xf3, 0x64, 0x8c, 0xc1, 0x26, 0x39, 0xd8,
	0x93, 0xcf, 0x92, 0x31, 0x9a, 0xca, 0xad, 0x39, 0x1d, 0x25, 0x4a, 0x05, 0x75, 0x5b, 0x39, 0xa9,
	0x3a, 0x46, 0xc3, 0x3e, 0x87, 0x46, 0x21, 0xb3, 0x69, 0xa2, 0x31, 0xce, 0x0a, 0x15, 0x40, 0xf3,
	0xde, 0x49, 0xe3, 0x69, 0xd0, 0x72, 0xc3, 0x6a, 0x11, 0x90, 0xcf, 0xad, 0x43, 0xb7, 0x50, 0x3d,
	0x28, 0xe6, 0xdf, 0xd1, 0x9b, 0x0d, 0xd8, 0x2b, 0xc1, 0xfd, 0xfd, 0x44, 0x1b, 0xbc, 0x0f, 0xa0,
	0x9a, 0xa4, 0x7a, 0x01, 0xb4, 0x93, 0xd8, 0x21, 0x6c, 0x4a, 0xd4, 0x71, 0x2a, 0x38, 0x3a, 0x80,
	0x6b, 0x12, 0x75, 0x47, 0x70, 0x64, 0x01, 0xd4, 0xc6, 0xa8, 0x54, 0x32, 0x44, 0x82, 0xb6, 0xde,
	0x9b, 0x89, 0x26, 0xc8, 0x96, 0x9f, 0x71, 0x42, 0xb5, 0xde, 0xab, 0x91, 0xdc, 0xe5, 0x2b, 0x8d,
	0xfb, 0xab, 0x8d, 0x5f, 0x1b, 0x7b, 0x75, 0xcd, 0xd8, 0x0f, 0x61, 0x93, 0x70, 0x36, 0xe9, 0x2d,
	0xb6, 0x35, 0x92, 0xbb, 0x9c, 0x7d, 0x08, 0xf7, 0x5d, 0x7a, 0xb3, 0x17, 0x19, 0x57, 0xc1, 0x66,
	0xf3, 0xde, 0x3c, 0xc1, 0x99, 0xe0, 0xd8, 0xe5, 0x2a, 0xfa, 0xcb, 0x83, 0x07, 0xa7, 0x34, 0x8d,
	0x81, 0x05, 0x41, 0xd9, 0xa5, 0x3b, 0x80, 0x2a, 0xf9, 0xa9, 0xc0, 0xa3, 0x28, 0x27, 0x91, 0x5e,
	0x27, 0x7a, 0xa2, 0x82, 0x0d, 0xa7, 0x27, 0x89, 0xc6, 0x84, 0x89, 0x4c, 0x2f, 0xe2, 0x4b, 0x21,
	0xb9, 0x43, 0x01, 0xac, 0xea, 0x67, 0x21, 0xb9, 0x59, 0x77, 0x9d, 0x0c, 0x55, 0x50, 0xa1, 0x30,
	0xfa, 0x36, 0xb0, 0x4d, 0x51, 0x0e, 0x84, 0x9a, 0x2d, 0xd4, 0x4c, 0x34, 0xc7, 0x88, 0xf3, 0x73,
	0xe5, 0xb6, 0xc9, 0xef, 0x39, 0xc9, 0x2c, 0xd9, 0x28, 0x1b, 0x67, 0x9a, 0x9a, 0xf5, 0x7b, 0x56,
	0x88, 0x7e, 0xf7, 0x60, 0x7f, 0xb9, 0x89, 0xdb, 0x18, 0xe5, 0x11, 0xd8, 0xe9, 0xc4, 0xa6, 0x2c,
	0xdb, 0x86, 0x9d, 0x6d, 0x1f, 0x89, 0x60, 0x5a, 0xe8, 0x64, 0xe4, 0x08, 0x68, 0xdb, 0x01, 0x52,
	0x11, 0x03, 0xa3, 0x3f, 0x3d, 0x60, 0xcf, 0x38, 0xef, 0xcf, 0xc0, 0x77, 0x38, 0xef, 0x83, 0x4f,
	0x39, 0x5c, 0x81, 0x56, 0xb8, 0x89, 0xcd, 0x1f, 0xc1, 0x4e, 0x69, 0xb4, 0xb4, 0x3e, 0xb6, 0xd6,
	0xed, 0xf9, 0x6c, 0x69, 0x85, 0x56, 0xa8, 0x51, 0xf9, 0x1f, 0xd4, 0xf8, 0xc3, 0x83, 0x07, 0x4b,
	0xe5, 0xde, 0x31, 0x39, 0x9e, 0x00, 0x73, 0x2d, 0xe2, 0xe5, 0x62, 0x83, 0x7d, 0x42, 0xdd, 0x36,
	0x7f, 0x86, 0x97, 0xb3, 0x25, 0x3e, 0x83, 0x87, 0xa7, 0x38, 0x42, 0x47, 0xe3, 0x1b, 0xd1, 0x9d,
	0x5f, 0x29, 0x26, 0xef, 0x6c, 0x91, 0x61, 0x0e, 0x9d, 0x8a, 0xde, 0x78, 0x70, 0xb0, 0x9a, 0xf0,
	0x6e, 0xfb, 0x8f, 0x3e, 0x81, 0xdd, 0xbe, 0x4e, 0xa4, 0x7e, 0x07, 0x4e, 0x46, 0x97, 0xb0, 0x57,
	0xf2, 0xbd, 0x8d, 0x42, 0x1f, 0x42, 0xf5, 0x85, 0x18, 0x2c, 0xca, 0xf4, 0x5f, 0x88, 0x41, 0x97,
	0x47, 0x1f, 0xc3, 0x4e, 0x5f, 0x8b, 0xe2, 0x5d, 0x6a, 0x9c, 0x9a, 0x7e, 0x66, 0xae, 0x77, 0x58,
	0xe2, 0x13, 0xd8, 0x2b, 0x4d, 0xf2, 0x86, 0x22, 0x7f, 0xf3, 0x80, 0x95, 0xbd, 0x6f, 0xf5, 0x16,
	0x31, 0xfb, 0x5c, 0xbe, 0x45, 0xba, 0x5c, 0x95, 0x9a, 0xf0, 0x4b, 0x4d, 0x3c, 0xfd, 0xbb, 0x02,
	0x5b, 0xe5, 0x57, 0x04, 0x3b, 0x85, 0x46, 0xe9, 0x7f, 0x8b, 0x1d, 0xce, 0x29, 0xbd, 0xfa, 0x78,
	0x08, 0xc3, 0x75, 0x26, 0xd7, 0xd7, 0x77, 0x70, 0x7f, 0xf9, 0xd6, 0x64, 0x8f, 0xe6, 0xde, 0x6b,
	0xfe, 0x13, 0xc2, 0xe3, 0xb7, 0x58, 0x5d, 0xba, 0x6f, 0x61, 0x7b, 0xe9, 0xc6, 0x60, 0x47, 0x73,
	0xff, 0xeb, 0x17, 0x5f, 0xf8, 0x68, 0xbd, 0xd1, 0xe5, 0xea, 0xc3, 0xee, 0x2a, 0x01, 0xd9, 0x7b,
	0xa5, 0xe3, 0xd7, 0x90, 0x3d, 0x7c, 0xfc, 0x56, 0xbb, 0x4b, 0x7a, 0x0a, 0x8d, 0x12, 0x4f, 0x4a,
	0xa8, 0xad, 0x32, 0xad, 0x84, 0xda, 0x75, 0x62, 0x3d, 0x03, 0x58, 0x6c, 0x32, 0x2b, 0xdd, 0xa6,
	0xcb, 0x4c, 0x08, 0x0f, 0xd7, 0x58, 0x5c, 0x8a, 0xaf, 0x61, 0xab, 0xbc, 0x66, 0x2c, 0x5c, 0x57,
	0xb9, 0x4b, 0x73, 0xb4, 0xd6, 0x66, 0x13, 0x85, 0x5f, 0xfd, 0xfa, 0x6f, 0xe5, 0x4b, 0x68, 0x5f,
	0x68, 0x5d, 0xa8, 0x2f, 0xda, 0x6d, 0x2e, 0x52, 0xd5, 0x9a, 0xbf, 0x47, 0x5b, 0xa9, 0x18, 0xb7,
	0x93, 0x22, 0x6b, 0xd3, 0x7a, 0xb5, 0xb3, 0x9c, 0xe3, 0xab, 0xd6, 0x85, 0x1e, 0x8f, 0xd8, 0xce,
	0x0f, 0x59, 0x3e, 0xec, 0x90, 0x0f, 0xa5, 0x1b, 0x54, 0xe9, 0xb1, 0xfa, 0xd9, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xed, 0x4b, 0x1f, 0xe6, 0xf0, 0x0a, 0x00, 0x00,
}
