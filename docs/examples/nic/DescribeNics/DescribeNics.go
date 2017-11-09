// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chai2010/qingcloud-go/pkg/config"
	pb "github.com/chai2010/qingcloud-go/pkg/service.pb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()

	conf := config.MustLoadUserConfig()
	conf.JSONDisableUnknownFields = false
	conf.LogLevel = "debug" // debug/warn

	qcService, err := pb.Init(conf)
	if err != nil {
		log.Fatal(err)
	}

	nicService, err := qcService.Nic("pek3a")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := nicService.DescribeNics(&pb.DescribeNicsInput{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonpbEncode(reply))
}

func jsonpbEncode(m proto.Message) string {
	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		Indent:       "  ",
		EnumsAsInts:  true,
		EmitDefaults: true,
	}

	s, err := jsonMarshaler.MarshalToString(m)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
