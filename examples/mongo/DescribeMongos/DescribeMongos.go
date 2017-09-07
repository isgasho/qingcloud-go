// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	qcConfig "github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/spec.pb"
)

func main() {
	flag.Parse()

	conf := loadUserConfig()
	conf.SetLogLevel("warn") // debug/warn
	conf.JSONAllowUnknownFields = true

	mgoService, err := pb.NewMongoService(conf, "pek3a")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := mgoService.DescribeMongos(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encodeJSON(reply)))
}

func loadUserConfig() *qcConfig.Config {
	conf, err := qcConfig.NewDefault()
	if err != nil {
		log.Fatal(err)
	}
	err = conf.LoadUserConfig()
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

func encodeJSON(m interface{}) []byte {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return nil
	}
	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	return data
}
