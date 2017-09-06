// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

// 测试 proto 数据转转 request
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	pb "github.com/chai2010/qingcloud-go/spec.pb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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
	reqInput := &pb.DescribeNicsInput{
		Instances: []string{"Instances-1", "Instances-2"},
		Limit:     12,
		NicName:   "NicName-111",
		Nics:      []string{"Nics-111", "Nics-222"},
		Offset:    123,
		Status:    "ok",
		VxnetType: 2,
		Vxnets:    []string{"Vxnets-11", "Vxnets-22"},
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName: true,
		Indent:   "  ",
	}

	s, err := jsonMarshaler.MarshalToString(reqInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("json:", s)

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(s), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["limit"].(float64)
	fmt.Println("limit:", num)

	fmt.Println("nic_name:", dat["nic_name"].(string))

	switch dat["nic_name"].(type) {
	case string:
		fmt.Println("nic_name type:", "string")
	}

	fmt.Println("vxnets[0]:", dat["vxnets"].([]interface{})[0].(string))
	fmt.Println("vxnets[1]:", dat["vxnets"].([]interface{})[1].(string))

	mx, err := protoMessageToMap(reqInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mx: %#v\n", mx)
}

func main2() {
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

func protoMessageToMap(msg proto.Message) (map[string]string, error) {
	jsonMarshaler := &jsonpb.Marshaler{
		OrigName: true,
		Indent:   "",
	}

	jsonString, err := jsonMarshaler.MarshalToString(msg)
	if err != nil {
		return nil, err
	}

	var mapx map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &mapx); err != nil {
		return nil, err
	}

	var m = map[string]string{}
	for k, v := range mapx {
		switch v := v.(type) {
		case string:
			m[k] = v
		case float64:
			m[k] = fmt.Sprintf("%v", v)
		case []interface{}:
			for i := 0; i < len(v); i++ {
				ki := k + "." + strconv.Itoa(i+1)

				switch vi := v[i].(type) {
				case string:
					m[ki] = vi
				case float64:
					m[ki] = fmt.Sprintf("%v", vi)
				default:
					// unreachable
				}
			}
		default:
			// unreachable
		}
	}

	return m, nil
}
