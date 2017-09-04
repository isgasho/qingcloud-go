// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package utils

import (
	"reflect"
	"testing"
)

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
