// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate go run makestatic.go

package gopherlua_lustache

import (
	"log"
	"strings"

	"github.com/yuin/gopher-lua"
)

var pkgModFileNameMap = map[string]string{
	"lustache":          Files["lustache-1.3.1-0/src/lustache.lua"],
	"lustache.context":  Files["lustache-1.3.1-0/src/lustache/context.lua"],
	"lustache.renderer": Files["lustache-1.3.1-0/src/lustache/renderer.lua"],
	"lustache.scanner":  Files["lustache-1.3.1-0/src/lustache/scanner.lua"],
}

func Preload(L *lua.LState) {
	for modName, modFileName := range pkgModFileNameMap {
		modName, modFileName := modName, modFileName

		L.PreloadModule(modName, func(L *lua.LState) int {
			if err := dostring(L, modFileName, Files[modFileName]); err != nil {
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
