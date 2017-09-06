
<p align="center"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/logo.jpg" alt="QingCloud"></p>

# qingcloud 自用 SDK

在线文档:

- https://godoc.org/github.com/chai2010/qingcloud-go
- https://godoc.org/github.com/chai2010/qingcloud-go/spec.pb
- https://docs.qingcloud.com

## 设计思路

- 基于 Protobuf3 语法对应的 json 格式文件定义 api 元数据
- 基于 Protobuf3 语法定义 rest 接口的请求和响应结构体, 只是用于结构体, 请求时转 json 处理(proto库自带)
- 内置 Go 语言代码模板基于标准库语法

小目标:

- api 元数据采用语言无关的 Proto3 语法定义, 文件采用 proto3 对应的 json 格式
- 生成的请求和响应结构体保持和 Proto3 生成的结构一致, Proto3 本身支持多种语言
- 更多的单元测试代码
- 更多的示例代码
- 最小化外部依赖

## 外部依赖

- [github.com/golang/glog](https://github.com/golang/glog)
- [github.com/golang/protobuf](https://github.com/golang/protobuf)

其中 glog 是全局的日志包, SDK 不能放到 vendor 目录(可能导致多个glog冲突).

## 代码预览

- [base.proto](./spec.pb/base.proto)
- [base.pb.go](./spec.pb/base.pb.go)
- [instances.proto](./spec.pb/instances.proto)
- [instances.pb.go](./spec.pb/instances.pb.go)

主机服务代码:

```go
type InstancesService interface {
	DescribeInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RunInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	TerminateInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StartInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	StopInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	RestartInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResetInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ResizeInstances(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	ModifyInstanceAttributes(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DescribeInstanceTypes(ctx context.Context, in *DescribeInstanceTypes_Request) (out *DescribeInstanceTypes_Reponse, err error)
	CreateBrokers(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
	DeleteBrokers(ctx context.Context, in *google_protobuf1.Empty) (out *google_protobuf1.Empty, err error)
}

type InstancesServiceClient struct{}

// NewInstancesServiceClient returns a InstancesService stub to handle
// requests to the set of InstancesService at the other end of the connection.
func NewInstancesServiceClient(opt *Options) *InstancesServiceClient {
	return &InstancesServiceClient{}
}

func (c *InstancesServiceClient) DescribeInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) RunInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) TerminateInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) StartInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) StopInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) RestartInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ResetInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ResizeInstances(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) ModifyInstanceAttributes(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) DescribeInstanceTypes(ctx context.Context, in *DescribeInstanceTypes_Request, opt ...*Options) (out *DescribeInstanceTypes_Reponse, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) CreateBrokers(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
func (c *InstancesServiceClient) DeleteBrokers(ctx context.Context, in *google_protobuf1.Empty, opt ...*Options) (out *google_protobuf1.Empty, err error) {
	panic("TODO")
}
```