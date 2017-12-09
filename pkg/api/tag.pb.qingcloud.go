// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: tag.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type TagServiceInterface interface {
	DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error)
	CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error)
	DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error)
	ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error)
	AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error)
	DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error)
}

type TagService struct {
	ServerInfo       *ServerInfo
	Properties       *TagServiceProperties
	LastResponseBody string
}

func NewTagService(server *ServerInfo, serviceProp *TagServiceProperties) (p *TagService) {
	return &TagService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *TagService) DescribeTags(input *DescribeTagsInput) (output *DescribeTagsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeTagsOutput)

	err = client.CallMethod(nil, "DescribeTags", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) CreateTag(input *CreateTagInput) (output *CreateTagOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateTagOutput)

	err = client.CallMethod(nil, "CreateTag", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) DeleteTags(input *DeleteTagsInput) (output *DeleteTagsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteTagsOutput)

	err = client.CallMethod(nil, "DeleteTags", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) ModifyTagAttributes(input *ModifyTagAttributesInput) (output *ModifyTagAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyTagAttributesOutput)

	err = client.CallMethod(nil, "ModifyTagAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) AttachTags(input *AttachTagsInput) (output *AttachTagsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AttachTagsOutput)

	err = client.CallMethod(nil, "AttachTags", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) DetachTags(input *DetachTagsInput) (output *DetachTagsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DetachTagsOutput)

	err = client.CallMethod(nil, "DetachTags", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
