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
	"fmt"
	"testing"
	"time"
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

func TestTimeToString(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	tAssert(t, err == nil)

	someTime := time.Date(2016, 9, 1, 15, 30, 0, 0, tz)
	tAssert(t, "Thu, 01 Sep 2016 07:30:00 GMT" == TimeToString(someTime, "RFC 822"))
	tAssert(t, "2016-09-01T07:30:00Z" == TimeToString(someTime, "ISO 8601"))
}

func TestStringToTime(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	tAssert(t, err == nil)
	someTime := time.Date(2016, 9, 1, 15, 30, 0, 0, tz)

	parsedTime, err := StringToTime("Thu, 01 Sep 2016 07:30:00 GMT", "RFC 822")
	tAssert(t, err == nil)
	tAssert(t, someTime.UTC() == parsedTime)

	parsedTime, err = StringToTime("2016-09-01T07:30:00Z", "ISO 8601")
	tAssert(t, err == nil)
	tAssert(t, someTime.UTC() == parsedTime)

	parsedTime, err = StringToTime("2016-09-01T07:30:00.000Z", "ISO 8601")
	tAssert(t, err == nil)
	tAssert(t, someTime.UTC() == parsedTime)
}

func TestStringToUnixString(t *testing.T) {
	tAssert(t, 1472715000 == StringToUnixInt("Thu, 01 Sep 2016 07:30:00 GMT", "RFC 822"))
	tAssert(t, 1472715000 == StringToUnixInt("2016-09-01T07:30:00Z", "ISO 8601"))
	tAssert(t, 1472715000 == StringToUnixInt("2016-09-01T07:30:00.000Z", "ISO 8601"))
}
