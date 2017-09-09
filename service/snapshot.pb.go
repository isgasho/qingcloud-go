// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot.proto

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

type SnapshotServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SnapshotServiceProperties) Reset()                    { *m = SnapshotServiceProperties{} }
func (m *SnapshotServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SnapshotServiceProperties) ProtoMessage()               {}
func (*SnapshotServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *SnapshotServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type DescribeSnapshotsInput struct {
}

func (m *DescribeSnapshotsInput) Reset()                    { *m = DescribeSnapshotsInput{} }
func (m *DescribeSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSnapshotsInput) ProtoMessage()               {}
func (*DescribeSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

type DescribeSnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeSnapshotsOutput) Reset()                    { *m = DescribeSnapshotsOutput{} }
func (m *DescribeSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeSnapshotsOutput) ProtoMessage()               {}
func (*DescribeSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func (m *DescribeSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateSnapshotsInput struct {
}

func (m *CreateSnapshotsInput) Reset()                    { *m = CreateSnapshotsInput{} }
func (m *CreateSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotsInput) ProtoMessage()               {}
func (*CreateSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{3} }

type CreateSnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CreateSnapshotsOutput) Reset()                    { *m = CreateSnapshotsOutput{} }
func (m *CreateSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*CreateSnapshotsOutput) ProtoMessage()               {}
func (*CreateSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{4} }

func (m *CreateSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteSnapshotsInput struct {
}

func (m *DeleteSnapshotsInput) Reset()                    { *m = DeleteSnapshotsInput{} }
func (m *DeleteSnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotsInput) ProtoMessage()               {}
func (*DeleteSnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{5} }

type DeleteSnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteSnapshotsOutput) Reset()                    { *m = DeleteSnapshotsOutput{} }
func (m *DeleteSnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteSnapshotsOutput) ProtoMessage()               {}
func (*DeleteSnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{6} }

func (m *DeleteSnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteSnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteSnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ApplySnapshotsInput struct {
}

func (m *ApplySnapshotsInput) Reset()                    { *m = ApplySnapshotsInput{} }
func (m *ApplySnapshotsInput) String() string            { return proto.CompactTextString(m) }
func (*ApplySnapshotsInput) ProtoMessage()               {}
func (*ApplySnapshotsInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{7} }

type ApplySnapshotsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ApplySnapshotsOutput) Reset()                    { *m = ApplySnapshotsOutput{} }
func (m *ApplySnapshotsOutput) String() string            { return proto.CompactTextString(m) }
func (*ApplySnapshotsOutput) ProtoMessage()               {}
func (*ApplySnapshotsOutput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{8} }

func (m *ApplySnapshotsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ApplySnapshotsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ApplySnapshotsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ModifySnapshotAttributesInput struct {
}

func (m *ModifySnapshotAttributesInput) Reset()                    { *m = ModifySnapshotAttributesInput{} }
func (m *ModifySnapshotAttributesInput) String() string            { return proto.CompactTextString(m) }
func (*ModifySnapshotAttributesInput) ProtoMessage()               {}
func (*ModifySnapshotAttributesInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{9} }

type ModifySnapshotAttributesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ModifySnapshotAttributesOutput) Reset()         { *m = ModifySnapshotAttributesOutput{} }
func (m *ModifySnapshotAttributesOutput) String() string { return proto.CompactTextString(m) }
func (*ModifySnapshotAttributesOutput) ProtoMessage()    {}
func (*ModifySnapshotAttributesOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{10}
}

func (m *ModifySnapshotAttributesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ModifySnapshotAttributesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ModifySnapshotAttributesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CaptureInstanceFromSnapshotInput struct {
}

func (m *CaptureInstanceFromSnapshotInput) Reset()         { *m = CaptureInstanceFromSnapshotInput{} }
func (m *CaptureInstanceFromSnapshotInput) String() string { return proto.CompactTextString(m) }
func (*CaptureInstanceFromSnapshotInput) ProtoMessage()    {}
func (*CaptureInstanceFromSnapshotInput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{11}
}

type CaptureInstanceFromSnapshotOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CaptureInstanceFromSnapshotOutput) Reset()         { *m = CaptureInstanceFromSnapshotOutput{} }
func (m *CaptureInstanceFromSnapshotOutput) String() string { return proto.CompactTextString(m) }
func (*CaptureInstanceFromSnapshotOutput) ProtoMessage()    {}
func (*CaptureInstanceFromSnapshotOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{12}
}

func (m *CaptureInstanceFromSnapshotOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CaptureInstanceFromSnapshotOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CaptureInstanceFromSnapshotOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateVolumeFromSnapshotInput struct {
}

func (m *CreateVolumeFromSnapshotInput) Reset()                    { *m = CreateVolumeFromSnapshotInput{} }
func (m *CreateVolumeFromSnapshotInput) String() string            { return proto.CompactTextString(m) }
func (*CreateVolumeFromSnapshotInput) ProtoMessage()               {}
func (*CreateVolumeFromSnapshotInput) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{13} }

type CreateVolumeFromSnapshotOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CreateVolumeFromSnapshotOutput) Reset()         { *m = CreateVolumeFromSnapshotOutput{} }
func (m *CreateVolumeFromSnapshotOutput) String() string { return proto.CompactTextString(m) }
func (*CreateVolumeFromSnapshotOutput) ProtoMessage()    {}
func (*CreateVolumeFromSnapshotOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{14}
}

func (m *CreateVolumeFromSnapshotOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreateVolumeFromSnapshotOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CreateVolumeFromSnapshotOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotServiceProperties)(nil), "service.SnapshotServiceProperties")
	proto.RegisterType((*DescribeSnapshotsInput)(nil), "service.DescribeSnapshotsInput")
	proto.RegisterType((*DescribeSnapshotsOutput)(nil), "service.DescribeSnapshotsOutput")
	proto.RegisterType((*CreateSnapshotsInput)(nil), "service.CreateSnapshotsInput")
	proto.RegisterType((*CreateSnapshotsOutput)(nil), "service.CreateSnapshotsOutput")
	proto.RegisterType((*DeleteSnapshotsInput)(nil), "service.DeleteSnapshotsInput")
	proto.RegisterType((*DeleteSnapshotsOutput)(nil), "service.DeleteSnapshotsOutput")
	proto.RegisterType((*ApplySnapshotsInput)(nil), "service.ApplySnapshotsInput")
	proto.RegisterType((*ApplySnapshotsOutput)(nil), "service.ApplySnapshotsOutput")
	proto.RegisterType((*ModifySnapshotAttributesInput)(nil), "service.ModifySnapshotAttributesInput")
	proto.RegisterType((*ModifySnapshotAttributesOutput)(nil), "service.ModifySnapshotAttributesOutput")
	proto.RegisterType((*CaptureInstanceFromSnapshotInput)(nil), "service.CaptureInstanceFromSnapshotInput")
	proto.RegisterType((*CaptureInstanceFromSnapshotOutput)(nil), "service.CaptureInstanceFromSnapshotOutput")
	proto.RegisterType((*CreateVolumeFromSnapshotInput)(nil), "service.CreateVolumeFromSnapshotInput")
	proto.RegisterType((*CreateVolumeFromSnapshotOutput)(nil), "service.CreateVolumeFromSnapshotOutput")
}

type SnapshotServiceInterface interface {
	DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error)
	CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error)
	DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error)
	ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error)
	ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error)
	CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error)
	CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error)
}

type SnapshotService struct {
	Config           *config.Config
	Properties       *SnapshotServiceProperties
	LastResponseBody string
}

func NewSnapshotService(conf *config.Config, zone string) (p *SnapshotService) {
	return &SnapshotService{
		Config:     conf,
		Properties: &SnapshotServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Snapshot(zone string) (*SnapshotService, error) {
	properties := &SnapshotServiceProperties{
		Zone: zone,
	}

	return &SnapshotService{Config: s.Config, Properties: properties}, nil
}

func (p *SnapshotService) DescribeSnapshots(in *DescribeSnapshotsInput) (out *DescribeSnapshotsOutput, err error) {
	if in == nil {
		in = &DescribeSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeSnapshotsOutput{}
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

func (p *DescribeSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) CreateSnapshots(in *CreateSnapshotsInput) (out *CreateSnapshotsOutput, err error) {
	if in == nil {
		in = &CreateSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateSnapshotsOutput{}
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

func (p *CreateSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) DeleteSnapshots(in *DeleteSnapshotsInput) (out *DeleteSnapshotsOutput, err error) {
	if in == nil {
		in = &DeleteSnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteSnapshotsOutput{}
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

func (p *DeleteSnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) ApplySnapshots(in *ApplySnapshotsInput) (out *ApplySnapshotsOutput, err error) {
	if in == nil {
		in = &ApplySnapshotsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplySnapshots",
		RequestMethod: "GET", // GET or POST
	}

	x := &ApplySnapshotsOutput{}
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

func (p *ApplySnapshotsInput) Validate() error {
	return nil
}

func (p *SnapshotService) ModifySnapshotAttributes(in *ModifySnapshotAttributesInput) (out *ModifySnapshotAttributesOutput, err error) {
	if in == nil {
		in = &ModifySnapshotAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySnapshotAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifySnapshotAttributesOutput{}
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

func (p *ModifySnapshotAttributesInput) Validate() error {
	return nil
}

func (p *SnapshotService) CaptureInstanceFromSnapshot(in *CaptureInstanceFromSnapshotInput) (out *CaptureInstanceFromSnapshotOutput, err error) {
	if in == nil {
		in = &CaptureInstanceFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CaptureInstanceFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CaptureInstanceFromSnapshotOutput{}
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

func (p *CaptureInstanceFromSnapshotInput) Validate() error {
	return nil
}

func (p *SnapshotService) CreateVolumeFromSnapshot(in *CreateVolumeFromSnapshotInput) (out *CreateVolumeFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateVolumeFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumeFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateVolumeFromSnapshotOutput{}
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

func (p *CreateVolumeFromSnapshotInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("snapshot.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x61, 0x5b, 0xb7, 0xfb, 0xb0, 0x09, 0xb3, 0x96, 0x34, 0xd0, 0xae, 0xe4, 0x01, 0x06,
	0x0f, 0x9d, 0x04, 0x5f, 0x30, 0x75, 0x42, 0xda, 0xc3, 0xc4, 0xd8, 0xa4, 0xbd, 0x56, 0x69, 0x72,
	0x57, 0xac, 0x25, 0xb6, 0xb1, 0xaf, 0x91, 0xe0, 0x13, 0xf8, 0x31, 0xbe, 0x87, 0x3f, 0x40, 0x6b,
	0x9d, 0x34, 0xcb, 0x92, 0x74, 0x2f, 0x79, 0xa9, 0x6a, 0xfb, 0x9c, 0x7b, 0xee, 0xb5, 0xce, 0x89,
	0x61, 0xdf, 0x88, 0x50, 0x99, 0xef, 0x92, 0x26, 0x4a, 0x4b, 0x92, 0xac, 0x6b, 0x50, 0xff, 0xe4,
	0x11, 0xfa, 0xc3, 0x1f, 0x5c, 0x2c, 0xa2, 0x44, 0xda, 0x78, 0x66, 0xe2, 0xbb, 0x99, 0xb6, 0x09,
	0x9e, 0xdc, 0xff, 0xac, 0x70, 0xc1, 0x09, 0x0c, 0xae, 0x1d, 0xf3, 0x7a, 0xc5, 0xb8, 0xd4, 0x52,
	0xa1, 0x26, 0x8e, 0x86, 0x31, 0xd8, 0xfa, 0x2d, 0x05, 0x7a, 0x9d, 0x71, 0xe7, 0x78, 0xef, 0x6a,
	0xf9, 0x3f, 0xf0, 0xa0, 0x7f, 0x86, 0x26, 0xd2, 0x7c, 0x8e, 0x19, 0xd1, 0x9c, 0x0b, 0x65, 0x29,
	0xb8, 0x85, 0x57, 0x8f, 0x4e, 0xbe, 0x5a, 0x52, 0x96, 0x58, 0x1f, 0x76, 0xc2, 0x88, 0xb8, 0x14,
	0xae, 0x94, 0x5b, 0xb1, 0x01, 0xec, 0x6a, 0xa4, 0x59, 0x24, 0x63, 0xf4, 0x9e, 0x8d, 0x3b, 0xc7,
	0xdb, 0x57, 0x5d, 0x8d, 0x34, 0x95, 0x31, 0x32, 0x0f, 0xba, 0x29, 0x1a, 0x13, 0x2e, 0xd0, 0x7b,
	0xbe, 0xe4, 0x64, 0xcb, 0xa0, 0x0f, 0x87, 0x53, 0x8d, 0x21, 0x95, 0xf5, 0x63, 0xe8, 0x95, 0xf6,
	0x5b, 0x52, 0x3f, 0xc3, 0x04, 0xab, 0xd4, 0x4b, 0xfb, 0x6d, 0xa8, 0xf7, 0xe0, 0xe5, 0xa9, 0x52,
	0xc9, 0xaf, 0x92, 0x78, 0x04, 0x87, 0x0f, 0xb7, 0xdb, 0xd0, 0x3e, 0x82, 0xe1, 0x85, 0x8c, 0xf9,
	0x6d, 0xae, 0x72, 0x4a, 0xa4, 0xf9, 0xdc, 0x12, 0xba, 0x2e, 0x52, 0x18, 0xd5, 0x01, 0xda, 0xe8,
	0x27, 0x80, 0xf1, 0x34, 0x54, 0x64, 0x35, 0x9e, 0x0b, 0x43, 0xa1, 0x88, 0xf0, 0x8b, 0x96, 0x69,
	0xa6, 0xbd, 0x6a, 0x49, 0xc1, 0xdb, 0x06, 0x4c, 0x4b, 0xb7, 0xb4, 0x72, 0xe1, 0x8d, 0x4c, 0x6c,
	0x5a, 0xd1, 0x52, 0x0a, 0xa3, 0x3a, 0x40, 0x0b, 0xfd, 0x7c, 0xfa, 0xbb, 0x0d, 0x07, 0xa5, 0x84,
	0xb3, 0x1b, 0x78, 0xf1, 0x28, 0xa9, 0xec, 0x68, 0xe2, 0x3e, 0x19, 0x93, 0xea, 0x7c, 0xfb, 0xe3,
	0x7a, 0x80, 0x6b, 0xfc, 0x12, 0x0e, 0x4a, 0x09, 0x64, 0xc3, 0x9c, 0x54, 0x95, 0x59, 0x7f, 0x54,
	0x77, 0xbc, 0xae, 0x58, 0x4a, 0x55, 0xa1, 0x62, 0x55, 0x0e, 0x0b, 0x15, 0xab, 0xe3, 0x78, 0x01,
	0xfb, 0x0f, 0xa3, 0xc2, 0xde, 0xe4, 0x8c, 0x8a, 0x68, 0xf9, 0xc3, 0x9a, 0x53, 0x57, 0xee, 0x0e,
	0xbc, 0x3a, 0xcf, 0xb3, 0x77, 0x39, 0xb5, 0x31, 0x37, 0xfe, 0xfb, 0x8d, 0x38, 0x27, 0x46, 0xf0,
	0xba, 0xc1, 0xcd, 0xec, 0xc3, 0xfa, 0x32, 0x37, 0xe4, 0xc2, 0xff, 0xf8, 0x14, 0xe8, 0x7a, 0xc4,
	0x3a, 0xc3, 0x16, 0x46, 0x6c, 0x34, 0x7d, 0x61, 0xc4, 0x66, 0xef, 0xfb, 0x83, 0x3f, 0xff, 0xb6,
	0x7a, 0xb0, 0xf7, 0x8d, 0x8b, 0xc5, 0xf4, 0xfe, 0xd1, 0x62, 0xbb, 0x19, 0x60, 0xbe, 0xb3, 0x7c,
	0xb1, 0x3e, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x27, 0xb3, 0x9c, 0xeb, 0x06, 0x00, 0x00,
}
