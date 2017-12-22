-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

-- https://github.com/chai2010/qingcloud-go

local qc = require("qingcloud.iaas")

local config = qc.LoadJSON("~/.qingcloud/qcli.json")
local client = qc.Client:new(config)

task("default", {"version", "doc"}, function()
	print("done")
end)

task("version", nil, "print version", function(task)
	print("qc.version")
	print("git sha1: " .. qc.version_info.git_sha1_version)
	print("build time: " .. qc.version_info.build_date)
end)

task("doc", {"install"}, function()
	print("doc")
end)

task("install", nil, function(task, destdir)
	print("install:", task.name, destdir)
end)

task("list.instance", nil, function(task, destdir)
	local reply, err = client:DescribeInstances {
		--owner = "usr-xxxxxxxx",
		zone = "pek3a",
		limit = 100
	}
	if err ~= nil then
		error(err)
	end

	if reply.ret_code ~= 0 then
		error(string.format("err: %d, %s", reply.ret_code, reply.message))
	end

	for i = 1, #reply.instance_set do
		local item = reply.instance_set[i]
		--if item.status ~= "ceased" then
			print(string.format(
				"%s %8s %10s %10s %s %s",
				item.instance_id,
				item.instance_type,
				item.memory_current..'MB',
				item.status,
				item.create_time,
				item.instance_name
			))
		--end
	end

	-- print('total: ' .. reply.total_count)
end)
