// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	ApiServer       string
	Zone            string
}

func LoadConfig(path string) (*Config, error) {
	panic("TODO")
}

func SaveConfig(path string, config *Config) error {
	panic("TODO")
}
