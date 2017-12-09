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
	Properties       *NicServiceProperties
	LastResponseBody string
}

func NewNicService(server *ServerInfo, serviceProp *NicServiceProperties) (p *NicService) {
	return &NicService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *NicService) CreateNics(input *CreateNicsInput) (output *CreateNicsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateNicsOutput)

	err = client.CallMethod(nil, "CreateNics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DescribeNics(input *DescribeNicsInput) (output *DescribeNicsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeNicsOutput)

	err = client.CallMethod(nil, "DescribeNics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) AttachNics(input *AttachNicsInput) (output *AttachNicsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AttachNicsOutput)

	err = client.CallMethod(nil, "AttachNics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DetachNics(input *DetachNicsInput) (output *DetachNicsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DetachNicsOutput)

	err = client.CallMethod(nil, "DetachNics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) ModifyNicAttributes(input *ModifyNicAttributesInput) (output *ModifyNicAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyNicAttributesOutput)

	err = client.CallMethod(nil, "ModifyNicAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *NicService) DeleteNics(input *DeleteNicsInput) (output *DeleteNicsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteNicsOutput)

	err = client.CallMethod(nil, "DeleteNics", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
