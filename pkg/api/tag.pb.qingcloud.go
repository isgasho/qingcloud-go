// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: tag.proto

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
	TypeInfoMap["DescribeTagsInput"] = reflect.TypeOf((*DescribeTagsInput)(nil))
	TypeInfoMap["DescribeTagsOutput"] = reflect.TypeOf((*DescribeTagsOutput)(nil))
	TypeInfoMap["CreateTagInput"] = reflect.TypeOf((*CreateTagInput)(nil))
	TypeInfoMap["CreateTagOutput"] = reflect.TypeOf((*CreateTagOutput)(nil))
	TypeInfoMap["DeleteTagsInput"] = reflect.TypeOf((*DeleteTagsInput)(nil))
	TypeInfoMap["DeleteTagsOutput"] = reflect.TypeOf((*DeleteTagsOutput)(nil))
	TypeInfoMap["ModifyTagAttributesInput"] = reflect.TypeOf((*ModifyTagAttributesInput)(nil))
	TypeInfoMap["ModifyTagAttributesOutput"] = reflect.TypeOf((*ModifyTagAttributesOutput)(nil))
	TypeInfoMap["AttachTagsInput"] = reflect.TypeOf((*AttachTagsInput)(nil))
	TypeInfoMap["AttachTagsOutput"] = reflect.TypeOf((*AttachTagsOutput)(nil))
	TypeInfoMap["DetachTagsInput"] = reflect.TypeOf((*DetachTagsInput)(nil))
	TypeInfoMap["DetachTagsOutput"] = reflect.TypeOf((*DetachTagsOutput)(nil))
}

type TagServiceInterface interface {
	DescribeTags(in *DescribeTagsInput) (out *DescribeTagsOutput, err error)
	CreateTag(in *CreateTagInput) (out *CreateTagOutput, err error)
	DeleteTags(in *DeleteTagsInput) (out *DeleteTagsOutput, err error)
	ModifyTagAttributes(in *ModifyTagAttributesInput) (out *ModifyTagAttributesOutput, err error)
	AttachTags(in *AttachTagsInput) (out *AttachTagsOutput, err error)
	DetachTags(in *DetachTagsInput) (out *DetachTagsOutput, err error)
}

type TagService struct {
	ServerInfo *ServerInfo
}

func NewTagService(server *ServerInfo) (p *TagService) {
	return &TagService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["DescribeTags"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "DescribeTags",
		InputTypeName:  "DescribeTagsInput",
		OutputTypeName: "DescribeTagsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*DescribeTagsInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeTagsOutput)(nil)),
	}
	ServiceApiSpecMap["CreateTag"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "CreateTag",
		InputTypeName:  "CreateTagInput",
		OutputTypeName: "CreateTagOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*CreateTagInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateTagOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteTags"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "DeleteTags",
		InputTypeName:  "DeleteTagsInput",
		OutputTypeName: "DeleteTagsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*DeleteTagsInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteTagsOutput)(nil)),
	}
	ServiceApiSpecMap["ModifyTagAttributes"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "ModifyTagAttributes",
		InputTypeName:  "ModifyTagAttributesInput",
		OutputTypeName: "ModifyTagAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*ModifyTagAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifyTagAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["AttachTags"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "AttachTags",
		InputTypeName:  "AttachTagsInput",
		OutputTypeName: "AttachTagsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*AttachTagsInput)(nil)),
		OutputType:  reflect.TypeOf((*AttachTagsOutput)(nil)),
	}
	ServiceApiSpecMap["DetachTags"] = ServiceApiSpec{
		ServiceName:    "TagService",
		ActionName:     "DetachTags",
		InputTypeName:  "DetachTagsInput",
		OutputTypeName: "DetachTagsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*TagService)(nil)),
		InputType:   reflect.TypeOf((*DetachTagsInput)(nil)),
		OutputType:  reflect.TypeOf((*DetachTagsOutput)(nil)),
	}
}

func (p *TagService) DescribeTags(input *DescribeTagsInput) (output *DescribeTagsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeTagsInput)
	}

	output = new(DescribeTagsOutput)
	err = client.CallMethod("DescribeTags", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) CreateTag(input *CreateTagInput) (output *CreateTagOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateTagInput)
	}

	output = new(CreateTagOutput)
	err = client.CallMethod("CreateTag", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) DeleteTags(input *DeleteTagsInput) (output *DeleteTagsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteTagsInput)
	}

	output = new(DeleteTagsOutput)
	err = client.CallMethod("DeleteTags", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) ModifyTagAttributes(input *ModifyTagAttributesInput) (output *ModifyTagAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyTagAttributesInput)
	}

	output = new(ModifyTagAttributesOutput)
	err = client.CallMethod("ModifyTagAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) AttachTags(input *AttachTagsInput) (output *AttachTagsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AttachTagsInput)
	}

	output = new(AttachTagsOutput)
	err = client.CallMethod("AttachTags", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *TagService) DetachTags(input *DetachTagsInput) (output *DetachTagsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DetachTagsInput)
	}

	output = new(DetachTagsOutput)
	err = client.CallMethod("DetachTags", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
