// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package golang_validate_v1

import (
	"bytes"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"

	spec_metadata "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"
	plugin "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud-validate"
)

func init() {
	plugin.RegisterValidateGenerator(new(pkgGenerator))
}

type pkgGenerator struct{}

func (p pkgGenerator) Name() string        { return "golang" }
func (p pkgGenerator) FileNameExt() string { return ".pb.validate.go" }

func (p pkgGenerator) HeaderCode(file *generator.FileDescriptor) string {
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

func (p pkgGenerator) ValidateCode(file *generator.FileDescriptor, message *descriptor.DescriptorProto) string {
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

	fieldOption := getMessageFieldOption(fieldDescriptor)
	if fieldOption == nil {
		return ""
	}

	switch {
	case pkgIsSupportedBool(fieldDescriptor):
	case pkgIsSupportedInt(fieldDescriptor):
	case pkgIsSupportedFloat(fieldDescriptor):
	case pkgIsSupportedString(fieldDescriptor):
	}

	// int type
	if pkgIsSupportedInt(fieldDescriptor) {
		//
	}

	// float type
	if pkgIsSupportedFloat(fieldDescriptor) {
		//
	}

	// string
	if pkgIsSupportedString(fieldDescriptor) {
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
