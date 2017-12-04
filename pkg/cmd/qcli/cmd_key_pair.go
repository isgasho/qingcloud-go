// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
describe-keypairs	获取密钥列表	modify-keypair-attributes	修改密钥基本属性
create-keypair	新加密钥对	delete-keypairs	删除密钥
attach-keypairs	加载密钥到主机	detach-keypairs
*/

var cmdKeyPair = cli.Command{
	Name:    "keypair",
	Aliases: []string{"ssh"},
	Usage:   "manage keypair",
	Subcommands: []cli.Command{
		{
			Name:    "add",
			Aliases: []string{},
			Usage:   "add a new template",
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
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
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
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
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
