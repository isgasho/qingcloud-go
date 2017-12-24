// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: router.proto

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
	GetVPNCerts(in *GetVPNCertsInput) (out *GetVPNCertsOutput, err error)
}

type RouterService struct {
	ServerInfo *ServerInfo
}

func NewRouterService(server *ServerInfo) (p *RouterService) {
	return &RouterService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["DescribeRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DescribeRouters",
		InputTypeName:  "DescribeRoutersInput",
		OutputTypeName: "DescribeRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["CreateRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "CreateRouters",
		InputTypeName:  "CreateRoutersInput",
		OutputTypeName: "CreateRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*CreateRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DeleteRouters",
		InputTypeName:  "DeleteRoutersInput",
		OutputTypeName: "DeleteRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DeleteRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["UpdateRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "UpdateRouters",
		InputTypeName:  "UpdateRoutersInput",
		OutputTypeName: "UpdateRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*UpdateRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*UpdateRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["PowerOffRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "PowerOffRouters",
		InputTypeName:  "PowerOffRoutersInput",
		OutputTypeName: "PowerOffRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*PowerOffRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*PowerOffRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["PowerOnRouters"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "PowerOnRouters",
		InputTypeName:  "PowerOnRoutersInput",
		OutputTypeName: "PowerOnRoutersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*PowerOnRoutersInput)(nil)),
		OutputType:  reflect.TypeOf((*PowerOnRoutersOutput)(nil)),
	}
	ServiceApiSpecMap["JoinRouter"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "JoinRouter",
		InputTypeName:  "JoinRouterInput",
		OutputTypeName: "JoinRouterOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*JoinRouterInput)(nil)),
		OutputType:  reflect.TypeOf((*JoinRouterOutput)(nil)),
	}
	ServiceApiSpecMap["LeaveRouter"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "LeaveRouter",
		InputTypeName:  "LeaveRouterInput",
		OutputTypeName: "LeaveRouterOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*LeaveRouterInput)(nil)),
		OutputType:  reflect.TypeOf((*LeaveRouterOutput)(nil)),
	}
	ServiceApiSpecMap["ModifyRouterAttributes"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "ModifyRouterAttributes",
		InputTypeName:  "ModifyRouterAttributesInput",
		OutputTypeName: "ModifyRouterAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*ModifyRouterAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifyRouterAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeRouterStatics"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DescribeRouterStatics",
		InputTypeName:  "DescribeRouterStaticsInput",
		OutputTypeName: "DescribeRouterStaticsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRouterStaticsInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRouterStaticsOutput)(nil)),
	}
	ServiceApiSpecMap["AddRouterStatics"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "AddRouterStatics",
		InputTypeName:  "AddRouterStaticsInput",
		OutputTypeName: "AddRouterStaticsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*AddRouterStaticsInput)(nil)),
		OutputType:  reflect.TypeOf((*AddRouterStaticsOutput)(nil)),
	}
	ServiceApiSpecMap["ModifyRouterStaticAttributes"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "ModifyRouterStaticAttributes",
		InputTypeName:  "ModifyRouterStaticAttributesInput",
		OutputTypeName: "ModifyRouterStaticAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*ModifyRouterStaticAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifyRouterStaticAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteRouterStatics"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DeleteRouterStatics",
		InputTypeName:  "DeleteRouterStaticsInput",
		OutputTypeName: "DeleteRouterStaticsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DeleteRouterStaticsInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteRouterStaticsOutput)(nil)),
	}
	ServiceApiSpecMap["CopyRouterStatics"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "CopyRouterStatics",
		InputTypeName:  "CopyRouterStaticsInput",
		OutputTypeName: "CopyRouterStaticsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*CopyRouterStaticsInput)(nil)),
		OutputType:  reflect.TypeOf((*CopyRouterStaticsOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeRouterVxnets"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DescribeRouterVxnets",
		InputTypeName:  "DescribeRouterVxnetsInput",
		OutputTypeName: "DescribeRouterVxnetsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRouterVxnetsInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRouterVxnetsOutput)(nil)),
	}
	ServiceApiSpecMap["AddRouterStaticEntries"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "AddRouterStaticEntries",
		InputTypeName:  "AddRouterStaticEntriesInput",
		OutputTypeName: "AddRouterStaticEntriesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*AddRouterStaticEntriesInput)(nil)),
		OutputType:  reflect.TypeOf((*AddRouterStaticEntriesOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteRouterStaticEntries"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DeleteRouterStaticEntries",
		InputTypeName:  "DeleteRouterStaticEntriesInput",
		OutputTypeName: "DeleteRouterStaticEntriesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DeleteRouterStaticEntriesInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteRouterStaticEntriesOutput)(nil)),
	}
	ServiceApiSpecMap["ModifyRouterStaticEntryAttributes"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "ModifyRouterStaticEntryAttributes",
		InputTypeName:  "ModifyRouterStaticEntryAttributesInput",
		OutputTypeName: "ModifyRouterStaticEntryAttributesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*ModifyRouterStaticEntryAttributesInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifyRouterStaticEntryAttributesOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeRouterStaticEntries"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "DescribeRouterStaticEntries",
		InputTypeName:  "DescribeRouterStaticEntriesInput",
		OutputTypeName: "DescribeRouterStaticEntriesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRouterStaticEntriesInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRouterStaticEntriesOutput)(nil)),
	}
	ServiceApiSpecMap["GetVPNCerts"] = ServiceApiSpec{
		ServiceName:    "RouterService",
		ActionName:     "GetVPNCerts",
		InputTypeName:  "GetVPNCertsInput",
		OutputTypeName: "GetVPNCertsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RouterService)(nil)),
		InputType:   reflect.TypeOf((*GetVPNCertsInput)(nil)),
		OutputType:  reflect.TypeOf((*GetVPNCertsOutput)(nil)),
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

func (p *RouterService) GetVPNCerts(input *GetVPNCertsInput) (output *GetVPNCertsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GetVPNCertsInput)
	}

	output = new(GetVPNCertsOutput)
	err = client.CallMethod("GetVPNCerts", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
