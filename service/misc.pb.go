// Code generated by protoc-gen-go. DO NOT EDIT.
// source: misc.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

var _ = config.Config{}
var _ = request.Request{}
var _ = data.Operation{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MiscServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *MiscServiceProperties) Reset()                    { *m = MiscServiceProperties{} }
func (m *MiscServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*MiscServiceProperties) ProtoMessage()               {}
func (*MiscServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *MiscServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type GrantQuotaIndepInput struct {
	User               string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Zone               string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	Instance           int32  `protobuf:"varint,3,opt,name=instance" json:"instance,omitempty"`
	Cpu                int32  `protobuf:"varint,4,opt,name=cpu" json:"cpu,omitempty"`
	Memory             int32  `protobuf:"varint,5,opt,name=memory" json:"memory,omitempty"`
	HpInstance         int32  `protobuf:"varint,6,opt,name=hp_instance,json=hpInstance" json:"hp_instance,omitempty"`
	HpCpu              int32  `protobuf:"varint,7,opt,name=hp_cpu,json=hpCpu" json:"hp_cpu,omitempty"`
	HpMemory           int32  `protobuf:"varint,8,opt,name=hp_memory,json=hpMemory" json:"hp_memory,omitempty"`
	Volume             int32  `protobuf:"varint,9,opt,name=volume" json:"volume,omitempty"`
	VolumeSize         int32  `protobuf:"varint,10,opt,name=volume_size,json=volumeSize" json:"volume_size,omitempty"`
	HcVolume           int32  `protobuf:"varint,11,opt,name=hc_volume,json=hcVolume" json:"hc_volume,omitempty"`
	HcVolumeSize       int32  `protobuf:"varint,12,opt,name=hc_volume_size,json=hcVolumeSize" json:"hc_volume_size,omitempty"`
	HppVolume          int32  `protobuf:"varint,13,opt,name=hpp_volume,json=hppVolume" json:"hpp_volume,omitempty"`
	HppVolumeSize      int32  `protobuf:"varint,14,opt,name=hpp_volume_size,json=hppVolumeSize" json:"hpp_volume_size,omitempty"`
	Image              int32  `protobuf:"varint,15,opt,name=image" json:"image,omitempty"`
	Loadbalancer       int32  `protobuf:"varint,16,opt,name=loadbalancer" json:"loadbalancer,omitempty"`
	LoadbalancerPolicy int32  `protobuf:"varint,17,opt,name=loadbalancer_policy,json=loadbalancerPolicy" json:"loadbalancer_policy,omitempty"`
	Vxnet              int32  `protobuf:"varint,18,opt,name=vxnet" json:"vxnet,omitempty"`
	Router             int32  `protobuf:"varint,19,opt,name=router" json:"router,omitempty"`
	Eip                int32  `protobuf:"varint,20,opt,name=eip" json:"eip,omitempty"`
	EipBandwidth       int32  `protobuf:"varint,21,opt,name=eip_bandwidth,json=eipBandwidth" json:"eip_bandwidth,omitempty"`
	Rdb                int32  `protobuf:"varint,22,opt,name=rdb" json:"rdb,omitempty"`
	HppRdb             int32  `protobuf:"varint,23,opt,name=hpp_rdb,json=hppRdb" json:"hpp_rdb,omitempty"`
	Cache              int32  `protobuf:"varint,24,opt,name=cache" json:"cache,omitempty"`
	HpCache            int32  `protobuf:"varint,25,opt,name=hp_cache,json=hpCache" json:"hp_cache,omitempty"`
	Mongo              int32  `protobuf:"varint,26,opt,name=mongo" json:"mongo,omitempty"`
	HpMongo            int32  `protobuf:"varint,27,opt,name=hp_mongo,json=hpMongo" json:"hp_mongo,omitempty"`
}

func (m *GrantQuotaIndepInput) Reset()                    { *m = GrantQuotaIndepInput{} }
func (m *GrantQuotaIndepInput) String() string            { return proto.CompactTextString(m) }
func (*GrantQuotaIndepInput) ProtoMessage()               {}
func (*GrantQuotaIndepInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *GrantQuotaIndepInput) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GrantQuotaIndepInput) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *GrantQuotaIndepInput) GetInstance() int32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpInstance() int32 {
	if m != nil {
		return m.HpInstance
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpCpu() int32 {
	if m != nil {
		return m.HpCpu
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpMemory() int32 {
	if m != nil {
		return m.HpMemory
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVolumeSize() int32 {
	if m != nil {
		return m.VolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHcVolume() int32 {
	if m != nil {
		return m.HcVolume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHcVolumeSize() int32 {
	if m != nil {
		return m.HcVolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppVolume() int32 {
	if m != nil {
		return m.HppVolume
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppVolumeSize() int32 {
	if m != nil {
		return m.HppVolumeSize
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetImage() int32 {
	if m != nil {
		return m.Image
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetLoadbalancer() int32 {
	if m != nil {
		return m.Loadbalancer
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetLoadbalancerPolicy() int32 {
	if m != nil {
		return m.LoadbalancerPolicy
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetVxnet() int32 {
	if m != nil {
		return m.Vxnet
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetRouter() int32 {
	if m != nil {
		return m.Router
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetEip() int32 {
	if m != nil {
		return m.Eip
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetEipBandwidth() int32 {
	if m != nil {
		return m.EipBandwidth
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetRdb() int32 {
	if m != nil {
		return m.Rdb
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHppRdb() int32 {
	if m != nil {
		return m.HppRdb
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetCache() int32 {
	if m != nil {
		return m.Cache
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpCache() int32 {
	if m != nil {
		return m.HpCache
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetMongo() int32 {
	if m != nil {
		return m.Mongo
	}
	return 0
}

func (m *GrantQuotaIndepInput) GetHpMongo() int32 {
	if m != nil {
		return m.HpMongo
	}
	return 0
}

type GrantQuotaIndepOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	UserId  string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ZoneId  string `protobuf:"bytes,5,opt,name=zone_id,json=zoneId" json:"zone_id,omitempty"`
}

func (m *GrantQuotaIndepOutput) Reset()                    { *m = GrantQuotaIndepOutput{} }
func (m *GrantQuotaIndepOutput) String() string            { return proto.CompactTextString(m) }
func (*GrantQuotaIndepOutput) ProtoMessage()               {}
func (*GrantQuotaIndepOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *GrantQuotaIndepOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GrantQuotaIndepOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GrantQuotaIndepOutput) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

type RevokeQuotaIndepInput struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Zone string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
}

func (m *RevokeQuotaIndepInput) Reset()                    { *m = RevokeQuotaIndepInput{} }
func (m *RevokeQuotaIndepInput) String() string            { return proto.CompactTextString(m) }
func (*RevokeQuotaIndepInput) ProtoMessage()               {}
func (*RevokeQuotaIndepInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *RevokeQuotaIndepInput) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *RevokeQuotaIndepInput) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type RevokeQuotaIndepOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *RevokeQuotaIndepOutput) Reset()                    { *m = RevokeQuotaIndepOutput{} }
func (m *RevokeQuotaIndepOutput) String() string            { return proto.CompactTextString(m) }
func (*RevokeQuotaIndepOutput) ProtoMessage()               {}
func (*RevokeQuotaIndepOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *RevokeQuotaIndepOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RevokeQuotaIndepOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RevokeQuotaIndepOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetQuotaLeftInput struct {
	ResourceTypes []string `protobuf:"bytes,1,rep,name=resource_types,json=resourceTypes" json:"resource_types,omitempty"`
}

func (m *GetQuotaLeftInput) Reset()                    { *m = GetQuotaLeftInput{} }
func (m *GetQuotaLeftInput) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaLeftInput) ProtoMessage()               {}
func (*GetQuotaLeftInput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *GetQuotaLeftInput) GetResourceTypes() []string {
	if m != nil {
		return m.ResourceTypes
	}
	return nil
}

type GetQuotaLeftOutput struct {
	Action       string                          `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode      int32                           `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message      string                          `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	QuotaLeftSet []*GetQuotaLeftOutput_QuotaLeft `protobuf:"bytes,4,rep,name=quota_left_set,json=quotaLeftSet" json:"quota_left_set,omitempty"`
}

func (m *GetQuotaLeftOutput) Reset()                    { *m = GetQuotaLeftOutput{} }
func (m *GetQuotaLeftOutput) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaLeftOutput) ProtoMessage()               {}
func (*GetQuotaLeftOutput) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *GetQuotaLeftOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetQuotaLeftOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetQuotaLeftOutput) GetMessage() string {
	if m != nil {
		return m.Message
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
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType" json:"resource_type,omitempty"`
	Left         int32  `protobuf:"varint,2,opt,name=left" json:"left,omitempty"`
}

func (m *GetQuotaLeftOutput_QuotaLeft) Reset()         { *m = GetQuotaLeftOutput_QuotaLeft{} }
func (m *GetQuotaLeftOutput_QuotaLeft) String() string { return proto.CompactTextString(m) }
func (*GetQuotaLeftOutput_QuotaLeft) ProtoMessage()    {}
func (*GetQuotaLeftOutput_QuotaLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{6, 0}
}

func (m *GetQuotaLeftOutput_QuotaLeft) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *GetQuotaLeftOutput_QuotaLeft) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func init() {
	proto.RegisterType((*MiscServiceProperties)(nil), "service.MiscServiceProperties")
	proto.RegisterType((*GrantQuotaIndepInput)(nil), "service.GrantQuotaIndepInput")
	proto.RegisterType((*GrantQuotaIndepOutput)(nil), "service.GrantQuotaIndepOutput")
	proto.RegisterType((*RevokeQuotaIndepInput)(nil), "service.RevokeQuotaIndepInput")
	proto.RegisterType((*RevokeQuotaIndepOutput)(nil), "service.RevokeQuotaIndepOutput")
	proto.RegisterType((*GetQuotaLeftInput)(nil), "service.GetQuotaLeftInput")
	proto.RegisterType((*GetQuotaLeftOutput)(nil), "service.GetQuotaLeftOutput")
	proto.RegisterType((*GetQuotaLeftOutput_QuotaLeft)(nil), "service.GetQuotaLeftOutput.QuotaLeft")
}

// See https://docs.qingcloud.com/api/misc/index.html
type MiscServiceInterface interface {
	GrantQuotaIndep(in *GrantQuotaIndepInput) (out *GrantQuotaIndepOutput, err error)
	RevokeQuotaIndep(in *RevokeQuotaIndepInput) (out *RevokeQuotaIndepOutput, err error)
	GetQuotaLeft(in *GetQuotaLeftInput) (out *GetQuotaLeftOutput, err error)
}

// See https://docs.qingcloud.com/api/misc/index.html
type MiscService struct {
	Config           *config.Config
	Properties       *MiscServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/misc/index.html
func NewMiscService(conf *config.Config, zone string) (p *MiscService) {
	return &MiscService{
		Config:     conf,
		Properties: &MiscServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/misc/index.html
func (s *QingCloudService) Misc(zone string) (*MiscService, error) {
	properties := &MiscServiceProperties{
		Zone: zone,
	}

	return &MiscService{Config: s.Config, Properties: properties}, nil
}

func (p *MiscService) GrantQuotaIndep(in *GrantQuotaIndepInput) (out *GrantQuotaIndepOutput, err error) {
	if in == nil {
		in = &GrantQuotaIndepInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GrantQuotaIndep",
		RequestMethod: "GET", // GET or POST
	}

	x := &GrantQuotaIndepOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MiscService) RevokeQuotaIndep(in *RevokeQuotaIndepInput) (out *RevokeQuotaIndepOutput, err error) {
	if in == nil {
		in = &RevokeQuotaIndepInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeQuotaIndep",
		RequestMethod: "GET", // GET or POST
	}

	x := &RevokeQuotaIndepOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MiscService) GetQuotaLeft(in *GetQuotaLeftInput) (out *GetQuotaLeftOutput, err error) {
	if in == nil {
		in = &GetQuotaLeftInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetQuotaLeft",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetQuotaLeftOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MiscServiceProperties) Validate() error {
	return nil // TODO
}

func (p *GrantQuotaIndepInput) Validate() error {
	return nil // TODO
}

func (p *GrantQuotaIndepOutput) Validate() error {
	return nil // TODO
}

func (p *RevokeQuotaIndepInput) Validate() error {
	return nil // TODO
}

func (p *RevokeQuotaIndepOutput) Validate() error {
	return nil // TODO
}

func (p *GetQuotaLeftInput) Validate() error {
	return nil // TODO
}

func (p *GetQuotaLeftOutput) Validate() error {
	return nil // TODO
}

func init() { proto.RegisterFile("misc.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xc1, 0x6e, 0x1b, 0x37,
	0x10, 0x85, 0x22, 0x4b, 0xb2, 0xc6, 0xb2, 0xec, 0x30, 0x96, 0xc2, 0xc8, 0x70, 0x63, 0xa8, 0x4d,
	0x61, 0xa0, 0x80, 0x04, 0xa4, 0xb7, 0x5c, 0x5a, 0xd4, 0x01, 0x0c, 0xa1, 0x35, 0xea, 0xc8, 0x45,
	0xaf, 0x8b, 0x15, 0x39, 0xf1, 0x12, 0xd9, 0x5d, 0xd2, 0x5c, 0xae, 0x1b, 0xfb, 0x13, 0xfa, 0x01,
	0xfd, 0xa1, 0x7e, 0x51, 0x81, 0x5e, 0x7a, 0x2b, 0x86, 0xdc, 0x5d, 0xab, 0xaa, 0xd3, 0x43, 0x01,
	0x5f, 0x16, 0x7c, 0xf3, 0x66, 0xde, 0x0c, 0x87, 0xc3, 0x25, 0x40, 0xa6, 0x0a, 0x31, 0x33, 0x56,
	0x3b, 0xcd, 0x7a, 0x05, 0xda, 0x1b, 0x25, 0x70, 0x72, 0x74, 0xad, 0xf2, 0x2b, 0x91, 0xea, 0x52,
	0x46, 0x85, 0xfc, 0x10, 0xd9, 0x32, 0xc5, 0x39, 0x7d, 0x82, 0xdf, 0xf4, 0x2b, 0x18, 0x9d, 0xab,
	0x42, 0x5c, 0x06, 0xef, 0x0b, 0xab, 0x0d, 0x5a, 0xa7, 0xb0, 0x60, 0x0c, 0xb6, 0xee, 0x74, 0x8e,
	0xbc, 0x75, 0xdc, 0x3a, 0xe9, 0x2f, 0xfd, 0x7a, 0xfa, 0x57, 0x07, 0x0e, 0xce, 0x6c, 0x9c, 0xbb,
	0x77, 0xa5, 0x76, 0xf1, 0x22, 0x97, 0x68, 0x16, 0xb9, 0x29, 0x1d, 0x39, 0x97, 0x05, 0xda, 0xda,
	0x99, 0xd6, 0x8d, 0xc0, 0x93, 0x7b, 0x01, 0x36, 0x81, 0x6d, 0x95, 0x17, 0x2e, 0xce, 0x05, 0xf2,
	0xf6, 0x71, 0xeb, 0xa4, 0xb3, 0x6c, 0x30, 0xdb, 0x87, 0xb6, 0x30, 0x25, 0xdf, 0xf2, 0x66, 0x5a,
	0xb2, 0x31, 0x74, 0x33, 0xcc, 0xb4, 0xbd, 0xe5, 0x1d, 0x6f, 0xac, 0x10, 0x7b, 0x09, 0x3b, 0x89,
	0x89, 0x1a, 0xa1, 0xae, 0x27, 0x21, 0x31, 0x8b, 0x5a, 0x6a, 0x04, 0xdd, 0xc4, 0x44, 0xa4, 0xd6,
	0xf3, 0x5c, 0x27, 0x31, 0xa7, 0xa6, 0x64, 0x87, 0xd0, 0x4f, 0x4c, 0x54, 0x49, 0x6e, 0x87, 0xf4,
	0x89, 0x39, 0x0f, 0xa2, 0x63, 0xe8, 0xde, 0xe8, 0xb4, 0xcc, 0x90, 0xf7, 0x43, 0xb2, 0x80, 0x28,
	0x59, 0x58, 0x45, 0x85, 0xba, 0x43, 0x0e, 0x21, 0x59, 0x30, 0x5d, 0xaa, 0x3b, 0xf4, 0xaa, 0x22,
	0xaa, 0x62, 0x77, 0x2a, 0x55, 0xf1, 0x73, 0x88, 0xfe, 0x02, 0x86, 0x0d, 0x19, 0x04, 0x06, 0xde,
	0x63, 0x50, 0x7b, 0x78, 0x89, 0x23, 0x80, 0xc4, 0x98, 0x5a, 0x63, 0xd7, 0x7b, 0xf4, 0x13, 0x63,
	0x2a, 0x91, 0x2f, 0x61, 0xef, 0x9e, 0x0e, 0x2a, 0x43, 0xef, 0xb3, 0xdb, 0xf8, 0x78, 0x99, 0x03,
	0xe8, 0xa8, 0x2c, 0xbe, 0x42, 0xbe, 0x17, 0x76, 0xed, 0x01, 0x9b, 0xc2, 0x20, 0xd5, 0xb1, 0x5c,
	0xc5, 0x29, 0xf5, 0xc6, 0xf2, 0xfd, 0x50, 0xc0, 0xba, 0x8d, 0xcd, 0xe1, 0xd9, 0x3a, 0x8e, 0x8c,
	0x4e, 0x95, 0xb8, 0xe5, 0x4f, 0xbd, 0x2b, 0x5b, 0xa7, 0x2e, 0x3c, 0x43, 0xa9, 0x6e, 0x3e, 0xe6,
	0xe8, 0x38, 0x0b, 0xa9, 0x3c, 0xa0, 0x1e, 0x5a, 0x5d, 0x3a, 0xb4, 0xfc, 0x59, 0xe8, 0x61, 0x40,
	0x74, 0xb4, 0xa8, 0x0c, 0x3f, 0x08, 0x47, 0x8b, 0xca, 0xb0, 0xcf, 0x61, 0x17, 0x95, 0x89, 0x56,
	0x71, 0x2e, 0x7f, 0x51, 0xd2, 0x25, 0x7c, 0x14, 0xaa, 0x42, 0x65, 0xbe, 0xab, 0x6d, 0x14, 0x66,
	0xe5, 0x8a, 0x8f, 0x43, 0x98, 0x95, 0x2b, 0xf6, 0x1c, 0x7a, 0xd4, 0x09, 0xb2, 0x3e, 0x0f, 0x19,
	0x12, 0x63, 0x96, 0x72, 0x45, 0xf5, 0x88, 0x58, 0x24, 0xc8, 0x79, 0xa8, 0xc7, 0x03, 0xf6, 0x02,
	0xb6, 0x69, 0x0e, 0x3c, 0xf1, 0xc2, 0x13, 0xbd, 0xc4, 0x9c, 0x7a, 0xea, 0x00, 0x3a, 0x99, 0xce,
	0xaf, 0x34, 0x9f, 0x84, 0x00, 0x0f, 0xaa, 0x80, 0x40, 0x1c, 0xd6, 0x01, 0xe7, 0x04, 0xa7, 0xbf,
	0xb5, 0x60, 0xb4, 0x31, 0xfb, 0x3f, 0x96, 0x8e, 0x86, 0x7f, 0x0c, 0xdd, 0x58, 0x38, 0xa5, 0xf3,
	0x6a, 0xfc, 0x2b, 0x44, 0x62, 0x16, 0x5d, 0x24, 0xb4, 0x0c, 0x97, 0xa0, 0xb3, 0xec, 0x59, 0x74,
	0xa7, 0x5a, 0x22, 0xe3, 0xd0, 0xcb, 0xb0, 0x28, 0xe8, 0xac, 0xda, 0x3e, 0xa6, 0x86, 0xb4, 0x43,
	0xba, 0x3d, 0x91, 0x92, 0xfe, 0x26, 0xf4, 0x97, 0x5d, 0x82, 0x0b, 0x49, 0x04, 0x5d, 0x21, 0x22,
	0x3a, 0x81, 0x20, 0xb8, 0x90, 0xd3, 0x6f, 0x60, 0xb4, 0xc4, 0x1b, 0xfd, 0x01, 0xff, 0xe7, 0xa5,
	0x9c, 0x22, 0x8c, 0x37, 0x05, 0x1e, 0x61, 0x67, 0xd3, 0x37, 0xf0, 0xf4, 0x0c, 0x43, 0xf7, 0x7e,
	0xc0, 0xf7, 0x2e, 0xd4, 0xf8, 0x0a, 0x86, 0x16, 0x0b, 0x5d, 0x5a, 0x81, 0x91, 0xbb, 0x35, 0x58,
	0xf0, 0xd6, 0x71, 0xfb, 0xa4, 0xbf, 0xdc, 0xad, 0xad, 0x3f, 0x91, 0x71, 0xfa, 0x67, 0x0b, 0xd8,
	0x7a, 0xf0, 0x63, 0x74, 0xfe, 0x7b, 0x18, 0x5e, 0x93, 0x7e, 0x94, 0xe2, 0x7b, 0x17, 0x15, 0xe8,
	0xf8, 0xd6, 0x71, 0xfb, 0x64, 0xe7, 0xf5, 0xab, 0x59, 0xf5, 0x2b, 0x9d, 0xfd, 0xbb, 0x82, 0x59,
	0x83, 0x97, 0x83, 0xeb, 0x7a, 0x79, 0x89, 0x6e, 0xf2, 0x16, 0xfa, 0x0d, 0x45, 0xc3, 0xfe, 0x8f,
	0x4d, 0x56, 0xd5, 0x0e, 0xd6, 0xf7, 0x48, 0x27, 0x43, 0x89, 0xab, 0x7a, 0xfd, 0xfa, 0xf5, 0xef,
	0x4f, 0x60, 0x67, 0xed, 0xef, 0xcc, 0x2e, 0x60, 0x6f, 0x63, 0x04, 0xd9, 0xd1, 0x7d, 0x75, 0x0f,
	0xfc, 0x98, 0x27, 0x9f, 0x7d, 0x8a, 0xae, 0x3a, 0x78, 0x09, 0xfb, 0x9b, 0x67, 0xcf, 0xee, 0x63,
	0x1e, 0x9c, 0xab, 0xc9, 0xcb, 0x4f, 0xf2, 0x95, 0xe8, 0x19, 0x0c, 0xd6, 0x5b, 0xc5, 0x26, 0x0f,
	0x76, 0x30, 0x88, 0x1d, 0xfe, 0x47, 0x77, 0x27, 0x6f, 0x7f, 0xfd, 0x63, 0xeb, 0x5b, 0x98, 0x25,
	0xce, 0x99, 0xe2, 0xcd, 0x7c, 0x2e, 0xb5, 0x28, 0x66, 0xcd, 0x73, 0x36, 0x13, 0x3a, 0x9b, 0xc7,
	0x46, 0xcd, 0xe9, 0xd5, 0x9b, 0xab, 0x5c, 0xe2, 0xc7, 0x59, 0xe2, 0xb2, 0x94, 0x0d, 0xdf, 0xa9,
	0xfc, 0xea, 0xd4, 0xbb, 0x50, 0xf3, 0x56, 0x5d, 0xff, 0xd2, 0x7d, 0xfd, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xce, 0xc3, 0x8c, 0xab, 0x1f, 0x07, 0x00, 0x00,
}
