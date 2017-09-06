// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_data.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/request"
import request_data_pkg "github.com/chai2010/qingcloud-go/request/data"
import "github.com/chai2010/qingcloud-go/request/errors"

var _ = config.Config{}
var _ = request.Request{}
var _ = request_data_pkg.Operation{}
var _ = errors.ParameterRequiredError{}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserDataServiceProperties struct {
	Zone string `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
}

func (m *UserDataServiceProperties) Reset()                    { *m = UserDataServiceProperties{} }
func (m *UserDataServiceProperties) String() string            { return proto.CompactTextString(m) }
func (*UserDataServiceProperties) ProtoMessage()               {}
func (*UserDataServiceProperties) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

func (m *UserDataServiceProperties) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDataServiceProperties)(nil), "spec.UserDataServiceProperties")
}

type UserDataServiceInterface interface {
	UploadUserDataAttachment(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type UserDataService struct {
	Config     *config.Config
	Properties *UserDataServiceProperties
}

func NewUserDataService(conf *config.Config, zone string) (p *UserDataService, err error) {
	return &UserDataService{
		Config:     conf,
		Properties: &UserDataServiceProperties{Zone: zone},
	}, nil
}

func (p *UserDataService) UploadUserDataAttachment(in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("user_data.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x41, 0x6b, 0x83, 0x40,
	0x10, 0x46, 0x11, 0xa4, 0xd0, 0xbd, 0x08, 0x7b, 0x28, 0x6a, 0xa1, 0x2d, 0x3d, 0xf5, 0xb4, 0x42,
	0xfb, 0x0b, 0x0a, 0xed, 0xa5, 0xa7, 0x92, 0xe0, 0x31, 0x84, 0x51, 0x27, 0x46, 0x70, 0xdd, 0x61,
	0x76, 0x4c, 0x48, 0x7e, 0x7d, 0x88, 0xc6, 0xcb, 0x42, 0x6e, 0xbb, 0xdf, 0x7b, 0x03, 0x4f, 0x25,
	0xa3, 0x47, 0xde, 0x36, 0x20, 0x60, 0x88, 0x9d, 0x38, 0x1d, 0x7b, 0xc2, 0x3a, 0xcf, 0x5a, 0xe7,
	0xda, 0x1e, 0x8b, 0x69, 0xab, 0xc6, 0x5d, 0x01, 0xc3, 0x69, 0x16, 0xf2, 0xe7, 0x10, 0xa1, 0x25,
	0x59, 0xe0, 0x6b, 0x08, 0xa5, 0xb3, 0xe8, 0x05, 0x2c, 0xdd, 0x84, 0x97, 0x50, 0x38, 0x32, 0x10,
	0x21, 0xfb, 0x99, 0xbf, 0x17, 0x2a, 0x2b, 0x3d, 0xf2, 0x0f, 0x08, 0xac, 0x91, 0x0f, 0x5d, 0x8d,
	0xff, 0xec, 0x08, 0x59, 0x3a, 0xf4, 0x5a, 0xab, 0xf8, 0xec, 0x06, 0x4c, 0xa3, 0xb7, 0xe8, 0xe3,
	0x71, 0x35, 0xbd, 0x3f, 0x37, 0x2a, 0x09, 0x0e, 0xf4, 0x9f, 0x4a, 0x4b, 0xea, 0x1d, 0x34, 0x0b,
	0xf8, 0x16, 0x81, 0x7a, 0x6f, 0x71, 0x10, 0xfd, 0x64, 0xe6, 0x00, 0xb3, 0x04, 0x98, 0xdf, 0x6b,
	0x7e, 0x7e, 0x67, 0xaf, 0x1e, 0xa6, 0xff, 0xd7, 0x25, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x68, 0x24,
	0x18, 0x28, 0x01, 0x00, 0x00,
}
