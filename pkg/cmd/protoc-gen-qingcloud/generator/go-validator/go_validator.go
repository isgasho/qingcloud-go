// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package golang_validator

import (
	"bytes"
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

func (p pkgGenerator) Name() string        { return "go-validator" }
func (p pkgGenerator) FileNameExt() string { return ".pb.validate.go" }

func (p pkgGenerator) HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplFileHeader))
	t.Execute(&buf,
		struct {
			File *generator.FileDescriptor
		}{
			File: file,
		},
	)
	return buf.String()
}

func (p pkgGenerator) ServiceCode(g *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string {
	return ""
}

func (p pkgGenerator) MessageCode(g *generator.Generator, file *generator.FileDescriptor, message *descriptor.DescriptorProto) string {
	var fieldValidate []string
	for _, field := range message.GetField() {
		if s := p.FieldValidateCode(file, message, field); s != "" {
			fieldValidate = append(fieldValidate, s)
		}
	}

	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(tmplFuncMap).Parse(tmplValidate))
	t.Execute(&buf,
		struct {
			File          *generator.FileDescriptor
			Message       *descriptor.DescriptorProto
			FieldValidate []string
		}{
			File:          file,
			Message:       message,
			FieldValidate: fieldValidate,
		},
	)
	return buf.String()
}
func (p pkgGenerator) FieldValidateCode(
	fileDescriptor *generator.FileDescriptor,
	messageDescriptor *descriptor.DescriptorProto,
	fieldDescriptor *descriptor.FieldDescriptorProto,
) string {
	fieldName := generator.CamelCase(fileDescriptor.GetName())

	fieldOption := utils.GetMessageFieldOption(fieldDescriptor)
	if fieldOption == nil {
		return ""
	}

	switch {
	case utils.IsSupportedBool(fieldDescriptor):
	case utils.IsSupportedInt(fieldDescriptor):
	case utils.IsSupportedFloat(fieldDescriptor):
	case utils.IsSupportedString(fieldDescriptor):
	}

	// int type
	if utils.IsSupportedInt(fieldDescriptor) {
		//
	}

	// float type
	if utils.IsSupportedFloat(fieldDescriptor) {
		//
	}

	// string
	if utils.IsSupportedString(fieldDescriptor) {
		//
	}

	// map time
	if messageDescriptor.GetOptions().GetMapEntry() {
		//
	}

	//msg.GetOptions().GetMapEntry()

	return "// Field: " + fieldName
}

/*

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		validator := getFieldValidatorIfAny(field)
		if validator != nil && validator.Regex != nil {
			fieldName := p.GetFieldName(message, field)
			p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, strconv.Quote(*validator.Regex), `)`)
		}
	}
}
*/
