// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"strings"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

var pkgGeneratorList []GeneratorInterface

type GeneratorInterface interface {
	Name() string
	FileNameExt() string

	HeaderCode(spec *spec_metadata.FileSpec) string
	ServiceCode(spec *spec_metadata.ServiceSpec) string
}

func RegisterGenerater(g GeneratorInterface) {
	pkgGeneratorList = append(pkgGeneratorList, g)
}

func getGenerater(name string) GeneratorInterface {
	for _, g := range pkgGeneratorList {
		if g.Name() == name {
			return g
		}
	}
	for _, g := range pkgGeneratorList {
		if strings.HasPrefix(g.Name(), name) {
			return g
		}
	}
	return nil
}
