<p align="center"><a href="http://qingcloud.com" target="_blank"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/docs/images/logo-01-600x130.png" alt="QingCloud"></a></p>

# 青云 SDK Go Version

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg?branch=master)](https://travis-ci.org/chai2010/qingcloud-go)
[![Docker Build Status](https://img.shields.io/docker/build/chai2010/qingcloud-go.svg)](https://hub.docker.com/r/chai2010/qingcloud-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/qingcloud-go)](https://goreportcard.com/report/github.com/chai2010/qingcloud-go)
[![GoDoc](https://godoc.org/github.com/chai2010/qingcloud-go?status.svg)](https://godoc.org/github.com/chai2010/qingcloud-go)
[![API Reference](http://img.shields.io/badge/api-reference-green.svg)](http://docs.qingcloud.com)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/qingcloud-go/blob/master/LICENSE)

项目特色:

- Go语言风格的 SDK 封装
- qcli 命令行工具, 完全等价 SDK 功能
- 基于 Protobuf 维护规范, 便于升级和维护
- 最小化外部包依赖: SDK 仅依赖 Protobuf 包
- 更多的单元闭环测试

在线文档:

- https://docs.qingcloud.com
- https://godoc.org/github.com/chai2010/qingcloud-go

国内镜像:

- https://gitee.com/chai2010/qingcloud-go

接口规范:

- [api](api)

## qcli 命令行

Docker 运行([配置中国区镜像](https://www.docker-cn.com/registry-mirror)):

- `docker run --rm -it -v $HOME:/root -w /root chai2010/qingcloud-go qcli`

从Go源码安装(Go1.9+):

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
   --config value, -c value             config file (default: "~/.qingcloud/qcli.json") [$QCLI_CONFIG_FILE]
   --api_server value, -s value         api server (default: "https://api.qingcloud.com/iaas/") [$QCLI_API_SERVER]
   --access_key_id value, -i value      access key id [$QCLI_ACCESS_KEY_ID]
   --secret_access_key value, -k value  secret access key [$QCLI_SECRET_ACCESS_KEY]
   --zone value, -z value               zone (pek3a,pek3b,gd1,sh1a,ap1,ap2a,...) (default: "pek3a") [$QCLI_ZONE]
   --debug, -d                          set debug mode [$QCLI_DEBUG]
   --help, -h                           show help
   --version, -v                        print the version
chai-mba:qingcloud-go chai$
```

为了避免在每次运行时输入密钥, 可以给 qcli 创建一个默认配置文件 (`~/.qingcloud/qcli.json`):

```json
{
	"api_server": "https://api.qingcloud.com/iaas/",
	"access_key_id": "QYACCESSKEYIDEXAMPLE",
	"secret_access_key": "SECRETACCESSKEY",
	"zone": ""
}
```

要查看主机数量, 可以输入以下命令:

	$ qcli instance DescribeInstances

加入 `-d` 选项可以开启调试模式执行:

	$ qcli -d instance DescribeInstances

*注意: 命令行还在开发中, 欢迎参与完善!*

## 快速入门

以下为 [./hello.go](./hello.go) 的内容:

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/pbutil"
)

var (
	flagId   = flag.String("id", "", "AccessKeyId")
	flagKey  = flag.String("key", "", "SecretAccessKey")
	flagZone = flag.String("zone", "pek3a", "zone")
)

func main() {
	flag.Parse()

	// 返回 NIC 服务, pek3a 为 北京3区-A
	qnic := pb.NewNicService(&pb.ServerInfo{
		AccessKeyId:     proto.String(*flagId),
		SecretAccessKey: proto.String(*flagKey),
		Zone:            proto.String(*flagZone),
	})

	// 列出所有网卡
	reply, err := qnic.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 格式打印
	s, _ := pbutil.EncodeJsonIndent(reply)
	fmt.Println(s)
}
```

运行例子:

	go run hello.go -id=QYACCESSKEYIDEXAMPLE -key=SECRETACCESSKEY

其中 `-id` 和 `-key` 分别为 AccessKey 的公钥和私钥.

更完整的例子可以参考 [./pkg/cmd/qcli](./pkg/cmd/qcli) 的实现.

## 版权

The Apache License.
