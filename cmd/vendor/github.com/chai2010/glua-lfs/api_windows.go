package lfs

import (
	"os"
	"syscall"

	"github.com/yuin/gopher-lua"
)

func attributesFill(tbl *lua.LTable, stat os.FileInfo) {
	sys, ok := stat.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return
	}
	_ = sys

	{
		var mode string
		switch stat.Mode() & syscall.S_IFMT {
		case syscall.S_IFREG:
			mode = "file"
		case syscall.S_IFDIR:
			mode = "directory"
		case syscall.S_IFLNK:
			mode = "link"
		case syscall.S_IFSOCK:
			mode = "socket"
		case syscall.S_IFIFO:
			mode = "named pipe"
		case syscall.S_IFCHR:
			mode = "char device"
		case syscall.S_IFBLK:
			mode = "block device"
		default:
			mode = "other"
		}
		tbl.RawSetH(lua.LString("mode"), lua.LString(mode))
	}
}
