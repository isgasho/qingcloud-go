// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
describe-images	获取自有映像列表	modify-image-attributes	修改自有映像的基本属性
capture-instance	将关闭的主机捕获为自有映像	delete-images
*/

var cmdImage = cli.Command{
	Name:    "image",
	Aliases: []string{"img"},
	Usage:   "manage image",
	Subcommands: []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"desc"},
			Usage:   "get image list",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "capture",
			Aliases: []string{"create", "new"},
			Usage:   "create a new image",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify",
			Aliases: []string{"update", "set"},
			Usage:   "modify image attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"del"},
			Usage:   "delete a image",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
