// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

// Package errors provides detailed error types for api field validation.
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

type QingCloudErrorInterface interface {
	GetRetCode() int32
	GetMessage() string
	error
}

// QingCloudError stores information of a QingCloud error response.
type QingCloudError struct {
	RetCode int    `json:"ret_code"`
	Message string `json:"message"`
}

// Error returns the description of QingCloud error response.
func (ise QingCloudError) Error() string {
	return fmt.Sprintf("QingCloud Error: Code (%d), Message (%s)", ise.RetCode, ise.Message)
}
