// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: volume.proto

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

type VolumesServiceInterface interface {
	DescribeVolumes(in *DescribeVolumesInput) (out *DescribeVolumesOutput, err error)
	CreateVolumes(in *CreateVolumesInput) (out *CreateVolumesOutput, err error)
	DeleteVolumes(in *DeleteVolumesInput) (out *DeleteVolumesOutput, err error)
	AttachVolumes(in *AttachVolumesInput) (out *AttachVolumesOutput, err error)
	DetachVolumes(in *DetachVolumesInput) (out *DetachVolumesOutput, err error)
	ResizeVolumes(in *ResizeVolumesInput) (out *ResizeVolumesOutput, err error)
	ModifyVolumeAttributes(in *ModifyVolumeAttributesInput) (out *ModifyVolumeAttributesOutput, err error)
}

type VolumesService struct {
	Config           *config.Config
	Properties       *VolumesServiceProperties
	LastResponseBody string
}

func NewVolumesService(conf *config.Config, zone string) (p *VolumesService) {
	return &VolumesService{
		Config:     conf,
		Properties: &VolumesServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *VolumesService) DescribeVolumes(in *DescribeVolumesInput) (out *DescribeVolumesOutput, err error) {
	if in == nil {
		in = &DescribeVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeVolumes",
		RequestMethod: "GET",
	}

	x := &DescribeVolumesOutput{}
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

func (p *VolumesService) CreateVolumes(in *CreateVolumesInput) (out *CreateVolumesOutput, err error) {
	if in == nil {
		in = &CreateVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateVolumes",
		RequestMethod: "GET",
	}

	x := &CreateVolumesOutput{}
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

func (p *VolumesService) DeleteVolumes(in *DeleteVolumesInput) (out *DeleteVolumesOutput, err error) {
	if in == nil {
		in = &DeleteVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteVolumes",
		RequestMethod: "GET",
	}

	x := &DeleteVolumesOutput{}
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

func (p *VolumesService) AttachVolumes(in *AttachVolumesInput) (out *AttachVolumesOutput, err error) {
	if in == nil {
		in = &AttachVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AttachVolumes",
		RequestMethod: "GET",
	}

	x := &AttachVolumesOutput{}
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

func (p *VolumesService) DetachVolumes(in *DetachVolumesInput) (out *DetachVolumesOutput, err error) {
	if in == nil {
		in = &DetachVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DetachVolumes",
		RequestMethod: "GET",
	}

	x := &DetachVolumesOutput{}
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

func (p *VolumesService) ResizeVolumes(in *ResizeVolumesInput) (out *ResizeVolumesOutput, err error) {
	if in == nil {
		in = &ResizeVolumesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeVolumes",
		RequestMethod: "GET",
	}

	x := &ResizeVolumesOutput{}
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

func (p *VolumesService) ModifyVolumeAttributes(in *ModifyVolumeAttributesInput) (out *ModifyVolumeAttributesOutput, err error) {
	if in == nil {
		in = &ModifyVolumeAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyVolumeAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyVolumeAttributesOutput{}
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
