// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdInstance = cli.Command{
	Name:    "instance",
	Aliases: []string{"ins"},
	Usage:   "manage instance",
	Subcommands: []cli.Command{
		{
			Name:    "describe-instances",
			Aliases: []string{},
			Usage:   "describe-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "run-instances",
			Aliases: []string{},
			Usage:   "run-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "start-instances",
			Aliases: []string{},
			Usage:   "start-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "stop-instances",
			Aliases: []string{},
			Usage:   "stop-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "restart-instances",
			Aliases: []string{},
			Usage:   "restart-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "terminate-instances",
			Aliases: []string{},
			Usage:   "terminate-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "resize-instances",
			Aliases: []string{},
			Usage:   "resize-instances",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "reset-instances",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-instance-attributes",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
