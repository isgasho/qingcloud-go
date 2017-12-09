// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	ApiServerHost   string
	LogLevel        string `yaml:"log_level"`
	Zone            string `yaml:"zone"`
}

func Default() *Config {
	return &Config{}
}

func Load(path string) (*Config, error) {
	panic("TODO")
}

func MustLoad(path string) *Config {
	panic("TODO")
}

func IsUserConfigExists() bool {
	panic("TODO")
}

func LoadUserConfig() (*Config, error) {
	return &Config{}, nil
}

func MustLoadUserConfig() *Config {
	return &Config{}
}

func Parse(content string) (*Config, error) {
	panic("TODO")
}

func (p *Config) Clone() *Config {
	panic("TODO")
}

func (p *Config) Save(path string) error {
	panic("TODO")
}

func (p *Config) String() string {
	return ""
}

func GetHomePath() string {
	panic("TODO")
}
