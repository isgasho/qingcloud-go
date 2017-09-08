// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// 青云集群管理工具
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chai2010/cli"
	_ "github.com/chai2010/cli/i18n/zh_CN"

	qc "github.com/chai2010/qingcloud-go"
)

const (
	AppName    = "qcluster"
	AppUsage   = "青云集群管理工具"
	AppVersion = qc.Version
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion

	app.UsageText = `qcluster
  qcluster [全局选项] 命令 [命令选项] [参数...]
  qcluster -h`

	app.Authors = []cli.Author{
		{
			Name:  "ChaiShushan",
			Email: "chaishushan@gmail.com",
		},
	}

	// 任何命令执行前
	// 如果 ini 和 public 不存在, 则尝试切换到当前目录(可能是服务模式)
	app.Before = func(context *cli.Context) error {
		flag.Parse() // 避免 glog 警告

		return nil
	}

	// 没有参数时启动
	app.Action = func(c *cli.Context) {
		if c.NArg() > 0 {
			fmt.Fprintf(c.App.Writer, "错误的参数: %s; 帮助请输入 -h\n", c.Args())
			return
		}

		return
	}

	// 未知命令帮助提示
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "没有找到 '%v' 相关的主题!\n", command)
	}

	// 其它子命令
	app.Commands = []cli.Command{
		{
			Name:      "info",
			Usage:     "显示信息",
			ArgsUsage: "[path]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "image-mode",
					Usage: "图像模式",
				},
			},
			Action: func(c *cli.Context) {
				return
			},
		},
		{
			Name:      "list",
			Usage:     "显示列表",
			ArgsUsage: "path",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "regexp",
					Usage: "正则表达式过滤",
				},
			},
			Action: func(c *cli.Context) {
				return
			},
		},
	}

	app.Run(os.Args)
}

func GetAppPath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("main.GetAppPath: %s is directory", p)
	}
	if runtime.GOOS == "windows" {
		if filepath.Ext(p) == "" {
			p += ".exe"
			fi, err := os.Stat(p)
			if err == nil {
				if !fi.Mode().IsDir() {
					return p, nil
				}
				err = fmt.Errorf("main.GetAppPath: %s is directory", p)
			}
		}
	}
	return "", err
}
