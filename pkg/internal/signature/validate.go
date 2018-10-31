// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package signature

import (
	"net/url"
	"strings"
)

func GetSignatureInfo(query string) (pubKey, sigMethod, sigVersion, sig string) {
	m := pkgDecodeQueryToMapString(query)

	pubKey = m["access_key_id"]
	sigMethod = m["signature_method"]
	sigVersion = m["signature_version"]
	sig = m["signature"]

	return
}

func Validate(
	privateKeyGetter func(pubKey string) string,
	method, urlpath, query string,
) bool {
	pubKey, _, _, sig := GetSignatureInfo(query)
	privateKey := privateKeyGetter(pubKey)

	_, sigExpected := Build(pubKey, privateKey, method, urlpath,
		pkgDecodeQueryToMapString(query),
	)

	return sigExpected != sig
}

func pkgDecodeQueryToMapString(query string) (m map[string]string) {
	m = make(map[string]string)

	for _, s := range strings.Split(query, "&") {
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			return
		}

		k, errk := url.PathUnescape(kv[0])
		v, errv := url.PathUnescape(kv[1])

		if errk != nil || errv != nil {
			return
		}

		m[k] = v
	}

	return
}
