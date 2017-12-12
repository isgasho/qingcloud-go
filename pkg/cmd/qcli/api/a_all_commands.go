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
	pbutil "github.com/chai2010/qingcloud-go/pkg/pbutil"
)

var (
	AllCommands []cli.Command
	DebugMode   = false
)

const pkgDefaultConfig = `{
	"api_server": "https://api.qingcloud.com/iaas/",
	"access_key_id": "",
	"secret_access_key": "",
	"zone": "pek3a"
}
`

func init() {
	if !pkgIsUserConfigExists() {
		pkgInstallUserConfig()
	}
}

func pkgGetServerInfo(c *cli.Context) (p *pb.ServerInfo) {
	p = pkgMustLoadConfig(c.GlobalString("config"))

	if c.GlobalIsSet("api_server") || p.GetApiServer() == "" {
		p.ApiServer = proto.String(c.GlobalString("api_server"))
	}
	if c.GlobalIsSet("access_key_id") || p.GetAccessKeyId() == "" {
		p.AccessKeyId = proto.String(c.GlobalString("access_key_id"))
	}
	if c.GlobalIsSet("secret_access_key") || p.GetSecretAccessKey() == "" {
		p.SecretAccessKey = proto.String(c.GlobalString("secret_access_key"))
	}
	if c.GlobalIsSet("zone") || p.GetZone() == "" {
		p.Zone = proto.String(c.GlobalString("zone"))
	}

	if DebugMode {
		s, _ := pbutil.EncodeJsonIndent(p)
		log.Printf("qcli/api.pkgGetServerInfo: p = %s", s)
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

func pkgMustLoadConfig(cfgpath string) *pb.ServerInfo {
	p := new(pb.ServerInfo)

	if strings.HasPrefix(cfgpath, "~/") || strings.HasPrefix(cfgpath, `~\`) {
		cfgpath = pkgGetHomePath() + cfgpath[1:]
	}

	data, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		log.Fatalf("qcli/api.pkgMustLoadConfig: %v", err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		log.Fatalf("qcli/api.pkgMustLoadConfig: %v", err)
	}

	if DebugMode {
		s, _ := pbutil.EncodeJsonIndent(p)
		log.Printf("qcli/api.pkgMustLoadConfig: p = %s", s)
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
