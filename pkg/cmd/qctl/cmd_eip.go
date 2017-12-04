// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdEip = cli.Command{
	Name:    "eip",
	Aliases: []string{},
	Usage:   "manage EIP",
	Subcommands: []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"desc"},
			Usage:   "describe eips",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "allocate-eips",
			Aliases: []string{},
			Usage:   "allocate-eips",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "associate-eip",
			Aliases: []string{},
			Usage:   "associate-eip",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "dissociate-eips",
			Aliases: []string{},
			Usage:   "dissociate-eips",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "change-eips-bandwidth",
			Aliases: []string{},
			Usage:   "change-eips-bandwidth",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "change-eips-billing-mode",
			Aliases: []string{},
			Usage:   "change-eips-billing-mode",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-eip-attributes",
			Aliases: []string{},
			Usage:   "modify-eip-attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "release-eips",
			Aliases: []string{},
			Usage:   "release-eips",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
