// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli_pb

import (
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

var AllCommands []cli.Command

func pkgGetServerInfo() *pb.ServerInfo {
	return &pb.ServerInfo{}
}
