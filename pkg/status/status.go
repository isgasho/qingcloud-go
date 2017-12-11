// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package status

import (
	"fmt"
	"net/http"
)

type Status interface {
	GetRetCode() int32 // 0 and http.StatusOK is OK
	GetMessage() string
}

func New(code int32, msg string) Status {
	return &pkgStatus{code: code, desc: msg}
}

func Newf(code int32, format string, a ...interface{}) Status {
	return &pkgStatus{code: code, desc: fmt.Sprintf(format, a...)}
}

func NewError(code int32, msg string) error {
	if code == 0 || code == int32(http.StatusOK) {
		return nil
	}

	return &pkgError{
		pkgStatus: pkgStatus{
			code: code,
			desc: msg,
		},
	}
}

func NewErrorf(code int32, format string, a ...interface{}) error {
	if code == 0 || code == int32(http.StatusOK) {
		return nil
	}

	return &pkgError{
		pkgStatus: pkgStatus{
			code: code,
			desc: fmt.Sprintf(format, a...),
		},
	}
}

func Error(s Status) error {
	if s == nil {
		return nil
	}
	if v := s.GetRetCode(); v == 0 || v == int32(http.StatusOK) {
		return nil
	}

	if err, ok := s.(error); ok {
		return err
	}

	return &pkgError{
		pkgStatus: pkgStatus{
			code: s.GetRetCode(),
			desc: s.GetMessage(),
		},
	}
}

type pkgStatus struct {
	code int32
	desc string
}

func (p *pkgStatus) GetRetCode() int32 {
	return p.code
}

func (p *pkgStatus) GetMessage() string {
	return p.desc
}

type pkgError struct {
	pkgStatus
}

func (p *pkgError) Error() string {
	return fmt.Sprintf(
		"error: code = %d, desc = %v",
		p.code,
		p.desc,
	)
}

func pkgCodeText(code int32) string {
	if s := http.StatusText(int(code)); s != "" {
		return s
	}
	return "unknown code"
}
