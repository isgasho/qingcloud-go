<p align="center"><a href="http://qingcloud.com" target="_blank"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/docs/images/logo-2018-01-490x130.png" alt="QingCloud"></a></p>

# 青云非官方 SDK Go/[GopherLua](https://github.com/yuin/gopher-lua)

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg?branch=master)](https://travis-ci.org/chai2010/qingcloud-go)
[![Docker Build Status](https://img.shields.io/docker/build/chai2010/qingcloud-go.svg)](https://hub.docker.com/r/chai2010/qingcloud-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/qingcloud-go)](https://goreportcard.com/report/github.com/chai2010/qingcloud-go)
[![codebeat badge](https://codebeat.co/badges/4dd71afb-7bde-4073-9c8c-cbe6275ba6eb)](https://codebeat.co/projects/github-com-chai2010-qingcloud-go-master)
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

接口规范:

- [api](api)

## qingcloud-go vs [Office SDK API](https://github.com/yunify/qingcloud-sdk-go)

[QingCloud Service](https://docs.qingcloud.com/api/common/action.html)            | [chai2010/qingcloud-go](https://github.com/chai2010/qingcloud-go) | [yunify/qingcloud-sdk-go](https://github.com/yunify/qingcloud-sdk-go)
:----------------- | :-------------------- | :---------------------
Alarm              | ✓                     | ✗
Cache              | ✓                     | ✓
Cluster            | ✓                     | ✓
DNSAlias           | ✓                     | ✓
EIP                | ✓                     | ✓
Hadoop             | ✓                     | ✗
Image              | ✓                     | ✓
Instance           | ✓                     | ✓
Job                | ✓                     | ✓
KeyPair            | ✓                     | ✓
LoadBalancer       | ✓                     | ✓
Misc               | ✓                     | ✗
Mongo              | ✓                     | ✓
Monitor            | ✓                     | ✓
Nic                | ✓                     | ✓
NotificationCenter | ✓                     | ✗
RDB                | ✓                     | ✓
ResourceACL        | ✓                     | ✗
Router             | ✓                     | ✓
S2                 | ✓                     | ✗
SecurityGroup      | ✓                     | ✓
Snapshot           | ✓                     | ✓
Span               | ✓                     | ✗
Spark              | ✓                     | ✗
Subuser            | ✓                     | ✗
Tag                | ✓                     | ✓
UserData           | ✓                     | ✓
Volumes            | ✓                     | ✓
Vxnet              | ✓                     | ✓
Zone               | ✓                     | ✓

## qcli 命令行

Docker 运行([配置中国区镜像](https://www.docker-cn.com/registry-mirror)):

- `docker run --rm -it -v $HOME:/root -w /root chai2010/qingcloud-go qcli`

从Go源码安装(Go1.11+):

- `go get github.com/chai2010/qingcloud-go/cmd/qcli`

输入 `qcli` 或 `qcli -h` 查看命令提示:

```
chai-mba:api chai$ qcli
NAME:
   qcli - QingCloud Command Line Interface

USAGE:
   qcli [global options] command [command options] [arguments...]

VERSION:
   v0.9.0-33-gbf90580

AUTHOR:
   ChaiShushan <chaishushan@gmail.com>

COMMANDS:
     lake, make  build target with lakefile.lua
     help, h     Shows a list of commands or help for one command

   SDK API Style Command:
     alarm               AlarmService
     cache               CacheService
     cluster             ClusterService
     dnsalias            DNSAliasService
     eip                 EIPService
     hadoop              HadoopService
     image               ImageService
     instance            InstanceService
     job                 JobService
     keypair             KeyPairService
     loadbalancer        LoadBalancerService
     misc                MiscService
     mongo               MongoService
     monitor             MonitorService
     nic                 NicService
     notificationcenter  NotificationCenterService
     rdb                 RDBService
     resourceacl         ResourceACLService
     router              RouterService
     s2                  S2Service
     securitygroup       SecurityGroupService
     snapshot            SnapshotService
     span                SpanService
     spark               SparkService
     subuser             SubuserService
     tag                 TagService
     userdata            UserDataService
     volumes             VolumesService
     vxnet               VxnetService
     zone                ZoneService

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
	"zone": "pek3a"
}
```

qcli 内置了一个类似 make 到构建系统, 输入以下命令查看 ([lakefile.lua](lakefile.lua)):

	$ qcli lake -h
	$ qcli lake -t
	$ qcli lake list.instance

要查看主机数量, 可以输入以下命令:

	$ qcli instance DescribeInstances

加入 `-d` 选项可以开启调试模式执行:

	$ qcli -d instance DescribeInstances

*注意: 命令行还在开发中, 欢迎参与完善!*

## 快速入门 (Go语言版本)

以下为 [./hello.go](./hello.go) 的内容:

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
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

	fmt.Println(reply)
}
```

运行例子:

	go run hello.go -id=QYACCESSKEYIDEXAMPLE -key=SECRETACCESSKEY

其中 `-id` 和 `-key` 分别为 AccessKey 的公钥和私钥.

更完整的例子可以参考 [./pkg/cmd/qcli](./pkg/cmd/qcli) 的实现.


## 快速入门([GopherLua](https://github.com/yuin/gopher-lua) 版本)

[hello-lua.go](hello-lua.go):

```go
package main

import (
	"github.com/yuin/gopher-lua"

	qc_iaas "github.com/chai2010/qingcloud-go/pkg/gopher-lua/qingcloud.iaas"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	qc_iaas.Preload(L)

	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}
```

[hello.lua](hello.lua):

```lua
local qc = require("qingcloud.iaas")

if #arg == 1 and arg[1] == '-v' then
	print(qc.version)
	print(qc.version_info.git_sha1_version)
	print(qc.version_info.build_date)
	do return end
end

if #arg == 1 and arg[1] == '-h' then
	print(qc.copyright)
	print("hello, 青云!")
	do return end
end

local config = qc.LoadJSON("~/.qingcloud/qcli.json")
local client = qc.Client:new(config)

local reply, err = client:DescribeInstances {
	--owner = "usr-xxxxxxxx",
	zone = "pek3a",
	limit = 100
}
if err ~= nil then
	print("error:", err)
	do return end
end

if reply.ret_code ~= 0 then
	print(reply.ret_code)
	print(reply.message)
	do return end
end

for i = 1, #reply.instance_set do
	local item = reply.instance_set[i]
	print(i,
		item.instance_id,
		item.instance_type,
		item.memory_current..'MB',
		item.status,
		item.create_time,
		item.instance_name
	)
end

print('total: ' .. reply.total_count)
```

或者通过 `qlua` 命令行解释执行:

	$ go install github.com/chai2010/qingcloud-go/cmd/qlua
	$ qlua hello.lua -h
	$ qlua hello.lua -v
	$ qlua hello.lua


## 快速入门 (qcli lake 版本)

[lakefile.lua](lakefile.lua):

```lua
task("default", {"version", "doc"}, function()
	print("done")
end)

task("version", nil, function()
	print("version")
end)

task("doc", {"install"}, function()
	print("doc")
end)

task("install", nil, function(task, destdir)
	print("install:", task.name, destdir)
end)
```

输入以下命令:

	$ qcli lake -h
	$ qcli lake -t
	$ qcli lake
	$ qcli make install[dir=/path/to/dir]


输入以下命令生成 [Graphviz](http://www.graphviz.org/) 格式的依赖图:

	$ qcli make -g > lakefile-graph.dot
	$ qcli make -g | dot -Tpng > lakefile-graph.png

[lakefile-graph.dot](lakefile-graph.dot):

```dot
digraph G {
	default [label = "default"];
	version [label = "version"];
	doc [label = "doc"];
	install [label = "install"];
	list_instance [label = "list.instance"];

	default -> version;
	default -> doc;
	doc -> install;
}
```

[lakefile-graph.png](lakefile-graph.png):

![](lakefile-graph.png):


## 版权

The Apache License.
