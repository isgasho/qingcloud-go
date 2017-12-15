// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
	"github.com/yuin/gopher-lua"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/client"
	pbCli "github.com/chai2010/qingcloud-go/pkg/cmd/qcli/api"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

var _ = fmt.Sprint

func luaGetStringField(L *lua.LState, tb *lua.LTable, key string, v string) string {
	ret := tb.RawGetString(key)
	if ln, ok := ret.(lua.LString); ok {
		return string(ln)
	}
	return v
}
func luaGetIntField(L *lua.LState, tb *lua.LTable, key string, v int) int {
	ret := tb.RawGetString(key)
	if ln, ok := ret.(lua.LNumber); ok {
		return int(ln)
	}
	return v
}

func luaGetBoolField(L *lua.LState, tb *lua.LTable, key string, v bool) bool {
	ret := tb.RawGetString(key)
	if lb, ok := ret.(lua.LBool); ok {
		return bool(lb)
	}
	return v
}

func luaOpenSDK(c *cli.Context, L *lua.LState) int {
	sdkmod := L.RegisterModule("qingcloud", luaSDK_Funcs)

	// version
	L.SetField(sdkmod, "version", lua.LString(verpkg.ShortVersion))
	L.SetField(sdkmod, "short_version", lua.LString(verpkg.ShortVersion))
	L.SetField(sdkmod, "git_sha1_version", lua.LString(verpkg.GitSha1Version))
	L.SetField(sdkmod, "build_date", lua.LString(verpkg.BuildDate))

	// server_info
	L.SetField(sdkmod, "server_info", func() *lua.LTable {
		tb := L.NewTable()
		info := pbCli.GetServerInfo(c)
		L.SetField(tb, "api_server", lua.LString(info.GetApiServer()))
		L.SetField(tb, "access_key_id", lua.LString(info.GetAccessKeyId()))
		L.SetField(tb, "secret_access_key", lua.LString(info.GetSecretAccessKey()))
		L.SetField(tb, "zone", lua.LString(info.GetZone()))
		return tb
	}())

	// all_method_list
	L.SetField(sdkmod, "all_method_list", func() (tb *lua.LTable) {
		tb = L.NewTable()
		for k, v := range pb.ServiceApiSpecMap {
			action := L.NewTable()
			L.SetField(tb, k, action)

			L.SetField(action, "name", lua.LString(v.ActionName))
			L.SetField(action, "method", lua.LString(v.HttpMethod))
		}
		return
	}())

	L.Push(sdkmod)
	return 1
}

var luaSDK_Funcs = map[string]lua.LGFunction{
	"call_method": luaSDK_CallMethod,
	"decode_json": luaSDK_decodeJSON,
	"encode_json": luaSDK_encodeJSON,
}

// call_method({server_info, action, request}): response[, error]
func luaSDK_CallMethod(L *lua.LState) int {
	call_info := L.CheckTable(1)

	server_info := L.GetField(call_info, "server_info")
	action := L.GetField(call_info, "action")
	request := L.GetField(call_info, "request")

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
		action.(lua.LString).String(), "GET", string(reqJsonData),
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

func luaSDK_decodeJSON(L *lua.LState) int {
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

func luaSDK_encodeJSON(L *lua.LState) int {
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
