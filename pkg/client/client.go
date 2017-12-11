// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/pbutil"
	"github.com/chai2010/qingcloud-go/pkg/signature"
	"github.com/chai2010/qingcloud-go/pkg/status"
)

type Validator interface {
	Validate() error
}

type Client struct {
	accessKeyId     string
	secretAccessKey string
	apiServer       string
	zone            string
}

func NewClient(apiServer, accessKeyId, secretAccessKey, zone string) *Client {
	return &Client{
		apiServer:       apiServer,
		accessKeyId:     accessKeyId,
		secretAccessKey: secretAccessKey,
		zone:            zone,
	}
}

func (p *Client) CallMethod(
	svcMethodName, httpMethod string,
	input, output proto.Message,
	opt *CallOptions,
) error {
	if x, ok := input.(Validator); ok {
		if err := x.Validate(); err != nil {
			return err
		}
	}

	inputMap, err := pbutil.ProtoMessageToMap(input)
	if err != nil {
		return err
	}

	inputMap["action"] = svcMethodName

	query, _ := signature.Build(
		p.accessKeyId, p.secretAccessKey,
		httpMethod, p.getApiServerPath(),
		inputMap,
	)

	switch httpMethod {
	case "GET":
		resp, err := opt.GetHttpClient().Get(p.apiServer + "?" + query)
		if err != nil {
			return err
		}
		err = DecodeResponse(resp, output)
		if err != nil {
			return err
		}
		return nil

	case "POST":
		resp, err := opt.GetHttpClient().Post(p.apiServer, "application/json", strings.NewReader(query))
		if err != nil {
			return err
		}
		err = DecodeResponse(resp, output)
		if err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("pkg/client.CallMethod: unsupport methond %v", httpMethod)
	}
}

func (p *Client) getApiServerPath() string {
	u, _ := url.Parse(p.apiServer)
	return u.Path
}

func DecodeResponse(resp *http.Response, output proto.Message) error {
	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "application/json" {
		return status.NewError(int32(resp.StatusCode), string(buf.Bytes()))
	}

	err := pbutil.DecodeJson(buf.Bytes(), output)
	if err != nil {
		return err
	}

	if s, ok := output.(status.Status); ok {
		if err := status.Error(s); err != nil {
			return err
		}
	}

	return nil
}
