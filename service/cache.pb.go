// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

package service

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

type CacheServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *CacheServiceProperties) Reset()                    { *m = CacheServiceProperties{} }
func (m *CacheServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*CacheServiceProperties) ProtoMessage()               {}
func (*CacheServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CacheServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*CacheServiceProperties)(nil), "service.CacheServiceProperties")
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

func NewCacheService(conf *config.Config, zone string) (p *CacheService) {
	return &CacheService{
		Config:     conf,
		Properties: &CacheServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) Cache(zone string) (*CacheService, error) {
	properties := &CacheServiceProperties{
		Zone: zone,
	}

	return &CacheService{Config: s.Config, Properties: properties}, nil
}

func (p *CacheService) DescribeCaches(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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
	o := &data.Operation{
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

func init() { proto.RegisterFile("cache.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0x11, 0x1e, 0xef, 0xf1, 0x46, 0x7d, 0xf8, 0x42, 0x6b, 0xad, 0xed, 0xa1, 0xb4, 0x97,
	0x1e, 0x4a, 0x84, 0xf6, 0xda, 0x5f, 0x1a, 0xab, 0xf4, 0x97, 0x48, 0x82, 0xbd, 0x6f, 0x92, 0x51,
	0x03, 0x9a, 0x5d, 0x76, 0xc7, 0x52, 0xfd, 0x97, 0xfa, 0x4f, 0x16, 0x13, 0x0a, 0x1b, 0x68, 0x0f,
	0xd9, 0xed, 0x2d, 0x19, 0xe6, 0xfb, 0x99, 0xcc, 0x7c, 0x67, 0x02, 0xd5, 0x88, 0x45, 0x73, 0x74,
	0x85, 0xe4, 0xc4, 0x9d, 0x3f, 0x0a, 0xe5, 0x6b, 0x12, 0x61, 0xfb, 0x60, 0xc6, 0xf9, 0x6c, 0x81,
	0x9d, 0x2c, 0x1c, 0xae, 0xa6, 0x1d, 0x5c, 0x0a, 0x5a, 0xe7, 0x59, 0xc7, 0x67, 0xd0, 0xf4, 0xb6,
	0xa2, 0x20, 0x4f, 0x1e, 0x4b, 0x2e, 0x50, 0x52, 0x82, 0xca, 0x71, 0xe0, 0xd7, 0x86, 0xa7, 0xd8,
	0xaa, 0x1c, 0x55, 0x4e, 0xff, 0xfa, 0xd9, 0xf3, 0xf9, 0x7b, 0x1d, 0x6a, 0x7a, 0xba, 0x73, 0x0b,
	0xff, 0xfa, 0xa8, 0x22, 0x99, 0x84, 0x98, 0xc5, 0x95, 0xd3, 0x74, 0xf3, 0x72, 0xee, 0x67, 0x39,
	0xf7, 0x6e, 0x5b, 0xae, 0xfd, 0x4d, 0xdc, 0xb9, 0x82, 0xaa, 0x27, 0x91, 0x51, 0xae, 0x2f, 0x2d,
	0xbf, 0x04, 0x08, 0x88, 0x0b, 0xf3, 0xe2, 0x01, 0x31, 0x49, 0x86, 0xf2, 0x1b, 0xa8, 0xfb, 0xa8,
	0x2c, 0x00, 0xd7, 0x50, 0xeb, 0xe3, 0x02, 0x09, 0xcd, 0xf5, 0x3e, 0xaa, 0x64, 0x63, 0x31, 0xfc,
	0x89, 0x88, 0x8d, 0x87, 0xdf, 0x83, 0x86, 0x37, 0x67, 0xe9, 0x2c, 0x97, 0xbf, 0xbc, 0xa5, 0x48,
	0xa5, 0x19, 0x43, 0xd8, 0x7d, 0xe6, 0x71, 0x32, 0x5d, 0x67, 0x8c, 0x2e, 0x91, 0x4c, 0xc2, 0x15,
	0x19, 0xf4, 0xd2, 0x07, 0xa7, 0xb0, 0x8a, 0x23, 0x1e, 0x9b, 0x59, 0xda, 0x8d, 0x63, 0x0b, 0x40,
	0x0f, 0x1a, 0x9a, 0xa5, 0x66, 0x0c, 0x0f, 0xfe, 0xeb, 0x7b, 0x65, 0x06, 0x79, 0x84, 0x7d, 0x6d,
	0xb0, 0x5b, 0x86, 0xc5, 0x70, 0xef, 0x61, 0x4f, 0xbb, 0xd2, 0x81, 0xe4, 0xcb, 0x20, 0x65, 0x42,
	0xcd, 0x79, 0x79, 0xc3, 0x47, 0x70, 0x58, 0xf0, 0x69, 0xcc, 0x24, 0x5b, 0x22, 0xa1, 0x1c, 0x4a,
	0xbe, 0x12, 0x46, 0x7d, 0x6a, 0x9f, 0x56, 0xa4, 0x95, 0x86, 0x3d, 0x40, 0xab, 0x2b, 0xc4, 0x62,
	0xfd, 0x13, 0xac, 0x27, 0x68, 0x6b, 0x9b, 0x60, 0xdb, 0xe6, 0x04, 0x4e, 0x34, 0x3b, 0x8b, 0x34,
	0x3b, 0x63, 0xbf, 0x76, 0x43, 0x99, 0x5c, 0xb2, 0xf6, 0x33, 0xb1, 0x00, 0x0d, 0x60, 0xc7, 0x47,
	0x85, 0x64, 0xc9, 0x09, 0x7f, 0x67, 0xef, 0x17, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xed,
	0xeb, 0x0e, 0x17, 0x07, 0x00, 0x00,
}
