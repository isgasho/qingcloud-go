// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: monitor.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type MonitorServiceInterface interface {
	GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error)
	GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error)
	GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error)
	GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error)
	GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error)
	GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error)
}

type MonitorService struct {
	ServerInfo       *ServerInfo
	LastResponseBody string
}

func NewMonitorService(server *ServerInfo) (p *MonitorService) {
	return &MonitorService{
		ServerInfo: server,
	}
}

func (p *MonitorService) GetMonitor(input *GetMonitorInput) (output *GetMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetMonitorOutput)

	err = client.CallMethod("GetMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MonitorService) GetLoadBalancerMonitor(input *GetLoadBalancerMonitorInput) (output *GetLoadBalancerMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetLoadBalancerMonitorOutput)

	err = client.CallMethod("GetLoadBalancerMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MonitorService) GetRDBMonitor(input *GetRDBMonitorInput) (output *GetRDBMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetRDBMonitorOutput)

	err = client.CallMethod("GetRDBMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MonitorService) GetCacheMonitor(input *GetCacheMonitorInput) (output *GetCacheMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetCacheMonitorOutput)

	err = client.CallMethod("GetCacheMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MonitorService) GetZooKeeperMonitor(input *GetZooKeeperMonitorInput) (output *GetZooKeeperMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetZooKeeperMonitorOutput)

	err = client.CallMethod("GetZooKeeperMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *MonitorService) GetQueueMonitor(input *GetQueueMonitorInput) (output *GetQueueMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)
	output = new(GetQueueMonitorOutput)

	err = client.CallMethod("GetQueueMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
