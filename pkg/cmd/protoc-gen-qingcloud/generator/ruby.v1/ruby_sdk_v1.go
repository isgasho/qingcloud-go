// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package ruby_sdk_v1

import (
	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"
)

func init() {
	plugin.RegisterGenerater(new(pkgGenerator))
}

type pkgGenerator struct{}

func (p pkgGenerator) Name() string        { return "ruby.v1" }
func (p pkgGenerator) FileNameExt() string { return ".pb.qingcloud.rb" }

func (p pkgGenerator) GenFileHeaderCode(spec *spec_metadata.FileSpec) (code string, err error) {
	panic("TODO")
}
func (p pkgGenerator) GenFileTailCode(spec *spec_metadata.FileSpec) (code string, err error) {
	return "", nil
}

func (p pkgGenerator) GenServiceCode(spec *spec_metadata.ServiceSpec) (code string, err error) {
	panic("TODO")
}
