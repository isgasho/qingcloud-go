// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// 青云SDK Protobuf接口规范插件.
package main

import (
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/go-server"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/go-validator"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/qcli"
)

func main() {
	plugin.Main()
}
