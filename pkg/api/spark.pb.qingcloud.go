// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: spark.proto

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

type SparkServiceInterface interface {
	CreateSpark(in *CreateSparkInput) (out *CreateSparkOutput, err error)
	DescribeSparks(in *DescribeSparksInput) (out *DescribeSparksOutput, err error)
	AddSparkNodes(in *AddSparkNodesInput) (out *AddSparkNodesOutput, err error)
	DeleteSparkNodes(in *DeleteSparkNodesInput) (out *DeleteSparkNodesOutput, err error)
	StartSparks(in *StartSparksInput) (out *StartSparksOutput, err error)
	StopSparks(in *StopSparksInput) (out *StopSparksOutput, err error)
	DeleteSparks(in *DeleteSparksInput) (out *DeleteSparksOutput, err error)
}

type SparkService struct {
	ServerInfo *ServerInfo
}

func NewSparkService(server *ServerInfo) (p *SparkService) {
	return &SparkService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["CreateSpark"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "CreateSpark",
		InputTypeName:  "CreateSparkInput",
		OutputTypeName: "CreateSparkOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*CreateSparkInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateSparkOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeSparks"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "DescribeSparks",
		InputTypeName:  "DescribeSparksInput",
		OutputTypeName: "DescribeSparksOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*DescribeSparksInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeSparksOutput)(nil)),
	}
	ServiceApiSpecMap["AddSparkNodes"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "AddSparkNodes",
		InputTypeName:  "AddSparkNodesInput",
		OutputTypeName: "AddSparkNodesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*AddSparkNodesInput)(nil)),
		OutputType:  reflect.TypeOf((*AddSparkNodesOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteSparkNodes"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "DeleteSparkNodes",
		InputTypeName:  "DeleteSparkNodesInput",
		OutputTypeName: "DeleteSparkNodesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*DeleteSparkNodesInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteSparkNodesOutput)(nil)),
	}
	ServiceApiSpecMap["StartSparks"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "StartSparks",
		InputTypeName:  "StartSparksInput",
		OutputTypeName: "StartSparksOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*StartSparksInput)(nil)),
		OutputType:  reflect.TypeOf((*StartSparksOutput)(nil)),
	}
	ServiceApiSpecMap["StopSparks"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "StopSparks",
		InputTypeName:  "StopSparksInput",
		OutputTypeName: "StopSparksOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*StopSparksInput)(nil)),
		OutputType:  reflect.TypeOf((*StopSparksOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteSparks"] = ServiceApiSpec{
		ServiceName:    "SparkService",
		ActionName:     "DeleteSparks",
		InputTypeName:  "DeleteSparksInput",
		OutputTypeName: "DeleteSparksOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*SparkService)(nil)),
		InputType:   reflect.TypeOf((*DeleteSparksInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteSparksOutput)(nil)),
	}
}

func (p *SparkService) CreateSpark(input *CreateSparkInput) (output *CreateSparkOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateSparkInput)
	}

	output = new(CreateSparkOutput)
	err = client.CallMethod("CreateSpark", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) DescribeSparks(input *DescribeSparksInput) (output *DescribeSparksOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeSparksInput)
	}

	output = new(DescribeSparksOutput)
	err = client.CallMethod("DescribeSparks", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) AddSparkNodes(input *AddSparkNodesInput) (output *AddSparkNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddSparkNodesInput)
	}

	output = new(AddSparkNodesOutput)
	err = client.CallMethod("AddSparkNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) DeleteSparkNodes(input *DeleteSparkNodesInput) (output *DeleteSparkNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSparkNodesInput)
	}

	output = new(DeleteSparkNodesOutput)
	err = client.CallMethod("DeleteSparkNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) StartSparks(input *StartSparksInput) (output *StartSparksOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StartSparksInput)
	}

	output = new(StartSparksOutput)
	err = client.CallMethod("StartSparks", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) StopSparks(input *StopSparksInput) (output *StopSparksOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StopSparksInput)
	}

	output = new(StopSparksOutput)
	err = client.CallMethod("StopSparks", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *SparkService) DeleteSparks(input *DeleteSparksInput) (output *DeleteSparksOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteSparksInput)
	}

	output = new(DeleteSparksOutput)
	err = client.CallMethod("DeleteSparks", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
