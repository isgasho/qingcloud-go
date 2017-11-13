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
	"fmt"
	"testing"
)

func tAssert(tb testing.TB, condition bool, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprint(a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}

func tAssertf(tb testing.TB, condition bool, format string, a ...interface{}) {
	tb.Helper()
	if !condition {
		if msg := fmt.Sprintf(format, a...); msg != "" {
			tb.Fatal("Assert failed: " + msg)
		} else {
			tb.Fatal("Assert failed")
		}
	}
}

func TestJSONDecode_Unknown(t *testing.T) {
	jsonString := `{
		"key1" : "This is a string.",
		"key2" : 10.50,
   		"key3": [null, {"nestedKey1": "Another string"}]
	}`

	anyData, err := JSONDecode([]byte(jsonString))
	tAssert(t, err == nil)
	data := anyData.(map[string]interface{})
	tAssert(t, 10.50 == data["key2"])

	var anotherData interface{}
	_, err = JSONDecode([]byte(jsonString), &anotherData)
	tAssert(t, err == nil)
	data = anyData.(map[string]interface{})
	tAssert(t, 10.50 == data["key2"])
}

func TestJSONDecode_Known(t *testing.T) {
	type SampleJSON struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	sampleJSONString := `{"name": "NAME"}`

	sample := SampleJSON{Name: "NaMe", Description: "DeScRiPtIoN"}
	anyDataPointer, err := JSONDecode([]byte(sampleJSONString), &sample)
	tAssert(t, err == nil)
	data := anyDataPointer.(*SampleJSON)
	tAssert(t, "NAME" == sample.Name)
	tAssert(t, "DeScRiPtIoN" == sample.Description)
	tAssert(t, "NAME" == (*data).Name)
	tAssert(t, "DeScRiPtIoN" == (*data).Description)
}

func TestJSONEncode(t *testing.T) {
	type SampleJSON struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	sample := SampleJSON{Name: "NaMe", Description: "DeScRiPtIoN"}

	jsonBytes, err := JSONEncode(sample, true)
	tAssert(t, err == nil)
	tAssert(t, `{"name":"NaMe","description":"DeScRiPtIoN"}` == string(jsonBytes))
}

func TestJSONFormatToReadable(t *testing.T) {
	sampleJSONString := `{"name": "NAME"}`

	jsonBytes, err := JSONFormatToReadable([]byte(sampleJSONString))
	tAssert(t, err == nil)
	tAssert(t, "{\n  \"name\": \"NAME\"\n}" == string(jsonBytes))
}
