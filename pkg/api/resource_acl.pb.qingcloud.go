// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: resource_acl.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type ResourceACLServiceInterface interface {
	DescribeSharedResourceGroups(in *DescribeSharedResourceGroupsInput) (out *DescribeSharedResourceGroupsOutput, err error)
	DescribeResourceGroups(in *DescribeResourceGroupsInput) (out *DescribeResourceGroupsOutput, err error)
	CreateResourceGroups(in *CreateResourceGroupsInput) (out *CreateResourceGroupsOutput, err error)
	ModifyResourceGroupAttributes(in *ModifyResourceGroupAttributesInput) (out *ModifyResourceGroupAttributesOutput, err error)
	DeleteResourceGroups(in *DeleteResourceGroupsInput) (out *DeleteResourceGroupsOutput, err error)
	DescribeResourceGroupItems(in *DescribeResourceGroupItemsInput) (out *DescribeResourceGroupItemsOutput, err error)
	AddResourceGroupItems(in *AddResourceGroupItemsInput) (out *AddResourceGroupItemsOutput, err error)
	DeleteResourceGroupItems(in *DeleteResourceGroupItemsInput) (out *DeleteResourceGroupItemsOutput, err error)
	DescribeUserGroups(in *DescribeUserGroupsInput) (out *DescribeUserGroupsOutput, err error)
	CreateUserGroups(in *CreateUserGroupsInput) (out *CreateUserGroupsOutput, err error)
	ModifyUserGroupAttributes(in *ModifyUserGroupAttributesInput) (out *ModifyUserGroupAttributesOutput, err error)
	DeleteUserGroups(in *DeleteUserGroupsInput) (out *DeleteUserGroupsOutput, err error)
	DescribeUserGroupMembers(in *DescribeUserGroupMembersInput) (out *DescribeUserGroupMembersOutput, err error)
	AddUserGroupMembers(in *AddUserGroupMembersInput) (out *AddUserGroupMembersOutput, err error)
	ModifyUserGroupMemberAttributes(in *ModifyUserGroupMemberAttributesInput) (out *ModifyUserGroupMemberAttributesOutput, err error)
	DeleteUserGroupMembers(in *DeleteUserGroupMembersInput) (out *DeleteUserGroupMembersOutput, err error)
	DescribeGroupRoles(in *DescribeGroupRolesInput) (out *DescribeGroupRolesOutput, err error)
	CreateGroupRoles(in *CreateGroupRolesInput) (out *CreateGroupRolesOutput, err error)
	ModifyGroupRoleAttributes(in *ModifyGroupRoleAttributesInput) (out *ModifyGroupRoleAttributesOutput, err error)
	DeleteGroupRoles(in *DeleteGroupRolesInput) (out *DeleteGroupRolesOutput, err error)
	DescribeGroupRoleRules(in *DescribeGroupRoleRulesInput) (out *DescribeGroupRoleRulesOutput, err error)
	AddGroupRoleRules(in *AddGroupRoleRulesInput) (out *AddGroupRoleRulesOutput, err error)
	ModifyGroupRoleRuleAttributes(in *ModifyGroupRoleRuleAttributesInput) (out *ModifyGroupRoleRuleAttributesOutput, err error)
	DeleteGroupRoleRules(in *DeleteGroupRoleRulesInput) (out *DeleteGroupRoleRulesOutput, err error)
	GrantResourceGroupsToUserGroups(in *GrantResourceGroupsToUserGroupsInput) (out *GrantResourceGroupsToUserGroupsOutput, err error)
	RevokeResourceGroupsFromUserGroups(in *RevokeResourceGroupsFromUserGroupsInput) (out *RevokeResourceGroupsFromUserGroupsOutput, err error)
	DescribeResourceUserGroups(in *DescribeResourceUserGroupsInput) (out *DescribeResourceUserGroupsOutput, err error)
}

type ResourceACLService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewResourceACLService(server *ServerInfo) (p *ResourceACLService) {
	return &ResourceACLService{
		ServerInfo: server,
	}
}

func (p *ResourceACLService) DescribeSharedResourceGroups(input *DescribeSharedResourceGroupsInput) (output *DescribeSharedResourceGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSharedResourceGroupsInput)
	}

	output = new(DescribeSharedResourceGroupsOutput)
	err = client.CallMethod("DescribeSharedResourceGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeResourceGroups(input *DescribeResourceGroupsInput) (output *DescribeResourceGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeResourceGroupsInput)
	}

	output = new(DescribeResourceGroupsOutput)
	err = client.CallMethod("DescribeResourceGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) CreateResourceGroups(input *CreateResourceGroupsInput) (output *CreateResourceGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateResourceGroupsInput)
	}

	output = new(CreateResourceGroupsOutput)
	err = client.CallMethod("CreateResourceGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) ModifyResourceGroupAttributes(input *ModifyResourceGroupAttributesInput) (output *ModifyResourceGroupAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyResourceGroupAttributesInput)
	}

	output = new(ModifyResourceGroupAttributesOutput)
	err = client.CallMethod("ModifyResourceGroupAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteResourceGroups(input *DeleteResourceGroupsInput) (output *DeleteResourceGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteResourceGroupsInput)
	}

	output = new(DeleteResourceGroupsOutput)
	err = client.CallMethod("DeleteResourceGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeResourceGroupItems(input *DescribeResourceGroupItemsInput) (output *DescribeResourceGroupItemsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeResourceGroupItemsInput)
	}

	output = new(DescribeResourceGroupItemsOutput)
	err = client.CallMethod("DescribeResourceGroupItems", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) AddResourceGroupItems(input *AddResourceGroupItemsInput) (output *AddResourceGroupItemsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddResourceGroupItemsInput)
	}

	output = new(AddResourceGroupItemsOutput)
	err = client.CallMethod("AddResourceGroupItems", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteResourceGroupItems(input *DeleteResourceGroupItemsInput) (output *DeleteResourceGroupItemsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteResourceGroupItemsInput)
	}

	output = new(DeleteResourceGroupItemsOutput)
	err = client.CallMethod("DeleteResourceGroupItems", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeUserGroups(input *DescribeUserGroupsInput) (output *DescribeUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeUserGroupsInput)
	}

	output = new(DescribeUserGroupsOutput)
	err = client.CallMethod("DescribeUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) CreateUserGroups(input *CreateUserGroupsInput) (output *CreateUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateUserGroupsInput)
	}

	output = new(CreateUserGroupsOutput)
	err = client.CallMethod("CreateUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) ModifyUserGroupAttributes(input *ModifyUserGroupAttributesInput) (output *ModifyUserGroupAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyUserGroupAttributesInput)
	}

	output = new(ModifyUserGroupAttributesOutput)
	err = client.CallMethod("ModifyUserGroupAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteUserGroups(input *DeleteUserGroupsInput) (output *DeleteUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteUserGroupsInput)
	}

	output = new(DeleteUserGroupsOutput)
	err = client.CallMethod("DeleteUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeUserGroupMembers(input *DescribeUserGroupMembersInput) (output *DescribeUserGroupMembersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeUserGroupMembersInput)
	}

	output = new(DescribeUserGroupMembersOutput)
	err = client.CallMethod("DescribeUserGroupMembers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) AddUserGroupMembers(input *AddUserGroupMembersInput) (output *AddUserGroupMembersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddUserGroupMembersInput)
	}

	output = new(AddUserGroupMembersOutput)
	err = client.CallMethod("AddUserGroupMembers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) ModifyUserGroupMemberAttributes(input *ModifyUserGroupMemberAttributesInput) (output *ModifyUserGroupMemberAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyUserGroupMemberAttributesInput)
	}

	output = new(ModifyUserGroupMemberAttributesOutput)
	err = client.CallMethod("ModifyUserGroupMemberAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteUserGroupMembers(input *DeleteUserGroupMembersInput) (output *DeleteUserGroupMembersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteUserGroupMembersInput)
	}

	output = new(DeleteUserGroupMembersOutput)
	err = client.CallMethod("DeleteUserGroupMembers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeGroupRoles(input *DescribeGroupRolesInput) (output *DescribeGroupRolesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeGroupRolesInput)
	}

	output = new(DescribeGroupRolesOutput)
	err = client.CallMethod("DescribeGroupRoles", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) CreateGroupRoles(input *CreateGroupRolesInput) (output *CreateGroupRolesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateGroupRolesInput)
	}

	output = new(CreateGroupRolesOutput)
	err = client.CallMethod("CreateGroupRoles", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) ModifyGroupRoleAttributes(input *ModifyGroupRoleAttributesInput) (output *ModifyGroupRoleAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyGroupRoleAttributesInput)
	}

	output = new(ModifyGroupRoleAttributesOutput)
	err = client.CallMethod("ModifyGroupRoleAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteGroupRoles(input *DeleteGroupRolesInput) (output *DeleteGroupRolesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteGroupRolesInput)
	}

	output = new(DeleteGroupRolesOutput)
	err = client.CallMethod("DeleteGroupRoles", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeGroupRoleRules(input *DescribeGroupRoleRulesInput) (output *DescribeGroupRoleRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeGroupRoleRulesInput)
	}

	output = new(DescribeGroupRoleRulesOutput)
	err = client.CallMethod("DescribeGroupRoleRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) AddGroupRoleRules(input *AddGroupRoleRulesInput) (output *AddGroupRoleRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddGroupRoleRulesInput)
	}

	output = new(AddGroupRoleRulesOutput)
	err = client.CallMethod("AddGroupRoleRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) ModifyGroupRoleRuleAttributes(input *ModifyGroupRoleRuleAttributesInput) (output *ModifyGroupRoleRuleAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyGroupRoleRuleAttributesInput)
	}

	output = new(ModifyGroupRoleRuleAttributesOutput)
	err = client.CallMethod("ModifyGroupRoleRuleAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DeleteGroupRoleRules(input *DeleteGroupRoleRulesInput) (output *DeleteGroupRoleRulesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteGroupRoleRulesInput)
	}

	output = new(DeleteGroupRoleRulesOutput)
	err = client.CallMethod("DeleteGroupRoleRules", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) GrantResourceGroupsToUserGroups(input *GrantResourceGroupsToUserGroupsInput) (output *GrantResourceGroupsToUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GrantResourceGroupsToUserGroupsInput)
	}

	output = new(GrantResourceGroupsToUserGroupsOutput)
	err = client.CallMethod("GrantResourceGroupsToUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) RevokeResourceGroupsFromUserGroups(input *RevokeResourceGroupsFromUserGroupsInput) (output *RevokeResourceGroupsFromUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RevokeResourceGroupsFromUserGroupsInput)
	}

	output = new(RevokeResourceGroupsFromUserGroupsOutput)
	err = client.CallMethod("RevokeResourceGroupsFromUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ResourceACLService) DescribeResourceUserGroups(input *DescribeResourceUserGroupsInput) (output *DescribeResourceUserGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeResourceUserGroupsInput)
	}

	output = new(DescribeResourceUserGroupsOutput)
	err = client.CallMethod("DescribeResourceUserGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
