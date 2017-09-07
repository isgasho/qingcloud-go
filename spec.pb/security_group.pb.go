// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security_group.proto

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

type SecurityGroupServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *SecurityGroupServiceProperties) Reset()                    { *m = SecurityGroupServiceProperties{} }
func (m *SecurityGroupServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*SecurityGroupServiceProperties) ProtoMessage()               {}
func (*SecurityGroupServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *SecurityGroupServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*SecurityGroupServiceProperties)(nil), "spec.SecurityGroupServiceProperties")
}

type SecurityGroupServiceInterface interface {
	DescribeSecurityGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateSecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSecurityGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ApplySecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifySecurityGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	AddSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifySecurityGroupRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateSecurityGroupSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeSecurityGroupSnapshots(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSecurityGroupSnapshots(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	RollbackSecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DescribeSecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CreateSecurityGroupIPSet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	DeleteSecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	ModifySecurityGroupIPSetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
	CopySecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error)
}

type SecurityGroupService struct {
	Config     *config.Config
	Properties *SecurityGroupServiceProperties
}

func NewSecurityGroupService(conf *config.Config, zone string) (p *SecurityGroupService, err error) {
	return &SecurityGroupService{
		Config:     conf,
		Properties: &SecurityGroupServiceProperties{Zone: zone},
	}, nil
}

func (p *SecurityGroupService) DescribeSecurityGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSecurityGroups",
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
func (p *SecurityGroupService) CreateSecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSecurityGroup",
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
func (p *SecurityGroupService) DeleteSecurityGroups(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSecurityGroups",
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
func (p *SecurityGroupService) ApplySecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplySecurityGroup",
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
func (p *SecurityGroupService) ModifySecurityGroupAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySecurityGroupAttributes",
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
func (p *SecurityGroupService) DescribeSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSecurityGroupRules",
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
func (p *SecurityGroupService) AddSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddSecurityGroupRules",
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
func (p *SecurityGroupService) DeleteSecurityGroupRules(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSecurityGroupRules",
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
func (p *SecurityGroupService) ModifySecurityGroupRuleAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySecurityGroupRuleAttributes",
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
func (p *SecurityGroupService) CreateSecurityGroupSnapshot(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSecurityGroupSnapshot",
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
func (p *SecurityGroupService) DescribeSecurityGroupSnapshots(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSecurityGroupSnapshots",
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
func (p *SecurityGroupService) DeleteSecurityGroupSnapshots(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSecurityGroupSnapshots",
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
func (p *SecurityGroupService) RollbackSecurityGroup(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RollbackSecurityGroup",
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
func (p *SecurityGroupService) DescribeSecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeSecurityGroupIPSets",
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
func (p *SecurityGroupService) CreateSecurityGroupIPSet(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateSecurityGroupIPSet",
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
func (p *SecurityGroupService) DeleteSecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteSecurityGroupIPSets",
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
func (p *SecurityGroupService) ModifySecurityGroupIPSetAttributes(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifySecurityGroupIPSetAttributes",
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
func (p *SecurityGroupService) CopySecurityGroupIPSets(in *google_protobuf.Empty) (out *google_protobuf.Empty, err error) {
	if in == nil {
		in = &google_protobuf.Empty{}
	}
	o := &request_data_pkg.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopySecurityGroupIPSets",
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

func init() { proto.RegisterFile("security_group.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0xb3, 0x64, 0x31, 0xf1, 0x79, 0xab, 0xdb, 0x9c, 0x9b, 0x2e, 0xba, 0x93, 0x27, 0x96,
	0xa8, 0xff, 0xc0, 0x32, 0x26, 0x4e, 0x45, 0x09, 0x78, 0x37, 0xfc, 0x78, 0xc3, 0xc6, 0xba, 0x36,
	0x6d, 0x31, 0xc1, 0x9b, 0xff, 0xb9, 0x81, 0xb9, 0x03, 0x49, 0x3d, 0x00, 0xbb, 0x41, 0xe1, 0x7d,
	0xf2, 0xe9, 0xf7, 0x7d, 0xa1, 0xa7, 0x30, 0xce, 0x24, 0xd5, 0xf9, 0x5b, 0x2a, 0x79, 0x26, 0x2c,
	0x21, 0xb9, 0xe6, 0xa4, 0xab, 0x04, 0xc6, 0xa3, 0x71, 0xca, 0x79, 0xca, 0x70, 0x56, 0x9e, 0x45,
	0xd9, 0x7a, 0x86, 0x9f, 0x42, 0xe7, 0xdb, 0x5f, 0xa6, 0xb7, 0x30, 0x09, 0xfe, 0x46, 0x9d, 0x62,
	0x32, 0x40, 0xf9, 0x45, 0x63, 0xf4, 0x24, 0x17, 0x28, 0x35, 0x45, 0x45, 0x08, 0x74, 0xbf, 0xf9,
	0x06, 0x87, 0x9d, 0x8b, 0xce, 0xd5, 0xa1, 0x5f, 0x3e, 0x5f, 0xff, 0x1c, 0x41, 0xcf, 0x34, 0x46,
	0xee, 0x61, 0x60, 0xa3, 0x8a, 0x25, 0x8d, 0xb0, 0xf2, 0x5d, 0x91, 0x81, 0xb5, 0xd5, 0xb0, 0x76,
	0x1a, 0xd6, 0xb2, 0xd0, 0x18, 0xfd, 0x73, 0x4e, 0x96, 0x70, 0xbc, 0x90, 0x18, 0xea, 0x2a, 0xa7,
	0x36, 0xe6, 0x0e, 0x7a, 0x36, 0x32, 0xd4, 0x6d, 0x75, 0x6c, 0x20, 0x73, 0x21, 0x58, 0xde, 0xce,
	0xe6, 0x05, 0xce, 0x5d, 0x9e, 0xd0, 0x75, 0x15, 0x33, 0xd7, 0x5a, 0xd2, 0x28, 0xd3, 0x58, 0x5f,
	0xeb, 0x09, 0x46, 0xc6, 0xbc, 0xfd, 0x8c, 0x35, 0xa0, 0x39, 0xd0, 0x9f, 0x27, 0xc9, 0x1e, 0x40,
	0x0f, 0x30, 0x34, 0xa4, 0xde, 0x8c, 0x15, 0xc0, 0xa5, 0x21, 0xb3, 0x82, 0xd5, 0x22, 0x37, 0x17,
	0xc6, 0x86, 0x76, 0x05, 0x9b, 0x50, 0xa8, 0x77, 0xae, 0x6b, 0xe3, 0x3c, 0x98, 0x18, 0xd7, 0xb0,
	0x03, 0xd6, 0x17, 0x7c, 0x86, 0x33, 0x43, 0x82, 0xcd, 0x79, 0x0e, 0xf4, 0x7d, 0xce, 0x58, 0x14,
	0xc6, 0x1f, 0xed, 0x2a, 0xec, 0xc2, 0xd8, 0x78, 0xd5, 0x95, 0x17, 0xa0, 0x6e, 0xd4, 0x14, 0xc3,
	0x22, 0x4a, 0x58, 0x6d, 0xd6, 0x23, 0x9c, 0x1a, 0x32, 0x6b, 0x28, 0xf6, 0x0a, 0x53, 0x43, 0xed,
	0x4a, 0x58, 0x8b, 0xde, 0xad, 0xe0, 0x64, 0xc1, 0x45, 0xbe, 0x07, 0xc1, 0xe8, 0xa0, 0x7c, 0xbf,
	0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x69, 0x68, 0xc4, 0x07, 0xfb, 0x05, 0x00, 0x00,
}
