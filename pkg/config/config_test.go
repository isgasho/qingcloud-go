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

package config

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

func TestConfig(t *testing.T) {
	config := Config{
		AccessKeyID:     "AccessKeyID",
		SecretAccessKey: "SecretAccessKey",
		Protocol:        "http",
		LogLevel:        "warn",
	}

	tAssert(t, "AccessKeyID" == config.AccessKeyID)
	tAssert(t, "SecretAccessKey" == config.SecretAccessKey)
	tAssert(t, "http" == config.Protocol)
	tAssert(t, "warn" == config.LogLevel)
}

func TestConfig_LoadDefaultConfig(t *testing.T) {
	config := Config{}
	config.LoadDefaultConfig()

	tAssert(t, "" == config.AccessKeyID)
	tAssert(t, "" == config.SecretAccessKey)
	tAssert(t, "https" == config.Protocol)
}

func TestConfig_LoadUserConfig(t *testing.T) {
	config := Config{}
	config.LoadUserConfig()

	tAssert(t, "https" == config.Protocol)
}

func TestConfig_LoadConfigFromContent(t *testing.T) {
	fileContent := `
qy_access_key_id: 'access_key_id'
qy_secret_access_key: 'secret_access_key'

log_level: 'debug'
`

	config := Config{}
	config.LoadConfigFromContent([]byte(fileContent))

	tAssert(t, "access_key_id" == config.AccessKeyID)
	tAssert(t, "secret_access_key" == config.SecretAccessKey)
	tAssert(t, "https" == config.Protocol)
}

func TestNewDefault(t *testing.T) {
	config, err := NewDefault()
	tAssert(t, err == nil)

	tAssert(t, "" == config.AccessKeyID)
	tAssert(t, "" == config.SecretAccessKey)
	tAssert(t, "https" == config.Protocol)
}

func TestNew(t *testing.T) {
	config, err := New("AccessKeyID", "SecretAccessKey")
	tAssert(t, err == nil)

	tAssert(t, "AccessKeyID" == config.AccessKeyID)
	tAssert(t, "SecretAccessKey" == config.SecretAccessKey)
	tAssert(t, "https" == config.Protocol)
}
