// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: user_data.proto

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

type UserDataServiceInterface interface {
	UploadUserDataAttachment(in *UploadUserDataAttachmentInput) (out *UploadUserDataAttachmentOutput, err error)
}

type UserDataService struct {
	Config           *config.Config
	Properties       *UserDataServiceProperties
	LastResponseBody string
}

func NewUserDataService(conf *config.Config, zone string) (p *UserDataService) {
	return &UserDataService{
		Config:     conf,
		Properties: &UserDataServiceProperties{Zone: proto.String(zone)},
	}
}

func (p *UserDataService) UploadUserDataAttachment(in *UploadUserDataAttachmentInput) (out *UploadUserDataAttachmentOutput, err error) {
	if in == nil {
		in = &UploadUserDataAttachmentInput{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "UploadUserDataAttachment",
		RequestMethod: "POST",
	}

	x := &UploadUserDataAttachmentOutput{}
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
