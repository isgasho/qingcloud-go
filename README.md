
<p align="center"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/logo.jpg" alt="QingCloud"></p>

# 青云 SDK Go Version

[![Build Status](https://travis-ci.org/chai2010/qingcloud-go.svg)](https://travis-ci.org/chai2010/qingcloud-go)
[![GoDoc](https://godoc.org/github.com/chai2010/qingcloud-go?status.svg)](https://godoc.org/github.com/chai2010/qingcloud-go)
[![API Reference](http://img.shields.io/badge/api-reference-green.svg)](http://docs.qingcloud.com)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/chai2010/qingcloud-go/blob/master/LICENSE)

在线文档:

- https://godoc.org/github.com/chai2010/qingcloud-go
- https://godoc.org/github.com/chai2010/qingcloud-go/config
- https://godoc.org/github.com/chai2010/qingcloud-go/request
- https://godoc.org/github.com/chai2010/qingcloud-go/spec.pb
- https://docs.qingcloud.com

可运行的例子:

- [examples/nic/DescribeNics/DescribeNics.go](./examples/nic/DescribeNics/DescribeNics.go)
- [examples/mongo/DescribeMongos/DescribeMongos.go](./examples/mongo/DescribeMongos/DescribeMongos.go)
- [examples/cluster/DescribeClusters/DescribeClusters.go](./examples/cluster/DescribeClusters/DescribeClusters.go)


## 运行流程

当前用户的默认配置文件在 `${HOME}/.qingcloud/config.yaml`, 内容如下:

```yaml
# QingCloud services configuration

qy_access_key_id: 'ACCESS_KEY_ID'
qy_secret_access_key: 'SECRET_ACCESS_KEY'

host: 'api.qingcloud.com'
port: 443
protocol: 'https'
uri: '/iaas'
connection_retries: 3

json_disable_unknown_fields: false

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: 'warn'
```

将 `qy_access_key_id` 和 `qy_secret_access_key` 字段替换为 API密钥 中的内容.

运行例子 [examples/nic/DescribeNics/DescribeNics.go](./examples/nic/DescribeNics/DescribeNics.go):

	go run ./examples/nic/DescribeNics/DescribeNics.go

接口规范文件在 [spec.pb/nic.proto](./spec.pb/nic.proto).

生成器代码在 [protoc-gen-go/qingcloud/qingcloud.go](protoc-gen-go/qingcloud/qingcloud.go).

*注: 目前只实现了该接口*

## 设计思路

- 基于 Protobuf3 语法对应的 json 格式文件定义 api 元数据
- 基于 Protobuf3 语法定义 rest 接口的请求和响应结构体, 只是用于结构体, 请求时转 json 处理(proto库自带)
- 内置 Go 语言代码模板基于标准库语法

小目标:

- api 元数据采用语言无关的 Proto3 语法定义, 文件采用 proto3 对应的 json 格式
- 生成的请求和响应结构体保持和 Proto3 生成的结构一致, Proto3 本身支持多种语言
- 更多的单元测试代码
- 更多的示例代码
- 最小化外部依赖

## protobuf 扩展信息

- [spec.pb/user_data.proto](./spec.pb/user_data.proto):

```proto
service UserDataService {
	rpc UploadUserDataAttachment(UploadUserDataAttachmentInput) returns (UploadUserDataAttachmentOutput) {
		option (google.api.http) = {
			custom { kind: "POST" }
		};
	}
}
```

扩展信息中的 `custom.kind` 用于表示 HTTP 方法, 默认是 GET, 极少数的API是 POST ([UploadUserDataAttachment](https://docs.qingcloud.com/api/userdata/upload_userdata_attachment.html)).

当采用 POST 方法时, 需要明确指定 `custom.kind`.

## 外部依赖

- [github.com/golang/glog](https://github.com/golang/glog)
- [github.com/golang/protobuf](https://github.com/golang/protobuf)

其中 glog 是全局的日志包, SDK 不能放到 vendor 目录(可能导致多个glog冲突).

-------

## API 状态

如下是请求的 API 列表，目前的状态：

### 主机

- [ ] DescribeInstances
- [ ] RunInstances
- [ ] TerminateInstances
- [ ] StartInstances
- [ ] StopInstances
- [ ] RestartInstances
- [ ] ResetInstances
- [ ] ResizeInstances
- [ ] ModifyInstanceAttributes
- [x] [DescribeInstanceTypes](spec.pb/instances.proto)
- [ ] CreateBrokers
- [ ] DeleteBrokers

### 硬盘

- [ ] DescribeVolumes
- [ ] CreateVolumes
- [ ] DeleteVolumes
- [ ] AttachVolumes
- [ ] DetachVolumes
- [ ] ResizeVolumes
- [ ] ModifyVolumeAttributes

## 网卡

- [ ] CreateNics
- [x] [DescribeNics](spec.pb/nic.proto)
- [ ] AttachNics
- [ ] DetachNics
- [ ] ModifyNicAttributes
- [ ] DeleteNics


### 私有网络

- [ ] DescribeVxnets
- [ ] CreateVxnets
- [ ] DeleteVxnets
- [ ] JoinVxnet
- [ ] LeaveVxnet
- [ ] ModifyVxnetAttributes
- [ ] DescribeVxnetInstances

### 路由器

- [ ] DescribeRouters
- [ ] CreateRouters
- [ ] DeleteRouters
- [ ] UpdateRouters
- [ ] PowerOffRouters
- [ ] PowerOnRouters
- [ ] JoinRouter
- [ ] LeaveRouter
- [ ] ModifyRouterAttributes
- [ ] DescribeRouterStatics
- [ ] AddRouterStatics
- [ ] ModifyRouterStaticAttributes
- [ ] DeleteRouterStatics
- [ ] DescribeRouterVxnets
- [ ] AddRouterStaticEntries
- [ ] DeleteRouterStaticEntries
- [ ] ModifyRouterStaticEntryAttributes
- [ ] DescribeRouterStaticEntries

### 公网IP

- [ ] DescribeEips
- [ ] AllocateEips
- [ ] ReleaseEips
- [ ] AssociateEip
- [ ] DissociateEips
- [ ] ChangeEipsBandwidth
- [ ] ChangeEipsBillingMode
- [ ] ModifyEipAttributes

### 防火墙

- [ ] DescribeSecurityGroups
- [ ] CreateSecurityGroup
- [ ] DeleteSecurityGroups
- [ ] ApplySecurityGroup
- [ ] ModifySecurityGroupAttributes
- [ ] DescribeSecurityGroupRules
- [ ] AddSecurityGroupRules
- [ ] DeleteSecurityGroupRules
- [ ] ModifySecurityGroupRuleAttributes
- [ ] CreateSecurityGroupSnapshot
- [ ] DescribeSecurityGroupSnapshots
- [ ] DeleteSecurityGroupSnapshots
- [ ] RollbackSecurityGroup

### SSH 密钥

- [ ] DescribeKeyPairs
- [ ] CreateKeyPair
- [ ] DeleteKeyPairs
- [ ] AttachKeyPairs
- [ ] DetachKeyPairs
- [ ] ModifyKeyPairAttributes

### 映像

- [ ] DescribeImages
- [ ] CaptureInstance
- [ ] DeleteImages
- [ ] ModifyImageAttributes
- [ ] GrantImageToUsers
- [ ] RevokeImageFromUsers
- [ ] DescribeImageUsers

### 负载均衡

- [ ] CreateLoadBalancer
- [ ] DescribeLoadBalancers
- [ ] DeleteLoadBalancers
- [ ] ModifyLoadBalancerAttributes
- [ ] StartLoadBalancers
- [ ] StopLoadBalancers
- [ ] UpdateLoadBalancers
- [ ] ResizeLoadBalancers
- [ ] AssociateEipsToLoadBalancer
- [ ] DissociateEipsFromLoadBalancer
- [ ] AddLoadBalancerListeners
- [ ] DescribeLoadBalancerListeners
- [ ] DeleteLoadBalancerListeners
- [ ] ModifyLoadBalancerListenerAttributes
- [ ] AddLoadBalancerBackends
- [ ] DescribeLoadBalancerBackends
- [ ] DeleteLoadBalancerBackends
- [ ] ModifyLoadBalancerBackendAttributes
- [ ] CreateLoadBalancerPolicy
- [ ] DescribeLoadBalancerPolicies
- [ ] ModifyLoadBalancerPolicyAttributes
- [ ] ApplyLoadBalancerPolicy
- [ ] DeleteLoadBalancerPolicies
- [ ] AddLoadBalancerPolicyRules
- [ ] DescribeLoadBalancerPolicyRules
- [ ] ModifyLoadBalancerPolicyRuleAttributes
- [ ] DeleteLoadBalancerPolicyRules
- [ ] CreateServerCertificate
- [ ] DescribeServerCertificates
- [ ] ModifyServerCertificateAttributes
- [ ] DeleteServerCertificates

### 资源监控

- [ ] GetMonitor
- [ ] GetLoadBalancerMonitor
- [ ] GetRDBMonitor
- [ ] GetCacheMonitor
- [ ] GetZooKeeperMonitor
- [ ] GetQueueMonitor

### 备份

- [ ] DescribeSnapshots
- [ ] CreateSnapshots
- [ ] DeleteSnapshots
- [ ] ApplySnapshots
- [ ] ModifySnapshotAttributes
- [ ] CaptureInstanceFromSnapshot
- [ ] CreateVolumeFromSnapshot

### User Data

- [x] [UploadUserDataAttachment](spec.pb/user_data.proto)

### 内网域名别名

- [ ] DescribeDNSAliases
- [ ] AssociateDNSAlias
- [ ] DissociateDNSAliases
- [ ] GetDNSLabel

### 操作日志

- [x] [DescribeJobs](spec.pb/job.proto)

### 标签

- [ ] DescribeTags
- [ ] CreateTag
- [ ] DeleteTags
- [ ] ModifyTagAttributes
- [ ] AttachTags
- [ ] DetachTags

### 区域

- [x] [DescribeZones](spec.pb/zone.proto)

### 数据库

- [ ] CreateRDB
- [ ] DescribeRDBs
- [ ] DeleteRDBs
- [ ] StartRDBs
- [ ] StopRDBs
- [ ] ResizeRDBs
- [ ] RDBsLeaveVxnet
- [ ] RDBsJoinVxnet
- [ ] CreateRDBFromSnapshot
- [ ] CreateTempRDBInstanceFromSnapshot
- [ ] GetRDBInstanceFiles
- [ ] CopyRDBInstanceFilesToFTP
- [ ] CeaseRDBInstance
- [ ] CreateTempRDBInstanceFromSnapshot
- [ ] GetRDBMonitor
- [ ] ModifyRDBParameters
- [ ] ApplyRDBParameterGroup
- [ ] DescribeRDBParameters

### Mongo 集群

- [ ] DescribeMongoNodes
- [ ] DescribeMongoParameters
- [ ] ResizeMongos
- [ ] CreateMongo
- [ ] StopMongos
- [ ] StartMongos
- [ ] DescribeMongos
- [ ] DeleteMongos
- [ ] CreateMongoFromSnapshot
- [ ] ChangeMongoVxnet
- [ ] AddMongoInstances
- [ ] RemoveMongoInstances
- [ ] ModifyMongoAttributes
- [ ] ModifyMongoInstances
- [ ] GetMongoMonitor

### 缓存服务

- [ ] DescribeCaches
- [ ] CreateCache
- [ ] StopCaches
- [ ] StartCaches
- [ ] RestartCaches
- [ ] DeleteCaches
- [ ] ResizeCaches
- [ ] UpdateCache
- [ ] ChangeCacheVxnet
- [ ] ModifyCacheAttributes
- [ ] DescribeCacheNodes
- [ ] AddCacheNodes
- [ ] DeleteCacheNodes
- [ ] RestartCacheNodes
- [ ] ModifyCacheNodeAttributes
- [ ] CreateCacheFromSnapshot
- [ ] DescribeCacheParameterGroups
- [ ] CreateCacheParameterGroup
- [ ] ApplyCacheParameterGroup
- [ ] DeleteCacheParameterGroups
- [ ] ModifyCacheParameterGroupAttributes
- [ ] DescribeCacheParameters
- [ ] UpdateCacheParameters
- [ ] ResetCacheParameters

### Virtual SAN

- [ ] CreateS2Server
- [ ] DescribeS2Servers
- [ ] ModifyS2Server
- [ ] ResizeS2Servers
- [ ] DeleteS2Servers
- [ ] PowerOnS2Servers
- [ ] PowerOffS2Servers
- [ ] UpdateS2Servers
- [ ] ChangeS2ServerVxnet
- [ ] CreateS2SharedTarget
- [ ] DescribeS2SharedTargets
- [ ] DeleteS2SharedTargets
- [ ] EnableS2SharedTargets
- [ ] DisableS2SharedTargets
- [ ] ModifyS2SharedTargets
- [ ] AttachToS2SharedTarget
- [ ] DetachFromS2SharedTarget
- [ ] DescribeS2DefaultParameters

### Spark

- [ ] AddSparkNodes
- [ ] DeleteSparkNodes
- [ ] StartSparks
- [ ] StopSparks

### Hadoop 服务

- [ ] AddHadoopNodes
- [ ] DeleteHadoopNodes
- [ ] StartHadoops
- [ ] StopHadoops

### 集群服务

- [x] [CreateCluster](spec.pb/cluster)
- [x] [DescribeClusters](spec.pb/cluster)
- [x] [DescribeClusterNodes](spec.pb/cluster)
- [ ] StopClusters
- [ ] StartClusters
- [ ] DeleteClusters
- [ ] lease
- [ ] AddClusterNodes
- [ ] DeleteClusterNodes
- [ ] ResizeCluster
- [ ] ChangeClusterVxnet
- [ ] SuspendClusters
- [ ] UpdateClusterEnvironment
- [ ] ModifyClusterAttributes
- [ ] ModifyClusterNodeAttributes
- [ ] GetClustersStats
- [ ] DescribeClusterUsers
- [ ] RestartClusterService
- [ ] UpgradeClusters
- [ ] AuthorizeClustersBrokerToDeveloper
- [ ] RevokeClustersBrokerFromDeveloper


### 资源协作中心

- [ ] DescribeSharedResourceGroups
- [ ] DescribeResourceGroups
- [ ] CreateResourceGroups
- [ ] ModifyResourceGroupAttributes
- [ ] DeleteResourceGroups
- [ ] DescribeResourceGroupItems
- [ ] AddResourceGroupItems
- [ ] DeleteResourceGroupItems
- [ ] DescribeUserGroups
- [ ] CreateUserGroups
- [ ] ModifyUserGroupAttributes
- [ ] DeleteUserGroups
- [ ] DescribeUserGroupMembers
- [ ] AddUserGroupMembers
- [ ] ModifyUserGroupMemberAttributes
- [ ] DeleteUserGroupMembers
- [ ] DescribeGroupRoles
- [ ] CreateGroupRoles
- [ ] ModifyGroupRoleAttributes
- [ ] DeleteGroupRoles
- [ ] DescribeGroupRoleRules
- [ ] AddGroupRoleRules
- [ ] ModifyGroupRoleRuleAttributes
- [ ] DeleteGroupRoleRules
- [ ] GrantResourceGroupsToUserGroups
- [ ] RevokeResourceGroupsFromUserGroups
- [ ] DescribeResourceUserGroups

### 消息中心

- [ ] DescribeNotificationCenterUserPosts

