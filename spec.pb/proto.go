// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// https://github.com/google/protobuf/releases
// go get github.com/golang/protobuf/protoc-gen-go

//go:generate protoc --go_out=. base.proto instances.proto

package spec
