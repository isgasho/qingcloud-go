// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/service"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// 初始化 青云 服务对象
	qcService, err := pb.Init(config.MustLoadUserConfig())
	if err != nil {
		log.Fatal(err)
	}

	// 返回 NIC 子服务, pek3a 为 北京3区-A
	nicService, err := qcService.Nic("pek3a")
	if err != nil {
		log.Fatal(err)
	}

	// 列出所以网卡
	reply, err := nicService.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 格式打印
	fmt.Println(jsonpbEncode(reply))
}

// pb转json, 采用原始名称, 不忽略空值
func jsonpbEncode(m proto.Message) string {
	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(m)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
