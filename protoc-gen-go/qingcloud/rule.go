// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	rule_pb "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
)

type pbServiceRule struct{ *rule_pb.ServiceRule }

func (p *pbServiceRule) IsMainService() bool        { return p.ServiceRule.GetSubServiceName() == "" }
func (p *pbServiceRule) IsSubService() bool         { return p.ServiceRule.GetSubServiceName() != "" }
func (p *pbServiceRule) GetMainServiceName() string { return p.ServiceRule.GetServiceName() }
func (p *pbServiceRule) GetSubServiceName() string  { return p.ServiceRule.GetSubServiceName() }
func (p *pbServiceRule) GetDocUrl() string          { return p.ServiceRule.GetDocUrl() }

type pbMethodRule struct{ *rule_pb.MethodRule }

func (p *pbMethodRule) GetHttpMethod() string {
	if s := p.MethodRule.GetHttpAction(); s != "" {
		return s
	}
	return "GET"
}

type pbMethodInputRule struct{ *rule_pb.MethodInputRule }

func (p *pbMethodInputRule) IsRequired(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) IsNumberType(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) IsStringType(filedName string) bool {
	return false
}

func (p *pbMethodInputRule) HasEnumValue(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) HasDefaultValue(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) HasMaxValue(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) HasMinValue(filedName string) bool {
	return false
}
func (p *pbMethodInputRule) HasMultipleOfValue(filedName string) bool {
	return false
}

func (p *pbMethodInputRule) GetRequiredFiledList() []string {
	return nil
}
func (p *pbMethodInputRule) GetEnumValueList() (names []string, values [][]string) {
	return nil, nil
}
func (p *pbMethodInputRule) GetDefaultValueList() (names []string, values []string) {
	return nil, nil
}
func (p *pbMethodInputRule) GetMaxValueList() (names []string, values []string) {
	return nil, nil
}
func (p *pbMethodInputRule) GetMinValueList() (names []string, values []string) {
	return nil, nil
}
func (p *pbMethodInputRule) GetMultipleOfValueList() (names []string, values []string) {
	return nil, nil
}
