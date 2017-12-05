// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

var pkgServiceGeneratorList []ServiceGenerator

type ServiceGenerator interface {
	Name() string
	FileNameExt() string

	HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string
	ServiceCode(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string
	MessageCode(p *generator.Generator, file *generator.FileDescriptor, msg *descriptor.DescriptorProto) string
}

func RegisterServiceGenerator(g ServiceGenerator) {
	pkgServiceGeneratorList = append(pkgServiceGeneratorList, g)
}

func getAllServiceGenerator() []ServiceGenerator {
	return pkgServiceGeneratorList
}

func getAllServiceGeneratorNames() (names []string) {
	for _, g := range pkgServiceGeneratorList {
		names = append(names, g.Name())
	}
	return
}

func getServiceGenerator(name string) ServiceGenerator {
	for _, g := range pkgServiceGeneratorList {
		if g.Name() == name {
			return g
		}
	}
	return nil
}
