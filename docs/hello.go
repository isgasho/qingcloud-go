// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/config"
	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func main() {
	// 返回 NIC 服务, pek3a 为 北京3区-A
	nicService := pb.NewNicService(config.MustLoadUserConfig(), "pek3a")

	// 列出所有网卡
	reply, err := nicService.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 原始返回的json数据
	// nicService.LastResponseBody

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
