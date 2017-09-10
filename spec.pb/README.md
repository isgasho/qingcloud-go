## 设计思路

- 基于 Protobuf3 语法定义青云服务接口规范
- 基于 Protobuf3 的扩展特性, 增加自定义的元数据
- 基于 Protobuf3 生成的输入和输出参数结构进行 json 的编码和解码
- 内置基于Go标准库模板语法的Go代码生成
- 外部模板暂不支持(以后会考虑)

## 构建流程

1. 安装官方的 [`protoc`](https://github.com/google/protobuf/releases) 程序, V3版本, 带了官方的扩展类型
1. 安装官方的 `protoc-gen-go` 插件, `go install github.com/golang/protobuf/protoc-gen-go`
1. 用官方的 `protoc-gen-go` 插件生成 [qingcloud_sdk_rule/rule.proto](./qingcloud_sdk_rule/rule.proto), 参考 `make rule` 命令
1. 将官方的 `protoc-gen-go` 插件升级为支持 青云 SDK 生成的版本, `go install github.com/chai2010/qingcloud-go/protoc-gen-go`
1. 根据 spec.pb 下各个服务的 proto 文件, 用升级后的 `protoc-gen-go` 构建出青云的 SDK 代码, 放在 [../service](../service) 目录, 参考 `make` 命令
1. 在上级目录运行单元测试 `make test ./...`
1. OK

**服务规则升级流程:**

1. 更新 [qingcloud_sdk_rule/rule.proto](./qingcloud_sdk_rule/rule.proto) 文件
1. 用 `protoc-gen-go` 插件重现生成 [qingcloud_sdk_rule/rule.proto](./qingcloud_sdk_rule/rule.proto), 参考 `make rule` 命令
1. 升级 `protoc-gen-go` 插件(依赖 rule.proto 文件生成的代码), `go install github.com/chai2010/qingcloud-go/protoc-gen-go`
1. 根据新的规则更新相应的 proto 接口文件, 并运行 `make`

**服务升级流程:**

1. 更新已有的 proto 接口文件, 或者新建 proto 接口文件
1. 运行 `make`

## Protobuf 扩展信息

扩展类型在 [spec.pb/qingcloud_sdk_rule/rule.proto](./qingcloud_sdk_rule/rule.proto) 文件中定义:

```proto
// 服务规则
// 有主服务和子服务之分, 子服务隶属于某个主服务
message ServiceRule {
	string doc_url = 1;          // 文档链接
	string service_name = 2;     // 主服务名(因为要生成一个Init函数, 只能有一个, 否则会重名)
	string sub_service_name = 3; // 子服务名(主服务可省略)
}

// 方法规则
message MethodRule {
	string doc_url = 1;          // 文档链接
	string http_method = 2;      // http 行为有 GET 和 POST 之分, 默认是 GET
}

// 输入参数规则
// 输入参数成员只有数值类型和字符串, 以及对应的数组, 不含嵌套结构
message MethodInputRule {
	string required_fileds = 1;  // 格式: "a, b, ..."
	string default_value = 2;    // 格式: "a:v, b:v, ..."
	string enum_value = 3;       // 格式: "a:a1,a2,a3; b:b1,b2; ..."
	string minimum = 4;          // 格式: "a:v, b:v, ...", 最小值, 仅数值类型或数组
	string maximum = 5;          // 格式: "a:v, b:v, ...", 最大值, 仅数值类型或数组
	string multipleOf = 6;       // 格式: "a:v, b:v, ...", 整倍数, 仅整数类型或数组
}

// 通过扩展信息给 method 增加约束
extend google.protobuf.ServiceOptions {
	ServiceRule service_rule = 10001;
}

// 通过扩展信息给 method 增加约束
extend google.protobuf.MethodOptions {
	MethodRule method_rule = 10001;
}

// 通过扩展信息给 message 增加约束
extend google.protobuf.MessageOptions {
	MethodInputRule method_input_rule = 10001;
}
```

通过扩展数据可以改变方法的行为, [user_data.proto](./user_data.proto) 指定了 POST 方法:

```proto
syntax = "proto3";

package service;

import "qingcloud_sdk_rule/rule.proto";

service UserDataService {
	option (qingcloud.sdk.rule.service_rule) = {
		service_name:     "QingCloud"
		sub_service_name: "UserData"
	};

	rpc UploadUserDataAttachment(UploadUserDataAttachmentInput) returns (UploadUserDataAttachmentOutput) {
		option (qingcloud.sdk.rule.method_rule) = {
			http_method: "POST"
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

message UploadUserDataAttachmentOutput {
	string action = 1;
	int32 ret_code = 2;
	string message = 3;

	string attachment_id = 4;
}
```

扩展信息中的 `http_action` 用于表示 HTTP 方法, 默认是 GET, 极少数的API是 POST ([UploadUserDataAttachment](https://docs.qingcloud.com/api/userdata/upload_userdata_attachment.html)). 当采用 POST 方法时,
需要明确指定 `http_action`.

输入的参数可以通过 `option (qingcloud.sdk.rule.method_input_rule)` 扩展来定义额外的约束, 主要是针对 必须成员/默认值/枚举字符串 几种类型.
不过目前还没有使用该信息.

<!--

## QingCloud API 状态

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

-->

