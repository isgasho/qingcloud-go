// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

// 测试 proto 数据转转 request
package main

import (
	"encoding/json"
	"fmt"
)

// 以前成员全部是定义为指针类型
// proto3 全部定义为值类型
// request/builder.go 只实现了指针类型的处理

// 思路:
// 1. proto 转 json
// 2. json 解码为 map[string]interface{}
// 3. 遍历 map

// 只需要支持 int/string 和对应的数组

// request/builder.go
// func (b *Builder) parseRequestParams() error {
// var requestParams map[string]string

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
}
