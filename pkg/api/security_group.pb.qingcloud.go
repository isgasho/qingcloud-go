// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: security_group.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type SecurityGroupServiceInterface interface {
	DescribeSecurityGroups(in *DescribeSecurityGroupsInput) (out *DescribeSecurityGroupsOutput, err error)
	CreateSecurityGroup(in *CreateSecurityGroupInput) (out *CreateSecurityGroupOutput, err error)
	DeleteSecurityGroups(in *DeleteSecurityGroupsInput) (out *DeleteSecurityGroupsOutput, err error)
	ApplySecurityGroup(in *ApplySecurityGroupInput) (out *ApplySecurityGroupOutpu, err error)
	ModifySecurityGroupAttributes(in *ModifySecurityGroupAttributesInput) (out *ModifySecurityGroupAttributesOutput, err error)
	DescribeSecurityGroupRules(in *DescribeSecurityGroupRulesInput) (out *DescribeSecurityGroupRulesOutput, err error)
	AddSecurityGroupRules(in *AddSecurityGroupRulesInput) (out *AddSecurityGroupRulesOutput, err error)
	DeleteSecurityGroupRules(in *DeleteSecurityGroupRulesInput) (out *DeleteSecurityGroupRulesOutput, err error)
	ModifySecurityGroupRuleAttributes(in *ModifySecurityGroupRuleAttributesInput) (out *ModifySecurityGroupRuleAttributesOutput, err error)
	CreateSecurityGroupSnapshot(in *CreateSecurityGroupSnapshotInput) (out *CreateSecurityGroupSnapshotOutput, err error)
	DescribeSecurityGroupSnapshots(in *DescribeSecurityGroupSnapshotsInput) (out *DescribeSecurityGroupSnapshotsOutput, err error)
	DeleteSecurityGroupSnapshots(in *DeleteSecurityGroupSnapshotsInput) (out *DeleteSecurityGroupSnapshotsOutput, err error)
	RollbackSecurityGroup(in *RollbackSecurityGroupInput) (out *RollbackSecurityGroupOutput, err error)
	DescribeSecurityGroupIPSets(in *DescribeSecurityGroupIPSetsInput) (out *DescribeSecurityGroupIPSetsOutput, err error)
	CreateSecurityGroupIPSet(in *CreateSecurityGroupIPSetInput) (out *CreateSecurityGroupIPSetOutput, err error)
	DeleteSecurityGroupIPSets(in *DeleteSecurityGroupIPSetsInput) (out *DeleteSecurityGroupIPSetsOutput, err error)
	ModifySecurityGroupIPSetAttributes(in *ModifySecurityGroupIPSetAttributesInput) (out *ModifySecurityGroupIPSetAttributesOutput, err error)
	CopySecurityGroupIPSets(in *CopySecurityGroupIPSetsInput) (out *CopySecurityGroupIPSetsOutput, err error)
}

type SecurityGroupService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewSecurityGroupService(server *ServerInfo) (p *SecurityGroupService) {
	return &SecurityGroupService{
		ServerInfo: server,
	}
}

func (p *SecurityGroupService) DescribeSecurityGroups(input *DescribeSecurityGroupsInput) (output *DescribeSecurityGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSecurityGroupsInput)
	}

	output = new(DescribeSecurityGroupsOutput)
	err = client.CallMethod("DescribeSecurityGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) CreateSecurityGroup(input *CreateSecurityGroupInput) (output *CreateSecurityGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSecurityGroupInput)
	}

	output = new(CreateSecurityGroupOutput)
	err = client.CallMethod("CreateSecurityGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DeleteSecurityGroups(input *DeleteSecurityGroupsInput) (output *DeleteSecurityGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSecurityGroupsInput)
	}

	output = new(DeleteSecurityGroupsOutput)
	err = client.CallMethod("DeleteSecurityGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) ApplySecurityGroup(input *ApplySecurityGroupInput) (output *ApplySecurityGroupOutpu, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ApplySecurityGroupInput)
	}

	output = new(ApplySecurityGroupOutpu)
	err = client.CallMethod("ApplySecurityGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) ModifySecurityGroupAttributes(input *ModifySecurityGroupAttributesInput) (output *ModifySecurityGroupAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifySecurityGroupAttributesInput)
	}

	output = new(ModifySecurityGroupAttributesOutput)
	err = client.CallMethod("ModifySecurityGroupAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DescribeSecurityGroupRules(input *DescribeSecurityGroupRulesInput) (output *DescribeSecurityGroupRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSecurityGroupRulesInput)
	}

	output = new(DescribeSecurityGroupRulesOutput)
	err = client.CallMethod("DescribeSecurityGroupRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) AddSecurityGroupRules(input *AddSecurityGroupRulesInput) (output *AddSecurityGroupRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddSecurityGroupRulesInput)
	}

	output = new(AddSecurityGroupRulesOutput)
	err = client.CallMethod("AddSecurityGroupRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DeleteSecurityGroupRules(input *DeleteSecurityGroupRulesInput) (output *DeleteSecurityGroupRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSecurityGroupRulesInput)
	}

	output = new(DeleteSecurityGroupRulesOutput)
	err = client.CallMethod("DeleteSecurityGroupRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) ModifySecurityGroupRuleAttributes(input *ModifySecurityGroupRuleAttributesInput) (output *ModifySecurityGroupRuleAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifySecurityGroupRuleAttributesInput)
	}

	output = new(ModifySecurityGroupRuleAttributesOutput)
	err = client.CallMethod("ModifySecurityGroupRuleAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) CreateSecurityGroupSnapshot(input *CreateSecurityGroupSnapshotInput) (output *CreateSecurityGroupSnapshotOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSecurityGroupSnapshotInput)
	}

	output = new(CreateSecurityGroupSnapshotOutput)
	err = client.CallMethod("CreateSecurityGroupSnapshot", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DescribeSecurityGroupSnapshots(input *DescribeSecurityGroupSnapshotsInput) (output *DescribeSecurityGroupSnapshotsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSecurityGroupSnapshotsInput)
	}

	output = new(DescribeSecurityGroupSnapshotsOutput)
	err = client.CallMethod("DescribeSecurityGroupSnapshots", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DeleteSecurityGroupSnapshots(input *DeleteSecurityGroupSnapshotsInput) (output *DeleteSecurityGroupSnapshotsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSecurityGroupSnapshotsInput)
	}

	output = new(DeleteSecurityGroupSnapshotsOutput)
	err = client.CallMethod("DeleteSecurityGroupSnapshots", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) RollbackSecurityGroup(input *RollbackSecurityGroupInput) (output *RollbackSecurityGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RollbackSecurityGroupInput)
	}

	output = new(RollbackSecurityGroupOutput)
	err = client.CallMethod("RollbackSecurityGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DescribeSecurityGroupIPSets(input *DescribeSecurityGroupIPSetsInput) (output *DescribeSecurityGroupIPSetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSecurityGroupIPSetsInput)
	}

	output = new(DescribeSecurityGroupIPSetsOutput)
	err = client.CallMethod("DescribeSecurityGroupIPSets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) CreateSecurityGroupIPSet(input *CreateSecurityGroupIPSetInput) (output *CreateSecurityGroupIPSetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSecurityGroupIPSetInput)
	}

	output = new(CreateSecurityGroupIPSetOutput)
	err = client.CallMethod("CreateSecurityGroupIPSet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) DeleteSecurityGroupIPSets(input *DeleteSecurityGroupIPSetsInput) (output *DeleteSecurityGroupIPSetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSecurityGroupIPSetsInput)
	}

	output = new(DeleteSecurityGroupIPSetsOutput)
	err = client.CallMethod("DeleteSecurityGroupIPSets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) ModifySecurityGroupIPSetAttributes(input *ModifySecurityGroupIPSetAttributesInput) (output *ModifySecurityGroupIPSetAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifySecurityGroupIPSetAttributesInput)
	}

	output = new(ModifySecurityGroupIPSetAttributesOutput)
	err = client.CallMethod("ModifySecurityGroupIPSetAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SecurityGroupService) CopySecurityGroupIPSets(input *CopySecurityGroupIPSetsInput) (output *CopySecurityGroupIPSetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CopySecurityGroupIPSetsInput)
	}

	output = new(CopySecurityGroupIPSetsOutput)
	err = client.CallMethod("CopySecurityGroupIPSets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
