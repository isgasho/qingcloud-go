// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	_ "github.com/chai2010/qingcloud-go/pkg/internal/glogger"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "qctl"
	app.Usage = "Query or send control commands to the QingCloud"
	app.Version = verpkg.Version

	app.Authors = []cli.Author{
		{Name: "ChaiShushan", Email: "chaishushan@gmail.com"},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "ERR: command %q not found!\n", command)
	}

	app.Commands = []cli.Command{
		cmdInstance,
		cmdCluster,
		cmdVolume,
		cmdNic,
		cmdVxnet,
		cmdRouter,
		cmdEip,
		cmdSecurityGroup,
		cmdKeyPair,
		cmdImage,
		cmdLoadBalancer,
		cmdMonitor,
		cmdSnapshot,
		cmdJob,
	}

	app.Run(os.Args)
}
