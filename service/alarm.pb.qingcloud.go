// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: alarm.proto

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

// See https://docs.qingcloud.com/api/alarm/index.html
type AlarmServiceInterface interface {
	DescribeAlarmPolicies(in *DescribeAlarmPoliciesInput) (out *DescribeAlarmPoliciesOutput, err error)
	CreateAlarmPolicy(in *CreateAlarmPolicyInput) (out *CreateAlarmPolicyOutput, err error)
	ModifyAlarmPolicyAttributes(in *ModifyAlarmPolicyAttributesInput) (out *ModifyAlarmPolicyAttributesOutput, err error)
	DeleteAlarmPolicies(in *DeleteAlarmPoliciesInput) (out *DeleteAlarmPoliciesOutput, err error)
	DescribeAlarmPolicyRules(in *DescribeAlarmPolicyRulesInput) (out *DescribeAlarmPolicyRulesOutput, err error)
	AddAlarmPolicyRules(in *AddAlarmPolicyRulesInput) (out *AddAlarmPolicyRulesOutput, err error)
	ModifyAlarmPolicyRuleAttributes(in *ModifyAlarmPolicyRuleAttributesInput) (out *ModifyAlarmPolicyRuleAttributesOutput, err error)
	DeleteAlarmPolicyRules(in *DeleteAlarmPolicyRulesInput) (out *DeleteAlarmPolicyRulesOutput, err error)
	DescribeAlarmPolicyActions(in *DescribeAlarmPolicyActionsInput) (out *DescribeAlarmPolicyActionsOutput, err error)
	AddAlarmPolicyActions(in *AddAlarmPolicyActionsInput) (out *AddAlarmPolicyActionsOutput, err error)
	ModifyAlarmPolicyActionAttributes(in *ModifyAlarmPolicyActionAttributesInput) (out *ModifyAlarmPolicyActionAttributesOutput, err error)
	DeleteAlarmPolicyActions(in *DeleteAlarmPolicyActionsInput) (out *DeleteAlarmPolicyActionsOutput, err error)
	AssociateAlarmPolicy(in *AssociateAlarmPolicyInput) (out *AssociateAlarmPolicyOutput, err error)
	DissociateAlarmPolicy(in *DissociateAlarmPolicyInput) (out *DissociateAlarmPolicyOutput, err error)
	ApplyAlarmPolicy(in *ApplyAlarmPolicyInput) (out *ApplyAlarmPolicyOutput, err error)
	DescribeAlarms(in *DescribeAlarmsInput) (out *DescribeAlarmsOutput, err error)
	DescribeAlarmHistory(in *DescribeAlarmHistoryInput) (out *DescribeAlarmHistoryOutput, err error)
}

// See https://docs.qingcloud.com/api/alarm/index.html
type AlarmService struct {
	Config           *config.Config
	Properties       *AlarmServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/alarm/index.html
func NewAlarmService(conf *config.Config, zone string) (p *AlarmService) {
	return &AlarmService{
		Config:     conf,
		Properties: &AlarmServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/alarm/index.html
func (s *QingCloudService) Alarm(zone string) (*AlarmService, error) {
	properties := &AlarmServiceProperties{
		Zone: zone,
	}

	return &AlarmService{Config: s.Config, Properties: properties}, nil
}

func (p *AlarmService) DescribeAlarmPolicies(in *DescribeAlarmPoliciesInput) (out *DescribeAlarmPoliciesOutput, err error) {
	if in == nil {
		in = &DescribeAlarmPoliciesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeAlarmPoliciesOutput{}
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

func (p *AlarmService) CreateAlarmPolicy(in *CreateAlarmPolicyInput) (out *CreateAlarmPolicyOutput, err error) {
	if in == nil {
		in = &CreateAlarmPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "CreateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &CreateAlarmPolicyOutput{}
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

func (p *AlarmService) ModifyAlarmPolicyAttributes(in *ModifyAlarmPolicyAttributesInput) (out *ModifyAlarmPolicyAttributesOutput, err error) {
	if in == nil {
		in = &ModifyAlarmPolicyAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyAlarmPolicyAttributesOutput{}
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

func (p *AlarmService) DeleteAlarmPolicies(in *DeleteAlarmPoliciesInput) (out *DeleteAlarmPoliciesOutput, err error) {
	if in == nil {
		in = &DeleteAlarmPoliciesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicies",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteAlarmPoliciesOutput{}
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

func (p *AlarmService) DescribeAlarmPolicyRules(in *DescribeAlarmPolicyRulesInput) (out *DescribeAlarmPolicyRulesOutput, err error) {
	if in == nil {
		in = &DescribeAlarmPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeAlarmPolicyRulesOutput{}
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

func (p *AlarmService) AddAlarmPolicyRules(in *AddAlarmPolicyRulesInput) (out *AddAlarmPolicyRulesOutput, err error) {
	if in == nil {
		in = &AddAlarmPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &AddAlarmPolicyRulesOutput{}
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

func (p *AlarmService) ModifyAlarmPolicyRuleAttributes(in *ModifyAlarmPolicyRuleAttributesInput) (out *ModifyAlarmPolicyRuleAttributesOutput, err error) {
	if in == nil {
		in = &ModifyAlarmPolicyRuleAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyRuleAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyAlarmPolicyRuleAttributesOutput{}
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

func (p *AlarmService) DeleteAlarmPolicyRules(in *DeleteAlarmPolicyRulesInput) (out *DeleteAlarmPolicyRulesOutput, err error) {
	if in == nil {
		in = &DeleteAlarmPolicyRulesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyRules",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteAlarmPolicyRulesOutput{}
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

func (p *AlarmService) DescribeAlarmPolicyActions(in *DescribeAlarmPolicyActionsInput) (out *DescribeAlarmPolicyActionsOutput, err error) {
	if in == nil {
		in = &DescribeAlarmPolicyActionsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeAlarmPolicyActionsOutput{}
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

func (p *AlarmService) AddAlarmPolicyActions(in *AddAlarmPolicyActionsInput) (out *AddAlarmPolicyActionsOutput, err error) {
	if in == nil {
		in = &AddAlarmPolicyActionsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AddAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &AddAlarmPolicyActionsOutput{}
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

func (p *AlarmService) ModifyAlarmPolicyActionAttributes(in *ModifyAlarmPolicyActionAttributesInput) (out *ModifyAlarmPolicyActionAttributesOutput, err error) {
	if in == nil {
		in = &ModifyAlarmPolicyActionAttributesInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ModifyAlarmPolicyActionAttributes",
		RequestMethod: "GET", // GET or POST
	}

	x := &ModifyAlarmPolicyActionAttributesOutput{}
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

func (p *AlarmService) DeleteAlarmPolicyActions(in *DeleteAlarmPolicyActionsInput) (out *DeleteAlarmPolicyActionsOutput, err error) {
	if in == nil {
		in = &DeleteAlarmPolicyActionsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DeleteAlarmPolicyActions",
		RequestMethod: "GET", // GET or POST
	}

	x := &DeleteAlarmPolicyActionsOutput{}
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

func (p *AlarmService) AssociateAlarmPolicy(in *AssociateAlarmPolicyInput) (out *AssociateAlarmPolicyOutput, err error) {
	if in == nil {
		in = &AssociateAlarmPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "AssociateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &AssociateAlarmPolicyOutput{}
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

func (p *AlarmService) DissociateAlarmPolicy(in *DissociateAlarmPolicyInput) (out *DissociateAlarmPolicyOutput, err error) {
	if in == nil {
		in = &DissociateAlarmPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DissociateAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &DissociateAlarmPolicyOutput{}
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

func (p *AlarmService) ApplyAlarmPolicy(in *ApplyAlarmPolicyInput) (out *ApplyAlarmPolicyOutput, err error) {
	if in == nil {
		in = &ApplyAlarmPolicyInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "ApplyAlarmPolicy",
		RequestMethod: "GET", // GET or POST
	}

	x := &ApplyAlarmPolicyOutput{}
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

func (p *AlarmService) DescribeAlarms(in *DescribeAlarmsInput) (out *DescribeAlarmsOutput, err error) {
	if in == nil {
		in = &DescribeAlarmsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarms",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeAlarmsOutput{}
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

func (p *AlarmService) DescribeAlarmHistory(in *DescribeAlarmHistoryInput) (out *DescribeAlarmHistoryOutput, err error) {
	if in == nil {
		in = &DescribeAlarmHistoryInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeAlarmHistory",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeAlarmHistoryOutput{}
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

func (p *AlarmServiceProperties) Validate() error {
	return nil
}

func (p *DescribeAlarmPoliciesInput) Validate() error {

	{
		var _enumValues = []string{
			"instance",
			"eip",
			"router",
			"loadbalancer_listener_tcp",
			"loadbalancer_listener_https",
			"loadbalancer_listener_http",
			"loadbalancer_backend_http",
			"loadbalancer_backend_tcp",
			"rdb_mysql",
			"rdb_psql",
			"elasticsearch_node",
			"cache_node_memcached",
			"cache_node_redis",
			"queue_node_kafka",
			"spark_node",
			"hadoop_node",
			"zookeeper_node",
			"hbase_node",
			"mongo_instance",
			"storm_node",
		}
		for _, v := range _enumValues {
			if p.GetAlarmPolicyType() != v {
				return fmt.Errorf("DescribeAlarmPoliciesInput.AlarmPolicyType: invalid enum value!")
			}
		}
	}

	return nil
}

func (p *DescribeAlarmPoliciesOutput) Validate() error {
	return nil
}

func (p *CreateAlarmPolicyInput) Validate() error {
	return nil
}

func (p *CreateAlarmPolicyOutput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyAttributesInput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyAttributesOutput) Validate() error {
	return nil
}

func (p *DeleteAlarmPoliciesInput) Validate() error {
	return nil
}

func (p *DeleteAlarmPoliciesOutput) Validate() error {
	return nil
}

func (p *DescribeAlarmPolicyRulesInput) Validate() error {

	return nil
}

func (p *DescribeAlarmPolicyRulesOutput) Validate() error {
	return nil
}

func (p *AddAlarmPolicyRulesInput) Validate() error {
	return nil
}

func (p *AddAlarmPolicyRulesOutput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyRuleAttributesInput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyRuleAttributesOutput) Validate() error {
	return nil
}

func (p *DeleteAlarmPolicyRulesInput) Validate() error {
	return nil
}

func (p *DeleteAlarmPolicyRulesOutput) Validate() error {
	return nil
}

func (p *DescribeAlarmPolicyActionsInput) Validate() error {
	return nil
}

func (p *DescribeAlarmPolicyActionsOutput) Validate() error {
	return nil
}

func (p *AddAlarmPolicyActionsInput) Validate() error {
	return nil
}

func (p *AddAlarmPolicyActionsOutput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyActionAttributesInput) Validate() error {
	return nil
}

func (p *ModifyAlarmPolicyActionAttributesOutput) Validate() error {
	return nil
}

func (p *DeleteAlarmPolicyActionsInput) Validate() error {
	return nil
}

func (p *DeleteAlarmPolicyActionsOutput) Validate() error {
	return nil
}

func (p *AssociateAlarmPolicyInput) Validate() error {
	return nil
}

func (p *AssociateAlarmPolicyOutput) Validate() error {
	return nil
}

func (p *DissociateAlarmPolicyInput) Validate() error {
	return nil
}

func (p *DissociateAlarmPolicyOutput) Validate() error {
	return nil
}

func (p *ApplyAlarmPolicyInput) Validate() error {
	return nil
}

func (p *ApplyAlarmPolicyOutput) Validate() error {
	return nil
}

func (p *DescribeAlarmsInput) Validate() error {
	return nil
}

func (p *DescribeAlarmsOutput) Validate() error {
	return nil
}

func (p *DescribeAlarmHistoryInput) Validate() error {
	return nil
}

func (p *DescribeAlarmHistoryOutput) Validate() error {
	return nil
}
