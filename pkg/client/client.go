// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/internal/pbutil"
	"github.com/chai2010/qingcloud-go/pkg/signature"
	"github.com/chai2010/qingcloud-go/pkg/status"
)

var DebugMode = false

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
	if apiServer == "" {
		log.Fatalf("qingcloud-go/pkg/client.NewClient: invalid apiServer: %q", apiServer)
	}
	if accessKeyId == "" {
		log.Fatalf("qingcloud-go/pkg/client.NewClient: invalid accessKeyId: %q", accessKeyId)
	}
	if secretAccessKey == "" {
		log.Fatalf("qingcloud-go/pkg/client.NewClient: invalid secretAccessKey: %q", secretAccessKey)
	}

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
	if DebugMode {
		log.Printf("input: %#v\n", input)
	}

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
	inputMap["version"] = "1"

	if inputMap["time_stamp"] == "" {
		inputMap["time_stamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	}
	if inputMap["zone"] == "" {
		inputMap["zone"] = p.zone
	}

	if DebugMode {
		log.Printf("inputMap: %#v\n", inputMap)
	}

	query, sig := signature.Build(
		p.accessKeyId, p.secretAccessKey,
		httpMethod, p.getApiServerPath(),
		inputMap,
	)

	if DebugMode {
		log.Printf("query: %v\n", query)
		log.Printf("sig: %v\n", sig)
	}

	switch httpMethod {
	case "GET":
		return p.doGet(opt.GetHttpClient(), query, output)
	case "POST":
		return p.doPost(opt.GetHttpClient(), query, output)
	default:
		return fmt.Errorf("pkg/client.CallMethod: unsupport methond %v", httpMethod)
	}
}

func (p *Client) CallMethodWithMapX(
	svcMethodName, httpMethod string, input map[string]interface{},
	opt *CallOptions,
) (
	output string, err error,
) {
	if DebugMode {
		log.Printf("input: %#v\n", input)
	}

	inputMap := pbutil.UnpackMapXToMapString(input)
	inputMap["action"] = svcMethodName
	inputMap["version"] = "1"

	if inputMap["time_stamp"] == "" {
		inputMap["time_stamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	}
	if inputMap["zone"] == "" {
		inputMap["zone"] = p.zone
	}

	if DebugMode {
		log.Printf("inputMap: %#v\n", inputMap)
	}

	query, sig := signature.Build(
		p.accessKeyId, p.secretAccessKey,
		httpMethod, p.getApiServerPath(),
		inputMap,
	)

	if DebugMode {
		log.Printf("query: %v\n", query)
		log.Printf("sig: %v\n", sig)
	}

	switch httpMethod {
	case "GET":
		return p.doGet_json(opt.GetHttpClient(), query)
	case "POST":
		return p.doPost_json(opt.GetHttpClient(), query)
	default:
		return "", fmt.Errorf("pkg/client.CallMethod: unsupport methond %v", httpMethod)
	}
}

func (p *Client) CallMethodWithJson(
	svcMethodName, httpMethod string, input string,
	opt *CallOptions,
) (
	output string, err error,
) {
	if DebugMode {
		log.Printf("input: %#v\n", input)
	}

	inputMap, err := pbutil.JsonToMapString(input)
	if err != nil {
		return "", err
	}

	inputMap["action"] = svcMethodName
	inputMap["version"] = "1"

	if inputMap["time_stamp"] == "" {
		inputMap["time_stamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	}
	if inputMap["zone"] == "" {
		inputMap["zone"] = p.zone
	}

	if DebugMode {
		log.Printf("inputMap: %#v\n", inputMap)
	}

	query, sig := signature.Build(
		p.accessKeyId, p.secretAccessKey,
		httpMethod, p.getApiServerPath(),
		inputMap,
	)

	if DebugMode {
		log.Printf("query: %v\n", query)
		log.Printf("sig: %v\n", sig)
	}

	switch httpMethod {
	case "GET":
		return p.doGet_json(opt.GetHttpClient(), query)
	case "POST":
		return p.doPost_json(opt.GetHttpClient(), query)
	default:
		return "", fmt.Errorf("pkg/client.CallMethod: unsupport methond %v", httpMethod)
	}
}

func (p *Client) getApiServerPath() string {
	u, _ := url.Parse(p.apiServer)
	return u.Path
}

func (p *Client) doGet_json(c *http.Client, query string) (string, error) {
	url := p.apiServer + "?" + query
	if DebugMode {
		log.Printf("GET: %v\n", url)
	}

	resp, err := c.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	return string(buf.Bytes()), err
}

func (p *Client) doGet(c *http.Client, query string, output proto.Message) error {
	url := p.apiServer + "?" + query
	if DebugMode {
		log.Printf("GET: %v\n", url)
	}

	resp, err := c.Get(url)
	if err != nil {
		return err
	}

	err = DecodeResponse(resp, output)
	if err != nil {
		return err
	}
	return nil
}

func (p *Client) doPost(c *http.Client, query string, output proto.Message) error {
	if DebugMode {
		log.Printf("POST: %v, %v\n", p.apiServer, "application/x-www-form-urlencoded")
		log.Printf("Body: %v\n", query)
	}

	resp, err := c.Post(p.apiServer, "application/x-www-form-urlencoded", strings.NewReader(query))
	if err != nil {
		return err
	}

	err = DecodeResponse(resp, output)
	if err != nil {
		return err
	}
	return nil
}

func (p *Client) doPost_json(c *http.Client, query string) (string, error) {
	if DebugMode {
		log.Printf("POST: %v, %v\n", p.apiServer, "application/json")
		log.Printf("Body: %v\n", query)
	}

	resp, err := c.Post(p.apiServer, "application/json", strings.NewReader(query))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	return string(buf.Bytes()), err
}

func DecodeResponse(resp *http.Response, output proto.Message) error {
	defer resp.Body.Close()

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)

	if DebugMode {
		log.Printf("RespBody: %s\n", buf.Bytes())
	}

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
