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
	"testing"
	"time"

	. "github.com/chai2010/qingcloud-go/internal/assert"
)

func TestTimeToString(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	AssertNil(t, err)

	someTime := time.Date(2016, 9, 1, 15, 30, 0, 0, tz)
	AssertEqual(t, "Thu, 01 Sep 2016 07:30:00 GMT", TimeToString(someTime, "RFC 822"))
	AssertEqual(t, "2016-09-01T07:30:00Z", TimeToString(someTime, "ISO 8601"))
}

func TestStringToTime(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	AssertNil(t, err)
	someTime := time.Date(2016, 9, 1, 15, 30, 0, 0, tz)

	parsedTime, err := StringToTime("Thu, 01 Sep 2016 07:30:00 GMT", "RFC 822")
	AssertNil(t, err)
	AssertEqual(t, someTime.UTC(), parsedTime)

	parsedTime, err = StringToTime("2016-09-01T07:30:00Z", "ISO 8601")
	AssertNil(t, err)
	AssertEqual(t, someTime.UTC(), parsedTime)

	parsedTime, err = StringToTime("2016-09-01T07:30:00.000Z", "ISO 8601")
	AssertNil(t, err)
	AssertEqual(t, someTime.UTC(), parsedTime)
}

func TestStringToUnixString(t *testing.T) {
	AssertEqual(t, 1472715000, StringToUnixInt("Thu, 01 Sep 2016 07:30:00 GMT", "RFC 822"))
	AssertEqual(t, 1472715000, StringToUnixInt("2016-09-01T07:30:00Z", "ISO 8601"))
	AssertEqual(t, 1472715000, StringToUnixInt("2016-09-01T07:30:00.000Z", "ISO 8601"))
}

func TestJSONDecode_Unknown(t *testing.T) {
	jsonString := `{
		"key1" : "This is a string.",
		"key2" : 10.50,
   		"key3": [null, {"nestedKey1": "Another string"}]
	}`

	anyData, err := JSONDecode([]byte(jsonString))
	AssertNil(t, err)
	data := anyData.(map[string]interface{})
	AssertEqual(t, 10.50, data["key2"])

	var anotherData interface{}
	_, err = JSONDecode([]byte(jsonString), &anotherData)
	AssertNil(t, err)
	data = anyData.(map[string]interface{})
	AssertEqual(t, 10.50, data["key2"])
}

func TestJSONDecode_Known(t *testing.T) {
	type SampleJSON struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	sampleJSONString := `{"name": "NAME"}`

	sample := SampleJSON{Name: "NaMe", Description: "DeScRiPtIoN"}
	anyDataPointer, err := JSONDecode([]byte(sampleJSONString), &sample)
	AssertNil(t, err)
	data := anyDataPointer.(*SampleJSON)
	AssertEqual(t, "NAME", sample.Name)
	AssertEqual(t, "DeScRiPtIoN", sample.Description)
	AssertEqual(t, "NAME", (*data).Name)
	AssertEqual(t, "DeScRiPtIoN", (*data).Description)
}

func TestJSONEncode(t *testing.T) {
	type SampleJSON struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	sample := SampleJSON{Name: "NaMe", Description: "DeScRiPtIoN"}

	jsonBytes, err := JSONEncode(sample, true)
	AssertNil(t, err)
	AssertEqual(t, `{"name":"NaMe","description":"DeScRiPtIoN"}`, string(jsonBytes))
}

func TestJSONFormatToReadable(t *testing.T) {
	sampleJSONString := `{"name": "NAME"}`

	jsonBytes, err := JSONFormatToReadable([]byte(sampleJSONString))
	AssertNil(t, err)
	AssertEqual(t, "{\n  \"name\": \"NAME\"\n}", string(jsonBytes))
}
