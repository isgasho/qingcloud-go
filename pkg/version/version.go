// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package version

//go:generate go run gen_helper.go
//go:generate go fmt

var (
	ShortVersion   = "dev"
	GitSha1Version = "git-sha1"
	BuildDate      = "2017-01-01"
)
