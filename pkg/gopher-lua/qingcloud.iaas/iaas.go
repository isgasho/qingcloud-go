// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_iaas

import (
	"fmt"
	"log"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
	"github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("qingcloud.iaas", Loader)
}

func Loader(L *lua.LState) int {
	mod := L.NewTable()

	initVersionTable(L, mod)
	initClientClass(L, mod)
	initClientServiceMethod(L, mod)

	L.SetFuncs(mod, _gosdk_Funcs)

	L.Push(mod)
	return 1
}

func initVersionTable(L *lua.LState, mod *lua.LTable) {
	L.SetField(mod, "version", func() *lua.LTable {
		tb := L.NewTable()
		L.SetField(tb, "version", lua.LString(verpkg.ShortVersion))
		L.SetField(tb, "short_version", lua.LString(verpkg.ShortVersion))
		L.SetField(tb, "git_sha1_version", lua.LString(verpkg.GitSha1Version))
		L.SetField(tb, "build_date", lua.LString(verpkg.BuildDate))
		return tb
	}())
}

func initClientClass(L *lua.LState, mod *lua.LTable) {
	const code = `
	Client = {
		api_server        = "",
		access_key_id     = "",
		secret_access_key = "",
		zone              = ""
	}

	function Client:new(obj)
		local obj = obj or {}
		setmetatable(obj, self)
		self.__index = self
		return obj
	end

	function Client:CallMethod()
		print("must to be replaced by Go function!!!")
	end

	function Client:sayhi()
		print("Client:sayhi")
	end

	return Client
	`

	if err := L.DoString(code); err != nil {
		log.Fatal(err)
	}

	L.SetField(mod, "Client", L.ToTable(-1))
}

func initClientServiceMethod(L *lua.LState, mod *lua.LTable) {
	for _, v := range pb.ServiceApiSpecMap {
		if v.ActionName == "DescribeInstances" {
			fmt.Println("action:", v)
		}
		code := fmt.Sprintf(`
			return function()
				return "todo"
				-- _gosdk_call_method(self, "%s", request, "%s")
			end`,
			v.ActionName,
			v.HttpMethod,
		)

		if err := L.DoString(code); err != nil {
			log.Fatal(err)
		}

		L.SetField(
			L.GetField(mod, "Client"), v.ActionName,
			L.ToFunction(-1),
		)
	}
}

var _gosdk_Funcs = map[string]lua.LGFunction{
	"_gosdk_call_method": gosdkCallMethod,
}

func gosdkCallMethod(L *lua.LState) int {
	L.Push(lua.LString("gosdkCallMethod"))
	return 1
}
