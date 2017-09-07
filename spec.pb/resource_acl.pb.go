// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_acl.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import request_data_pkg "github.com/chai2010/qingcloud-go/request/data"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = request_data_pkg.Operation{}
var _ = errors.ParameterRequiredError{}

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
	proto.RegisterType((*ResourceACLServiceProperties)(nil), "spec.ResourceACLServiceProperties")
}

type ResourceACLServiceInterface interface {
	DescribeSharedResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyResourceGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyUserGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyUserGroupMemberAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyGroupRoleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifyGroupRoleRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	GrantResourceGroupsToUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RevokeResourceGroupsFromUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeResourceUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type ResourceACLService struct {
	Config     *config.Config
	Properties *ResourceACLServiceProperties
}

func NewResourceACLService(conf *config.Config, zone string) (p *ResourceACLService, err error) {
	return &ResourceACLService{
		Config:     conf,
		Properties: &ResourceACLServiceProperties{Zone: zone},
	}, nil
}

func (p *ResourceACLService) DescribeSharedResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSharedResourceGroups",
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
func (p *ResourceACLService) DescribeResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceGroups",
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
func (p *ResourceACLService) CreateResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateResourceGroups",
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
func (p *ResourceACLService) ModifyResourceGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyResourceGroupAttributes",
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
func (p *ResourceACLService) DeleteResourceGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteResourceGroups",
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
func (p *ResourceACLService) DescribeResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceGroupItems",
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
func (p *ResourceACLService) AddResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddResourceGroupItems",
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
func (p *ResourceACLService) DeleteResourceGroupItems(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteResourceGroupItems",
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
func (p *ResourceACLService) DescribeUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeUserGroups",
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
func (p *ResourceACLService) CreateUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateUserGroups",
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
func (p *ResourceACLService) ModifyUserGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyUserGroupAttributes",
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
func (p *ResourceACLService) DeleteUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteUserGroups",
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
func (p *ResourceACLService) DescribeUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeUserGroupMembers",
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
func (p *ResourceACLService) AddUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddUserGroupMembers",
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
func (p *ResourceACLService) ModifyUserGroupMemberAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyUserGroupMemberAttributes",
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
func (p *ResourceACLService) DeleteUserGroupMembers(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteUserGroupMembers",
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
func (p *ResourceACLService) DescribeGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeGroupRoles",
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
func (p *ResourceACLService) CreateGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateGroupRoles",
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
func (p *ResourceACLService) ModifyGroupRoleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyGroupRoleAttributes",
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
func (p *ResourceACLService) DeleteGroupRoles(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteGroupRoles",
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
func (p *ResourceACLService) DescribeGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeGroupRoleRules",
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
func (p *ResourceACLService) AddGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddGroupRoleRules",
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
func (p *ResourceACLService) ModifyGroupRoleRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyGroupRoleRuleAttributes",
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
func (p *ResourceACLService) DeleteGroupRoleRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteGroupRoleRules",
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
func (p *ResourceACLService) GrantResourceGroupsToUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GrantResourceGroupsToUserGroups",
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
func (p *ResourceACLService) RevokeResourceGroupsFromUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RevokeResourceGroupsFromUserGroups",
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
func (p *ResourceACLService) DescribeResourceUserGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeResourceUserGroups",
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

func init() { proto.RegisterFile("resource_acl.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x85, 0x84, 0x2a, 0x75, 0x2e, 0xa5, 0xdb, 0x16, 0x51, 0x4a, 0x45, 0xc5, 0xa9, 0x27,
	0x23, 0xd1, 0x4f, 0x60, 0xfe, 0x36, 0x09, 0x24, 0xc1, 0x90, 0x73, 0x64, 0x7b, 0x07, 0x62, 0xc5,
	0x66, 0xad, 0xf1, 0x1a, 0x89, 0x7c, 0xf1, 0x5c, 0xa3, 0xd8, 0xb1, 0x85, 0x09, 0x39, 0xe0, 0xe1,
	0x66, 0xaf, 0x76, 0x7f, 0x7a, 0xfb, 0xf4, 0xf6, 0x0d, 0x08, 0xc2, 0x48, 0xc5, 0xe4, 0xe2, 0xbd,
	0xed, 0xfa, 0x46, 0x48, 0x4a, 0x2b, 0x51, 0x8d, 0x42, 0x74, 0x9b, 0xbf, 0xd6, 0x4a, 0xad, 0x7d,
	0xec, 0x26, 0x6b, 0x4e, 0xbc, 0xea, 0x62, 0x10, 0xea, 0x5d, 0xba, 0xa5, 0xd3, 0x83, 0x96, 0xf5,
	0x76, 0xd0, 0x1c, 0x4c, 0x17, 0x48, 0x5b, 0xcf, 0xc5, 0x5b, 0x52, 0x21, 0x92, 0xf6, 0x30, 0x12,
	0x02, 0xaa, 0x4f, 0x6a, 0x83, 0x8d, 0xca, 0x9f, 0xca, 0xdf, 0xcf, 0x56, 0xf2, 0xdd, 0x7b, 0xfe,
	0x02, 0xe2, 0xfd, 0x21, 0x71, 0x0d, 0xad, 0x21, 0x46, 0x2e, 0x79, 0x0e, 0x2e, 0x1e, 0x6c, 0x42,
	0x99, 0xed, 0x99, 0x90, 0x8a, 0xc3, 0x48, 0xd4, 0x8d, 0x54, 0x88, 0x91, 0x09, 0x31, 0x46, 0xaf,
	0x42, 0x9a, 0x1f, 0xac, 0x8b, 0xff, 0x50, 0xcf, 0x78, 0x4c, 0xd2, 0x18, 0xbe, 0x0f, 0x08, 0x6d,
	0xcd, 0xe5, 0xdc, 0xc0, 0xef, 0x99, 0x92, 0xde, 0x6a, 0x57, 0xe0, 0x98, 0x5a, 0x93, 0xe7, 0xc4,
	0x1a, 0x4b, 0x09, 0x1b, 0xa2, 0x8f, 0x6c, 0x61, 0x53, 0x68, 0x1e, 0xb5, 0xea, 0x42, 0x63, 0x70,
	0x3a, 0x6d, 0x02, 0x3f, 0x4c, 0x29, 0xcf, 0x00, 0xba, 0x84, 0xc6, 0x91, 0xeb, 0x95, 0x63, 0x0d,
	0x41, 0x64, 0x57, 0xbc, 0x8b, 0x90, 0x4a, 0x1a, 0xd5, 0x87, 0x5a, 0x9a, 0x04, 0x06, 0xe3, 0x0a,
	0x7e, 0xa6, 0x29, 0xc8, 0x19, 0x8c, 0x04, 0xf4, 0xa1, 0x96, 0x5a, 0xc4, 0x10, 0x94, 0xd8, 0x7c,
	0x60, 0xcd, 0x0c, 0x03, 0x07, 0xe9, 0x74, 0xd6, 0x08, 0xbe, 0x99, 0x52, 0xb2, 0x31, 0x73, 0x68,
	0x1f, 0x78, 0x94, 0x92, 0x18, 0x4e, 0x25, 0x75, 0x50, 0x70, 0xaa, 0xac, 0xb8, 0xbd, 0x28, 0x25,
	0x1c, 0x4b, 0xf9, 0xc8, 0x88, 0x12, 0x83, 0x91, 0x47, 0x29, 0x67, 0x9c, 0x23, 0x4a, 0x0c, 0x41,
	0x7b, 0x9d, 0x9b, 0x53, 0xac, 0xb8, 0x0c, 0x69, 0x00, 0x5f, 0x4d, 0x29, 0x99, 0x90, 0xbc, 0x70,
	0x0b, 0x9c, 0x73, 0x14, 0x2e, 0x53, 0xd8, 0x1c, 0xda, 0x13, 0xb2, 0x37, 0xba, 0xd8, 0xdb, 0x4b,
	0xc5, 0x78, 0xc5, 0x4b, 0xe8, 0x58, 0xb8, 0x55, 0x8f, 0x07, 0xb3, 0x60, 0x4c, 0x2a, 0x60, 0x50,
	0x8f, 0x4c, 0x86, 0xf2, 0x34, 0xe7, 0x53, 0xf2, 0xff, 0xef, 0x25, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x3b, 0x8e, 0x83, 0x6d, 0x08, 0x00, 0x00,
}
