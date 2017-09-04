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
	"net/url"
	"strings"
)

// URLQueryEscape escapes the original string.
func URLQueryEscape(origin string) string {
	escaped := url.QueryEscape(origin)
	escaped = strings.Replace(escaped, "%2F", "/", -1)
	escaped = strings.Replace(escaped, "%3D", "=", -1)
	escaped = strings.Replace(escaped, "+", "%20", -1)
	return escaped
}

// URLQueryUnescape unescapes the escaped string.
func URLQueryUnescape(escaped string) (string, error) {
	escaped = strings.Replace(escaped, "/", "%2F", -1)
	escaped = strings.Replace(escaped, "=", "%3D", -1)
	escaped = strings.Replace(escaped, "%20", " ", -1)
	return url.QueryUnescape(escaped)
}
