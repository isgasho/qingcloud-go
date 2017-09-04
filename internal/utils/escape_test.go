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
	"reflect"
	"testing"
)

func TestURLQueryEscape(t *testing.T) {
	tAssertEQ(t, "/", URLQueryEscape("/"))
	tAssertEQ(t, "%20", URLQueryEscape(" "))
	tAssertEQ(t,
		"%21%40%23%24%25%5E%26%2A%28%29_%2B%7B%7D%7C",
		URLQueryEscape("!@#$%^&*()_+{}|"),
	)
}

func TestURLQueryUnescape(t *testing.T) {
	x, err := URLQueryUnescape("/")
	tAssertNil(t, err)
	tAssertEQ(t, "/", x)

	x, err = URLQueryUnescape("%20")
	tAssertNil(t, err)
	tAssertEQ(t, " ", x)

	x, err = URLQueryUnescape("%21%40%23%24%25%5E%26%2A%28%29_%2B%7B%7D%7C")
	tAssertNil(t, err)
	tAssertEQ(t, "!@#$%^&*()_+{}|", x)
}

func tAssertNil(tb testing.TB, x interface{}) {
	if x != nil {
		tb.Fatalf("not nil: %v", x)
	}
}

func tAssertEQ(tb testing.TB, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		tb.Fatalf("not equal: %v, %v", a, b)
	}
}
