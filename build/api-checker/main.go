// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate go run gen_helper.go
//go:generate go fmt

package main

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
)

var typeInfoMap = map[string]interface{}{}

func main() {
	for typeName, typeValue := range typeInfoMap {
		for _, field := range structs.Fields(typeValue) {
			originName := strings.Split(field.Tag("json"), ",")[0]
			if originName == "" || originName == "-" {
				continue
			}

			s := fmt.Sprintf("%s.%v: %T", typeName, originName, field.Value())
			s = strings.Replace(s, "service.", "", -1)
			fmt.Println(s)
		}
	}
}
