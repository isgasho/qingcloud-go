// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"net/http"
)

type CallOptions struct {
	HttpClient *http.Client
}

func (p *CallOptions) GetHttpClient() *http.Client {
	if p == nil || p.HttpClient == nil {
		return http.DefaultClient
	}
	return p.HttpClient
}
