// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: instances.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}

type InstanceServiceInterface interface {
	DescribeInstances(in *DescribeInstancesInput) (out *DescribeInstancesOutput, err error)
	RunInstances(in *RunInstancesInput) (out *RunInstancesOutput, err error)
	TerminateInstances(in *TerminateInstancesInput) (out *TerminateInstancesOutput, err error)
	StartInstances(in *StartInstancesInput) (out *StartInstancesOutput, err error)
	StopInstances(in *StopInstancesInput) (out *StopInstancesOutput, err error)
	RestartInstances(in *RestartInstancesInput) (out *RestartInstancesOutput, err error)
	ResetInstances(in *ResetInstancesInput) (out *ResetInstancesOutput, err error)
	ResizeInstances(in *ResizeInstancesInput) (out *ResizeInstancesOutput, err error)
	ModifyInstanceAttributes(in *ModifyInstanceAttributesInput) (out *ModifyInstanceAttributesOutput, err error)
	DescribeInstanceTypes(in *DescribeInstanceTypesInput) (out *DescribeInstanceTypesOutput, err error)
	CreateBrokers(in *CreateBrokersInput) (out *CreateBrokersOutput, err error)
	DeleteBrokers(in *DeleteBrokersInput) (out *DeleteBrokersOutput, err error)
}

type InstanceService struct {
	Config           *config.Config
	Properties       *InstanceServiceProperties
	LastResponseBody string
}

func NewInstanceService(conf *config.Config, zone string) (p *InstanceService) {
	return &InstanceService{
		Config:     conf,
		Properties: &InstanceServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *InstanceService) DescribeInstances(in *DescribeInstancesInput) (out *DescribeInstancesOutput, err error) {
	if in == nil {
		in = &DescribeInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeInstances",
		RequestMethod: "GET",
	}

	x := &DescribeInstancesOutput{}
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

func (p *InstanceService) RunInstances(in *RunInstancesInput) (out *RunInstancesOutput, err error) {
	if in == nil {
		in = &RunInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RunInstances",
		RequestMethod: "GET",
	}

	x := &RunInstancesOutput{}
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

func (p *InstanceService) TerminateInstances(in *TerminateInstancesInput) (out *TerminateInstancesOutput, err error) {
	if in == nil {
		in = &TerminateInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "TerminateInstances",
		RequestMethod: "GET",
	}

	x := &TerminateInstancesOutput{}
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

func (p *InstanceService) StartInstances(in *StartInstancesInput) (out *StartInstancesOutput, err error) {
	if in == nil {
		in = &StartInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartInstances",
		RequestMethod: "GET",
	}

	x := &StartInstancesOutput{}
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

func (p *InstanceService) StopInstances(in *StopInstancesInput) (out *StopInstancesOutput, err error) {
	if in == nil {
		in = &StopInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopInstances",
		RequestMethod: "GET",
	}

	x := &StopInstancesOutput{}
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

func (p *InstanceService) RestartInstances(in *RestartInstancesInput) (out *RestartInstancesOutput, err error) {
	if in == nil {
		in = &RestartInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestartInstances",
		RequestMethod: "GET",
	}

	x := &RestartInstancesOutput{}
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

func (p *InstanceService) ResetInstances(in *ResetInstancesInput) (out *ResetInstancesOutput, err error) {
	if in == nil {
		in = &ResetInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResetInstances",
		RequestMethod: "GET",
	}

	x := &ResetInstancesOutput{}
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

func (p *InstanceService) ResizeInstances(in *ResizeInstancesInput) (out *ResizeInstancesOutput, err error) {
	if in == nil {
		in = &ResizeInstancesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeInstances",
		RequestMethod: "GET",
	}

	x := &ResizeInstancesOutput{}
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

func (p *InstanceService) ModifyInstanceAttributes(in *ModifyInstanceAttributesInput) (out *ModifyInstanceAttributesOutput, err error) {
	if in == nil {
		in = &ModifyInstanceAttributesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyInstanceAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyInstanceAttributesOutput{}
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

func (p *InstanceService) DescribeInstanceTypes(in *DescribeInstanceTypesInput) (out *DescribeInstanceTypesOutput, err error) {
	if in == nil {
		in = &DescribeInstanceTypesInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeInstanceTypes",
		RequestMethod: "GET",
	}

	x := &DescribeInstanceTypesOutput{}
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

func (p *InstanceService) CreateBrokers(in *CreateBrokersInput) (out *CreateBrokersOutput, err error) {
	if in == nil {
		in = &CreateBrokersInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateBrokers",
		RequestMethod: "GET",
	}

	x := &CreateBrokersOutput{}
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

func (p *InstanceService) DeleteBrokers(in *DeleteBrokersInput) (out *DeleteBrokersOutput, err error) {
	if in == nil {
		in = &DeleteBrokersInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteBrokers",
		RequestMethod: "GET",
	}

	x := &DeleteBrokersOutput{}
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
