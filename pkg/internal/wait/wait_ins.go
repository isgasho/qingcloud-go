// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package wait

import (
	"fmt"
	"time"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	statuspkg "github.com/chai2010/qingcloud-go/pkg/status"
)

func WaitInstanceStatus(
	server *pb.ServerInfo, insId string,
	status statuspkg.InstanceStatus,
	timeout time.Duration,
) error {
	return WaitForIntervalWorkDone(
		fmt.Sprintf("ins:%v", insId), timeout, func() (done bool, err error) {
			return waitInstanceStatus(server, insId, status)
		},
	)
}

func WaitInstanceNetwork(
	server *pb.ServerInfo, insId string,
	timeout time.Duration,
) error {
	return WaitForIntervalWorkDone(
		fmt.Sprintf("ins:%v", insId), timeout, func() (done bool, err error) {
			return waitInstanceNetwork(server, insId)
		},
	)
}

func waitInstanceStatus(
	server *pb.ServerInfo, insId string,
	expectedStatus statuspkg.InstanceStatus,
) (done bool, err error) {
	reply, err := pb.NewInstanceService(server).DescribeInstances(&pb.DescribeInstancesInput{
		Instances: []string{insId},
	})
	if err != nil {
		return false, err
	}

	if len(reply.GetInstanceSet()) != 1 {
		return false, fmt.Errorf("can not find instance [%s]", insId)
	}

	insStatus := statuspkg.InstanceStatus(reply.GetInstanceSet()[0].GetStatus())
	if insStatus == expectedStatus {
		return true, nil
	}

	return false, nil
}

func waitInstanceNetwork(
	server *pb.ServerInfo, insId string,
) (done bool, err error) {
	reply, err := pb.NewInstanceService(server).DescribeInstances(&pb.DescribeInstancesInput{
		Instances: []string{insId},
	})
	if err != nil {
		return false, err
	}

	if len(reply.GetInstanceSet()) != 1 {
		return false, fmt.Errorf("can not find instance [%s]", insId)
	}

	vxnets := reply.GetInstanceSet()[0].GetVxnets()
	return len(vxnets) > 0 && vxnets[0].GetPrivateIp() != "", nil
}
