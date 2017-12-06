// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: subuser.proto

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

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	Config           *config.Config
	Properties       *SubuserServiceProperties
	LastResponseBody string
}

func NewSubuserService(conf *config.Config, zone string) (p *SubuserService) {
	return &SubuserService{
		Config:     conf,
		Properties: &SubuserServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Subuser(zone string) (*SubuserService, error) {
	properties := &SubuserServiceProperties{
		Zone: proto.String(zone),
	}

	return &SubuserService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *SubuserService) DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error) {
	if in == nil {
		in = &DescribeSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSubUsers",
		RequestMethod: "GET",
	}

	x := &DescribeSubUsersOutput{}
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

func (p *SubuserService) CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error) {
	if in == nil {
		in = &CreateSubUserInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSubUser",
		RequestMethod: "GET",
	}

	x := &CreateSubUserOutput{}
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

func (p *SubuserService) ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error) {
	if in == nil {
		in = &ModifySubUserAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySubUserAttributes",
		RequestMethod: "GET",
	}

	x := &ModifySubUserAttributesOutput{}
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

func (p *SubuserService) DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error) {
	if in == nil {
		in = &DeleteSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSubUsers",
		RequestMethod: "GET",
	}

	x := &DeleteSubUsersOutput{}
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

func (p *SubuserService) RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error) {
	if in == nil {
		in = &RestoreSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RestoreSubUsers",
		RequestMethod: "GET",
	}

	x := &RestoreSubUsersOutput{}
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
