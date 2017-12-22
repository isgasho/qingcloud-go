// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/yuin/gopher-lua"

	strings "github.com/chai2010/qingcloud-go/pkg/gopher-lua/strings"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	strings.Preload(L)

	if err := L.DoString(code); err != nil {
		panic(err)
	}
}

const code = `
local strings = require("strings")

print(strings.ToUpper("abc"))

local ss = strings.Split("aa,b,,c", ",")
for i = 1, #ss do
	print(i, ss[i])
end
`
