// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/pbutil"
)

var (
	flagId   = flag.String("id", "", "AccessKeyId")
	flagKey  = flag.String("key", "", "SecretAccessKey")
	flagZone = flag.String("zone", "pek3a", "zone")
)

func main() {
	flag.Parse()

	// 返回 NIC 服务, pek3a 为 北京3区-A
	qnic := pb.NewNicService(&pb.ServerInfo{
		AccessKeyId:     proto.String(*flagId),
		SecretAccessKey: proto.String(*flagKey),
		Zone:            proto.String(*flagZone),
	})

	// 列出所有网卡
	reply, err := qnic.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 格式打印
	s, _ := pbutil.EncodeJsonIndent(reply)
	fmt.Println(s)
}
