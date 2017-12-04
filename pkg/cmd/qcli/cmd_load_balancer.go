// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
describe-loadbalancers	获取负载均衡器列表	modify-loadbalancer-attributes	修改负载均衡器的基本属性
create-loadbalancer	创建负载均衡器	delete-loadbalancers	删除负载均衡器
start-loadbalancers	启动负载均衡器	stop-loadbalancers	关闭负载均衡器
associate-eips-to-loadbalancer	给负载均衡器绑定公网IP	dissociate-eips-from-loadbalancer	将公网IP从负载均衡器上解绑
add-loadbalancer-listeners	添加监听器	delete-loadbalancer-listeners	删除监听器
add-loadbalancer-backends	添加监听器下的后端服务	delete-loadbalancer-backends	删除后端服务
describe-loadbalancer-listeners	获取负载均衡器监听器列表	describe-loadbalancer-backends	获取负载均衡器后端服务列表
modify-loadbalancer-listener-attributes	修改监听器基本属性	modify-loadbalancer-backend-attributes	修改后端服务基本属性
update-loadbalancers	更新负载均衡器配置
*/

var cmdLoadBalancer = cli.Command{
	Name:    "load-balancer",
	Aliases: []string{"ld"},
	Usage:   "manage load balancer",
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
