// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"time"
)

var (
	DefaultHost     = "api.qingcloud.com"
	DefaultUri      = "/iaas"
	DefaultProtocol = "https"
	DefaultPort     = 443

	DefaultCallTimeout       = time.Duration(0)
	DefaultCallRequestMethod = "GET"
)

type Options struct {
	Host     *string
	Uri      *string
	Protocol *string
	Port     *int
}

func (p *Options) GetHost() string {
	if p == nil || p.Host == nil {
		return DefaultHost
	}
	return *p.Host
}

func (p *Options) GetUri() string {
	if p == nil || p.Uri == nil {
		return DefaultUri
	}
	return *p.Uri
}

func (p *Options) GetProtocol() string {
	if p == nil || p.Protocol == nil {
		return DefaultProtocol
	}
	return *p.Protocol
}

func (p *Options) GetPort() int {
	if p == nil || p.Port == nil {
		return DefaultPort
	}
	return *p.Port
}

type CallOptions struct {
	RequestMethod *string
	Timeout       *time.Duration
}

func (p *CallOptions) GetRequestMethod() string {
	if p == nil || p.RequestMethod == nil {
		return DefaultCallRequestMethod
	}
	return *p.RequestMethod
}

func (p *CallOptions) GetTimeout() time.Duration {
	if p == nil || p.Timeout == nil {
		return DefaultCallTimeout
	}
	return *p.Timeout
}
