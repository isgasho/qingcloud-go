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
	"testing"

	. "github.com/chai2010/assert"
)

func TestConfig(t *testing.T) {
	config := Config{
		AccessKeyID:       "AccessKeyID",
		SecretAccessKey:   "SecretAccessKey",
		Protocol:          "http",
		ConnectionRetries: 10,
		LogLevel:          "warn",
	}

	AssertEqual(t, "AccessKeyID", config.AccessKeyID)
	AssertEqual(t, "SecretAccessKey", config.SecretAccessKey)
	AssertEqual(t, "http", config.Protocol)
	AssertEqual(t, 10, config.ConnectionRetries)
	AssertEqual(t, "warn", config.LogLevel)
}

func TestConfig_LoadDefaultConfig(t *testing.T) {
	config := NewDefault()

	AssertEqual(t, "", config.AccessKeyID)
	AssertEqual(t, "", config.SecretAccessKey)
	AssertEqual(t, "https", config.Protocol)
}

func TestConfig_LoadUserConfig(t *testing.T) {
	config, err := LoadUserConfig()
	AssertNil(t, err)

	AssertEqual(t, "https", config.Protocol)
}

func TestConfig_LoadConfigFromContent(t *testing.T) {
	fileContent := `
qy_access_key_id: 'access_key_id'
qy_secret_access_key: 'secret_access_key'

log_level: 'debug'
`

	config, err := LoadConfigFromContent([]byte(fileContent))
	AssertNil(t, err)

	AssertEqual(t, "access_key_id", config.AccessKeyID)
	AssertEqual(t, "secret_access_key", config.SecretAccessKey)
	AssertEqual(t, "https", config.Protocol)
}

func TestNewDefault(t *testing.T) {
	config := NewDefault()

	AssertEqual(t, "", config.AccessKeyID)
	AssertEqual(t, "", config.SecretAccessKey)
	AssertEqual(t, "https", config.Protocol)
	AssertEqual(t, 3, config.ConnectionRetries)
}
