// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// See https://docs.qingcloud.com/api/common/signature.html

package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

func Build(pubKey, privateKey, method, urlpath string, m map[string]string) (query, signature string) {
	pubKey = strings.TrimSpace(pubKey)
	privateKey = strings.TrimSpace(privateKey)
	method = strings.TrimSpace(method)
	urlpath = strings.TrimSpace(urlpath)
	m = cloneMapStringString(m)

	m["access_key_id"] = pubKey
	m["signature_method"] = "HmacSHA256"
	m["signature_version"] = "1"

	sortedQueryString := makeSortedUrlQueryString(m)
	stringToSign := makeStringToSign(method, urlpath, sortedQueryString)

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(stringToSign))

	signature = strings.TrimSpace(base64.StdEncoding.EncodeToString(h.Sum(nil)))

	signature = strings.Replace(signature, " ", "+", -1)
	signature = url.QueryEscape(signature)

	sortedQueryString = sortedQueryString + "&signature=" + signature
	return sortedQueryString, signature
}

func makeStringToSign(method, urlpath, sortedQueryString string) string {
	return method + "\n" + urlpath + "\n" + sortedQueryString
}

func makeSortedUrlQueryString(m map[string]string) string {
	keys, values := getMapSortedKeyValues(m)

	var parts []string
	for i := 0; i < len(keys) && i < len(values); i++ {
		ki, vi := keys[i], values[i]
		vi = strings.Replace(url.QueryEscape(vi), "+", "%20", -1)
		parts = append(parts, ki+"="+vi)
	}
	return strings.Join(parts, "&")
}

func getMapSortedKeyValues(m map[string]string) (keys, values []string) {
	keys = make([]string, 0, len(m))
	values = make([]string, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		values = append(values, m[k])
	}

	return
}

func cloneMapStringString(m0 map[string]string) map[string]string {
	var m1 = make(map[string]string)
	for k, v := range m0 {
		if k = strings.TrimSpace(k); k != "" {
			m1[k] = v
		}
	}
	return m1
}
