// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate go run makestatic.go

package gopherlua_inspect

import (
	"log"

	"github.com/yuin/gopher-lua"
)

var pkgModCodeMap = map[string]string{
	"inspect": Files["inspect.lua-3.1.0/inspect.lua"],
}

func Preload(L *lua.LState) {
	for modName, modCode := range pkgModCodeMap {
		modName, modCode := modName, modCode

		L.PreloadModule(modName, func(L *lua.LState) int {
			if err := L.DoString(modCode); err != nil {
				log.Fatal(err)
			}

			if tb := L.ToTable(-1); tb != nil {
				L.Push(tb)
				return 1
			}
			return 0
		})
	}
}
