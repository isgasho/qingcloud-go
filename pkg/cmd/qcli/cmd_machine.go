// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdMachine = cli.Command{
	Name:    "machine",
	Aliases: []string{},
	Usage:   "docker machine style command",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "regexp",
			Usage: "regexp filter",
			Value: ".*",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Println("TODO")
		return nil
	},
}
