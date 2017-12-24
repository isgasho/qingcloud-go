// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: misc.proto

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

type MiscServiceInterface interface {
	GrantQuotaIndep(in *GrantQuotaIndepInput) (out *GrantQuotaIndepOutput, err error)
	RevokeQuotaIndep(in *RevokeQuotaIndepInput) (out *RevokeQuotaIndepOutput, err error)
	GetQuotaLeft(in *GetQuotaLeftInput) (out *GetQuotaLeftOutput, err error)
}

type MiscService struct {
	ServerInfo *ServerInfo
}

func NewMiscService(server *ServerInfo) (p *MiscService) {
	return &MiscService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["GrantQuotaIndep"] = ServiceApiSpec{
		ServiceName:    "MiscService",
		ActionName:     "GrantQuotaIndep",
		InputTypeName:  "GrantQuotaIndepInput",
		OutputTypeName: "GrantQuotaIndepOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MiscService)(nil)),
		InputType:   reflect.TypeOf((*GrantQuotaIndepInput)(nil)),
		OutputType:  reflect.TypeOf((*GrantQuotaIndepOutput)(nil)),
	}
	ServiceApiSpecMap["RevokeQuotaIndep"] = ServiceApiSpec{
		ServiceName:    "MiscService",
		ActionName:     "RevokeQuotaIndep",
		InputTypeName:  "RevokeQuotaIndepInput",
		OutputTypeName: "RevokeQuotaIndepOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MiscService)(nil)),
		InputType:   reflect.TypeOf((*RevokeQuotaIndepInput)(nil)),
		OutputType:  reflect.TypeOf((*RevokeQuotaIndepOutput)(nil)),
	}
	ServiceApiSpecMap["GetQuotaLeft"] = ServiceApiSpec{
		ServiceName:    "MiscService",
		ActionName:     "GetQuotaLeft",
		InputTypeName:  "GetQuotaLeftInput",
		OutputTypeName: "GetQuotaLeftOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MiscService)(nil)),
		InputType:   reflect.TypeOf((*GetQuotaLeftInput)(nil)),
		OutputType:  reflect.TypeOf((*GetQuotaLeftOutput)(nil)),
	}
}

func (p *MiscService) GrantQuotaIndep(input *GrantQuotaIndepInput) (output *GrantQuotaIndepOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GrantQuotaIndepInput)
	}

	output = new(GrantQuotaIndepOutput)
	err = client.CallMethod("GrantQuotaIndep", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MiscService) RevokeQuotaIndep(input *RevokeQuotaIndepInput) (output *RevokeQuotaIndepOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RevokeQuotaIndepInput)
	}

	output = new(RevokeQuotaIndepOutput)
	err = client.CallMethod("RevokeQuotaIndep", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MiscService) GetQuotaLeft(input *GetQuotaLeftInput) (output *GetQuotaLeftOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GetQuotaLeftInput)
	}

	output = new(GetQuotaLeftOutput)
	err = client.CallMethod("GetQuotaLeft", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
