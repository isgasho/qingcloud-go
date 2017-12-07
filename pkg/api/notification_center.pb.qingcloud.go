// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: notification_center.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}

type NotificationCenterServiceInterface interface {
	DescribeNotificationCenterUserPosts(in *DescribeNotificationCenterUserPostsInput) (out *DescribeNotificationCenterUserPostsOutput, err error)
}

type NotificationCenterService struct {
	Config           *config.Config
	Properties       *NotificationCenterServiceProperties
	LastResponseBody string
}

func NewNotificationCenterService(conf *config.Config, zone string) (p *NotificationCenterService) {
	return &NotificationCenterService{
		Config:     conf,
		Properties: &NotificationCenterServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *NotificationCenterService) DescribeNotificationCenterUserPosts(in *DescribeNotificationCenterUserPostsInput) (out *DescribeNotificationCenterUserPostsOutput, err error) {
	if in == nil {
		in = &DescribeNotificationCenterUserPostsInput{}
	}
	o := &request.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "DescribeNotificationCenterUserPosts",
		RequestMethod: "GET",
	}

	x := &DescribeNotificationCenterUserPostsOutput{}
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
