// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/yuin/gopher-lua"

	inspect "github.com/chai2010/qingcloud-go/pkg/gopher-lua/inspect"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	inspect.Preload(L)

	if err := L.DoString(code); err != nil {
		panic(err)
	}
}

const code = `
local inspect = require("inspect")

print("https://github.com/kikito/inspect.lua")

print(inspect(1))
print(inspect("Hello"))

print(inspect({1,2,3,4}))
print(inspect({a=1,b=2}))

print(inspect({1,2,3,b=2,a=1}))
print(inspect({a={b=2}}))

local a = {1, 2}
local b = {3, 4, a}
a[3] = b -- a references b, and b references a
print(inspect(a))
`
