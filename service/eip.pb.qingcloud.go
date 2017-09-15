// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: eip.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/logger"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

// See https://docs.qingcloud.com/api/eip/index.html
type EIPServiceInterface interface {
	DescribeEips(in *DescribeEipsInput) (out *DescribeEipsOutput, err error)
	AllocateEips(in *AllocateEipsInput) (out *AllocateEipsOutput, err error)
	ReleaseEips(in *ReleaseEipsInput) (out *ReleaseEipsOutput, err error)
	AssociateEip(in *AssociateEipInput) (out *AssociateEipOutput, err error)
	DissociateEips(in *DissociateEipsInput) (out *DissociateEipsOutput, err error)
	ChangeEipsBandwidth(in *ChangeEipsBandwidthInput) (out *ChangeEipsBandwidthOutput, err error)
	ChangeEipsBillingMode(in *ChangeEipsBillingModeInput) (out *ChangeEipsBillingModeOutput, err error)
	ModifyEipAttributes(in *ModifyEipAttributesInput) (out *ModifyEipAttributesOutput, err error)
}

// See https://docs.qingcloud.com/api/eip/index.html
type EIPService struct {
	Config           *config.Config
	Properties       *EIPServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/eip/index.html
func NewEIPService(conf *config.Config, zone string) (p *EIPService) {
	return &EIPService{
		Config:     conf,
		Properties: &EIPServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/eip/index.html
func (s *QingCloudService) EIP(zone string) (*EIPService, error) {
	properties := &EIPServiceProperties{
		Zone: zone,
	}

	return &EIPService{Config: s.Config, Properties: properties}, nil
}

func (p *EIPService) DescribeEips(in *DescribeEipsInput) (out *DescribeEipsOutput, err error) {
	if in == nil {
		in = &DescribeEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeEipsOutput{}
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

func (p *EIPService) AllocateEips(in *AllocateEipsInput) (out *AllocateEipsOutput, err error) {
	if in == nil {
		in = &AllocateEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AllocateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &AllocateEipsOutput{}
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

func (p *EIPService) ReleaseEips(in *ReleaseEipsInput) (out *ReleaseEipsOutput, err error) {
	if in == nil {
		in = &ReleaseEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ReleaseEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &ReleaseEipsOutput{}
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

func (p *EIPService) AssociateEip(in *AssociateEipInput) (out *AssociateEipOutput, err error) {
	if in == nil {
		in = &AssociateEipInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateEip",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateEipOutput{}
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

func (p *EIPService) DissociateEips(in *DissociateEipsInput) (out *DissociateEipsOutput, err error) {
	if in == nil {
		in = &DissociateEipsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateEips",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateEipsOutput{}
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

func (p *EIPService) ChangeEipsBandwidth(in *ChangeEipsBandwidthInput) (out *ChangeEipsBandwidthOutput, err error) {
	if in == nil {
		in = &ChangeEipsBandwidthInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBandwidth",
		RequestMethod: "GET", // GET or POST
	}

	x := &ChangeEipsBandwidthOutput{}
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

func (p *EIPService) ChangeEipsBillingMode(in *ChangeEipsBillingModeInput) (out *ChangeEipsBillingModeOutput, err error) {
	if in == nil {
		in = &ChangeEipsBillingModeInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeEipsBillingMode",
		RequestMethod: "GET", // GET or POST
	}

	x := &ChangeEipsBillingModeOutput{}
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

func (p *EIPService) ModifyEipAttributes(in *ModifyEipAttributesInput) (out *ModifyEipAttributesOutput, err error) {
	if in == nil {
		in = &ModifyEipAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyEipAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyEipAttributesOutput{}
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

func (p *EIPServiceProperties) Validate() error {
	return nil
}

func (p *DescribeEipsInput) Validate() error {
	return nil
}

func (p *DescribeEipsOutput) Validate() error {
	return nil
}

func (p *AllocateEipsInput) Validate() error {
	return nil
}

func (p *AllocateEipsOutput) Validate() error {
	return nil
}

func (p *ReleaseEipsInput) Validate() error {
	return nil
}

func (p *ReleaseEipsOutput) Validate() error {
	return nil
}

func (p *AssociateEipInput) Validate() error {
	return nil
}

func (p *AssociateEipOutput) Validate() error {
	return nil
}

func (p *DissociateEipsInput) Validate() error {
	return nil
}

func (p *DissociateEipsOutput) Validate() error {
	return nil
}

func (p *ChangeEipsBandwidthInput) Validate() error {
	return nil
}

func (p *ChangeEipsBandwidthOutput) Validate() error {
	return nil
}

func (p *ChangeEipsBillingModeInput) Validate() error {
	return nil
}

func (p *ChangeEipsBillingModeOutput) Validate() error {
	return nil
}

func (p *ModifyEipAttributesInput) Validate() error {
	return nil
}

func (p *ModifyEipAttributesOutput) Validate() error {
	return nil
}
