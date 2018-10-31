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

func WaitLoadBalancerStatus(
	server *pb.ServerInfo, lbId string,
	status statuspkg.LoadBalancerStatus,
	timeout time.Duration,
) error {
	return WaitForIntervalWorkDone(
		fmt.Sprintf("lb:%v", lbId), timeout, func() (done bool, err error) {
			return waitLoadBalancerStatus(server, lbId, status)
		},
	)
}

func waitLoadBalancerStatus(
	server *pb.ServerInfo, lbId string,
	expectedStatus statuspkg.LoadBalancerStatus,
) (done bool, err error) {
	reply, err := pb.NewLoadBalancerService(server).DescribeLoadBalancers(&pb.DescribeLoadBalancersInput{
		Loadbalancers: []string{lbId},
	})
	if err != nil {
		return false, err
	}

	if len(reply.GetLoadbalancerSet()) != 1 {
		return false, fmt.Errorf("can not find lb [%s]", lbId)
	}

	insStatus := statuspkg.LoadBalancerStatus(reply.GetLoadbalancerSet()[0].GetStatus())
	if insStatus == expectedStatus {
		return true, nil
	}

	return false, nil
}
