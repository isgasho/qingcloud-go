// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/go-sdk
// source: cluster.proto

package service

import proto "github.com/golang/protobuf/proto"
import "fmt"

import "github.com/chai2010/qingcloud-go/pkg/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = proto.Marshal

var _ = client.NewClient

type ClusterServiceInterface interface {
	CreateCluster(in *CreateClusterInput) (out *CreateClusterOutput, err error)
	DescribeClusters(in *DescribeClustersInput) (out *DescribeClustersOutput, err error)
	DescribeClusterNodes(in *DescribeClusterNodesInput) (out *DescribeClusterNodesOutput, err error)
	StopClusters(in *StopClustersInput) (out *StopClustersOutput, err error)
	StartClusters(in *StartClustersInput) (out *StartClustersOutput, err error)
	DeleteClusters(in *DeleteClustersInput) (out *DeleteClustersOutput, err error)
	Lease(in *LeaseInput) (out *LeaseOutput, err error)
	AddClusterNodes(in *AddClusterNodesInput) (out *AddClusterNodesOutput, err error)
	DeleteClusterNodes(in *DeleteClusterNodesInput) (out *DeleteClusterNodesOutput, err error)
	ResizeCluster(in *ResizeClusterInput) (out *ResizeClusterOutput, err error)
	ChangeClusterVxnet(in *ChangeClusterVxnetInput) (out *ChangeClusterVxnetOutput, err error)
	SuspendClusters(in *SuspendClustersInput) (out *SuspendClustersOutput, err error)
	UpdateClusterEnvironment(in *UpdateClusterEnvironmentInput) (out *UpdateClusterEnvironmentOutput, err error)
	ModifyClusterAttributes(in *ModifyClusterAttributesInput) (out *ModifyClusterAttributesOutput, err error)
	ModifyClusterNodeAttributes(in *ModifyClusterNodeAttributesInput) (out *ModifyClusterNodeAttributesOutput, err error)
	GetClustersStats(in *GetClustersStatsInput) (out *GetClustersStatsOutput, err error)
	DescribeClusterUsers(in *DescribeClusterUsersInput) (out *DescribeClusterUsersOutput, err error)
	RestartClusterService(in *RestartClusterServiceInput) (out *RestartClusterServiceOutput, err error)
	UpgradeClusters(in *UpgradeClustersInput) (out *UpgradeClustersOutput, err error)
	AuthorizeClustersBrokerToDeveloper(in *AuthorizeClustersBrokerToDeveloperInput) (out *AuthorizeClustersBrokerToDeveloperOutput, err error)
	RevokeClustersBrokerFromDeveloper(in *RevokeClustersBrokerFromDeveloperInput) (out *RevokeClustersBrokerFromDeveloperOutput, err error)
}

type ClusterService struct {
	ServerInfo       *ServerInfo
	Properties       *ClusterServiceProperties
	LastResponseBody string
}

func NewClusterService(server *ServerInfo, serviceProp *ClusterServiceProperties) (p *ClusterService) {
	return &ClusterService{
		ServerInfo: server,
		Properties: serviceProp,
	}
}

func (p *ClusterService) CreateCluster(input *CreateClusterInput) (output *CreateClusterOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(CreateClusterOutput)

	err = client.CallMethod(nil, "CreateCluster", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) DescribeClusters(input *DescribeClustersInput) (output *DescribeClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeClustersOutput)

	err = client.CallMethod(nil, "DescribeClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) DescribeClusterNodes(input *DescribeClusterNodesInput) (output *DescribeClusterNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeClusterNodesOutput)

	err = client.CallMethod(nil, "DescribeClusterNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) StopClusters(input *StopClustersInput) (output *StopClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StopClustersOutput)

	err = client.CallMethod(nil, "StopClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) StartClusters(input *StartClustersInput) (output *StartClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(StartClustersOutput)

	err = client.CallMethod(nil, "StartClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) DeleteClusters(input *DeleteClustersInput) (output *DeleteClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteClustersOutput)

	err = client.CallMethod(nil, "DeleteClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) Lease(input *LeaseInput) (output *LeaseOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(LeaseOutput)

	err = client.CallMethod(nil, "Lease", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) AddClusterNodes(input *AddClusterNodesInput) (output *AddClusterNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AddClusterNodesOutput)

	err = client.CallMethod(nil, "AddClusterNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) DeleteClusterNodes(input *DeleteClusterNodesInput) (output *DeleteClusterNodesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DeleteClusterNodesOutput)

	err = client.CallMethod(nil, "DeleteClusterNodes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) ResizeCluster(input *ResizeClusterInput) (output *ResizeClusterOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ResizeClusterOutput)

	err = client.CallMethod(nil, "ResizeCluster", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) ChangeClusterVxnet(input *ChangeClusterVxnetInput) (output *ChangeClusterVxnetOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ChangeClusterVxnetOutput)

	err = client.CallMethod(nil, "ChangeClusterVxnet", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) SuspendClusters(input *SuspendClustersInput) (output *SuspendClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(SuspendClustersOutput)

	err = client.CallMethod(nil, "SuspendClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) UpdateClusterEnvironment(input *UpdateClusterEnvironmentInput) (output *UpdateClusterEnvironmentOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(UpdateClusterEnvironmentOutput)

	err = client.CallMethod(nil, "UpdateClusterEnvironment", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) ModifyClusterAttributes(input *ModifyClusterAttributesInput) (output *ModifyClusterAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyClusterAttributesOutput)

	err = client.CallMethod(nil, "ModifyClusterAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) ModifyClusterNodeAttributes(input *ModifyClusterNodeAttributesInput) (output *ModifyClusterNodeAttributesOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(ModifyClusterNodeAttributesOutput)

	err = client.CallMethod(nil, "ModifyClusterNodeAttributes", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) GetClustersStats(input *GetClustersStatsInput) (output *GetClustersStatsOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(GetClustersStatsOutput)

	err = client.CallMethod(nil, "GetClustersStats", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) DescribeClusterUsers(input *DescribeClusterUsersInput) (output *DescribeClusterUsersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(DescribeClusterUsersOutput)

	err = client.CallMethod(nil, "DescribeClusterUsers", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) RestartClusterService(input *RestartClusterServiceInput) (output *RestartClusterServiceOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(RestartClusterServiceOutput)

	err = client.CallMethod(nil, "RestartClusterService", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) UpgradeClusters(input *UpgradeClustersInput) (output *UpgradeClustersOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(UpgradeClustersOutput)

	err = client.CallMethod(nil, "UpgradeClusters", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) AuthorizeClustersBrokerToDeveloper(input *AuthorizeClustersBrokerToDeveloperInput) (output *AuthorizeClustersBrokerToDeveloperOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(AuthorizeClustersBrokerToDeveloperOutput)

	err = client.CallMethod(nil, "AuthorizeClustersBrokerToDeveloper", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (p *ClusterService) RevokeClustersBrokerFromDeveloper(input *RevokeClustersBrokerFromDeveloperInput) (output *RevokeClustersBrokerFromDeveloperOutput, err error) {
	client := client.NewClient("", "", nil)
	output = new(RevokeClustersBrokerFromDeveloperOutput)

	err = client.CallMethod(nil, "RevokeClustersBrokerFromDeveloper", input, output, nil)
	if err != nil {
		return nil, err
	}

	return
}
