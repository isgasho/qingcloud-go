// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdSecurityGroup = cli.Command{
	Name:    "security-group",
	Aliases: []string{"sg"},
	Usage:   "manage security group",
	Subcommands: []cli.Command{
		{
			Name:    "describe-security-groups",
			Aliases: []string{},
			Usage:   "add a new template",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-security-group-attributes",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "create-security-group",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete-security-groups",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "describe-security-group-rules",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-security-group-rule-attributes",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "add-security-group-rules",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete-security-group-rules",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "apply-security-group",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
