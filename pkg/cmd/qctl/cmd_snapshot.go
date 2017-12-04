// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
describe-snapshots	获取备份列表	modify-snapshot-attributes	修改备份的基本属性
create-snapshots	创建备份	delete-snapshots	删除备份
create-volume-from-snapshot	从备份创建硬盘	capture-instance-from-snapshot	从备份创建自有映像
apply-snapshots	回滚到指定备份点
*/

var cmdSnapshot = cli.Command{
	Name:    "snapshot",
	Aliases: []string{"snap"},
	Usage:   "manage snapshot",
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
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{},
			Usage:   "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
