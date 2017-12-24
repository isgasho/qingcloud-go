// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: monitor.proto

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
	TypeInfoMap["GetMonitorInput"] = reflect.TypeOf((*GetMonitorInput)(nil))
	TypeInfoMap["GetMonitorOutput"] = reflect.TypeOf((*GetMonitorOutput)(nil))
	TypeInfoMap["GetLoadBalancerMonitorInput"] = reflect.TypeOf((*GetLoadBalancerMonitorInput)(nil))
	TypeInfoMap["GetLoadBalancerMonitorOutput"] = reflect.TypeOf((*GetLoadBalancerMonitorOutput)(nil))
	TypeInfoMap["GetRDBMonitorInput"] = reflect.TypeOf((*GetRDBMonitorInput)(nil))
	TypeInfoMap["GetRDBMonitorOutput"] = reflect.TypeOf((*GetRDBMonitorOutput)(nil))
	TypeInfoMap["GetCacheMonitorInput"] = reflect.TypeOf((*GetCacheMonitorInput)(nil))
	TypeInfoMap["GetCacheMonitorOutput"] = reflect.TypeOf((*GetCacheMonitorOutput)(nil))
	TypeInfoMap["GetZooKeeperMonitorInput"] = reflect.TypeOf((*GetZooKeeperMonitorInput)(nil))
	TypeInfoMap["GetZooKeeperMonitorOutput"] = reflect.TypeOf((*GetZooKeeperMonitorOutput)(nil))
	TypeInfoMap["GetQueueMonitorInput"] = reflect.TypeOf((*GetQueueMonitorInput)(nil))
	TypeInfoMap["GetQueueMonitorOutput"] = reflect.TypeOf((*GetQueueMonitorOutput)(nil))
}

type MonitorServiceInterface interface {
	GetMonitor(in *GetMonitorInput) (out *GetMonitorOutput, err error)
	GetLoadBalancerMonitor(in *GetLoadBalancerMonitorInput) (out *GetLoadBalancerMonitorOutput, err error)
	GetRDBMonitor(in *GetRDBMonitorInput) (out *GetRDBMonitorOutput, err error)
	GetCacheMonitor(in *GetCacheMonitorInput) (out *GetCacheMonitorOutput, err error)
	GetZooKeeperMonitor(in *GetZooKeeperMonitorInput) (out *GetZooKeeperMonitorOutput, err error)
	GetQueueMonitor(in *GetQueueMonitorInput) (out *GetQueueMonitorOutput, err error)
}

type MonitorService struct {
	ServerInfo *ServerInfo
}

func NewMonitorService(server *ServerInfo) (p *MonitorService) {
	return &MonitorService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["GetMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetMonitor",
		InputTypeName:  "GetMonitorInput",
		OutputTypeName: "GetMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetMonitorOutput)(nil)),
	}
	ServiceApiSpecMap["GetLoadBalancerMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetLoadBalancerMonitor",
		InputTypeName:  "GetLoadBalancerMonitorInput",
		OutputTypeName: "GetLoadBalancerMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetLoadBalancerMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetLoadBalancerMonitorOutput)(nil)),
	}
	ServiceApiSpecMap["GetRDBMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetRDBMonitor",
		InputTypeName:  "GetRDBMonitorInput",
		OutputTypeName: "GetRDBMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetRDBMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetRDBMonitorOutput)(nil)),
	}
	ServiceApiSpecMap["GetCacheMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetCacheMonitor",
		InputTypeName:  "GetCacheMonitorInput",
		OutputTypeName: "GetCacheMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetCacheMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetCacheMonitorOutput)(nil)),
	}
	ServiceApiSpecMap["GetZooKeeperMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetZooKeeperMonitor",
		InputTypeName:  "GetZooKeeperMonitorInput",
		OutputTypeName: "GetZooKeeperMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetZooKeeperMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetZooKeeperMonitorOutput)(nil)),
	}
	ServiceApiSpecMap["GetQueueMonitor"] = ServiceApiSpec{
		ServiceName:    "MonitorService",
		ActionName:     "GetQueueMonitor",
		InputTypeName:  "GetQueueMonitorInput",
		OutputTypeName: "GetQueueMonitorOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*MonitorService)(nil)),
		InputType:   reflect.TypeOf((*GetQueueMonitorInput)(nil)),
		OutputType:  reflect.TypeOf((*GetQueueMonitorOutput)(nil)),
	}
}

func (p *MonitorService) GetMonitor(input *GetMonitorInput) (output *GetMonitorOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GetMonitorInput)
	}

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

	if input == nil {
		input = new(GetLoadBalancerMonitorInput)
	}

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

	if input == nil {
		input = new(GetRDBMonitorInput)
	}

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

	if input == nil {
		input = new(GetCacheMonitorInput)
	}

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

	if input == nil {
		input = new(GetZooKeeperMonitorInput)
	}

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

	if input == nil {
		input = new(GetQueueMonitorInput)
	}

	output = new(GetQueueMonitorOutput)
	err = client.CallMethod("GetQueueMonitor", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
