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
		cli.BoolFlag{
			Name:  "tasks, t",
			Usage: "show task list",
		},
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
			Name:  "stdin, s",
			Usage: "read from stdin",
		},
		cli.BoolFlag{
			Name:  "graph, g",
			Usage: "generate graphviz file",
		},
	},
	Action: func(c *cli.Context) error {
		L := lua.NewState()
		defer L.Close()

		if c.IsSet("dir") {
			if newdir := c.String("dir"); newdir != "" {
				if err := os.Chdir(newdir); err != nil {
					log.Fatal(err)
				}
			}
		}

		L.SetGlobal("arg", func() *lua.LTable {
			tb := L.NewTable()
			for i := 0; i < c.NArg(); i++ {
				tb.Append(lua.LString(c.Args()[i]))
			}
			return tb
		}())

		L.SetGlobal("lake_flags", func() *lua.LTable {
			tb := L.NewTable()
			if c.IsSet("file") {
				if s := c.String("file"); s != "" {
					tb.RawSetString("lakefile_name", lua.LString(s))
				}
			}
			if c.IsSet("stdin") {
				// EOF: UNIX Ctrl+D, Windows Ctrl+Z
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Fatal(err)
				}
				tb.RawSetString("lakefile_content", lua.LString(b))
			}
			if c.IsSet("tasks") {
				tb.RawSetString("show_tasks", lua.LTrue)
			}
			if c.IsSet("graph") {
				tb.RawSetString("gen_graph", lua.LTrue)
			}
			return tb
		}())

		luaPreload(L)

		if err := luaDostring(L, ":/lake/main.lua", lakeMainCode); err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

const lakeMainCode = `
local Lake = require("Lake")

local lake = Lake.Application:new()

if lake_flags.lakefile_name then
	lake.lakefile = lake_flags.lakefile_name
end
if lake_flags.lakefile_content then
	lake.lakefile_content = lake_flags.lakefile_content
end

-- The Lake "DSL"
-- These functions allow to interface with the build engine in a succinct way
function task(name, prerequisites, describeOrAction, action)
	if type(describeOrAction) == "function" then
		lake:defineTask(name, prerequisites, describeOrAction)
	else
		assert(type(describeOrAction) == "string")
		assert(type(action) == "function")

		lake:defineTask(name, prerequisites, action)
		lake.tasks[name].describe = describeOrAction
	end
end

if lake_flags.show_tasks then
	lake:show_tasks()
elseif lake_flags.gen_graph then
	print(lake:gen_graph())
else
	lake:run(arg)
end
`
