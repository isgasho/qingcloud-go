<p align="center"><a href="http://qingcloud.com"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/logo.jpg" alt="QingCloud"></a></p>

# 青云 SDK Go Version (兼容 [yunify/qingcloud-sdk-go](https://github.com/yunify/qingcloud-sdk-go))

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg)](https://travis-ci.org/chai2010/qingcloud-go)
[![GoDoc](https://godoc.org/github.com/chai2010/qingcloud-go?status.svg)](https://godoc.org/github.com/chai2010/qingcloud-go)
[![API Reference](http://img.shields.io/badge/api-reference-green.svg)](http://docs.qingcloud.com)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/qingcloud-go/blob/master/LICENSE)

在线文档:

- https://docs.qingcloud.com
- https://godoc.org/github.com/chai2010/qingcloud-go

接口规范:

- [spec.pb](spec.pb)

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

	"github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/service"
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

	// 列出所以网卡
	reply, err := nicService.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

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

运行例子:

	go run hello.go

[更多例子](examples).
