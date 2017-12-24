// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: rdb.proto

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

type RDBServiceInterface interface {
	CreateRDB(in *CreateRDBInput) (out *CreateRDBOutput, err error)
	DescribeRDBs(in *DescribeRDBsInput) (out *DescribeRDBsOutput, err error)
	DeleteRDBs(in *DeleteRDBsInput) (out *DeleteRDBsOutput, err error)
	StartRDBs(in *StartRDBsInput) (out *StartRDBsOutput, err error)
	StopRDBs(in *StopRDBsInput) (out *StopRDBsOutput, err error)
	ResizeRDBs(in *ResizeRDBsInput) (out *ResizeRDBsOutput, err error)
	RDBsLeaveVxnet(in *RDBsLeaveVxnetInput) (out *RDBsLeaveVxnetOutput, err error)
	RDBsJoinVxnet(in *RDBsJoinVxnetInput) (out *RDBsJoinVxnetOutput, err error)
	CreateRDBFromSnapshot(in *CreateRDBFromSnapshotInput) (out *CreateRDBFromSnapshotOutput, err error)
	CreateTempRDBInstanceFromSnapshot(in *CreateTempRDBInstanceFromSnapshotInput) (out *CreateTempRDBInstanceFromSnapshotOutput, err error)
	GetRDBInstanceFiles(in *GetRDBInstanceFilesInput) (out *GetRDBInstanceFilesOutput, err error)
	CopyRDBInstanceFilesToFTP(in *CopyRDBInstanceFilesToFTPInput) (out *CopyRDBInstanceFilesToFTPOutput, err error)
	PurgeRDBLogs(in *PurgeRDBLogsInput) (out *PurgeRDBLogsOutput, err error)
	CeaseRDBInstance(in *CeaseRDBInstanceInput) (out *CeaseRDBInstanceOutput, err error)
	ModifyRDBParameters(in *ModifyRDBParametersInput) (out *ModifyRDBParametersOutput, err error)
	ApplyRDBParameterGroup(in *ApplyRDBParameterGroupInput) (out *ApplyRDBParameterGroupOutput, err error)
	DescribeRDBParameters(in *DescribeRDBParametersInput) (out *DescribeRDBParametersOutput, err error)
}

type RDBService struct {
	ServerInfo *ServerInfo
}

func NewRDBService(server *ServerInfo) (p *RDBService) {
	return &RDBService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["CreateRDB"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "CreateRDB",
		InputTypeName:  "CreateRDBInput",
		OutputTypeName: "CreateRDBOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*CreateRDBInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateRDBOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeRDBs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "DescribeRDBs",
		InputTypeName:  "DescribeRDBsInput",
		OutputTypeName: "DescribeRDBsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRDBsInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRDBsOutput)(nil)),
	}
	ServiceApiSpecMap["DeleteRDBs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "DeleteRDBs",
		InputTypeName:  "DeleteRDBsInput",
		OutputTypeName: "DeleteRDBsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*DeleteRDBsInput)(nil)),
		OutputType:  reflect.TypeOf((*DeleteRDBsOutput)(nil)),
	}
	ServiceApiSpecMap["StartRDBs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "StartRDBs",
		InputTypeName:  "StartRDBsInput",
		OutputTypeName: "StartRDBsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*StartRDBsInput)(nil)),
		OutputType:  reflect.TypeOf((*StartRDBsOutput)(nil)),
	}
	ServiceApiSpecMap["StopRDBs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "StopRDBs",
		InputTypeName:  "StopRDBsInput",
		OutputTypeName: "StopRDBsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*StopRDBsInput)(nil)),
		OutputType:  reflect.TypeOf((*StopRDBsOutput)(nil)),
	}
	ServiceApiSpecMap["ResizeRDBs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "ResizeRDBs",
		InputTypeName:  "ResizeRDBsInput",
		OutputTypeName: "ResizeRDBsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*ResizeRDBsInput)(nil)),
		OutputType:  reflect.TypeOf((*ResizeRDBsOutput)(nil)),
	}
	ServiceApiSpecMap["RDBsLeaveVxnet"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "RDBsLeaveVxnet",
		InputTypeName:  "RDBsLeaveVxnetInput",
		OutputTypeName: "RDBsLeaveVxnetOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*RDBsLeaveVxnetInput)(nil)),
		OutputType:  reflect.TypeOf((*RDBsLeaveVxnetOutput)(nil)),
	}
	ServiceApiSpecMap["RDBsJoinVxnet"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "RDBsJoinVxnet",
		InputTypeName:  "RDBsJoinVxnetInput",
		OutputTypeName: "RDBsJoinVxnetOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*RDBsJoinVxnetInput)(nil)),
		OutputType:  reflect.TypeOf((*RDBsJoinVxnetOutput)(nil)),
	}
	ServiceApiSpecMap["CreateRDBFromSnapshot"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "CreateRDBFromSnapshot",
		InputTypeName:  "CreateRDBFromSnapshotInput",
		OutputTypeName: "CreateRDBFromSnapshotOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*CreateRDBFromSnapshotInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateRDBFromSnapshotOutput)(nil)),
	}
	ServiceApiSpecMap["CreateTempRDBInstanceFromSnapshot"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "CreateTempRDBInstanceFromSnapshot",
		InputTypeName:  "CreateTempRDBInstanceFromSnapshotInput",
		OutputTypeName: "CreateTempRDBInstanceFromSnapshotOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*CreateTempRDBInstanceFromSnapshotInput)(nil)),
		OutputType:  reflect.TypeOf((*CreateTempRDBInstanceFromSnapshotOutput)(nil)),
	}
	ServiceApiSpecMap["GetRDBInstanceFiles"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "GetRDBInstanceFiles",
		InputTypeName:  "GetRDBInstanceFilesInput",
		OutputTypeName: "GetRDBInstanceFilesOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*GetRDBInstanceFilesInput)(nil)),
		OutputType:  reflect.TypeOf((*GetRDBInstanceFilesOutput)(nil)),
	}
	ServiceApiSpecMap["CopyRDBInstanceFilesToFTP"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "CopyRDBInstanceFilesToFTP",
		InputTypeName:  "CopyRDBInstanceFilesToFTPInput",
		OutputTypeName: "CopyRDBInstanceFilesToFTPOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*CopyRDBInstanceFilesToFTPInput)(nil)),
		OutputType:  reflect.TypeOf((*CopyRDBInstanceFilesToFTPOutput)(nil)),
	}
	ServiceApiSpecMap["PurgeRDBLogs"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "PurgeRDBLogs",
		InputTypeName:  "PurgeRDBLogsInput",
		OutputTypeName: "PurgeRDBLogsOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*PurgeRDBLogsInput)(nil)),
		OutputType:  reflect.TypeOf((*PurgeRDBLogsOutput)(nil)),
	}
	ServiceApiSpecMap["CeaseRDBInstance"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "CeaseRDBInstance",
		InputTypeName:  "CeaseRDBInstanceInput",
		OutputTypeName: "CeaseRDBInstanceOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*CeaseRDBInstanceInput)(nil)),
		OutputType:  reflect.TypeOf((*CeaseRDBInstanceOutput)(nil)),
	}
	ServiceApiSpecMap["ModifyRDBParameters"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "ModifyRDBParameters",
		InputTypeName:  "ModifyRDBParametersInput",
		OutputTypeName: "ModifyRDBParametersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*ModifyRDBParametersInput)(nil)),
		OutputType:  reflect.TypeOf((*ModifyRDBParametersOutput)(nil)),
	}
	ServiceApiSpecMap["ApplyRDBParameterGroup"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "ApplyRDBParameterGroup",
		InputTypeName:  "ApplyRDBParameterGroupInput",
		OutputTypeName: "ApplyRDBParameterGroupOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*ApplyRDBParameterGroupInput)(nil)),
		OutputType:  reflect.TypeOf((*ApplyRDBParameterGroupOutput)(nil)),
	}
	ServiceApiSpecMap["DescribeRDBParameters"] = ServiceApiSpec{
		ServiceName:    "RDBService",
		ActionName:     "DescribeRDBParameters",
		InputTypeName:  "DescribeRDBParametersInput",
		OutputTypeName: "DescribeRDBParametersOutput",
		HttpMethod:     "GET",

		ServiceType: reflect.TypeOf((*RDBService)(nil)),
		InputType:   reflect.TypeOf((*DescribeRDBParametersInput)(nil)),
		OutputType:  reflect.TypeOf((*DescribeRDBParametersOutput)(nil)),
	}
}

func (p *RDBService) CreateRDB(input *CreateRDBInput) (output *CreateRDBOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateRDBInput)
	}

	output = new(CreateRDBOutput)
	err = client.CallMethod("CreateRDB", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) DescribeRDBs(input *DescribeRDBsInput) (output *DescribeRDBsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRDBsInput)
	}

	output = new(DescribeRDBsOutput)
	err = client.CallMethod("DescribeRDBs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) DeleteRDBs(input *DeleteRDBsInput) (output *DeleteRDBsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DeleteRDBsInput)
	}

	output = new(DeleteRDBsOutput)
	err = client.CallMethod("DeleteRDBs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) StartRDBs(input *StartRDBsInput) (output *StartRDBsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StartRDBsInput)
	}

	output = new(StartRDBsOutput)
	err = client.CallMethod("StartRDBs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) StopRDBs(input *StopRDBsInput) (output *StopRDBsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(StopRDBsInput)
	}

	output = new(StopRDBsOutput)
	err = client.CallMethod("StopRDBs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) ResizeRDBs(input *ResizeRDBsInput) (output *ResizeRDBsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ResizeRDBsInput)
	}

	output = new(ResizeRDBsOutput)
	err = client.CallMethod("ResizeRDBs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) RDBsLeaveVxnet(input *RDBsLeaveVxnetInput) (output *RDBsLeaveVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RDBsLeaveVxnetInput)
	}

	output = new(RDBsLeaveVxnetOutput)
	err = client.CallMethod("RDBsLeaveVxnet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) RDBsJoinVxnet(input *RDBsJoinVxnetInput) (output *RDBsJoinVxnetOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(RDBsJoinVxnetInput)
	}

	output = new(RDBsJoinVxnetOutput)
	err = client.CallMethod("RDBsJoinVxnet", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) CreateRDBFromSnapshot(input *CreateRDBFromSnapshotInput) (output *CreateRDBFromSnapshotOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateRDBFromSnapshotInput)
	}

	output = new(CreateRDBFromSnapshotOutput)
	err = client.CallMethod("CreateRDBFromSnapshot", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) CreateTempRDBInstanceFromSnapshot(input *CreateTempRDBInstanceFromSnapshotInput) (output *CreateTempRDBInstanceFromSnapshotOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CreateTempRDBInstanceFromSnapshotInput)
	}

	output = new(CreateTempRDBInstanceFromSnapshotOutput)
	err = client.CallMethod("CreateTempRDBInstanceFromSnapshot", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) GetRDBInstanceFiles(input *GetRDBInstanceFilesInput) (output *GetRDBInstanceFilesOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(GetRDBInstanceFilesInput)
	}

	output = new(GetRDBInstanceFilesOutput)
	err = client.CallMethod("GetRDBInstanceFiles", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) CopyRDBInstanceFilesToFTP(input *CopyRDBInstanceFilesToFTPInput) (output *CopyRDBInstanceFilesToFTPOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CopyRDBInstanceFilesToFTPInput)
	}

	output = new(CopyRDBInstanceFilesToFTPOutput)
	err = client.CallMethod("CopyRDBInstanceFilesToFTP", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) PurgeRDBLogs(input *PurgeRDBLogsInput) (output *PurgeRDBLogsOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(PurgeRDBLogsInput)
	}

	output = new(PurgeRDBLogsOutput)
	err = client.CallMethod("PurgeRDBLogs", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) CeaseRDBInstance(input *CeaseRDBInstanceInput) (output *CeaseRDBInstanceOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(CeaseRDBInstanceInput)
	}

	output = new(CeaseRDBInstanceOutput)
	err = client.CallMethod("CeaseRDBInstance", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) ModifyRDBParameters(input *ModifyRDBParametersInput) (output *ModifyRDBParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ModifyRDBParametersInput)
	}

	output = new(ModifyRDBParametersOutput)
	err = client.CallMethod("ModifyRDBParameters", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) ApplyRDBParameterGroup(input *ApplyRDBParameterGroupInput) (output *ApplyRDBParameterGroupOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(ApplyRDBParameterGroupInput)
	}

	output = new(ApplyRDBParameterGroupOutput)
	err = client.CallMethod("ApplyRDBParameterGroup", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *RDBService) DescribeRDBParameters(input *DescribeRDBParametersInput) (output *DescribeRDBParametersOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(DescribeRDBParametersInput)
	}

	output = new(DescribeRDBParametersOutput)
	err = client.CallMethod("DescribeRDBParameters", "GET", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
