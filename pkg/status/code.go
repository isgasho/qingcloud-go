// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// https://docs.qingcloud.com/api/common/error_code.html

package status

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

type (
	JobStatus          string
	InstanceStatus     string
	LoadBalancerStatus string
)

const (
	JobStatus_Unknown    JobStatus = "unknown"
	JobStatus_Successful JobStatus = "successful"
	JobStatus_Failed     JobStatus = "failed"
	JobStatus_Pending    JobStatus = "pending"
	JobStatus_Working    JobStatus = "working"
)

const (
	InstanceStatus_Pending    InstanceStatus = "pending"
	InstanceStatus_Running    InstanceStatus = "running"
	InstanceStatus_Stopped    InstanceStatus = "stopped"
	InstanceStatus_Suspended  InstanceStatus = "suspended"
	InstanceStatus_Terminated InstanceStatus = "terminated"
	InstanceStatus_Ceased     InstanceStatus = "ceased"
)

const (
	LoadBalancerStatus_Pending   LoadBalancerStatus = "pending"
	LoadBalancerStatus_Active    LoadBalancerStatus = "active"
	LoadBalancerStatus_Stopped   LoadBalancerStatus = "stopped"
	LoadBalancerStatus_Suspended LoadBalancerStatus = "suspended"
	LoadBalancerStatus_Deleted   LoadBalancerStatus = "deleted"
	LoadBalancerStatus_Ceased    LoadBalancerStatus = "ceased"
)
