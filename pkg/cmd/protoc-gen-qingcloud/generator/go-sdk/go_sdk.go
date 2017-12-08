// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package golang_sdk_v1

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"
	"github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils"
)

func init() {
	plugin.RegisterServiceGenerator(new(pkgGenerator))
}

type pkgGenerator struct{}

func (p pkgGenerator) Name() string        { return "go-sdk" }
func (p pkgGenerator) FileNameExt() string { return ".pb.qingcloud.go" }

func (pkgGenerator) HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string {
	spec := utils.BuildFileSpec(g, file)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplFileHeader))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func (pkgGenerator) ServiceCode(g *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string {
	spec := utils.BuildServiceSpec(g, file, svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
func (pkgGenerator) MessageCode(p *generator.Generator, file *generator.FileDescriptor, msg *descriptor.DescriptorProto) string {
	return ""
}

var tmplFuncMap = template.FuncMap{
	"strings_Title":   strings.Title,
	"strings_ToLower": strings.ToLower,
	"strings_ToUpper": strings.ToUpper,

	"strings_TrimPrefix": strings.TrimPrefix,
	"strings_TrimSuffix": strings.TrimSuffix,
}

const tmplFileHeader = `// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: {{.GetFileName}}

package {{.GetPackageName}}

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"
import "github.com/chai2010/qingcloud-go/pkg/config"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = config.Config{}
var _ = client.NewClient

`

const tmplService = `
{{$service := .}}

type {{.GetServiceName}}Interface interface {
	{{- range $_, $m := .GetMethodList}}
		{{$m.GetMethodName}}(in *{{$m.GetInputTypeName}}) (out *{{$m.GetOutputTypeName}}, err error)
	{{- end}}
}

type {{.GetServiceName}} struct {
	Config           *config.Config
	Properties       *{{.GetServiceName}}Properties
	LastResponseBody string
}

func New{{.GetServiceName}}(conf *config.Config, zone string) (p *{{.GetServiceName}}) {
	return &{{.GetServiceName}}{
		Config:     conf,
		Properties: &{{.GetServiceName}}Properties{ Zone: proto.String(zone) },
	}
}

{{range $_, $m := .GetMethodList}}
func (p *{{$service.GetServiceName}}) {{$m.GetMethodName}}(input *{{$m.GetInputTypeName}}) (output *{{$m.OutputTypeName}}, err error) {
	client := client.NewClient("", "", nil)
	output = new({{$m.GetOutputTypeName}})

	err = client.CallMethod(nil, "{{$m.GetMethodName}}", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
{{end}}
`
