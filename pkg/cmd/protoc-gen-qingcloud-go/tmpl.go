// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud_plugin

import (
	"bytes"
	"text/template"
)

type FileSpec struct {
	IsProto3    bool
	FileName    string
	PackageName string
}

type ServiceSpec struct {
	IsProto3        bool
	DocUrl          string
	ServiceName     string
	MainServiceName string
	MethodList      []MethodSpec
}

type MethodSpec struct {
	DocUrl         string
	MethodName     string
	InputTypeName  string
	OutputTypeName string
	HttpMethod     string
}

type MessageOptionsSpec struct {
	IsProto3           bool
	MessageName        string
	FieldNameMap       map[string]string
	RequiredFieldMap   map[string]bool
	RepeatedFieldMap   map[string]bool
	DefaultValueMap    map[string]string
	EnumValueListMap   map[string][]string
	MinValueMap        map[string]string
	MaxValueMap        map[string]string
	MultipleOfValueMap map[string]string
	RegexpValueMap     map[string]string
	FiledTypeMap       map[string]string // bool/int/float/string/bytes/message/?
}

func NewMessageOptionsSpec() *MessageOptionsSpec {
	return &MessageOptionsSpec{
		FieldNameMap:       map[string]string{},
		RequiredFieldMap:   map[string]bool{},
		RepeatedFieldMap:   map[string]bool{},
		DefaultValueMap:    map[string]string{},
		EnumValueListMap:   map[string][]string{},
		MinValueMap:        map[string]string{},
		MaxValueMap:        map[string]string{},
		MultipleOfValueMap: map[string]string{},
		RegexpValueMap:     map[string]string{},
		FiledTypeMap:       map[string]string{},
	}
}

func (spec *FileSpec) HeaderCode() string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplFile))
	t.Execute(&buf, spec)
	return buf.String()
}

func (spec *ServiceSpec) Code() string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	t.Execute(&buf, spec)
	return buf.String()
}

func (spec *MessageOptionsSpec) HasOptions() bool {
	if len(spec.RequiredFieldMap) > 0 {
		return true
	}
	if len(spec.RepeatedFieldMap) > 0 {
		// skip
	}
	if len(spec.DefaultValueMap) > 0 {
		return true
	}
	if len(spec.EnumValueListMap) > 0 {
		return true
	}
	if len(spec.MinValueMap) > 0 {
		return true
	}
	if len(spec.MaxValueMap) > 0 {
		return true
	}
	if len(spec.MultipleOfValueMap) > 0 {
		return true
	}
	if len(spec.RegexpValueMap) > 0 {
		return true
	}
	if len(spec.FiledTypeMap) > 0 {
		// skip
	}

	return false
}

// ----------------------------------------------------------------------------

const tmplFile = `// Code generated by protoc-gen-qingcloud-go. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go
// source: {{.FileName}}

package {{.PackageName}}

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
var _ = logger.SetLevel
var _ = request.Request{}
var _ = data.Operation{}

`

// ----------------------------------------------------------------------------

const tmplService = `
{{$service := .}}

{{if .DocUrl}}// See {{.DocUrl}}{{end}}
type {{.ServiceName}}ServiceInterface interface {
	{{- range $_, $m := .MethodList}}
		{{$m.MethodName}}(in *{{$m.InputTypeName}}) (out *{{$m.OutputTypeName}}, err error)
	{{- end}}
}

{{if .DocUrl}}// See {{.DocUrl}}{{end}}
type {{.ServiceName}}Service struct {
	Config           *config.Config
	Properties       *{{.ServiceName}}ServiceProperties
	LastResponseBody string
}

{{if not .MainServiceName}}
func Init(c *config.Config) (*{{.ServiceName}}Service, error) {
	properties := &{{.ServiceName}}ServiceProperties{}
	logger.SetLevel(c.LogLevel)
	return &{{.ServiceName}}Service{Config: c, Properties: properties}, nil
}
{{end}}

{{if .MainServiceName}}
{{if .DocUrl}}// See {{.DocUrl}}{{end}}
func New{{.ServiceName}}Service(conf *config.Config, zone string) (p *{{.ServiceName}}Service) {
	return &{{.ServiceName}}Service{
		Config:     conf,
		Properties: &{{.ServiceName}}ServiceProperties{ Zone: {{if .IsProto3}}zone{{else}}proto.String(zone){{end}} },
	}
}

{{if .DocUrl}}// See {{.DocUrl}}{{end}}
func (s *{{.MainServiceName}}Service) {{.ServiceName}}(zone string) (*{{.ServiceName}}Service, error) {
	properties := &{{.ServiceName}}ServiceProperties{
		Zone: {{if .IsProto3}}zone{{else}}proto.String(zone){{end}},
	}

	return &{{.ServiceName}}Service{Config: s.Config, Properties: properties}, nil
}
{{end}}

{{range $_, $m := .MethodList}}
{{if .DocUrl}}// See {{.DocUrl}}{{end}}
func (p *{{$service.ServiceName}}Service) {{$m.MethodName}}(in *{{$m.InputTypeName}}) (out *{{$m.OutputTypeName}}, err error) {
	if in == nil {
		in = &{{$m.InputTypeName}}{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "{{$m.MethodName}}",
		RequestMethod: "{{$m.HttpMethod}}", // GET or POST
	}

	x := &{{$m.OutputTypeName}}{}
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

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
