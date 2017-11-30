// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: mongo.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"
import "github.com/chai2010/qingcloud-go/pkg/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}
var _ = data.Operation{}

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
	Config           *config.Config
	Properties       *MongoServiceProperties
	LastResponseBody string
}

func NewMongoService(conf *config.Config, zone string) (p *MongoService) {
	return &MongoService{
		Config:     conf,
		Properties: &MongoServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Mongo(zone string) (*MongoService, error) {
	properties := &MongoServiceProperties{
		Zone: proto.String(zone),
	}

	return &MongoService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *MongoService) DescribeMongoNodes(in *DescribeMongoNodesInput) (out *DescribeMongoNodesOutput, err error) {
	if in == nil {
		in = &DescribeMongoNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoNodes",
		RequestMethod: "GET",
	}

	x := &DescribeMongoNodesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) DescribeMongoParameters(in *DescribeMongoParametersInput) (out *DescribeMongoParametersOutput, err error) {
	if in == nil {
		in = &DescribeMongoParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongoParameters",
		RequestMethod: "GET",
	}

	x := &DescribeMongoParametersOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) ResizeMongos(in *ResizeMongosInput) (out *ResizeMongosOutput, err error) {
	if in == nil {
		in = &ResizeMongosInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeMongos",
		RequestMethod: "GET",
	}

	x := &ResizeMongosOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) CreateMongo(in *CreateMongoInput) (out *CreateMongoOutput, err error) {
	if in == nil {
		in = &CreateMongoInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongo",
		RequestMethod: "GET",
	}

	x := &CreateMongoOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) StopMongos(in *StopMongosInput) (out *StopMongosOutput, err error) {
	if in == nil {
		in = &StopMongosInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopMongos",
		RequestMethod: "GET",
	}

	x := &StopMongosOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) StartMongos(in *StartMongosInput) (out *StartMongosOutput, err error) {
	if in == nil {
		in = &StartMongosInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartMongos",
		RequestMethod: "GET",
	}

	x := &StartMongosOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) DescribeMongos(in *DescribeMongosInput) (out *DescribeMongosOutput, err error) {
	if in == nil {
		in = &DescribeMongosInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeMongos",
		RequestMethod: "GET",
	}

	x := &DescribeMongosOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) DeleteMongos(in *DeleteMongosInput) (out *DeleteMongosOutput, err error) {
	if in == nil {
		in = &DeleteMongosInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteMongos",
		RequestMethod: "GET",
	}

	x := &DeleteMongosOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) CreateMongoFromSnapshot(in *CreateMongoFromSnapshotInput) (out *CreateMongoFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateMongoFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateMongoFromSnapshot",
		RequestMethod: "GET",
	}

	x := &CreateMongoFromSnapshotOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) ChangeMongoVxnet(in *ChangeMongoVxnetInput) (out *ChangeMongoVxnetOutput, err error) {
	if in == nil {
		in = &ChangeMongoVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ChangeMongoVxnet",
		RequestMethod: "GET",
	}

	x := &ChangeMongoVxnetOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) AddMongoInstances(in *AddMongoInstancesInput) (out *AddMongoInstancesOutput, err error) {
	if in == nil {
		in = &AddMongoInstancesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddMongoInstances",
		RequestMethod: "GET",
	}

	x := &AddMongoInstancesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) RemoveMongoInstances(in *RemoveMongoInstancesInput) (out *RemoveMongoInstancesOutput, err error) {
	if in == nil {
		in = &RemoveMongoInstancesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RemoveMongoInstances",
		RequestMethod: "GET",
	}

	x := &RemoveMongoInstancesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) ModifyMongoAttributes(in *ModifyMongoAttributesInput) (out *ModifyMongoAttributesOutput, err error) {
	if in == nil {
		in = &ModifyMongoAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoAttributes",
		RequestMethod: "GET",
	}

	x := &ModifyMongoAttributesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}

func (p *MongoService) ModifyMongoInstances(in *ModifyMongoInstancesInput) (out *ModifyMongoInstancesOutput, err error) {
	if in == nil {
		in = &ModifyMongoInstancesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyMongoInstances",
		RequestMethod: "GET",
	}

	x := &ModifyMongoInstancesOutput{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}
