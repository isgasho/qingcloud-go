// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	rule_pb "github.com/chai2010/qingcloud-go/spec.pb/qingcloud_sdk_rule"
)

type ServiceSpec struct {
	MainServiceName string
	SubServiceName  string
	MethodList      []MethodSpec
	*rule_pb.ServiceOptionsRule
}

type MethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
	*rule_pb.MethodOptionsRule
}

type MessageSpec struct {
	MessageTypeName string
	*rule_pb.MessageOptionsRule
}

const tmplService = `
{{/* TODO */}}
`

const tmplMessageValidate = `
{{/* TODO */}}
`
