// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin_validate

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func Main() {
	qcPlugin := new(validatePlugin)
	generator.RegisterPlugin(qcPlugin)

	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	g := generator.New()

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

	if len(getAllValidateGenerator()) == 0 {
		g.Fail("no validate generator plugin")
	}

	// set default plugins: qingcloud
	// protoc --qingcloud_out=plugin=golang:. x.proto
	if s := getParameterValue(g.Request.GetParameter(), "plugin"); s != "" {
		if x := getValidateGenerator(s); x != nil {
			qcPlugin.InitValidateGenerator(x)
		} else {
			log.Print("protoc-gen-qingcloud-validate: registor plugins:", getAllValidateGeneratorNames())
			g.Fail("invalid plugin option:", s)
		}
	} else {
		if x := getAllValidateGenerator(); len(x) == 1 {
			qcPlugin.InitValidateGenerator(x[0])
		} else {
			log.Print("protoc-gen-qingcloud-validate: registor plugins:", getAllValidateGeneratorNames())
			g.Fail("no plugin option")
		}
	}

	// parse command line parameters
	g.CommandLineParameters("plugins=" + qcPlugin.Name())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// skip *.pb.go
	respFileList := g.Response.File[:0]
	for _, file := range g.Response.File {
		if !strings.HasSuffix(file.GetName(), ".pb.go") {
			respFileList = append(respFileList, file)
		}
	}
	g.Response.File = respFileList

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

func getParameterValue(parameter, key string) string {
	for _, p := range strings.Split(parameter, ",") {
		if i := strings.Index(p, "="); i > 0 {
			if p[0:i] == key {
				return p[i+1:]
			}
		}
	}
	return ""
}
