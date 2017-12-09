// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: subuser.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewSubuserService(server *ServerInfo) (p *SubuserService) {
	return &SubuserService{
		ServerInfo: server,
	}
}

func (p *SubuserService) DescribeSubUsers(input *DescribeSubUsersInput) (output *DescribeSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeSubUsersOutput)

	err = client.CallMethod(nil, "DescribeSubUsers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) CreateSubUser(input *CreateSubUserInput) (output *CreateSubUserOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(CreateSubUserOutput)

	err = client.CallMethod(nil, "CreateSubUser", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) ModifySubUserAttributes(input *ModifySubUserAttributesInput) (output *ModifySubUserAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ModifySubUserAttributesOutput)

	err = client.CallMethod(nil, "ModifySubUserAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) DeleteSubUsers(input *DeleteSubUsersInput) (output *DeleteSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DeleteSubUsersOutput)

	err = client.CallMethod(nil, "DeleteSubUsers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) RestoreSubUsers(input *RestoreSubUsersInput) (output *RestoreSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(RestoreSubUsersOutput)

	err = client.CallMethod(nil, "RestoreSubUsers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
