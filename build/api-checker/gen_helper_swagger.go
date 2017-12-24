// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/template"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func main() {
	os.Remove("./z_types_swagger.go")

	var nameList []string
	for typeName, _ := range pb.TypeInfoMap {
		if isInIgnredService(typeName) {
			continue
		}
		if isIgnoredType(typeName) {
			continue
		}
		nameList = append(nameList, typeName)
	}
	sort.Strings(nameList)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(&buf,
		struct {
			NameList []string
		}{
			NameList: nameList,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./z_types_swagger.go", buf.Bytes(), 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}
}

func getServiceName(typeName string) string {
	for _, serviceSpec := range pb.ServiceApiSpecMap {
		if serviceSpec.InputTypeName == typeName {
			return serviceSpec.ServiceName
		}
		if serviceSpec.OutputTypeName == typeName {
			return serviceSpec.ServiceName
		}
	}
	return ""
}

func isInIgnredService(typeName string) bool {
	serviceName := getServiceName(typeName)
	for _, s := range ignoredServiceNameList {
		if serviceName == s {
			return true
		}
	}
	return false
}

func isIgnoredType(name string) bool {
	for _, s := range ignoredTypeNameList {
		if name == s {
			return true
		}
	}
	return false
}

var ignoredServiceNameList = []string{
	"AlarmService",
	"ResourceACLService",
	"HadoopService",
	"SpanService",
	"SparkService",
	"S2Service",
}

var ignoredTypeNameList = []string{}

const tmpl = `
{{- $NameList := .NameList -}}

// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// +build swagger_api
// go test -tags=swagger_api

package main

import (
	service "github.com/yunify/qingcloud-sdk-go/service"
)

func init() {
	{{range $_, $name := $NameList -}}
		pbTypeInfoMap["{{$name}}"] = &service.{{$name}}{}
	{{end -}}
}
`
