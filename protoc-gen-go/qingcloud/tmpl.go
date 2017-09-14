// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"bytes"
	"text/template"
)

type ServiceSpec struct {
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

func GetImportsCode() string {
	return tmplImports
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

func (spec *MessageOptionsSpec) ValidateCode() string {
	if !spec.HasOptions() {
		var buf bytes.Buffer
		t := template.Must(template.New("").Parse(tmplMessageValidateEmpty))
		t.Execute(&buf, spec)
		return buf.String()
	}

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplMessageValidate))
	t.Execute(&buf, spec)
	return buf.String()
}

// ----------------------------------------------------------------------------

const tmplImports = `
import "regexp"

import "github.com/chai2010/qingcloud-go/config"
import "github.com/chai2010/qingcloud-go/logger"
import "github.com/chai2010/qingcloud-go/request"
import "github.com/chai2010/qingcloud-go/request/data"

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
		Properties: &{{.ServiceName}}ServiceProperties{Zone: zone},
	}
}

{{if .DocUrl}}// See {{.DocUrl}}{{end}}
func (s *{{.MainServiceName}}Service) {{.ServiceName}}(zone string) (*{{.ServiceName}}Service, error) {
	properties := &{{.ServiceName}}ServiceProperties{
		Zone: zone,
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

const tmplMessageValidateEmpty = `
func (p *{{.MessageName}}) Validate() error {
	return nil
}
`

const tmplMessageValidate = `
{{$msg := .}}

func (p *{{$msg.MessageName}}) Validate() error {
	{{/* required fields */}}
	{{range $k, $_ := $msg.RequiredFieldMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if or $isRepeaded (eq $type "string") (eq $type "bytes") -}}
			if len(p.{{$name}}) == 0 {
				return fmt.Errorf("{{$msg.MessageName}}.{{$name}} required field missing!")
			}
		{{else}}
			{{if eq $type "message"}}
				if {{$msg.MessageName}}.{{$name}} == nil {
					return fmt.Errorf("{{$msg.MessageName}}.{{$name}} required field missing!")
				}
			{{end}}
		{{end}}
	{{end}}

	{{/* enum fields */}}
	{{range $k, $v := $msg.EnumValueListMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if $isRepeaded}}
			{{if eq $type "int"}}
				{
					var _enumValues = []int{
						{{range $_, $vi := $v -}}
							{{$vi}},
						{{end}}
					}
					var found = false
					for _, v := range _enumValues {
						if int(p.{{$name}}) == v {
							found = true
						}
						if !found {
							return fmt.Errorf("{{$msg.MessageName}}.{{$name}} invalid enum value!")
						}
					}
				}
			{{end}}
			{{if eq $type "string"}}
				{
					var _enumValues = []string{
						{{range $_, $vi := $v -}}
							"{{$vi}}",
						{{end}}
					}
					for _, v := range _enumValues {
						if p.{{$name}} != v {
							return fmt.Errorf("{{$msg.MessageName}}.{{$name}} invalid enum value!")
						}
					}
				}
			{{end}}
		{{else}}
			{{if eq $type "int"}}
				{
					var _enumValues = []int{
						{{range $_, $vi := $v -}}
							{{$vi}},
						{{end}}
					}
					for _, v := range _enumValues {
						if int(p.{{$name}}) != v {
							return fmt.Errorf("{{$msg.MessageName}}.{{$name}} invalid enum value!")
						}
					}
				}
			{{end}}
			{{if eq $type "string"}}
				{
					var _enumValues = []string{
						{{range $_, $vi := $v -}}
							"{{$vi}}",
						{{end}}
					}
					for _, v := range _enumValues {
						if p.{{$name}} != v {
							return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: invalid enum value!")
						}
					}
				}
			{{end}}
		{{end}}
	{{end}}

	{{/* min_value */}}
	{{range $k, $v := $msg.MinValueMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if $isRepeaded}}
			{{if eq $type "int"}}
				for _, vi := range {{$v}} {
					if int(p.{{$name}}) < vi {
						return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check min_value failed!")
					}
				}
			{{end}}
		{{else}}
			{{if eq $type "int"}}
				if int(p.{{$name}}) < {{$v}} {
					return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check min_value failed!")
				}
			{{end}}
		{{end}}
	{{end}}

	{{/* max_value */}}
	{{range $k, $v := $msg.MaxValueMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if $isRepeaded}}
			{{if eq $type "int"}}
				for _, vi := range {{$v}} {
					if int(vi) > {{$v}} {
						return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check max_value failed!")
					}
				}
			{{end}}
		{{else}}
			{{if eq $type "int"}}
				if int(p.{{$name}}) > {{$v}} {
					return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check max_value failed!")
				}
			{{end}}
		{{end}}
	{{end}}

	{{/* multiple_of_value */}}
	{{range $k, $v := $msg.MaxValueMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if $isRepeaded}}
			{{if eq $type "int"}}
				for _, vi := range {{$v}} {
					if int(p.{{$name}}) % {{$v}} != 0 {
						return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check multiple_of_value failed!")
					}
				}
			{{end}}
		{{else}}
			{{if eq $type "int"}}
				if int(p.{{$name}}) % {{$v}} != 0 {
					return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check multiple_of_value failed!")
				}
			{{end}}
		{{end}}
	{{end}}

	{{/* regexp_value */}}
	{{range $k, $v := $msg.RegexpValueMap}}
		{{$name := index $msg.FieldNameMap $k}}
		{{$type := index $msg.FiledTypeMap $k}}
		{{$isRepeaded := index $msg.RepeatedFieldMap $k}}

		{{if $isRepeaded}}
			{{if eq $type "string"}}
				for _, vi := range {{$v}} {
					if ok, err := regexp.MatchString("{{$v}}", p.{{$name}}); err != nil || !ok {
						return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check regexp failed!")
					}
				}
			{{end}}
		{{else}}
			{{if eq $type "string"}}
				if ok, err := regexp.MatchString("{{$v}}", p.{{$name}}); err != nil || !ok {
					return fmt.Errorf("{{$msg.MessageName}}.{{$name}}: check regexp failed!")
				}
			{{end}}
		{{end}}
	{{end}}

	return nil
}
`

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
