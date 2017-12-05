// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/go-validator"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/golang"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/generator/qcli"
)

func main() {
	plugin.Main()
}
