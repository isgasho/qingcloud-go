<p align="center"><a href="http://qingcloud.com" target="_blank"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/logo.jpg" alt="QingCloud"></a></p>

# 青云 SDK Go Version

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg)](https://travis-ci.org/chai2010/qingcloud-go)
[![GoDoc](https://godoc.org/github.com/chai2010/qingcloud-go?status.svg)](https://godoc.org/github.com/chai2010/qingcloud-go)
[![API Reference](http://img.shields.io/badge/api-reference-green.svg)](http://docs.qingcloud.com)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/qingcloud-go/blob/master/LICENSE)

项目目标:

- Go语言风格的接口, 简单好用是第一目标
- 基于 Protobuf-V2 语法维护规范, 便于升级和维护
- 更完备的接口, 更多的测试 ....

在线文档:

- https://docs.qingcloud.com
- https://godoc.org/github.com/chai2010/qingcloud-go

接口规范:

- [api](api)

## 配置文件

当前用户的配置文件在 `${HOME}/.qingcloud/config.yaml`, 内容如下:

```yaml
# QingCloud services configuration

qy_access_key_id: 'ACCESS_KEY_ID'
qy_secret_access_key: 'SECRET_ACCESS_KEY'

host: 'api.qingcloud.com'
port: 443
protocol: 'https'
uri: '/iaas'
connection_retries: 3

json_disable_unknown_fields: false

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: 'warn'
```

将 `qy_access_key_id` 和 `qy_secret_access_key` 字段替换为 API密钥 中的内容.

其中 `json_disable_unknown_fields` 是新加的变量, 表示在JSON解码时忽略 proto.Message 遇到未定义成员的错误.

## 快速入门

以下为 [hello.go](./hello.go) 的内容:

```go
package main

import (
	"fmt"
	"log"

	"github.com/chai2010/qingcloud-go/pkg/config"
	pb "github.com/chai2010/qingcloud-go/pkg/service.pb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// 初始化 青云 服务对象
	qcService, err := pb.Init(config.MustLoadUserConfig())
	if err != nil {
		log.Fatal(err)
	}

	// 返回 NIC 子服务, pek3a 为 北京3区-A
	nicService, err := qcService.Nic("pek3a")
	if err != nil {
		log.Fatal(err)
	}

	// 列出所有网卡
	reply, err := nicService.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 原始返回的json数据
	// nicService.LastResponseBody

	// JSON 格式打印
	fmt.Println(jsonpbEncode(reply))
}

// pb转json, 采用原始名称, 不忽略空值
func jsonpbEncode(m proto.Message) string {
	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(m)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
```

初始化子服务也可以用以下方式:

```go
nicService := pb.NewNicService(config.MustLoadUserConfig(), "pek3a")
```

运行例子:

	go run hello.go

[更多例子](docs/examples).

## 文档指南

使用青云SDK一般是以下步骤:

1. 用 [pkg/config](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/config) 包构造一个配置对象, 里面含有最重要的 API密钥, 还包含日志级别等信息.
2. 基于配置对象调用 [pkg/service.pb](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/service.pb) 包的 [`Init`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/service.pb#Init) 函数构造一个青云主服务对象 [`qcService`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/service.pb#QingCloudService), 其中会根据配置文件设置日志级别.
3. 假设有一个 [UserData](./api/user_data.proto) 子服务, 那么调用 [`qcService.UserData("pek3a")`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/service.pb#QingCloudService.UserData) 方法将返回子服务对象, 其中参数是区域
4. 使用子服务对象就可以调用每个子对象的方法了

我们可以查看子服务对应的接口规范, 在 [api/user_data.proto](./api/user_data.proto) 文件定义 ([青云文档](https://docs.qingcloud.com/api/userdata/index.html)):

```proto
service UserDataService {
	rpc UploadUserDataAttachment(UploadUserDataAttachmentInput) returns (UploadUserDataAttachmentOutput);
}

message UploadUserDataAttachmentInput {
	bytes attachment_content = 2;
	string attachment_name = 1;
}

message UploadUserDataAttachmentOutput {
	string action = 1;
	int32 ret_code = 2;
	string message = 3;

	string attachment_id = 4;
}
```

其中`service`关键字开头的表示定义一组子服务, 其中`rpc`开头的表示子服务中每个具体的方法. 方法的输入参数和返回值分别为`UploadUserDataAttachmentInput`和`UploadUserDataAttachmentInput`结构体类型, 它们由后面的`message`关键字定义.

[SDK的代码生成插件](./pkg/cmd/protoc-gen-qingcloud-go/qingcloud/qingcloud.go) 会生成以下的Go语言代码:

```go
type UserDataService struct {
	// ...
}

func NewUserDataService(conf *config.Config, zone string) *UserDataService {
	// ...
}

func (p *QingCloudService) UserData(zone string) (*UserDataService, error) {
	// ...
}

type UploadUserDataAttachmentInput struct {
	// ...
}
type UploadUserDataAttachmentOutput struct {
	// ...
}

func (p *UserDataService) UploadUserDataAttachment(
	in *UploadUserDataAttachmentInput,
) (
	*UploadUserDataAttachmentOutput,
	error,
) {
	// ...
}
```

规范文件的语法细节可以参考 [spec.pb/README.md](./api/README.md), proto3 文件语法可以参考 [Protobuf](https://developers.google.cn/protocol-buffers/docs/proto3) 的官方文档.

## 与官方SDK的兼容性

因为 protobuf 语法的限制, 它无法实现对 多维数组/字典数组/异构数据 对支持. 目前对于异构数据对处理思路是, 避免用 protobuf 实现异构数据, 采用纯Go语言对结构体实现(内部不得使用protobuf定义对时间类型). 青云的服务中绝大部分参数都是对 protobuf 友好对. 我们尽量做到以下几点:

- 即使有不兼容的地方, API 也是非常相似的
- 该 SDK 和 官方 SDK 的 API 保持最大的兼容性

假设青云的REST规范的文档中有一个名为 `job_id` 的输入参数, 对应 `XXXInput` 结构体的成员.

官方文档是根据 [json定义的规范](https://github.com/yunify/qingcloud-api-specs/tree/master/2013-08-30/swagger), 然后通过一个名为 [snips](https://github.com/yunify/snips) 的工具加自己定义的 [模板](https://github.com/yunify/qingcloud-sdk-go/tree/master/template) 生成的代码, `XXXInput` 输入参数生成的代码可能类似以下结构:

```go
type XXXInput struct {
	JobID *string `json:"job_id" name:"job_id" location:"elements"`
}
```

而我们的SDK采用Protobuf3标准工具生成的代码:

```go
type XXXInput struct {
	JobId string   `protobuf:"bytes,5,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}
```

其中有两个大的差异: 一个是成员名称不同, 分别为 `JobID` 和 `JobId`; 另一个为类型不同, 分别为 `*string` 和 `string`.

[snips](https://github.com/yunify/snips) 采用和 Protobuf-V2 类似的生成规则, 零值是 `nil`, 空值是空字符串, 二者是不等价的. 在 Protobuf3 的生成规则中, 默认将零值和空值等价.

注意: 经过慎重考虑, 我们绝对转为 Protobuf-V2 语法, 实现对默认值的支持.


## 为何不用官方SDK, 为何要重新做一个?

官方的 SDK 对青云的服务支持还不完全, 也没有发布正式的版本. 原因之一是, 官方SDK采用json格式定义规范很难维护, 特别是有几十甚至上百个成员时, 手工维护json规范极其困难, 目前官方SDK已经落后于当前服务(官方的json规范共有约1.9万行, 我们的proto规范文件不到8千行). 另外还要花很大精力维护 snips 本身的开发, 完善各种语言多各种复合数据类型的支持. 而使用protobuf标准, 可以很容易编写服务规范文件, 官方的编译工具可以针对各种主流编程语言生成代码(我们只需要定制服务部分的代码生成规则), 也便于以后和docker和k8s等平台互联网(它们都是提供的grpc接口, 也是protobuf规范定义的).

可以看看 [Volume](https://docs.qingcloud.com/api/volume/index.html) 服务规范的对比:

- proto3 格式: [chai2010/qingcloud-go/api/volume.proto](./api/volume.proto)
- snips 格式: [yunify/qingcloud-api-specs/2013-08-30/swagger/volume.json](https://github.com/yunify/qingcloud-api-specs/blob/master/2013-08-30/swagger/volume.json)

对比可发现, protobuf 比 json 更容易维护.

<!-- 以后或许可以通过pb自动生成json规范 -->

## 外部依赖

- Protobuf: [https://github.com/golang/protobuf](https://github.com/golang/protobuf)

## 版权

The Apache License.
