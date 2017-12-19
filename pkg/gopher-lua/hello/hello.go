// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package gopherlua_hello

import (
	"log"
	"strings"

	"github.com/yuin/gopher-lua"
)

var pkgModName = "hello"

var pkgModCode = `
module("{{.ModName}}", package.seeall)

function sayHello()
	print("hello")
end
`

func Preload(L *lua.LState) {
	code := strings.Replace(pkgModCode, "{{.ModName}}", pkgModName, -1)

	L.PreloadModule(pkgModName, func(L *lua.LState) int {
		if err := L.DoString(code); err != nil {
			log.Fatal(err)
		}

		if tb := L.ToTable(-1); tb != nil {
			L.Push(tb)
			return 1
		}
		return 0
	})
}
