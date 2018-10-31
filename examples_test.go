// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	clientpkg "github.com/chai2010/qingcloud-go/pkg/client"
	statuspkg "github.com/chai2010/qingcloud-go/pkg/status"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

func Example_helloSDK() {
	// import pb "github.com/chai2010/qingcloud-go/pkg/api"

	qnic := pb.NewNicService(&pb.ServerInfo{
		AccessKeyId:     proto.String("QYACCESSKEYIDEXAMPLE"),
		SecretAccessKey: proto.String("SECRETACCESSKEY"),
		Zone:            proto.String("pek3a"),
	})

	reply, err := qnic.DescribeNics(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

func Example_clientCallMethod() {
	// import pb "github.com/chai2010/qingcloud-go/pkg/api"
	// import clientpkg "github.com/chai2010/qingcloud-go/pkg/client"

	serverInfo := &pb.ServerInfo{}
	client := clientpkg.NewClient(
		serverInfo.GetApiServer(),
		serverInfo.GetAccessKeyId(),
		serverInfo.GetSecretAccessKey(),
		serverInfo.GetZone(),
	)

	input := &pb.DescribeAlarmPoliciesInput{}
	output := &pb.DescribeAlarmPoliciesOutput{}

	err := client.CallMethod("DescribeAlarmPolicies", "GET", input, output, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}

func Example_status() {
	// import pb "github.com/chai2010/qingcloud-go/pkg/api"
	// import statuspkg "github.com/chai2010/qingcloud-go/pkg/status"

	reply := &pb.DescribeInstancesOutput{}
	err := statuspkg.Error(reply)

	fmt.Println(err)

	// Output:
	// <nil>
}

func Example_version() {
	// import verpkg "github.com/chai2010/qingcloud-go/pkg/version"

	fmt.Println(verpkg.GetVersionString())
}
