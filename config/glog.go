// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// https://godoc.org/github.com/golang/glog

package config

import (
	"flag"
	"strconv"
)

type Glog_LevelType string

const (
	GLOG_NULL    Glog_LevelType = ""
	GLOG_INFO    Glog_LevelType = "INFO"
	GLOG_WARNING Glog_LevelType = "WARNING"
	GLOG_ERROR   Glog_LevelType = "ERROR"
	GLOG_FATAL   Glog_LevelType = "FATAL"
)

func (s Glog_LevelType) isValid() bool {
	switch s {
	case GLOG_INFO, GLOG_WARNING, GLOG_ERROR, GLOG_FATAL:
		return true
	default:
		return false
	}
}

func ParseGlogLevel(level string) Glog_LevelType {
	switch level {
	case "debug", "info", "INFO":
		return GLOG_INFO
	case "warn", "WARNING":
		return GLOG_WARNING
	case "error", "ERROR":
		return GLOG_ERROR
	case "fatal", "FATAL":
		return GLOG_FATAL
	}
	return GLOG_NULL
}

func Glog_SetLogDir(dir string) {
	flag.CommandLine.Set("log_dir", dir)
}

func Glog_SetAlsoToStderr(s bool) {
	flag.CommandLine.Set("alsologtostderr", strconv.FormatBool(s))
}

func Glog_SetStderrThreshold(s Glog_LevelType) {
	if s.isValid() {
		flag.CommandLine.Set("stderrthreshold", string(s))
	}
}
