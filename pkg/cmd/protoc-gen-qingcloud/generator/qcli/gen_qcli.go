// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package gen_qcli

import (
	"bytes"
	"fmt"
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

func (p pkgGenerator) Name() string        { return "qcli" }
func (p pkgGenerator) FileNameExt() string { return ".pb.qcli.go" }

func (pkgGenerator) HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplFileHeader))
	err := t.Execute(&buf,
		struct {
			G    *generator.Generator
			File *generator.FileDescriptor
		}{
			G:    g,
			File: file,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func (pkgGenerator) ServiceCode(g *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplService))
	err := t.Execute(&buf,
		struct {
			G       *generator.Generator
			File    *generator.FileDescriptor
			Service *descriptor.ServiceDescriptorProto
		}{
			G:       g,
			File:    file,
			Service: svc,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func (pkgGenerator) MessageCode(p *generator.Generator, file *generator.FileDescriptor, msg *descriptor.DescriptorProto) string {
	return ""
}

var tmplFuncMap = template.FuncMap{
	"strings_Title":      strings.Title,
	"strings_ToLower":    strings.ToLower,
	"strings_ToUpper":    strings.ToUpper,
	"strings_TrimPrefix": strings.TrimPrefix,
	"strings_TrimSuffix": strings.TrimSuffix,

	"generator_CamelCase":             generator.CamelCase,
	"utils_GetMethodInputDescriptor":  utils.GetMethodInputDescriptor,
	"utils_GetMethodOutputDescriptor": utils.GetMethodOutputDescriptor,

	"pkgGenCmdFlagType":                pkgGenCmdFlagType,
	"pkgGenCmdFlagName":                pkgGenCmdFlagName,
	"pkgGenCmdFlagUsage":               pkgGenCmdFlagUsage,
	"pkgGenCmdFlagDefaultValue":        pkgGenCmdFlagDefaultValue,
	"pkgGenCmdFlagDefaultValueComment": pkgGenCmdFlagDefaultValueComment,
	"pkgGenCmdFlagFillField":           pkgGenCmdFlagFillField,
}

func pkgGenCmdFlagType(field *descriptor.FieldDescriptorProto) string {
	switch {
	case utils.IsSupportedBool(field):
		return "BoolFlag"
	case utils.IsSupportedInt(field):
		return "IntFlag"
	case utils.IsSupportedFloat(field):
		return "Float64Flag"
	case utils.IsSupportedString(field):
		return "StringFlag"
	default: // json: slice/message/map
		return "StringFlag"
	}
}

func pkgGenCmdFlagName(field *descriptor.FieldDescriptorProto) string {
	return field.GetName()
}
func pkgGenCmdFlagUsage(field *descriptor.FieldDescriptorProto) string {
	s := field.GetName()
	s = strings.Replace(s, "_", " ", -1)
	return s
}

func pkgGenCmdFlagDefaultValue(field *descriptor.FieldDescriptorProto) string {
	switch {
	case utils.IsSupportedBool(field):
		return "false"
	case utils.IsSupportedInt(field):
		return "0"
	case utils.IsSupportedFloat(field):
		return "0"
	case utils.IsSupportedString(field):
		return `""`
	default: // json: slice/message/map
		return `""`
	}
}
func pkgGenCmdFlagDefaultValueComment(field *descriptor.FieldDescriptorProto) string {
	switch {
	case utils.IsSupportedBool(field):
		return ""
	case utils.IsSupportedInt(field):
		return ""
	case utils.IsSupportedFloat(field):
		return ""
	case utils.IsSupportedString(field):
		return ""
	default:
		return `// json: slice/message/map/time`
	}
}

func pkgGenCmdFlagFillField(field *descriptor.FieldDescriptorProto) string {
	if utils.IsSupportedRepeated(field) {
		return fmt.Sprintf(
			`if err := json.Unmarshal([]byte(c.String("%s")), &in.%s); err != nil {
				log.Fatal(err)
			}`,
			field.GetName(),
			generator.CamelCase(field.GetName()),
		)
	}
	switch {
	case utils.IsSupportedBool(field):
		return fmt.Sprintf(
			`in.%s = proto.Bool(c.Bool("%s"))`,
			generator.CamelCase(field.GetName()),
			field.GetName(),
		)
	case utils.IsSupportedInt(field):
		switch _type := field.GetType(); _type {
		case descriptor.FieldDescriptorProto_TYPE_INT32,
			descriptor.FieldDescriptorProto_TYPE_SINT32,
			descriptor.FieldDescriptorProto_TYPE_FIXED32:
			return fmt.Sprintf(
				`in.%s = proto.Int32(int32(c.Int("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		case descriptor.FieldDescriptorProto_TYPE_INT64,
			descriptor.FieldDescriptorProto_TYPE_SINT64,
			descriptor.FieldDescriptorProto_TYPE_FIXED64:
			return fmt.Sprintf(
				`in.%s = proto.Int64(int64(c.Int64("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		case descriptor.FieldDescriptorProto_TYPE_UINT32:
			return fmt.Sprintf(
				`in.%s = proto.Uint32(uint32(c.Int("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		case descriptor.FieldDescriptorProto_TYPE_UINT64:
			return fmt.Sprintf(
				`in.%s = proto.Uint64(uint64(c.Int64("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		default:
			return fmt.Sprintf("unknown type: %v", _type)
		}
	case utils.IsSupportedFloat(field):
		switch _type := field.GetType(); _type {
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			return fmt.Sprintf(
				`in.%s = proto.Float32(float32(c.Float64("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			return fmt.Sprintf(
				`in.%s = proto.Float64(float64(c.Float64("%s")))`,
				generator.CamelCase(field.GetName()),
				field.GetName(),
			)
		default:
			return fmt.Sprintf("unknown type: %v", _type)
		}
	case utils.IsSupportedString(field):
		return fmt.Sprintf(
			`in.%s = proto.String(c.String("%s"))`,
			generator.CamelCase(field.GetName()),
			field.GetName(),
		)
	default:
		return fmt.Sprintf(
			`if err := json.Unmarshal([]byte(c.String("%s")), &in.%s); err != nil {
				log.Fatal(err)
			}`,
			field.GetName(),
			generator.CamelCase(field.GetName()),
		)
	}
}

const tmplFileHeader = `
{{- $G := .G -}}
{{- $File := .File -}}

// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: {{$File.GetName}}

package qcli_pb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = json.Marshal
	_ = log.Print
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = pb.AlarmService{}
)

`

const tmplService = `
{{- $G := .G -}}
{{- $File := .File -}}
{{- $Service := .Service -}}

{{- $ServiceName := generator_CamelCase $Service.GetName -}}
{{- $ServiceCommandName := strings_TrimSuffix $ServiceName "Service" | strings_ToLower -}}

func init() {
	AllCommands = append(AllCommands, Cmd{{$ServiceName}})
}

var Cmd{{$ServiceName}} = cli.Command{
	Name:    "{{$ServiceCommandName}}",
	Aliases: []string{},
	Usage:   "manage {{$ServiceName}}",
	Subcommands: []cli.Command{
		{{range $_, $Method := $Service.Method -}}
			{{- $MethodName := generator_CamelCase $Method.GetName -}}
			{{- $MethodInput := utils_GetMethodInputDescriptor $G $Method -}}
			{{- $MethodOutput := utils_GetMethodOutputDescriptor $G $Method -}}

			{{- $MethodInputName := generator_CamelCase $MethodInput.GetName -}}
			{{- $MethodOutputName := generator_CamelCase $MethodOutput.GetName -}}

			{
				Name:    "{{$MethodName}}",
				Aliases: []string{},
				Usage:   "{{$MethodName}}",
				Flags:   _flag_{{$ServiceName}}_{{$MethodName}},
				Action:  _func_{{$ServiceName}}_{{$MethodName}},
			},
		{{end}}
	},
}

{{/* generate command flags and functions */}}
{{range $_, $Method := $Service.Method}}

{{- $MethodName := generator_CamelCase $Method.GetName -}}
{{- $MethodInput := utils_GetMethodInputDescriptor $G $Method -}}
{{- $MethodOutput := utils_GetMethodOutputDescriptor $G $Method -}}

{{- $MethodInputName := generator_CamelCase $MethodInput.GetName -}}
{{- $MethodOutputName := generator_CamelCase $MethodOutput.GetName -}}

var _flag_{{$ServiceName}}_{{$MethodName}} = []cli.Flag{
	{{range $_, $MethodInputField := $MethodInput.Field -}}
		cli.{{pkgGenCmdFlagType $MethodInputField}}{
			Name:  "{{pkgGenCmdFlagName $MethodInputField}}",
			Usage: "{{pkgGenCmdFlagUsage $MethodInputField}}",

			{{- if $v := pkgGenCmdFlagDefaultValue $MethodInputField}}
			Value: {{$v}}, {{pkgGenCmdFlagDefaultValueComment $MethodInputField}}
			{{- end}}
		},
	{{end}}
}

func _func_{{$ServiceName}}_{{$MethodName}}(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.New{{$ServiceName}}(conf, zone)

	in := new(pb.{{$MethodInputName}})

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		{{range $_, $MethodInputField := $MethodInput.Field -}}
			if c.IsSet("{{pkgGenCmdFlagName $MethodInputField}}") {
				{{pkgGenCmdFlagFillField $MethodInputField}}
			}
		{{end -}}
	}

	out, err := qc.{{$MethodName}}(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}
{{end}}
`
