-- Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
-- Use of this source code is governed by a Apache
-- license that can be found in the LICENSE file.

task("default", {"verion"}, function()
	print("done")
end)

task("verion", nil, function()
	print("version")
end)

task("doc", {"install"}, function()
	print("doc")
end)

task("install", nil, function(task, destdir)
	print("install")
end)
