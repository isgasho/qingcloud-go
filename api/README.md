## 设计思路

- 基于 Protobuf2 语法定义青云服务接口规范
- 基于 Protobuf2 的扩展特性, 增加自定义的元数据
- 基于 Protobuf2 生成的输入和输出参数结构进行 json 的编码和解码
- 内置基于Go标准库模板语法的Go代码生成
- 外部模板暂不支持(以后会考虑)

## 构建流程

基于标准的 protobuf 插件生成代码:

1. 安装官方的 [`protoc`](https://github.com/google/protobuf/releases) 程序, V3版本, 带了官方的扩展类型
1. 安装官方的 `protoc-gen-go` 插件, `go install github.com/golang/protobuf/protoc-gen-go`
1. 用官方的 `protoc-gen-go` 插件编译当前目录下全部的 proto 文件(生成的文件后缀名为`*.pb.go`, 不含青云SDK代码)

用针对青云定制的插件生成青云SDK相关代码:

1. 安装定制的`protoc-gen-qingcloud`插件, 命令: `go install github.com/chai2010/qingcloud-go/cmd/protoc-gen-qingcloud`
1. 用定制的`protoc-gen-qingcloud`插件编译 spec.pb 目录下全部的 proto 文件(默认是Go语言, 生成的文件后缀名为`*.pb.qingcloud.go`, 含有青云SDK代码), 参考 `make` 命令
1. 在上级目录运行单元测试 `make test ./...`
1. OK


**服务升级流程:**

1. 更新已有的 proto 接口文件, 或者新建 proto 接口文件
1. 运行 `make`

## Protobuf 扩展信息

扩展类型在 [api/spec_metadata/spec.proto](./spec_metadata/spec.proto) 文件中定义:

```proto
// 文件的扩展信息
message FileOption {
	optional ExternalDocumentation external_docs = 1;
}

// 服务的扩展信息
message ServiceOption {
	optional ExternalDocumentation external_docs = 3;
}

// 服务方法的扩展信息
message MethodOption {
	optional string http_method = 1;
	optional ExternalDocumentation external_docs = 2;
}

// 消息的扩展信息
message MessageOption {
	optional ExternalDocumentation external_docs = 1;
}

// 消息成员的扩展信息
// 这是重要信息, 在运行时可动态获取改信息对 message 做合法性验证
message FieldOption {
	optional double min_value = 1;
	optional double max_value = 2;
	optional double multiple_of_value = 3;
	optional double exclusive_min_value = 4;
	optional double exclusive_max_value = 5;
	optional int32 min_length = 6;
	optional int32 max_length = 7;
	optional string regexp_value = 8;
	repeated string enum_value = 9;
	optional ExternalDocumentation external_docs = 10;
	optional string struct_tag = 11;
}

// 扩展文档
message ExternalDocumentation {
	optional string title = 1;
	optional string description = 2;
	optional string url = 3;
}
```

通过扩展数据可以改变方法的行为, [user_data.proto](./user_data.proto) 指定了 POST 方法:

```proto
syntax = "proto2";

package service;

import "spec_metadata/spec.proto";

// https://docs.qingcloud.com/api/userdata/index.html

message UserDataServiceProperties {
	optional string zone = 1;
}

service UserDataService {
	option (qingcloud.api.spec_metadata.service_option) = {
		external_docs: {
			url: "https://docs.qingcloud.com/api/userdata/index.html"
		}
	};

	rpc UploadUserDataAttachment(UploadUserDataAttachmentInput) returns (UploadUserDataAttachmentOutput) {
		option (qingcloud.api.spec_metadata.method_option) = {
			http_method: "POST"
		};
	}
}

message UploadUserDataAttachmentInput {
	optional string attachment_name = 1;
	required bytes attachment_content = 2;
}

message UploadUserDataAttachmentOutput {
	optional string action = 1;
	optional int32 ret_code = 2;
	optional string message = 3;

	optional string attachment_id = 4;
}
```

扩展信息中的 `http_method` 用于表示 HTTP 方法, 默认是 GET, 极少数的API是 POST ([UploadUserDataAttachment](https://docs.qingcloud.com/api/userdata/upload_userdata_attachment.html)). 当采用 POST 方法时,
需要明确指定 `http_method`.

输入的参数可以通过 `(qingcloud.api.spec_metadata.field_option)` 扩展来定义额外的约束, 用于验证整个结构体.

目前生成的代码还没有使用该信息.

