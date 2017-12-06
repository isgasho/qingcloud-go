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

package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/chai2010/qingcloud-go/pkg/errors"
	"github.com/chai2010/qingcloud-go/pkg/request/data"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Unpacker is the response unpacker.
type Unpacker struct {
	operation    *data.Operation
	httpResponse *http.Response
	output       interface{}
}

// UnpackHTTPRequest unpack the http response with an operation, http response and an output.
func (u *Unpacker) UnpackHTTPRequest(o *data.Operation, r *http.Response, x interface{}) (respBody string, err error) {
	u.operation = o
	u.httpResponse = r
	u.output = x

	respBody, err = u.parseResponse()
	if err != nil {
		return respBody, err
	}

	err = u.parseError()
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

func (u *Unpacker) parseResponse() (respBody string, err error) {
	if u.httpResponse.StatusCode == 200 {
		if u.httpResponse.Header.Get("Content-Type") == "application/json" {
			buffer := &bytes.Buffer{}
			buffer.ReadFrom(u.httpResponse.Body)
			u.httpResponse.Body.Close()

			respBody = string(buffer.Bytes())
			err := u.decodeJson(buffer.Bytes(), u.output)
			if err != nil {
				return respBody, err
			}
		}
	}

	return respBody, nil
}

func (u *Unpacker) parseError() error {
	if x, ok := u.output.(errors.QingCloudErrorInterface); ok {
		if x.GetRetCode() != 0 {
			return &errors.QingCloudError{
				RetCode: int(x.GetRetCode()),
				Message: x.GetMessage(),
			}
		}
	}

	return nil
}

func (u *Unpacker) decodeJson(data []byte, x interface{}) error {
	if x, ok := x.(proto.Message); ok {
		decoder := &jsonpb.Unmarshaler{
			AllowUnknownFields: !u.operation.Config.JSONDisableUnknownFields,
		}
		err := decoder.Unmarshal(bytes.NewReader(data), x)
		if err != nil {
			return err
		}
		return nil
	}

	return json.Unmarshal(data, x)
}
