
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
