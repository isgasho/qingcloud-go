// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: subuser.proto

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

// See https://docs.qingcloud.com/api/subuser/index.html
type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

// See https://docs.qingcloud.com/api/subuser/index.html
type SubuserService struct {
	Config           *config.Config
	Properties       *SubuserServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/subuser/index.html
func NewSubuserService(conf *config.Config, zone string) (p *SubuserService) {
	return &SubuserService{
		Config:     conf,
		Properties: &SubuserServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/subuser/index.html
func (s *QingCloudService) Subuser(zone string) (*SubuserService, error) {
	properties := &SubuserServiceProperties{
		Zone: zone,
	}

	return &SubuserService{Config: s.Config, Properties: properties}, nil
}

func (p *SubuserService) DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error) {
	if in == nil {
		in = &DescribeSubUsersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSubUsers",
		RequestMethod: "GET", // GET or POST
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
		RequestMethod: "GET", // GET or POST
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
		RequestMethod: "GET", // GET or POST
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
		RequestMethod: "GET", // GET or POST
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
		RequestMethod: "GET", // GET or POST
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

func (p *SubuserServiceProperties) Validate() error {
	return nil
}

func (p *DescribeSubUsersInput) Validate() error {
	return nil
}

func (p *DescribeSubUsersOutput) Validate() error {
	return nil
}

func (p *CreateSubUserInput) Validate() error {
	return nil
}

func (p *CreateSubUserOutput) Validate() error {
	return nil
}

func (p *ModifySubUserAttributesInput) Validate() error {
	return nil
}

func (p *ModifySubUserAttributesOutput) Validate() error {
	return nil
}

func (p *DeleteSubUsersInput) Validate() error {
	return nil
}

func (p *DeleteSubUsersOutput) Validate() error {
	return nil
}

func (p *RestoreSubUsersInput) Validate() error {
	return nil
}

func (p *RestoreSubUsersOutput) Validate() error {
	return nil
}