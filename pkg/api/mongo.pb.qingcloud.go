// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: mongo.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type MongoServiceInterface interface {
	DescribeMongoNodes(in *DescribeMongoNodesInput) (out *DescribeMongoNodesOutput, err error)
	DescribeMongoParameters(in *DescribeMongoParametersInput) (out *DescribeMongoParametersOutput, err error)
	ResizeMongos(in *ResizeMongosInput) (out *ResizeMongosOutput, err error)
	CreateMongo(in *CreateMongoInput) (out *CreateMongoOutput, err error)
	StopMongos(in *StopMongosInput) (out *StopMongosOutput, err error)
	StartMongos(in *StartMongosInput) (out *StartMongosOutput, err error)
	DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error)
	DeleteMongos(in *DeleteMongosInput) (out *DeleteMongosOutput, err error)
	CreateMongoFromSnapshot(in *CreateMongoFromSnapshotInput) (out *CreateMongoFromSnapshotOutput, err error)
	ChangeMongoVxnet(in *ChangeMongoVxnetInput) (out *ChangeMongoVxnetOutput, err error)
	AddMongoInstances(in *AddMongoInstancesInput) (out *AddMongoInstancesOutput, err error)
	RemoveMongoInstances(in *RemoveMongoInstancesInput) (out *RemoveMongoInstancesOutput, err error)
	ModifyMongoAttributes(in *ModifyMongoAttributesInput) (out *ModifyMongoAttributesOutput, err error)
	ModifyMongoInstances(in *ModifyMongoInstancesInput) (out *ModifyMongoInstancesOutput, err error)
}

type MongoService struct {
	ServerInfo       *ServerInfo
	Properties       *MongoServiceProperties
	LastResponseBody string
}

func NewMongoService(server *ServerInfo, serviceProp *MongoServiceProperties) (p *MongoService) {
	return &MongoService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *MongoService) DescribeMongoNodes(input *DescribeMongoNodesInput) (output *DescribeMongoNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeMongoNodesOutput)

	err = client.CallMethod(nil, "DescribeMongoNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) DescribeMongoParameters(input *DescribeMongoParametersInput) (output *DescribeMongoParametersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeMongoParametersOutput)

	err = client.CallMethod(nil, "DescribeMongoParameters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) ResizeMongos(input *ResizeMongosInput) (output *ResizeMongosOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ResizeMongosOutput)

	err = client.CallMethod(nil, "ResizeMongos", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) CreateMongo(input *CreateMongoInput) (output *CreateMongoOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateMongoOutput)

	err = client.CallMethod(nil, "CreateMongo", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) StopMongos(input *StopMongosInput) (output *StopMongosOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StopMongosOutput)

	err = client.CallMethod(nil, "StopMongos", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) StartMongos(input *StartMongosInput) (output *StartMongosOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StartMongosOutput)

	err = client.CallMethod(nil, "StartMongos", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) DescribeMongos(input *DescribeMongosInput) (output *DescribeMongosOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeMongosOutput)

	err = client.CallMethod(nil, "DescribeMongos", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) DeleteMongos(input *DeleteMongosInput) (output *DeleteMongosOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteMongosOutput)

	err = client.CallMethod(nil, "DeleteMongos", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) CreateMongoFromSnapshot(input *CreateMongoFromSnapshotInput) (output *CreateMongoFromSnapshotOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateMongoFromSnapshotOutput)

	err = client.CallMethod(nil, "CreateMongoFromSnapshot", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) ChangeMongoVxnet(input *ChangeMongoVxnetInput) (output *ChangeMongoVxnetOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ChangeMongoVxnetOutput)

	err = client.CallMethod(nil, "ChangeMongoVxnet", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) AddMongoInstances(input *AddMongoInstancesInput) (output *AddMongoInstancesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddMongoInstancesOutput)

	err = client.CallMethod(nil, "AddMongoInstances", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) RemoveMongoInstances(input *RemoveMongoInstancesInput) (output *RemoveMongoInstancesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(RemoveMongoInstancesOutput)

	err = client.CallMethod(nil, "RemoveMongoInstances", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) ModifyMongoAttributes(input *ModifyMongoAttributesInput) (output *ModifyMongoAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyMongoAttributesOutput)

	err = client.CallMethod(nil, "ModifyMongoAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MongoService) ModifyMongoInstances(input *ModifyMongoInstancesInput) (output *ModifyMongoInstancesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyMongoInstancesOutput)

	err = client.CallMethod(nil, "ModifyMongoInstances", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
