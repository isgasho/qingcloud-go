// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/yuin/gopher-lua"
)

var (
	luaErrFunction = errors.New("lua: cannot encode function to JSON")
	luaErrChannel  = errors.New("lua: cannot encode channel to JSON")
	luaErrState    = errors.New("lua: cannot encode state to JSON")
	luaErrUserData = errors.New("lua: cannot encode userdata to JSON")
	luaErrNested   = errors.New("lua: cannot encode recursively nested tables to JSON")
)

type luaJsonValue struct {
	lua.LValue
	visited map[*lua.LTable]bool
}

func (p luaJsonValue) MarshalJSON() ([]byte, error) {
	return luaToJSON(p.LValue, p.visited)
}

func luaToJSON(value lua.LValue, visited map[*lua.LTable]bool) (data []byte, err error) {
	switch converted := value.(type) {
	case lua.LBool:
		data, err = json.Marshal(converted)
	case lua.LChannel:
		err = luaErrChannel
	case lua.LNumber:
		data, err = json.Marshal(converted)
	case *lua.LFunction:
		err = luaErrFunction
	case *lua.LNilType:
		data, err = json.Marshal(converted)
	case *lua.LState:
		err = luaErrState
	case lua.LString:
		data, err = json.Marshal(converted)
	case *lua.LTable:
		var arr []luaJsonValue
		var obj map[string]luaJsonValue

		if visited[converted] {
			panic(luaErrNested)
		}
		visited[converted] = true

		converted.ForEach(func(k lua.LValue, v lua.LValue) {
			i, numberKey := k.(lua.LNumber)
			if numberKey && obj == nil {
				index := int(i) - 1
				if index != len(arr) {
					// map out of order; convert to map
					obj = make(map[string]luaJsonValue)
					for i, value := range arr {
						obj[strconv.Itoa(i+1)] = value
					}
					obj[strconv.Itoa(index+1)] = luaJsonValue{v, visited}
					return
				}
				arr = append(arr, luaJsonValue{v, visited})
				return
			}
			if obj == nil {
				obj = make(map[string]luaJsonValue)
				for i, value := range arr {
					obj[strconv.Itoa(i+1)] = value
				}
			}
			obj[k.String()] = luaJsonValue{v, visited}
		})
		if obj != nil {
			data, err = json.Marshal(obj)
		} else {
			data, err = json.Marshal(arr)
		}
	case *lua.LUserData:
		// TODO: call metatable __tostring?
		err = luaErrUserData
	}
	return
}

func luaFromJSON(L *lua.LState, value interface{}) lua.LValue {
	switch converted := value.(type) {
	case bool:
		return lua.LBool(converted)
	case float64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case []interface{}:
		arr := L.CreateTable(len(converted), 0)
		for _, item := range converted {
			arr.Append(luaFromJSON(L, item))
		}
		return arr
	case map[string]interface{}:
		tbl := L.CreateTable(0, len(converted))
		for key, item := range converted {
			tbl.RawSetH(lua.LString(key), luaFromJSON(L, item))
		}
		return tbl
	}
	return lua.LNil
}
