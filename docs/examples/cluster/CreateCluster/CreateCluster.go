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

	clusterService := pb.NewClusterService(conf, "pek3a")

	reply, err := clusterService.CreateCluster(&pb.CreateClusterInput{
		Conf: proto.String(confCluster),
	})
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

const confCluster = `
{
	"app_id": "app-tg3lbp0a",
	"app_version": "1.0",
	"vxnet": "vxnet-etjv01u",
	"node": {
		"instance_class": 0,
		"count": 3,
		"cpu": 1,
		"memory": 512,
		"server_id_upper_bound":255
	}
}
`
