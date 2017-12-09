// Code generated by protoc-gen-go. DO NOT EDIT.
// source: misc.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GrantQuotaIndepInput struct {
	User               *string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Zone               *string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	Instance           *int32  `protobuf:"varint,3,opt,name=instance" json:"instance,omitempty"`
	Cpu                *int32  `protobuf:"varint,4,opt,name=cpu" json:"cpu,omitempty"`
	Memory             *int32  `protobuf:"varint,5,opt,name=memory" json:"memory,omitempty"`
	HpInstance         *int32  `protobuf:"varint,6,opt,name=hp_instance,json=hpInstance" json:"hp_instance,omitempty"`
	HpCpu              *int32  `protobuf:"varint,7,opt,name=hp_cpu,json=hpCpu" json:"hp_cpu,omitempty"`
	HpMemory           *int32  `protobuf:"varint,8,opt,name=hp_memory,json=hpMemory" json:"hp_memory,omitempty"`
	Volume             *int32  `protobuf:"varint,9,opt,name=volume" json:"volume,omitempty"`
	VolumeSize         *int32  `protobuf:"varint,10,opt,name=volume_size,json=volumeSize" json:"volume_size,omitempty"`
	HcVolume           *int32  `protobuf:"varint,11,opt,name=hc_volume,json=hcVolume" json:"hc_volume,omitempty"`
	HcVolumeSize       *int32  `protobuf:"varint,12,opt,name=hc_volume_size,json=hcVolumeSize" json:"hc_volume_size,omitempty"`
	HppVolume          *int32  `protobuf:"varint,13,opt,name=hpp_volume,json=hppVolume" json:"hpp_volume,omitempty"`
	HppVolumeSize      *int32  `protobuf:"varint,14,opt,name=hpp_volume_size,json=hppVolumeSize" json:"hpp_volume_size,omitempty"`
	Image              *int32  `protobuf:"varint,15,opt,name=image" json:"image,omitempty"`
	Loadbalancer       *int32  `protobuf:"varint,16,opt,name=loadbalancer" json:"loadbalancer,omitempty"`
	LoadbalancerPolicy *int32  `protobuf:"varint,17,opt,name=loadbalancer_policy,json=loadbalancerPolicy" json:"loadbalancer_policy,omitempty"`
	Vxnet              *int32  `protobuf:"varint,18,opt,name=vxnet" json:"vxnet,omitempty"`
	Router             *int32  `protobuf:"varint,19,opt,name=router" json:"router,omitempty"`
	Eip                *int32  `protobuf:"varint,20,opt,name=eip" json:"eip,omitempty"`
	EipBandwidth       *int32  `protobuf:"varint,21,opt,name=eip_bandwidth,json=eipBandwidth" json:"eip_bandwidth,omitempty"`
	Rdb                *int32  `protobuf:"varint,22,opt,name=rdb" json:"rdb,omitempty"`
	HppRdb             *int32  `protobuf:"varint,23,opt,name=hpp_rdb,json=hppRdb" json:"hpp_rdb,omitempty"`
	Cache              *int32  `protobuf:"varint,24,opt,name=cache" json:"cache,omitempty"`
	HpCache            *int32  `protobuf:"varint,25,opt,name=hp_cache,json=hpCache" json:"hp_cache,omitempty"`
	Mongo              *int32  `protobuf:"varint,26,opt,name=mongo" json:"mongo,omitempty"`
	HpMongo            *int32  `protobuf:"varint,27,opt,name=hp_mongo,json=hpMongo" json:"hp_mongo,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *GrantQuotaIndepInput) Reset()                    { *m = GrantQuotaIndepInput{} }
func (m *GrantQuotaIndepInput) String() string            { return proto.CompactTextString(m) }
func (*GrantQuotaIndepInput) ProtoMessage()               {}
func (*GrantQuotaIndepInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *GrantQuotaIndepInput) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *GrantQuotaIndepInput) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

func (m *GrantQuotaIndepInput) GetInstance() int32 {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetCpu() int32 {
	if m != nil && m.Cpu != nil {
		return *m.Cpu
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetMemory() int32 {
	if m != nil && m.Memory != nil {
		return *m.Memory
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpInstance() int32 {
	if m != nil && m.HpInstance != nil {
		return *m.HpInstance
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpCpu() int32 {
	if m != nil && m.HpCpu != nil {
		return *m.HpCpu
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpMemory() int32 {
	if m != nil && m.HpMemory != nil {
		return *m.HpMemory
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVolume() int32 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVolumeSize() int32 {
	if m != nil && m.VolumeSize != nil {
		return *m.VolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHcVolume() int32 {
	if m != nil && m.HcVolume != nil {
		return *m.HcVolume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHcVolumeSize() int32 {
	if m != nil && m.HcVolumeSize != nil {
		return *m.HcVolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppVolume() int32 {
	if m != nil && m.HppVolume != nil {
		return *m.HppVolume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppVolumeSize() int32 {
	if m != nil && m.HppVolumeSize != nil {
		return *m.HppVolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetImage() int32 {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetLoadbalancer() int32 {
	if m != nil && m.Loadbalancer != nil {
		return *m.Loadbalancer
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetLoadbalancerPolicy() int32 {
	if m != nil && m.LoadbalancerPolicy != nil {
		return *m.LoadbalancerPolicy
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVxnet() int32 {
	if m != nil && m.Vxnet != nil {
		return *m.Vxnet
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetRouter() int32 {
	if m != nil && m.Router != nil {
		return *m.Router
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetEip() int32 {
	if m != nil && m.Eip != nil {
		return *m.Eip
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetEipBandwidth() int32 {
	if m != nil && m.EipBandwidth != nil {
		return *m.EipBandwidth
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetRdb() int32 {
	if m != nil && m.Rdb != nil {
		return *m.Rdb
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppRdb() int32 {
	if m != nil && m.HppRdb != nil {
		return *m.HppRdb
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetCache() int32 {
	if m != nil && m.Cache != nil {
		return *m.Cache
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpCache() int32 {
	if m != nil && m.HpCache != nil {
		return *m.HpCache
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetMongo() int32 {
	if m != nil && m.Mongo != nil {
		return *m.Mongo
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpMongo() int32 {
	if m != nil && m.HpMongo != nil {
		return *m.HpMongo
	}
	return 0
}

type GrantQuotaIndepOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	UserId           *string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ZoneId           *string `protobuf:"bytes,5,opt,name=zone_id,json=zoneId" json:"zone_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GrantQuotaIndepOutput) Reset()                    { *m = GrantQuotaIndepOutput{} }
func (m *GrantQuotaIndepOutput) String() string            { return proto.CompactTextString(m) }
func (*GrantQuotaIndepOutput) ProtoMessage()               {}
func (*GrantQuotaIndepOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *GrantQuotaIndepOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GrantQuotaIndepOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetZoneId() string {
	if m != nil && m.ZoneId != nil {
		return *m.ZoneId
	}
	return ""
}

type RevokeQuotaIndepInput struct {
	User             *string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Zone             *string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RevokeQuotaIndepInput) Reset()                    { *m = RevokeQuotaIndepInput{} }
func (m *RevokeQuotaIndepInput) String() string            { return proto.CompactTextString(m) }
func (*RevokeQuotaIndepInput) ProtoMessage()               {}
func (*RevokeQuotaIndepInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *RevokeQuotaIndepInput) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *RevokeQuotaIndepInput) GetZone() string {
	if m != nil && m.Zone != nil {
		return *m.Zone
	}
	return ""
}

type RevokeQuotaIndepOutput struct {
	Action           *string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RevokeQuotaIndepOutput) Reset()                    { *m = RevokeQuotaIndepOutput{} }
func (m *RevokeQuotaIndepOutput) String() string            { return proto.CompactTextString(m) }
func (*RevokeQuotaIndepOutput) ProtoMessage()               {}
func (*RevokeQuotaIndepOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *RevokeQuotaIndepOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *RevokeQuotaIndepOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *RevokeQuotaIndepOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type GetQuotaLeftInput struct {
	ResourceTypes    []string `protobuf:"bytes,1,rep,name=resource_types,json=resourceTypes" json:"resource_types,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetQuotaLeftInput) Reset()                    { *m = GetQuotaLeftInput{} }
func (m *GetQuotaLeftInput) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaLeftInput) ProtoMessage()               {}
func (*GetQuotaLeftInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *GetQuotaLeftInput) GetResourceTypes() []string {
	if m != nil {
		return m.ResourceTypes
	}
	return nil
}

type GetQuotaLeftOutput struct {
	Action           *string                         `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode          *int32                          `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message          *string                         `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	QuotaLeftSet     []*GetQuotaLeftOutput_QuotaLeft `protobuf:"bytes,4,rep,name=quota_left_set,json=quotaLeftSet" json:"quota_left_set,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *GetQuotaLeftOutput) Reset()                    { *m = GetQuotaLeftOutput{} }
func (m *GetQuotaLeftOutput) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaLeftOutput) ProtoMessage()               {}
func (*GetQuotaLeftOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *GetQuotaLeftOutput) GetAction() string {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return ""
}

func (m *GetQuotaLeftOutput) GetRetCode() int32 {
	if m != nil && m.RetCode != nil {
		return *m.RetCode
	}
	return 0
}

func (m *GetQuotaLeftOutput) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *GetQuotaLeftOutput) GetQuotaLeftSet() []*GetQuotaLeftOutput_QuotaLeft {
	if m != nil {
		return m.QuotaLeftSet
	}
	return nil
}

type GetQuotaLeftOutput_QuotaLeft struct {
	ResourceType     *string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	Left             *int32  `protobuf:"varint,2,opt,name=left" json:"left,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetQuotaLeftOutput_QuotaLeft) Reset()         { *m = GetQuotaLeftOutput_QuotaLeft{} }
func (m *GetQuotaLeftOutput_QuotaLeft) String() string { return proto.CompactTextString(m) }
func (*GetQuotaLeftOutput_QuotaLeft) ProtoMessage()    {}
func (*GetQuotaLeftOutput_QuotaLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{5, 0}
}

func (m *GetQuotaLeftOutput_QuotaLeft) GetResourceType() string {
	if m != nil && m.ResourceType != nil {
		return *m.ResourceType
	}
	return ""
}

func (m *GetQuotaLeftOutput_QuotaLeft) GetLeft() int32 {
	if m != nil && m.Left != nil {
		return *m.Left
	}
	return 0
}

func init() {
	proto.RegisterType((*GrantQuotaIndepInput)(nil), "service.GrantQuotaIndepInput")
	proto.RegisterType((*GrantQuotaIndepOutput)(nil), "service.GrantQuotaIndepOutput")
	proto.RegisterType((*RevokeQuotaIndepInput)(nil), "service.RevokeQuotaIndepInput")
	proto.RegisterType((*RevokeQuotaIndepOutput)(nil), "service.RevokeQuotaIndepOutput")
	proto.RegisterType((*GetQuotaLeftInput)(nil), "service.GetQuotaLeftInput")
	proto.RegisterType((*GetQuotaLeftOutput)(nil), "service.GetQuotaLeftOutput")
	proto.RegisterType((*GetQuotaLeftOutput_QuotaLeft)(nil), "service.GetQuotaLeftOutput.QuotaLeft")
}

func init() { proto.RegisterFile("misc.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 777 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x86, 0x22, 0x53, 0xb2, 0xc6, 0xf2, 0x4f, 0x36, 0xb6, 0xb2, 0xa1, 0x91, 0xc6, 0x50, 0x9b,
	0x42, 0x27, 0xaa, 0xf0, 0xa5, 0x40, 0x2e, 0x05, 0xea, 0x02, 0x86, 0xd0, 0x18, 0x4d, 0xa9, 0xa2,
	0x57, 0x82, 0xda, 0x9d, 0x98, 0x8b, 0x8a, 0xdc, 0x35, 0xb9, 0x54, 0x63, 0x3f, 0x48, 0x1e, 0xa5,
	0x87, 0x3e, 0x57, 0x2f, 0xb9, 0x15, 0xb3, 0x4b, 0xca, 0xaa, 0xea, 0xf4, 0x50, 0xc0, 0xb7, 0xf9,
	0xe6, 0x9b, 0xf9, 0x66, 0x76, 0x76, 0x96, 0x04, 0xc8, 0x55, 0x25, 0x22, 0x53, 0x6a, 0xab, 0x59,
	0xbf, 0xc2, 0x72, 0xa5, 0x04, 0x86, 0xbc, 0x32, 0x28, 0x92, 0x1c, 0x6d, 0x2a, 0x53, 0x9b, 0x4e,
	0x09, 0xf9, 0x90, 0xf1, 0xa7, 0x00, 0x8e, 0x2f, 0xcb, 0xb4, 0xb0, 0x3f, 0xd7, 0xda, 0xa6, 0xb3,
	0x42, 0xa2, 0x99, 0x15, 0xa6, 0xb6, 0x8c, 0xc1, 0x4e, 0x5d, 0x61, 0xc9, 0x3b, 0x67, 0x9d, 0xc9,
	0x20, 0x76, 0x36, 0xf9, 0xee, 0x74, 0x81, 0xfc, 0x89, 0xf7, 0x91, 0xcd, 0x42, 0xd8, 0x55, 0x45,
	0x65, 0xd3, 0x42, 0x20, 0xef, 0x9e, 0x75, 0x26, 0x41, 0xbc, 0xc6, 0xec, 0x08, 0xba, 0xc2, 0xd4,
	0x7c, 0xc7, 0xb9, 0xc9, 0x64, 0x23, 0xe8, 0xe5, 0x98, 0xeb, 0xf2, 0x96, 0x07, 0xce, 0xd9, 0x20,
	0xf6, 0x0a, 0xf6, 0x32, 0x93, 0xac, 0x85, 0x7a, 0x8e, 0x84, 0xcc, 0xcc, 0x5a, 0xa9, 0x13, 0xe8,
	0x65, 0x26, 0x21, 0xb5, 0xbe, 0xe3, 0x82, 0xcc, 0x5c, 0x98, 0x9a, 0x9d, 0xc2, 0x20, 0x33, 0x49,
	0x23, 0xb9, 0xeb, 0xcb, 0x67, 0xe6, 0xca, 0x8b, 0x8e, 0xa0, 0xb7, 0xd2, 0xcb, 0x3a, 0x47, 0x3e,
	0xf0, 0xc5, 0x3c, 0xa2, 0x62, 0xde, 0x4a, 0x2a, 0x75, 0x87, 0x1c, 0x7c, 0x31, 0xef, 0x9a, 0xab,
	0x3b, 0x74, 0xaa, 0x22, 0x69, 0x72, 0xf7, 0x1a, 0x55, 0xf1, 0xab, 0xcf, 0xfe, 0x0a, 0x0e, 0xd6,
	0xa4, 0x17, 0x18, 0xba, 0x88, 0x61, 0x1b, 0xe1, 0x24, 0x5e, 0x02, 0x64, 0xc6, 0xb4, 0x1a, 0xfb,
	0x2e, 0x62, 0x90, 0x19, 0xd3, 0x88, 0x7c, 0x0d, 0x87, 0xf7, 0xb4, 0x57, 0x39, 0x70, 0x31, 0xfb,
	0xeb, 0x18, 0x27, 0x73, 0x0c, 0x81, 0xca, 0xd3, 0x6b, 0xe4, 0x87, 0xfe, 0xd4, 0x0e, 0xb0, 0x31,
	0x0c, 0x97, 0x3a, 0x95, 0x8b, 0x74, 0x49, 0xb3, 0x29, 0xf9, 0x91, 0x6f, 0x60, 0xd3, 0xc7, 0xa6,
	0xf0, 0x6c, 0x13, 0x27, 0x46, 0x2f, 0x95, 0xb8, 0xe5, 0x4f, 0x5d, 0x28, 0xdb, 0xa4, 0xde, 0x39,
	0x86, 0x4a, 0xad, 0x3e, 0x14, 0x68, 0x39, 0xf3, 0xa5, 0x1c, 0xa0, 0x19, 0x96, 0xba, 0xb6, 0x58,
	0xf2, 0x67, 0x7e, 0x86, 0x1e, 0xd1, 0xd5, 0xa2, 0x32, 0xfc, 0xd8, 0x5f, 0x2d, 0x2a, 0xc3, 0xbe,
	0x84, 0x7d, 0x54, 0x26, 0x59, 0xa4, 0x85, 0xfc, 0x5d, 0x49, 0x9b, 0xf1, 0x13, 0xdf, 0x15, 0x2a,
	0xf3, 0x7d, 0xeb, 0xa3, 0xb4, 0x52, 0x2e, 0xf8, 0xc8, 0xa7, 0x95, 0x72, 0xc1, 0x9e, 0x43, 0x9f,
	0x26, 0x41, 0xde, 0xe7, 0xbe, 0x42, 0x66, 0x4c, 0x2c, 0x17, 0xd4, 0x8f, 0x48, 0x45, 0x86, 0x9c,
	0xfb, 0x7e, 0x1c, 0x60, 0x2f, 0x60, 0x97, 0xf6, 0xc0, 0x11, 0x2f, 0x1c, 0xd1, 0xcf, 0xcc, 0x85,
	0xa3, 0x8e, 0x21, 0xc8, 0x75, 0x71, 0xad, 0x79, 0xe8, 0x13, 0x1c, 0x68, 0x12, 0x3c, 0x71, 0xda,
	0x26, 0x5c, 0x11, 0x1c, 0x7f, 0xec, 0xc0, 0xc9, 0xd6, 0xee, 0xff, 0x54, 0x5b, 0x5a, 0xfe, 0x11,
	0xf4, 0x52, 0x61, 0x95, 0x2e, 0x9a, 0xf5, 0x6f, 0x10, 0x89, 0x95, 0x68, 0x13, 0xa1, 0xa5, 0x7f,
	0x04, 0x41, 0xdc, 0x2f, 0xd1, 0x5e, 0x68, 0x89, 0x8c, 0x43, 0x3f, 0xc7, 0xaa, 0xa2, 0xbb, 0xea,
	0xba, 0x9c, 0x16, 0xd2, 0x09, 0xe9, 0xf5, 0x24, 0x4a, 0xba, 0x97, 0x30, 0x88, 0x7b, 0x04, 0x67,
	0x92, 0x08, 0x7a, 0x42, 0x44, 0x04, 0x9e, 0x20, 0x38, 0x93, 0xe3, 0xef, 0xe0, 0x24, 0xc6, 0x95,
	0xfe, 0x0d, 0xff, 0xe7, 0xa3, 0x1c, 0x23, 0x8c, 0xb6, 0x05, 0x1e, 0xe1, 0x64, 0xe3, 0x37, 0xf0,
	0xf4, 0x12, 0xfd, 0xf4, 0xde, 0xe2, 0x7b, 0xeb, 0x7b, 0x7c, 0x0d, 0x07, 0x25, 0x56, 0xba, 0x2e,
	0x05, 0x26, 0xf6, 0xd6, 0x60, 0xc5, 0x3b, 0x67, 0xdd, 0xc9, 0x20, 0xde, 0x6f, 0xbd, 0xbf, 0x90,
	0x73, 0xfc, 0x57, 0x07, 0xd8, 0x66, 0xf2, 0x63, 0x4c, 0xfe, 0x47, 0x38, 0xb8, 0x21, 0xfd, 0x64,
	0x89, 0xef, 0x6d, 0x52, 0xa1, 0xe5, 0x3b, 0x67, 0xdd, 0xc9, 0xde, 0xf9, 0xeb, 0xa8, 0xf9, 0x30,
	0x46, 0xff, 0xee, 0x20, 0x5a, 0xe3, 0x78, 0x78, 0xd3, 0x9a, 0x73, 0xb4, 0xe1, 0x0f, 0x30, 0x58,
	0x53, 0xb4, 0xec, 0xff, 0x38, 0x64, 0xd3, 0xed, 0x70, 0xf3, 0x8c, 0x74, 0x33, 0x54, 0xb8, 0xe9,
	0xd7, 0xd9, 0xe7, 0x7f, 0x3c, 0x81, 0xbd, 0x2b, 0x55, 0x89, 0xb9, 0x6f, 0x80, 0xbd, 0x83, 0xc3,
	0xad, 0x15, 0x64, 0x2f, 0xef, 0xbb, 0x7b, 0xe0, 0xc3, 0x1c, 0x7e, 0xf1, 0x39, 0xba, 0x99, 0xe0,
	0x1c, 0x8e, 0xb6, 0xef, 0x9e, 0xdd, 0xe7, 0x3c, 0xb8, 0x57, 0xe1, 0xab, 0xcf, 0xf2, 0x8d, 0xe8,
	0x25, 0x0c, 0x37, 0x47, 0xc5, 0xc2, 0x07, 0x27, 0xe8, 0xc5, 0x4e, 0xff, 0x63, 0xba, 0xe1, 0xb7,
	0x7f, 0x7e, 0xfc, 0xf4, 0xf6, 0x1c, 0xbe, 0x09, 0xa3, 0xcc, 0x5a, 0x53, 0xbd, 0x99, 0x4e, 0xa5,
	0x16, 0x55, 0x74, 0xa3, 0x8a, 0x6b, 0xb1, 0xd4, 0xb5, 0x8c, 0x84, 0xce, 0xa7, 0xa9, 0x51, 0x53,
	0xfa, 0x8d, 0x4d, 0x55, 0x21, 0xf1, 0x43, 0x94, 0xd9, 0x7c, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x61, 0xf1, 0x57, 0xd8, 0x06, 0x00, 0x00,
}
