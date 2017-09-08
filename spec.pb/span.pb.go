// Code generated by protoc-gen-go. DO NOT EDIT.
// source: span.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"

var _ = config.Config{}
var _ = request.Request{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SpanServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SpanServiceProperties) Reset()                    { *m = SpanServiceProperties{} }
func (m *SpanServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SpanServiceProperties) ProtoMessage()               {}
func (*SpanServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *SpanServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*SpanServiceProperties)(nil), "spec.SpanServiceProperties")
}

type SpanServiceInterface interface {
	CreateSpan(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeSpans(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSpans(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddSpanMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RemoveSpanMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifySpanAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateSpan(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type SpanService struct {
	Config     *config.Config
	Properties *SpanServiceProperties
}

func NewSpanService(conf *config.Config, zone string) (p *SpanService, err error) {
	return &SpanService{
		Config:     conf,
		Properties: &SpanServiceProperties{Zone: zone},
	}, nil
}

func (p *SpanService) CreateSpan(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSpan",
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

func (p *SpanService) DescribeSpans(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSpans",
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

func (p *SpanService) DeleteSpans(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSpans",
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

func (p *SpanService) AddSpanMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddSpanMembers",
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

func (p *SpanService) RemoveSpanMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RemoveSpanMembers",
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

func (p *SpanService) ModifySpanAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySpanAttributes",
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

func (p *SpanService) UpdateSpan(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateSpan",
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

func init() { proto.RegisterFile("span.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0x42, 0x41,
	0x14, 0x85, 0x91, 0x24, 0xe8, 0x4a, 0x41, 0x43, 0x45, 0xd8, 0x26, 0x5a, 0x05, 0xc1, 0x13, 0x6a,
	0x6b, 0x94, 0x68, 0xed, 0x84, 0x50, 0xfa, 0x01, 0x6f, 0xde, 0x1c, 0x65, 0xc0, 0xf7, 0xee, 0x70,
	0xe7, 0x2a, 0xd8, 0x2f, 0xed, 0xe7, 0xc4, 0x8c, 0x04, 0x6d, 0x5a, 0x3c, 0xdd, 0xcd, 0x1c, 0xce,
	0xf9, 0xbe, 0xc5, 0x25, 0x8a, 0xa1, 0x6c, 0x8a, 0x20, 0xac, 0x6c, 0xba, 0x31, 0xa0, 0xea, 0xdf,
	0x2c, 0x99, 0x97, 0x2b, 0x0c, 0x72, 0x66, 0xd7, 0x8b, 0x01, 0xea, 0xa0, 0xdb, 0x5d, 0xe5, 0xee,
	0x81, 0x2e, 0xe7, 0xa1, 0x6c, 0xe6, 0x90, 0x8d, 0xaf, 0xf0, 0x21, 0x1c, 0x20, 0xea, 0x11, 0x8d,
	0xa1, 0xee, 0x17, 0x37, 0xb8, 0xee, 0xdc, 0x76, 0xee, 0x4f, 0x66, 0xf9, 0xfd, 0xf8, 0x7d, 0x44,
	0xbd, 0x3f, 0x6d, 0x33, 0x24, 0x1a, 0x0b, 0x4a, 0x45, 0x0a, 0xcd, 0x55, 0xb1, 0x13, 0x15, 0xbf,
	0xa2, 0xe2, 0x2d, 0x89, 0xfa, 0xff, 0xe4, 0xe6, 0x85, 0x4e, 0x27, 0x88, 0x95, 0x78, 0x9b, 0xf7,
	0xb1, 0x35, 0xe0, 0x99, 0x7a, 0x13, 0xac, 0xa0, 0x7b, 0xce, 0x5f, 0xe9, 0x6c, 0xe4, 0x5c, 0xda,
	0x4e, 0x51, 0x5b, 0x48, 0x7b, 0xc2, 0x98, 0xce, 0x67, 0xa8, 0x79, 0x83, 0x43, 0x20, 0xef, 0x74,
	0x31, 0x65, 0xe7, 0x17, 0xdb, 0x04, 0x19, 0xa9, 0x8a, 0xb7, 0x6b, 0x45, 0x7b, 0xce, 0x90, 0xe8,
	0x33, 0xb8, 0x3d, 0x8f, 0x61, 0x8f, 0xf3, 0xff, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x9d,
	0x6d, 0xd7, 0x3f, 0x02, 0x00, 0x00,
}
