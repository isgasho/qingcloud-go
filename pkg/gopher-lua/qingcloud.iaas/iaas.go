// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_iaas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/client"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
	"github.com/yuin/gopher-lua"
)

var _ = fmt.Sprint

func Preload(L *lua.LState) {
	L.PreloadModule("qingcloud.iaas", Loader)
}

func Loader(L *lua.LState) int {
	mod := L.NewTable()

	initCopyright(L, mod)

	initVersionTable(L, mod)
	initServiceMethodTable(L, mod)

	initClientClass(L, mod)
	initClientServiceMethod(L, mod)

	L.SetFuncs(mod, _gosdk_Funcs)

	L.Push(mod)
	return 1
}

var _gosdk_Funcs = map[string]lua.LGFunction{
	"DecodeJSON": apiDecodeJSON,
	"EncodeJSON": apiEncodeJSON,
	"LoadJSON":   apiLoadJSON,
}

func initCopyright(L *lua.LState, mod *lua.LTable) {
	const s = `Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
Use of this source code is governed by a Apache
license that can be found in the LICENSE file.
`
	L.SetField(mod, "copyright", lua.LString(s))
}

func initVersionTable(L *lua.LState, mod *lua.LTable) {
	L.SetField(mod, "version", lua.LString(verpkg.ShortVersion))

	L.SetField(mod, "version_info", func() *lua.LTable {
		tb := L.NewTable()
		L.SetField(tb, "version", lua.LString(verpkg.ShortVersion))
		L.SetField(tb, "short_version", lua.LString(verpkg.ShortVersion))
		L.SetField(tb, "git_sha1_version", lua.LString(verpkg.GitSha1Version))
		L.SetField(tb, "build_date", lua.LString(verpkg.BuildDate))
		return tb
	}())
}

func initServiceMethodTable(L *lua.LState, mod *lua.LTable) {
	L.SetField(mod, "method_list", func() *lua.LTable {
		tb := L.NewTable()

		var keys = make([]string, 0, len(pb.ServiceApiSpecMap))
		for k := range pb.ServiceApiSpecMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := pb.ServiceApiSpecMap[k]

			info := L.NewTable()
			L.SetField(info, "action_name", lua.LString(v.ActionName))
			L.SetField(info, "http_method", lua.LString(v.HttpMethod))
			tb.Append(info)
		}

		return tb
	}())
}

func initClientClass(L *lua.LState, mod *lua.LTable) {
	code := fmt.Sprintf(`
		Client = {
			api_server        = %[1]q,
			access_key_id     = %[2]q,
			secret_access_key = %[3]q,
			zone              = %[4]q
		}

		function Client:new(obj)
			local p = {}
			setmetatable(p, self)
			self.__index = self

			obj = obj or {}
			for k, v in pairs(obj) do
				p[k] = v
			end

			if p.api_server == nil or p.api_server == '' then
				p.api_server = %[1]q
			end
			if p.access_key_id == nil or p.access_key_id == '' then
				p.access_key_id = %[2]q
			end
			if p.secret_access_key == nil or p.secret_access_key == '' then
				p.secret_access_key = %[3]q
			end
			if p.zone == nil or p.zone == '' then
				p.zone = %[4]q
			end

			return p
		end

		return Client
	`,
		(*pb.ServerInfo)(nil).GetApiServer(),
		(*pb.ServerInfo)(nil).GetAccessKeyId(),
		(*pb.ServerInfo)(nil).GetSecretAccessKey(),
		(*pb.ServerInfo)(nil).GetZone(),
	)

	if err := L.DoString(code); err != nil {
		log.Fatal(err)
	}

	L.SetField(mod, "Client", L.ToTable(-1))
}

func initClientServiceMethod(L *lua.LState, mod *lua.LTable) {
	self := L.GetField(mod, "Client")
	for _, v := range pb.ServiceApiSpecMap {
		L.SetField(
			self, v.ActionName,
			L.NewClosure(gosdkCallMethod,
				lua.LString(v.HttpMethod),
				lua.LString(v.ActionName),
			),
		)
	}
}

// closure
// upvalue: http_method, action
// function arguments: server_info, request
// return: reply[, error]
func gosdkCallMethod(L *lua.LState) int {
	http_method := L.CheckString(lua.UpvalueIndex(1))
	action := L.CheckString(lua.UpvalueIndex(2))

	server_info := L.CheckTable(1)
	request := L.CheckTable(2)

	client := client.NewClient(
		L.GetField(server_info, "api_server").(lua.LString).String(),
		L.GetField(server_info, "access_key_id").(lua.LString).String(),
		L.GetField(server_info, "secret_access_key").(lua.LString).String(),
		L.GetField(server_info, "zone").(lua.LString).String(),
	)

	visited := make(map[*lua.LTable]bool)
	reqJsonData, err := luaToJSON(request, visited)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	output, err := client.CallMethodWithJson(
		action, http_method, string(reqJsonData),
		nil,
	)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var value interface{}
	err = json.Unmarshal([]byte(output), &value)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(luaFromJSON(L, value))
	return 1
}

func apiDecodeJSON(L *lua.LState) int {
	str := L.CheckString(1)

	var value interface{}
	err := json.Unmarshal([]byte(str), &value)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(luaFromJSON(L, value))
	return 1
}

func apiEncodeJSON(L *lua.LState) int {
	value := L.CheckAny(1)

	visited := make(map[*lua.LTable]bool)
	data, err := luaToJSON(value, visited)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(string(data)))
	return 1
}

func apiLoadJSON(L *lua.LState) int {
	cfgpath := L.CheckString(1)
	if strings.HasPrefix(cfgpath, "~/") || strings.HasPrefix(cfgpath, `~\`) {
		cfgpath = pkgGetHomePath() + cfgpath[1:]
	}

	data, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var value interface{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(luaFromJSON(L, value))
	return 1
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
