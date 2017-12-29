// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/yuin/gopher-lua"

	getopt "github.com/chai2010/glua-getopt"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	getopt.Preload(L)

	if err := L.DoString(code); err != nil {
		panic(err)
	}
}

const code = `
local getopt = require("getopt")

local arg = {"-aone", "-btwo", "-c=three", "-d", "--file=foo.lua", "abc", "123"}
local opts, arg = getopt( arg, "abc" )

for k, v in pairs(opts) do
  print( k, v )
end

for k, v in pairs(arg) do
  print( k, v )
end

--[[
output:
	a       one
	b       two
	c       =three
	d       true
	file    foo.lua
	1       abc
	2       123
--]]
`
