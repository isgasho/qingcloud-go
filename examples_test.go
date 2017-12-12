// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qingcloud

import (
	"fmt"

	sigpkg "github.com/chai2010/qingcloud-go/pkg/signature"
)

func Example_signature_Build() {
	query, signature := sigpkg.Build(
		"QYACCESSKEYIDEXAMPLE", "SECRETACCESSKEY",
		"GET", "/iaas/", map[string]string{},
	)

	fmt.Println(query)
	fmt.Println(signature)
	// Output:
	// access_key_id=QYACCESSKEYIDEXAMPLE&signature_method=HmacSHA256&signature_version=1&signature=O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D
	// O5EhQeUqTF00g59t5Pb46QPfnPMhUOAcxTWvzlnraeE%3D
}
