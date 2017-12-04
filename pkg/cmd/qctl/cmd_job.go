// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdJob = cli.Command{
	Name:    "job",
	Aliases: []string{},
	Usage:   "manage jobs",
	Subcommands: []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"desc"},
			Usage:   "describe jobs",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
	},
}
