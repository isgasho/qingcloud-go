// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: router.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type RouterServiceInterface interface {
	DescribeRouters(in *DescribeRoutersInput) (out *DescribeRoutersOutput, err error)
	CreateRouters(in *CreateRoutersInput) (out *CreateRoutersOutput, err error)
	DeleteRouters(in *DeleteRoutersInput) (out *DeleteRoutersOutput, err error)
	UpdateRouters(in *UpdateRoutersInput) (out *UpdateRoutersOutput, err error)
	PowerOffRouters(in *PowerOffRoutersInput) (out *PowerOffRoutersOutput, err error)
	PowerOnRouters(in *PowerOnRoutersInput) (out *PowerOnRoutersOutput, err error)
	JoinRouter(in *JoinRouterInput) (out *JoinRouterOutput, err error)
	LeaveRouter(in *LeaveRouterInput) (out *LeaveRouterOutput, err error)
	ModifyRouterAttributes(in *ModifyRouterAttributesInput) (out *ModifyRouterAttributesOutput, err error)
	DescribeRouterStatics(in *DescribeRouterStaticsInput) (out *DescribeRouterStaticsOutput, err error)
	AddRouterStatics(in *AddRouterStaticsInput) (out *AddRouterStaticsOutput, err error)
	ModifyRouterStaticAttributes(in *ModifyRouterStaticAttributesInput) (out *ModifyRouterStaticAttributesOutput, err error)
	DeleteRouterStatics(in *DeleteRouterStaticsInput) (out *DeleteRouterStaticsOutput, err error)
	CopyRouterStatics(in *CopyRouterStaticsInput) (out *CopyRouterStaticsOutput, err error)
	DescribeRouterVxnets(in *DescribeRouterVxnetsInput) (out *DescribeRouterVxnetsOutput, err error)
	AddRouterStaticEntries(in *AddRouterStaticEntriesInput) (out *AddRouterStaticEntriesOutput, err error)
	DeleteRouterStaticEntries(in *DeleteRouterStaticEntriesInput) (out *DeleteRouterStaticEntriesOutput, err error)
	ModifyRouterStaticEntryAttributes(in *ModifyRouterStaticEntryAttributesInput) (out *ModifyRouterStaticEntryAttributesOutput, err error)
	DescribeRouterStaticEntries(in *DescribeRouterStaticEntriesInput) (out *DescribeRouterStaticEntriesOutput, err error)
}

type RouterService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewRouterService(server *ServerInfo) (p *RouterService) {
	return &RouterService{
		ServerInfo: server,
	}
}

func (p *RouterService) DescribeRouters(input *DescribeRoutersInput) (output *DescribeRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DescribeRoutersOutput)

	err = client.CallMethod(nil, "DescribeRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) CreateRouters(input *CreateRoutersInput) (output *CreateRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(CreateRoutersOutput)

	err = client.CallMethod(nil, "CreateRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DeleteRouters(input *DeleteRoutersInput) (output *DeleteRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DeleteRoutersOutput)

	err = client.CallMethod(nil, "DeleteRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) UpdateRouters(input *UpdateRoutersInput) (output *UpdateRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(UpdateRoutersOutput)

	err = client.CallMethod(nil, "UpdateRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) PowerOffRouters(input *PowerOffRoutersInput) (output *PowerOffRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(PowerOffRoutersOutput)

	err = client.CallMethod(nil, "PowerOffRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) PowerOnRouters(input *PowerOnRoutersInput) (output *PowerOnRoutersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(PowerOnRoutersOutput)

	err = client.CallMethod(nil, "PowerOnRouters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) JoinRouter(input *JoinRouterInput) (output *JoinRouterOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(JoinRouterOutput)

	err = client.CallMethod(nil, "JoinRouter", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) LeaveRouter(input *LeaveRouterInput) (output *LeaveRouterOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(LeaveRouterOutput)

	err = client.CallMethod(nil, "LeaveRouter", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) ModifyRouterAttributes(input *ModifyRouterAttributesInput) (output *ModifyRouterAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(ModifyRouterAttributesOutput)

	err = client.CallMethod(nil, "ModifyRouterAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DescribeRouterStatics(input *DescribeRouterStaticsInput) (output *DescribeRouterStaticsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DescribeRouterStaticsOutput)

	err = client.CallMethod(nil, "DescribeRouterStatics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) AddRouterStatics(input *AddRouterStaticsInput) (output *AddRouterStaticsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(AddRouterStaticsOutput)

	err = client.CallMethod(nil, "AddRouterStatics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) ModifyRouterStaticAttributes(input *ModifyRouterStaticAttributesInput) (output *ModifyRouterStaticAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(ModifyRouterStaticAttributesOutput)

	err = client.CallMethod(nil, "ModifyRouterStaticAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DeleteRouterStatics(input *DeleteRouterStaticsInput) (output *DeleteRouterStaticsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DeleteRouterStaticsOutput)

	err = client.CallMethod(nil, "DeleteRouterStatics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) CopyRouterStatics(input *CopyRouterStaticsInput) (output *CopyRouterStaticsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(CopyRouterStaticsOutput)

	err = client.CallMethod(nil, "CopyRouterStatics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DescribeRouterVxnets(input *DescribeRouterVxnetsInput) (output *DescribeRouterVxnetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DescribeRouterVxnetsOutput)

	err = client.CallMethod(nil, "DescribeRouterVxnets", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) AddRouterStaticEntries(input *AddRouterStaticEntriesInput) (output *AddRouterStaticEntriesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(AddRouterStaticEntriesOutput)

	err = client.CallMethod(nil, "AddRouterStaticEntries", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DeleteRouterStaticEntries(input *DeleteRouterStaticEntriesInput) (output *DeleteRouterStaticEntriesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DeleteRouterStaticEntriesOutput)

	err = client.CallMethod(nil, "DeleteRouterStaticEntries", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) ModifyRouterStaticEntryAttributes(input *ModifyRouterStaticEntryAttributesInput) (output *ModifyRouterStaticEntryAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(ModifyRouterStaticEntryAttributesOutput)

	err = client.CallMethod(nil, "ModifyRouterStaticEntryAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RouterService) DescribeRouterStaticEntries(input *DescribeRouterStaticEntriesInput) (output *DescribeRouterStaticEntriesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
	)
	output = new(DescribeRouterStaticEntriesOutput)

	err = client.CallMethod(nil, "DescribeRouterStaticEntries", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
