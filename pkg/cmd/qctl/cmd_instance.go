// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
describe-instances	获取主机列表	run-instances	创建主机
start-instances	开启主机	stop-instances	关闭主机
restart-instances	重启主机	terminate-instances	销毁主机
resize-instances	修改主机配置	reset-instances	重置操作系统
modify-instance-attributes	修改主机基本属性
*/

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
