-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

-- https://github.com/chai2010/qingcloud-go

local qc = require("qingcloud.iaas")

if #arg == 1 and arg[1] == '-v' then
	print(qc.version)
	print(qc.version_info.git_sha1_version)
	print(qc.version_info.build_date)
	do return end
end

if #arg == 1 and arg[1] == '-h' then
	print(qc.copyright)
	print("hello, 青云!")
	do return end
end

local config = qc.LoadJSON("~/.qingcloud/qcli.json")
local client = qc.Client:new(config)

local reply, err = client:DescribeInstances {
	--owner = "usr-xxxxxxxx",
	zone = "pek3a",
	limit = 100
}
if err ~= nil then
	print("error:", err)
	do return end
end

if reply.ret_code ~= 0 then
	print(reply.ret_code)
	print(reply.message)
	do return end
end

for i = 1, #reply.instance_set do
	local item = reply.instance_set[i]
	print(i,
		item.instance_id,
		item.instance_type,
		item.memory_current..'MB',
		item.status,
		item.create_time,
		item.instance_name
	)
end

print('total: ' .. reply.total_count)
