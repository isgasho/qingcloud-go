// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: cache.proto

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

type CacheServiceInterface interface {
	DescribeCaches(in *DescribeCachesInput) (out *DescribeCachesOutput, err error)
	CreateCache(in *CreateCacheInput) (out *CreateCacheOutput, err error)
	StopCaches(in *StopCachesInput) (out *StopCachesOutput, err error)
	StartCaches(in *StartCachesInput) (out *StartCachesOutput, err error)
	RestartCaches(in *RestartCachesInput) (out *RestartCachesOutput, err error)
	DeleteCaches(in *DeleteCachesInput) (out *DeleteCachesOutput, err error)
	ResizeCaches(in *ResizeCachesInput) (out *ResizeCachesOutput, err error)
	UpdateCache(in *UpdateCacheInput) (out *UpdateCacheOutput, err error)
	ChangeCacheVxnet(in *ChangeCacheVxnetInput) (out *ChangeCacheVxnetOutput, err error)
	ModifyCacheAttributes(in *ModifyCacheAttributesInput) (out *ModifyCacheAttributesOutput, err error)
	DescribeCacheNodes(in *DescribeCacheNodesInput) (out *DescribeCacheNodesOutput, err error)
	AddCacheNodes(in *AddCacheNodesInput) (out *AddCacheNodesOutput, err error)
	DeleteCacheNodes(in *DeleteCacheNodesInput) (out *DeleteCacheNodesOutput, err error)
	RestartCacheNodes(in *RestartCacheNodesInput) (out *RestartCacheNodesOutput, err error)
	ModifyCacheNodeAttributes(in *ModifyCacheNodeAttributesInput) (out *ModifyCacheNodeAttributesOutput, err error)
	CreateCacheFromSnapshot(in *CreateCacheFromSnapshotInput) (out *CreateCacheFromSnapshotOutput, err error)
	DescribeCacheParameterGroups(in *DescribeCacheParameterGroupsInput) (out *DescribeCacheParameterGroupsOutput, err error)
	CreateCacheParameterGroup(in *CreateCacheParameterGroupInput) (out *CreateCacheParameterGroupOutput, err error)
	ApplyCacheParameterGroup(in *ApplyCacheParameterGroupInput) (out *ApplyCacheParameterGroupOutput, err error)
	DeleteCacheParameterGroups(in *DeleteCacheParameterGroupsInput) (out *DeleteCacheParameterGroupsOutput, err error)
	ModifyCacheParameterGroupAttributes(in *ModifyCacheParameterGroupAttributesInput) (out *ModifyCacheParameterGroupAttributesOutput, err error)
	DescribeCacheParameters(in *DescribeCacheParametersInput) (out *DescribeCacheParametersOutput, err error)
	UpdateCacheParameters(in *UpdateCacheParametersInput) (out *UpdateCacheParametersOutput, err error)
	ResetCacheParameters(in *ResetCacheParametersInput) (out *ResetCacheParametersOutput, err error)
}

type CacheService struct {
	ServerInfo *ServerInfo
}

func NewCacheService(server *ServerInfo) (p *CacheService) {
	return &CacheService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["DescribeCaches"] = ServiceApiSpec{
		ActionName: "DescribeCaches",
		InputType:  reflect.TypeOf((*DescribeCachesInput)(nil)),
		OutputType: reflect.TypeOf((*DescribeCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["CreateCache"] = ServiceApiSpec{
		ActionName: "CreateCache",
		InputType:  reflect.TypeOf((*CreateCacheInput)(nil)),
		OutputType: reflect.TypeOf((*CreateCacheOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["StopCaches"] = ServiceApiSpec{
		ActionName: "StopCaches",
		InputType:  reflect.TypeOf((*StopCachesInput)(nil)),
		OutputType: reflect.TypeOf((*StopCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["StartCaches"] = ServiceApiSpec{
		ActionName: "StartCaches",
		InputType:  reflect.TypeOf((*StartCachesInput)(nil)),
		OutputType: reflect.TypeOf((*StartCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["RestartCaches"] = ServiceApiSpec{
		ActionName: "RestartCaches",
		InputType:  reflect.TypeOf((*RestartCachesInput)(nil)),
		OutputType: reflect.TypeOf((*RestartCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DeleteCaches"] = ServiceApiSpec{
		ActionName: "DeleteCaches",
		InputType:  reflect.TypeOf((*DeleteCachesInput)(nil)),
		OutputType: reflect.TypeOf((*DeleteCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ResizeCaches"] = ServiceApiSpec{
		ActionName: "ResizeCaches",
		InputType:  reflect.TypeOf((*ResizeCachesInput)(nil)),
		OutputType: reflect.TypeOf((*ResizeCachesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["UpdateCache"] = ServiceApiSpec{
		ActionName: "UpdateCache",
		InputType:  reflect.TypeOf((*UpdateCacheInput)(nil)),
		OutputType: reflect.TypeOf((*UpdateCacheOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ChangeCacheVxnet"] = ServiceApiSpec{
		ActionName: "ChangeCacheVxnet",
		InputType:  reflect.TypeOf((*ChangeCacheVxnetInput)(nil)),
		OutputType: reflect.TypeOf((*ChangeCacheVxnetOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ModifyCacheAttributes"] = ServiceApiSpec{
		ActionName: "ModifyCacheAttributes",
		InputType:  reflect.TypeOf((*ModifyCacheAttributesInput)(nil)),
		OutputType: reflect.TypeOf((*ModifyCacheAttributesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DescribeCacheNodes"] = ServiceApiSpec{
		ActionName: "DescribeCacheNodes",
		InputType:  reflect.TypeOf((*DescribeCacheNodesInput)(nil)),
		OutputType: reflect.TypeOf((*DescribeCacheNodesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["AddCacheNodes"] = ServiceApiSpec{
		ActionName: "AddCacheNodes",
		InputType:  reflect.TypeOf((*AddCacheNodesInput)(nil)),
		OutputType: reflect.TypeOf((*AddCacheNodesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DeleteCacheNodes"] = ServiceApiSpec{
		ActionName: "DeleteCacheNodes",
		InputType:  reflect.TypeOf((*DeleteCacheNodesInput)(nil)),
		OutputType: reflect.TypeOf((*DeleteCacheNodesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["RestartCacheNodes"] = ServiceApiSpec{
		ActionName: "RestartCacheNodes",
		InputType:  reflect.TypeOf((*RestartCacheNodesInput)(nil)),
		OutputType: reflect.TypeOf((*RestartCacheNodesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ModifyCacheNodeAttributes"] = ServiceApiSpec{
		ActionName: "ModifyCacheNodeAttributes",
		InputType:  reflect.TypeOf((*ModifyCacheNodeAttributesInput)(nil)),
		OutputType: reflect.TypeOf((*ModifyCacheNodeAttributesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["CreateCacheFromSnapshot"] = ServiceApiSpec{
		ActionName: "CreateCacheFromSnapshot",
		InputType:  reflect.TypeOf((*CreateCacheFromSnapshotInput)(nil)),
		OutputType: reflect.TypeOf((*CreateCacheFromSnapshotOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DescribeCacheParameterGroups"] = ServiceApiSpec{
		ActionName: "DescribeCacheParameterGroups",
		InputType:  reflect.TypeOf((*DescribeCacheParameterGroupsInput)(nil)),
		OutputType: reflect.TypeOf((*DescribeCacheParameterGroupsOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["CreateCacheParameterGroup"] = ServiceApiSpec{
		ActionName: "CreateCacheParameterGroup",
		InputType:  reflect.TypeOf((*CreateCacheParameterGroupInput)(nil)),
		OutputType: reflect.TypeOf((*CreateCacheParameterGroupOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ApplyCacheParameterGroup"] = ServiceApiSpec{
		ActionName: "ApplyCacheParameterGroup",
		InputType:  reflect.TypeOf((*ApplyCacheParameterGroupInput)(nil)),
		OutputType: reflect.TypeOf((*ApplyCacheParameterGroupOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DeleteCacheParameterGroups"] = ServiceApiSpec{
		ActionName: "DeleteCacheParameterGroups",
		InputType:  reflect.TypeOf((*DeleteCacheParameterGroupsInput)(nil)),
		OutputType: reflect.TypeOf((*DeleteCacheParameterGroupsOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ModifyCacheParameterGroupAttributes"] = ServiceApiSpec{
		ActionName: "ModifyCacheParameterGroupAttributes",
		InputType:  reflect.TypeOf((*ModifyCacheParameterGroupAttributesInput)(nil)),
		OutputType: reflect.TypeOf((*ModifyCacheParameterGroupAttributesOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["DescribeCacheParameters"] = ServiceApiSpec{
		ActionName: "DescribeCacheParameters",
		InputType:  reflect.TypeOf((*DescribeCacheParametersInput)(nil)),
		OutputType: reflect.TypeOf((*DescribeCacheParametersOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["UpdateCacheParameters"] = ServiceApiSpec{
		ActionName: "UpdateCacheParameters",
		InputType:  reflect.TypeOf((*UpdateCacheParametersInput)(nil)),
		OutputType: reflect.TypeOf((*UpdateCacheParametersOutput)(nil)),
		HttpMethod: "GET",
	}
	ServiceApiSpecMap["ResetCacheParameters"] = ServiceApiSpec{
		ActionName: "ResetCacheParameters",
		InputType:  reflect.TypeOf((*ResetCacheParametersInput)(nil)),
		OutputType: reflect.TypeOf((*ResetCacheParametersOutput)(nil)),
		HttpMethod: "GET",
	}
}

func (p *CacheService) DescribeCaches(input *DescribeCachesInput) (output *DescribeCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeCachesInput)
	}

	output = new(DescribeCachesOutput)
	err = client.CallMethod("DescribeCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCache(input *CreateCacheInput) (output *CreateCacheOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateCacheInput)
	}

	output = new(CreateCacheOutput)
	err = client.CallMethod("CreateCache", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) StopCaches(input *StopCachesInput) (output *StopCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StopCachesInput)
	}

	output = new(StopCachesOutput)
	err = client.CallMethod("StopCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) StartCaches(input *StartCachesInput) (output *StartCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StartCachesInput)
	}

	output = new(StartCachesOutput)
	err = client.CallMethod("StartCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) RestartCaches(input *RestartCachesInput) (output *RestartCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RestartCachesInput)
	}

	output = new(RestartCachesOutput)
	err = client.CallMethod("RestartCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCaches(input *DeleteCachesInput) (output *DeleteCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteCachesInput)
	}

	output = new(DeleteCachesOutput)
	err = client.CallMethod("DeleteCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ResizeCaches(input *ResizeCachesInput) (output *ResizeCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ResizeCachesInput)
	}

	output = new(ResizeCachesOutput)
	err = client.CallMethod("ResizeCaches", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) UpdateCache(input *UpdateCacheInput) (output *UpdateCacheOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(UpdateCacheInput)
	}

	output = new(UpdateCacheOutput)
	err = client.CallMethod("UpdateCache", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ChangeCacheVxnet(input *ChangeCacheVxnetInput) (output *ChangeCacheVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ChangeCacheVxnetInput)
	}

	output = new(ChangeCacheVxnetOutput)
	err = client.CallMethod("ChangeCacheVxnet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheAttributes(input *ModifyCacheAttributesInput) (output *ModifyCacheAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyCacheAttributesInput)
	}

	output = new(ModifyCacheAttributesOutput)
	err = client.CallMethod("ModifyCacheAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheNodes(input *DescribeCacheNodesInput) (output *DescribeCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeCacheNodesInput)
	}

	output = new(DescribeCacheNodesOutput)
	err = client.CallMethod("DescribeCacheNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) AddCacheNodes(input *AddCacheNodesInput) (output *AddCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(AddCacheNodesInput)
	}

	output = new(AddCacheNodesOutput)
	err = client.CallMethod("AddCacheNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCacheNodes(input *DeleteCacheNodesInput) (output *DeleteCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteCacheNodesInput)
	}

	output = new(DeleteCacheNodesOutput)
	err = client.CallMethod("DeleteCacheNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) RestartCacheNodes(input *RestartCacheNodesInput) (output *RestartCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RestartCacheNodesInput)
	}

	output = new(RestartCacheNodesOutput)
	err = client.CallMethod("RestartCacheNodes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheNodeAttributes(input *ModifyCacheNodeAttributesInput) (output *ModifyCacheNodeAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyCacheNodeAttributesInput)
	}

	output = new(ModifyCacheNodeAttributesOutput)
	err = client.CallMethod("ModifyCacheNodeAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCacheFromSnapshot(input *CreateCacheFromSnapshotInput) (output *CreateCacheFromSnapshotOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateCacheFromSnapshotInput)
	}

	output = new(CreateCacheFromSnapshotOutput)
	err = client.CallMethod("CreateCacheFromSnapshot", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheParameterGroups(input *DescribeCacheParameterGroupsInput) (output *DescribeCacheParameterGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeCacheParameterGroupsInput)
	}

	output = new(DescribeCacheParameterGroupsOutput)
	err = client.CallMethod("DescribeCacheParameterGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCacheParameterGroup(input *CreateCacheParameterGroupInput) (output *CreateCacheParameterGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateCacheParameterGroupInput)
	}

	output = new(CreateCacheParameterGroupOutput)
	err = client.CallMethod("CreateCacheParameterGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ApplyCacheParameterGroup(input *ApplyCacheParameterGroupInput) (output *ApplyCacheParameterGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ApplyCacheParameterGroupInput)
	}

	output = new(ApplyCacheParameterGroupOutput)
	err = client.CallMethod("ApplyCacheParameterGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCacheParameterGroups(input *DeleteCacheParameterGroupsInput) (output *DeleteCacheParameterGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteCacheParameterGroupsInput)
	}

	output = new(DeleteCacheParameterGroupsOutput)
	err = client.CallMethod("DeleteCacheParameterGroups", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheParameterGroupAttributes(input *ModifyCacheParameterGroupAttributesInput) (output *ModifyCacheParameterGroupAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyCacheParameterGroupAttributesInput)
	}

	output = new(ModifyCacheParameterGroupAttributesOutput)
	err = client.CallMethod("ModifyCacheParameterGroupAttributes", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheParameters(input *DescribeCacheParametersInput) (output *DescribeCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeCacheParametersInput)
	}

	output = new(DescribeCacheParametersOutput)
	err = client.CallMethod("DescribeCacheParameters", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) UpdateCacheParameters(input *UpdateCacheParametersInput) (output *UpdateCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(UpdateCacheParametersInput)
	}

	output = new(UpdateCacheParametersOutput)
	err = client.CallMethod("UpdateCacheParameters", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ResetCacheParameters(input *ResetCacheParametersInput) (output *ResetCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ResetCacheParametersInput)
	}

	output = new(ResetCacheParametersOutput)
	err = client.CallMethod("ResetCacheParameters", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
