// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: cache.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

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
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewCacheService(server *ServerInfo) (p *CacheService) {
	return &CacheService{
		ServerInfo: server,
	}
}

func (p *CacheService) DescribeCaches(input *DescribeCachesInput) (output *DescribeCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeCachesOutput)

	err = client.CallMethod(nil, "DescribeCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCache(input *CreateCacheInput) (output *CreateCacheOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(CreateCacheOutput)

	err = client.CallMethod(nil, "CreateCache", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) StopCaches(input *StopCachesInput) (output *StopCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(StopCachesOutput)

	err = client.CallMethod(nil, "StopCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) StartCaches(input *StartCachesInput) (output *StartCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(StartCachesOutput)

	err = client.CallMethod(nil, "StartCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) RestartCaches(input *RestartCachesInput) (output *RestartCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(RestartCachesOutput)

	err = client.CallMethod(nil, "RestartCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCaches(input *DeleteCachesInput) (output *DeleteCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DeleteCachesOutput)

	err = client.CallMethod(nil, "DeleteCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ResizeCaches(input *ResizeCachesInput) (output *ResizeCachesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ResizeCachesOutput)

	err = client.CallMethod(nil, "ResizeCaches", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) UpdateCache(input *UpdateCacheInput) (output *UpdateCacheOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(UpdateCacheOutput)

	err = client.CallMethod(nil, "UpdateCache", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ChangeCacheVxnet(input *ChangeCacheVxnetInput) (output *ChangeCacheVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ChangeCacheVxnetOutput)

	err = client.CallMethod(nil, "ChangeCacheVxnet", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheAttributes(input *ModifyCacheAttributesInput) (output *ModifyCacheAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ModifyCacheAttributesOutput)

	err = client.CallMethod(nil, "ModifyCacheAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheNodes(input *DescribeCacheNodesInput) (output *DescribeCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeCacheNodesOutput)

	err = client.CallMethod(nil, "DescribeCacheNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) AddCacheNodes(input *AddCacheNodesInput) (output *AddCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(AddCacheNodesOutput)

	err = client.CallMethod(nil, "AddCacheNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCacheNodes(input *DeleteCacheNodesInput) (output *DeleteCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DeleteCacheNodesOutput)

	err = client.CallMethod(nil, "DeleteCacheNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) RestartCacheNodes(input *RestartCacheNodesInput) (output *RestartCacheNodesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(RestartCacheNodesOutput)

	err = client.CallMethod(nil, "RestartCacheNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheNodeAttributes(input *ModifyCacheNodeAttributesInput) (output *ModifyCacheNodeAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ModifyCacheNodeAttributesOutput)

	err = client.CallMethod(nil, "ModifyCacheNodeAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCacheFromSnapshot(input *CreateCacheFromSnapshotInput) (output *CreateCacheFromSnapshotOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(CreateCacheFromSnapshotOutput)

	err = client.CallMethod(nil, "CreateCacheFromSnapshot", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheParameterGroups(input *DescribeCacheParameterGroupsInput) (output *DescribeCacheParameterGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeCacheParameterGroupsOutput)

	err = client.CallMethod(nil, "DescribeCacheParameterGroups", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) CreateCacheParameterGroup(input *CreateCacheParameterGroupInput) (output *CreateCacheParameterGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(CreateCacheParameterGroupOutput)

	err = client.CallMethod(nil, "CreateCacheParameterGroup", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ApplyCacheParameterGroup(input *ApplyCacheParameterGroupInput) (output *ApplyCacheParameterGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ApplyCacheParameterGroupOutput)

	err = client.CallMethod(nil, "ApplyCacheParameterGroup", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DeleteCacheParameterGroups(input *DeleteCacheParameterGroupsInput) (output *DeleteCacheParameterGroupsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DeleteCacheParameterGroupsOutput)

	err = client.CallMethod(nil, "DeleteCacheParameterGroups", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ModifyCacheParameterGroupAttributes(input *ModifyCacheParameterGroupAttributesInput) (output *ModifyCacheParameterGroupAttributesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ModifyCacheParameterGroupAttributesOutput)

	err = client.CallMethod(nil, "ModifyCacheParameterGroupAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) DescribeCacheParameters(input *DescribeCacheParametersInput) (output *DescribeCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(DescribeCacheParametersOutput)

	err = client.CallMethod(nil, "DescribeCacheParameters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) UpdateCacheParameters(input *UpdateCacheParametersInput) (output *UpdateCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(UpdateCacheParametersOutput)

	err = client.CallMethod(nil, "UpdateCacheParameters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *CacheService) ResetCacheParameters(input *ResetCacheParametersInput) (output *ResetCacheParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		nil,
	)
	output = new(ResetCacheParametersOutput)

	err = client.CallMethod(nil, "ResetCacheParameters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
