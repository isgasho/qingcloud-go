// Code generated by protoc-gen-go. DO NOT EDIT.
// source: instances.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type InstanceServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *InstanceServiceProperties) Reset()                    { *m = InstanceServiceProperties{} }
func (m *InstanceServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*InstanceServiceProperties) ProtoMessage()               {}
func (*InstanceServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *InstanceServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeInstanceTypesInput struct {
	Action        string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Zone          string   `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	InstanceTypes []string `protobuf:"bytes,3,rep,name=instance_types,json=instanceTypes" json:"instance_types,omitempty"`
}

func (m *DescribeInstanceTypesInput) Reset()                    { *m = DescribeInstanceTypesInput{} }
func (m *DescribeInstanceTypesInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeInstanceTypesInput) ProtoMessage()               {}
func (*DescribeInstanceTypesInput) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DescribeInstanceTypesInput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeInstanceTypesInput) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *DescribeInstanceTypesInput) GetInstanceTypes() []string {
	if m != nil {
		return m.InstanceTypes
	}
	return nil
}

type DescribeInstanceTypesOutput struct {
	Action          string                                             `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode         int32                                              `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	InstanceTypeSet []*DescribeInstanceTypesOutput_InstanceTypeSetElem `protobuf:"bytes,3,rep,name=instance_type_set,json=instanceTypeSet" json:"instance_type_set,omitempty"`
	TotalCount      int32                                              `protobuf:"varint,4,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *DescribeInstanceTypesOutput) Reset()                    { *m = DescribeInstanceTypesOutput{} }
func (m *DescribeInstanceTypesOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeInstanceTypesOutput) ProtoMessage()               {}
func (*DescribeInstanceTypesOutput) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DescribeInstanceTypesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeInstanceTypesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeInstanceTypesOutput) GetInstanceTypeSet() []*DescribeInstanceTypesOutput_InstanceTypeSetElem {
	if m != nil {
		return m.InstanceTypeSet
	}
	return nil
}

func (m *DescribeInstanceTypesOutput) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type DescribeInstanceTypesOutput_InstanceTypeSetElem struct {
	InstanceTypeId   string `protobuf:"bytes,1,opt,name=instance_type_id,json=instanceTypeId" json:"instance_type_id,omitempty"`
	InstanceTypeName string `protobuf:"bytes,2,opt,name=instance_type_name,json=instanceTypeName" json:"instance_type_name,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	VcpusCurrent     int32  `protobuf:"varint,4,opt,name=vcpus_current,json=vcpusCurrent" json:"vcpus_current,omitempty"`
	MemoryCurrent    int32  `protobuf:"varint,5,opt,name=memory_current,json=memoryCurrent" json:"memory_current,omitempty"`
	Status           string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) Reset() {
	*m = DescribeInstanceTypesOutput_InstanceTypeSetElem{}
}
func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) String() string {
	return proto.CompactTextString(m)
}
func (*DescribeInstanceTypesOutput_InstanceTypeSetElem) ProtoMessage() {}
func (*DescribeInstanceTypesOutput_InstanceTypeSetElem) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{2, 0}
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetInstanceTypeId() string {
	if m != nil {
		return m.InstanceTypeId
	}
	return ""
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetInstanceTypeName() string {
	if m != nil {
		return m.InstanceTypeName
	}
	return ""
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetVcpusCurrent() int32 {
	if m != nil {
		return m.VcpusCurrent
	}
	return 0
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetMemoryCurrent() int32 {
	if m != nil {
		return m.MemoryCurrent
	}
	return 0
}

func (m *DescribeInstanceTypesOutput_InstanceTypeSetElem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*InstanceServiceProperties)(nil), "spec.InstanceServiceProperties")
	proto.RegisterType((*DescribeInstanceTypesInput)(nil), "spec.DescribeInstanceTypesInput")
	proto.RegisterType((*DescribeInstanceTypesOutput)(nil), "spec.DescribeInstanceTypesOutput")
	proto.RegisterType((*DescribeInstanceTypesOutput_InstanceTypeSetElem)(nil), "spec.DescribeInstanceTypesOutput.InstanceTypeSetElem")
}

type InstanceServiceInterface interface {
	DescribeInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RunInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	TerminateInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RestartInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResetInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyInstanceAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeInstanceTypes(in *DescribeInstanceTypesInput) (out *DescribeInstanceTypesOutput, err error)
	CreateBrokers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteBrokers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type InstanceService struct {
	Config     *config.Config
	Properties *InstanceServiceProperties
}

func NewInstanceService(conf *config.Config, zone string) (p *InstanceService, err error) {
	return &InstanceService{
		Config:     conf,
		Properties: &InstanceServiceProperties{Zone: zone},
	}, nil
}

func (p *InstanceService) DescribeInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) RunInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RunInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) TerminateInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "TerminateInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) StartInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) StopInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) RestartInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) ResetInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResetInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) ResizeInstances(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeInstances",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) ModifyInstanceAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyInstanceAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) DescribeInstanceTypes(in *DescribeInstanceTypesInput) (out *DescribeInstanceTypesOutput, err error) {
	if in == nil {
		in = &DescribeInstanceTypesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeInstanceTypes",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeInstanceTypesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *DescribeInstanceTypesInput) Validate() error {
	return nil
}

func (p *InstanceService) CreateBrokers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateBrokers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *InstanceService) DeleteBrokers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteBrokers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf.Empty{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

func init() { proto.RegisterFile("instances.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x25, 0xb5, 0xa3, 0x36, 0xe3, 0xf8, 0x23, 0x5b, 0x6a, 0x14, 0xe5, 0x50, 0xd7, 0xa5, 0xe0,
	0x43, 0x91, 0x21, 0xa5, 0xd7, 0x36, 0x89, 0x9c, 0x83, 0x0b, 0xfd, 0x40, 0xce, 0xb1, 0x20, 0xd6,
	0xd2, 0x24, 0x2c, 0xb5, 0xb4, 0x62, 0x77, 0x14, 0x70, 0x7e, 0x62, 0xcf, 0xfd, 0x27, 0xfd, 0x03,
	0x45, 0x2b, 0xcb, 0x95, 0x83, 0xeb, 0x82, 0x7c, 0xd3, 0x3e, 0xde, 0xbc, 0xf7, 0x34, 0x33, 0xbb,
	0xd0, 0x15, 0x89, 0x26, 0x9e, 0x84, 0xa8, 0xdd, 0x54, 0x49, 0x92, 0xac, 0xa9, 0x53, 0x0c, 0x9d,
	0xb3, 0x3b, 0x29, 0xef, 0x16, 0x38, 0x36, 0xd8, 0x3c, 0xbb, 0x1d, 0x63, 0x9c, 0xd2, 0xb2, 0xa0,
	0x0c, 0xc7, 0x70, 0x3a, 0x5d, 0x55, 0xcd, 0x50, 0xdd, 0x8b, 0x10, 0xbf, 0x29, 0x99, 0xa2, 0x22,
	0x81, 0x9a, 0x31, 0x68, 0x3e, 0xc8, 0x04, 0xed, 0x83, 0xc1, 0xc1, 0xe8, 0xc8, 0x37, 0xdf, 0x43,
	0x09, 0xce, 0x04, 0x75, 0xa8, 0xc4, 0x1c, 0xcb, 0xc2, 0x9b, 0x65, 0x8a, 0x7a, 0x9a, 0xa4, 0x19,
	0xb1, 0x3e, 0x58, 0x3c, 0x24, 0x21, 0x93, 0x55, 0xcd, 0xea, 0xb4, 0x56, 0x7a, 0xf2, 0x57, 0x89,
	0xbd, 0x81, 0x4e, 0x19, 0x38, 0xa0, 0x5c, 0xc2, 0x6e, 0x0c, 0x1a, 0xa3, 0x23, 0xbf, 0x2d, 0xaa,
	0xba, 0xc3, 0x9f, 0x0d, 0x38, 0xdb, 0xea, 0xf8, 0x35, 0xa3, 0x5d, 0x96, 0xa7, 0xf0, 0x4c, 0x21,
	0x05, 0xa1, 0x8c, 0x0a, 0xdb, 0x43, 0xff, 0xa9, 0x42, 0xf2, 0x64, 0x84, 0x8c, 0xc3, 0xc9, 0x86,
	0x73, 0xa0, 0x91, 0x8c, 0x79, 0xeb, 0xfc, 0xbd, 0x9b, 0xf7, 0xcc, 0xdd, 0x61, 0xe8, 0x56, 0xb1,
	0x19, 0xd2, 0xf5, 0x02, 0x63, 0xbf, 0x2b, 0x36, 0x41, 0xf6, 0x12, 0x5a, 0x24, 0x89, 0x2f, 0x82,
	0x50, 0x66, 0x09, 0xd9, 0x4d, 0x13, 0x00, 0x0c, 0xe4, 0xe5, 0x88, 0xf3, 0xfb, 0x00, 0x9e, 0x6f,
	0x51, 0x62, 0x23, 0xe8, 0x6d, 0x66, 0x13, 0xd1, 0xea, 0xc7, 0x3a, 0x55, 0x8f, 0x69, 0xc4, 0xde,
	0x02, 0xdb, 0x64, 0x26, 0x3c, 0x2e, 0x3b, 0xdc, 0xab, 0x72, 0xbf, 0xf0, 0x18, 0xd9, 0x00, 0x5a,
	0x91, 0xf9, 0xa9, 0xd4, 0xf4, 0xaa, 0x61, 0x68, 0x55, 0x88, 0xbd, 0x86, 0xf6, 0x7d, 0x98, 0x66,
	0x3a, 0x08, 0x33, 0xa5, 0x70, 0x1d, 0xfa, 0xd8, 0x80, 0x5e, 0x81, 0xe5, 0x43, 0x8b, 0x31, 0x96,
	0x6a, 0xb9, 0x66, 0x1d, 0x1a, 0x56, 0xbb, 0x40, 0x4b, 0x5a, 0x1f, 0x2c, 0x4d, 0x9c, 0x32, 0x6d,
	0x5b, 0xc5, 0x50, 0x8a, 0xd3, 0xf9, 0x2f, 0x0b, 0xba, 0x8f, 0xf6, 0x8d, 0x79, 0x70, 0xf2, 0xb8,
	0xdd, 0x9a, 0xf5, 0xdd, 0x62, 0x6b, 0xdd, 0x72, 0x6b, 0xdd, 0xeb, 0x7c, 0x6b, 0x9d, 0x7f, 0xe0,
	0xec, 0x03, 0x1c, 0xfb, 0x59, 0x52, 0xbf, 0x7e, 0x02, 0xec, 0x06, 0x55, 0x2c, 0x12, 0x4e, 0x7b,
	0xa4, 0xb8, 0x80, 0xce, 0x8c, 0xb8, 0xa2, 0xfa, 0x0a, 0x1f, 0xa1, 0x3d, 0x23, 0x99, 0xd6, 0x17,
	0xb8, 0x82, 0x9e, 0x8f, 0x7a, 0xbf, 0x10, 0x17, 0xd0, 0xf1, 0x51, 0xe3, 0x1e, 0x0a, 0x97, 0xd0,
	0xf5, 0x51, 0x8b, 0x87, 0x3d, 0x7a, 0xf9, 0x09, 0xec, 0xcf, 0x32, 0x12, 0xb7, 0xcb, 0x52, 0xe2,
	0x92, 0x48, 0x89, 0x79, 0x46, 0x35, 0xb4, 0xbe, 0xc3, 0x8b, 0xad, 0x37, 0x9a, 0x0d, 0x76, 0x5c,
	0x77, 0xf3, 0xa2, 0x39, 0xaf, 0xfe, 0xfb, 0x20, 0xe4, 0x33, 0xf3, 0x14, 0x72, 0xc2, 0x2b, 0x25,
	0x7f, 0xa0, 0xaa, 0x35, 0xf4, 0x09, 0x2e, 0xb0, 0xb6, 0xc0, 0xdc, 0x32, 0xe7, 0x77, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xeb, 0x98, 0xd1, 0x02, 0x06, 0x00, 0x00,
}
