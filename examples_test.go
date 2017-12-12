// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/pbutil"
	sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"
)

func Example_signatureBuild() {
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
	// TODO
}

func Example_helloSDK() {
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
