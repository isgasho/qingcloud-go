// Code generated by protoc-gen-go. DO NOT EDIT.
// source: load_balancer.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

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

type DescribeLoadBalancersInput struct {
	Limit         int32    `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Loadbalancers []string `protobuf:"bytes,2,rep,name=loadbalancers" json:"loadbalancers,omitempty"`
	Offset        int32    `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	SearchWord    string   `protobuf:"bytes,4,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Status        []string `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	Tags          []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Verbose       int32    `protobuf:"varint,7,opt,name=verbose" json:"verbose,omitempty"`
}

func (m *DescribeLoadBalancersInput) Reset()                    { *m = DescribeLoadBalancersInput{} }
func (m *DescribeLoadBalancersInput) String() string            { return proto.CompactTextString(m) }
func (*DescribeLoadBalancersInput) ProtoMessage()               {}
func (*DescribeLoadBalancersInput) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *DescribeLoadBalancersInput) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeLoadBalancersInput) GetLoadbalancers() []string {
	if m != nil {
		return m.Loadbalancers
	}
	return nil
}

func (m *DescribeLoadBalancersInput) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeLoadBalancersInput) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *DescribeLoadBalancersInput) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeLoadBalancersInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DescribeLoadBalancersInput) GetVerbose() int32 {
	if m != nil {
		return m.Verbose
	}
	return 0
}

type DescribeLoadBalancersOutput struct {
	Action          string          `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	RetCode         int32           `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Message         string          `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	LoadbalancerSet []*LoadBalancer `protobuf:"bytes,4,rep,name=loadbalancer_set,json=loadbalancerSet" json:"loadbalancer_set,omitempty"`
}

func (m *DescribeLoadBalancersOutput) Reset()                    { *m = DescribeLoadBalancersOutput{} }
func (m *DescribeLoadBalancersOutput) String() string            { return proto.CompactTextString(m) }
func (*DescribeLoadBalancersOutput) ProtoMessage()               {}
func (*DescribeLoadBalancersOutput) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *DescribeLoadBalancersOutput) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DescribeLoadBalancersOutput) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *DescribeLoadBalancersOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DescribeLoadBalancersOutput) GetLoadbalancerSet() []*LoadBalancer {
	if m != nil {
		return m.LoadbalancerSet
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadBalancerServiceProperties)(nil), "service.LoadBalancerServiceProperties")
	proto.RegisterType((*DescribeLoadBalancersInput)(nil), "service.DescribeLoadBalancersInput")
	proto.RegisterType((*DescribeLoadBalancersOutput)(nil), "service.DescribeLoadBalancersOutput")
}

type LoadBalancerServiceInterface interface {
	CreateLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeLoadBalancers(in *DescribeLoadBalancersInput) (out *DescribeLoadBalancersOutput, err error)
	DeleteLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyLoadBalancerAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	StartLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	StopLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	UpdateLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ResizeLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AssociateEipsToLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DissociateEipsFromLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyLoadBalancerListenerAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyLoadBalancerBackendAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateLoadBalancerPolicy(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeLoadBalancerPolicies(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyLoadBalancerPolicyAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ApplyLoadBalancerPolicy(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteLoadBalancerPolicies(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateServerCertificate(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeServerCertificates(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyServerCertificateAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteServerCertificates(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
}

type LoadBalancerService struct {
	Config           *config.Config
	Properties       *LoadBalancerServiceProperties
	LastResponseBody string
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

func (p *LoadBalancerService) CreateLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancer",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeLoadBalancers(in *DescribeLoadBalancersInput) (out *DescribeLoadBalancersOutput, err error) {
	if in == nil {
		in = &DescribeLoadBalancersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeLoadBalancersOutput{}
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

func (p *DescribeLoadBalancersInput) Validate() error {
	return nil
}

func (p *LoadBalancerService) DeleteLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyLoadBalancerAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) StartLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) StopLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) UpdateLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UpdateLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ResizeLoadBalancers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeLoadBalancers",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) AssociateEipsToLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEipsToLoadBalancer",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DissociateEipsFromLoadBalancer(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEipsFromLoadBalancer",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) AddLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerListeners",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerListeners",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DeleteLoadBalancerListeners(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerListeners",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyLoadBalancerListenerAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerListenerAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) AddLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerBackends",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerBackends",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DeleteLoadBalancerBackends(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerBackends",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyLoadBalancerBackendAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerBackendAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) CreateLoadBalancerPolicy(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateLoadBalancerPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicies(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ApplyLoadBalancerPolicy(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyLoadBalancerPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicies(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) AddLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddLoadBalancerPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeLoadBalancerPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyLoadBalancerPolicyRuleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyLoadBalancerPolicyRuleAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DeleteLoadBalancerPolicyRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteLoadBalancerPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) CreateServerCertificate(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateServerCertificate",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DescribeServerCertificates(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeServerCertificates",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) ModifyServerCertificateAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyServerCertificateAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func (p *LoadBalancerService) DeleteServerCertificates(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteServerCertificates",
		RequestMethod: "GET", // GET or POST
	}

	x := &google_protobuf2.Empty{}
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

func init() { proto.RegisterFile("load_balancer.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x61, 0x4e, 0xdb, 0x30,
	0x18, 0x55, 0xa1, 0x94, 0xf5, 0x63, 0xd3, 0x36, 0x77, 0x40, 0x08, 0xeb, 0x60, 0x05, 0x4d, 0xfc,
	0x2a, 0x12, 0x5c, 0x60, 0xd0, 0x76, 0x12, 0xa8, 0x0c, 0x48, 0x61, 0xdb, 0xbf, 0xca, 0x49, 0xbe,
	0x76, 0x16, 0x69, 0x9c, 0xd9, 0x0e, 0x53, 0x39, 0xc2, 0xce, 0xb2, 0x63, 0xec, 0x0e, 0x3b, 0xc7,
	0x6e, 0x30, 0x39, 0x49, 0xa7, 0xd0, 0xa6, 0x93, 0x48, 0xf8, 0x53, 0xe5, 0x73, 0xec, 0xe7, 0xf7,
	0xde, 0xf7, 0xe2, 0x1a, 0x6a, 0x1e, 0xa7, 0x6e, 0xdf, 0xa6, 0x1e, 0xf5, 0x1d, 0x14, 0xcd, 0x40,
	0x70, 0xc5, 0xc9, 0xb2, 0x44, 0x71, 0xcb, 0x1c, 0x34, 0x57, 0xd4, 0x38, 0x40, 0x19, 0x8f, 0x9a,
	0xf5, 0x6f, 0xcc, 0x1f, 0x3a, 0x1e, 0x0f, 0xdd, 0xbe, 0x74, 0x6f, 0xfa, 0x22, 0xf4, 0x70, 0x5f,
	0xff, 0x24, 0xaf, 0x37, 0x87, 0x9c, 0x0f, 0x3d, 0xdc, 0x8f, 0x2a, 0x3b, 0x1c, 0xec, 0xe3, 0x28,
	0x50, 0xe3, 0xf8, 0x65, 0xe3, 0x10, 0xea, 0x5d, 0x4e, 0xdd, 0xe3, 0x64, 0x9f, 0x5e, 0x8c, 0x7f,
	0x21, 0x78, 0x80, 0x42, 0x31, 0x94, 0x84, 0x40, 0xf9, 0x8e, 0xfb, 0x68, 0x94, 0xb6, 0x4b, 0x7b,
	0x55, 0x2b, 0x7a, 0x6e, 0xfc, 0x2e, 0x81, 0xd9, 0x46, 0xe9, 0x08, 0x66, 0x63, 0x7a, 0xb5, 0x3c,
	0xf1, 0x83, 0x50, 0x91, 0x57, 0xb0, 0xe4, 0xb1, 0x11, 0x53, 0xd1, 0x9a, 0x25, 0x2b, 0x2e, 0xc8,
	0x2e, 0x3c, 0xd3, 0x92, 0x26, 0x8a, 0xa4, 0xb1, 0xb0, 0xbd, 0xb8, 0x57, 0xb5, 0xee, 0x0f, 0x92,
	0x35, 0xa8, 0xf0, 0xc1, 0x40, 0xa2, 0x32, 0x16, 0xa3, 0xc5, 0x49, 0x45, 0xb6, 0x60, 0x45, 0x22,
	0x15, 0xce, 0xd7, 0xfe, 0x77, 0x2e, 0x5c, 0xa3, 0x1c, 0xb1, 0x81, 0x78, 0xe8, 0x33, 0x17, 0xae,
	0x5e, 0x28, 0x15, 0x55, 0xa1, 0x34, 0x96, 0x22, 0xdc, 0xa4, 0xd2, 0xfc, 0x15, 0x1d, 0x4a, 0xa3,
	0x12, 0x8d, 0x46, 0xcf, 0xc4, 0x80, 0xe5, 0x5b, 0x14, 0x36, 0x97, 0x68, 0x2c, 0x47, 0xbb, 0x4c,
	0xca, 0xc6, 0xcf, 0x12, 0x6c, 0x66, 0x2a, 0x3b, 0x0f, 0x95, 0x96, 0xb6, 0x06, 0x15, 0xea, 0x28,
	0xc6, 0xfd, 0xc4, 0x8f, 0xa4, 0x22, 0x1b, 0xf0, 0x44, 0xa0, 0xea, 0x3b, 0xdc, 0x45, 0x63, 0x21,
	0x86, 0x14, 0xa8, 0x5a, 0xdc, 0x45, 0xbd, 0xd9, 0x08, 0xa5, 0xa4, 0x43, 0x8c, 0x24, 0x55, 0xad,
	0x49, 0x49, 0xde, 0xc3, 0x8b, 0xb4, 0xf8, 0xbe, 0x56, 0x5d, 0xde, 0x5e, 0xdc, 0x5b, 0x39, 0x58,
	0x6d, 0x26, 0x8d, 0x6e, 0xa6, 0x49, 0x58, 0xcf, 0xd3, 0xd3, 0x7b, 0xa8, 0x0e, 0x7e, 0xd5, 0xa0,
	0x96, 0xd1, 0x3e, 0xd2, 0x06, 0xd2, 0x12, 0x48, 0xd5, 0x3d, 0x0d, 0x64, 0xad, 0x19, 0x27, 0xa1,
	0x39, 0x49, 0x42, 0xb3, 0xa3, 0x93, 0x60, 0xce, 0x19, 0x27, 0x36, 0xac, 0x66, 0x7a, 0x41, 0x76,
	0xfe, 0xd1, 0x9b, 0x9f, 0x02, 0x73, 0xf7, 0xff, 0x93, 0x12, 0x43, 0x3b, 0x50, 0x6b, 0xa3, 0x87,
	0x6a, 0x6a, 0x87, 0x87, 0x52, 0xfd, 0x08, 0xaf, 0xcf, 0xb8, 0xcb, 0x06, 0xe3, 0x34, 0xcc, 0x91,
	0x52, 0x82, 0xd9, 0xa1, 0xc2, 0x87, 0xe3, 0xb5, 0x81, 0xf4, 0x14, 0x15, 0xaa, 0x18, 0xab, 0x16,
	0xbc, 0xec, 0x29, 0x1e, 0x14, 0x03, 0xe9, 0x40, 0xed, 0x3a, 0x70, 0x69, 0x51, 0x87, 0x3a, 0x50,
	0xb3, 0x50, 0xb2, 0xbb, 0x82, 0x30, 0x67, 0xb0, 0x79, 0x24, 0x25, 0x77, 0x18, 0x55, 0xd8, 0x61,
	0x81, 0xbc, 0xe2, 0x85, 0x22, 0x76, 0x01, 0x6f, 0xda, 0x2c, 0x8d, 0xf7, 0x41, 0xf0, 0x51, 0x21,
	0xc4, 0x53, 0x30, 0x8e, 0x5c, 0x37, 0x0d, 0xd1, 0x65, 0x52, 0xa1, 0x9f, 0x47, 0xec, 0x39, 0xd4,
	0xb3, 0xb2, 0x9b, 0x1f, 0xf0, 0x4c, 0x9f, 0x2e, 0xd3, 0x69, 0xcf, 0x0f, 0xf7, 0x09, 0x76, 0x67,
	0x53, 0x3f, 0x81, 0x2b, 0x90, 0xfe, 0x13, 0x58, 0x9f, 0xf2, 0xf0, 0x98, 0x3a, 0x37, 0xe8, 0xbb,
	0xb9, 0x3e, 0xcc, 0x2c, 0x0b, 0x73, 0xe3, 0x75, 0xf5, 0x3f, 0xcf, 0xb4, 0x83, 0xb9, 0xd1, 0xae,
	0x61, 0x67, 0xd6, 0xc0, 0x04, 0xad, 0x80, 0x7f, 0xa7, 0x60, 0xcc, 0x1e, 0xbf, 0x17, 0xdc, 0x63,
	0xce, 0xf8, 0xb1, 0x0c, 0x8c, 0xd0, 0x58, 0x0e, 0x6e, 0x57, 0xd0, 0x98, 0x95, 0x1c, 0x73, 0x2b,
	0x98, 0x98, 0x20, 0xf0, 0xc6, 0x8f, 0x20, 0x38, 0xb3, 0xc3, 0xb9, 0xe5, 0x76, 0xc1, 0x9c, 0x8a,
	0x72, 0x4c, 0xcb, 0x0a, 0xbd, 0x1c, 0x68, 0x97, 0xb0, 0x35, 0xb7, 0x19, 0x39, 0x21, 0xbf, 0xc0,
	0xbb, 0x79, 0xfd, 0xd0, 0x80, 0x05, 0x7a, 0x12, 0x9d, 0x5e, 0x99, 0x46, 0xe6, 0xa4, 0x7a, 0x02,
	0xeb, 0x71, 0xac, 0xf5, 0x35, 0x03, 0x45, 0x4b, 0xdf, 0x10, 0x07, 0xcc, 0xa1, 0x0a, 0xf3, 0x35,
	0x39, 0x36, 0x72, 0x06, 0xec, 0xe1, 0xc4, 0x7a, 0xf0, 0x36, 0xf6, 0x70, 0x06, 0xab, 0xd8, 0x47,
	0x1c, 0xdb, 0x57, 0x9c, 0xa0, 0x59, 0xff, 0xf1, 0xa7, 0xbc, 0x01, 0xd5, 0x4b, 0xe6, 0x0f, 0x5b,
	0xfa, 0x9e, 0x4e, 0x9e, 0xa6, 0x7b, 0x62, 0x57, 0xa2, 0xe9, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0x7f, 0xb3, 0x41, 0xf4, 0x0b, 0x00, 0x00,
}
