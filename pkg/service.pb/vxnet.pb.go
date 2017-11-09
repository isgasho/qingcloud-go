// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vxnet.proto

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

type VxnetServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *VxnetServiceProperties) Reset()                    { *m = VxnetServiceProperties{} }
func (m *VxnetServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*VxnetServiceProperties) ProtoMessage()               {}
func (*VxnetServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *VxnetServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeVxnetsInput struct {
	Vxnets     []string `protobuf:"bytes,1,rep,name=vxnets" json:"vxnets,omitempty"`
	VxnetType  int32    `protobuf:"varint,2,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	SearchWord string   `protobuf:"bytes,3,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags       []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Verbose    int32    `protobuf:"varint,5,opt,name=verbose" json:"verbose,omitempty"`
	Offset     int32    `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit      int32    `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeVxnetsInput) Reset()                    { *m = DescribeVxnetsInput{} }
func (m *DescribeVxnetsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVxnetsInput) ProtoMessage()               {}
func (*DescribeVxnetsInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *DescribeVxnetsInput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

func (m *DescribeVxnetsInput) GetVxnetType() int32 {
	if m != nil {
		return m.VxnetType
	}
	return 0
}

func (m *DescribeVxnetsInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeVxnetsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeVxnetsInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

func (m *DescribeVxnetsInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeVxnetsInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeVxnetsOutput struct {
	Action     string                               `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                                `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                               `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	VxnetSet   []*DescribeVxnetsOutput_ResponseItem `protobuf:"bytes,4,rep,name=vxnet_set,json=vxnetSet" json:"vxnet_set,omitempty"`
	TotalCount int32                                `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeVxnetsOutput) Reset()                    { *m = DescribeVxnetsOutput{} }
func (m *DescribeVxnetsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVxnetsOutput) ProtoMessage()               {}
func (*DescribeVxnetsOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{2} }

func (m *DescribeVxnetsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeVxnetsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeVxnetsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeVxnetsOutput) GetVxnetSet() []*DescribeVxnetsOutput_ResponseItem {
	if m != nil {
		return m.VxnetSet
	}
	return nil
}

func (m *DescribeVxnetsOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeVxnetsOutput_ResponseItem struct {
	VxnetType   int32                       `protobuf:"varint,1,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	VxnetId     string                      `protobuf:"bytes,2,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	VxnetName   string                      `protobuf:"bytes,3,opt,name=vxnet_name,json=vxnetName" json:"vxnet_name,omitempty"`
	CreateTime  *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	Description string                      `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	InstanceIds []string                    `protobuf:"bytes,6,rep,name=instance_ids,json=instanceIds" json:"instance_ids,omitempty"`
	Router      *Router                     `protobuf:"bytes,7,opt,name=router" json:"router,omitempty"`
}

func (m *DescribeVxnetsOutput_ResponseItem) Reset()         { *m = DescribeVxnetsOutput_ResponseItem{} }
func (m *DescribeVxnetsOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeVxnetsOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeVxnetsOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor30, []int{2, 0}
}

func (m *DescribeVxnetsOutput_ResponseItem) GetVxnetType() int32 {
	if m != nil {
		return m.VxnetType
	}
	return 0
}

func (m *DescribeVxnetsOutput_ResponseItem) GetVxnetId() string {
	if m != nil {
		return m.VxnetId
	}
	return ""
}

func (m *DescribeVxnetsOutput_ResponseItem) GetVxnetName() string {
	if m != nil {
		return m.VxnetName
	}
	return ""
}

func (m *DescribeVxnetsOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeVxnetsOutput_ResponseItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeVxnetsOutput_ResponseItem) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

func (m *DescribeVxnetsOutput_ResponseItem) GetRouter() *Router {
	if m != nil {
		return m.Router
	}
	return nil
}

type CreateVxnetsInput struct {
	VxnetName  string `protobuf:"bytes,1,opt,name=vxnet_name,json=vxnetName" json:"vxnet_name,omitempty"`
	VxnetType  int32  `protobuf:"varint,2,opt,name=vxnet_type,json=vxnetType" json:"vxnet_type,omitempty"`
	Count      int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	TargetUser string `protobuf:"bytes,4,opt,name=target_user,json=targetUser" json:"target_user,omitempty"`
}

func (m *CreateVxnetsInput) Reset()                    { *m = CreateVxnetsInput{} }
func (m *CreateVxnetsInput) String() string            { return proto.CompactTextString(m) }
func (*CreateVxnetsInput) ProtoMessage()               {}
func (*CreateVxnetsInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{3} }

func (m *CreateVxnetsInput) GetVxnetName() string {
	if m != nil {
		return m.VxnetName
	}
	return ""
}

func (m *CreateVxnetsInput) GetVxnetType() int32 {
	if m != nil {
		return m.VxnetType
	}
	return 0
}

func (m *CreateVxnetsInput) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CreateVxnetsInput) GetTargetUser() string {
	if m != nil {
		return m.TargetUser
	}
	return ""
}

type CreateVxnetsOutput struct {
	Action  string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Vxnets  []string `protobuf:"bytes,4,rep,name=vxnets" json:"vxnets,omitempty"`
}

func (m *CreateVxnetsOutput) Reset()                    { *m = CreateVxnetsOutput{} }
func (m *CreateVxnetsOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateVxnetsOutput) ProtoMessage()               {}
func (*CreateVxnetsOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{4} }

func (m *CreateVxnetsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateVxnetsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateVxnetsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateVxnetsOutput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

type DeleteVxnetsInput struct {
	Vxnets []string `protobuf:"bytes,1,rep,name=vxnets" json:"vxnets,omitempty"`
}

func (m *DeleteVxnetsInput) Reset()                    { *m = DeleteVxnetsInput{} }
func (m *DeleteVxnetsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteVxnetsInput) ProtoMessage()               {}
func (*DeleteVxnetsInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{5} }

func (m *DeleteVxnetsInput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

type DeleteVxnetsOutput struct {
	Action  string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Vxnets  []string `protobuf:"bytes,4,rep,name=vxnets" json:"vxnets,omitempty"`
}

func (m *DeleteVxnetsOutput) Reset()                    { *m = DeleteVxnetsOutput{} }
func (m *DeleteVxnetsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteVxnetsOutput) ProtoMessage()               {}
func (*DeleteVxnetsOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{6} }

func (m *DeleteVxnetsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteVxnetsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteVxnetsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteVxnetsOutput) GetVxnets() []string {
	if m != nil {
		return m.Vxnets
	}
	return nil
}

type JoinVxnetInput struct {
	Vxnet     string   `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	Instances []string `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *JoinVxnetInput) Reset()                    { *m = JoinVxnetInput{} }
func (m *JoinVxnetInput) String() string            { return proto.CompactTextString(m) }
func (*JoinVxnetInput) ProtoMessage()               {}
func (*JoinVxnetInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{7} }

func (m *JoinVxnetInput) GetVxnet() string {
	if m != nil {
		return m.Vxnet
	}
	return ""
}

func (m *JoinVxnetInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

type JoinVxnetOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *JoinVxnetOutput) Reset()                    { *m = JoinVxnetOutput{} }
func (m *JoinVxnetOutput) String() string            { return proto.CompactTextString(m) }
func (*JoinVxnetOutput) ProtoMessage()               {}
func (*JoinVxnetOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{8} }

func (m *JoinVxnetOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *JoinVxnetOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *JoinVxnetOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *JoinVxnetOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type LeaveVxnetInput struct {
	Vxnet     string   `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	Instances []string `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
}

func (m *LeaveVxnetInput) Reset()                    { *m = LeaveVxnetInput{} }
func (m *LeaveVxnetInput) String() string            { return proto.CompactTextString(m) }
func (*LeaveVxnetInput) ProtoMessage()               {}
func (*LeaveVxnetInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{9} }

func (m *LeaveVxnetInput) GetVxnet() string {
	if m != nil {
		return m.Vxnet
	}
	return ""
}

func (m *LeaveVxnetInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

type LeaveVxnetOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId   string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *LeaveVxnetOutput) Reset()                    { *m = LeaveVxnetOutput{} }
func (m *LeaveVxnetOutput) String() string            { return proto.CompactTextString(m) }
func (*LeaveVxnetOutput) ProtoMessage()               {}
func (*LeaveVxnetOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{10} }

func (m *LeaveVxnetOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *LeaveVxnetOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *LeaveVxnetOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LeaveVxnetOutput) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ModifyVxnetAttributesInput struct {
	Vxnet       string `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	VxnetName   string `protobuf:"bytes,2,opt,name=vxnet_name,json=vxnetName" json:"vxnet_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ModifyVxnetAttributesInput) Reset()                    { *m = ModifyVxnetAttributesInput{} }
func (m *ModifyVxnetAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyVxnetAttributesInput) ProtoMessage()               {}
func (*ModifyVxnetAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{11} }

func (m *ModifyVxnetAttributesInput) GetVxnet() string {
	if m != nil {
		return m.Vxnet
	}
	return ""
}

func (m *ModifyVxnetAttributesInput) GetVxnetName() string {
	if m != nil {
		return m.VxnetName
	}
	return ""
}

func (m *ModifyVxnetAttributesInput) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ModifyVxnetAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifyVxnetAttributesOutput) Reset()                    { *m = ModifyVxnetAttributesOutput{} }
func (m *ModifyVxnetAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyVxnetAttributesOutput) ProtoMessage()               {}
func (*ModifyVxnetAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{12} }

func (m *ModifyVxnetAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifyVxnetAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifyVxnetAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DescribeVxnetInstancesInput struct {
	Vxnet        string   `protobuf:"bytes,1,opt,name=vxnet" json:"vxnet,omitempty"`
	Instances    []string `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
	InstanceType string   `protobuf:"bytes,3,opt,name=instance_type,json=instanceType" json:"instance_type,omitempty"`
	Status       string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Image        string   `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
	Offset       int32    `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit        int32    `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
}

func (m *DescribeVxnetInstancesInput) Reset()                    { *m = DescribeVxnetInstancesInput{} }
func (m *DescribeVxnetInstancesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVxnetInstancesInput) ProtoMessage()               {}
func (*DescribeVxnetInstancesInput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{13} }

func (m *DescribeVxnetInstancesInput) GetVxnet() string {
	if m != nil {
		return m.Vxnet
	}
	return ""
}

func (m *DescribeVxnetInstancesInput) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *DescribeVxnetInstancesInput) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *DescribeVxnetInstancesInput) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeVxnetInstancesInput) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DescribeVxnetInstancesInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeVxnetInstancesInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type DescribeVxnetInstancesOutput struct {
	Action      string                                       `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode     int32                                        `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message     string                                       `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	InstanceSet []*DescribeVxnetInstancesOutput_ResponseItem `protobuf:"bytes,4,rep,name=instance_set,json=instanceSet" json:"instance_set,omitempty"`
	TotalCount  int32                                        `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeVxnetInstancesOutput) Reset()                    { *m = DescribeVxnetInstancesOutput{} }
func (m *DescribeVxnetInstancesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeVxnetInstancesOutput) ProtoMessage()               {}
func (*DescribeVxnetInstancesOutput) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{14} }

func (m *DescribeVxnetInstancesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeVxnetInstancesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput) GetInstanceSet() []*DescribeVxnetInstancesOutput_ResponseItem {
	if m != nil {
		return m.InstanceSet
	}
	return nil
}

func (m *DescribeVxnetInstancesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeVxnetInstancesOutput_ResponseItem struct {
	VxnetId          string                      `protobuf:"bytes,1,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	InstanceId       string                      `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	InstanceName     string                      `protobuf:"bytes,3,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	Description      string                      `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	InstanceType     string                      `protobuf:"bytes,5,opt,name=instance_type,json=instanceType" json:"instance_type,omitempty"`
	VcpusCurrent     int32                       `protobuf:"varint,6,opt,name=vcpus_current,json=vcpusCurrent" json:"vcpus_current,omitempty"`
	MemoryCurrent    int32                       `protobuf:"varint,7,opt,name=memory_current,json=memoryCurrent" json:"memory_current,omitempty"`
	Status           string                      `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	TransitionStatus string                      `protobuf:"bytes,9,opt,name=transition_status,json=transitionStatus" json:"transition_status,omitempty"`
	CreateTime       *google_protobuf1.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime       *google_protobuf1.Timestamp `protobuf:"bytes,11,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	ImageId          string                      `protobuf:"bytes,12,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	DhcpOptions      map[string]string           `protobuf:"bytes,13,rep,name=dhcp_options,json=dhcpOptions" json:"dhcp_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PrivateIp        string                      `protobuf:"bytes,14,opt,name=private_ip,json=privateIp" json:"private_ip,omitempty"`
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) Reset() {
	*m = DescribeVxnetInstancesOutput_ResponseItem{}
}
func (m *DescribeVxnetInstancesOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*DescribeVxnetInstancesOutput_ResponseItem) ProtoMessage()    {}
func (*DescribeVxnetInstancesOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor30, []int{14, 0}
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetVxnetId() string {
	if m != nil {
		return m.VxnetId
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetVcpusCurrent() int32 {
	if m != nil {
		return m.VcpusCurrent
	}
	return 0
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetMemoryCurrent() int32 {
	if m != nil {
		return m.MemoryCurrent
	}
	return 0
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetTransitionStatus() string {
	if m != nil {
		return m.TransitionStatus
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetDhcpOptions() map[string]string {
	if m != nil {
		return m.DhcpOptions
	}
	return nil
}

func (m *DescribeVxnetInstancesOutput_ResponseItem) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

func init() {
	proto.RegisterType((*VxnetServiceProperties)(nil), "service.VxnetServiceProperties")
	proto.RegisterType((*DescribeVxnetsInput)(nil), "service.DescribeVxnetsInput")
	proto.RegisterType((*DescribeVxnetsOutput)(nil), "service.DescribeVxnetsOutput")
	proto.RegisterType((*DescribeVxnetsOutput_ResponseItem)(nil), "service.DescribeVxnetsOutput.ResponseItem")
	proto.RegisterType((*CreateVxnetsInput)(nil), "service.CreateVxnetsInput")
	proto.RegisterType((*CreateVxnetsOutput)(nil), "service.CreateVxnetsOutput")
	proto.RegisterType((*DeleteVxnetsInput)(nil), "service.DeleteVxnetsInput")
	proto.RegisterType((*DeleteVxnetsOutput)(nil), "service.DeleteVxnetsOutput")
	proto.RegisterType((*JoinVxnetInput)(nil), "service.JoinVxnetInput")
	proto.RegisterType((*JoinVxnetOutput)(nil), "service.JoinVxnetOutput")
	proto.RegisterType((*LeaveVxnetInput)(nil), "service.LeaveVxnetInput")
	proto.RegisterType((*LeaveVxnetOutput)(nil), "service.LeaveVxnetOutput")
	proto.RegisterType((*ModifyVxnetAttributesInput)(nil), "service.ModifyVxnetAttributesInput")
	proto.RegisterType((*ModifyVxnetAttributesOutput)(nil), "service.ModifyVxnetAttributesOutput")
	proto.RegisterType((*DescribeVxnetInstancesInput)(nil), "service.DescribeVxnetInstancesInput")
	proto.RegisterType((*DescribeVxnetInstancesOutput)(nil), "service.DescribeVxnetInstancesOutput")
	proto.RegisterType((*DescribeVxnetInstancesOutput_ResponseItem)(nil), "service.DescribeVxnetInstancesOutput.ResponseItem")
}

func init() { proto.RegisterFile("vxnet.proto", fileDescriptor30) }

var fileDescriptor30 = []byte{
	// 1156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x06, 0xad, 0x93, 0x35, 0x94, 0x0f, 0xd9, 0xdf, 0xf1, 0x4f, 0xd3, 0x36, 0xec, 0xca, 0x09,
	0x6a, 0x34, 0x85, 0x04, 0xb8, 0x37, 0x45, 0x02, 0x04, 0x48, 0xe5, 0x34, 0x50, 0xd1, 0x34, 0x2d,
	0x93, 0xb4, 0x97, 0x04, 0x45, 0xae, 0x65, 0x3a, 0x22, 0x97, 0xdd, 0x5d, 0xaa, 0x51, 0xef, 0x0b,
	0xb4, 0xbd, 0xec, 0x2b, 0xf4, 0x69, 0x7a, 0xdb, 0x47, 0xe8, 0x1b, 0xf4, 0x0d, 0x0a, 0xce, 0xf2,
	0x24, 0x89, 0xf1, 0xa1, 0x85, 0x73, 0x23, 0x70, 0x66, 0x67, 0xe7, 0xb0, 0xf3, 0xed, 0x7c, 0x2b,
	0xd0, 0xa7, 0x6f, 0x43, 0x2a, 0x7b, 0x11, 0x67, 0x92, 0x91, 0x96, 0xa0, 0x7c, 0xea, 0xbb, 0xd4,
	0xd4, 0xe5, 0x2c, 0xa2, 0x42, 0x69, 0xcd, 0xfd, 0xef, 0xfd, 0x70, 0xec, 0x4e, 0x58, 0xec, 0xd9,
	0xc2, 0x7b, 0x63, 0xf3, 0x78, 0x42, 0xfb, 0xc9, 0x4f, 0xba, 0x7c, 0x30, 0x66, 0x6c, 0x3c, 0xa1,
	0x7d, 0x94, 0x46, 0xf1, 0x59, 0x5f, 0xfa, 0x01, 0x15, 0xd2, 0x09, 0x22, 0x65, 0xd0, 0xfd, 0x18,
	0xb6, 0xbf, 0x4d, 0x82, 0xbc, 0x54, 0xce, 0xbf, 0xe6, 0x2c, 0xa2, 0x5c, 0xfa, 0x54, 0x10, 0x02,
	0xf5, 0x1f, 0x59, 0x48, 0x0d, 0xed, 0x50, 0x3b, 0x6e, 0x5b, 0xf8, 0xdd, 0xfd, 0x43, 0x83, 0xff,
	0x9d, 0x52, 0xe1, 0x72, 0x7f, 0x44, 0x71, 0x9b, 0x18, 0x86, 0x51, 0x2c, 0xc9, 0x36, 0x34, 0x31,
	0x55, 0x61, 0x68, 0x87, 0xb5, 0xe3, 0xb6, 0x95, 0x4a, 0x64, 0x1f, 0x00, 0xbf, 0xec, 0x24, 0x65,
	0x63, 0xe5, 0x50, 0x3b, 0x6e, 0x58, 0x6d, 0xd4, 0xbc, 0x9a, 0x45, 0x94, 0x1c, 0x80, 0x2e, 0xa8,
	0xc3, 0xdd, 0x73, 0xfb, 0x07, 0xc6, 0x3d, 0xa3, 0x86, 0x91, 0x40, 0xa9, 0xbe, 0x63, 0xdc, 0x4b,
	0x72, 0x90, 0xce, 0x58, 0x18, 0x75, 0xf4, 0x8a, 0xdf, 0xc4, 0x80, 0xd6, 0x94, 0xf2, 0x11, 0x13,
	0xd4, 0x68, 0xa0, 0xc3, 0x4c, 0x4c, 0xb2, 0x60, 0x67, 0x67, 0x82, 0x4a, 0xa3, 0x89, 0x0b, 0xa9,
	0x44, 0xb6, 0xa0, 0x31, 0xf1, 0x03, 0x5f, 0x1a, 0x2d, 0x54, 0x2b, 0xa1, 0xfb, 0x57, 0x0d, 0xb6,
	0xe6, 0x6b, 0x79, 0x11, 0xcb, 0xb4, 0x18, 0xc7, 0x95, 0x3e, 0x0b, 0xd3, 0xd2, 0x53, 0x89, 0xec,
	0xc0, 0x2a, 0xa7, 0xd2, 0x76, 0x99, 0x97, 0x95, 0xd2, 0xe2, 0x54, 0x0e, 0x98, 0x47, 0x93, 0x9c,
	0x02, 0x2a, 0x84, 0x33, 0xa6, 0x69, 0x11, 0x99, 0x48, 0x9e, 0x81, 0xaa, 0xd7, 0x4e, 0xd2, 0x4a,
	0xca, 0xd0, 0x4f, 0x3e, 0xea, 0xa5, 0x9d, 0xec, 0x55, 0x85, 0xef, 0x59, 0x54, 0x44, 0x2c, 0x14,
	0x74, 0x28, 0x69, 0x60, 0xad, 0x4e, 0x55, 0x73, 0x64, 0x72, 0x56, 0x92, 0x49, 0x67, 0x62, 0xbb,
	0x2c, 0x0e, 0x65, 0x5a, 0x3a, 0xa0, 0x6a, 0x90, 0x68, 0xcc, 0xdf, 0x56, 0xa0, 0x53, 0xde, 0xbb,
	0x70, 0xf8, 0xda, 0xe2, 0xe1, 0xef, 0x80, 0x72, 0x6e, 0xfb, 0x1e, 0x96, 0xd3, 0xb6, 0x5a, 0x28,
	0x0f, 0xbd, 0x62, 0x67, 0xe8, 0x04, 0x59, 0x45, 0x6a, 0xe7, 0x57, 0x4e, 0x40, 0xc9, 0x23, 0xd0,
	0x5d, 0x4e, 0x1d, 0x49, 0xed, 0x04, 0x4d, 0x46, 0xfd, 0x50, 0x3b, 0xd6, 0x4f, 0xcc, 0x9e, 0x82,
	0x5a, 0x2f, 0x83, 0x5a, 0xef, 0x55, 0x06, 0x35, 0x0b, 0x94, 0x79, 0xa2, 0x20, 0x87, 0xa0, 0x7b,
	0x58, 0x76, 0x84, 0x47, 0xdc, 0x40, 0xe7, 0x65, 0x15, 0xf9, 0x00, 0x3a, 0x7e, 0x28, 0xa4, 0x13,
	0xba, 0xd4, 0xf6, 0x3d, 0x61, 0x34, 0xb1, 0xf9, 0x7a, 0xa6, 0x1b, 0x7a, 0x82, 0x7c, 0x08, 0x4d,
	0xce, 0x62, 0x49, 0x39, 0xb6, 0x54, 0x3f, 0xd9, 0xc8, 0x8f, 0xd4, 0x42, 0xb5, 0x95, 0x2e, 0x77,
	0x7f, 0xd6, 0xe0, 0xce, 0x00, 0x83, 0x97, 0xe1, 0x3a, 0x5f, 0x9f, 0xb6, 0x58, 0xdf, 0x15, 0xa8,
	0xdd, 0x82, 0x86, 0xea, 0x41, 0x4d, 0xc1, 0x09, 0x05, 0xec, 0x8f, 0xc3, 0xc7, 0x54, 0xda, 0xb1,
	0xa0, 0x1c, 0x0f, 0xa5, 0x6d, 0x81, 0x52, 0xbd, 0x16, 0x94, 0x77, 0x67, 0x40, 0xca, 0x99, 0xdc,
	0x06, 0xd8, 0x8a, 0x6b, 0x58, 0x2f, 0x5f, 0xc3, 0xee, 0x03, 0xb8, 0x73, 0x4a, 0x27, 0x54, 0x5e,
	0xe7, 0xce, 0x26, 0x79, 0x96, 0x8d, 0xdf, 0x67, 0x9e, 0xa7, 0xb0, 0xfe, 0x05, 0xf3, 0x43, 0x0c,
	0xac, 0x92, 0xdc, 0x82, 0x06, 0xae, 0xa5, 0x51, 0x95, 0x40, 0xf6, 0xa0, 0x9d, 0xa1, 0x41, 0x18,
	0x2b, 0xe8, 0xa2, 0x50, 0x74, 0x63, 0xd8, 0xc8, 0xbd, 0xdc, 0x46, 0xf6, 0x77, 0xa1, 0x79, 0xc1,
	0x46, 0xc9, 0xb5, 0x51, 0x4d, 0x6e, 0x5c, 0xb0, 0xd1, 0xd0, 0xeb, 0x3e, 0x85, 0x8d, 0x2f, 0xa9,
	0x33, 0xa5, 0xff, 0x31, 0xfb, 0x29, 0x6c, 0x16, 0x6e, 0xde, 0x63, 0xfa, 0x02, 0xcc, 0xe7, 0xcc,
	0xf3, 0xcf, 0x66, 0x18, 0xf8, 0x89, 0x94, 0xdc, 0x1f, 0xc5, 0x92, 0x8a, 0xcb, 0x2a, 0x99, 0xbf,
	0x47, 0x2b, 0x8b, 0xf7, 0x68, 0xe1, 0xaa, 0xd7, 0x96, 0xae, 0x7a, 0xf7, 0x02, 0x76, 0x2b, 0x83,
	0xde, 0x42, 0xdd, 0xdd, 0x3f, 0x35, 0xd8, 0x9d, 0x1b, 0xb8, 0xc3, 0xec, 0xcc, 0xff, 0x75, 0xb3,
	0xc8, 0x11, 0xac, 0xe5, 0xa3, 0x0a, 0x87, 0x85, 0x8a, 0x99, 0xcf, 0x2f, 0x9c, 0x17, 0xdb, 0xd0,
	0x14, 0xd2, 0x91, 0xb1, 0x48, 0x0f, 0x3c, 0x95, 0x92, 0x80, 0x7e, 0x90, 0x24, 0xaa, 0x66, 0xa0,
	0x12, 0x6e, 0x48, 0x62, 0xbf, 0xb4, 0x60, 0xaf, 0xba, 0xa8, 0xdb, 0x80, 0xce, 0xeb, 0xd2, 0x64,
	0x2e, 0xf8, 0xec, 0xa4, 0x9a, 0xcf, 0x16, 0x32, 0x99, 0xe7, 0xb5, 0x7c, 0x9a, 0x5f, 0x8b, 0xda,
	0x7e, 0x6f, 0x2c, 0x50, 0x5b, 0x99, 0xbb, 0xb4, 0x79, 0xee, 0x3a, 0x00, 0xbd, 0xc4, 0x1e, 0x29,
	0x28, 0xa1, 0x20, 0x8f, 0xb9, 0x9e, 0x95, 0xf8, 0x2d, 0xaf, 0xac, 0x0a, 0xba, 0xf5, 0x65, 0x96,
	0x5a, 0x6a, 0x7d, 0xa3, 0xa2, 0xf5, 0x47, 0xb0, 0x36, 0x75, 0xa3, 0x58, 0xd8, 0x6e, 0xcc, 0x39,
	0x0d, 0xb3, 0x9e, 0x76, 0x50, 0x39, 0x50, 0x3a, 0x72, 0x1f, 0xd6, 0x03, 0x1a, 0x30, 0x3e, 0xcb,
	0xad, 0x54, 0x8b, 0xd7, 0x94, 0x36, 0x33, 0x2b, 0x60, 0xb4, 0x3a, 0x07, 0xa3, 0x07, 0x70, 0x47,
	0x72, 0x27, 0x14, 0x7e, 0x92, 0x96, 0x9d, 0x9a, 0xb4, 0xd1, 0x64, 0xb3, 0x58, 0x78, 0xa9, 0x8c,
	0x17, 0xa8, 0x1b, 0x6e, 0x44, 0xdd, 0x8f, 0x40, 0x57, 0xee, 0xd5, 0x66, 0xfd, 0xea, 0xcd, 0xca,
	0x1c, 0x37, 0xef, 0xc0, 0x2a, 0x02, 0x3c, 0x69, 0x4a, 0x47, 0xb5, 0x0c, 0xe5, 0xa1, 0x47, 0xce,
	0xa0, 0xe3, 0x9d, 0xbb, 0x91, 0xcd, 0xf0, 0x64, 0x85, 0xb1, 0x86, 0xb0, 0x1a, 0xdc, 0x1c, 0x56,
	0xbd, 0xd3, 0x73, 0x37, 0x7a, 0xa1, 0xbc, 0x3c, 0x0d, 0x25, 0x9f, 0x59, 0xba, 0x57, 0x68, 0x92,
	0x71, 0x15, 0x71, 0x7f, 0x9a, 0x54, 0xef, 0x47, 0xc6, 0xba, 0x1a, 0x57, 0xa9, 0x66, 0x18, 0x99,
	0x8f, 0x61, 0x73, 0x71, 0x3f, 0xd9, 0x84, 0xda, 0x1b, 0x3a, 0x4b, 0x31, 0x96, 0x7c, 0xe2, 0x98,
	0x70, 0x26, 0x71, 0x36, 0xee, 0x94, 0xf0, 0x70, 0xe5, 0x53, 0xed, 0xe4, 0xa7, 0x06, 0x74, 0xca,
	0x6f, 0x69, 0xf2, 0x1c, 0xd6, 0xe7, 0x5f, 0x78, 0x64, 0xef, 0x1d, 0x4f, 0x3f, 0x9c, 0x40, 0xe6,
	0xfe, 0xa5, 0x0f, 0x43, 0xf2, 0x0c, 0x3a, 0xe5, 0x07, 0x04, 0x31, 0x73, 0xf3, 0xa5, 0x17, 0x8e,
	0xb9, 0x5b, 0xb9, 0x56, 0x38, 0x2a, 0x33, 0x7c, 0xc9, 0xd1, 0xd2, 0x2b, 0xa1, 0xe4, 0xa8, 0xe2,
	0x51, 0xf0, 0x18, 0xda, 0x39, 0xd3, 0x92, 0xff, 0xe7, 0x96, 0xf3, 0x1c, 0x6e, 0x1a, 0xcb, 0x0b,
	0xe9, 0xfe, 0x27, 0x00, 0x05, 0xd7, 0x91, 0xc2, 0x6e, 0x81, 0x47, 0xcd, 0x9d, 0x8a, 0x95, 0xd4,
	0xc5, 0x08, 0xee, 0x56, 0x32, 0x08, 0x39, 0xca, 0xf7, 0xbc, 0x9b, 0xd6, 0xcc, 0x7b, 0x97, 0x1b,
	0xa5, 0x31, 0x28, 0x6c, 0x57, 0x43, 0x90, 0xdc, 0xbb, 0x02, 0xa3, 0x2a, 0xca, 0xfd, 0x6b, 0x21,
	0xd9, 0xfc, 0xfc, 0xd7, 0xbf, 0xeb, 0x9f, 0x41, 0xff, 0x5c, 0xca, 0x48, 0x3c, 0xec, 0xf7, 0x3d,
	0xe6, 0x8a, 0x5e, 0xfe, 0xef, 0xae, 0xe7, 0xb2, 0xa0, 0xef, 0x44, 0x7e, 0x1f, 0xe7, 0x5c, 0xdf,
	0x0f, 0x3d, 0xfa, 0xb6, 0x77, 0x2e, 0x83, 0x09, 0xd9, 0xf8, 0xc6, 0x0f, 0xc7, 0x03, 0xb4, 0x41,
	0xb7, 0xa3, 0x26, 0xde, 0xc4, 0x4f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xca, 0xb9, 0x22, 0x46,
	0x3e, 0x0e, 0x00, 0x00,
}