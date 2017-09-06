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
	"time"
)

var layoutMap = map[string]string{
	"RFC 822":  "Mon, 02 Jan 2006 15:04:05 GMT",
	"ISO 8601": "2006-01-02T15:04:05Z",
}

// TimeToString transforms given time to string.
func TimeToString(timeValue time.Time, format string) string {
	return timeValue.UTC().Format(layoutMap[format])
}

// StringToTime transforms given string to time.
func StringToTime(timeString, format string) (time.Time, error) {
	return time.Parse(layoutMap[format], timeString)
}

// StringToUnixInt transforms given string to unix time int.
func StringToUnixInt(timeString, format string) int {
	t, err := StringToTime(timeString, format)
	if err != nil {
		return 0
	}

	return int(t.Unix())
}

// JSONEncode encode given interface to json byte slice.
func JSONEncode(source interface{}, unescape bool) ([]byte, error) {
	bytesResult, err := json.Marshal(source)
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

// JSONDecode decode given json byte slice to corresponding struct.
func JSONDecode(content []byte, destinations ...interface{}) (interface{}, error) {
	var destination interface{}
	var err error
	if len(destinations) == 1 {
		destination = destinations[0]
		err = json.Unmarshal(content, destination)
	} else {
		err = json.Unmarshal(content, &destination)
	}

	if err != nil {
		return nil, err
	}
	return destination, err
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
