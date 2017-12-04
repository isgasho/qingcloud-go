// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"

	"github.com/urfave/cli"
)

var cmdLoadBalancer = cli.Command{
	Name:    "loadbalancer",
	Aliases: []string{"ld"},
	Usage:   "manage loadbalancer",
	Subcommands: []cli.Command{
		{
			Name:    "describe",
			Aliases: []string{"desc"},
			Usage:   "describe loadbalancers",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify",
			Aliases: []string{"set"},
			Usage:   "modify loadbalancer attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"new"},
			Usage:   "create loadbalancer",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"del"},
			Usage:   "delete loadbalancers",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{},
			Usage:   "start loadbalancers",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"stop"},
			Usage:   "stop loadbalancers",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "associate-eips",
			Aliases: []string{},
			Usage:   "associate eips to loadbalancer",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "dissociate-eips",
			Aliases: []string{},
			Usage:   "dissociate eips from loadbalancer",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "add-listeners",
			Aliases: []string{},
			Usage:   "add loadbalancer listeners",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete-listeners",
			Aliases: []string{"del-listeners"},
			Usage:   "delete loadbalancer listeners",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "add-backends",
			Aliases: []string{},
			Usage:   "add loadbalancer backends",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "delete-backends",
			Aliases: []string{"del-backends"},
			Usage:   "delete loadbalancer backends",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "describe-listeners",
			Aliases: []string{"desc-listeners"},
			Usage:   "describe loadbalancer listeners",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "describe-backends",
			Aliases: []string{"desc-backends"},
			Usage:   "describe loadbalancer backends",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-listener-attributes",
			Aliases: []string{},
			Usage:   "modify loadbalancer listener attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "modify-backend-attributes",
			Aliases: []string{},
			Usage:   "modify loadbalancer backend attributes",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{},
			Usage:   "update loadbalancers",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
