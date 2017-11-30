// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/chai2010/qingcloud-go/pkg/request"
	"github.com/chai2010/qingcloud-go/pkg/request/data"
)

// QingCloudService: QingCloud provides a platform which can make the delivery of computing resources more simple, efficient and reliable, even more environmental.
type QingCloudService struct {
	Config     *config.Config
	Properties *QingCloudServiceProperties
}

type QingCloudServiceProperties struct {
}

func Init(c *config.Config) (*QingCloudService, error) {
	properties := &QingCloudServiceProperties{}
	return &QingCloudService{Config: c, Properties: properties}, nil
}

// Documentation URL: https://docs.qingcloud.com/api/zone/describe_zones.html
func (s *QingCloudService) DescribeZones(i *DescribeZonesInput) (*DescribeZonesOutput, error) {
	if i == nil {
		i = &DescribeZonesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET",
	}

	x := &DescribeZonesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
