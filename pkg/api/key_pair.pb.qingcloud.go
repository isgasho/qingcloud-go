// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: key_pair.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type KeyPairServiceInterface interface {
	DescribeKeyPairs(in *DescribeKeyPairsInput) (out *DescribeKeyPairsOutput, err error)
	CreateKeyPair(in *CreateKeyPairInput) (out *CreateKeyPairOutput, err error)
	DeleteKeyPairs(in *DeleteKeyPairsInput) (out *DeleteKeyPairsOutput, err error)
	AttachKeyPairs(in *AttachKeyPairsInput) (out *AttachKeyPairsOutput, err error)
	DetachKeyPairs(in *DetachKeyPairsInput) (out *DetachKeyPairsOutput, err error)
	ModifyKeyPairAttributes(in *ModifyKeyPairAttributesInput) (out *ModifyKeyPairAttributesOutput, err error)
}

type KeyPairService struct {
	ServerInfo       *ServerInfo
	Properties       *KeyPairServiceProperties
	LastResponseBody string
}

func NewKeyPairService(server *ServerInfo, serviceProp *KeyPairServiceProperties) (p *KeyPairService) {
	return &KeyPairService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *KeyPairService) DescribeKeyPairs(input *DescribeKeyPairsInput) (output *DescribeKeyPairsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeKeyPairsOutput)

	err = client.CallMethod(nil, "DescribeKeyPairs", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *KeyPairService) CreateKeyPair(input *CreateKeyPairInput) (output *CreateKeyPairOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateKeyPairOutput)

	err = client.CallMethod(nil, "CreateKeyPair", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *KeyPairService) DeleteKeyPairs(input *DeleteKeyPairsInput) (output *DeleteKeyPairsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteKeyPairsOutput)

	err = client.CallMethod(nil, "DeleteKeyPairs", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *KeyPairService) AttachKeyPairs(input *AttachKeyPairsInput) (output *AttachKeyPairsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AttachKeyPairsOutput)

	err = client.CallMethod(nil, "AttachKeyPairs", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *KeyPairService) DetachKeyPairs(input *DetachKeyPairsInput) (output *DetachKeyPairsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DetachKeyPairsOutput)

	err = client.CallMethod(nil, "DetachKeyPairs", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *KeyPairService) ModifyKeyPairAttributes(input *ModifyKeyPairAttributesInput) (output *ModifyKeyPairAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyKeyPairAttributesOutput)

	err = client.CallMethod(nil, "ModifyKeyPairAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
