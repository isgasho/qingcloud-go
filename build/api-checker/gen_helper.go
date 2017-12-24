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
	"text/template"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func main() {
	os.Remove("./z_types.go")

	var nameList []string
	for typeName, _ := range pb.TypeInfoMap {
		nameList = append(nameList, typeName)
	}

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

	err = ioutil.WriteFile("./z_types.go", buf.Bytes(), 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}
}

const tmpl = `
{{- $NameList := .NameList -}}

// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

func init() {
	{{range $_, $name := $NameList -}}
		typeInfoMap["{{$name}}"] = &pb.{{$name}}{}
	{{end -}}
}
`
