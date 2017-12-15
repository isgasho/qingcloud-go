// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package service

import (
	"reflect"
)

type ServiceApiSpec struct {
	ActionName string
	InputType  reflect.Type
	OutputType reflect.Type
	HttpMethod string
}

var ServiceApiSpecMap = map[string]ServiceApiSpec{}
