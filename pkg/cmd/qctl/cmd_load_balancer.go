// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdLoadBalancer = cli.Command{
	Name:    "load-balancer",
	Aliases: []string{"ld"},
	Usage:   "manage load balancer",
	Subcommands: []cli.Command{
		{
			Name:  "add",
			Usage: "add a new template",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
