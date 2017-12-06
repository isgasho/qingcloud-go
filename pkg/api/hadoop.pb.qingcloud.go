// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: hadoop.proto

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

type HadoopServiceInterface interface {
	AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error)
	DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error)
	StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error)
	StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error)
}

type HadoopService struct {
	Config           *config.Config
	Properties       *HadoopServiceProperties
	LastResponseBody string
}

func NewHadoopService(conf *config.Config, zone string) (p *HadoopService) {
	return &HadoopService{
		Config:     conf,
		Properties: &HadoopServiceProperties{Zone: proto.String(zone)},
	}
}

func (s *QingCloudService) Hadoop(zone string) (*HadoopService, error) {
	properties := &HadoopServiceProperties{
		Zone: proto.String(zone),
	}

	return &HadoopService{
		Config:     s.Config,
		Properties: properties,
	}, nil
}

func (p *HadoopService) AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error) {
	if in == nil {
		in = &AddHadoopNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddHadoopNodes",
		RequestMethod: "GET",
	}

	x := &AddHadoopNodesOutput{}
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

func (p *HadoopService) DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error) {
	if in == nil {
		in = &DeleteHadoopNodesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteHadoopNodes",
		RequestMethod: "GET",
	}

	x := &DeleteHadoopNodesOutput{}
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

func (p *HadoopService) StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error) {
	if in == nil {
		in = &StartHadoopsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartHadoops",
		RequestMethod: "GET",
	}

	x := &StartHadoopsOutput{}
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

func (p *HadoopService) StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error) {
	if in == nil {
		in = &StopHadoopsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopHadoops",
		RequestMethod: "GET",
	}

	x := &StopHadoopsOutput{}
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
