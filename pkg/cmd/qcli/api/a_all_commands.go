// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli_pb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

var AllCommands []cli.Command

const pkgDefaultConfig = `{
	"api_server": "https://api.qingcloud.com/iaas/",
	"access_key_id": "",
	"secret_access_key": "",
	"zone": ""
}
`

func init() {
	if !pkgIsUserConfigExists() {
		pkgInstallUserConfig()
	}
}

func pkgGetServerInfo(c *cli.Context) (p *pb.ServerInfo) {
	if c.IsSet("config") {
		p = pkgMustLoadConfig(c.GlobalString("config"))
	} else {
		p = pkgLoadConfigIgnoreAnyError(c.GlobalString("config"))
	}

	if c.IsSet("api_server") || p.GetApiServer() == "" {
		p.ApiServer = proto.String(c.GlobalString("api_server"))
	}
	if c.IsSet("access_key_id") || p.GetAccessKeyId() == "" {
		p.AccessKeyId = proto.String(c.GlobalString("access_key_id"))
	}
	if c.IsSet("secret_access_key") || p.GetSecretAccessKey() == "" {
		p.SecretAccessKey = proto.String(c.GlobalString("secret_access_key"))
	}
	if c.IsSet("zone") || p.GetZone() == "" {
		p.Zone = proto.String(c.GlobalString("zone"))
	}

	return
}

func pkgIsUserConfigExists() bool {
	cfgpath := "~/.qingcloud/qcli.json"
	if strings.HasPrefix(cfgpath, "~/") || strings.HasPrefix(cfgpath, `~\`) {
		cfgpath = pkgGetHomePath() + cfgpath[1:]
	}

	fi, err := os.Stat(cfgpath)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}

	return true
}

func pkgMustLoadConfig(path string) *pb.ServerInfo {
	p := new(pb.ServerInfo)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func pkgLoadConfigIgnoreAnyError(path string) *pb.ServerInfo {
	p := new(pb.ServerInfo)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return p
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		return p
	}

	return p
}

func pkgInstallUserConfig() error {
	cfgpath := "~/.qingcloud/qcli.json"
	if strings.HasPrefix(cfgpath, "~/") || strings.HasPrefix(cfgpath, `~\`) {
		cfgpath = pkgGetHomePath() + cfgpath[1:]
	}

	return ioutil.WriteFile(cfgpath, []byte(pkgDefaultConfig), 0644)
}

func pkgGetHomePath() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	if home == "" {
		home = "~"
	}

	return home
}
