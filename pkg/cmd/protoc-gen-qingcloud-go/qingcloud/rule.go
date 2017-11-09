// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"strings"

	rule_pb "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
)

type FieldRule struct {
	*rule_pb.FieldOptionsRule
}

type ServiceRule struct {
	*rule_pb.ServiceOptionsRule
}

func (p *ServiceRule) GetDocUrl() string {
	return p.ServiceOptionsRule.DocUrl
}
func (p *ServiceRule) GetServiceName() string {
	name := p.ServiceOptionsRule.ServiceName
	if idx := strings.Index(name, "."); idx >= 0 {
		return name[idx+1:]
	}
	return name
}
func (p *ServiceRule) GetMainServiceName() string {
	name := p.ServiceOptionsRule.ServiceName
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		return name[:idx]
	}
	return ""
}

type MethodRule struct {
	*rule_pb.MethodOptionsRule
}

func (p *MethodRule) GetDocUrl() string {
	return p.MethodOptionsRule.DocUrl
}
func (p *MethodRule) GetHttpMethod() string {
	return p.MethodOptionsRule.HttpMethod
}
func (p *MethodRule) GetInputType() string {
	return p.MethodOptionsRule.InputType
}
func (p *MethodRule) GetOutputType() string {
	return p.MethodOptionsRule.OutputType
}

type MessageRule struct {
	*rule_pb.MessageOptionsRule
}

func (p *MessageRule) IsRequired(filedName string) bool {
	for _, s := range p.getRequiredFiledList() {
		if s == filedName {
			return true
		}
	}
	return false
}

func (p *MessageRule) HasDefaultValue(filedName string) bool {
	return len(p.GetDefaultValue(filedName)) > 0
}
func (p *MessageRule) GetDefaultValue(filedName string) string {
	names, values := p.getDefaultValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *MessageRule) HasEnumValue(filedName string) bool {
	return len(p.GetEnumValueList(filedName)) > 0
}
func (p *MessageRule) GetEnumValueList(filedName string) []string {
	names, values := p.getEnumValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return nil
}

func (p *MessageRule) HasMinValue(filedName string) bool {
	return len(p.GetMinValue(filedName)) > 0
}
func (p *MessageRule) GetMinValue(filedName string) string {
	names, values := p.getMinValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *MessageRule) HasMaxValue(filedName string) bool {
	return len(p.GetMaxValue(filedName)) > 0
}
func (p *MessageRule) GetMaxValue(filedName string) string {
	names, values := p.getMaxValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *MessageRule) HasMultipleOfValue(filedName string) bool {
	return len(p.GetMultipleOfValue(filedName)) > 0
}
func (p *MessageRule) GetMultipleOfValue(filedName string) string {
	names, values := p.getMultipleOfValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *MessageRule) HasRegexpValue(filedName string) bool {
	return len(p.GetRegexpValue(filedName)) > 0
}
func (p *MessageRule) GetRegexpValue(filedName string) string {
	names, values := p.getRegexpValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

// data: "a; b; ..."
func (p *MessageRule) getRequiredFiledList() []string {
	return splitString(p.MessageOptionsRule.RequiredFileds, ";")
}

// data: "a:v; b:v; ..."
func (p *MessageRule) getDefaultValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.DefaultValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:a1,a2,a3; b:b1,b2; ..."
func (p *MessageRule) getEnumValueList() (names []string, values [][]string) {
	for _, s := range splitString(p.MessageOptionsRule.EnumValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, splitString(kv[1], ","))
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *MessageRule) getMinValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MinValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *MessageRule) getMaxValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MaxValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *MessageRule) getMultipleOfValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MultipleOfValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// "a:{{...}}; b:{{...}};"
func (p *MessageRule) getRegexpValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MultipleOfValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, strings.TrimSuffix(strings.TrimPrefix(kv[1], "{{"), "}}"))
		}
	}
	return
}

func splitString(s, sep string) []string {
	var ss []string
	for _, s := range strings.Split(s, sep) {
		if s := strings.TrimSpace(s); s != "" {
			ss = append(ss, s)
		}
	}
	return ss
}
