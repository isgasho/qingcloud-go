// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/yuin/gopher-lua"

	lua_socket "github.com/BixData/gluasocket"
	lua_http "github.com/cjoudrey/gluahttp"
	lua_url "github.com/cjoudrey/gluaurl"

	lua_json "github.com/chai2010/glua-json"
	lua_lfs "github.com/chai2010/glua-lfs"
	lua_getopt "github.com/chai2010/qingcloud-go/pkg/gopher-lua/getopt"
	lua_inspect "github.com/chai2010/qingcloud-go/pkg/gopher-lua/inspect"
	lua_lake "github.com/chai2010/qingcloud-go/pkg/gopher-lua/lake"
	lua_lustache "github.com/chai2010/qingcloud-go/pkg/gopher-lua/lustache"
	lua_qc_iaas "github.com/chai2010/qingcloud-go/pkg/gopher-lua/qingcloud.iaas"
	lua_strings "github.com/chai2010/qingcloud-go/pkg/gopher-lua/strings"
)

func init() {
	if os.PathSeparator == '/' { // unix-like
		lua.LuaPathDefault = "./?.lua;"
		lua.LuaPathDefault += pkgGetHomePath() + "/.qingcloud/lua/?.lua;"
		lua.LuaPathDefault += lua.LuaLDir + "/?.lua;"
		lua.LuaPathDefault += lua.LuaLDir + "/?/init.lua"
	} else { // windows
		lua.LuaPathDefault = ".\\?.lua;"
		lua.LuaPathDefault += pkgGetHomePath() + "\\.qingcloud\\lua\\?.lua;"
		lua.LuaPathDefault += lua.LuaLDir + "\\?.lua;"
		lua.LuaPathDefault += lua.LuaLDir + "\\?\\init.lua"
	}
}

func luaPreload(L *lua.LState) {
	lua_strings.Preload(L)

	lua_json.Preload(L)
	lua_lfs.Preload(L)
	lua_socket.Preload(L)

	L.PreloadModule("http", lua_http.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("url", lua_url.Loader)

	lua_lake.Preload(L)
	lua_lustache.Preload(L)
	lua_inspect.Preload(L)
	lua_getopt.Preload(L)

	lua_qc_iaas.Preload(L)
}

func luaDostring(ls *lua.LState, name, source string) error {
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

func pkgGetHomePath() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	if home == "" {
		home = "~"
	}

	return home
}
