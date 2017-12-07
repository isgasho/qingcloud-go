// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()

	conf := config.MustLoadUserConfig()
	conf.JSONDisableUnknownFields = false
	conf.LogLevel = "debug" // debug/warn

	nicService := pb.NewNicService(conf, "pek3a")

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
