// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
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
	"call_method": luaSDK_CallMethod,
}

func luaSDK_CallMethod(L *lua.LState) int {
	L.Push(lua.LString("TODO"))
	return 1
}
