// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: eip.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type EIPServiceInterface interface {
	DescribeEips(in *DescribeEipsInput) (out *DescribeEipsOutput, err error)
	AllocateEips(in *AllocateEipsInput) (out *AllocateEipsOutput, err error)
	ReleaseEips(in *ReleaseEipsInput) (out *ReleaseEipsOutput, err error)
	AssociateEip(in *AssociateEipInput) (out *AssociateEipOutput, err error)
	DissociateEips(in *DissociateEipsInput) (out *DissociateEipsOutput, err error)
	ChangeEipsBandwidth(in *ChangeEipsBandwidthInput) (out *ChangeEipsBandwidthOutput, err error)
	ChangeEipsBillingMode(in *ChangeEipsBillingModeInput) (out *ChangeEipsBillingModeOutput, err error)
	ModifyEipAttributes(in *ModifyEipAttributesInput) (out *ModifyEipAttributesOutput, err error)
}

type EIPService struct {
	Properties       *EIPServiceProperties
	LastResponseBody string
}

func NewEIPService(accessKeyId, secretAccessKey, zone string) (p *EIPService) {
	return &EIPService{
		Properties: &EIPServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *EIPService) DescribeEips(input *DescribeEipsInput) (output *DescribeEipsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeEipsOutput)

	err = client.CallMethod(nil, "DescribeEips", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) AllocateEips(input *AllocateEipsInput) (output *AllocateEipsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AllocateEipsOutput)

	err = client.CallMethod(nil, "AllocateEips", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) ReleaseEips(input *ReleaseEipsInput) (output *ReleaseEipsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ReleaseEipsOutput)

	err = client.CallMethod(nil, "ReleaseEips", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) AssociateEip(input *AssociateEipInput) (output *AssociateEipOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AssociateEipOutput)

	err = client.CallMethod(nil, "AssociateEip", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) DissociateEips(input *DissociateEipsInput) (output *DissociateEipsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DissociateEipsOutput)

	err = client.CallMethod(nil, "DissociateEips", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) ChangeEipsBandwidth(input *ChangeEipsBandwidthInput) (output *ChangeEipsBandwidthOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ChangeEipsBandwidthOutput)

	err = client.CallMethod(nil, "ChangeEipsBandwidth", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) ChangeEipsBillingMode(input *ChangeEipsBillingModeInput) (output *ChangeEipsBillingModeOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ChangeEipsBillingModeOutput)

	err = client.CallMethod(nil, "ChangeEipsBillingMode", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *EIPService) ModifyEipAttributes(input *ModifyEipAttributesInput) (output *ModifyEipAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyEipAttributesOutput)

	err = client.CallMethod(nil, "ModifyEipAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
