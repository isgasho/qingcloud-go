// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package gopherlua_getopt

import (
	"log"
	"strings"

	"github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	code := strings.Replace(pkgModCode, "{{.ModName}}", pkgModName, -1)

	L.PreloadModule(pkgModName, func(L *lua.LState) int {
		if err := dostring(L, pkgModName+":/gopherlua_getopt/main.lua", code); err != nil {
			log.Fatal(err)
		}

		if tb := L.ToTable(-1); tb != nil {
			L.Push(tb)
			return 1
		}
		return 0
	})
}

func dostring(ls *lua.LState, name, source string) error {
	if name == "" {
		name = "<string>"
	}

	fn, err := ls.Load(strings.NewReader(source), name)
	if err != nil {
		return err
	}

	ls.Push(fn)
	return ls.PCall(0, lua.MultRet, nil)
}

var pkgModName = "getopt"

var pkgModCode = `
local getopt ={
	_URL = 'http://lua-users.org/wiki/AlternativeGetOpt',
	_DESCRIPTION = 'Alternative Get Opt',
}

-- getopt_alt.lua

-- getopt, POSIX style command line argument parser
-- param arg contains the command line arguments in a standard table.
-- param options is a string with the letters that expect string values.
-- returns a table where associated keys are true, nil, or a string value.
-- The following example styles are supported
--   -aone   ==> opts["a"]=="one"
--   -btwo   ==> opts["b"]=="two"
--   -c      ==> opts["c"]==true
--   --c=one ==> opts["c"]=="one"
--   -cdaone ==> opts["c"]==true opts["d"]==true opts["a"]=="one"
-- note POSIX demands the parser ends at the first non option
--      this behavior isn't implemented.

function getopt.getopt( arg, options )
	local tab = {}
	local newarg = {}

	options = options or ""

	for k, v in ipairs(arg) do
		if string.sub( v, 1, 2) == "--" then
			local x = string.find( v, "=", 1, true )
			if x then
				tab[ string.sub( v, 3, x-1 ) ] = string.sub( v, x+1 )
			else
				tab[ string.sub( v, 3 ) ] = true
			end
		elseif string.sub( v, 1, 1 ) == "-" then
			local y = 2
			local l = string.len(v)
			local jopt
			while ( y <= l ) do
				jopt = string.sub( v, y, y )
				if string.find( options, jopt, 1, true ) then
					if y < l then
						tab[ jopt ] = string.sub( v, y+1 )
						y = l
					else
						tab[ jopt ] = arg[ k + 1 ]
					end
				else
					tab[ jopt ] = true
				end
				y = y + 1
			end
		else
			table.insert(newarg, v)
		end
	end

	return tab, newarg
end

--[[
-- Test code
local getopt = require("getopt")

opts = getopt( arg, "ab" )
for k, v in pairs(opts) do
  print( k, v )
end

-- EOF
--]]

setmetatable(getopt, { __call = function(_, ...) return getopt.getopt(...) end })

return getopt
`
