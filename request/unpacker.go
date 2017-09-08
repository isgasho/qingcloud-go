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
	"fmt"
	"net/http"

	"github.com/chai2010/qingcloud-go/request/errors"
	"github.com/golang/glog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Unpacker is the response unpacker.
type Unpacker struct {
	operation    *Operation
	httpResponse *http.Response
	output       proto.Message
}

// UnpackHTTPRequest unpack the http response with an operation, http response and an output.
func (u *Unpacker) UnpackHTTPRequest(o *Operation, r *http.Response, x proto.Message) error {
	u.operation = o
	u.httpResponse = r
	u.output = x

	err := u.parseResponse()
	if err != nil {
		return err
	}

	err = u.parseError()
	if err != nil {
		return err
	}

	return nil
}

func (u *Unpacker) parseResponse() error {
	if u.httpResponse.StatusCode == 200 {
		if u.httpResponse.Header.Get("Content-Type") == "application/json" {
			buffer := &bytes.Buffer{}
			buffer.ReadFrom(u.httpResponse.Body)
			u.httpResponse.Body.Close()

			glog.Info(fmt.Sprintf(
				"Response json string: [%d] %s",
				StringToUnixInt(u.httpResponse.Header.Get("Date"), "RFC 822"),
				string(buffer.Bytes())))

			decoder := &jsonpb.Unmarshaler{
				AllowUnknownFields: u.operation.Config.JSONAllowUnknownFields,
			}

			err := decoder.Unmarshal(buffer, u.output)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
