// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package signature

import (
	"testing"
)

//func Build(pubKey, privateKey, method, urlpath string, m map[string]string) (query, signature string) {
func TestBuild(t *testing.T) {
	_, signature := Build(
		"QYACCESSKEYIDEXAMPLE", "SECRETACCESSKEY",
		"GET", "/iaas/", tMap1_unsorted,
	)
	if signature != tSignature {
		t.Fatalf("expected = %s, got = %s", tSignature, signature)
	}
}
func TestUtils(t *testing.T) {
	sortedQueryString := makeSortedUrlQueryString(tMap1_unsorted)
	if sortedQueryString != tSortedQueryString {
		t.Fatalf("expected = %s, got = %s", tSortedQueryString, sortedQueryString)
	}

	sToSign := makeStringToSign("GET", "/iaas/", sortedQueryString)
	if sToSign != tStringToSign {
		t.Fatalf("expected = %s, got = %s", tStringToSign, sToSign)
	}
}

func TestMakeSortedUrlQueryString2(t *testing.T) {
	s := makeSortedUrlQueryString(tMap1_unsorted)
	if s != tSortedQueryString {
		t.Fatalf("expected = %s, got = %s", tSortedQueryString, s)
	}
}

var tMap1_unsorted = map[string]string{
	"count":             "1",
	"vxnets.1":          "vxnet-0",
	"zone":              "pek1",
	"instance_type":     "small_b",
	"signature_version": "1",
	"signature_method":  "HmacSHA256",
	"instance_name":     "demo",
	"image_id":          "centos64x86a",
	"login_mode":        "passwd",
	"login_passwd":      "QingCloud20130712",
	"version":           "1",
	"access_key_id":     "QYACCESSKEYIDEXAMPLE",
	"action":            "RunInstances",
	"time_stamp":        "2013-08-27T14:30:10Z",
}

var tMap2_sorted = map[string]string{
	"access_key_id":     "QYACCESSKEYIDEXAMPLE",
	"action":            "RunInstances",
	"count":             "1",
	"image_id":          "centos64x86a",
	"instance_name":     "demo",
	"instance_type":     "small_b",
	"login_mode":        "passwd",
	"login_passwd":      "QingCloud20130712",
	"signature_method":  "HmacSHA256",
	"signature_version": "1",
	"time_stamp":        "2013-08-27T14:30:10Z",
	"version":           "1",
	"vxnets.1":          "vxnet-0",
	"zone":              "pek1",
}

var tMap3_url_encoded = map[string]string{
	"access_key_id":     "QYACCESSKEYIDEXAMPLE",
	"action":            "RunInstances",
	"count":             "1",
	"image_id":          "centos64x86a",
	"instance_name":     "demo",
	"instance_type":     "small_b",
	"login_mode":        "passwd",
	"login_passwd":      "QingCloud20130712",
	"signature_method":  "HmacSHA256",
	"signature_version": "1",
	"time_stamp":        "2013-08-27T14%3A30%3A10Z",
	"version":           "1",
	"vxnets.1":          "vxnet-0",
	"zone":              "pek1",
}

var tSortedQueryString = `access_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1`

var tStringToSign = `GET
/iaas/
access_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1`

var tSignature = `32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D`
