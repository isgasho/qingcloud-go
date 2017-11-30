// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-validate"
	_ "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-validate/generator/golang"
)

func main() {
	plugin.Main()
}
