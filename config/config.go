// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package config

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
)

// A Config stores a configuration of this sdk.
type Config struct {
	AccessKeyID     string `yaml:"qy_access_key_id"`
	SecretAccessKey string `yaml:"qy_secret_access_key"`

	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Protocol          string `yaml:"protocol"`
	URI               string `yaml:"uri"`
	ConnectionRetries int    `yaml:"connection_retries"`
	ConnectionTimeout int    `yaml:"connection_timeout"`

	JSONAllowUnknownFields bool   `yaml:"json_allow_unknown_fields"`
	LogLevel               string `yaml:"log_level"`
	Zone                   string `yaml:"zone"`

	Connection *http.Client
}

// NewDefault loads the default configuration for Config.
func NewDefault() *Config {
	c := &Config{}

	_, err := yamlDecode([]byte(DefaultConfigFileContent), c)
	if err != nil {
		glog.Fatal("Config parse error: ", err)
	}

	timeout := time.Duration(c.ConnectionTimeout) * time.Second
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
	}
	c.Connection = &http.Client{
		Transport: transport,
	}

	return c
}

// LoadUserConfig loads user configuration in ~/.qingcloud/config.yaml for Config.
// It returns error if file not found.
func LoadUserConfig() (*Config, error) {
	_, err := os.Stat(GetUserConfigFilePath())
	if err != nil {
		glog.Warning("Installing default config file to \"" + GetUserConfigFilePath() + "\"")
		InstallDefaultUserConfig()
	}

	return LoadConfigFromFilepath(GetUserConfigFilePath())
}

func MustLoadUserConfig() *Config {
	c, err := LoadUserConfig()
	if err != nil {
		panic(err)
	}
	return c
}

// LoadConfigFromFilepath loads configuration from a specified local path.
// It returns error if file not found or yaml decode failed.
func LoadConfigFromFilepath(filepath string) (*Config, error) {
	if strings.Index(filepath, "~/") == 0 {
		filepath = strings.Replace(filepath, "~/", getHome()+"/", 1)
	}

	configYAML, err := ioutil.ReadFile(filepath)
	if err != nil {
		glog.Error("File not found: " + filepath)
		return nil, err
	}

	return LoadConfigFromContent(configYAML)
}

// LoadConfigFromContent loads configuration from a given byte slice.
// It returns error if yaml decode failed.
func LoadConfigFromContent(content []byte) (*Config, error) {
	c := NewDefault()

	_, err := yamlDecode(content, c)
	if err != nil {
		glog.Error("Config parse error: " + err.Error())
		return nil, err
	}

	timeout := time.Duration(c.ConnectionTimeout) * time.Second
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
	}
	c.Connection = &http.Client{
		Transport: transport,
	}

	return c, nil
}
