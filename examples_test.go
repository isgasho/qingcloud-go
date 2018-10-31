// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	clientpkg "github.com/chai2010/qingcloud-go/pkg/client"
	"github.com/chai2010/qingcloud-go/pkg/pbutil"
	sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"
	statuspkg "github.com/chai2010/qingcloud-go/pkg/status"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
	"github.com/chai2010/qingcloud-go/pkg/wait"
)

func Example_signatureBuild() {
	// import sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"

	query, signature := sigpkg.Build(
		"QYACCESSKEYIDEXAMPLE", "SECRETACCESSKEY",
		"GET", "/iaas/", map[string]string{},
	)

	fmt.Println(query)
	fmt.Println(signature)

	// Output:
	// access_key_id=QYACCESSKEYIDEXAMPLE&signature_method=HmacSHA256&signature_version=1&signature=O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D
	// O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D
}

func Example_signatureGetSignatureInfo() {
	// import sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"

	s := "access_key_id=QYACCESSKEYIDEXAMPLE&signature_method=HmacSHA256&signature_version=1&signature=O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D"
	pubKey, sigMethod, sigVersion, sig := sigpkg.GetSignatureInfo(s)

	fmt.Println(pubKey)
	fmt.Println(sigMethod)
	fmt.Println(sigVersion)
	fmt.Println(sig)

	// Output:
	// QYACCESSKEYIDEXAMPLE
	// HmacSHA256
	// 1
	// O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE=
}

func Example_signatureValidate() {
	// import sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"

	ok := sigpkg.Validate(
		func(pubKey string) string {
			if pubKey != "QYACCESSKEYIDEXAMPLE" {
				log.Fatal("invalid pubKey:", pubKey)
			}
			return "SECRETACCESSKEY"
		},
		"GET", "/iaas/",
		"access_key_id=QYACCESSKEYIDEXAMPLE&signature_method=HmacSHA256&signature_version=1&signature=O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D",
	)

	fmt.Println(ok)

	// Output:
	// true
}

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

	s, _ := pbutil.EncodeJsonIndent(reply)
	fmt.Println(s)
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

func Example_waitJob() {
	// import pb "github.com/chai2010/qingcloud-go/pkg/api"
	// import "github.com/chai2010/qingcloud-go/pkg/wait"

	qc := pb.NewInstanceService(&pb.ServerInfo{
		AccessKeyId:     proto.String("QYACCESSKEYIDEXAMPLE"),
		SecretAccessKey: proto.String("SECRETACCESSKEY"),
		Zone:            proto.String("pek3a"),
	})

	reply, err := qc.RunInstances(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = wait.WaitJob(qc.ServerInfo, reply.GetJobId(), time.Second*5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("job done")
}

func Example_waitInstance() {
	// import pb "github.com/chai2010/qingcloud-go/pkg/api"
	// import statuspkg "github.com/chai2010/qingcloud-go/pkg/status"
	// import "github.com/chai2010/qingcloud-go/pkg/wait"

	qc := pb.NewInstanceService(&pb.ServerInfo{
		AccessKeyId:     proto.String("QYACCESSKEYIDEXAMPLE"),
		SecretAccessKey: proto.String("SECRETACCESSKEY"),
		Zone:            proto.String("pek3a"),
	})

	reply, err := qc.RunInstances(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = wait.WaitInstanceStatus(
		qc.ServerInfo, reply.GetInstances()[0],
		statuspkg.InstanceStatus_Running,
		time.Second*5,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("instance running")
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
