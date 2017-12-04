// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdKeyPair = cli.Command{
	Name:    "keypair",
	Aliases: []string{"ssh"},
	Usage:   "manage keypair",
	Subcommands: []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"desc"},
			Usage:   "describe keypairs",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify",
			Aliases: []string{"set"},
			Usage:   "modify keypair attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"new"},
			Usage:   "create keypair",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"del"},
			Usage:   "delete keypairs",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "attach",
			Aliases: []string{"bind"},
			Usage:   "attach keypairs",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "detach",
			Aliases: []string{"unbind"},
			Usage:   "detach keypairs",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
