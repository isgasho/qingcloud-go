// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package pbencoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtoMessageToMap(msg interface{}) (m map[string]string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("protoMessageToMap failed: %v", x)
		}
	}()

	d, err := EncodeJson(msg)
	if err != nil {
		return nil, err
	}

	var mapx map[string]interface{}
	if err := json.Unmarshal(d, &mapx); err != nil {
		return nil, err
	}

	m = UnpackMapXToMapString(mapx)
	return m, nil
}

// X is oneof string/float64/[]interface/map[string]interface{}
func UnpackMapXToMapString(mapx map[string]interface{}) map[string]string {
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
					for kk, vv := range UnpackMapXToMapString(vi) {
						m[ki+"."+kk] = vv
					}
				default:
					// unreachable
				}
			}
		case map[string]interface{}:
			for kk, vv := range UnpackMapXToMapString(v) {
				m[k+"."+kk] = vv
			}
		default:
			// unreachable
		}
	}
	return m
}

func EncodeJson(x interface{}) ([]byte, error) {
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
