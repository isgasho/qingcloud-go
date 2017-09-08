
<p align="center"><a href="http://qingcloud.com"><img src="https://raw.githubusercontent.com/chai2010/qingcloud-go/master/logo.jpg" alt="QingCloud"></a></p>

# 青云 SDK Go Version (兼容 [yunify/qingcloud-sdk-go](https://github.com/yunify/qingcloud-sdk-go))

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


## 设计思路

- 基于 Protobuf3 语法对应的 json 格式文件定义 api 元数据
- 基于 Protobuf3 语法定义 rest 接口的请求和响应结构体, 只是用于结构体, 请求时转 json 处理(proto库自带)
- 基于 Protobuf3 的扩展特性, 增加自定义的元数据
- 内置 Go 语言代码模板基于标准库语法


## 可运行的例子

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


## protobuf 扩展信息

扩展类型在 [spec.pb/qingcloud_sdk_rule/rule.proto](spec.pb/qingcloud_sdk_rule/rule.proto) 文件中定义:

```proto
message MethodRule {
	string http_action = 1;
}

message MethodInputRule {
	string required_fileds = 1; // 格式: "a, b, ..."
	string default_value = 2;   // 格式: "a:v, b:v, ..."
	string enum_value = 3;      // 格式: "a:a1,a2,a3; b:b1,b2; ..."
}

// 通过扩展信息给 method 增加约束
extend google.protobuf.MethodOptions {
	MethodRule method_rule = 10001;
}

// 通过扩展信息给 message 增加约束
extend google.protobuf.MessageOptions {
	MethodInputRule method_input_rule = 10002;
}
```

通过扩展数据可以改变方法的行为, [spec.pb/user_data.proto](./spec.pb/user_data.proto) 指定了 POST 方法:

```proto
service UserDataService {
	rpc UploadUserDataAttachment(UploadUserDataAttachmentInput) returns (UploadUserDataAttachmentOutput) {
		option (qingcloud.sdk.rule.method_rule) = {
			http_action: "POST"
		};
	}
}

message UploadUserDataAttachmentInput {
	bytes attachment_content = 2;
	string attachment_name = 1;

	option (qingcloud.sdk.rule.method_input_rule) = {
		required_fileds: "attachment_content"
		default_value: ""
		enum_value: ""
	};
}
```

扩展信息中的 `http_action` 用于表示 HTTP 方法, 默认是 GET, 极少数的API是 POST ([UploadUserDataAttachment](https://docs.qingcloud.com/api/userdata/upload_userdata_attachment.html)). 当采用 POST 方法时,
需要明确指定 `http_action`.

输入的参数可以通过 `option (qingcloud.sdk.rule.method_input_rule)` 扩展来定义额外的约束, 主要是针对 必须成员/默认值/枚举字符串 几种类型.
不过目前还没有使用该信息.


## 外部依赖

- [github.com/golang/protobuf](https://github.com/golang/protobuf)

稍后会放到 vendor 目录

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

- [x] [DescribeVolumes](spec.pb/volume.proto)
- [x] [CreateVolumes](spec.pb/volume.proto)
- [x] [DeleteVolumes](spec.pb/volume.proto)
- [x] [AttachVolumes](spec.pb/volume.proto)
- [x] [DetachVolumes](spec.pb/volume.proto)
- [x] [ResizeVolumes](spec.pb/volume.proto)
- [x] [ModifyVolumeAttributes](spec.pb/volume.proto)

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

