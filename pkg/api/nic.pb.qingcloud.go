// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: nic.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type NicServiceInterface interface {
	CreateNics(in *CreateNicsInput) (out *CreateNicsOutput, err error)
	DescribeNics(in *DescribeNicsInput) (out *DescribeNicsOutput, err error)
	AttachNics(in *AttachNicsInput) (out *AttachNicsOutput, err error)
	DetachNics(in *DetachNicsInput) (out *DetachNicsOutput, err error)
	ModifyNicAttributes(in *ModifyNicAttributesInput) (out *ModifyNicAttributesOutput, err error)
	DeleteNics(in *DeleteNicsInput) (out *DeleteNicsOutput, err error)
}

type NicService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewNicService(server *ServerInfo) (p *NicService) {
	return &NicService{
		ServerInfo: server,
	}
}

func (p *NicService) CreateNics(input *CreateNicsInput) (output *CreateNicsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateNicsInput)
	}

	output = new(CreateNicsOutput)
	err = client.CallMethod("CreateNics", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DescribeNics(input *DescribeNicsInput) (output *DescribeNicsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeNicsInput)
	}

	output = new(DescribeNicsOutput)
	err = client.CallMethod("DescribeNics", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) AttachNics(input *AttachNicsInput) (output *AttachNicsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AttachNicsInput)
	}

	output = new(AttachNicsOutput)
	err = client.CallMethod("AttachNics", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DetachNics(input *DetachNicsInput) (output *DetachNicsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DetachNicsInput)
	}

	output = new(DetachNicsOutput)
	err = client.CallMethod("DetachNics", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) ModifyNicAttributes(input *ModifyNicAttributesInput) (output *ModifyNicAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyNicAttributesInput)
	}

	output = new(ModifyNicAttributesOutput)
	err = client.CallMethod("ModifyNicAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DeleteNics(input *DeleteNicsInput) (output *DeleteNicsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteNicsInput)
	}

	output = new(DeleteNicsOutput)
	err = client.CallMethod("DeleteNics", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
