// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/yuin/gopher-lua"

	lustache "github.com/chai2010/qingcloud-go/pkg/gopher-lua/lustache"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	lustache.Preload(L)

	if err := L.DoString(code); err != nil {
		panic(err)
	}
}

const code = `
local lustache = require("lustache")

print("https://github.com/Olivine-Labs/lustache")

print("----------")

view_model = {
	title = "Joe",
	calc = function ()
		return 2 + 4
	end
}

output = lustache:render("{{title}} spends {{calc}}", view_model)
print(output)

print("----------")

local tmpl = [[
* {{name}}
* {{age}}
* {{company}}
* {{{company}}}
* {{&company}}
]]

local model = {
	name = "Chris",
	company = "<b>GitHub</b>"
}

print(lustache:render(tmpl, model))

print("----------")

local tmpl = [[
* {{name.first}} {{name.last}}
* {{age}}
]]

local model = {
	name = {
		first = "Michael",
		last = "Jackson",
	},
	age = "RIP",
}

print(lustache:render(tmpl, model))
`
