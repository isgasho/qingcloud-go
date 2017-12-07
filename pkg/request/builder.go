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
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Builder is the request builder for QingCloud service.
type Builder struct {
	parsedURL        string
	parsedProperties *map[string]string
	parsedParams     *map[string]string

	operation *Operation
	input     interface{}
}

// BuildHTTPRequest builds http request with an operation and an input.
func (b *Builder) BuildHTTPRequest(o *Operation, i interface{}) (*http.Request, error) {
	b.operation = o
	b.input = i

	err := b.parse()
	if err != nil {
		return nil, err
	}

	return b.build()
}

func (b *Builder) build() (*http.Request, error) {
	httpRequest, err := http.NewRequest(b.operation.RequestMethod, b.parsedURL, nil)
	if err != nil {
		return nil, err
	}

	return httpRequest, nil
}

func (b *Builder) parse() error {

	err := b.parseRequestProperties()
	if err != nil {
		return err
	}
	err = b.parseRequestParams()
	if err != nil {
		return err
	}
	err = b.parseRequestURL()
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) parseRequestProperties() error {
	propertiesMap, err := protoMessageToMap(b.operation.Properties)
	if err != nil {
		return err
	}

	b.parsedProperties = &propertiesMap
	return nil
}

func (b *Builder) parseRequestParams() error {
	requestParams, err := protoMessageToMap(b.input)
	if err != nil {
		return err
	}

	requestParams["action"] = b.operation.APIName

	b.parsedParams = &requestParams
	return nil
}

func (b *Builder) parseRequestURL() error {
	conf := b.operation.Config

	endpoint := conf.Protocol + "://" + conf.Host + ":" + strconv.Itoa(conf.Port)
	requestURI := regexp.MustCompile(`/+`).ReplaceAllString(conf.URI, "/")

	b.parsedURL = endpoint + requestURI

	if b.parsedParams != nil {
		zone := (*b.parsedProperties)["zone"]
		if zone != "" {
			(*b.parsedParams)["zone"] = zone
		}
		paramsParts := []string{}
		for key, value := range *b.parsedParams {
			paramsParts = append(paramsParts, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
		}

		joined := strings.Join(paramsParts, "&")
		if joined != "" {
			b.parsedURL += "?" + joined
		}
	}

	return nil
}

func protoMessageToMap(msg interface{}) (m map[string]string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("protoMessageToMap failed: %v", x)
		}
	}()

	d, err := encodeJson(msg)
	if err != nil {
		return nil, err
	}

	var mapx map[string]interface{}
	if err := json.Unmarshal(d, &mapx); err != nil {
		return nil, err
	}

	m = unpackMapXToMapString(mapx)
	return m, nil
}

// X is oneof string/float64/[]interface/map[string]interface{}
func unpackMapXToMapString(mapx map[string]interface{}) map[string]string {
	var m = map[string]string{}
	for k, v := range mapx {
		switch v := v.(type) {
		case string:
			m[k] = v
		case float64:
			m[k] = fmt.Sprintf("%v", v)
		case []interface{}:
			for i := 0; i < len(v); i++ {
				ki := k + "." + strconv.Itoa(i+1)
				switch vi := v[i].(type) {
				case string:
					m[ki] = vi
				case float64:
					m[ki] = fmt.Sprintf("%v", vi)
				case map[string]interface{}:
					for kk, vv := range unpackMapXToMapString(vi) {
						m[ki+"."+kk] = vv
					}
				default:
					// unreachable
				}
			}
		case map[string]interface{}:
			for kk, vv := range unpackMapXToMapString(v) {
				m[k+"."+kk] = vv
			}
		default:
			// unreachable
		}
	}
	return m
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
