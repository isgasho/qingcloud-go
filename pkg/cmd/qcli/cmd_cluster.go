// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdCluster = cli.Command{
	Name:    "cluster",
	Aliases: []string{},
	Usage:   "manage cluster",
	Subcommands: []cli.Command{
		{
			Name:    "todo",
			Aliases: []string{},
			Usage:   "cluster todo ...",
			Action: func(c *cli.Context) error {
				fmt.Println("todo: ", c.Args().First())
				return nil
			},
		},
	},
}
