// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"encoding/json"

	"github.com/yuin/gopher-lua"

	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

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

func luaOpenSDK(L *lua.LState) int {
	sdkmod := L.RegisterModule("qingcloud", luaSDK_Funcs)
	for k, v := range luaSDK_values {
		L.SetField(sdkmod, k, v)
	}
	L.Push(sdkmod)
	return 1
}

var luaSDK_values = map[string]lua.LValue{
	"version":          lua.LString(verpkg.ShortVersion),
	"short_version":    lua.LString(verpkg.ShortVersion),
	"git_sha1_version": lua.LString(verpkg.GitSha1Version),
	"build_date":       lua.LString(verpkg.BuildDate),
}

var luaSDK_Funcs = map[string]lua.LGFunction{
	"decode_json": luaSDK_decodeJSON,
	"encode_json": luaSDK_encodeJSON,
	"call_method": luaSDK_CallMethod,
}

func luaSDK_CallMethod(L *lua.LState) int {
	L.Push(lua.LString("TODO"))
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
