// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_acl.proto

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

type ResourceACLServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *ResourceACLServiceProperties) Reset()                    { *m = ResourceACLServiceProperties{} }
func (m *ResourceACLServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*ResourceACLServiceProperties) ProtoMessage()               {}
func (*ResourceACLServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *ResourceACLServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceACLServiceProperties)(nil), "service.ResourceACLServiceProperties")
}

type ResourceACLServiceInterface interface {
	DescribeSharedResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyResourceGroupAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyUserGroupAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyUserGroupMemberAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	CreateGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyGroupRoleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	AddGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	ModifyGroupRoleRuleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DeleteGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	GrantResourceGroupsToUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	RevokeResourceGroupsFromUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
	DescribeResourceUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error)
}

type ResourceACLService struct {
	Config           *config.Config
	Properties       *ResourceACLServiceProperties
	LastResponseBody string
}

func NewResourceACLService(conf *config.Config, zone string) (p *ResourceACLService) {
	return &ResourceACLService{
		Config:     conf,
		Properties: &ResourceACLServiceProperties{Zone: zone},
	}
}

func (s *QingCloudService) ResourceACL(zone string) (*ResourceACLService, error) {
	properties := &ResourceACLServiceProperties{
		Zone: zone,
	}

	return &ResourceACLService{Config: s.Config, Properties: properties}, nil
}

func (p *ResourceACLService) DescribeSharedResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSharedResourceGroups",
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

func (p *ResourceACLService) DescribeResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceGroups",
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

func (p *ResourceACLService) CreateResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateResourceGroups",
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

func (p *ResourceACLService) ModifyResourceGroupAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyResourceGroupAttributes",
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

func (p *ResourceACLService) DeleteResourceGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteResourceGroups",
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

func (p *ResourceACLService) DescribeResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceGroupItems",
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

func (p *ResourceACLService) AddResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddResourceGroupItems",
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

func (p *ResourceACLService) DeleteResourceGroupItems(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteResourceGroupItems",
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

func (p *ResourceACLService) DescribeUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeUserGroups",
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

func (p *ResourceACLService) CreateUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateUserGroups",
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

func (p *ResourceACLService) ModifyUserGroupAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyUserGroupAttributes",
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

func (p *ResourceACLService) DeleteUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteUserGroups",
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

func (p *ResourceACLService) DescribeUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeUserGroupMembers",
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

func (p *ResourceACLService) AddUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddUserGroupMembers",
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

func (p *ResourceACLService) ModifyUserGroupMemberAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyUserGroupMemberAttributes",
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

func (p *ResourceACLService) DeleteUserGroupMembers(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteUserGroupMembers",
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

func (p *ResourceACLService) DescribeGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeGroupRoles",
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

func (p *ResourceACLService) CreateGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateGroupRoles",
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

func (p *ResourceACLService) ModifyGroupRoleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyGroupRoleAttributes",
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

func (p *ResourceACLService) DeleteGroupRoles(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteGroupRoles",
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

func (p *ResourceACLService) DescribeGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeGroupRoleRules",
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

func (p *ResourceACLService) AddGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddGroupRoleRules",
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

func (p *ResourceACLService) ModifyGroupRoleRuleAttributes(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyGroupRoleRuleAttributes",
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

func (p *ResourceACLService) DeleteGroupRoleRules(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteGroupRoleRules",
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

func (p *ResourceACLService) GrantResourceGroupsToUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GrantResourceGroupsToUserGroups",
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

func (p *ResourceACLService) RevokeResourceGroupsFromUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeResourceGroupsFromUserGroups",
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

func (p *ResourceACLService) DescribeResourceUserGroups(in *google_protobuf2.Empty) (out *google_protobuf2.Empty, err error) {
	if in == nil {
		in = &google_protobuf2.Empty{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceUserGroups",
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

func init() { proto.RegisterFile("resource_acl.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x73, 0x93, 0x1b, 0xcd, 0x3d, 0x6e, 0x70, 0xd4, 0x1b, 0xac, 0xdc, 0x60, 0x58, 0xb9,
	0x2a, 0x09, 0x3e, 0x41, 0x29, 0x7f, 0xfc, 0x03, 0x2a, 0x05, 0xd7, 0xa4, 0xed, 0x1c, 0x6a, 0x43,
	0xdb, 0xa9, 0xa7, 0x53, 0x12, 0x7c, 0x04, 0xf7, 0x3e, 0xa8, 0x6f, 0x60, 0xe8, 0xd8, 0x86, 0x22,
	0x2e, 0xe8, 0xb0, 0x21, 0xb4, 0x3d, 0xfd, 0xe5, 0x9b, 0x2f, 0x5f, 0xbf, 0x19, 0x60, 0x84, 0x99,
	0xc8, 0xc9, 0xc7, 0xb5, 0xeb, 0x47, 0x66, 0x4a, 0x42, 0x0a, 0xf6, 0x38, 0x43, 0xda, 0x85, 0x3e,
	0x1a, 0x0f, 0xdf, 0xc3, 0x24, 0xf0, 0x23, 0x91, 0xf3, 0x75, 0xc6, 0xb7, 0x6b, 0xca, 0x23, 0xec,
	0x1f, 0x7e, 0xd4, 0x9c, 0xf1, 0x2a, 0x10, 0x22, 0x88, 0xb0, 0x5f, 0x5c, 0x79, 0xf9, 0xa6, 0x8f,
	0x71, 0x2a, 0xf7, 0xea, 0x61, 0x6f, 0x00, 0x1d, 0xe7, 0x2f, 0xda, 0xb2, 0x67, 0x4b, 0x45, 0xfc,
	0x42, 0x22, 0x45, 0x92, 0x21, 0x66, 0x8c, 0xc1, 0xed, 0x0f, 0x91, 0x60, 0xfb, 0xe6, 0xf5, 0xcd,
	0x9b, 0x3b, 0xa7, 0xf8, 0x3f, 0xf8, 0xd5, 0x02, 0xf6, 0xef, 0x4b, 0xec, 0x13, 0x74, 0x46, 0x98,
	0xf9, 0x14, 0x7a, 0xb8, 0xfc, 0xe6, 0x12, 0xf2, 0x72, 0x66, 0x4a, 0x22, 0x4f, 0x33, 0x76, 0x6f,
	0x2a, 0x21, 0x66, 0x29, 0xc4, 0x1c, 0x1f, 0x84, 0x18, 0xff, 0xb9, 0xcf, 0xde, 0xc1, 0x7d, 0xc9,
	0xd3, 0x24, 0x4d, 0xe0, 0xb9, 0x4d, 0xe8, 0x4a, 0x5d, 0xce, 0x67, 0x78, 0x98, 0x0b, 0x1e, 0x6e,
	0xf6, 0x35, 0x8e, 0x25, 0x25, 0x85, 0x5e, 0x2e, 0xb1, 0x91, 0xb0, 0x11, 0x46, 0xa8, 0x2d, 0x6c,
	0x06, 0xc6, 0x59, 0xab, 0xde, 0x4b, 0x8c, 0x2f, 0xa7, 0x4d, 0xe1, 0x85, 0xc5, 0xf9, 0x15, 0x40,
	0x1f, 0xa0, 0x7d, 0x66, 0x79, 0xcd, 0x58, 0x23, 0x60, 0xe5, 0x12, 0xbf, 0x66, 0x48, 0x0d, 0x8d,
	0x1a, 0x42, 0x4b, 0x25, 0x41, 0x83, 0xf1, 0x11, 0x5e, 0xaa, 0x14, 0x54, 0x0c, 0x8d, 0x04, 0x0c,
	0xa1, 0xa5, 0x2c, 0xd2, 0x10, 0x54, 0xd8, 0x7c, 0x62, 0xcd, 0x1c, 0x63, 0x0f, 0xe9, 0x72, 0xd6,
	0x18, 0x9e, 0x59, 0x9c, 0x6b, 0x63, 0x16, 0xd0, 0x3d, 0xf1, 0x48, 0x91, 0x34, 0x9c, 0x2a, 0xea,
	0xa0, 0xe6, 0x54, 0x53, 0x71, 0x47, 0x51, 0x2a, 0x38, 0x8e, 0x88, 0x50, 0x23, 0x4a, 0x1a, 0x8c,
	0x2a, 0x4a, 0x15, 0xe3, 0x1a, 0x51, 0xd2, 0x10, 0x74, 0xd4, 0xb9, 0x15, 0xc5, 0xc9, 0x9b, 0x90,
	0x6c, 0x78, 0x6a, 0x71, 0xae, 0x09, 0xa9, 0x0a, 0xb7, 0xc6, 0xb9, 0x46, 0xe1, 0x6a, 0x0a, 0x5b,
	0x40, 0x77, 0x4a, 0x6e, 0x22, 0xeb, 0xbd, 0xbd, 0x12, 0x1a, 0x5f, 0xf1, 0x0a, 0x7a, 0x0e, 0xee,
	0xc4, 0xf6, 0x64, 0x2f, 0x98, 0x90, 0x88, 0x35, 0xa8, 0x67, 0x76, 0x86, 0xe6, 0x34, 0xa3, 0xf3,
	0xf3, 0xf7, 0x6d, 0x1b, 0xee, 0x16, 0x61, 0x12, 0xd8, 0x87, 0xf3, 0x06, 0x7b, 0x72, 0x74, 0x0e,
	0xf0, 0x1e, 0x15, 0xd3, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xcd, 0x5c, 0x47, 0xad,
	0x08, 0x00, 0x00,
}
