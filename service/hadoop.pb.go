// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hadoop.proto

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

type HadoopServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *HadoopServiceProperties) Reset()                    { *m = HadoopServiceProperties{} }
func (m *HadoopServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*HadoopServiceProperties) ProtoMessage()               {}
func (*HadoopServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *HadoopServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type AddHadoopNodesInput struct {
}

func (m *AddHadoopNodesInput) Reset()                    { *m = AddHadoopNodesInput{} }
func (m *AddHadoopNodesInput) String() string            { return proto.CompactTextString(m) }
func (*AddHadoopNodesInput) ProtoMessage()               {}
func (*AddHadoopNodesInput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type AddHadoopNodesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AddHadoopNodesOutput) Reset()                    { *m = AddHadoopNodesOutput{} }
func (m *AddHadoopNodesOutput) String() string            { return proto.CompactTextString(m) }
func (*AddHadoopNodesOutput) ProtoMessage()               {}
func (*AddHadoopNodesOutput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *AddHadoopNodesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AddHadoopNodesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *AddHadoopNodesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteHadoopNodesInput struct {
}

func (m *DeleteHadoopNodesInput) Reset()                    { *m = DeleteHadoopNodesInput{} }
func (m *DeleteHadoopNodesInput) String() string            { return proto.CompactTextString(m) }
func (*DeleteHadoopNodesInput) ProtoMessage()               {}
func (*DeleteHadoopNodesInput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type DeleteHadoopNodesOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteHadoopNodesOutput) Reset()                    { *m = DeleteHadoopNodesOutput{} }
func (m *DeleteHadoopNodesOutput) String() string            { return proto.CompactTextString(m) }
func (*DeleteHadoopNodesOutput) ProtoMessage()               {}
func (*DeleteHadoopNodesOutput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *DeleteHadoopNodesOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteHadoopNodesOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DeleteHadoopNodesOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StartHadoopsInput struct {
}

func (m *StartHadoopsInput) Reset()                    { *m = StartHadoopsInput{} }
func (m *StartHadoopsInput) String() string            { return proto.CompactTextString(m) }
func (*StartHadoopsInput) ProtoMessage()               {}
func (*StartHadoopsInput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

type StartHadoopsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *StartHadoopsOutput) Reset()                    { *m = StartHadoopsOutput{} }
func (m *StartHadoopsOutput) String() string            { return proto.CompactTextString(m) }
func (*StartHadoopsOutput) ProtoMessage()               {}
func (*StartHadoopsOutput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *StartHadoopsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *StartHadoopsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *StartHadoopsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StopHadoopsInput struct {
}

func (m *StopHadoopsInput) Reset()                    { *m = StopHadoopsInput{} }
func (m *StopHadoopsInput) String() string            { return proto.CompactTextString(m) }
func (*StopHadoopsInput) ProtoMessage()               {}
func (*StopHadoopsInput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

type StopHadoopsOutput struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *StopHadoopsOutput) Reset()                    { *m = StopHadoopsOutput{} }
func (m *StopHadoopsOutput) String() string            { return proto.CompactTextString(m) }
func (*StopHadoopsOutput) ProtoMessage()               {}
func (*StopHadoopsOutput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *StopHadoopsOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *StopHadoopsOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *StopHadoopsOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HadoopServiceProperties)(nil), "service.HadoopServiceProperties")
	proto.RegisterType((*AddHadoopNodesInput)(nil), "service.AddHadoopNodesInput")
	proto.RegisterType((*AddHadoopNodesOutput)(nil), "service.AddHadoopNodesOutput")
	proto.RegisterType((*DeleteHadoopNodesInput)(nil), "service.DeleteHadoopNodesInput")
	proto.RegisterType((*DeleteHadoopNodesOutput)(nil), "service.DeleteHadoopNodesOutput")
	proto.RegisterType((*StartHadoopsInput)(nil), "service.StartHadoopsInput")
	proto.RegisterType((*StartHadoopsOutput)(nil), "service.StartHadoopsOutput")
	proto.RegisterType((*StopHadoopsInput)(nil), "service.StopHadoopsInput")
	proto.RegisterType((*StopHadoopsOutput)(nil), "service.StopHadoopsOutput")
}

type HadoopServiceInterface interface {
	AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error)
	DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error)
	StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error)
	StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error)
}

type HadoopService struct {
	Config           *config.Config
	Properties       *HadoopServiceProperties
	LastResponseBody string
}

func NewHadoopService(conf *config.Config, zone string) (p *HadoopService) {
	return &HadoopService{
		Config:     conf,
		Properties: &HadoopServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Hadoop(zone string) (*HadoopService, error) {
	properties := &HadoopServiceProperties{
		Zone: zone,
	}

	return &HadoopService{Config: s.Config, Properties: properties}, nil
}

func (p *HadoopService) AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error) {
	if in == nil {
		in = &AddHadoopNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddHadoopNodes",
		RequestMethod: "GET", // GET or POST
	}

	x := &AddHadoopNodesOutput{}
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

func (p *AddHadoopNodesInput) Validate() error {
	return nil
}

func (p *HadoopService) DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error) {
	if in == nil {
		in = &DeleteHadoopNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteHadoopNodes",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteHadoopNodesOutput{}
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

func (p *DeleteHadoopNodesInput) Validate() error {
	return nil
}

func (p *HadoopService) StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error) {
	if in == nil {
		in = &StartHadoopsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartHadoops",
		RequestMethod: "GET", // GET or POST
	}

	x := &StartHadoopsOutput{}
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

func (p *StartHadoopsInput) Validate() error {
	return nil
}

func (p *HadoopService) StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error) {
	if in == nil {
		in = &StopHadoopsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopHadoops",
		RequestMethod: "GET", // GET or POST
	}

	x := &StopHadoopsOutput{}
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

func (p *StopHadoopsInput) Validate() error {
	return nil
}

func init() { proto.RegisterFile("hadoop.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x0e, 0x88, 0x20, 0x23, 0x1a, 0x19, 0x14, 0xca, 0x2a, 0x91, 0xf4, 0xc4, 0x45, 0x4c, 0xf4,
	0x09, 0x0c, 0x24, 0xea, 0xc1, 0x3f, 0x48, 0xbc, 0x62, 0xed, 0x8e, 0xd8, 0x88, 0xdd, 0xba, 0xdd,
	0x7a, 0xf0, 0x11, 0x7c, 0x48, 0xdf, 0xc1, 0x37, 0x30, 0x74, 0x57, 0xd2, 0xa6, 0xe5, 0xd8, 0x4b,
	0xd3, 0x9d, 0x6f, 0xbe, 0xef, 0x9b, 0xec, 0x37, 0x0b, 0x8d, 0x57, 0x87, 0x0b, 0x11, 0x0c, 0x03,
	0x29, 0x94, 0xc0, 0x5a, 0x48, 0xf2, 0xd3, 0x73, 0x89, 0xf5, 0x3e, 0x3c, 0x7f, 0xee, 0x2e, 0x44,
	0xc4, 0x67, 0x21, 0x7f, 0x9b, 0xc9, 0x68, 0x41, 0xa7, 0xcb, 0x8f, 0xee, 0xb3, 0x4f, 0xa0, 0x73,
	0x15, 0xf3, 0xa6, 0xba, 0xff, 0x5e, 0x8a, 0x80, 0xa4, 0xf2, 0x28, 0x44, 0x84, 0xca, 0x97, 0xf0,
	0xc9, 0x2a, 0xf5, 0x4b, 0x83, 0xfa, 0x24, 0xfe, 0xb7, 0x0f, 0xa0, 0x75, 0xc1, 0xb9, 0x66, 0xdc,
	0x0a, 0x4e, 0xe1, 0xb5, 0x1f, 0x44, 0xca, 0x76, 0x61, 0x3f, 0x5d, 0xbe, 0x8b, 0x54, 0x10, 0x29,
	0x6c, 0x43, 0xd5, 0x71, 0x95, 0x27, 0x7c, 0x23, 0x62, 0x4e, 0xd8, 0x85, 0x2d, 0x49, 0x6a, 0xe6,
	0x0a, 0x4e, 0x56, 0xb9, 0x5f, 0x1a, 0x6c, 0x4e, 0x6a, 0x92, 0xd4, 0x48, 0x70, 0x42, 0x0b, 0x6a,
	0xef, 0x14, 0x86, 0xce, 0x9c, 0xac, 0x8d, 0x98, 0xf3, 0x7f, 0xb4, 0x2d, 0x68, 0x8f, 0x69, 0x41,
	0x8a, 0x32, 0xf6, 0x2f, 0xd0, 0xc9, 0x20, 0x45, 0x4c, 0xd0, 0x82, 0xe6, 0x54, 0x39, 0x52, 0x69,
	0x1b, 0x63, 0xee, 0x00, 0x26, 0x8b, 0x45, 0xf8, 0x22, 0xec, 0x4d, 0x95, 0x08, 0x52, 0xb6, 0x4f,
	0xcb, 0x59, 0x56, 0xb5, 0x02, 0x5c, 0xcf, 0x7e, 0xca, 0xb0, 0x93, 0xda, 0x0d, 0xbc, 0x81, 0xdd,
	0x74, 0xcc, 0x78, 0x34, 0x34, 0x7b, 0x36, 0xcc, 0x59, 0x0b, 0xd6, 0x5b, 0x83, 0x9a, 0x69, 0x1f,
	0xa1, 0x99, 0x89, 0x0d, 0x8f, 0x57, 0x9c, 0xfc, 0xb0, 0x59, 0x7f, 0x7d, 0x83, 0xd1, 0xbd, 0x84,
	0x46, 0x32, 0x11, 0x64, 0x2b, 0x46, 0x26, 0x3d, 0x76, 0x98, 0x8b, 0x19, 0xa1, 0x31, 0x6c, 0x27,
	0xee, 0x18, 0xbb, 0x89, 0xde, 0x74, 0x1a, 0x8c, 0xe5, 0x41, 0x5a, 0x85, 0x75, 0xbe, 0x7f, 0x2b,
	0x2d, 0xa8, 0x3f, 0x78, 0xfe, 0x7c, 0xb4, 0x7c, 0x87, 0x58, 0xd5, 0xf8, 0x73, 0x35, 0x7e, 0x82,
	0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0xa3, 0x99, 0x7d, 0xba, 0x03, 0x00, 0x00,
}
