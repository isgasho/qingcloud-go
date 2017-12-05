// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eip.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EIPServiceProperties struct {
	Zone             *string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EIPServiceProperties) Reset()                    { *m = EIPServiceProperties{} }
func (m *EIPServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*EIPServiceProperties) ProtoMessage()               {}
func (*EIPServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *EIPServiceProperties) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type DescribeEipsInput struct {
	Eips             []string `protobuf:"bytes,1,rep,name=eips" json:"eips,omitempty"`
	InstanceId       *string  `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Status           *string  `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	SearchWord       *string  `protobuf:"bytes,4,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Tags             []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	Verbose          *int32   `protobuf:"varint,6,opt,name=verbose" json:"verbose,omitempty"`
	Offset           *int32   `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
	Limit            *int32   `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DescribeEipsInput) Reset()                    { *m = DescribeEipsInput{} }
func (m *DescribeEipsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeEipsInput) ProtoMessage()               {}
func (*DescribeEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *DescribeEipsInput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

func (m *DescribeEipsInput) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *DescribeEipsInput) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *DescribeEipsInput) GetSearchWord() string {
	if m != nil && m.SearchWord != nil {
		return *m.SearchWord
	}
	return ""
}

func (m *DescribeEipsInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeEipsInput) GetVerbose() int32 {
	if m != nil && m.Verbose != nil {
		return *m.Verbose
	}
	return 0
}

func (m *DescribeEipsInput) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *DescribeEipsInput) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type DescribeEipsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	EipSet           []*EIP  `protobuf:"bytes,4,rep,name=eip_set,json=eipSet" json:"eip_set,omitempty"`
	TotalCount       *int32  `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DescribeEipsOutput) Reset()                    { *m = DescribeEipsOutput{} }
func (m *DescribeEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeEipsOutput) ProtoMessage()               {}
func (*DescribeEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DescribeEipsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DescribeEipsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DescribeEipsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DescribeEipsOutput) GetEipSet() []*EIP {
	if m != nil {
		return m.EipSet
	}
	return nil
}

func (m *DescribeEipsOutput) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

type AllocateEipsInput struct {
	Bandwidth        *int32  `protobuf:"varint,1,opt,name=bandwidth" json:"bandwidth,omitempty"`
	BillingMode      *string `protobuf:"bytes,2,opt,name=billing_mode,json=billingMode" json:"billing_mode,omitempty"`
	EipName          *string `protobuf:"bytes,3,opt,name=eip_name,json=eipName" json:"eip_name,omitempty"`
	Count            *int32  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	NeedIcp          *int32  `protobuf:"varint,5,opt,name=need_icp,json=needIcp" json:"need_icp,omitempty"`
	TargetUser       *string `protobuf:"bytes,6,opt,name=target_user,json=targetUser" json:"target_user,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AllocateEipsInput) Reset()                    { *m = AllocateEipsInput{} }
func (m *AllocateEipsInput) String() string            { return proto.CompactTextString(m) }
func (*AllocateEipsInput) ProtoMessage()               {}
func (*AllocateEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *AllocateEipsInput) GetBandwidth() int32 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

func (m *AllocateEipsInput) GetBillingMode() string {
	if m != nil && m.BillingMode != nil {
		return *m.BillingMode
	}
	return ""
}

func (m *AllocateEipsInput) GetEipName() string {
	if m != nil && m.EipName != nil {
		return *m.EipName
	}
	return ""
}

func (m *AllocateEipsInput) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *AllocateEipsInput) GetNeedIcp() int32 {
	if m != nil && m.NeedIcp != nil {
		return *m.NeedIcp
	}
	return 0
}

func (m *AllocateEipsInput) GetTargetUser() string {
	if m != nil && m.TargetUser != nil {
		return *m.TargetUser
	}
	return ""
}

type AllocateEipsOutput struct {
	Action           *string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32   `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Eips             []string `protobuf:"bytes,4,rep,name=eips" json:"eips,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AllocateEipsOutput) Reset()                    { *m = AllocateEipsOutput{} }
func (m *AllocateEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*AllocateEipsOutput) ProtoMessage()               {}
func (*AllocateEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *AllocateEipsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AllocateEipsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AllocateEipsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AllocateEipsOutput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

type ReleaseEipsInput struct {
	Eips             []string `protobuf:"bytes,1,rep,name=eips" json:"eips,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ReleaseEipsInput) Reset()                    { *m = ReleaseEipsInput{} }
func (m *ReleaseEipsInput) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEipsInput) ProtoMessage()               {}
func (*ReleaseEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *ReleaseEipsInput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

type ReleaseEipsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReleaseEipsOutput) Reset()                    { *m = ReleaseEipsOutput{} }
func (m *ReleaseEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*ReleaseEipsOutput) ProtoMessage()               {}
func (*ReleaseEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ReleaseEipsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ReleaseEipsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ReleaseEipsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *ReleaseEipsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type AssociateEipInput struct {
	Eip              *string `protobuf:"bytes,1,opt,name=eip" json:"eip,omitempty"`
	Instance         *string `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AssociateEipInput) Reset()                    { *m = AssociateEipInput{} }
func (m *AssociateEipInput) String() string            { return proto.CompactTextString(m) }
func (*AssociateEipInput) ProtoMessage()               {}
func (*AssociateEipInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *AssociateEipInput) GetEip() string {
	if m != nil && m.Eip != nil {
		return *m.Eip
	}
	return ""
}

func (m *AssociateEipInput) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

type AssociateEipOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AssociateEipOutput) Reset()                    { *m = AssociateEipOutput{} }
func (m *AssociateEipOutput) String() string            { return proto.CompactTextString(m) }
func (*AssociateEipOutput) ProtoMessage()               {}
func (*AssociateEipOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *AssociateEipOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *AssociateEipOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *AssociateEipOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *AssociateEipOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type DissociateEipsInput struct {
	Eips             []string `protobuf:"bytes,1,rep,name=eips" json:"eips,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DissociateEipsInput) Reset()                    { *m = DissociateEipsInput{} }
func (m *DissociateEipsInput) String() string            { return proto.CompactTextString(m) }
func (*DissociateEipsInput) ProtoMessage()               {}
func (*DissociateEipsInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *DissociateEipsInput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

type DissociateEipsOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DissociateEipsOutput) Reset()                    { *m = DissociateEipsOutput{} }
func (m *DissociateEipsOutput) String() string            { return proto.CompactTextString(m) }
func (*DissociateEipsOutput) ProtoMessage()               {}
func (*DissociateEipsOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *DissociateEipsOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *DissociateEipsOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *DissociateEipsOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *DissociateEipsOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ChangeEipsBandwidthInput struct {
	Eips             []string `protobuf:"bytes,1,rep,name=eips" json:"eips,omitempty"`
	Bandwidth        *int32   `protobuf:"varint,2,opt,name=bandwidth" json:"bandwidth,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ChangeEipsBandwidthInput) Reset()                    { *m = ChangeEipsBandwidthInput{} }
func (m *ChangeEipsBandwidthInput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBandwidthInput) ProtoMessage()               {}
func (*ChangeEipsBandwidthInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *ChangeEipsBandwidthInput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

func (m *ChangeEipsBandwidthInput) GetBandwidth() int32 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

type ChangeEipsBandwidthOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeEipsBandwidthOutput) Reset()                    { *m = ChangeEipsBandwidthOutput{} }
func (m *ChangeEipsBandwidthOutput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBandwidthOutput) ProtoMessage()               {}
func (*ChangeEipsBandwidthOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *ChangeEipsBandwidthOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ChangeEipsBandwidthOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ChangeEipsBandwidthOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *ChangeEipsBandwidthOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ChangeEipsBillingModeInput struct {
	Eips             []string `protobuf:"bytes,1,rep,name=eips" json:"eips,omitempty"`
	BillingMode      *string  `protobuf:"bytes,2,opt,name=billing_mode,json=billingMode" json:"billing_mode,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ChangeEipsBillingModeInput) Reset()                    { *m = ChangeEipsBillingModeInput{} }
func (m *ChangeEipsBillingModeInput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBillingModeInput) ProtoMessage()               {}
func (*ChangeEipsBillingModeInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *ChangeEipsBillingModeInput) GetEips() []string {
	if m != nil {
		return m.Eips
	}
	return nil
}

func (m *ChangeEipsBillingModeInput) GetBillingMode() string {
	if m != nil && m.BillingMode != nil {
		return *m.BillingMode
	}
	return ""
}

type ChangeEipsBillingModeOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	JobId            *string `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeEipsBillingModeOutput) Reset()                    { *m = ChangeEipsBillingModeOutput{} }
func (m *ChangeEipsBillingModeOutput) String() string            { return proto.CompactTextString(m) }
func (*ChangeEipsBillingModeOutput) ProtoMessage()               {}
func (*ChangeEipsBillingModeOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *ChangeEipsBillingModeOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ChangeEipsBillingModeOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ChangeEipsBillingModeOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *ChangeEipsBillingModeOutput) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

type ModifyEipAttributesInput struct {
	Eip              *string `protobuf:"bytes,1,opt,name=eip" json:"eip,omitempty"`
	EipName          *string `protobuf:"bytes,2,opt,name=eip_name,json=eipName" json:"eip_name,omitempty"`
	Description      *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyEipAttributesInput) Reset()                    { *m = ModifyEipAttributesInput{} }
func (m *ModifyEipAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifyEipAttributesInput) ProtoMessage()               {}
func (*ModifyEipAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *ModifyEipAttributesInput) GetEip() string {
	if m != nil && m.Eip != nil {
		return *m.Eip
	}
	return ""
}

func (m *ModifyEipAttributesInput) GetEipName() string {
	if m != nil && m.EipName != nil {
		return *m.EipName
	}
	return ""
}

func (m *ModifyEipAttributesInput) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

type ModifyEipAttributesOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ModifyEipAttributesOutput) Reset()                    { *m = ModifyEipAttributesOutput{} }
func (m *ModifyEipAttributesOutput) String() string            { return proto.CompactTextString(m) }
func (*ModifyEipAttributesOutput) ProtoMessage()               {}
func (*ModifyEipAttributesOutput) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

func (m *ModifyEipAttributesOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *ModifyEipAttributesOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *ModifyEipAttributesOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EIPServiceProperties)(nil), "service.EIPServiceProperties")
	proto.RegisterType((*DescribeEipsInput)(nil), "service.DescribeEipsInput")
	proto.RegisterType((*DescribeEipsOutput)(nil), "service.DescribeEipsOutput")
	proto.RegisterType((*AllocateEipsInput)(nil), "service.AllocateEipsInput")
	proto.RegisterType((*AllocateEipsOutput)(nil), "service.AllocateEipsOutput")
	proto.RegisterType((*ReleaseEipsInput)(nil), "service.ReleaseEipsInput")
	proto.RegisterType((*ReleaseEipsOutput)(nil), "service.ReleaseEipsOutput")
	proto.RegisterType((*AssociateEipInput)(nil), "service.AssociateEipInput")
	proto.RegisterType((*AssociateEipOutput)(nil), "service.AssociateEipOutput")
	proto.RegisterType((*DissociateEipsInput)(nil), "service.DissociateEipsInput")
	proto.RegisterType((*DissociateEipsOutput)(nil), "service.DissociateEipsOutput")
	proto.RegisterType((*ChangeEipsBandwidthInput)(nil), "service.ChangeEipsBandwidthInput")
	proto.RegisterType((*ChangeEipsBandwidthOutput)(nil), "service.ChangeEipsBandwidthOutput")
	proto.RegisterType((*ChangeEipsBillingModeInput)(nil), "service.ChangeEipsBillingModeInput")
	proto.RegisterType((*ChangeEipsBillingModeOutput)(nil), "service.ChangeEipsBillingModeOutput")
	proto.RegisterType((*ModifyEipAttributesInput)(nil), "service.ModifyEipAttributesInput")
	proto.RegisterType((*ModifyEipAttributesOutput)(nil), "service.ModifyEipAttributesOutput")
}

func init() { proto.RegisterFile("eip.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0x12, 0xcb, 0x8e, 0xc7, 0x41, 0xb1, 0xe1, 0x66, 0x17, 0xb2, 0x76, 0x8b, 0x7a, 0xd5,
	0x1f, 0xa4, 0x05, 0x6a, 0xa3, 0x7b, 0xe8, 0xa1, 0xb7, 0x34, 0x09, 0x16, 0x06, 0x36, 0x6d, 0xa0,
	0xa0, 0xe8, 0xa5, 0x80, 0x40, 0x89, 0x13, 0x9b, 0x0b, 0x49, 0x64, 0x45, 0x6a, 0xb3, 0x9b, 0x4b,
	0xd1, 0x97, 0xe8, 0x3b, 0xf4, 0x15, 0x7a, 0xeb, 0x3b, 0xf4, 0xda, 0x07, 0xe9, 0xad, 0xa0, 0x24,
	0xdb, 0x4c, 0x22, 0x7b, 0x4f, 0xf1, 0x8d, 0x33, 0x1c, 0xce, 0x7c, 0xdf, 0x90, 0xf3, 0x11, 0xfa,
	0xc8, 0xe5, 0x58, 0x16, 0x42, 0x0b, 0xd2, 0x53, 0x58, 0xbc, 0xe5, 0x09, 0xfa, 0x03, 0xfd, 0x5e,
	0xa2, 0xaa, 0xbd, 0xbe, 0xa7, 0x24, 0x26, 0x51, 0x86, 0x9a, 0x32, 0xaa, 0xe9, 0xc4, 0x58, 0xf5,
	0x4e, 0xf0, 0x15, 0x1c, 0x9e, 0x4d, 0x2f, 0x2e, 0xeb, 0x43, 0x17, 0x85, 0x90, 0x58, 0x68, 0x8e,
	0x8a, 0x10, 0xe8, 0xdc, 0x88, 0x1c, 0x3d, 0x67, 0xe4, 0x1c, 0xf5, 0xc3, 0x6a, 0x1d, 0xfc, 0xeb,
	0xc0, 0xc1, 0x29, 0xaa, 0xa4, 0xe0, 0x31, 0x9e, 0x71, 0xa9, 0xa6, 0xb9, 0x2c, 0xb5, 0x89, 0x44,
	0x2e, 0x95, 0xe7, 0x8c, 0x76, 0x4d, 0xa4, 0x59, 0x93, 0x4f, 0x60, 0xc0, 0x73, 0xa5, 0x69, 0x9e,
	0x60, 0xc4, 0x99, 0xb7, 0x53, 0x25, 0x81, 0x85, 0x6b, 0xca, 0xc8, 0x53, 0xe8, 0x2a, 0x4d, 0x75,
	0xa9, 0xbc, 0xdd, 0x6a, 0xaf, 0xb1, 0xcc, 0x41, 0x85, 0xb4, 0x48, 0xe6, 0xd1, 0xb5, 0x28, 0x98,
	0xd7, 0xa9, 0x0f, 0xd6, 0xae, 0x9f, 0x45, 0xc1, 0x4c, 0x35, 0x4d, 0x67, 0xca, 0x73, 0xeb, 0x6a,
	0x66, 0x4d, 0x3c, 0xe8, 0xbd, 0xc5, 0x22, 0x16, 0x0a, 0xbd, 0xee, 0xc8, 0x39, 0x72, 0xc3, 0x85,
	0x69, 0xca, 0x88, 0xab, 0x2b, 0x85, 0xda, 0xeb, 0x55, 0x1b, 0x8d, 0x45, 0x0e, 0xc1, 0x4d, 0x79,
	0xc6, 0xb5, 0xb7, 0x57, 0xb9, 0x6b, 0x23, 0xf8, 0xd3, 0x01, 0x62, 0xf3, 0xfb, 0xb1, 0xd4, 0x86,
	0xe0, 0x53, 0xe8, 0xd2, 0x44, 0x73, 0x91, 0x37, 0xcd, 0x68, 0x2c, 0x32, 0x84, 0xbd, 0x02, 0x75,
	0x94, 0x08, 0x86, 0x15, 0x43, 0x37, 0xec, 0x15, 0xa8, 0x4f, 0x04, 0x43, 0x83, 0x28, 0x43, 0xa5,
	0xe8, 0x0c, 0x1b, 0x7e, 0x0b, 0x93, 0x7c, 0x0e, 0x3d, 0xe4, 0x32, 0x32, 0x90, 0x3a, 0xa3, 0xdd,
	0xa3, 0xc1, 0xcb, 0xfd, 0x71, 0x73, 0x63, 0xe3, 0xb3, 0xe9, 0x45, 0xd8, 0x45, 0x2e, 0x2f, 0x51,
	0x9b, 0x3e, 0x68, 0xa1, 0x69, 0x1a, 0x25, 0xa2, 0xcc, 0xb5, 0xe7, 0x56, 0xe9, 0xa1, 0x72, 0x9d,
	0x18, 0x4f, 0xf0, 0xb7, 0x03, 0x07, 0xc7, 0x69, 0x2a, 0x12, 0xaa, 0xad, 0xbb, 0x78, 0x0e, 0xfd,
	0x98, 0xe6, 0xec, 0x9a, 0x33, 0x3d, 0xaf, 0xd0, 0xba, 0xe1, 0xca, 0x41, 0x5e, 0xc0, 0x7e, 0xcc,
	0xd3, 0x94, 0xe7, 0xb3, 0x28, 0x5b, 0x80, 0xee, 0x87, 0x83, 0xc6, 0x77, 0x6e, 0x80, 0x0f, 0x61,
	0xcf, 0xc0, 0xcb, 0x69, 0xb6, 0x44, 0x8e, 0x5c, 0xfe, 0x40, 0x33, 0x34, 0x3d, 0xab, 0xc1, 0x74,
	0xea, 0x9e, 0x55, 0x86, 0x39, 0x90, 0x23, 0xb2, 0x88, 0x27, 0xb2, 0x41, 0xd9, 0x33, 0xf6, 0x34,
	0x91, 0x15, 0x07, 0x5a, 0xcc, 0x50, 0x47, 0xa5, 0xc2, 0xa2, 0xba, 0x9a, 0x7e, 0x08, 0xb5, 0xeb,
	0x27, 0x85, 0x45, 0x50, 0x02, 0xb1, 0x29, 0x3c, 0x44, 0xbb, 0x17, 0x8f, 0xb3, 0xb3, 0x7a, 0x9c,
	0xc1, 0x17, 0xf0, 0x28, 0xc4, 0x14, 0xa9, 0xda, 0xfc, 0x88, 0x83, 0x6b, 0x38, 0xb0, 0xe2, 0x1e,
	0x02, 0xdd, 0x13, 0xe8, 0xbe, 0x11, 0xb1, 0x99, 0x90, 0xfa, 0xa1, 0xbb, 0x6f, 0x44, 0x3c, 0x65,
	0xc1, 0x31, 0x1c, 0x1c, 0x2b, 0x25, 0x12, 0x5e, 0x37, 0xa6, 0x46, 0xf8, 0x08, 0x76, 0x91, 0xcb,
	0xa6, 0xaa, 0x59, 0x12, 0x1f, 0xf6, 0x16, 0x13, 0xd5, 0x5c, 0xe5, 0xd2, 0x0e, 0xde, 0x01, 0xb1,
	0x53, 0x6c, 0x11, 0xfc, 0x97, 0xf0, 0xf8, 0x94, 0x5b, 0xa5, 0x37, 0x34, 0xf8, 0x06, 0x0e, 0x6f,
	0x87, 0x6e, 0x11, 0xe6, 0x6b, 0xf0, 0x4e, 0xe6, 0x34, 0x9f, 0x55, 0x75, 0xbf, 0x5f, 0x8c, 0xc8,
	0x7a, 0x45, 0xbb, 0x35, 0x59, 0x3b, 0x77, 0x26, 0x2b, 0xf8, 0x0d, 0x86, 0x2d, 0xd9, 0xb6, 0x48,
	0xe7, 0x12, 0x7c, 0x0b, 0xc0, 0x6a, 0xa0, 0xd7, 0x13, 0xfa, 0xb0, 0x18, 0x04, 0xbf, 0x3b, 0xf0,
	0xac, 0x35, 0xeb, 0x16, 0x89, 0x71, 0xf0, 0xce, 0x05, 0xe3, 0x57, 0xef, 0xcf, 0xb8, 0x3c, 0xd6,
	0xba, 0xe0, 0x71, 0xa9, 0x51, 0xad, 0x1b, 0x09, 0x5b, 0xbe, 0x76, 0x6e, 0xcb, 0xd7, 0x08, 0x06,
	0xac, 0xd2, 0x76, 0x59, 0x21, 0xae, 0xab, 0xdb, 0xae, 0x60, 0x0e, 0xc3, 0x96, 0x52, 0x0f, 0xc0,
	0xf5, 0xe5, 0x3f, 0x2e, 0xc0, 0xea, 0xd7, 0x25, 0xaf, 0x60, 0xdf, 0xfe, 0x76, 0x88, 0xbf, 0xfc,
	0x12, 0xee, 0xfd, 0xb6, 0xfe, 0xb3, 0xd6, 0xbd, 0x06, 0xe4, 0x2b, 0xd8, 0xb7, 0x05, 0xd5, 0x4a,
	0x74, 0xef, 0xab, 0xb0, 0x12, 0xb5, 0x68, 0xf0, 0x29, 0x0c, 0x2c, 0xe9, 0x23, 0xc3, 0x65, 0xec,
	0x5d, 0xe1, 0xf4, 0xfd, 0xb6, 0x2d, 0x0b, 0x8e, 0x35, 0xde, 0x36, 0x9c, 0xbb, 0xf2, 0x66, 0xc3,
	0xb9, 0xaf, 0x5b, 0xe7, 0xf0, 0xd1, 0x6d, 0xa1, 0x20, 0xcf, 0x57, 0x6d, 0xb8, 0x2f, 0x36, 0xfe,
	0xc7, 0x6b, 0x76, 0x9b, 0x74, 0xbf, 0xc0, 0xe3, 0x96, 0x69, 0x25, 0x2f, 0x96, 0xa7, 0xd6, 0x29,
	0x83, 0x1f, 0x6c, 0x0a, 0x69, 0xb2, 0xc7, 0xf0, 0xa4, 0x75, 0x68, 0xc8, 0xa7, 0x6d, 0x87, 0xef,
	0x8c, 0xaa, 0xff, 0xd9, 0xe6, 0xa0, 0x15, 0x83, 0x96, 0xa7, 0x6a, 0x31, 0x58, 0x37, 0x33, 0x16,
	0x83, 0xb5, 0x6f, 0xdd, 0xff, 0xf6, 0xaf, 0x3f, 0xfe, 0x7b, 0xfd, 0x0d, 0x4c, 0xfc, 0xaf, 0xe7,
	0x5a, 0x4b, 0xf5, 0xdd, 0x64, 0xc2, 0x44, 0xa2, 0xc6, 0xbf, 0xf2, 0x7c, 0x96, 0xa4, 0xa2, 0x64,
	0xe3, 0x44, 0x64, 0x13, 0x2a, 0xf9, 0x04, 0xb9, 0x9c, 0xf0, 0x9c, 0xe1, 0xbb, 0xf1, 0x5c, 0x67,
	0xe9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x69, 0xc1, 0xe7, 0x87, 0x0a, 0x00, 0x00,
}
