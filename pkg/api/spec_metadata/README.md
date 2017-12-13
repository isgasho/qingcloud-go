

# spec_metadata
`import "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package spec_metadata is a generated protocol buffer package.

It is generated from these files:


	spec_metadata/spec.proto

It has these top-level messages:


	FileSpec
	ServiceSpec
	ServiceMethodSpec
	FileOption
	ServiceOption
	MethodOption
	MessageOption
	FieldOption
	ExternalDocumentation
	CommandInfo




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type CommandInfo](#CommandInfo)
  * [func (*CommandInfo) Descriptor() ([]byte, []int)](#CommandInfo.Descriptor)
  * [func (m *CommandInfo) GetAliases() []string](#CommandInfo.GetAliases)
  * [func (m *CommandInfo) GetName() string](#CommandInfo.GetName)
  * [func (m *CommandInfo) GetUsage() string](#CommandInfo.GetUsage)
  * [func (*CommandInfo) ProtoMessage()](#CommandInfo.ProtoMessage)
  * [func (m *CommandInfo) Reset()](#CommandInfo.Reset)
  * [func (m *CommandInfo) String() string](#CommandInfo.String)
* [type ExternalDocumentation](#ExternalDocumentation)
  * [func (*ExternalDocumentation) Descriptor() ([]byte, []int)](#ExternalDocumentation.Descriptor)
  * [func (m *ExternalDocumentation) GetDescription() string](#ExternalDocumentation.GetDescription)
  * [func (m *ExternalDocumentation) GetTitle() string](#ExternalDocumentation.GetTitle)
  * [func (m *ExternalDocumentation) GetUrl() string](#ExternalDocumentation.GetUrl)
  * [func (*ExternalDocumentation) ProtoMessage()](#ExternalDocumentation.ProtoMessage)
  * [func (m *ExternalDocumentation) Reset()](#ExternalDocumentation.Reset)
  * [func (m *ExternalDocumentation) String() string](#ExternalDocumentation.String)
* [type FieldOption](#FieldOption)
  * [func (*FieldOption) Descriptor() ([]byte, []int)](#FieldOption.Descriptor)
  * [func (m *FieldOption) GetEnumValue() []string](#FieldOption.GetEnumValue)
  * [func (m *FieldOption) GetExclusiveMaxValue() float64](#FieldOption.GetExclusiveMaxValue)
  * [func (m *FieldOption) GetExclusiveMinValue() float64](#FieldOption.GetExclusiveMinValue)
  * [func (m *FieldOption) GetExternalDocs() *ExternalDocumentation](#FieldOption.GetExternalDocs)
  * [func (m *FieldOption) GetMaxLength() int32](#FieldOption.GetMaxLength)
  * [func (m *FieldOption) GetMaxValue() float64](#FieldOption.GetMaxValue)
  * [func (m *FieldOption) GetMinLength() int32](#FieldOption.GetMinLength)
  * [func (m *FieldOption) GetMinValue() float64](#FieldOption.GetMinValue)
  * [func (m *FieldOption) GetMultipleOfValue() float64](#FieldOption.GetMultipleOfValue)
  * [func (m *FieldOption) GetRegexpValue() string](#FieldOption.GetRegexpValue)
  * [func (m *FieldOption) GetStructTag() string](#FieldOption.GetStructTag)
  * [func (*FieldOption) ProtoMessage()](#FieldOption.ProtoMessage)
  * [func (m *FieldOption) Reset()](#FieldOption.Reset)
  * [func (m *FieldOption) String() string](#FieldOption.String)
* [type FileOption](#FileOption)
  * [func (*FileOption) Descriptor() ([]byte, []int)](#FileOption.Descriptor)
  * [func (m *FileOption) GetExternalDocs() *ExternalDocumentation](#FileOption.GetExternalDocs)
  * [func (*FileOption) ProtoMessage()](#FileOption.ProtoMessage)
  * [func (m *FileOption) Reset()](#FileOption.Reset)
  * [func (m *FileOption) String() string](#FileOption.String)
* [type FileSpec](#FileSpec)
  * [func (*FileSpec) Descriptor() ([]byte, []int)](#FileSpec.Descriptor)
  * [func (m *FileSpec) GetExternalDocs() *ExternalDocumentation](#FileSpec.GetExternalDocs)
  * [func (m *FileSpec) GetFileName() string](#FileSpec.GetFileName)
  * [func (m *FileSpec) GetPackageName() string](#FileSpec.GetPackageName)
  * [func (*FileSpec) ProtoMessage()](#FileSpec.ProtoMessage)
  * [func (m *FileSpec) Reset()](#FileSpec.Reset)
  * [func (m *FileSpec) String() string](#FileSpec.String)
* [type MessageOption](#MessageOption)
  * [func (*MessageOption) Descriptor() ([]byte, []int)](#MessageOption.Descriptor)
  * [func (m *MessageOption) GetExternalDocs() *ExternalDocumentation](#MessageOption.GetExternalDocs)
  * [func (*MessageOption) ProtoMessage()](#MessageOption.ProtoMessage)
  * [func (m *MessageOption) Reset()](#MessageOption.Reset)
  * [func (m *MessageOption) String() string](#MessageOption.String)
* [type MethodOption](#MethodOption)
  * [func (*MethodOption) Descriptor() ([]byte, []int)](#MethodOption.Descriptor)
  * [func (m *MethodOption) GetCmdInfo() *CommandInfo](#MethodOption.GetCmdInfo)
  * [func (m *MethodOption) GetExternalDocs() *ExternalDocumentation](#MethodOption.GetExternalDocs)
  * [func (m *MethodOption) GetHttpMethod() string](#MethodOption.GetHttpMethod)
  * [func (*MethodOption) ProtoMessage()](#MethodOption.ProtoMessage)
  * [func (m *MethodOption) Reset()](#MethodOption.Reset)
  * [func (m *MethodOption) String() string](#MethodOption.String)
* [type ServiceMethodSpec](#ServiceMethodSpec)
  * [func (*ServiceMethodSpec) Descriptor() ([]byte, []int)](#ServiceMethodSpec.Descriptor)
  * [func (m *ServiceMethodSpec) GetExternalDocs() *ExternalDocumentation](#ServiceMethodSpec.GetExternalDocs)
  * [func (m *ServiceMethodSpec) GetHttpMethod() string](#ServiceMethodSpec.GetHttpMethod)
  * [func (m *ServiceMethodSpec) GetInputTypeName() string](#ServiceMethodSpec.GetInputTypeName)
  * [func (m *ServiceMethodSpec) GetMethodName() string](#ServiceMethodSpec.GetMethodName)
  * [func (m *ServiceMethodSpec) GetOutputTypeName() string](#ServiceMethodSpec.GetOutputTypeName)
  * [func (*ServiceMethodSpec) ProtoMessage()](#ServiceMethodSpec.ProtoMessage)
  * [func (m *ServiceMethodSpec) Reset()](#ServiceMethodSpec.Reset)
  * [func (m *ServiceMethodSpec) String() string](#ServiceMethodSpec.String)
* [type ServiceOption](#ServiceOption)
  * [func (*ServiceOption) Descriptor() ([]byte, []int)](#ServiceOption.Descriptor)
  * [func (m *ServiceOption) GetCmdInfo() *CommandInfo](#ServiceOption.GetCmdInfo)
  * [func (m *ServiceOption) GetExternalDocs() *ExternalDocumentation](#ServiceOption.GetExternalDocs)
  * [func (*ServiceOption) ProtoMessage()](#ServiceOption.ProtoMessage)
  * [func (m *ServiceOption) Reset()](#ServiceOption.Reset)
  * [func (m *ServiceOption) String() string](#ServiceOption.String)
* [type ServiceSpec](#ServiceSpec)
  * [func (*ServiceSpec) Descriptor() ([]byte, []int)](#ServiceSpec.Descriptor)
  * [func (m *ServiceSpec) GetExternalDocs() *ExternalDocumentation](#ServiceSpec.GetExternalDocs)
  * [func (m *ServiceSpec) GetMethodList() []*ServiceMethodSpec](#ServiceSpec.GetMethodList)
  * [func (m *ServiceSpec) GetServiceName() string](#ServiceSpec.GetServiceName)
  * [func (*ServiceSpec) ProtoMessage()](#ServiceSpec.ProtoMessage)
  * [func (m *ServiceSpec) Reset()](#ServiceSpec.Reset)
  * [func (m *ServiceSpec) String() string](#ServiceSpec.String)


#### <a name="pkg-files">Package files</a>
[spec.pb.go](/src/github.com/chai2010/qingcloud-go/pkg/api/spec_metadata/spec.pb.go) 


## <a name="pkg-constants">Constants</a>
``` go
const Default_ServiceMethodSpec_HttpMethod string = "GET"
```

## <a name="pkg-variables">Variables</a>
``` go
var E_FieldOption = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*FieldOption)(nil),
    Field:         20171111,
    Name:          "qingcloud.api.spec_metadata.field_option",
    Tag:           "bytes,20171111,opt,name=field_option,json=fieldOption",
    Filename:      "spec_metadata/spec.proto",
}
```
``` go
var E_FileOption = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*FileOption)(nil),
    Field:         20171111,
    Name:          "qingcloud.api.spec_metadata.file_option",
    Tag:           "bytes,20171111,opt,name=file_option,json=fileOption",
    Filename:      "spec_metadata/spec.proto",
}
```
``` go
var E_MessageOption = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*MessageOption)(nil),
    Field:         20171111,
    Name:          "qingcloud.api.spec_metadata.message_option",
    Tag:           "bytes,20171111,opt,name=message_option,json=messageOption",
    Filename:      "spec_metadata/spec.proto",
}
```
``` go
var E_MethodOption = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MethodOptions)(nil),
    ExtensionType: (*MethodOption)(nil),
    Field:         20171111,
    Name:          "qingcloud.api.spec_metadata.method_option",
    Tag:           "bytes,20171111,opt,name=method_option,json=methodOption",
    Filename:      "spec_metadata/spec.proto",
}
```
``` go
var E_ServiceOption = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.ServiceOptions)(nil),
    ExtensionType: (*ServiceOption)(nil),
    Field:         20171111,
    Name:          "qingcloud.api.spec_metadata.service_option",
    Tag:           "bytes,20171111,opt,name=service_option,json=serviceOption",
    Filename:      "spec_metadata/spec.proto",
}
```



## <a name="CommandInfo">type</a> [CommandInfo](/src/target/spec.pb.go?s=13032:13357#L390)
``` go
type CommandInfo struct {
    Name             *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
    Aliases          []string `protobuf:"bytes,2,rep,name=aliases" json:"aliases,omitempty"`
    Usage            *string  `protobuf:"bytes,3,opt,name=usage" json:"usage,omitempty"`
    XXX_unrecognized []byte   `json:"-"`
}
```
命令行扩展










### <a name="CommandInfo.Descriptor">func</a> (\*CommandInfo) [Descriptor](/src/target/spec.pb.go?s=13570:13618#L400)
``` go
func (*CommandInfo) Descriptor() ([]byte, []int)
```



### <a name="CommandInfo.GetAliases">func</a> (\*CommandInfo) [GetAliases](/src/target/spec.pb.go?s=13764:13807#L409)
``` go
func (m *CommandInfo) GetAliases() []string
```



### <a name="CommandInfo.GetName">func</a> (\*CommandInfo) [GetName](/src/target/spec.pb.go?s=13657:13695#L402)
``` go
func (m *CommandInfo) GetName() string
```



### <a name="CommandInfo.GetUsage">func</a> (\*CommandInfo) [GetUsage](/src/target/spec.pb.go?s=13862:13901#L416)
``` go
func (m *CommandInfo) GetUsage() string
```



### <a name="CommandInfo.ProtoMessage">func</a> (\*CommandInfo) [ProtoMessage](/src/target/spec.pb.go?s=13518:13552#L399)
``` go
func (*CommandInfo) ProtoMessage()
```



### <a name="CommandInfo.Reset">func</a> (\*CommandInfo) [Reset](/src/target/spec.pb.go?s=13359:13388#L397)
``` go
func (m *CommandInfo) Reset()
```



### <a name="CommandInfo.String">func</a> (\*CommandInfo) [String](/src/target/spec.pb.go?s=13431:13468#L398)
``` go
func (m *CommandInfo) String() string
```



## <a name="ExternalDocumentation">type</a> [ExternalDocumentation](/src/target/spec.pb.go?s=11954:12291#L356)
``` go
type ExternalDocumentation struct {
    Title            *string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
    Description      *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
    Url              *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
    XXX_unrecognized []byte  `json:"-"`
}
```
扩展文档










### <a name="ExternalDocumentation.Descriptor">func</a> (\*ExternalDocumentation) [Descriptor](/src/target/spec.pb.go?s=12544:12602#L366)
``` go
func (*ExternalDocumentation) Descriptor() ([]byte, []int)
```



### <a name="ExternalDocumentation.GetDescription">func</a> (\*ExternalDocumentation) [GetDescription](/src/target/spec.pb.go?s=12761:12816#L375)
``` go
func (m *ExternalDocumentation) GetDescription() string
```



### <a name="ExternalDocumentation.GetTitle">func</a> (\*ExternalDocumentation) [GetTitle](/src/target/spec.pb.go?s=12641:12690#L368)
``` go
func (m *ExternalDocumentation) GetTitle() string
```



### <a name="ExternalDocumentation.GetUrl">func</a> (\*ExternalDocumentation) [GetUrl](/src/target/spec.pb.go?s=12899:12946#L382)
``` go
func (m *ExternalDocumentation) GetUrl() string
```



### <a name="ExternalDocumentation.ProtoMessage">func</a> (\*ExternalDocumentation) [ProtoMessage](/src/target/spec.pb.go?s=12482:12526#L365)
``` go
func (*ExternalDocumentation) ProtoMessage()
```



### <a name="ExternalDocumentation.Reset">func</a> (\*ExternalDocumentation) [Reset](/src/target/spec.pb.go?s=12293:12332#L363)
``` go
func (m *ExternalDocumentation) Reset()
```



### <a name="ExternalDocumentation.String">func</a> (\*ExternalDocumentation) [String](/src/target/spec.pb.go?s=12385:12432#L364)
``` go
func (m *ExternalDocumentation) String() string
```



## <a name="FieldOption">type</a> [FieldOption](/src/target/spec.pb.go?s=8693:10254#L258)
``` go
type FieldOption struct {
    MinValue          *float64               `protobuf:"fixed64,1,opt,name=min_value,json=minValue" json:"min_value,omitempty"`
    MaxValue          *float64               `protobuf:"fixed64,2,opt,name=max_value,json=maxValue" json:"max_value,omitempty"`
    MultipleOfValue   *float64               `protobuf:"fixed64,3,opt,name=multiple_of_value,json=multipleOfValue" json:"multiple_of_value,omitempty"`
    ExclusiveMinValue *float64               `protobuf:"fixed64,4,opt,name=exclusive_min_value,json=exclusiveMinValue" json:"exclusive_min_value,omitempty"`
    ExclusiveMaxValue *float64               `protobuf:"fixed64,5,opt,name=exclusive_max_value,json=exclusiveMaxValue" json:"exclusive_max_value,omitempty"`
    MinLength         *int32                 `protobuf:"varint,6,opt,name=min_length,json=minLength" json:"min_length,omitempty"`
    MaxLength         *int32                 `protobuf:"varint,7,opt,name=max_length,json=maxLength" json:"max_length,omitempty"`
    RegexpValue       *string                `protobuf:"bytes,8,opt,name=regexp_value,json=regexpValue" json:"regexp_value,omitempty"`
    EnumValue         []string               `protobuf:"bytes,9,rep,name=enum_value,json=enumValue" json:"enum_value,omitempty"`
    ExternalDocs      *ExternalDocumentation `protobuf:"bytes,10,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    StructTag         *string                `protobuf:"bytes,11,opt,name=struct_tag,json=structTag" json:"struct_tag,omitempty"`
    XXX_unrecognized  []byte                 `json:"-"`
}
```
消息成员的扩展信息
这是重要信息, 在运行时可动态获取改信息对 message 做合法性验证










### <a name="FieldOption.Descriptor">func</a> (\*FieldOption) [Descriptor](/src/target/spec.pb.go?s=10467:10515#L276)
``` go
func (*FieldOption) Descriptor() ([]byte, []int)
```



### <a name="FieldOption.GetEnumValue">func</a> (\*FieldOption) [GetEnumValue](/src/target/spec.pb.go?s=11592:11637#L334)
``` go
func (m *FieldOption) GetEnumValue() []string
```



### <a name="FieldOption.GetExclusiveMaxValue">func</a> (\*FieldOption) [GetExclusiveMaxValue](/src/target/spec.pb.go?s=11078:11130#L306)
``` go
func (m *FieldOption) GetExclusiveMaxValue() float64
```



### <a name="FieldOption.GetExclusiveMinValue">func</a> (\*FieldOption) [GetExclusiveMinValue](/src/target/spec.pb.go?s=10932:10984#L299)
``` go
func (m *FieldOption) GetExclusiveMinValue() float64
```



### <a name="FieldOption.GetExternalDocs">func</a> (\*FieldOption) [GetExternalDocs](/src/target/spec.pb.go?s=11694:11756#L341)
``` go
func (m *FieldOption) GetExternalDocs() *ExternalDocumentation
```



### <a name="FieldOption.GetMaxLength">func</a> (\*FieldOption) [GetMaxLength](/src/target/spec.pb.go?s=11344:11386#L320)
``` go
func (m *FieldOption) GetMaxLength() int32
```



### <a name="FieldOption.GetMaxValue">func</a> (\*FieldOption) [GetMaxValue](/src/target/spec.pb.go?s=10673:10716#L285)
``` go
func (m *FieldOption) GetMaxValue() float64
```



### <a name="FieldOption.GetMinLength">func</a> (\*FieldOption) [GetMinLength](/src/target/spec.pb.go?s=11224:11266#L313)
``` go
func (m *FieldOption) GetMinLength() int32
```



### <a name="FieldOption.GetMinValue">func</a> (\*FieldOption) [GetMinValue](/src/target/spec.pb.go?s=10554:10597#L278)
``` go
func (m *FieldOption) GetMinValue() float64
```



### <a name="FieldOption.GetMultipleOfValue">func</a> (\*FieldOption) [GetMultipleOfValue](/src/target/spec.pb.go?s=10792:10842#L292)
``` go
func (m *FieldOption) GetMultipleOfValue() float64
```



### <a name="FieldOption.GetRegexpValue">func</a> (\*FieldOption) [GetRegexpValue](/src/target/spec.pb.go?s=11464:11509#L327)
``` go
func (m *FieldOption) GetRegexpValue() string
```



### <a name="FieldOption.GetStructTag">func</a> (\*FieldOption) [GetStructTag](/src/target/spec.pb.go?s=11816:11859#L348)
``` go
func (m *FieldOption) GetStructTag() string
```



### <a name="FieldOption.ProtoMessage">func</a> (\*FieldOption) [ProtoMessage](/src/target/spec.pb.go?s=10415:10449#L275)
``` go
func (*FieldOption) ProtoMessage()
```



### <a name="FieldOption.Reset">func</a> (\*FieldOption) [Reset](/src/target/spec.pb.go?s=10256:10285#L273)
``` go
func (m *FieldOption) Reset()
```



### <a name="FieldOption.String">func</a> (\*FieldOption) [String](/src/target/spec.pb.go?s=10328:10365#L274)
``` go
func (m *FieldOption) String() string
```



## <a name="FileOption">type</a> [FileOption](/src/target/spec.pb.go?s=5224:5436#L161)
``` go
type FileOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
文件的扩展信息










### <a name="FileOption.Descriptor">func</a> (\*FileOption) [Descriptor](/src/target/spec.pb.go?s=5645:5692#L169)
``` go
func (*FileOption) Descriptor() ([]byte, []int)
```



### <a name="FileOption.GetExternalDocs">func</a> (\*FileOption) [GetExternalDocs](/src/target/spec.pb.go?s=5731:5792#L171)
``` go
func (m *FileOption) GetExternalDocs() *ExternalDocumentation
```



### <a name="FileOption.ProtoMessage">func</a> (\*FileOption) [ProtoMessage](/src/target/spec.pb.go?s=5594:5627#L168)
``` go
func (*FileOption) ProtoMessage()
```



### <a name="FileOption.Reset">func</a> (\*FileOption) [Reset](/src/target/spec.pb.go?s=5438:5466#L166)
``` go
func (m *FileOption) Reset()
```



### <a name="FileOption.String">func</a> (\*FileOption) [String](/src/target/spec.pb.go?s=5508:5544#L167)
``` go
func (m *FileOption) String() string
```



## <a name="FileSpec">type</a> [FileSpec](/src/target/spec.pb.go?s=1055:1518#L41)
``` go
type FileSpec struct {
    FileName         *string                `protobuf:"bytes,1,req,name=file_name,json=fileName" json:"file_name,omitempty"`
    PackageName      *string                `protobuf:"bytes,2,req,name=package_name,json=packageName" json:"package_name,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
proto 文件信息










### <a name="FileSpec.Descriptor">func</a> (\*FileSpec) [Descriptor](/src/target/spec.pb.go?s=1719:1764#L51)
``` go
func (*FileSpec) Descriptor() ([]byte, []int)
```



### <a name="FileSpec.GetExternalDocs">func</a> (\*FileSpec) [GetExternalDocs](/src/target/spec.pb.go?s=2044:2103#L67)
``` go
func (m *FileSpec) GetExternalDocs() *ExternalDocumentation
```



### <a name="FileSpec.GetFileName">func</a> (\*FileSpec) [GetFileName](/src/target/spec.pb.go?s=1803:1842#L53)
``` go
func (m *FileSpec) GetFileName() string
```



### <a name="FileSpec.GetPackageName">func</a> (\*FileSpec) [GetPackageName](/src/target/spec.pb.go?s=1919:1961#L60)
``` go
func (m *FileSpec) GetPackageName() string
```



### <a name="FileSpec.ProtoMessage">func</a> (\*FileSpec) [ProtoMessage](/src/target/spec.pb.go?s=1670:1701#L50)
``` go
func (*FileSpec) ProtoMessage()
```



### <a name="FileSpec.Reset">func</a> (\*FileSpec) [Reset](/src/target/spec.pb.go?s=1520:1546#L48)
``` go
func (m *FileSpec) Reset()
```



### <a name="FileSpec.String">func</a> (\*FileSpec) [String](/src/target/spec.pb.go?s=1586:1620#L49)
``` go
func (m *FileSpec) String() string
```



## <a name="MessageOption">type</a> [MessageOption](/src/target/spec.pb.go?s=7923:8138#L239)
``` go
type MessageOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
消息的扩展信息










### <a name="MessageOption.Descriptor">func</a> (\*MessageOption) [Descriptor](/src/target/spec.pb.go?s=8359:8409#L247)
``` go
func (*MessageOption) Descriptor() ([]byte, []int)
```



### <a name="MessageOption.GetExternalDocs">func</a> (\*MessageOption) [GetExternalDocs](/src/target/spec.pb.go?s=8448:8512#L249)
``` go
func (m *MessageOption) GetExternalDocs() *ExternalDocumentation
```



### <a name="MessageOption.ProtoMessage">func</a> (\*MessageOption) [ProtoMessage](/src/target/spec.pb.go?s=8305:8341#L246)
``` go
func (*MessageOption) ProtoMessage()
```



### <a name="MessageOption.Reset">func</a> (\*MessageOption) [Reset](/src/target/spec.pb.go?s=8140:8171#L244)
``` go
func (m *MessageOption) Reset()
```



### <a name="MessageOption.String">func</a> (\*MessageOption) [String](/src/target/spec.pb.go?s=8216:8255#L245)
``` go
func (m *MessageOption) String() string
```



## <a name="MethodOption">type</a> [MethodOption](/src/target/spec.pb.go?s=6780:7241#L205)
``` go
type MethodOption struct {
    HttpMethod       *string                `protobuf:"bytes,1,opt,name=http_method,json=httpMethod" json:"http_method,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,2,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    CmdInfo          *CommandInfo           `protobuf:"bytes,3,opt,name=cmd_info,json=cmdInfo" json:"cmd_info,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
服务方法的扩展信息










### <a name="MethodOption.Descriptor">func</a> (\*MethodOption) [Descriptor](/src/target/spec.pb.go?s=7458:7507#L215)
``` go
func (*MethodOption) Descriptor() ([]byte, []int)
```



### <a name="MethodOption.GetCmdInfo">func</a> (\*MethodOption) [GetCmdInfo](/src/target/spec.pb.go?s=7795:7843#L231)
``` go
func (m *MethodOption) GetCmdInfo() *CommandInfo
```



### <a name="MethodOption.GetExternalDocs">func</a> (\*MethodOption) [GetExternalDocs](/src/target/spec.pb.go?s=7672:7735#L224)
``` go
func (m *MethodOption) GetExternalDocs() *ExternalDocumentation
```



### <a name="MethodOption.GetHttpMethod">func</a> (\*MethodOption) [GetHttpMethod](/src/target/spec.pb.go?s=7546:7591#L217)
``` go
func (m *MethodOption) GetHttpMethod() string
```



### <a name="MethodOption.ProtoMessage">func</a> (\*MethodOption) [ProtoMessage](/src/target/spec.pb.go?s=7405:7440#L214)
``` go
func (*MethodOption) ProtoMessage()
```



### <a name="MethodOption.Reset">func</a> (\*MethodOption) [Reset](/src/target/spec.pb.go?s=7243:7273#L212)
``` go
func (m *MethodOption) Reset()
```



### <a name="MethodOption.String">func</a> (\*MethodOption) [String](/src/target/spec.pb.go?s=7317:7355#L213)
``` go
func (m *MethodOption) String() string
```



## <a name="ServiceMethodSpec">type</a> [ServiceMethodSpec](/src/target/spec.pb.go?s=3339:4103#L109)
``` go
type ServiceMethodSpec struct {
    MethodName       *string                `protobuf:"bytes,1,req,name=method_name,json=methodName" json:"method_name,omitempty"`
    InputTypeName    *string                `protobuf:"bytes,2,req,name=input_type_name,json=inputTypeName" json:"input_type_name,omitempty"`
    OutputTypeName   *string                `protobuf:"bytes,3,req,name=output_type_name,json=outputTypeName" json:"output_type_name,omitempty"`
    HttpMethod       *string                `protobuf:"bytes,4,opt,name=http_method,json=httpMethod,def=GET" json:"http_method,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,5,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
服务方法列表










### <a name="ServiceMethodSpec.Descriptor">func</a> (\*ServiceMethodSpec) [Descriptor](/src/target/spec.pb.go?s=4340:4394#L121)
``` go
func (*ServiceMethodSpec) Descriptor() ([]byte, []int)
```



### <a name="ServiceMethodSpec.GetExternalDocs">func</a> (\*ServiceMethodSpec) [GetExternalDocs](/src/target/spec.pb.go?s=5071:5139#L153)
``` go
func (m *ServiceMethodSpec) GetExternalDocs() *ExternalDocumentation
```



### <a name="ServiceMethodSpec.GetHttpMethod">func</a> (\*ServiceMethodSpec) [GetHttpMethod](/src/target/spec.pb.go?s=4906:4956#L146)
``` go
func (m *ServiceMethodSpec) GetHttpMethod() string
```



### <a name="ServiceMethodSpec.GetInputTypeName">func</a> (\*ServiceMethodSpec) [GetInputTypeName](/src/target/spec.pb.go?s=4623:4676#L132)
``` go
func (m *ServiceMethodSpec) GetInputTypeName() string
```



### <a name="ServiceMethodSpec.GetMethodName">func</a> (\*ServiceMethodSpec) [GetMethodName](/src/target/spec.pb.go?s=4492:4542#L125)
``` go
func (m *ServiceMethodSpec) GetMethodName() string
```



### <a name="ServiceMethodSpec.GetOutputTypeName">func</a> (\*ServiceMethodSpec) [GetOutputTypeName](/src/target/spec.pb.go?s=4763:4817#L139)
``` go
func (m *ServiceMethodSpec) GetOutputTypeName() string
```



### <a name="ServiceMethodSpec.ProtoMessage">func</a> (\*ServiceMethodSpec) [ProtoMessage](/src/target/spec.pb.go?s=4282:4322#L120)
``` go
func (*ServiceMethodSpec) ProtoMessage()
```



### <a name="ServiceMethodSpec.Reset">func</a> (\*ServiceMethodSpec) [Reset](/src/target/spec.pb.go?s=4105:4140#L118)
``` go
func (m *ServiceMethodSpec) Reset()
```



### <a name="ServiceMethodSpec.String">func</a> (\*ServiceMethodSpec) [String](/src/target/spec.pb.go?s=4189:4232#L119)
``` go
func (m *ServiceMethodSpec) String() string
```



## <a name="ServiceOption">type</a> [ServiceOption](/src/target/spec.pb.go?s=5877:6211#L179)
``` go
type ServiceOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    CmdInfo          *CommandInfo           `protobuf:"bytes,2,opt,name=cmd_info,json=cmdInfo" json:"cmd_info,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
服务的扩展信息










### <a name="ServiceOption.Descriptor">func</a> (\*ServiceOption) [Descriptor](/src/target/spec.pb.go?s=6432:6482#L188)
``` go
func (*ServiceOption) Descriptor() ([]byte, []int)
```



### <a name="ServiceOption.GetCmdInfo">func</a> (\*ServiceOption) [GetCmdInfo](/src/target/spec.pb.go?s=6645:6694#L197)
``` go
func (m *ServiceOption) GetCmdInfo() *CommandInfo
```



### <a name="ServiceOption.GetExternalDocs">func</a> (\*ServiceOption) [GetExternalDocs](/src/target/spec.pb.go?s=6521:6585#L190)
``` go
func (m *ServiceOption) GetExternalDocs() *ExternalDocumentation
```



### <a name="ServiceOption.ProtoMessage">func</a> (\*ServiceOption) [ProtoMessage](/src/target/spec.pb.go?s=6378:6414#L187)
``` go
func (*ServiceOption) ProtoMessage()
```



### <a name="ServiceOption.Reset">func</a> (\*ServiceOption) [Reset](/src/target/spec.pb.go?s=6213:6244#L185)
``` go
func (m *ServiceOption) Reset()
```



### <a name="ServiceOption.String">func</a> (\*ServiceOption) [String](/src/target/spec.pb.go?s=6289:6328#L186)
``` go
func (m *ServiceOption) String() string
```



## <a name="ServiceSpec">type</a> [ServiceSpec](/src/target/spec.pb.go?s=2179:2651#L75)
``` go
type ServiceSpec struct {
    ServiceName      *string                `protobuf:"bytes,1,req,name=service_name,json=serviceName" json:"service_name,omitempty"`
    MethodList       []*ServiceMethodSpec   `protobuf:"bytes,2,rep,name=method_list,json=methodList" json:"method_list,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```
服务信息










### <a name="ServiceSpec.Descriptor">func</a> (\*ServiceSpec) [Descriptor](/src/target/spec.pb.go?s=2864:2912#L85)
``` go
func (*ServiceSpec) Descriptor() ([]byte, []int)
```



### <a name="ServiceSpec.GetExternalDocs">func</a> (\*ServiceSpec) [GetExternalDocs](/src/target/spec.pb.go?s=3195:3257#L101)
``` go
func (m *ServiceSpec) GetExternalDocs() *ExternalDocumentation
```



### <a name="ServiceSpec.GetMethodList">func</a> (\*ServiceSpec) [GetMethodList](/src/target/spec.pb.go?s=3079:3137#L94)
``` go
func (m *ServiceSpec) GetMethodList() []*ServiceMethodSpec
```



### <a name="ServiceSpec.GetServiceName">func</a> (\*ServiceSpec) [GetServiceName](/src/target/spec.pb.go?s=2951:2996#L87)
``` go
func (m *ServiceSpec) GetServiceName() string
```



### <a name="ServiceSpec.ProtoMessage">func</a> (\*ServiceSpec) [ProtoMessage](/src/target/spec.pb.go?s=2812:2846#L84)
``` go
func (*ServiceSpec) ProtoMessage()
```



### <a name="ServiceSpec.Reset">func</a> (\*ServiceSpec) [Reset](/src/target/spec.pb.go?s=2653:2682#L82)
``` go
func (m *ServiceSpec) Reset()
```



### <a name="ServiceSpec.String">func</a> (\*ServiceSpec) [String](/src/target/spec.pb.go?s=2725:2762#L83)
``` go
func (m *ServiceSpec) String() string
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
