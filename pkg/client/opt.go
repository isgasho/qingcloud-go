// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"net/http"
)

type CallOptions struct {
	HttpClient *http.Client
	HttpMethod *string
}

func (p *CallOptions) GetHttpMethod() string {
	if p == nil || p.HttpMethod == nil {
		return "GET"
	}
	return *p.HttpMethod
}

func (p *CallOptions) GetHttpClient() *http.Client {
	if p == nil || p.HttpMethod == nil {
		return http.DefaultClient
	}
	return p.HttpClient
}
