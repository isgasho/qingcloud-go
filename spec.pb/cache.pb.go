// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

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

type CacheServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *CacheServiceProperties) Reset()                    { *m = CacheServiceProperties{} }
func (m *CacheServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*CacheServiceProperties) ProtoMessage()               {}
func (*CacheServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CacheServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*CacheServiceProperties)(nil), "spec.CacheServiceProperties")
}

type CacheServiceInterface interface {
	DescribeCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateCache(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RestartCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateCache(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ChangeCacheVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyCacheAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RestartCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyCacheNodeAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateCacheFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeCacheParameterGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateCacheParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplyCacheParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteCacheParameterGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyCacheParameterGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResetCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type CacheService struct {
	Config     *config.Config
	Properties *CacheServiceProperties
}

func NewCacheService(conf *config.Config, zone string) (p *CacheService, err error) {
	return &CacheService{
		Config:     conf,
		Properties: &CacheServiceProperties{Zone: zone},
	}, nil
}

func (p *CacheService) DescribeCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCaches",
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
func (p *CacheService) CreateCache(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCache",
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
func (p *CacheService) StopCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopCaches",
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
func (p *CacheService) StartCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartCaches",
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
func (p *CacheService) RestartCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartCaches",
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
func (p *CacheService) DeleteCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCaches",
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
func (p *CacheService) ResizeCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeCaches",
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
func (p *CacheService) UpdateCache(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateCache",
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
func (p *CacheService) ChangeCacheVxnet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeCacheVxnet",
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
func (p *CacheService) ModifyCacheAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheAttributes",
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
func (p *CacheService) DescribeCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheNodes",
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
func (p *CacheService) AddCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddCacheNodes",
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
func (p *CacheService) DeleteCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCacheNodes",
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
func (p *CacheService) RestartCacheNodes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartCacheNodes",
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
func (p *CacheService) ModifyCacheNodeAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheNodeAttributes",
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
func (p *CacheService) CreateCacheFromSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCacheFromSnapshot",
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
func (p *CacheService) DescribeCacheParameterGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheParameterGroups",
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
func (p *CacheService) CreateCacheParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateCacheParameterGroup",
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
func (p *CacheService) ApplyCacheParameterGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyCacheParameterGroup",
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
func (p *CacheService) DeleteCacheParameterGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteCacheParameterGroups",
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
func (p *CacheService) ModifyCacheParameterGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyCacheParameterGroupAttributes",
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
func (p *CacheService) DescribeCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeCacheParameters",
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
func (p *CacheService) UpdateCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateCacheParameters",
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
func (p *CacheService) ResetCacheParameters(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResetCacheParameters",
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

func init() { proto.RegisterFile("cache.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x11, 0xe4, 0xc1, 0xbb, 0xea, 0xc3, 0x17, 0xde, 0xb3, 0xd6, 0x76, 0x51, 0xda, 0x4d,
	0x17, 0x25, 0x42, 0xbb, 0xed, 0x3f, 0x8d, 0x55, 0xfa, 0x4f, 0x24, 0xc1, 0xee, 0x27, 0xc9, 0x55,
	0x03, 0x9a, 0x19, 0x66, 0xae, 0xa5, 0xfa, 0x95, 0xfa, 0x25, 0x8b, 0x09, 0x85, 0x09, 0xb4, 0x8b,
	0xcc, 0x74, 0x97, 0x5c, 0xe6, 0xfc, 0x6e, 0xee, 0x39, 0x77, 0x02, 0xb5, 0x88, 0x45, 0x0b, 0x74,
	0x85, 0xe4, 0xc4, 0x9d, 0xaa, 0x12, 0x18, 0x75, 0x0e, 0xe6, 0x9c, 0xcf, 0x97, 0xd8, 0xcd, 0x6a,
	0xe1, 0x7a, 0xd6, 0xc5, 0x95, 0xa0, 0x4d, 0x7e, 0xe4, 0xf8, 0x0c, 0x5a, 0xde, 0x4e, 0x11, 0xa0,
	0x7c, 0x4d, 0x22, 0x9c, 0x48, 0x2e, 0x50, 0x52, 0x82, 0xca, 0x71, 0xa0, 0xba, 0xe5, 0x29, 0xb6,
	0x2b, 0x47, 0x95, 0xd3, 0xdf, 0x7e, 0xf6, 0x7c, 0xfe, 0xde, 0x80, 0xba, 0x7e, 0xdc, 0xb9, 0x85,
	0x3f, 0x03, 0x54, 0x91, 0x4c, 0x42, 0xcc, 0xea, 0xca, 0x69, 0xb9, 0x79, 0x3b, 0xf7, 0xb3, 0x9d,
	0x7b, 0xb7, 0x6b, 0xd7, 0xf9, 0xa6, 0xee, 0x5c, 0x41, 0xcd, 0x93, 0xc8, 0x28, 0xd7, 0x97, 0x96,
	0x5f, 0x02, 0x04, 0xc4, 0x85, 0x79, 0xf3, 0x80, 0x98, 0x24, 0x43, 0xf9, 0x0d, 0x34, 0x7c, 0x54,
	0x16, 0x80, 0x6b, 0xa8, 0x0f, 0x70, 0x89, 0x84, 0xe6, 0x7a, 0x1f, 0x55, 0xb2, 0xb5, 0x30, 0x7f,
	0x2a, 0x62, 0x63, 0xf3, 0xfb, 0xd0, 0xf4, 0x16, 0x2c, 0x9d, 0xe7, 0xf2, 0x97, 0xb7, 0x14, 0xa9,
	0x34, 0x63, 0x04, 0xff, 0x9f, 0x79, 0x9c, 0xcc, 0x36, 0x19, 0xa3, 0x47, 0x24, 0x93, 0x70, 0x4d,
	0x06, 0xb3, 0x0c, 0xc0, 0x29, 0xac, 0xe2, 0x98, 0xc7, 0x66, 0x91, 0xf6, 0xe2, 0xd8, 0x02, 0xd0,
	0x87, 0xa6, 0x16, 0xa9, 0x19, 0xc3, 0x83, 0xbf, 0xfa, 0x5e, 0x99, 0x41, 0x1e, 0x61, 0x5f, 0x33,
	0x76, 0xc7, 0xb0, 0x30, 0xf7, 0x1e, 0xf6, 0xb4, 0x5b, 0x3a, 0x94, 0x7c, 0x15, 0xa4, 0x4c, 0xa8,
	0x05, 0x2f, 0x1f, 0xf8, 0x18, 0x0e, 0x0b, 0x39, 0x4d, 0x98, 0x64, 0x2b, 0x24, 0x94, 0x23, 0xc9,
	0xd7, 0xc2, 0x68, 0x4e, 0xed, 0xd3, 0x8a, 0xb4, 0xd2, 0xb0, 0x07, 0x68, 0xf7, 0x84, 0x58, 0x6e,
	0x7e, 0x82, 0xf5, 0x04, 0x1d, 0x6d, 0x13, 0x6c, 0xc7, 0x9c, 0xc2, 0x89, 0x16, 0x67, 0x91, 0x66,
	0x17, 0xec, 0xd7, 0x69, 0x28, 0x93, 0x9b, 0xac, 0xfd, 0x4c, 0x2c, 0x40, 0x43, 0xf8, 0xe7, 0xa3,
	0x42, 0xb2, 0xe4, 0x84, 0xbf, 0xb2, 0xf7, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x50,
	0xc3, 0x43, 0x14, 0x07, 0x00, 0x00,
}
