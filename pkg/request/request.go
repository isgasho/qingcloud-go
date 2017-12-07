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

// Package request contains everything for build a http request and unpack a http response.
package request

import (
	"net/http"

	"github.com/chai2010/qingcloud-go/pkg/config"
)

// Operation stores information of an operation.
type Operation struct {
	Config     *config.Config
	Properties interface{}

	APIName       string
	RequestMethod string
	RequestURI    string
	ResponseBody  string
}

// A Request can build, sign, send and unpack API request.
type Request struct {
	Operation *Operation
	Input     interface{}
	Output    interface{}

	HTTPRequest  *http.Request
	HTTPResponse *http.Response
}

// New create a Request from given Operation, Input and Output.
// It returns a Request.
func New(o *Operation, i, x interface{}) (*Request, error) {
	type Validator interface {
		Validate() error
	}
	if x, ok := i.(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	return &Request{
		Operation: o,
		Input:     i,
		Output:    x,
	}, nil
}

// Send sends API request.
// It returns error if error occurred.
func (r *Request) Send() error {
	httpRequest, err := new(Builder).BuildHTTPRequest(r.Operation, r.Input)
	if err != nil {
		return err
	}
	r.HTTPRequest = httpRequest

	s := &Signer{
		AccessKeyID:     r.Operation.Config.AccessKeyID,
		SecretAccessKey: r.Operation.Config.SecretAccessKey,
	}
	err = s.WriteSignature(r.HTTPRequest)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(r.HTTPRequest)
	if err != nil {
		return err
	}
	r.HTTPResponse = response

	respBody, err := new(Unpacker).UnpackHTTPRequest(r.Operation, r.HTTPResponse, r.Output)
	if err != nil {
		return err
	}

	r.Operation.ResponseBody = respBody
	if err != nil {
		return err
	}

	return nil
}
