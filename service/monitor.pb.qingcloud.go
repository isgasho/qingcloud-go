// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: monitor.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/logger"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

// See https://docs.qingcloud.com/api/monitor/index.html
type MonitorServiceInterface interface {
	GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error)
	GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error)
	GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error)
	GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error)
	GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error)
	GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error)
}

// See https://docs.qingcloud.com/api/monitor/index.html
type MonitorService struct {
	Config           *config.Config
	Properties       *MonitorServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/monitor/index.html
func NewMonitorService(conf *config.Config, zone string) (p *MonitorService) {
	return &MonitorService{
		Config:     conf,
		Properties: &MonitorServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/monitor/index.html
func (s *QingCloudService) Monitor(zone string) (*MonitorService, error) {
	properties := &MonitorServiceProperties{
		Zone: zone,
	}

	return &MonitorService{Config: s.Config, Properties: properties}, nil
}

func (p *MonitorService) GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error) {
	if in == nil {
		in = &GetMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetMonitorOutput{}
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

func (p *MonitorService) GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error) {
	if in == nil {
		in = &GetLoadBalancerMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetLoadBalancerMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetLoadBalancerMonitorOutput{}
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

func (p *MonitorService) GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error) {
	if in == nil {
		in = &GetRDBMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetRDBMonitorOutput{}
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

func (p *MonitorService) GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error) {
	if in == nil {
		in = &GetCacheMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetCacheMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetCacheMonitorOutput{}
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

func (p *MonitorService) GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error) {
	if in == nil {
		in = &GetZooKeeperMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetZooKeeperMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetZooKeeperMonitorOutput{}
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

func (p *MonitorService) GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error) {
	if in == nil {
		in = &GetQueueMonitorInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetQueueMonitor",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetQueueMonitorOutput{}
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

func (p *MonitorServiceProperties) Validate() error {
	return nil
}

func (p *GetMonitorInput) Validate() error {
	return nil
}

func (p *GetMonitorOutput) Validate() error {
	return nil
}

func (p *GetLoadBalancerMonitorInput) Validate() error {
	return nil
}

func (p *GetLoadBalancerMonitorOutput) Validate() error {
	return nil
}

func (p *GetRDBMonitorInput) Validate() error {
	return nil
}

func (p *GetRDBMonitorOutput) Validate() error {
	return nil
}

func (p *GetCacheMonitorInput) Validate() error {
	return nil
}

func (p *GetCacheMonitorOutput) Validate() error {
	return nil
}

func (p *GetZooKeeperMonitorInput) Validate() error {
	return nil
}

func (p *GetZooKeeperMonitorOutput) Validate() error {
	return nil
}

func (p *GetQueueMonitorInput) Validate() error {
	return nil
}

func (p *GetQueueMonitorOutput) Validate() error {
	return nil
}
