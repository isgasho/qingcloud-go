// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	urlpkg "net/url"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/errors"
	"github.com/chai2010/qingcloud-go/pkg/pbutil"
	"github.com/chai2010/qingcloud-go/pkg/signature"
)

type Client struct {
	accessKeyId     string
	secretAccessKey string
	opt             *Options
}

func NewClient(accessKeyId, secretAccessKey string, opt *Options) *Client {
	return &Client{
		accessKeyId:     accessKeyId,
		secretAccessKey: secretAccessKey,
		opt:             opt,
	}
}

func (p *Client) CallMethod(
	c context.Context, svcMethod string,
	input, output proto.Message,
	opt *CallOptions,
) error {
	inputMap, err := pbutil.ProtoMessageToMap(input)
	if err != nil {
		return err
	}

	inputMap["action"] = svcMethod

	query, _ := signature.Build(
		p.accessKeyId, p.secretAccessKey,
		opt.GetRequestMethod(), p.opt.GetUri(),
		inputMap,
	)

	switch method := opt.GetRequestMethod(); method {
	case "GET":
		var url = urlpkg.URL{
			Scheme:   p.opt.GetProtocol(),
			Host:     p.opt.GetHost() + ":" + strconv.Itoa(p.opt.GetPort()),
			Path:     p.opt.GetUri(),
			RawQuery: query,
		}
		resp, err := http.Get(url.String())
		if err != nil {
			return err
		}
		err = DecodeResponse(resp, output)
		if err != nil {
			return err
		}
		return nil

	case "POST":
		var url = urlpkg.URL{
			Scheme: p.opt.GetProtocol(),
			Host:   p.opt.GetHost() + ":" + strconv.Itoa(p.opt.GetPort()),
			Path:   p.opt.GetUri(),
		}
		resp, err := http.Post(url.String(), "application/json", strings.NewReader(query))
		if err != nil {
			return err
		}
		err = DecodeResponse(resp, output)
		if err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("pkg/client.CallMethod: unsupport methond %v", method)
	}
}

func DecodeResponse(resp *http.Response, output proto.Message) error {
	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "application/json" {
		return errors.New(int32(resp.StatusCode), string(buf.Bytes()))
	}

	err := pbutil.DecodeJson(buf.Bytes(), output)
	if err != nil {
		return err
	}

	if x, ok := output.(errors.Error); ok {
		if x.GetRetCode() != 0 {
			return errors.New(x.GetRetCode(), x.GetMessage())
		}
	}

	return nil
}
