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

package utils

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// JSONEncode encode given interface to json byte slice.
func JSONEncode(source interface{}, unescape bool) ([]byte, error) {
	bytesResult, err := encodeJson(source)
	if err != nil {
		return []byte{}, err
	}

	if unescape {
		bytesResult = bytes.Replace(bytesResult, []byte("\\u003c"), []byte("<"), -1)
		bytesResult = bytes.Replace(bytesResult, []byte("\\u003e"), []byte(">"), -1)
		bytesResult = bytes.Replace(bytesResult, []byte("\\u0026"), []byte("&"), -1)
	}

	return bytesResult, nil
}

func encodeJson(x interface{}) ([]byte, error) {
	if x, ok := x.(proto.Message); ok {
		jsonMarshaler := &jsonpb.Marshaler{
			OrigName: true,
			Indent:   "",
		}

		var buf bytes.Buffer
		err := jsonMarshaler.Marshal(&buf, x)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return json.Marshal(x)
}

// JSONDecode decode given json byte slice to corresponding struct.
func JSONDecode(content []byte, destinations ...interface{}) (interface{}, error) {
	var destination interface{}
	var err error
	if len(destinations) == 1 {
		destination = destinations[0]
		err = decodeJson(content, destination)
	} else {
		err = decodeJson(content, &destination)
	}

	if err != nil {
		return nil, err
	}
	return destination, err
}

func decodeJson(data []byte, x interface{}) error {
	if x, ok := x.(proto.Message); ok {
		decoder := &jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		err := decoder.Unmarshal(bytes.NewReader(data), x)
		if err != nil {
			return err
		}
		return nil
	}

	return json.Unmarshal(data, x)
}

// JSONFormatToReadable formats given json byte slice prettily.
func JSONFormatToReadable(source []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, source, "", "  ") // Using 2 space indent
	if err != nil {
		return []byte{}, err
	}

	return out.Bytes(), nil
}
