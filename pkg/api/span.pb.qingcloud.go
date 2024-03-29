// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: span.proto

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
	TypeInfoMap["CreateSpanInput"] = reflect.TypeOf((*CreateSpanInput)(nil))
	TypeInfoMap["CreateSpanOutput"] = reflect.TypeOf((*CreateSpanOutput)(nil))
	TypeInfoMap["DescribeSpansInput"] = reflect.TypeOf((*DescribeSpansInput)(nil))
	TypeInfoMap["DescribeSpansOutput"] = reflect.TypeOf((*DescribeSpansOutput)(nil))
	TypeInfoMap["DeleteSpansInput"] = reflect.TypeOf((*DeleteSpansInput)(nil))
	TypeInfoMap["DeleteSpansOutput"] = reflect.TypeOf((*DeleteSpansOutput)(nil))
	TypeInfoMap["AddSpanMembersInput"] = reflect.TypeOf((*AddSpanMembersInput)(nil))
	TypeInfoMap["AddSpanMembersOutput"] = reflect.TypeOf((*AddSpanMembersOutput)(nil))
	TypeInfoMap["RemoveSpanMembersInput"] = reflect.TypeOf((*RemoveSpanMembersInput)(nil))
	TypeInfoMap["RemoveSpanMembersOutput"] = reflect.TypeOf((*RemoveSpanMembersOutput)(nil))
	TypeInfoMap["ModifySpanAttributesInput"] = reflect.TypeOf((*ModifySpanAttributesInput)(nil))
	TypeInfoMap["ModifySpanAttributesOutput"] = reflect.TypeOf((*ModifySpanAttributesOutput)(nil))
	TypeInfoMap["UpdateSpanInput"] = reflect.TypeOf((*UpdateSpanInput)(nil))
	TypeInfoMap["UpdateSpanOutput"] = reflect.TypeOf((*UpdateSpanOutput)(nil))
}

type SpanServiceInterface interface {
	CreateSpan(in *CreateSpanInput) (out *CreateSpanOutput, err error)
	DescribeSpans(in *DescribeSpansInput) (out *DescribeSpansOutput, err error)
	DeleteSpans(in *DeleteSpansInput) (out *DeleteSpansOutput, err error)
	AddSpanMembers(in *AddSpanMembersInput) (out *AddSpanMembersOutput, err error)
	RemoveSpanMembers(in *RemoveSpanMembersInput) (out *RemoveSpanMembersOutput, err error)
	ModifySpanAttributes(in *ModifySpanAttributesInput) (out *ModifySpanAttributesOutput, err error)
	UpdateSpan(in *UpdateSpanInput) (out *UpdateSpanOutput, err error)
}

type SpanService struct {
	ServerInfo *ServerInfo
}

func NewSpanService(server *ServerInfo) (p *SpanService) {
	return &SpanService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["CreateSpan"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "CreateSpan",
		InputTypeName:  "CreateSpanInput",
		OutputTypeName: "CreateSpanOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*CreateSpanInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateSpanOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeSpans"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "DescribeSpans",
		InputTypeName:  "DescribeSpansInput",
		OutputTypeName: "DescribeSpansOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*DescribeSpansInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeSpansOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteSpans"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "DeleteSpans",
		InputTypeName:  "DeleteSpansInput",
		OutputTypeName: "DeleteSpansOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*DeleteSpansInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteSpansOutput)(nil)),
	}
	ServiceApiSpecMap["AddSpanMembers"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "AddSpanMembers",
		InputTypeName:  "AddSpanMembersInput",
		OutputTypeName: "AddSpanMembersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*AddSpanMembersInput)(nil)),
		OutputType:  reflect.TypeOf((*AddSpanMembersOutput)(nil)),
	}
	ServiceApiSpecMap["RemoveSpanMembers"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "RemoveSpanMembers",
		InputTypeName:  "RemoveSpanMembersInput",
		OutputTypeName: "RemoveSpanMembersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*RemoveSpanMembersInput)(nil)),
		OutputType:  reflect.TypeOf((*RemoveSpanMembersOutput)(nil)),
	}
	ServiceApiSpecMap["ModifySpanAttributes"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "ModifySpanAttributes",
		InputTypeName:  "ModifySpanAttributesInput",
		OutputTypeName: "ModifySpanAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*ModifySpanAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifySpanAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["UpdateSpan"] = ServiceApiSpec{
		ServiceName:    "SpanService",
		ActionName:     "UpdateSpan",
		InputTypeName:  "UpdateSpanInput",
		OutputTypeName: "UpdateSpanOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SpanService)(nil)),
		InputType:   reflect.TypeOf((*UpdateSpanInput)(nil)),
		OutputType:  reflect.TypeOf((*UpdateSpanOutput)(nil)),
	}
}

func (p *SpanService) CreateSpan(input *CreateSpanInput) (output *CreateSpanOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSpanInput)
	}

	output = new(CreateSpanOutput)
	err = client.CallMethod("CreateSpan", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) DescribeSpans(input *DescribeSpansInput) (output *DescribeSpansOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSpansInput)
	}

	output = new(DescribeSpansOutput)
	err = client.CallMethod("DescribeSpans", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) DeleteSpans(input *DeleteSpansInput) (output *DeleteSpansOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSpansInput)
	}

	output = new(DeleteSpansOutput)
	err = client.CallMethod("DeleteSpans", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) AddSpanMembers(input *AddSpanMembersInput) (output *AddSpanMembersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddSpanMembersInput)
	}

	output = new(AddSpanMembersOutput)
	err = client.CallMethod("AddSpanMembers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) RemoveSpanMembers(input *RemoveSpanMembersInput) (output *RemoveSpanMembersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RemoveSpanMembersInput)
	}

	output = new(RemoveSpanMembersOutput)
	err = client.CallMethod("RemoveSpanMembers", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) ModifySpanAttributes(input *ModifySpanAttributesInput) (output *ModifySpanAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifySpanAttributesInput)
	}

	output = new(ModifySpanAttributesOutput)
	err = client.CallMethod("ModifySpanAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SpanService) UpdateSpan(input *UpdateSpanInput) (output *UpdateSpanOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(UpdateSpanInput)
	}

	output = new(UpdateSpanOutput)
	err = client.CallMethod("UpdateSpan", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
