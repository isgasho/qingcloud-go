// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
)

var pkgGeneratorList []GeneratorInterface

type GeneratorInterface interface {
	Name() string
	FileNameExt() string

	GenFileHeaderCode(spec *spec_metadata.FileSpec) (code string, err error)
	GenFileTailCode(spec *spec_metadata.FileSpec) (code string, err error)
	GenServiceCode(spec *spec_metadata.ServiceSpec) (code string, err error)
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
	return nil
}

func genFileHeaderCode(name string, spec *spec_metadata.FileSpec) (code string, err error) {
	return getGenerater(name).GenFileHeaderCode(spec)
}
func genFileTailCode(name string, spec *spec_metadata.FileSpec) (code string, err error) {
	return getGenerater(name).GenFileTailCode(spec)
}
func genServiceCode(name string, spec *spec_metadata.ServiceSpec) (code string, err error) {
	return getGenerater(name).GenServiceCode(spec)
}
