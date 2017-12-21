-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

local getopt = require("getopt")
local Lake = require("Lake")

local lake = Lake.Application:new()

local opts, newarg = getopt(arg)

if opts.h or opts.help then
	print [[
qlua qlake.lua [options] targets...
Usage:
  qlua qlake.lua -h
  qlua qlake.lua --lakefile=other.lua
  qlua qlake.lua --stdin
  qlua qlake.lua

Options:
  -t --show_tasks show target list
  -s --stdin      read lakefile from stdin
  -f --lakefile   set lakefile
  -h --help       show help info
]]

	do return end
end

-- -f --lakefile
if opts.f and opts.f ~= "" then
	lake.lakefile = opts.f
end
if opts.lakefile and opts.lakefile ~= "" then
	lake.lakefile = opts.lakefile
end

-- -s --stdin
if opts.s or opts.stdin then
	lake.lakefile_content = io.read("*a")
end

-- The Lake "DSL"
-- These functions allow to interface with the build engine in a succinct way
function task(name, prerequisites, action)
	lake:defineTask(name, prerequisites, action)
end

-- -t --show_tasks
if opts.t or opts.show_tasks then
	lake:show_tasks()
else
	lake:run(arg)
end
