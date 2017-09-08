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

	"github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/service"
)

func main() {
	flag.Parse()

	conf := config.MustLoadUserConfig()
	conf.JSONDisableUnknownFields = false
	conf.LogLevel = "debug" // debug/warn

	qcService, err := pb.Init(conf)
	if err != nil {
		log.Fatal(err)
	}

	nicService, err := qcService.Nic("pek3a")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := nicService.DescribeNics(&pb.DescribeNicsInput{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encodeJSON(reply)))
}

func encodeJSON(m interface{}) []byte {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return nil
	}
	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	return data
}
