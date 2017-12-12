// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: vxnet.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type VxnetServiceInterface interface {
	DescribeVxnets(in *DescribeVxnetsInput) (out *DescribeVxnetsOutput, err error)
	CreateVxnets(in *CreateVxnetsInput) (out *CreateVxnetsOutput, err error)
	DeleteVxnets(in *DeleteVxnetsInput) (out *DeleteVxnetsOutput, err error)
	JoinVxnet(in *JoinVxnetInput) (out *JoinVxnetOutput, err error)
	LeaveVxnet(in *LeaveVxnetInput) (out *LeaveVxnetOutput, err error)
	ModifyVxnetAttributes(in *ModifyVxnetAttributesInput) (out *ModifyVxnetAttributesOutput, err error)
	DescribeVxnetInstances(in *DescribeVxnetInstancesInput) (out *DescribeVxnetInstancesOutput, err error)
}

type VxnetService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewVxnetService(server *ServerInfo) (p *VxnetService) {
	return &VxnetService{
		ServerInfo: server,
	}
}

func (p *VxnetService) DescribeVxnets(input *DescribeVxnetsInput) (output *DescribeVxnetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeVxnetsInput)
	}

	output = new(DescribeVxnetsOutput)
	err = client.CallMethod("DescribeVxnets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) CreateVxnets(input *CreateVxnetsInput) (output *CreateVxnetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateVxnetsInput)
	}

	output = new(CreateVxnetsOutput)
	err = client.CallMethod("CreateVxnets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) DeleteVxnets(input *DeleteVxnetsInput) (output *DeleteVxnetsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteVxnetsInput)
	}

	output = new(DeleteVxnetsOutput)
	err = client.CallMethod("DeleteVxnets", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) JoinVxnet(input *JoinVxnetInput) (output *JoinVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(JoinVxnetInput)
	}

	output = new(JoinVxnetOutput)
	err = client.CallMethod("JoinVxnet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) LeaveVxnet(input *LeaveVxnetInput) (output *LeaveVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(LeaveVxnetInput)
	}

	output = new(LeaveVxnetOutput)
	err = client.CallMethod("LeaveVxnet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) ModifyVxnetAttributes(input *ModifyVxnetAttributesInput) (output *ModifyVxnetAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyVxnetAttributesInput)
	}

	output = new(ModifyVxnetAttributesOutput)
	err = client.CallMethod("ModifyVxnetAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *VxnetService) DescribeVxnetInstances(input *DescribeVxnetInstancesInput) (output *DescribeVxnetInstancesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeVxnetInstancesInput)
	}

	output = new(DescribeVxnetInstancesOutput)
	err = client.CallMethod("DescribeVxnetInstances", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
