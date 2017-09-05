// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eip.proto

package spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"

import "context"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EipService interface {
	DescribeEips(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AllocateEips(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ReleaseEips(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	AssociateEip(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DissociateEips(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ChangeEipsBandwidth(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ChangeEipsBillingMode(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyEipAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type EipServiceClient struct{}

// NewEipServiceClient returns a EipService stub to handle
// requests to the set of EipService at the other end of the connection.
func NewEipServiceClient(opt *Options) *EipServiceClient {
	return &EipServiceClient{}
}

func (c *EipServiceClient) DescribeEips(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) AllocateEips(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) ReleaseEips(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) AssociateEip(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) DissociateEips(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) ChangeEipsBandwidth(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) ChangeEipsBillingMode(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *EipServiceClient) ModifyEipAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}

func init() { proto.RegisterFile("eip.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x2f, 0x45, 0x70, 0x15, 0x0f, 0x11, 0x05, 0x23, 0xe8, 0x1b, 0xa4, 0xa0, 0x67, 0xc5,
	0x6a, 0x17, 0x4f, 0xbd, 0xe8, 0x13, 0x6c, 0x36, 0xd3, 0x74, 0x60, 0x93, 0x19, 0x76, 0xa6, 0x96,
	0xbe, 0x8f, 0x0f, 0x2a, 0x4d, 0x2c, 0x84, 0x40, 0x0f, 0x9b, 0xe3, 0xce, 0x3f, 0xdf, 0xf7, 0xcf,
	0x9a, 0x73, 0x40, 0x2e, 0x38, 0x92, 0x52, 0x36, 0x13, 0x06, 0x9f, 0xdf, 0xd5, 0x44, 0x75, 0x80,
	0x79, 0x37, 0x2b, 0xb7, 0xeb, 0xb9, 0x6b, 0xf7, 0xfd, 0x42, 0x7e, 0x3f, 0x8e, 0xa0, 0x61, 0x3d,
	0x86, 0x8f, 0xe3, 0x50, 0xb1, 0x01, 0x51, 0xd7, 0xfc, 0xeb, 0xf3, 0x87, 0xf1, 0xc2, 0x2e, 0x3a,
	0x66, 0x88, 0xd2, 0xe7, 0x4f, 0xbf, 0x33, 0x63, 0x2c, 0xf2, 0x37, 0xc4, 0x1f, 0xf4, 0x90, 0xbd,
	0x9a, 0xcb, 0x25, 0x88, 0x8f, 0x58, 0x82, 0x45, 0x96, 0xec, 0xb6, 0xe8, 0xf9, 0xe2, 0xc8, 0x17,
	0xf6, 0xd0, 0x9e, 0x9f, 0x98, 0x1f, 0xf8, 0x45, 0x08, 0xe4, 0x9d, 0x4e, 0xe3, 0x5f, 0xcc, 0xc5,
	0x17, 0x04, 0x70, 0x32, 0xbd, 0x5e, 0x84, 0x3c, 0xf6, 0xfd, 0xc9, 0xfc, 0x9b, 0xb9, 0x5a, 0xe2,
	0x40, 0x90, 0x7e, 0x81, 0x35, 0xd7, 0x1f, 0x1b, 0xd7, 0xd6, 0x1d, 0xfd, 0xee, 0xda, 0x6a, 0x87,
	0x95, 0x6e, 0x92, 0x35, 0x9f, 0xe6, 0x66, 0xa0, 0xc1, 0x10, 0xb0, 0xad, 0x57, 0x54, 0xc1, 0x94,
	0x7b, 0x56, 0x54, 0xe1, 0x7a, 0x6f, 0x91, 0x17, 0xaa, 0x11, 0xcb, 0xad, 0x42, 0xf2, 0xb7, 0xca,
	0xb3, 0xee, 0xfd, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x00, 0xdc, 0x33, 0x7e, 0xb9, 0x02, 0x00,
	0x00,
}
