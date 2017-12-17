// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/yuin/gopher-lua"

	qc_iaas "github.com/chai2010/qingcloud-go/pkg/gopher-lua/qingcloud.iaas"
)

// touch qclifile.lua
// qcli make
// qcli make -f other.lua
// cat other.lua | qcli make -stdin

var cmdLuaMake = cli.Command{
	Name:    "lake",
	Aliases: []string{"make"},
	Usage:   "make target with qclifile.lua",
	Hidden:  false,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "set lake/make file",
			Value: "qclifile.lua",
		},
		cli.StringFlag{
			Name:  "dir, C",
			Usage: "change work directory",
			Value: "",
		},
		cli.BoolFlag{
			Name:  "stdin",
			Usage: "read from stdin",
		},
	},
	Action: func(c *cli.Context) error {
		L := lua.NewState()
		defer L.Close()

		argtb := L.NewTable()
		for i := 0; i < c.NArg(); i++ {
			L.RawSet(argtb, lua.LNumber(i+1), lua.LString(c.Args()[i]))
		}
		L.SetGlobal("arg", argtb)

		luaOpenSDK(c, L)
		qc_iaas.Preload(L)

		if c.IsSet("dir") {
			if newdir := c.String("dir"); newdir != "" {
				if err := os.Chdir(newdir); err != nil {
					log.Fatal(err)
				}
			}
		}

		if c.IsSet("stdin") {
			// EOF: UNIX Ctrl+D, Windows Ctrl+Z
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			prog := string(b)
			if err := L.DoString(prog); err != nil {
				log.Fatal(err)
			}
			return nil
		}

		lakefile := c.String("file")
		fi, err := os.Stat(lakefile)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("%q is not exists", lakefile)
			} else {
				log.Fatal(err)
			}
		}
		if fi.IsDir() {
			log.Fatalf("%q is not file", lakefile)
		}

		if err := L.DoFile(lakefile); err != nil {
			log.Fatal(err)
		}

		return nil
	},
}
