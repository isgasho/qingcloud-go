// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli_pb

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

var AllCommands []cli.Command

func pkgGetServerInfo(c *cli.Context) (p *pb.ServerInfo) {
	if c.IsSet("config") {
		p = pkgMustLoadConfig(c.GlobalString("config"))
	}
	if c.IsSet("api_server") {
		p.ApiServer = proto.String(c.GlobalString("api_server"))
	}
	if c.IsSet("access_key_id") {
		p.AccessKeyId = proto.String(c.GlobalString("access_key_id"))
	}
	if c.IsSet("secret_access_key") {
		p.SecretAccessKey = proto.String(c.GlobalString("secret_access_key"))
	}
	if c.IsSet("zone") {
		p.Zone = proto.String(c.GlobalString("zone"))
	}
	return
}

func pkgMustLoadConfig(path string) *pb.ServerInfo {
	p := new(pb.ServerInfo)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
