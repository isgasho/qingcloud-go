// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package wait

import (
	"fmt"
	"time"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
)

type JobStatus string

const (
	JobStatus_Unknown    JobStatus = "unknown"
	JobStatus_Successful JobStatus = "successful"
	JobStatus_Failed     JobStatus = "failed"
	JobStatus_Pending    JobStatus = "pending"
	JobStatus_Working    JobStatus = "working"
)

func WaitJob(server *pb.ServerInfo, jobId string, timeout time.Duration) error {
	return WaitForIntervalWorkDone(
		fmt.Sprintf("job:%v", jobId), timeout, func() (done bool, err error) {
			return waitJob(server, jobId)
		},
	)
}

func waitJob(server *pb.ServerInfo, jobId string) (done bool, err error) {
	reply, err := pb.NewJobService(server).DescribeJobs(&pb.DescribeJobsInput{
		Jobs: []string{jobId},
	})
	if err != nil {
		return false, err
	}

	if len(reply.GetJobSet()) != 1 {
		return false, fmt.Errorf("can not find job [%s]", jobId)
	}

	status := JobStatus(reply.GetJobSet()[0].GetStatus())
	switch status {
	case JobStatus_Successful:
		return true, nil // OK
	case JobStatus_Pending, JobStatus_Working:
		return false, nil
	case JobStatus_Failed:
		return false, fmt.Errorf("job [%s] failed", jobId)
	default:
		return false, fmt.Errorf("unknow status [%s] for job [%s]", status, jobId)
	}
}
