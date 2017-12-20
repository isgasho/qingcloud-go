-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

-- https://github.com/chai2010/qingcloud-go

local qc = require("qingcloud.iaas")

task("default", {"verion"}, function()
	print("done")
end)

task("verion", nil, function()
	print("qc.version")
	print("git sha1: " .. qc.version_info.git_sha1_version)
	print("build time: " .. qc.version_info.build_date)
end)

task("doc", {"install"}, function()
	print("doc")
end)

task("install", nil, function(task, destdir)
	print("install")
end)
