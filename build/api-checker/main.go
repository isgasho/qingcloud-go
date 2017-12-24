// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func main() {
	for typeName, typeType := range pb.TypeInfoMap {
		for i := 0; i < typeType.NumField(); i++ {
			_ = typeType.Field(i).Type
			_ = typeType.Field(i).Tag
		}
		_, _ = typeName, typeType
	}
}
