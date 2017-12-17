-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

-- https://github.com/chai2010/qingcloud-go

local qc = require("qingcloud.iaas")

print("hello, 青云!")

print(qc.copyright)

print(qc.version)
print(qc.version_info.git_sha1_version)
print(qc.version_info.build_date)

local config = qc.LoadJSON("~/.qingcloud/qcli.json")
local client = qc.Client:new(config)

local reply, err = client:DescribeInstances({})
assert(err == nil)
print(type(reply))

print(reply.action)
print(reply.ret_code)
print(reply.message)
