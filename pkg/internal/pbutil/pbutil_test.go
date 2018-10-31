// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package pbutil_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/chai2010/qingcloud-go/pkg/internal/pbutil"
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

func TestEncodeJson(t *testing.T) {
	// TODO
}

func TestDecodeJson(t *testing.T) {
	// TODO
}

func TestProtoMessageToMap(t *testing.T) {
	msg := &TestUnknownService{
		Action: proto.String("CreateMongo"),
		Vxnet:  proto.String("vxnet-dls87x2"),
		Zone:   proto.String("pk1a"),
		AList:  []string{"a", "b", "c"},
	}

	m, err := pbutil.ProtoMessageToMap(msg)
	tAssert(t, err == nil, err)
	tAssert(t, len(m) == 6, m)

	tAssert(t, m["action"] == msg.GetAction())
	tAssert(t, m["vxnet"] == msg.GetVxnet())
	tAssert(t, m["zone"] == msg.GetZone())

	for i := 0; i < len(msg.AList); i++ {
		tAssert(t, m["a_list."+strconv.Itoa(i+1)] == msg.AList[i])
	}
}
