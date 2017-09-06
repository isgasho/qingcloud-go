// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	qcConfig "github.com/yunify/qingcloud-sdk-go/config"
	qcSservice "github.com/yunify/qingcloud-sdk-go/service"
)

func main() {
	conf := loadUserConfig()
	conf.LogLevel = "debug"

	service, err := qcSservice.Init(conf)
	if err != nil {
		log.Fatal(err)
	}

	mgoService, err := service.Mongo("pek3a")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := mgoService.DescribeMongos(&qcSservice.DescribeMongosInput{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(encodeJSON(reply)))
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