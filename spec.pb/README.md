## 设计思路

- 基于 Protobuf3 语法定义青云服务接口规范
- 基于 Protobuf3 的扩展特性, 增加自定义的元数据
- 基于 Protobuf3 生成的输入和输出参数结构进行 json 的编码和解码
- 内置基于Go标准库模板语法的Go代码生成
- 外部模板暂不支持(以后会考虑)

## 构建流程

基于标准的 protobuf 插件生成代码:

1. 安装官方的 [`protoc`](https://github.com/google/protobuf/releases) 程序, V3版本, 带了官方的扩展类型
1. 安装官方的 `protoc-gen-go` 插件, `go install github.com/golang/protobuf/protoc-gen-go`
1. 用官方的 `protoc-gen-go` 插件编译 spec.pb 目录下全部的 proto 文件(生成的文件后缀名为`*.pb.go`, 不含青云SDK代码)

用针对青云定制的插件生成青云SDK相关代码:

1. 安装定制的`protoc-gen-qingcloud-go`插件, 命令: `go install github.com/chai2010/qingcloud-go/protoc-gen-qingcloud-go`
1. 用定制的`protoc-gen-qingcloud-go`插件编译 spec.pb 目录下全部的 proto 文件(生成的文件后缀名为`*.pb.qingcloud.go`, 含有青云SDK代码), 参考 `make` 命令
1. 在上级目录运行单元测试 `make test ./...`
1. OK


**服务升级流程:**

1. 更新已有的 proto 接口文件, 或者新建 proto 接口文件
1. 运行 `make`

## Protobuf 扩展信息

扩展类型在 [spec.pb/qingcloud_sdk_rule/rule.proto](./qingcloud_sdk_rule/rule.proto) 文件中定义:

```proto
// 服务规则
// 有主服务和子服务之分, 子服务隶属于某个主服务
message ServiceOptionsRule {
	string doc_url = 1;          // 文档链接
	string service_name = 2;     // 服务名, 格式: QingCloud, QingCloud.Alarm
}

// 方法规则
message MethodOptionsRule {
	string doc_url = 1;          // 文档链接
	string http_method = 2;      // http 行为有 GET 和 POST 之分, 默认是 GET
	string input_type = 3;       // 输入参数类型
	string output_type = 4;      // 输出参数类型
}

// 输入参数规则
// 输入参数成员只有数值类型和字符串, 以及对应的数组, 不含嵌套结构
// 元信息部分只能包含 字符串/数字/链接符号 等普通的字符
message MessageOptionsRule {
	string required_fileds = 1;   // 格式: "a; b; ..."
	string default_value = 2;     // 格式: "a:v; b:v; ..."
	string enum_value = 3;        // 格式: "a:a1,a2,a3; b:b1,b2; ..."
	string min_value = 4;         // 格式: "a:v; b:v ...", 最小值, 仅数值类型或数组
	string max_value = 5;         // 格式: "a:v; b:v; ...", 最大值, 仅数值类型或数组
	string multiple_of_value = 6; // 格式: "a:v; b:v; ...", 整倍数, 仅整数类型或数组
	string regexp_value = 7;      // 格式: "a:{{...}}; b:{{...}};", 简单正则, 不要挑战复杂格式
}

// 通过扩展信息给 method 增加约束
extend google.protobuf.ServiceOptions {
	ServiceOptionsRule service_rule = 10001;
}

// 通过扩展信息给 method 增加约束
extend google.protobuf.MethodOptions {
	MethodOptionsRule method_rule = 10001;
}

// 通过扩展信息给 message 增加约束
extend google.protobuf.MessageOptions {
	MessageOptionsRule message_rule = 10001;
}
```

通过扩展数据可以改变方法的行为, [user_data.proto](./user_data.proto) 指定了 POST 方法:

```proto
syntax = "proto3";

package service;

import "qingcloud_sdk_rule/rule.proto";

// https://docs.qingcloud.com/api/userdata/index.html

message UserDataServiceProperties {
	string zone = 1;
}

service UserDataService {
	option (qingcloud.sdk.rule.service_rule) = {
		service_name: "QingCloud.UserData"
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

	option (qingcloud.sdk.rule.message_rule) = {
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

扩展信息中的 `http_method` 用于表示 HTTP 方法, 默认是 GET, 极少数的API是 POST ([UploadUserDataAttachment](https://docs.qingcloud.com/api/userdata/upload_userdata_attachment.html)). 当采用 POST 方法时,
需要明确指定 `http_method`.

输入的参数可以通过 `option (qingcloud.sdk.rule.message_rule)` 扩展来定义额外的约束, 用于验证整个结构体.
目前生成的代码还没有使用该信息.
