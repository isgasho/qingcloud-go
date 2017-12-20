// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"os"
	"strings"

	"github.com/yuin/gopher-lua"

	lake "github.com/chai2010/qingcloud-go/pkg/gopher-lua/lake"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	lake.Preload(L)

	// define arg

	L.SetGlobal("arg", func() *lua.LTable {
		tb := L.NewTable()
		for i := 1; i < len(os.Args); i++ {
			tb.Append(lua.LString(os.Args[i]))
		}
		return tb
	}())

	if err := dostring(L, "main.lua", code); err != nil {
		log.Fatal(err)
	}
}

func dostring(ls *lua.LState, name, source string) error {
	if name == "" {
		name = "<string>"
	}

	fn, err := ls.Load(strings.NewReader(source), name)
	if err != nil {
		return err
	}

	ls.Push(fn)
	return ls.PCall(0, lua.MultRet, nil)
}

const code = `
-- https://github.com/sebnow/lake

local Lake = require("Lake")

local lake = Lake.Application:new()

-- The Lake "DSL"
-- These functions allow to interface with the build engine in a succinct way
function task(name, prerequisites, action)
	lake:defineTask(name, prerequisites, action)
end

lake:run(arg)
`
