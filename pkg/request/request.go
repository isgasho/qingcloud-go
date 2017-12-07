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
	"errors"
	"net/http"
)

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
	err := r.check()
	if err != nil {
		return err
	}

	err = r.build()
	if err != nil {
		return err
	}

	err = r.sign()
	if err != nil {
		return err
	}

	err = r.send()
	if err != nil {
		return err
	}

	respBody, err := r.unpack()
	r.Operation.ResponseBody = respBody
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) check() error {
	if r.Operation.Config.AccessKeyID == "" {
		return errors.New("access key not provided")
	}

	if r.Operation.Config.SecretAccessKey == "" {
		return errors.New("secret access key not provided")
	}

	return nil
}

func (r *Request) build() error {
	b := &Builder{}
	httpRequest, err := b.BuildHTTPRequest(r.Operation, r.Input)
	if err != nil {
		return err
	}

	r.HTTPRequest = httpRequest
	return nil
}

func (r *Request) sign() error {
	s := &Signer{
		AccessKeyID:     r.Operation.Config.AccessKeyID,
		SecretAccessKey: r.Operation.Config.SecretAccessKey,
	}
	err := s.WriteSignature(r.HTTPRequest)
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) send() error {
	response, err := http.DefaultClient.Do(r.HTTPRequest)
	if err != nil {
		return err
	}

	r.HTTPResponse = response
	return nil
}

func (r *Request) unpack() (respBody string, err error) {
	u := &Unpacker{}

	respBody, err = u.UnpackHTTPRequest(r.Operation, r.HTTPResponse, r.Output)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
