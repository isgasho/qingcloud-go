# 更新日志

### next

TODO


### v1.0.0

- 增加 GopherLua 接口
- qcli 增加 lake 子命令行, 内置青云 Lua 接口, 支持类似 makefile 方式管理任务
- 增加 qlua 命令,  内置青云 Lua 接口, 用于执行普通的 lua 脚本
- 增加 qlake 脚本(需要从 qlua 启动), 执行独立的 lakefile 任务

TODO:

- 接口文件审核

### v0.9.0

- 项目彻底重构, 完全去掉对官方 SDK 的代码引用, 现在的代码是 100% 独立实现的
- 放弃对官方 SDK 接口的兼容性, 致力于简化接口和简化实现
- 基本实现外部依赖最小化: SDK 仅仅依赖 Protobuf 包
- 增加 qcli 命令行, 实现和 SDK 函数的逐一对应

待完善部分:

- 缺少 `Validate` 代码, 稍后会加入
- qcli 的命令行参数最好能和 query 参数统一

### v0.2.0

- 重构插件代码
- 临时移除生成 `Validate` 代码, 稍后会重新加入

### v0.1.0

- 实现 `protoMessage.Validate` 方法, 可实现用户输入参数验证

### v0.0.9

- spec: 基本完成全部接口规范(少数proto无法支持的接口除外)

### v0.0.4

- spec: 增加对服务的扩展类型
- 生成代码增加返回原始 response 的途径
- cluster 接口规范初稿

### v0.0.3

- 最大程度兼容 yunify/qingcloud-sdk-go (pb生成成员的分词会有些差异)
- 实现了 message 的扩展, 可用于生成输入参数的验证函数

### v0.0.2

- 支持 POST 和 GET 标记
- 3 个例子可运行

### v0.0.1

- 第一个例子可以正常运行
