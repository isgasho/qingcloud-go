// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: rdb.proto

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

// See https://docs.qingcloud.com/api/rdb/index.html
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

// See https://docs.qingcloud.com/api/rdb/index.html
type RDBService struct {
	Config           *config.Config
	Properties       *RDBServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/rdb/index.html
func NewRDBService(conf *config.Config, zone string) (p *RDBService) {
	return &RDBService{
		Config:     conf,
		Properties: &RDBServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/rdb/index.html
func (s *QingCloudService) RDB(zone string) (*RDBService, error) {
	properties := &RDBServiceProperties{
		Zone: zone,
	}

	return &RDBService{Config: s.Config, Properties: properties}, nil
}

func (p *RDBService) CreateRDB(in *CreateRDBInput) (out *CreateRDBOutput, err error) {
	if in == nil {
		in = &CreateRDBInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDB",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateRDBOutput{}
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

func (p *RDBService) DescribeRDBs(in *DescribeRDBsInput) (out *DescribeRDBsOutput, err error) {
	if in == nil {
		in = &DescribeRDBsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeRDBsOutput{}
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

func (p *RDBService) DeleteRDBs(in *DeleteRDBsInput) (out *DeleteRDBsOutput, err error) {
	if in == nil {
		in = &DeleteRDBsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteRDBs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteRDBsOutput{}
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

func (p *RDBService) StartRDBs(in *StartRDBsInput) (out *StartRDBsOutput, err error) {
	if in == nil {
		in = &StartRDBsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StartRDBs",
		RequestMethod: "GET", // GET or POST
	}

	x := &StartRDBsOutput{}
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

func (p *RDBService) StopRDBs(in *StopRDBsInput) (out *StopRDBsOutput, err error) {
	if in == nil {
		in = &StopRDBsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "StopRDBs",
		RequestMethod: "GET", // GET or POST
	}

	x := &StopRDBsOutput{}
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

func (p *RDBService) ResizeRDBs(in *ResizeRDBsInput) (out *ResizeRDBsOutput, err error) {
	if in == nil {
		in = &ResizeRDBsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ResizeRDBs",
		RequestMethod: "GET", // GET or POST
	}

	x := &ResizeRDBsOutput{}
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

func (p *RDBService) RDBsLeaveVxnet(in *RDBsLeaveVxnetInput) (out *RDBsLeaveVxnetOutput, err error) {
	if in == nil {
		in = &RDBsLeaveVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsLeaveVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &RDBsLeaveVxnetOutput{}
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

func (p *RDBService) RDBsJoinVxnet(in *RDBsJoinVxnetInput) (out *RDBsJoinVxnetOutput, err error) {
	if in == nil {
		in = &RDBsJoinVxnetInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "RDBsJoinVxnet",
		RequestMethod: "GET", // GET or POST
	}

	x := &RDBsJoinVxnetOutput{}
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

func (p *RDBService) CreateRDBFromSnapshot(in *CreateRDBFromSnapshotInput) (out *CreateRDBFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateRDBFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateRDBFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateRDBFromSnapshotOutput{}
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

func (p *RDBService) CreateTempRDBInstanceFromSnapshot(in *CreateTempRDBInstanceFromSnapshotInput) (out *CreateTempRDBInstanceFromSnapshotOutput, err error) {
	if in == nil {
		in = &CreateTempRDBInstanceFromSnapshotInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateTempRDBInstanceFromSnapshot",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateTempRDBInstanceFromSnapshotOutput{}
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

func (p *RDBService) GetRDBInstanceFiles(in *GetRDBInstanceFilesInput) (out *GetRDBInstanceFilesOutput, err error) {
	if in == nil {
		in = &GetRDBInstanceFilesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "GetRDBInstanceFiles",
		RequestMethod: "GET", // GET or POST
	}

	x := &GetRDBInstanceFilesOutput{}
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

func (p *RDBService) CopyRDBInstanceFilesToFTP(in *CopyRDBInstanceFilesToFTPInput) (out *CopyRDBInstanceFilesToFTPOutput, err error) {
	if in == nil {
		in = &CopyRDBInstanceFilesToFTPInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CopyRDBInstanceFilesToFTP",
		RequestMethod: "GET", // GET or POST
	}

	x := &CopyRDBInstanceFilesToFTPOutput{}
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

func (p *RDBService) PurgeRDBLogs(in *PurgeRDBLogsInput) (out *PurgeRDBLogsOutput, err error) {
	if in == nil {
		in = &PurgeRDBLogsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "PurgeRDBLogs",
		RequestMethod: "GET", // GET or POST
	}

	x := &PurgeRDBLogsOutput{}
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

func (p *RDBService) CeaseRDBInstance(in *CeaseRDBInstanceInput) (out *CeaseRDBInstanceOutput, err error) {
	if in == nil {
		in = &CeaseRDBInstanceInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CeaseRDBInstance",
		RequestMethod: "GET", // GET or POST
	}

	x := &CeaseRDBInstanceOutput{}
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

func (p *RDBService) ModifyRDBParameters(in *ModifyRDBParametersInput) (out *ModifyRDBParametersOutput, err error) {
	if in == nil {
		in = &ModifyRDBParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyRDBParameters",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyRDBParametersOutput{}
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

func (p *RDBService) ApplyRDBParameterGroup(in *ApplyRDBParameterGroupInput) (out *ApplyRDBParameterGroupOutput, err error) {
	if in == nil {
		in = &ApplyRDBParameterGroupInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyRDBParameterGroup",
		RequestMethod: "GET", // GET or POST
	}

	x := &ApplyRDBParameterGroupOutput{}
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

func (p *RDBService) DescribeRDBParameters(in *DescribeRDBParametersInput) (out *DescribeRDBParametersOutput, err error) {
	if in == nil {
		in = &DescribeRDBParametersInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeRDBParameters",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeRDBParametersOutput{}
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

func (p *RDBServiceProperties) Validate() error {
	return nil
}

func (p *CreateRDBInput) Validate() error {
	return nil
}

func (p *CreateRDBOutput) Validate() error {
	return nil
}

func (p *DescribeRDBsInput) Validate() error {
	return nil
}

func (p *DescribeRDBsOutput) Validate() error {
	return nil
}

func (p *DeleteRDBsInput) Validate() error {
	return nil
}

func (p *DeleteRDBsOutput) Validate() error {
	return nil
}

func (p *StartRDBsInput) Validate() error {
	return nil
}

func (p *StartRDBsOutput) Validate() error {
	return nil
}

func (p *StopRDBsInput) Validate() error {
	return nil
}

func (p *StopRDBsOutput) Validate() error {
	return nil
}

func (p *ResizeRDBsInput) Validate() error {
	return nil
}

func (p *ResizeRDBsOutput) Validate() error {
	return nil
}

func (p *RDBsLeaveVxnetInput) Validate() error {
	return nil
}

func (p *RDBsLeaveVxnetOutput) Validate() error {
	return nil
}

func (p *RDBsJoinVxnetInput) Validate() error {
	return nil
}

func (p *RDBsJoinVxnetOutput) Validate() error {
	return nil
}

func (p *CreateRDBFromSnapshotInput) Validate() error {
	return nil
}

func (p *CreateRDBFromSnapshotOutput) Validate() error {
	return nil
}

func (p *CreateTempRDBInstanceFromSnapshotInput) Validate() error {
	return nil
}

func (p *CreateTempRDBInstanceFromSnapshotOutput) Validate() error {
	return nil
}

func (p *GetRDBInstanceFilesInput) Validate() error {
	return nil
}

func (p *GetRDBInstanceFilesOutput) Validate() error {
	return nil
}

func (p *CopyRDBInstanceFilesToFTPInput) Validate() error {
	return nil
}

func (p *CopyRDBInstanceFilesToFTPOutput) Validate() error {
	return nil
}

func (p *PurgeRDBLogsInput) Validate() error {
	return nil
}

func (p *PurgeRDBLogsOutput) Validate() error {
	return nil
}

func (p *CeaseRDBInstanceInput) Validate() error {
	return nil
}

func (p *CeaseRDBInstanceOutput) Validate() error {
	return nil
}

func (p *ModifyRDBParametersInput) Validate() error {
	return nil
}

func (p *ModifyRDBParametersOutput) Validate() error {
	return nil
}

func (p *ApplyRDBParameterGroupInput) Validate() error {
	return nil
}

func (p *ApplyRDBParameterGroupOutput) Validate() error {
	return nil
}

func (p *DescribeRDBParametersInput) Validate() error {
	return nil
}

func (p *DescribeRDBParametersOutput) Validate() error {
	return nil
}
