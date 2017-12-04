// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
get-monitoring-data	获取主机、公网、路由器的监控	get-loadbalancer-monitoring-data
*/

var cmdMonitor = cli.Command{
	Name:    "monitor",
	Aliases: []string{},
	Usage:   "manage monitor",
	Subcommands: []cli.Command{
		{
			Name:    "get",
			Aliases: []string{},
			Usage:   "get monitoring data",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "get-ld",
			Aliases: []string{},
			Usage:   "get loadbalancer monitoring data",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
