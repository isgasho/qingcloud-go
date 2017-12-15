-- Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

-- https://github.com/chai2010/qingcloud-go

local qc = require("qingcloud")

print("hello, 青云!")
print(qc.version)
print(qc.git_sha1_version)

local tb = qc.decode_json [[
	{
		"count":1,
		"vxnets.1":"vxnet-0",
		"zone":"pek1",
		"instance_type":"small_b",
		"signature_version":1,
		"signature_method":"HmacSHA256",
		"instance_name":"demo",
		"image_id":"centos64x86a",
		"login_mode":"passwd",
		"login_passwd":"QingCloud20130712",
		"version":1,
		"access_key_id":"QYACCESSKEYIDEXAMPLE",
		"action":"RunInstances",
		"time_stamp":"2013-08-27T14:30:10Z"
	}
]]

for k, v in pairs(tb) do
	print(string.format("%s: %s", k, v))
end

for k, v in pairs(qc.all_method_list) do
	--print(k, v.method)
end

for k, v in pairs(qc.server_info) do
	print(k, v)
end

local out, err = qc.call_method {
	server_info = qc.server_info,
	action = "DescribeInstances",
	request = {},
}

print("---------")

for k, v in pairs(out) do
	print(k, v)
end

print(type(out))
print(out.action)
print(out.ret_code)
