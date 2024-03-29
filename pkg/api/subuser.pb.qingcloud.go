// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: subuser.proto

package service

import (
	"fmt"
	"reflect"

	proto "github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/client"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Invalid

	_ = proto.Marshal
	_ = client.NewClient
)

func init() {
	TypeInfoMap["DescribeSubUsersInput"] = reflect.TypeOf((*DescribeSubUsersInput)(nil))
	TypeInfoMap["DescribeSubUsersOutput"] = reflect.TypeOf((*DescribeSubUsersOutput)(nil))
	TypeInfoMap["CreateSubUserInput"] = reflect.TypeOf((*CreateSubUserInput)(nil))
	TypeInfoMap["CreateSubUserOutput"] = reflect.TypeOf((*CreateSubUserOutput)(nil))
	TypeInfoMap["ModifySubUserAttributesInput"] = reflect.TypeOf((*ModifySubUserAttributesInput)(nil))
	TypeInfoMap["ModifySubUserAttributesOutput"] = reflect.TypeOf((*ModifySubUserAttributesOutput)(nil))
	TypeInfoMap["DeleteSubUsersInput"] = reflect.TypeOf((*DeleteSubUsersInput)(nil))
	TypeInfoMap["DeleteSubUsersOutput"] = reflect.TypeOf((*DeleteSubUsersOutput)(nil))
	TypeInfoMap["RestoreSubUsersInput"] = reflect.TypeOf((*RestoreSubUsersInput)(nil))
	TypeInfoMap["RestoreSubUsersOutput"] = reflect.TypeOf((*RestoreSubUsersOutput)(nil))
}

type SubuserServiceInterface interface {
	DescribeSubUsers(in *DescribeSubUsersInput) (out *DescribeSubUsersOutput, err error)
	CreateSubUser(in *CreateSubUserInput) (out *CreateSubUserOutput, err error)
	ModifySubUserAttributes(in *ModifySubUserAttributesInput) (out *ModifySubUserAttributesOutput, err error)
	DeleteSubUsers(in *DeleteSubUsersInput) (out *DeleteSubUsersOutput, err error)
	RestoreSubUsers(in *RestoreSubUsersInput) (out *RestoreSubUsersOutput, err error)
}

type SubuserService struct {
	ServerInfo *ServerInfo
}

func NewSubuserService(server *ServerInfo) (p *SubuserService) {
	return &SubuserService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["DescribeSubUsers"] = ServiceApiSpec{
		ServiceName:    "SubuserService",
		ActionName:     "DescribeSubUsers",
		InputTypeName:  "DescribeSubUsersInput",
		OutputTypeName: "DescribeSubUsersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SubuserService)(nil)),
		InputType:   reflect.TypeOf((*DescribeSubUsersInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeSubUsersOutput)(nil)),
	}
	ServiceApiSpecMap["CreateSubUser"] = ServiceApiSpec{
		ServiceName:    "SubuserService",
		ActionName:     "CreateSubUser",
		InputTypeName:  "CreateSubUserInput",
		OutputTypeName: "CreateSubUserOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SubuserService)(nil)),
		InputType:   reflect.TypeOf((*CreateSubUserInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateSubUserOutput)(nil)),
	}
	ServiceApiSpecMap["ModifySubUserAttributes"] = ServiceApiSpec{
		ServiceName:    "SubuserService",
		ActionName:     "ModifySubUserAttributes",
		InputTypeName:  "ModifySubUserAttributesInput",
		OutputTypeName: "ModifySubUserAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SubuserService)(nil)),
		InputType:   reflect.TypeOf((*ModifySubUserAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifySubUserAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteSubUsers"] = ServiceApiSpec{
		ServiceName:    "SubuserService",
		ActionName:     "DeleteSubUsers",
		InputTypeName:  "DeleteSubUsersInput",
		OutputTypeName: "DeleteSubUsersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SubuserService)(nil)),
		InputType:   reflect.TypeOf((*DeleteSubUsersInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteSubUsersOutput)(nil)),
	}
	ServiceApiSpecMap["RestoreSubUsers"] = ServiceApiSpec{
		ServiceName:    "SubuserService",
		ActionName:     "RestoreSubUsers",
		InputTypeName:  "RestoreSubUsersInput",
		OutputTypeName: "RestoreSubUsersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SubuserService)(nil)),
		InputType:   reflect.TypeOf((*RestoreSubUsersInput)(nil)),
		OutputType:  reflect.TypeOf((*RestoreSubUsersOutput)(nil)),
	}
}

func (p *SubuserService) DescribeSubUsers(input *DescribeSubUsersInput) (output *DescribeSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSubUsersInput)
	}

	output = new(DescribeSubUsersOutput)
	err = client.CallMethod("DescribeSubUsers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) CreateSubUser(input *CreateSubUserInput) (output *CreateSubUserOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSubUserInput)
	}

	output = new(CreateSubUserOutput)
	err = client.CallMethod("CreateSubUser", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) ModifySubUserAttributes(input *ModifySubUserAttributesInput) (output *ModifySubUserAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifySubUserAttributesInput)
	}

	output = new(ModifySubUserAttributesOutput)
	err = client.CallMethod("ModifySubUserAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) DeleteSubUsers(input *DeleteSubUsersInput) (output *DeleteSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSubUsersInput)
	}

	output = new(DeleteSubUsersOutput)
	err = client.CallMethod("DeleteSubUsers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SubuserService) RestoreSubUsers(input *RestoreSubUsersInput) (output *RestoreSubUsersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RestoreSubUsersInput)
	}

	output = new(RestoreSubUsersOutput)
	err = client.CallMethod("RestoreSubUsers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
