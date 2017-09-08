// Code generated by protoc-gen-go. DO NOT EDIT.
// source: load_balancer.proto

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

type LoadBalancerServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *LoadBalancerServiceProperties) Reset()                    { *m = LoadBalancerServiceProperties{} }
func (m *LoadBalancerServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerServiceProperties) ProtoMessage()               {}
func (*LoadBalancerServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *LoadBalancerServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*LoadBalancerServiceProperties)(nil), "service.LoadBalancerServiceProperties")
}

type LoadBalancerServiceInterface interface {
	CreateLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StartLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	StopLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	UpdateLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ResizeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AssociateEipsToLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DissociateEipsFromLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerListenerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerBackendAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplyLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateServerCertificate(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyServerCertificateAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type LoadBalancerService struct {
	Config     *config.Config
	Properties *LoadBalancerServiceProperties
}

func NewLoadBalancerService(conf *config.Config, zone string) (p *LoadBalancerService) {
	return &LoadBalancerService{
		Config:     conf,
		Properties: &LoadBalancerServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) LoadBalancer(zone string) (*LoadBalancerService, error) {
	properties := &LoadBalancerServiceProperties{
		Zone: zone,
	}

	return &LoadBalancerService{Config: s.Config, Properties: properties}, nil
}

func (p *LoadBalancerService) CreateLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancer",
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

func (p *LoadBalancerService) DescribeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancers",
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

func (p *LoadBalancerService) DeleteLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancers",
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

func (p *LoadBalancerService) ModifyLoadBalancerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerAttributes",
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

func (p *LoadBalancerService) StartLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartLoadBalancers",
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

func (p *LoadBalancerService) StopLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopLoadBalancers",
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

func (p *LoadBalancerService) UpdateLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateLoadBalancers",
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

func (p *LoadBalancerService) ResizeLoadBalancers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeLoadBalancers",
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

func (p *LoadBalancerService) AssociateEipsToLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEipsToLoadBalancer",
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

func (p *LoadBalancerService) DissociateEipsFromLoadBalancer(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEipsFromLoadBalancer",
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

func (p *LoadBalancerService) AddLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerListeners",
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

func (p *LoadBalancerService) DescribeLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerListeners",
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

func (p *LoadBalancerService) DeleteLoadBalancerListeners(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerListeners",
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

func (p *LoadBalancerService) ModifyLoadBalancerListenerAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerListenerAttributes",
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

func (p *LoadBalancerService) AddLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerBackends",
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

func (p *LoadBalancerService) DescribeLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerBackends",
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

func (p *LoadBalancerService) DeleteLoadBalancerBackends(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerBackends",
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

func (p *LoadBalancerService) ModifyLoadBalancerBackendAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerBackendAttributes",
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

func (p *LoadBalancerService) CreateLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancerPolicy",
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicies",
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyAttributes",
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

func (p *LoadBalancerService) ApplyLoadBalancerPolicy(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyLoadBalancerPolicy",
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicies(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicies",
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

func (p *LoadBalancerService) AddLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerPolicyRules",
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicyRules",
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyRuleAttributes",
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicyRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicyRules",
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

func (p *LoadBalancerService) CreateServerCertificate(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateServerCertificate",
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

func (p *LoadBalancerService) DescribeServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeServerCertificates",
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

func (p *LoadBalancerService) ModifyServerCertificateAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyServerCertificateAttributes",
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

func (p *LoadBalancerService) DeleteServerCertificates(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteServerCertificates",
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

func init() { proto.RegisterFile("load_balancer.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x35, 0x09, 0x81, 0x38, 0x77, 0x38, 0x82, 0x55, 0x29, 0xe3, 0xcf, 0x40, 0x88, 0xab,
	0x4c, 0x62, 0x4f, 0xd0, 0x35, 0x05, 0x6d, 0x4a, 0xa1, 0x34, 0x1b, 0xe2, 0x0e, 0x39, 0xc9, 0x69,
	0x65, 0x91, 0xc6, 0x96, 0xed, 0x22, 0xa5, 0x8f, 0xc0, 0x53, 0xa3, 0xc4, 0xad, 0x14, 0xea, 0xf4,
	0xa2, 0x76, 0xef, 0x12, 0x27, 0xe7, 0xa7, 0xef, 0x7c, 0xe7, 0xb3, 0x65, 0x08, 0x4a, 0x4e, 0x8b,
	0x5f, 0x19, 0x2d, 0x69, 0x95, 0xa3, 0x8c, 0x84, 0xe4, 0x9a, 0x93, 0x27, 0x0a, 0xe5, 0x1f, 0x96,
	0x63, 0x38, 0x5c, 0x72, 0xbe, 0x2c, 0xf1, 0xaa, 0x5d, 0xce, 0xd6, 0x8b, 0x2b, 0x5c, 0x09, 0x5d,
	0x9b, 0xbf, 0x2e, 0xaf, 0xe1, 0x22, 0xe1, 0xb4, 0xb8, 0xd9, 0xd6, 0xa6, 0xa6, 0x66, 0x26, 0xb9,
	0x40, 0xa9, 0x19, 0x2a, 0x42, 0xe0, 0xd1, 0x86, 0x57, 0x38, 0x38, 0x7b, 0x73, 0xf6, 0xf1, 0xe9,
	0xbc, 0x7d, 0xfe, 0xf4, 0x37, 0x80, 0xa0, 0xa7, 0x8a, 0xc4, 0x40, 0xc6, 0x12, 0xa9, 0xc6, 0xee,
	0x47, 0xf2, 0x22, 0x32, 0x02, 0xa2, 0x9d, 0x80, 0x68, 0xd2, 0x08, 0x08, 0x0f, 0xac, 0x93, 0x2f,
	0xf0, 0x3c, 0x46, 0x95, 0x4b, 0x96, 0xfd, 0xc7, 0x51, 0x47, 0x83, 0x26, 0x10, 0xc4, 0x58, 0xa2,
	0xf6, 0xc4, 0x7c, 0x85, 0x97, 0x53, 0x5e, 0xb0, 0x45, 0xdd, 0xc5, 0x8c, 0xb4, 0x96, 0x2c, 0x5b,
	0x6b, 0x3c, 0x9e, 0x17, 0x03, 0x49, 0x35, 0x95, 0xda, 0x4f, 0xd5, 0x18, 0x9e, 0xa5, 0x9a, 0x0b,
	0x6f, 0x87, 0x1e, 0x44, 0x41, 0xb5, 0xbf, 0xd1, 0x73, 0x54, 0x6c, 0xe3, 0x89, 0x99, 0xc2, 0x70,
	0xa4, 0x14, 0xcf, 0x19, 0xd5, 0x38, 0x61, 0x42, 0xdd, 0x73, 0xaf, 0x1c, 0xcd, 0xe0, 0x55, 0xcc,
	0xba, 0xbc, 0xcf, 0x92, 0xaf, 0xbc, 0x88, 0x77, 0x30, 0x18, 0x15, 0x45, 0x17, 0x91, 0x30, 0xa5,
	0xb1, 0x72, 0x69, 0xf6, 0x1b, 0x5c, 0xf4, 0xa5, 0xdc, 0x1d, 0x38, 0x85, 0xa1, 0x9d, 0x76, 0x77,
	0xdc, 0x0f, 0x78, 0x6f, 0xa7, 0x7e, 0x87, 0xf3, 0x48, 0xff, 0x2d, 0x9c, 0xef, 0x79, 0x78, 0x43,
	0xf3, 0xdf, 0x58, 0x15, 0x4e, 0x1b, 0xb3, 0xcf, 0x42, 0x67, 0x5e, 0x02, 0xa1, 0xed, 0xa0, 0x33,
	0xed, 0x01, 0xde, 0xd9, 0x06, 0x6e, 0x69, 0x1e, 0xfe, 0xdd, 0xc1, 0xc0, 0x3e, 0x63, 0x67, 0xbc,
	0x64, 0x79, 0x7d, 0x2a, 0x03, 0x5b, 0x1a, 0x73, 0xd0, 0x76, 0x0f, 0x97, 0x76, 0xcb, 0x46, 0x9b,
	0x67, 0x62, 0x84, 0x28, 0xeb, 0x13, 0x34, 0xdc, 0x3b, 0x61, 0xe7, 0x76, 0x13, 0x08, 0xf7, 0xa2,
	0x6c, 0x64, 0xcd, 0xd7, 0xa5, 0x03, 0xed, 0x3b, 0xbc, 0x3e, 0x38, 0x0c, 0x47, 0xe4, 0x4f, 0xf8,
	0x70, 0x68, 0x1e, 0x0d, 0xd0, 0x63, 0x26, 0xed, 0xe9, 0xd5, 0x6b, 0xa4, 0xa3, 0xd4, 0x5b, 0x38,
	0x37, 0xb1, 0x6e, 0xee, 0x12, 0x28, 0xc7, 0xcd, 0xed, 0x63, 0xc1, 0x72, 0xaa, 0xd1, 0x6d, 0xc8,
	0xc6, 0x48, 0x0b, 0x76, 0xbc, 0xb0, 0x14, 0xde, 0x1a, 0x0f, 0x2d, 0x96, 0xdf, 0x26, 0x36, 0xf6,
	0xf9, 0x0b, 0xcc, 0x1e, 0xb7, 0xef, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x32, 0x80,
	0xb9, 0x05, 0x0a, 0x00, 0x00,
}
