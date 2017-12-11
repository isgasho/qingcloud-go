<p align="center"><a href="http://qingcloud.com" target="_blank"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/docs/images/logo-01-600x130.png" alt="QingCloud"></a></p>

# 青云 SDK Go Version

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg?branch=master)](https://travis-ci.org/chai2010/qingcloud-go)
[![Docker Build Status](https://img.shields.io/docker/build/chai2010/qingcloud-go.svg)](https://hub.docker.com/r/chai2010/qingcloud-go/)
[![GoDoc](https://godoc.org/github.com/chai2010/qingcloud-go?status.svg)](https://godoc.org/github.com/chai2010/qingcloud-go)
[![API Reference](http://img.shields.io/badge/api-reference-green.svg)](http://docs.qingcloud.com)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/qingcloud-go/blob/master/LICENSE)

<!--

新开 SDK 的原因:

- 官方 SDK 迟迟没有 release, 项目活跃度较低
- 官方 SDK 的 [近2万行手写的json](https://github.com/yunify/qingcloud-api-specs/tree/master/2013-08-30/swagger) 维护极其困难, 已经失去继续进化的可能
- 官方 SDK 需要用 Go 语言的模板语言维护 snips 自定义的 [极其复杂的判断逻辑](https://github.com/yunify/qingcloud-sdk-go/blob/master/template/shared.tmpl) , 这是错误的!
- 基于 Ptotobuf 构建, 解决方案简单优美, 稳定性和可扩展性足够好, Docker/Kubernetes 都用它
- 其它通过非主流的工具构建的方式, 除了作者本人根本没有投入的必要
- 缺少一个完备的和 SDK 基本等价的命令行工具

-->

项目目标:

- Go语言风格的接口, 简单好用是第一目标
- 一个好用的命令行工具, 希望一切可以通过命令行脚本化
- 基于 Protobuf-V2 语法维护规范, 便于升级和维护
- 更完备的接口, 更多的测试 ....

在线文档:

- https://docs.qingcloud.com
- https://godoc.org/github.com/chai2010/qingcloud-go

国内镜像:

- https://gitee.com/chai2010/qingcloud-go

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

json_disable_unknown_fields: false

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: 'warn'
```

将 `qy_access_key_id` 和 `qy_secret_access_key` 字段替换为 API密钥 中的内容.

其中 `json_disable_unknown_fields` 是新加的变量, 表示在JSON解码时忽略 proto.Message 遇到未定义成员的错误.

## qcli 命令行

Docker 运行:

- `docker run --rm -it -v $HOME:/root -w /root chai2010/qingcloud-go qcli`

从Go源码安装:

- `go get github.com/chai2010/qingcloud-go/cmd/qcli`

或生成版本号后安装:

- `go generate github.com/chai2010/qingcloud-go/pkg/version`
- `go install  github.com/chai2010/qingcloud-go/cmd/qcli`

输入 `qcli` 或 `qcli -h` 查看命令提示:

```
chai-mba:api chai$ qcli
NAME:
   qcli - QingCloud Command Line Interface

USAGE:
   qcli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   ChaiShushan <chaishushan@gmail.com>

COMMANDS:
     alarm               manage AlarmService
     cache               manage CacheService
     cluster             manage ClusterService
     dnsalias            manage DNSAliasService
     eip                 manage EIPService
     hadoop              manage HadoopService
     image               manage ImageService
     instance            manage InstanceService
     job                 manage JobService
     keypair             manage KeyPairService
     loadbalancer        manage LoadBalancerService
     misc                manage MiscService
     mongo               manage MongoService
     monitor             manage MonitorService
     nic                 manage NicService
     notificationcenter  manage NotificationCenterService
     rdb                 manage RDBService
     resourceacl         manage ResourceACLService
     router              manage RouterService
     s2                  manage S2Service
     securitygroup       manage SecurityGroupService
     snapshot            manage SnapshotService
     span                manage SpanService
     spark               manage SparkService
     subuser             manage SubuserService
     tag                 manage TagService
     userdata            manage UserDataService
     volumes             manage VolumesService
     vxnet               manage VxnetService
     zone                manage ZoneService
     help, h             Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  config file (default: "~/.qingcloud/config.yaml") [$QCLI_CONFIG_FILE]
   --zone value, -z value    zone (pk3a,pk3b,gd1,sh1a,ap1,ap2a,...) (default: "pk3a") [$QCLI_ZONE]
   --glog_level value        glog level to stderr (INFO,WARNING,ERROR,FATAL) (default: "WARNING") [$QCLI_GLOG_LEVEL]
   --help, -h                show help
   --version, -v             print the version
chai-mba:qingcloud-go chai$
```

*注意: 命令行还在开发中, 欢迎参与完善!*

## 快速入门

以下为 [docs/hello.go](./docs/hello.go) 的内容:

```go
package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/config"
	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func main() {
	// 返回 NIC 服务, pek3a 为 北京3区-A
	nicService := pb.NewNicService(config.MustLoadUserConfig(), "pek3a")

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

	go run ./docs/hello.go

[更多例子](docs/examples).

## 文档指南

使用青云SDK一般是以下步骤:

1. 用 [pkg/config](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/config) 包构造一个配置对象, 里面含有最重要的 API密钥, 还包含日志级别等信息.
2. 基于配置对象调用 [pkg/api](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/api) 包的 [`Init`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/api#Init) 函数构造一个青云主服务对象 [`qcService`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/api#QingCloudService), 其中会根据配置文件设置日志级别.
3. 假设有一个 [UserData](./api/user_data.proto) 子服务, 那么调用 [`qcService.UserData("pek3a")`](https://godoc.org/github.com/chai2010/qingcloud-go/pkg/api#QingCloudService.UserData) 方法将返回子服务对象, 其中参数是区域
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

规范文件的语法细节可以参考 [api/README.md](./api/README.md), proto3 文件语法可以参考 [Protobuf](https://developers.google.cn/protocol-buffers/docs/proto3) 的官方文档.

## 官方 SDK 对比

[Volume](https://docs.qingcloud.com/api/volume/index.html) 服务接口规范文件对比:

- protobuf 格式: [chai2010/qingcloud-go/api/volume.proto](./api/volume.proto)
- snips 格式: [yunify/qingcloud-api-specs/2013-08-30/swagger/volume.json](https://github.com/yunify/qingcloud-api-specs/blob/master/2013-08-30/swagger/volume.json)


## 版权

The Apache License.
