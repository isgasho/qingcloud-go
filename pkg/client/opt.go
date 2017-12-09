// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"time"
)

var (
	DefaultCallTimeout       = time.Duration(0)
	DefaultCallRequestMethod = "GET"
)

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
