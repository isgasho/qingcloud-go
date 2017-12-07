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
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/chai2010/qingcloud-go/pkg/pbutil"
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

	propertiesMap, err := pbutil.ProtoMessageToMap(b.operation.Properties)
	if err != nil {
		return nil, err
	}
	b.parsedProperties = &propertiesMap

	requestParams, err := pbutil.ProtoMessageToMap(b.input)
	if err != nil {
		return nil, err
	}
	requestParams["action"] = b.operation.APIName
	b.parsedParams = &requestParams

	err = b.parseRequestURL()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(b.operation.RequestMethod, b.parsedURL, nil)
	if err != nil {
		return nil, err
	}

	return httpRequest, nil
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
