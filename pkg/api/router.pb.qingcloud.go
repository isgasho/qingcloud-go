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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRoutersInput)
	}

	output = new(DescribeRoutersOutput)
	err = client.CallMethod("DescribeRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateRoutersInput)
	}

	output = new(CreateRoutersOutput)
	err = client.CallMethod("CreateRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteRoutersInput)
	}

	output = new(DeleteRoutersOutput)
	err = client.CallMethod("DeleteRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(UpdateRoutersInput)
	}

	output = new(UpdateRoutersOutput)
	err = client.CallMethod("UpdateRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(PowerOffRoutersInput)
	}

	output = new(PowerOffRoutersOutput)
	err = client.CallMethod("PowerOffRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(PowerOnRoutersInput)
	}

	output = new(PowerOnRoutersOutput)
	err = client.CallMethod("PowerOnRouters", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(JoinRouterInput)
	}

	output = new(JoinRouterOutput)
	err = client.CallMethod("JoinRouter", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(LeaveRouterInput)
	}

	output = new(LeaveRouterOutput)
	err = client.CallMethod("LeaveRouter", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyRouterAttributesInput)
	}

	output = new(ModifyRouterAttributesOutput)
	err = client.CallMethod("ModifyRouterAttributes", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRouterStaticsInput)
	}

	output = new(DescribeRouterStaticsOutput)
	err = client.CallMethod("DescribeRouterStatics", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddRouterStaticsInput)
	}

	output = new(AddRouterStaticsOutput)
	err = client.CallMethod("AddRouterStatics", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyRouterStaticAttributesInput)
	}

	output = new(ModifyRouterStaticAttributesOutput)
	err = client.CallMethod("ModifyRouterStaticAttributes", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteRouterStaticsInput)
	}

	output = new(DeleteRouterStaticsOutput)
	err = client.CallMethod("DeleteRouterStatics", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CopyRouterStaticsInput)
	}

	output = new(CopyRouterStaticsOutput)
	err = client.CallMethod("CopyRouterStatics", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRouterVxnetsInput)
	}

	output = new(DescribeRouterVxnetsOutput)
	err = client.CallMethod("DescribeRouterVxnets", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddRouterStaticEntriesInput)
	}

	output = new(AddRouterStaticEntriesOutput)
	err = client.CallMethod("AddRouterStaticEntries", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteRouterStaticEntriesInput)
	}

	output = new(DeleteRouterStaticEntriesOutput)
	err = client.CallMethod("DeleteRouterStaticEntries", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyRouterStaticEntryAttributesInput)
	}

	output = new(ModifyRouterStaticEntryAttributesOutput)
	err = client.CallMethod("ModifyRouterStaticEntryAttributes", "GET", input, output, nil)
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
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRouterStaticEntriesInput)
	}

	output = new(DescribeRouterStaticEntriesOutput)
	err = client.CallMethod("DescribeRouterStaticEntries", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
