// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: user_data.proto

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

type UserDataServiceInterface interface {
	UploadUserDataAttachment(in *UploadUserDataAttachmentInput) (out *UploadUserDataAttachmentOutput, err error)
}

type UserDataService struct {
	ServerInfo *ServerInfo
}

func NewUserDataService(server *ServerInfo) (p *UserDataService) {
	return &UserDataService{
		ServerInfo: server,
	}
}

func init() {
	ServiceApiSpecMap["UploadUserDataAttachment"] = ServiceApiSpec{
		ServiceName:    "UserDataService",
		ActionName:     "UploadUserDataAttachment",
		InputTypeName:  "UploadUserDataAttachmentInput",
		OutputTypeName: "UploadUserDataAttachmentOutput",
		HttpMethod:     "POST",

		ServiceType: reflect.TypeOf((*UserDataService)(nil)),
		InputType:   reflect.TypeOf((*UploadUserDataAttachmentInput)(nil)),
		OutputType:  reflect.TypeOf((*UploadUserDataAttachmentOutput)(nil)),
	}
}

func (p *UserDataService) UploadUserDataAttachment(input *UploadUserDataAttachmentInput) (output *UploadUserDataAttachmentOutput, err error) {
	client := client.NewClient(
		p.ServerInfo.GetApiServer(),
		p.ServerInfo.GetAccessKeyId(),
		p.ServerInfo.GetSecretAccessKey(),
		p.ServerInfo.GetZone(),
	)

	if input == nil {
		input = new(UploadUserDataAttachmentInput)
	}

	output = new(UploadUserDataAttachmentOutput)
	err = client.CallMethod("UploadUserDataAttachment", "POST", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
