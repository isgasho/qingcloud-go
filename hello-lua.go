// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

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
