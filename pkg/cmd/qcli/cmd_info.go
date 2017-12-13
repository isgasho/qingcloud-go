// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdInfo = cli.Command{
	Name:    "info",
	Aliases: []string{},
	Usage:   "show resource info",
	Hidden:  true,
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
