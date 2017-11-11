// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-go"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-go/generator/go.v1"
)

func main() {
	plugin.Main()
}
