// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: hadoop.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type HadoopServiceInterface interface {
	AddHadoopNodes(in *AddHadoopNodesInput) (out *AddHadoopNodesOutput, err error)
	DeleteHadoopNodes(in *DeleteHadoopNodesInput) (out *DeleteHadoopNodesOutput, err error)
	StartHadoops(in *StartHadoopsInput) (out *StartHadoopsOutput, err error)
	StopHadoops(in *StopHadoopsInput) (out *StopHadoopsOutput, err error)
}

type HadoopService struct {
	Properties       *HadoopServiceProperties
	LastResponseBody string
}

func NewHadoopService(accessKeyId, secretAccessKey, zone string) (p *HadoopService) {
	return &HadoopService{
		Properties: &HadoopServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *HadoopService) AddHadoopNodes(input *AddHadoopNodesInput) (output *AddHadoopNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddHadoopNodesOutput)

	err = client.CallMethod(nil, "AddHadoopNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) DeleteHadoopNodes(input *DeleteHadoopNodesInput) (output *DeleteHadoopNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteHadoopNodesOutput)

	err = client.CallMethod(nil, "DeleteHadoopNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) StartHadoops(input *StartHadoopsInput) (output *StartHadoopsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StartHadoopsOutput)

	err = client.CallMethod(nil, "StartHadoops", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *HadoopService) StopHadoops(input *StopHadoopsInput) (output *StopHadoopsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StopHadoopsOutput)

	err = client.CallMethod(nil, "StopHadoops", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
