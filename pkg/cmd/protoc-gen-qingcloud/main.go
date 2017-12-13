// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// protoc-gen-go is a plugin for the Google protocol buffer compiler to generate
// Go code.  Run it by building this program and putting it in your path with
// the name
// 	protoc-gen-go
// That word 'go' at the end becomes part of the option string set for the
// protocol compiler, so once the protocol compiler (protoc) is installed
// you can run
// 	protoc --go_out=output_directory input_directory/file.proto
// to generate Go bindings for the protocol defined by file.proto.
// With that input, the output will be written to
// 	output_directory/file.pb.go
//
// The generated code is documented in the package comment for
// the library.
//
// See the README and documentation for protocol buffers to learn more:
// 	https://developers.google.com/protocol-buffers/
package qingcloud_plugin

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func Main() {
	qcPlugin := new(qingcloudPlugin)
	generator.RegisterPlugin(qcPlugin)

	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()
	if len(getAllServiceGenerator()) == 0 {
		g.Fail("no service generator plugin")
	}

	pkgReadRequetFromStdin(g)
	pkgGenerateAllFiles(g, qcPlugin)
	pkgWriteResponseToStdout(g)
}

func pkgGenerateAllFiles(g *generator.Generator, qcPlugin *qingcloudPlugin) {
	// set default plugins: qingcloud
	// protoc --qingcloud_out=plugin=golang:. x.proto
	qcPlugin.InitService(pkgGetUserPlugin(g))

	// parse command line parameters
	g.CommandLineParameters("plugins=" + qcPlugin.Name())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// skip non *.pb.qingcloud.go
	respFileList := g.Response.File[:0]
	for _, file := range g.Response.File {
		fileName := file.GetName()
		extName := qcPlugin.FileNameExt()

		if strings.HasSuffix(fileName, extName) {
			respFileList = append(respFileList, file)
		}
	}
	g.Response.File = respFileList
}

func pkgReadRequetFromStdin(g *generator.Generator) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
}

func pkgWriteResponseToStdout(g *generator.Generator) {
	data, err := proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func pkgGetUserPlugin(g *generator.Generator) ServiceGenerator {
	args := g.Request.GetParameter()
	userPluginName := pkgGetParameterValue(args, "plugin")
	if userPluginName == "" {
		userPluginName = getFirstServiceGeneratorName()
	}

	userPlugin := getServiceGenerator(userPluginName)
	if userPlugin == nil {
		log.Print("protoc-gen-qingcloud: registor plugins:", getAllServiceGeneratorNames())
		g.Fail("invalid plugin option:", userPluginName)
	}

	return userPlugin
}

func pkgGetParameterValue(parameter, key string) string {
	for _, p := range strings.Split(parameter, ",") {
		if i := strings.Index(p, "="); i > 0 {
			if p[0:i] == key {
				return p[i+1:]
			}
		}
	}
	return ""
}
