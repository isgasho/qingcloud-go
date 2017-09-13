// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

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

type MessageSpec struct {
	MessageTypeName string
	RequiredFileds  []string
	DefaultValue    map[string]string
	EnumValue       map[string][]string
	MinValue        map[string]string
	MaxValue        map[string]string
	MultipleOfValue map[string]string
	RegexpValue     map[string]string
}

const tmplService = `

{{if .DocUrl}}// See {{.DocUrl}}{{end}}
type {{.ServiceName}}Service struct {
	Config           *config.Config
	Properties       *{{.ServiceName}}Properties
	LastResponseBody string
}

func New{{.ServiceName}}Service(conf *config.Config, zone string) (p *{{.ServiceName}}Service) {
	return &{{.ServiceName}}Service{
		Config:     conf,
		Properties: &{{.ServiceName}}Properties{Zone: zone},
	}
}

{{if not .MainServiceName}}
func (s *{{.MainServiceName}}Service) {{.ServiceName}}(zone string) (*{{.ServiceName}}Service, error) {
	properties := &{{.ServiceName}}ServiceProperties{
		Zone: zone,
	}

	return &{{.ServiceName}}{Config: s.Config, Properties: properties}, nil
}
{{end}}

{{range $_, $m := .MethodList}}
{{if .DocUrl}}// See {{.DocUrl}}{{end}}
func (p *{{.ServiceName}}Service) {{$m.MethodName}}(in *{{$m.InputTypeName}}) (out *{{$m.OutputTypeName}}, err error) {
	if in == nil {
		in = &{{$m.InputTypeName}}{}
	}
	o := &data.Operation{
		Config:        p.Config,
		Properties:    p.Properties,
		APIName:       "{{$m.MethodName}}",
		RequestMethod: "{{$m.HttpMethod}}",
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

const tmplMessageValidate = `
func (p *{{.MessageTypeName}}) Validate() error {
	return nil
}
`
