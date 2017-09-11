// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"strings"

	rule_pb "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
)

type pbServiceRule struct{ *rule_pb.ServiceOptionsRule }

func (p *pbServiceRule) IsMainService() bool        { return p.ServiceOptionsRule.GetSubServiceName() == "" }
func (p *pbServiceRule) IsSubService() bool         { return p.ServiceOptionsRule.GetSubServiceName() != "" }
func (p *pbServiceRule) GetMainServiceName() string { return p.ServiceOptionsRule.GetServiceName() }
func (p *pbServiceRule) GetSubServiceName() string  { return p.ServiceOptionsRule.GetSubServiceName() }
func (p *pbServiceRule) GetDocUrl() string          { return p.ServiceOptionsRule.GetDocUrl() }

type pbMethodRule struct{ *rule_pb.MethodOptionsRule }

func (p *pbMethodRule) GetHttpMethod() string {
	if s := p.MethodOptionsRule.GetHttpMethod(); s != "" {
		return s
	}
	return "GET"
}

type pbMethodInputRule struct{ *rule_pb.MessageOptionsRule }

func (p *pbMethodInputRule) IsRequired(filedName string) bool {
	for _, s := range p.getRequiredFiledList() {
		if s == filedName {
			return true
		}
	}
	return false
}

func (p *pbMethodInputRule) HasDefaultValue(filedName string) bool {
	return len(p.GetDefaultValue(filedName)) > 0
}
func (p *pbMethodInputRule) GetDefaultValue(filedName string) string {
	names, values := p.getDefaultValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *pbMethodInputRule) HasEnumValue(filedName string) bool {
	return len(p.GetEnumValue(filedName)) > 0
}
func (p *pbMethodInputRule) GetEnumValue(filedName string) []string {
	names, values := p.getEnumValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return nil
}

func (p *pbMethodInputRule) HasMinValue(filedName string) bool {
	return len(p.GetMinValue(filedName)) > 0
}
func (p *pbMethodInputRule) GetMinValue(filedName string) string {
	names, values := p.getMinValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *pbMethodInputRule) HasMaxValue(filedName string) bool {
	return len(p.GetMaxValue(filedName)) > 0
}
func (p *pbMethodInputRule) GetMaxValue(filedName string) string {
	names, values := p.getMaxValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

func (p *pbMethodInputRule) HasMultipleOfValue(filedName string) bool {
	return len(p.GetMultipleOfValue(filedName)) > 0
}
func (p *pbMethodInputRule) GetMultipleOfValue(filedName string) string {
	names, values := p.getMultipleOfValueList()
	for i, name := range names {
		if filedName == name {
			return values[i]
		}
	}
	return ""
}

// data: "a; b; ..."
func (p *pbMethodInputRule) getRequiredFiledList() []string {
	return splitString(p.MessageOptionsRule.RequiredFileds, ";")
}

// data: "a:v; b:v; ..."
func (p *pbMethodInputRule) getDefaultValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.DefaultValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:a1,a2,a3; b:b1,b2; ..."
func (p *pbMethodInputRule) getEnumValueList() (names []string, values [][]string) {
	for _, s := range splitString(p.MessageOptionsRule.EnumValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, splitString(kv[1], ","))
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *pbMethodInputRule) getMinValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MinValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *pbMethodInputRule) getMaxValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MaxValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
		}
	}
	return
}

// data: "a:v; b:v; ..."
func (p *pbMethodInputRule) getMultipleOfValueList() (names []string, values []string) {
	for _, s := range splitString(p.MessageOptionsRule.MultipleOfValue, ";") {
		if kv := splitString(s, ":"); len(kv) == 2 {
			names = append(names, kv[0])
			values = append(values, kv[1])
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
