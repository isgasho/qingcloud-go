// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/chai2010/qingcloud-go/pkg/client"
	pb "github.com/chai2010/qingcloud-go/pkg/cmd/qcli/api"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "qcli"
	app.Usage = "QingCloud Command Line Interface"
	app.Version = verpkg.ShortVersion

	app.Authors = []cli.Author{
		{Name: "ChaiShushan", Email: "chaishushan@gmail.com"},
	}

	app.Before = func(c *cli.Context) error {
		if c.GlobalIsSet("debug") {
			client.DebugMode = c.GlobalBool("debug")
			pb.DebugMode = c.GlobalBool("debug")
		}
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "ERR: command %q not found!\n", command)
	}

	app.Flags = pkgGlobalFlags
	app.Commands = pb.AllCommands
	app.EnableBashCompletion = true

	app.Run(os.Args)
}

var pkgGlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Usage:  "config file",
		Value:  "~/.qingcloud/qcli.json",
		EnvVar: "QCLI_CONFIG_FILE",
	},
	cli.StringFlag{
		Name:   "api_server, s",
		Usage:  "api server",
		Value:  "https://api.qingcloud.com/iaas/",
		EnvVar: "QCLI_API_SERVER",
	},
	cli.StringFlag{
		Name:   "access_key_id, i",
		Usage:  "access key id",
		Value:  "",
		EnvVar: "QCLI_ACCESS_KEY_ID",
	},
	cli.StringFlag{
		Name:   "secret_access_key, k",
		Usage:  "secret access key",
		Value:  "",
		EnvVar: "QCLI_SECRET_ACCESS_KEY",
	},
	cli.StringFlag{
		Name:   "zone, z",
		Usage:  "zone (pek3a,pek3b,gd1,sh1a,ap1,ap2a,...)",
		Value:  "pek3a",
		EnvVar: "QCLI_ZONE",
	},
	cli.BoolFlag{
		Name:   "debug, d",
		Usage:  "set debug mode",
		EnvVar: "QCLI_DEBUG",
	},
}
