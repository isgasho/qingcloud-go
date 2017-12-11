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
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, ki := range keys {
		vi := strings.Replace(url.QueryEscape(m[ki]), "+", "%20", -1)
		parts = append(parts, ki+"="+vi)
	}

	return strings.Join(parts, "&")
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
