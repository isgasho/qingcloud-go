// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: vxnet.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"
import "github.com/chai2010/qingcloud-go/pkg/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}
var _ = data.Operation{}

type VxnetServiceInterface interface {
	DescribeVxnets(in *DescribeVxnetsInput) (out *DescribeVxnetsOutput, err error)
	CreateVxnets(in *CreateVxnetsInput) (out *CreateVxnetsOutput, err error)
	DeleteVxnets(in *DeleteVxnetsInput) (out *DeleteVxnetsOutput, err error)
	JoinVxnet(in *JoinVxnetInput) (out *JoinVxnetOutput, err error)
	LeaveVxnet(in *LeaveVxnetInput) (out *LeaveVxnetOutput, err error)
	ModifyVxnetAttributes(in *ModifyVxnetAttributesInput) (out *ModifyVxnetAttributesOutput, err error)
	DescribeVxnetInstances(in *DescribeVxnetInstancesInput) (out *DescribeVxnetInstancesOutput, err error)
}

type VxnetService struct {
	Config           *config.Config
	Properties       *VxnetServiceProperties
	LastResponseBody string
}

func NewVxnetService(conf *config.Config, zone string) (p *VxnetService) {
	return &VxnetService{
		Config:     conf,
		Properties: &VxnetServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Vxnet(zone string) (*VxnetService, error) {
	properties := &VxnetServiceProperties{
		Zone: proto.String(zone),
	}

	return &VxnetService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *VxnetService) DescribeVxnets(in *DescribeVxnetsInput) (out *DescribeVxnetsOutput, err error) {
	if in == nil {
		in = &DescribeVxnetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnets",
		RequestMethod: "GET",
	}

	x := &DescribeVxnetsOutput{}
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

func (p *VxnetService) CreateVxnets(in *CreateVxnetsInput) (out *CreateVxnetsOutput, err error) {
	if in == nil {
		in = &CreateVxnetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVxnets",
		RequestMethod: "GET",
	}

	x := &CreateVxnetsOutput{}
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

func (p *VxnetService) DeleteVxnets(in *DeleteVxnetsInput) (out *DeleteVxnetsOutput, err error) {
	if in == nil {
		in = &DeleteVxnetsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVxnets",
		RequestMethod: "GET",
	}

	x := &DeleteVxnetsOutput{}
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

func (p *VxnetService) JoinVxnet(in *JoinVxnetInput) (out *JoinVxnetOutput, err error) {
	if in == nil {
		in = &JoinVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "JoinVxnet",
		RequestMethod: "GET",
	}

	x := &JoinVxnetOutput{}
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

func (p *VxnetService) LeaveVxnet(in *LeaveVxnetInput) (out *LeaveVxnetOutput, err error) {
	if in == nil {
		in = &LeaveVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "LeaveVxnet",
		RequestMethod: "GET",
	}

	x := &LeaveVxnetOutput{}
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

func (p *VxnetService) ModifyVxnetAttributes(in *ModifyVxnetAttributesInput) (out *ModifyVxnetAttributesOutput, err error) {
	if in == nil {
		in = &ModifyVxnetAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVxnetAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyVxnetAttributesOutput{}
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

func (p *VxnetService) DescribeVxnetInstances(in *DescribeVxnetInstancesInput) (out *DescribeVxnetInstancesOutput, err error) {
	if in == nil {
		in = &DescribeVxnetInstancesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVxnetInstances",
		RequestMethod: "GET",
	}

	x := &DescribeVxnetInstancesOutput{}
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