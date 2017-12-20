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
)

// touch qclifile.lua
// qcli make
// qcli make -f other.lua
// cat other.lua | qcli make -stdin

var cmdLake = cli.Command{
	Name:    "lake",
	Aliases: []string{"make"},
	Usage:   "build target with lakefile.lua",
	Hidden:  false,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "set lake/make file",
			Value: "lakefile.lua",
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

		luaPreload(L)

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
