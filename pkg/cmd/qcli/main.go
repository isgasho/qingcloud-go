// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"flag"
	"fmt"
	"os"

	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/cmd/qcli/api"
	_ "github.com/chai2010/qingcloud-go/pkg/internal/glogger"
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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "config file",
			Value:  "~/.qingcloud/config.yaml",
			EnvVar: "QCLI_CONFIG_FILE",
		},
		cli.StringFlag{
			Name:   "zone, z",
			Usage:  "zone (pk3a,pk3b,gd1,sh1a,ap1,ap2a,...)",
			Value:  "pk3a",
			EnvVar: "QCLI_ZONE",
		},
		cli.StringFlag{
			Name:   "glog_level",
			Usage:  "glog level to stderr (INFO,WARNING,ERROR,FATAL)",
			Value:  "WARNING",
			EnvVar: "QCLI_GLOG_LEVEL",
		},
	}

	app.Before = func(c *cli.Context) error {
		flag.Parse()
		if c.IsSet("glog_level") {
			flag.CommandLine.Set("stderrthreshold", c.String("glog_level"))
		}
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "ERR: command %q not found!\n", command)
	}

	app.EnableBashCompletion = true
	app.Commands = pb.AllCommands
	app.Run(os.Args)
}
