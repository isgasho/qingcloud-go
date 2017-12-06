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
import fmt "fmt"
import math "math"
import "regexp"

import "github.com/chai2010/qingcloud-go/pkg/config"
import "github.com/chai2010/qingcloud-go/pkg/logger"
import "github.com/chai2010/qingcloud-go/pkg/request"
import "github.com/chai2010/qingcloud-go/pkg/request/data"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = regexp.Match
var _ = config.Config{}
var _ = logger.Info
var _ = request.Request{}
var _ = data.Operation{}

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

func (s *QingCloudService) {{strings_TrimSuffix .GetServiceName "Service"}}(zone string) (*{{.GetServiceName}}, error) {
	properties := &{{.GetServiceName}}Properties{
		Zone: proto.String(zone),
	}

	return &{{.GetServiceName}}{
		Config: s.Config,
		Properties: properties,
	}, nil
}

{{range $_, $m := .GetMethodList}}
func (p *{{$service.GetServiceName}}) {{$m.GetMethodName}}(in *{{$m.GetInputTypeName}}) (out *{{$m.OutputTypeName}}, err error) {
	if in == nil {
		in = &{{$m.GetInputTypeName}}{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "{{$m.GetMethodName}}",
		RequestMethod: "{{$m.GetHttpMethod}}",
	}

	x := &{{$m.GetOutputTypeName}}{}
	r, err := request.New(o, in, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	p.LastResponseBody = o.ResponseBody

	if err != nil {
		return nil, err
	}

	return x, err
}
{{end}}
`