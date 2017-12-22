// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated by "makestatic"; DO NOT EDIT.

package gopherlua_lake

var Files = map[string]string{
	"lake-0.1.0/Lake/init.lua": `-- Copyright (c) 2009 Sebastian Nowicki <sebnow@gmail.com>
--
-- Permission is hereby granted, free of charge, to any person obtaining a copy
-- of this software and associated documentation files (the "Software"), to deal
-- in the Software without restriction, including without limitation the rights
-- to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-- copies of the Software, and to permit persons to whom the Software is
-- furnished to do so, subject to the following conditions:
--
-- The above copyright notice and this permission notice shall be included in
-- all copies or substantial portions of the Software.
--
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-- IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-- FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-- AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-- LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-- OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
-- SOFTWARE.

-- Export the Lake namespace
return {
	Application = require("Lake.Application"),
	Task = require("Lake.Task"),
}
`,

	"lake-0.1.0/Lake/Application.lua": `-- Copyright (c) 2009 Sebastian Nowicki <sebnow@gmail.com>
--
-- Permission is hereby granted, free of charge, to any person obtaining a copy
-- of this software and associated documentation files (the "Software"), to deal
-- in the Software without restriction, including without limitation the rights
-- to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-- copies of the Software, and to permit persons to whom the Software is
-- furnished to do so, subject to the following conditions:
--
-- The above copyright notice and this permission notice shall be included in
-- all copies or substantial portions of the Software.
--
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-- IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-- FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-- AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-- LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-- OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
-- SOFTWARE.

local Object = require("Lake.Object")
local Task = require("Lake.Task")

-- The application interface for the build engine
local Application = Object:clone{
	name = nil,
	lakefiles = nil,
	lakefile = nil,
	lakefile_content = nil,
	prerequisites = {},
	tasks = {},
}

--- Initialize a new application.
-- @param name the name of the application.
-- @return A new application.
function Application:new(name)
	local DEFAULT_LAKEFILES = {
		"lakefile", "Lakefile",
		"lakefile.lua", "Lakefile.lua"
	}
	self = Application:clone{
		name = name or "lake",
		lakefiles = DEFAULT_LAKEFILES,
	}
	return self
end

function Application:gen_graph()
	self:loadLakefile()

	-- digraph begin
	local out = "digraph G {\n"

	-- generate nodes
	for name, task in pairs(self.tasks) do
		local id = string.gsub(name, "[^a-zA-Z0-9]", "_")
		out = out .. string.format('\t%s [label = "%s"];\n', id, name)
	end

	-- generate edge
	if #self.prerequisites > 0 then
		out = out .. '\n'
	end
	for name, task in pairs(self.tasks) do
		local id0 = string.gsub(name, "[^a-zA-Z0-9]", "_")
		for _, prerequisite in ipairs(task.prerequisites or {}) do
			local id1 = string.gsub(prerequisite.name, "[^a-zA-Z0-9]", "_")
			out = out .. string.format('\t%s -> %s;\n', id0, id1)
		end
	end

	-- digraph end
	out = out .. "}"

	-- ok
	return out
end

function Application:show_tasks()
	self:loadLakefile()
	local maxNameLen = 1
	for name, _ in pairs(self.tasks) do
		if #name > maxNameLen then
			maxNameLen = #name
		end
	end

	local fmt = "%-" .. maxNameLen .. "s - %s"
	for _, task in pairs(self.tasks) do
		if task.describe and task.describe ~= "" then
			print(string.format(fmt, task.name, task.describe))
		else
			print(string.format(fmt, task.name, task.name .. " target"))
		end
	end
end

--- Invoke all tasks in args.
-- @param args a list of options, and tasks to be invoked.
function Application:run(args)
	-- Collect tasks to be run
	local targets = {}
	for _, task in ipairs(args or {}) do
		table.insert(targets, task)
	end
	-- Add default task if none were specified
	if #targets == 0 then
		table.insert(targets, "default")
	end
	self:loadLakefile()
	-- Invoke all tasks
	for _, name in ipairs(targets) do
		self:invokeTask(name)
	end
end

--- Parse a lakefile.
function Application:loadLakefile()
	local file
	if self.lakefile then
		file = self.lakefile
	else
		file = self:findLakefile()
		self.lakefile = file
	end
	if self.lakefile_content then
		local f = assert(loadstring(self.lakefile_content))
		f()
	else
		self.lakefile = file
		local f = assert(loadfile(file))
		f()
	end
end

-- Attempt to resolve prerequisites in self.prerequisites to task
-- prerequisites
function Application:_resolvePrerequisites()
	for i, edge in ipairs(self.prerequisites) do
		local task, prerequisite = unpack(edge)
		if self.tasks[prerequisite] then
			task:enhance({self.tasks[prerequisite]})
			-- self.prerequisites[i] = nil
		end
	end
end

--- Define a new task.
--
-- This should be used from the lakefile, as opposed to using Task:new()
-- directly. It defines tasks in a textual manner (without actually creating
-- Task objects), tracking prerequisites. When a task is to be invoked, they
-- are resolved and real Task objects are created.
-- @param name the name of a task.
-- @param prerequisites the tasks this task depends on.
-- @param action the action to be executed.
function Application:defineTask(name, prerequisites, action)
	assert(name and type(name) == "string")
	local task = self.tasks[name] or Task:new(name, nil, action)
	self.tasks[name] = task
	-- prerequisites must be a table
	if type(prerequisites) == "string" then
		prerequisites = {prerequisites}
	end
	-- Track prerequisites for this task
	for _, prerequisite in ipairs(prerequisites or {}) do
		table.insert(self.prerequisites, {task, prerequisite})
	end
	-- Attempt to resolve prerequisites and add to task
	self:_resolvePrerequisites()
end

-- Split task name and argument list.
-- @param raw_args a task deifnition string, e.g. "foo[bar, baz]".
-- @return The name (e.g. "foo") and a list of arguments
-- (e.g. {"bar", "baz"}) or nil if no arguments are given.
local function parseTaskArguments(raw_args)
	-- Split task name and argument list (e.g. "[foo, bar]")
	local name, arg_list = raw_args:match("([^%[]+)%[(.-)%]")
	local args = {}
	if arg_list then
		-- Split argument list and add each element to args
		for arg in arg_list:gmatch("([^,]+)[%s,]*") do
			table.insert(args, arg)
		end
	end
	return name, args
end

--- Invoke a task.
-- If tasks have not yet been resolved, they will be now.
-- @param task a string representing the task to be invoked.
function Application:invokeTask(task)
	local name, args = parseTaskArguments(task)
	name = name or task
	-- Ensure all tasks are defined
	self:_resolvePrerequisites()
	if self.tasks[name] then
		self.tasks[name]:invoke(unpack(args))
	end
end

--- Find a lakefile to execute.
-- This searches self.lakefiles for possible lakefile filenames.
-- @return The path to a lakefile.
function Application:findLakefile()
	for _, file in ipairs(self.lakefiles) do
		local fp, _, code = io.open(file, "r")
		if code == nil then
			io.close(fp)
			return file;
		end
	end
end

return Application
`,

	"lake-0.1.0/Lake/Task.lua": `-- Copyright (c) 2009 Sebastian Nowicki <sebnow@gmail.com>
--
-- Permission is hereby granted, free of charge, to any person obtaining a copy
-- of this software and associated documentation files (the "Software"), to deal
-- in the Software without restriction, including without limitation the rights
-- to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-- copies of the Software, and to permit persons to whom the Software is
-- furnished to do so, subject to the following conditions:
--
-- The above copyright notice and this permission notice shall be included in
-- all copies or substantial portions of the Software.
--
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-- IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-- FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-- AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-- LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-- OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
-- SOFTWARE.

local Object = require("Lake.Object")

--- The base unit of work.
-- A task can depend on other tasks, and when invoked executes its action.
local Task = Object:clone()

--- The name of a task.
Task.name = nil
--- The describe of a task.
Task.describe = nil
--- Other tasks this task depends on.
Task.prerequisites = {}
--- Actions to be executed for this task.
Task.actions = {}
--- Whether the task was already invoked. This gets reset when Task:enable()
-- is called.
Task.wasInvoked = false

--- Initialize a new task.
-- @param name the name of the task (required).
-- @param prerequisites a list of tasks this task depends on.
-- @param action the function to be called when the task is invoked.
-- @return A new task.
function Task:new(name, prerequisites, action)
	assert(name and type(name) == "string")
	return Task:clone{
		name = name,
		describe = "",
		prerequisites = prerequisites or {},
		actions = {action}
	}
end

--- Execute the actions associated with this task.
-- This does not execute the actions of a task's prerequisites.
-- @param ... arguments to be passed to the action.
function Task:execute(...)
	for _, v in ipairs(self.actions) do
		v(self, ...)
	end
end

--- Invoke a task and its prerequisites.
-- The task will be executed when all prerequisites have been executed. It
-- will not be executed if it was already invoked before. To allow it to be
-- invoked again, use Task:enable().
-- @param ... arguments to be passed to all actions being executed.
-- @see Task:enable()
function Task:invoke(...)
	if not self.wasInvoked then
		self.wasInvoked = true
		self:_invokePrerequisites(...)
		self:execute(...)
	end
end

--- Add prerequisites and actions to a task.
-- @param prerequisites a list of prerequisites to be added.
-- @param actions a list of actions to be added.
function Task:enhance(prerequisites, actions)
	for _, task in ipairs(prerequisites or {}) do
		local hasPrerequisite = false
		for _, prereq in ipairs(self.prerequisites) do
			if prereq == task then
				hasPrerequisite = true
				break
			end
		end
		if not hasPrerequisite then
			table.insert(self.prerequisites, task)
		end
	end
	for _, action in ipairs(actions or {}) do
		table.insert(self.actions, action)
	end
end

-- Invoke a task's prerequisites.
function Task:_invokePrerequisites(...)
	for _, task in ipairs(self.prerequisites) do
		task:invoke(...)
	end
end

--- Enable a task to be invoked, even if it has already been invoked.
function Task:enable()
	self.wasInvoked = false
end

return Task
`,

	"lake-0.1.0/Lake/Object.lua": `-- Copyright (c) 2009 Sebastian Nowicki <sebnow@gmail.com>
--
-- Permission is hereby granted, free of charge, to any person obtaining a copy
-- of this software and associated documentation files (the "Software"), to deal
-- in the Software without restriction, including without limitation the rights
-- to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-- copies of the Software, and to permit persons to whom the Software is
-- furnished to do so, subject to the following conditions:
--
-- The above copyright notice and this permission notice shall be included in
-- all copies or substantial portions of the Software.
--
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-- IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-- FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-- AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-- LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-- OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
-- SOFTWARE.

-- Object Root object prototype
local Object = {}

--- Clone an object.
-- @param object the object to be cloned.
-- @return A clone of object.
function Object:clone(object)
	object = object or {}
	setmetatable(object, self)
	self.__index = self
	return object
end

return Object
`,
}
