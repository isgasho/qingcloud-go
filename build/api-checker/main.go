// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate go run gen_helper.go
//go:generate go run gen_helper_swagger.go
//go:generate go fmt

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/fatih/structs"
)

var (
	pbTypeInfoMap      = map[string]interface{}{}
	swaggerTypeInfoMap = map[string]interface{}{}
)

func main() {
	makePbTypeInfoFile("./api_pb_types.txt")
}

func makePbTypeInfoFile(name string) {
	var buf bytes.Buffer
	for _, s := range scanPbTypeInfo() {
		fmt.Fprintln(&buf, s)
	}

	err := ioutil.WriteFile(name, buf.Bytes(), 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}
}

func scanPbTypeInfo() []string {
	var ss []string
	for typeName, typeValue := range pbTypeInfoMap {
		if typeName == "ServerInfo" {
			continue // skip
		}
		for _, field := range structs.Fields(typeValue) {
			filedName := strings.Split(field.Tag("json"), ",")[0]
			if filedName == "" || filedName == "-" {
				continue // skip
			}

			typeName := fixPbTypeName(typeName)
			typeInfo := fmt.Sprintf("%T", field.Value())

			ss = append(ss, typeName+"."+filedName+":"+fixPbTypeInfo(typeInfo))
		}
	}
	sort.Strings(ss)
	return ss
}

func fixPbTypeName(s string) string {
	return s
}

func fixPbTypeInfo(s string) string {
	return s
}
