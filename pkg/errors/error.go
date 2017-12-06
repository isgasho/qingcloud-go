// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
)

const (
	ErrCode_NULL = 0

	// Client
	ErrCode_1100 = 1100
	ErrCode_1200 = 1200
	ErrCode_1300 = 1300
	ErrCode_1400 = 1400

	ErrCode_2100 = 2100
	ErrCode_2400 = 2400
	ErrCode_2500 = 2500

	// Server
	ErrCode_5000 = 5000
	ErrCode_5100 = 5100
	ErrCode_5200 = 5200
	ErrCode_5300 = 5300
)

type Error interface {
	GetRetCode() int32
	GetMessage() string
	error
}

func New(code int32, msg string) error {
	return &qcError{code: code, desc: msg}
}

func Newf(code int32, format string, a ...interface{}) error {
	return &qcError{code: code, desc: fmt.Sprintf(format, a...)}
}

type qcError struct {
	code int32
	desc string
}

func (p *qcError) GetRetCode() int32 {
	return p.code
}

func (p *qcError) GetMessage() string {
	return p.desc
}

func (p *qcError) Error() string {
	return fmt.Sprintf("error: code = %v, desc = %v", p.code, p.desc)
}
