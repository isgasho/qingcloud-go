// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"strings"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

var pkgServiceGeneratorList []ServiceGenerator

type ServiceGenerator interface {
	Name() string
	FileNameExt() string

	HeaderCode(spec *spec_metadata.FileSpec) string
	ServiceCode(spec *spec_metadata.ServiceSpec) string
}

func RegisterServiceGenerater(g ServiceGenerator) {
	pkgServiceGeneratorList = append(pkgServiceGeneratorList, g)
}

func getServiceGenerater(name string) ServiceGenerator {
	for _, g := range pkgServiceGeneratorList {
		if g.Name() == name {
			return g
		}
	}
	for _, g := range pkgServiceGeneratorList {
		if strings.HasPrefix(g.Name(), name) {
			return g
		}
	}
	return nil
}
