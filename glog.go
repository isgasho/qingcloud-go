// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// https://godoc.org/github.com/golang/glog

package qingcloud

import (
	"flag"
	"strconv"
)

type glogLevelType string

const (
	_GLOG_NULL    glogLevelType = ""
	_GLOG_INFO    glogLevelType = "INFO"
	_GLOG_WARNING glogLevelType = "WARNING"
	_GLOG_ERROR   glogLevelType = "ERROR"
	_GLOG_FATAL   glogLevelType = "FATAL"
)

func (s glogLevelType) isValid() bool {
	switch s {
	case _GLOG_INFO, _GLOG_WARNING, _GLOG_ERROR, _GLOG_FATAL:
		return true
	default:
		return false
	}
}

func glogParseLevel(level string) glogLevelType {
	switch level {
	case "debug", "info", "INFO":
		return _GLOG_INFO
	case "warn", "WARNING":
		return _GLOG_WARNING
	case "error", "ERROR":
		return _GLOG_ERROR
	case "fatal", "FATAL":
		return _GLOG_FATAL
	}
	return _GLOG_NULL
}

// Logs are written to standard error instead of to files.
func SetLogToStderr(s bool) {
	flag.CommandLine.Set("logtostderr", strconv.FormatBool(s))
}

// Logs are written to standard error as well as to files.
func SetLogAlsoToStderr(s bool) {
	flag.CommandLine.Set("alsologtostderr", strconv.FormatBool(s))
}

// ame as SetLogLevel, ;og events at or above this severity are logged to standard
// error as well as to files.
func SetLogLevel(s string) {
	if level := glogParseLevel(s); level.isValid() {
		flag.CommandLine.Set("stderrthreshold", string(s))
	}
}

// Log events at or above this severity are logged to standard
// error as well as to files.
func SetLogStderrLevel(s string) {
	if level := glogParseLevel(s); level.isValid() {
		flag.CommandLine.Set("stderrthreshold", string(s))
	}
}

// Log files will be written to this directory instead of the
// default temporary directory.
func SetLogDir(dir string) {
	flag.CommandLine.Set("log_dir", dir)
}
