// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type MonitorServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *MonitorServiceProperties) Reset()                    { *m = MonitorServiceProperties{} }
func (m *MonitorServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*MonitorServiceProperties) ProtoMessage()               {}
func (*MonitorServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *MonitorServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type GetMonitorInput struct {
	Resource  string                      `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Meters    []string                    `protobuf:"bytes,2,rep,name=meters" json:"meters,omitempty"`
	Step      string                      `protobuf:"bytes,3,opt,name=step" json:"step,omitempty"`
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime   *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *GetMonitorInput) Reset()                    { *m = GetMonitorInput{} }
func (m *GetMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetMonitorInput) ProtoMessage()               {}
func (*GetMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *GetMonitorInput) GetResource() string {
	if m != nil {
		return m.Resource
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
	if m != nil {
		return m.Step
	}
	return ""
}

func (m *GetMonitorInput) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetMonitorInput) GetEndTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetMonitorOutput struct {
	Action     string                           `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                            `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                           `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ResourceId string                           `protobuf:"bytes,4,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	MeterSet   []*GetMonitorOutput_ResponseItem `protobuf:"bytes,5,rep,name=meter_set,json=meterSet" json:"meter_set,omitempty"`
}

func (m *GetMonitorOutput) Reset()                    { *m = GetMonitorOutput{} }
func (m *GetMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetMonitorOutput) ProtoMessage()               {}
func (*GetMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *GetMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetMonitorOutput) GetResourceId() string {
	if m != nil {
		return m.ResourceId
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
	VxnetId  string `protobuf:"bytes,2,opt,name=vxnet_id,json=vxnetId" json:"vxnet_id,omitempty"`
	MeterId  string `protobuf:"bytes,3,opt,name=meter_id,json=meterId" json:"meter_id,omitempty"`
	Sequence int32  `protobuf:"varint,4,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *GetMonitorOutput_ResponseItem) Reset()         { *m = GetMonitorOutput_ResponseItem{} }
func (m *GetMonitorOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*GetMonitorOutput_ResponseItem) ProtoMessage()    {}
func (*GetMonitorOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{2, 0}
}

func (m *GetMonitorOutput_ResponseItem) GetVxnetId() string {
	if m != nil {
		return m.VxnetId
	}
	return ""
}

func (m *GetMonitorOutput_ResponseItem) GetMeterId() string {
	if m != nil {
		return m.MeterId
	}
	return ""
}

func (m *GetMonitorOutput_ResponseItem) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type GetLoadBalancerMonitorInput struct {
	Resource  string                      `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Meters    []string                    `protobuf:"bytes,2,rep,name=meters" json:"meters,omitempty"`
	Step      string                      `protobuf:"bytes,3,opt,name=step" json:"step,omitempty"`
	StartTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime   *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *GetLoadBalancerMonitorInput) Reset()                    { *m = GetLoadBalancerMonitorInput{} }
func (m *GetLoadBalancerMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorInput) ProtoMessage()               {}
func (*GetLoadBalancerMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *GetLoadBalancerMonitorInput) GetResource() string {
	if m != nil {
		return m.Resource
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
	if m != nil {
		return m.Step
	}
	return ""
}

func (m *GetLoadBalancerMonitorInput) GetStartTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *GetLoadBalancerMonitorInput) GetEndTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetLoadBalancerMonitorOutput struct {
	Action     string                                       `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode    int32                                        `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message    string                                       `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	ResourceId string                                       `protobuf:"bytes,4,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	MeterSet   []*GetLoadBalancerMonitorOutput_ResponseItem `protobuf:"bytes,5,rep,name=meter_set,json=meterSet" json:"meter_set,omitempty"`
}

func (m *GetLoadBalancerMonitorOutput) Reset()                    { *m = GetLoadBalancerMonitorOutput{} }
func (m *GetLoadBalancerMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorOutput) ProtoMessage()               {}
func (*GetLoadBalancerMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *GetLoadBalancerMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetLoadBalancerMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput) GetResourceId() string {
	if m != nil {
		return m.ResourceId
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
	EipId string `protobuf:"bytes,1,opt,name=eip_id,json=eipId" json:"eip_id,omitempty"`
	// data: ...
	MeterId string `protobuf:"bytes,3,opt,name=meter_id,json=meterId" json:"meter_id,omitempty"`
}

func (m *GetLoadBalancerMonitorOutput_ResponseItem) Reset() {
	*m = GetLoadBalancerMonitorOutput_ResponseItem{}
}
func (m *GetLoadBalancerMonitorOutput_ResponseItem) String() string { return proto.CompactTextString(m) }
func (*GetLoadBalancerMonitorOutput_ResponseItem) ProtoMessage()    {}
func (*GetLoadBalancerMonitorOutput_ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor13, []int{4, 0}
}

func (m *GetLoadBalancerMonitorOutput_ResponseItem) GetEipId() string {
	if m != nil {
		return m.EipId
	}
	return ""
}

func (m *GetLoadBalancerMonitorOutput_ResponseItem) GetMeterId() string {
	if m != nil {
		return m.MeterId
	}
	return ""
}

type GetRDBMonitorInput struct {
}

func (m *GetRDBMonitorInput) Reset()                    { *m = GetRDBMonitorInput{} }
func (m *GetRDBMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetRDBMonitorInput) ProtoMessage()               {}
func (*GetRDBMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

type GetRDBMonitorOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GetRDBMonitorOutput) Reset()                    { *m = GetRDBMonitorOutput{} }
func (m *GetRDBMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetRDBMonitorOutput) ProtoMessage()               {}
func (*GetRDBMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *GetRDBMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetRDBMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetRDBMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetCacheMonitorInput struct {
}

func (m *GetCacheMonitorInput) Reset()                    { *m = GetCacheMonitorInput{} }
func (m *GetCacheMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetCacheMonitorInput) ProtoMessage()               {}
func (*GetCacheMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

type GetCacheMonitorOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GetCacheMonitorOutput) Reset()                    { *m = GetCacheMonitorOutput{} }
func (m *GetCacheMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetCacheMonitorOutput) ProtoMessage()               {}
func (*GetCacheMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *GetCacheMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetCacheMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetCacheMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetZooKeeperMonitorInput struct {
}

func (m *GetZooKeeperMonitorInput) Reset()                    { *m = GetZooKeeperMonitorInput{} }
func (m *GetZooKeeperMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetZooKeeperMonitorInput) ProtoMessage()               {}
func (*GetZooKeeperMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

type GetZooKeeperMonitorOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GetZooKeeperMonitorOutput) Reset()                    { *m = GetZooKeeperMonitorOutput{} }
func (m *GetZooKeeperMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetZooKeeperMonitorOutput) ProtoMessage()               {}
func (*GetZooKeeperMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *GetZooKeeperMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetZooKeeperMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetZooKeeperMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetQueueMonitorInput struct {
}

func (m *GetQueueMonitorInput) Reset()                    { *m = GetQueueMonitorInput{} }
func (m *GetQueueMonitorInput) String() string            { return proto.CompactTextString(m) }
func (*GetQueueMonitorInput) ProtoMessage()               {}
func (*GetQueueMonitorInput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{11} }

type GetQueueMonitorOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *GetQueueMonitorOutput) Reset()                    { *m = GetQueueMonitorOutput{} }
func (m *GetQueueMonitorOutput) String() string            { return proto.CompactTextString(m) }
func (*GetQueueMonitorOutput) ProtoMessage()               {}
func (*GetQueueMonitorOutput) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{12} }

func (m *GetQueueMonitorOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *GetQueueMonitorOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *GetQueueMonitorOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MonitorServiceProperties)(nil), "service.MonitorServiceProperties")
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

// See https://docs.qingcloud.com/api/monitor/index.html
type MonitorServiceInterface interface {
	GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error)
	GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error)
	GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error)
	GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error)
	GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error)
	GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error)
}

// See https://docs.qingcloud.com/api/monitor/index.html
type MonitorService struct {
	Config           *config.Config
	Properties       *MonitorServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/monitor/index.html
func NewMonitorService(conf *config.Config, zone string) (p *MonitorService) {
	return &MonitorService{
		Config:     conf,
		Properties: &MonitorServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/monitor/index.html
func (s *QingCloudService) Monitor(zone string) (*MonitorService, error) {
	properties := &MonitorServiceProperties{
		Zone: zone,
	}

	return &MonitorService{Config: s.Config, Properties: properties}, nil
}

func (p *MonitorService) GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error) {
	if in == nil {
		in = &GetMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetMonitorOutput{}
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

func (p *MonitorService) GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error) {
	if in == nil {
		in = &GetLoadBalancerMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetLoadBalancerMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetLoadBalancerMonitorOutput{}
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

func (p *MonitorService) GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error) {
	if in == nil {
		in = &GetRDBMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetRDBMonitorOutput{}
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

func (p *MonitorService) GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error) {
	if in == nil {
		in = &GetCacheMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetCacheMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetCacheMonitorOutput{}
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

func (p *MonitorService) GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error) {
	if in == nil {
		in = &GetZooKeeperMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetZooKeeperMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetZooKeeperMonitorOutput{}
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

func (p *MonitorService) GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error) {
	if in == nil {
		in = &GetQueueMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetQueueMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetQueueMonitorOutput{}
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

func (p *MonitorServiceProperties) Validate() error {
	return nil // TODO
}

func (p *GetMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetMonitorOutput) Validate() error {
	return nil // TODO
}

func (p *GetLoadBalancerMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetLoadBalancerMonitorOutput) Validate() error {
	return nil // TODO
}

func (p *GetRDBMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetRDBMonitorOutput) Validate() error {
	return nil // TODO
}

func (p *GetCacheMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetCacheMonitorOutput) Validate() error {
	return nil // TODO
}

func (p *GetZooKeeperMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetZooKeeperMonitorOutput) Validate() error {
	return nil // TODO
}

func (p *GetQueueMonitorInput) Validate() error {
	return nil // TODO
}

func (p *GetQueueMonitorOutput) Validate() error {
	return nil // TODO
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x56, 0xd2, 0x5c, 0x9a, 0xd3, 0xbf, 0x3f, 0x30, 0xb4, 0x91, 0xeb, 0xb6, 0x34, 0x58, 0x80,
	0xb2, 0xb2, 0x45, 0x10, 0x0b, 0x58, 0x41, 0x83, 0x88, 0x52, 0x40, 0x6d, 0x5d, 0x56, 0x08, 0x29,
	0xb8, 0x9e, 0x43, 0x62, 0x11, 0x7b, 0x5c, 0xcf, 0x71, 0x55, 0xf1, 0x08, 0xbc, 0x04, 0x4b, 0x1e,
	0x86, 0x2d, 0x12, 0xcf, 0xc1, 0x1b, 0x20, 0x5f, 0xd2, 0xd8, 0x89, 0xd3, 0xb2, 0x89, 0x90, 0xd8,
	0x58, 0x3e, 0x73, 0x6e, 0xdf, 0x77, 0xe6, 0xf3, 0x78, 0x60, 0xdd, 0x15, 0x9e, 0x43, 0x22, 0xd0,
	0xfd, 0x40, 0x90, 0x60, 0x75, 0x89, 0xc1, 0xb9, 0x63, 0xa3, 0xba, 0x7b, 0xe6, 0x78, 0x43, 0x7b,
	0x2c, 0x42, 0x3e, 0x90, 0xfc, 0xd3, 0x20, 0x08, 0xc7, 0x68, 0x44, 0x8f, 0x24, 0x4e, 0xdd, 0x1b,
	0x0a, 0x31, 0x1c, 0xa3, 0x11, 0x5b, 0xa7, 0xe1, 0x47, 0x83, 0x1c, 0x17, 0x25, 0x59, 0xae, 0x9f,
	0x04, 0x68, 0x3a, 0x28, 0x6f, 0x92, 0xca, 0x27, 0x49, 0xc5, 0xa3, 0x40, 0xf8, 0x18, 0x90, 0x83,
	0x92, 0x31, 0xa8, 0x7c, 0x16, 0x1e, 0x2a, 0xa5, 0x56, 0xa9, 0xdd, 0x30, 0xe3, 0x77, 0xed, 0x7b,
	0x09, 0x6e, 0xf4, 0x90, 0xd2, 0x9c, 0xbe, 0xe7, 0x87, 0xc4, 0x54, 0x58, 0x0d, 0x50, 0x8a, 0x30,
	0xb0, 0x27, 0xb1, 0x97, 0x36, 0x6b, 0x42, 0xcd, 0x45, 0xc2, 0x40, 0x2a, 0xe5, 0xd6, 0x4a, 0xbb,
	0x61, 0xa6, 0x56, 0x54, 0x5b, 0x12, 0xfa, 0xca, 0x4a, 0x52, 0x3b, 0x7a, 0x67, 0x4f, 0x00, 0x24,
	0x59, 0x01, 0x0d, 0x22, 0x90, 0x4a, 0xa5, 0x55, 0x6a, 0xaf, 0x75, 0x54, 0x3d, 0x61, 0xa0, 0x4f,
	0x18, 0xe8, 0x6f, 0x27, 0x0c, 0xcc, 0x46, 0x1c, 0x1d, 0xd9, 0xec, 0x31, 0xac, 0xa2, 0xc7, 0x93,
	0xc4, 0xea, 0xb5, 0x89, 0x75, 0xf4, 0x78, 0x64, 0x69, 0xdf, 0xca, 0x70, 0x73, 0xca, 0xe6, 0x30,
	0xa4, 0x88, 0x4e, 0x13, 0x6a, 0x96, 0x4d, 0x8e, 0xf0, 0x52, 0x32, 0xa9, 0xc5, 0xb6, 0x22, 0x9a,
	0x34, 0xb0, 0x05, 0x47, 0xa5, 0xdc, 0x2a, 0xb5, 0xab, 0x66, 0x3d, 0x40, 0xea, 0x0a, 0x8e, 0x4c,
	0x81, 0xba, 0x8b, 0x52, 0x5a, 0x43, 0x4c, 0x09, 0x4d, 0x4c, 0xb6, 0x07, 0x6b, 0x93, 0x59, 0x0c,
	0x1c, 0x1e, 0x93, 0x6a, 0x98, 0x30, 0x59, 0xea, 0x73, 0xd6, 0x85, 0x46, 0x3c, 0x92, 0x81, 0x44,
	0x52, 0xaa, 0xad, 0x95, 0xf6, 0x5a, 0xe7, 0x81, 0x9e, 0xee, 0xae, 0x3e, 0x8b, 0x4d, 0x37, 0x51,
	0xfa, 0xc2, 0x93, 0xd8, 0x27, 0x74, 0xcd, 0xd5, 0x38, 0xf1, 0x04, 0x49, 0xfd, 0x00, 0xff, 0x65,
	0x3d, 0x11, 0xd4, 0xf3, 0x0b, 0x0f, 0x29, 0x6a, 0x59, 0x4e, 0x00, 0xc5, 0x76, 0x9f, 0x47, 0xae,
	0xa4, 0x9f, 0xc3, 0xa7, 0x58, 0x09, 0x83, 0x3e, 0x8f, 0xf6, 0x51, 0xe2, 0x59, 0x88, 0x9e, 0x9d,
	0x4c, 0xbf, 0x6a, 0x5e, 0xda, 0xda, 0xcf, 0x12, 0x6c, 0xf7, 0x90, 0x5e, 0x0b, 0x8b, 0xef, 0x5b,
	0x63, 0xcb, 0xb3, 0x31, 0xf8, 0x47, 0x34, 0xf0, 0xb5, 0x0c, 0x3b, 0xc5, 0xcc, 0xfe, 0x8a, 0x1e,
	0x0e, 0xe7, 0xf5, 0xd0, 0xc9, 0xea, 0x61, 0x21, 0xce, 0x45, 0xda, 0x78, 0x36, 0xa3, 0x8d, 0x4d,
	0xa8, 0xa1, 0xe3, 0x47, 0xcd, 0x13, 0x3a, 0x55, 0x74, 0xfc, 0x2b, 0x75, 0xa1, 0x6d, 0x00, 0xeb,
	0x21, 0x99, 0x2f, 0xf6, 0xb3, 0x3b, 0xae, 0x9d, 0xc2, 0xed, 0xdc, 0xea, 0x12, 0xa6, 0xa5, 0x35,
	0x61, 0xa3, 0x87, 0xd4, 0xb5, 0xec, 0x11, 0xe6, 0x7a, 0x73, 0xd8, 0x9c, 0x59, 0x5f, 0x46, 0x77,
	0x15, 0x94, 0x1e, 0xd2, 0x3b, 0x21, 0x5e, 0x21, 0xfa, 0x79, 0xbd, 0x6b, 0x23, 0xd8, 0x2a, 0xf0,
	0x2d, 0x6f, 0x06, 0xc7, 0x21, 0x86, 0x45, 0x33, 0xc8, 0xae, 0x2f, 0xa1, 0x7b, 0xe7, 0x47, 0x05,
	0xfe, 0xcf, 0xff, 0x20, 0xd8, 0x73, 0x80, 0xe9, 0xb9, 0xc4, 0x94, 0x82, 0xc3, 0x2a, 0x06, 0xa8,
	0x6e, 0x2d, 0x3c, 0xc6, 0x18, 0x42, 0xb3, 0x58, 0xca, 0xec, 0xde, 0x35, 0x5a, 0x4f, 0x4a, 0xdf,
	0xff, 0xa3, 0x2f, 0x82, 0x1d, 0xc0, 0x7a, 0x4e, 0xa2, 0x6c, 0x3b, 0x9b, 0x37, 0x23, 0x68, 0x75,
	0xa7, 0xd8, 0x99, 0xd6, 0x3a, 0x8a, 0xff, 0x7b, 0x59, 0xc9, 0xb1, 0xdd, 0x6c, 0xc2, 0x9c, 0x48,
	0xd5, 0x3b, 0x8b, 0xdc, 0x69, 0xc5, 0xf7, 0xf1, 0x07, 0x34, 0x2b, 0x21, 0x76, 0x37, 0x9b, 0x56,
	0x28, 0x3e, 0x55, 0xbb, 0x2a, 0x24, 0x87, 0x37, 0x2b, 0x8f, 0x3c, 0xde, 0x39, 0x41, 0xe5, 0xf1,
	0xce, 0xeb, 0x4a, 0x3d, 0xf8, 0xf2, 0xab, 0xf2, 0x12, 0x1e, 0x8e, 0x88, 0x7c, 0xf9, 0xd4, 0x30,
	0xb8, 0xb0, 0xa5, 0x7e, 0x79, 0xfb, 0xd0, 0x6d, 0xe1, 0x1a, 0x96, 0xef, 0x18, 0xe9, 0x3d, 0xc5,
	0x70, 0x3c, 0x8e, 0x17, 0xfa, 0x88, 0xdc, 0x31, 0xbb, 0x75, 0xec, 0x78, 0xc3, 0x6e, 0x1c, 0x95,
	0x56, 0x3c, 0xad, 0xc5, 0x27, 0xf2, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x6e, 0xba,
	0x0d, 0xd7, 0x08, 0x00, 0x00,
}
