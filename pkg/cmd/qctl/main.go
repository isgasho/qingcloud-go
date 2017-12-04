// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qctl

import (
	"os"

	"github.com/urfave/cli"

	_ "github.com/chai2010/qingcloud-go/pkg/internal/glogger"
)

func Main() {
	app := cli.NewApp()
	app.Name = "qctl"
	app.Usage = "QingCloud SDK control"
	app.Action = func(c *cli.Context) error {
	  println("TODO")
	  return nil
	}

	app.Run(os.Args)
}
