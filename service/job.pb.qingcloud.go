// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: job.proto

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

// See https://docs.qingcloud.com/api/job/index.html
type JobServiceInterface interface {
	DescribeJobs(in *DescribeJobsInput) (out *DescribeJobsOutput, err error)
}

// See https://docs.qingcloud.com/api/job/index.html
type JobService struct {
	Config           *config.Config
	Properties       *JobServiceProperties
	LastResponseBody string
}

// See https://docs.qingcloud.com/api/job/index.html
func NewJobService(conf *config.Config, zone string) (p *JobService) {
	return &JobService{
		Config:     conf,
		Properties: &JobServiceProperties{Zone: zone},
	}
}

// See https://docs.qingcloud.com/api/job/index.html
func (s *QingCloudService) Job(zone string) (*JobService, error) {
	properties := &JobServiceProperties{
		Zone: zone,
	}

	return &JobService{Config: s.Config, Properties: properties}, nil
}

func (p *JobService) DescribeJobs(in *DescribeJobsInput) (out *DescribeJobsOutput, err error) {
	if in == nil {
		in = &DescribeJobsInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeJobs",
		RequestMethod: "GET", // GET or POST
	}

	x := &DescribeJobsOutput{}
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

func (p *JobServiceProperties) Validate() error {
	return nil
}

func (p *DescribeJobsInput) Validate() error {
	return nil
}

func (p *DescribeJobsOutput) Validate() error {
	return nil
}