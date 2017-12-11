// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package pbutil

import (
	_ "github.com/fatih/structs"
	"github.com/golang/protobuf/proto"
)

func EncodeToMap(msg proto.Message) (m map[string]string, err error) {
	return ProtoMessageToMap(msg)
}

func DecodeFromMap(m map[string]string, x proto.Message) error {
	panic("TODO")
}
