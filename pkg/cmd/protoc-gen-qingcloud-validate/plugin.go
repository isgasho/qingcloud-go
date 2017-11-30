// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin_validate

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

var pkgValidateGeneratorList []ValidateGenerator

type ValidateGenerator interface {
	Name() string
	FileNameExt() string

	HeaderCode(file *generator.FileDescriptor) string
	ValidateCode(file *generator.FileDescriptor, msgDesc *descriptor.DescriptorProto) string
}

func RegisterValidateGenerator(g ValidateGenerator) {
	pkgValidateGeneratorList = append(pkgValidateGeneratorList, g)
}

func getAllValidateGenerator() []ValidateGenerator {
	return pkgValidateGeneratorList
}

func getAllValidateGeneratorNames() (names []string) {
	for _, g := range pkgValidateGeneratorList {
		names = append(names, g.Name())
	}
	return
}

func getValidateGenerator(name string) ValidateGenerator {
	for _, g := range pkgValidateGeneratorList {
		if g.Name() == name {
			return g
		}
	}
	return nil
}
