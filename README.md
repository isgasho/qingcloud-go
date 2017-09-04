# qingcloud 自用 SDK

基本思路:

- 基于 Protobuf3 语法对应的 json 格式文件定义 api 元数据
- 基于 Protobuf3 语法定义 rest 接口的请求和响应结构体, 只是用于结构体, 请求时转 json 处理(proto库自带)
- 基于 [Mustache](http://mustache.github.io) 模板定义最终要渲染的生成代码

最终目标:

- api 元数据采用语言无关的 Proto3 语法定义, 文件采用 proto3 对应的 json 格式
- 生成的请求和响应结构体保持和 Proto3 生成的结构一致, Proto3 本身支持多种语言
- Go 标准库的模板语法太难用, [Mustache](http://mustache.github.io) 语言无关
- 更多的单元测试代码
- 没有外部依赖!

基于现有的 proto3 和 Mustache 构造语言无关的格式.
