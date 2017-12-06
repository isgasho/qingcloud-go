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

// Package config contains a Config struct for QingCloud SDK.
package config

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/chai2010/qingcloud-go/pkg/logger"
)

// A Config stores a configuration of this sdk.
type Config struct {
	AccessKeyID     string `yaml:"qy_access_key_id"`
	SecretAccessKey string `yaml:"qy_secret_access_key"`

	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Protocol          string `yaml:"protocol"`
	URI               string `yaml:"uri"`
	ConnectionTimeout int    `yaml:"connection_timeout"`

	JSONDisableUnknownFields bool   `yaml:"json_disable_unknown_fields"`
	LogLevel                 string `yaml:"log_level"`
	Zone                     string `yaml:"zone"`
}

// NewDefault loads the default configuration for Config.
func NewDefault() (*Config, error) {
	c := &Config{}

	err := yaml.Unmarshal([]byte(DefaultConfigFileContent), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// LoadUserConfig loads user configuration in ~/.qingcloud/config.yaml for Config.
// It returns error if file not found.
func LoadUserConfig() (*Config, error) {
	return LoadConfigFromFilepath(strings.Replace(DefaultConfigFile, "~/", getHome()+"/", 1))
}

// MustLoadUserConfig loads user configuration in ~/.qingcloud/config.yaml for Config.
// It panic if failed.
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
		logger.Error("File not found: " + filepath)
		return nil, err
	}

	return LoadConfigFromContent(configYAML)
}
func MustLoadConfigFromFilepath(filepath string) *Config {
	cfg, err := LoadConfigFromFilepath(filepath)
	if err != nil {
		logger.Fatal(err)
	}
	return cfg
}

// LoadConfigFromContent loads configuration from a given byte slice.
// It returns error if yaml decode failed.
func LoadConfigFromContent(content []byte) (*Config, error) {
	c, err := NewDefault()
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, c)
	if err != nil {
		logger.Error("Config parse error: " + err.Error())
		return nil, err
	}

	return c, nil
}

func getHome() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}

	return home
}

// DefaultConfigFile is the default config file.
const DefaultConfigFile = "~/.qingcloud/config.yaml"

// DefaultConfigFileContent is the default config file content.
const DefaultConfigFileContent = `# QingCloud services configuration

#qy_access_key_id: 'ACCESS_KEY_ID'
#qy_secret_access_key: 'SECRET_ACCESS_KEY'

host: 'api.qingcloud.com'
port: 443
protocol: 'https'
uri: '/iaas'
connection_retries: 3
connection_timeout: 30

json_allow_unknown_fields: false

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: 'warn'

`
