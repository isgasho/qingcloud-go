// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: hadoop.proto

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
	TypeInfoMap["AddHadoopNodesInput"] = reflect.TypeOf((*AddHadoopNodesInput)(nil))
	TypeInfoMap["AddHadoopNodesOutput"] = reflect.TypeOf((*AddHadoopNodesOutput)(nil))
	TypeInfoMap["DeleteHadoopNodesInput"] = reflect.TypeOf((*DeleteHadoopNodesInput)(nil))
	TypeInfoMap["DeleteHadoopNodesOutput"] = reflect.TypeOf((*DeleteHadoopNodesOutput)(nil))
	TypeInfoMap["StartHadoopsInput"] = reflect.TypeOf((*StartHadoopsInput)(nil))
	TypeInfoMap["StartHadoopsOutput"] = reflect.TypeOf((*StartHadoopsOutput)(nil))
	TypeInfoMap["StopHadoopsInput"] = reflect.TypeOf((*StopHadoopsInput)(nil))
	TypeInfoMap["StopHadoopsOutput"] = reflect.TypeOf((*StopHadoopsOutput)(nil))
}

type HadoopServiceInterface interface {
	AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error)
	DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error)
	StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error)
	StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error)
}

type HadoopService struct {
	ServerInfo *ServerInfo
}

func NewHadoopService(server *ServerInfo) (p *HadoopService) {
	return &HadoopService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["AddHadoopNodes"] = ServiceApiSpec{
		ServiceName:    "HadoopService",
		ActionName:     "AddHadoopNodes",
		InputTypeName:  "AddHadoopNodesInput",
		OutputTypeName: "AddHadoopNodesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*HadoopService)(nil)),
		InputType:   reflect.TypeOf((*AddHadoopNodesInput)(nil)),
		OutputType:  reflect.TypeOf((*AddHadoopNodesOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteHadoopNodes"] = ServiceApiSpec{
		ServiceName:    "HadoopService",
		ActionName:     "DeleteHadoopNodes",
		InputTypeName:  "DeleteHadoopNodesInput",
		OutputTypeName: "DeleteHadoopNodesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*HadoopService)(nil)),
		InputType:   reflect.TypeOf((*DeleteHadoopNodesInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteHadoopNodesOutput)(nil)),
	}
	ServiceApiSpecMap["StartHadoops"] = ServiceApiSpec{
		ServiceName:    "HadoopService",
		ActionName:     "StartHadoops",
		InputTypeName:  "StartHadoopsInput",
		OutputTypeName: "StartHadoopsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*HadoopService)(nil)),
		InputType:   reflect.TypeOf((*StartHadoopsInput)(nil)),
		OutputType:  reflect.TypeOf((*StartHadoopsOutput)(nil)),
	}
	ServiceApiSpecMap["StopHadoops"] = ServiceApiSpec{
		ServiceName:    "HadoopService",
		ActionName:     "StopHadoops",
		InputTypeName:  "StopHadoopsInput",
		OutputTypeName: "StopHadoopsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*HadoopService)(nil)),
		InputType:   reflect.TypeOf((*StopHadoopsInput)(nil)),
		OutputType:  reflect.TypeOf((*StopHadoopsOutput)(nil)),
	}
}

func (p *HadoopService) AddHadoopNodes(input *AddHadoopNodesInput) (output *AddHadoopNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddHadoopNodesInput)
	}

	output = new(AddHadoopNodesOutput)
	err = client.CallMethod("AddHadoopNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) DeleteHadoopNodes(input *DeleteHadoopNodesInput) (output *DeleteHadoopNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteHadoopNodesInput)
	}

	output = new(DeleteHadoopNodesOutput)
	err = client.CallMethod("DeleteHadoopNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) StartHadoops(input *StartHadoopsInput) (output *StartHadoopsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StartHadoopsInput)
	}

	output = new(StartHadoopsOutput)
	err = client.CallMethod("StartHadoops", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) StopHadoops(input *StopHadoopsInput) (output *StopHadoopsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StopHadoopsInput)
	}

	output = new(StopHadoopsOutput)
	err = client.CallMethod("StopHadoops", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
