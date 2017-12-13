# spec_metadata
`import "github.com/chai2010/qingcloud-go/pkg/api/spec_metadata"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
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

**Package spec_metadata is a generated [Protobuf](https://developers.google.com/protocol-buffers/)-compatible package.**

It is generated from these files:

- [spec_metadata/spec.proto](./spec_metadata/spec.proto)

## <a name="pkg-imports">Imported Packages</a>

- [github.com/golang/protobuf/proto](https://godoc.org/github.com/golang/protobuf/proto)
- [github.com/golang/protobuf/protoc-gen-go/descriptor](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/descriptor)

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type CommandInfo](#CommandInfo)
  * [func (\*CommandInfo) Descriptor() ([]byte, []int)](#CommandInfo.Descriptor)
  * [func (m \*CommandInfo) GetAliases() []string](#CommandInfo.GetAliases)
  * [func (m \*CommandInfo) GetName() string](#CommandInfo.GetName)
  * [func (m \*CommandInfo) GetUsage() string](#CommandInfo.GetUsage)
  * [func (\*CommandInfo) ProtoMessage()](#CommandInfo.ProtoMessage)
  * [func (m \*CommandInfo) Reset()](#CommandInfo.Reset)
  * [func (m \*CommandInfo) String() string](#CommandInfo.String)
* [type ExternalDocumentation](#ExternalDocumentation)
  * [func (\*ExternalDocumentation) Descriptor() ([]byte, []int)](#ExternalDocumentation.Descriptor)
  * [func (m \*ExternalDocumentation) GetDescription() string](#ExternalDocumentation.GetDescription)
  * [func (m \*ExternalDocumentation) GetTitle() string](#ExternalDocumentation.GetTitle)
  * [func (m \*ExternalDocumentation) GetUrl() string](#ExternalDocumentation.GetUrl)
  * [func (\*ExternalDocumentation) ProtoMessage()](#ExternalDocumentation.ProtoMessage)
  * [func (m \*ExternalDocumentation) Reset()](#ExternalDocumentation.Reset)
  * [func (m \*ExternalDocumentation) String() string](#ExternalDocumentation.String)
* [type FieldOption](#FieldOption)
  * [func (\*FieldOption) Descriptor() ([]byte, []int)](#FieldOption.Descriptor)
  * [func (m \*FieldOption) GetEnumValue() []string](#FieldOption.GetEnumValue)
  * [func (m \*FieldOption) GetExclusiveMaxValue() float64](#FieldOption.GetExclusiveMaxValue)
  * [func (m \*FieldOption) GetExclusiveMinValue() float64](#FieldOption.GetExclusiveMinValue)
  * [func (m \*FieldOption) GetExternalDocs() \*ExternalDocumentation](#FieldOption.GetExternalDocs)
  * [func (m \*FieldOption) GetMaxLength() int32](#FieldOption.GetMaxLength)
  * [func (m \*FieldOption) GetMaxValue() float64](#FieldOption.GetMaxValue)
  * [func (m \*FieldOption) GetMinLength() int32](#FieldOption.GetMinLength)
  * [func (m \*FieldOption) GetMinValue() float64](#FieldOption.GetMinValue)
  * [func (m \*FieldOption) GetMultipleOfValue() float64](#FieldOption.GetMultipleOfValue)
  * [func (m \*FieldOption) GetRegexpValue() string](#FieldOption.GetRegexpValue)
  * [func (m \*FieldOption) GetStructTag() string](#FieldOption.GetStructTag)
  * [func (\*FieldOption) ProtoMessage()](#FieldOption.ProtoMessage)
  * [func (m \*FieldOption) Reset()](#FieldOption.Reset)
  * [func (m \*FieldOption) String() string](#FieldOption.String)
* [type FileOption](#FileOption)
  * [func (\*FileOption) Descriptor() ([]byte, []int)](#FileOption.Descriptor)
  * [func (m \*FileOption) GetExternalDocs() \*ExternalDocumentation](#FileOption.GetExternalDocs)
  * [func (\*FileOption) ProtoMessage()](#FileOption.ProtoMessage)
  * [func (m \*FileOption) Reset()](#FileOption.Reset)
  * [func (m \*FileOption) String() string](#FileOption.String)
* [type FileSpec](#FileSpec)
  * [func (\*FileSpec) Descriptor() ([]byte, []int)](#FileSpec.Descriptor)
  * [func (m \*FileSpec) GetExternalDocs() \*ExternalDocumentation](#FileSpec.GetExternalDocs)
  * [func (m \*FileSpec) GetFileName() string](#FileSpec.GetFileName)
  * [func (m \*FileSpec) GetPackageName() string](#FileSpec.GetPackageName)
  * [func (\*FileSpec) ProtoMessage()](#FileSpec.ProtoMessage)
  * [func (m \*FileSpec) Reset()](#FileSpec.Reset)
  * [func (m \*FileSpec) String() string](#FileSpec.String)
* [type MessageOption](#MessageOption)
  * [func (\*MessageOption) Descriptor() ([]byte, []int)](#MessageOption.Descriptor)
  * [func (m \*MessageOption) GetExternalDocs() \*ExternalDocumentation](#MessageOption.GetExternalDocs)
  * [func (\*MessageOption) ProtoMessage()](#MessageOption.ProtoMessage)
  * [func (m \*MessageOption) Reset()](#MessageOption.Reset)
  * [func (m \*MessageOption) String() string](#MessageOption.String)
* [type MethodOption](#MethodOption)
  * [func (\*MethodOption) Descriptor() ([]byte, []int)](#MethodOption.Descriptor)
  * [func (m \*MethodOption) GetCmdInfo() \*CommandInfo](#MethodOption.GetCmdInfo)
  * [func (m \*MethodOption) GetExternalDocs() \*ExternalDocumentation](#MethodOption.GetExternalDocs)
  * [func (m \*MethodOption) GetHttpMethod() string](#MethodOption.GetHttpMethod)
  * [func (\*MethodOption) ProtoMessage()](#MethodOption.ProtoMessage)
  * [func (m \*MethodOption) Reset()](#MethodOption.Reset)
  * [func (m \*MethodOption) String() string](#MethodOption.String)
* [type ServiceMethodSpec](#ServiceMethodSpec)
  * [func (\*ServiceMethodSpec) Descriptor() ([]byte, []int)](#ServiceMethodSpec.Descriptor)
  * [func (m \*ServiceMethodSpec) GetExternalDocs() \*ExternalDocumentation](#ServiceMethodSpec.GetExternalDocs)
  * [func (m \*ServiceMethodSpec) GetHttpMethod() string](#ServiceMethodSpec.GetHttpMethod)
  * [func (m \*ServiceMethodSpec) GetInputTypeName() string](#ServiceMethodSpec.GetInputTypeName)
  * [func (m \*ServiceMethodSpec) GetMethodName() string](#ServiceMethodSpec.GetMethodName)
  * [func (m \*ServiceMethodSpec) GetOutputTypeName() string](#ServiceMethodSpec.GetOutputTypeName)
  * [func (\*ServiceMethodSpec) ProtoMessage()](#ServiceMethodSpec.ProtoMessage)
  * [func (m \*ServiceMethodSpec) Reset()](#ServiceMethodSpec.Reset)
  * [func (m \*ServiceMethodSpec) String() string](#ServiceMethodSpec.String)
* [type ServiceOption](#ServiceOption)
  * [func (\*ServiceOption) Descriptor() ([]byte, []int)](#ServiceOption.Descriptor)
  * [func (m \*ServiceOption) GetCmdInfo() \*CommandInfo](#ServiceOption.GetCmdInfo)
  * [func (m \*ServiceOption) GetExternalDocs() \*ExternalDocumentation](#ServiceOption.GetExternalDocs)
  * [func (\*ServiceOption) ProtoMessage()](#ServiceOption.ProtoMessage)
  * [func (m \*ServiceOption) Reset()](#ServiceOption.Reset)
  * [func (m \*ServiceOption) String() string](#ServiceOption.String)
* [type ServiceSpec](#ServiceSpec)
  * [func (\*ServiceSpec) Descriptor() ([]byte, []int)](#ServiceSpec.Descriptor)
  * [func (m \*ServiceSpec) GetExternalDocs() \*ExternalDocumentation](#ServiceSpec.GetExternalDocs)
  * [func (m \*ServiceSpec) GetMethodList() []\*ServiceMethodSpec](#ServiceSpec.GetMethodList)
  * [func (m \*ServiceSpec) GetServiceName() string](#ServiceSpec.GetServiceName)
  * [func (\*ServiceSpec) ProtoMessage()](#ServiceSpec.ProtoMessage)
  * [func (m \*ServiceSpec) Reset()](#ServiceSpec.Reset)
  * [func (m \*ServiceSpec) String() string](#ServiceSpec.String)

#### <a name="pkg-files">Package files</a>
[spec.pb.go](./spec.pb.go) 

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

## <a name="CommandInfo">type</a> [CommandInfo](./spec.pb.go#L379-L384)
``` go
type CommandInfo struct {
    Name             *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
    Aliases          []string `protobuf:"bytes,2,rep,name=aliases" json:"aliases,omitempty"`
    Usage            *string  `protobuf:"bytes,3,opt,name=usage" json:"usage,omitempty"`
    XXX_unrecognized []byte   `json:"-"`
}
```

### <a name="CommandInfo.Descriptor">func</a> (\*CommandInfo) [Descriptor](./spec.pb.go#L389)
``` go
func (*CommandInfo) Descriptor() ([]byte, []int)
```

### <a name="CommandInfo.GetAliases">func</a> (\*CommandInfo) [GetAliases](./spec.pb.go#L398)
``` go
func (m *CommandInfo) GetAliases() []string
```

### <a name="CommandInfo.GetName">func</a> (\*CommandInfo) [GetName](./spec.pb.go#L391)
``` go
func (m *CommandInfo) GetName() string
```

### <a name="CommandInfo.GetUsage">func</a> (\*CommandInfo) [GetUsage](./spec.pb.go#L405)
``` go
func (m *CommandInfo) GetUsage() string
```

### <a name="CommandInfo.ProtoMessage">func</a> (\*CommandInfo) [ProtoMessage](./spec.pb.go#L388)
``` go
func (*CommandInfo) ProtoMessage()
```

### <a name="CommandInfo.Reset">func</a> (\*CommandInfo) [Reset](./spec.pb.go#L386)
``` go
func (m *CommandInfo) Reset()
```

### <a name="CommandInfo.String">func</a> (\*CommandInfo) [String](./spec.pb.go#L387)
``` go
func (m *CommandInfo) String() string
```

## <a name="ExternalDocumentation">type</a> [ExternalDocumentation](./spec.pb.go#L346-L351)
``` go
type ExternalDocumentation struct {
    Title            *string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
    Description      *string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
    Url              *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
    XXX_unrecognized []byte  `json:"-"`
}
```

### <a name="ExternalDocumentation.Descriptor">func</a> (\*ExternalDocumentation) [Descriptor](./spec.pb.go#L356)
``` go
func (*ExternalDocumentation) Descriptor() ([]byte, []int)
```

### <a name="ExternalDocumentation.GetDescription">func</a> (\*ExternalDocumentation) [GetDescription](./spec.pb.go#L365)
``` go
func (m *ExternalDocumentation) GetDescription() string
```

### <a name="ExternalDocumentation.GetTitle">func</a> (\*ExternalDocumentation) [GetTitle](./spec.pb.go#L358)
``` go
func (m *ExternalDocumentation) GetTitle() string
```

### <a name="ExternalDocumentation.GetUrl">func</a> (\*ExternalDocumentation) [GetUrl](./spec.pb.go#L372)
``` go
func (m *ExternalDocumentation) GetUrl() string
```

### <a name="ExternalDocumentation.ProtoMessage">func</a> (\*ExternalDocumentation) [ProtoMessage](./spec.pb.go#L355)
``` go
func (*ExternalDocumentation) ProtoMessage()
```

### <a name="ExternalDocumentation.Reset">func</a> (\*ExternalDocumentation) [Reset](./spec.pb.go#L353)
``` go
func (m *ExternalDocumentation) Reset()
```

### <a name="ExternalDocumentation.String">func</a> (\*ExternalDocumentation) [String](./spec.pb.go#L354)
``` go
func (m *ExternalDocumentation) String() string
```

## <a name="FieldOption">type</a> [FieldOption](./spec.pb.go#L249-L262)
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

### <a name="FieldOption.Descriptor">func</a> (\*FieldOption) [Descriptor](./spec.pb.go#L267)
``` go
func (*FieldOption) Descriptor() ([]byte, []int)
```

### <a name="FieldOption.GetEnumValue">func</a> (\*FieldOption) [GetEnumValue](./spec.pb.go#L325)
``` go
func (m *FieldOption) GetEnumValue() []string
```

### <a name="FieldOption.GetExclusiveMaxValue">func</a> (\*FieldOption) [GetExclusiveMaxValue](./spec.pb.go#L297)
``` go
func (m *FieldOption) GetExclusiveMaxValue() float64
```

### <a name="FieldOption.GetExclusiveMinValue">func</a> (\*FieldOption) [GetExclusiveMinValue](./spec.pb.go#L290)
``` go
func (m *FieldOption) GetExclusiveMinValue() float64
```

### <a name="FieldOption.GetExternalDocs">func</a> (\*FieldOption) [GetExternalDocs](./spec.pb.go#L332)
``` go
func (m *FieldOption) GetExternalDocs() *ExternalDocumentation
```

### <a name="FieldOption.GetMaxLength">func</a> (\*FieldOption) [GetMaxLength](./spec.pb.go#L311)
``` go
func (m *FieldOption) GetMaxLength() int32
```

### <a name="FieldOption.GetMaxValue">func</a> (\*FieldOption) [GetMaxValue](./spec.pb.go#L276)
``` go
func (m *FieldOption) GetMaxValue() float64
```

### <a name="FieldOption.GetMinLength">func</a> (\*FieldOption) [GetMinLength](./spec.pb.go#L304)
``` go
func (m *FieldOption) GetMinLength() int32
```

### <a name="FieldOption.GetMinValue">func</a> (\*FieldOption) [GetMinValue](./spec.pb.go#L269)
``` go
func (m *FieldOption) GetMinValue() float64
```

### <a name="FieldOption.GetMultipleOfValue">func</a> (\*FieldOption) [GetMultipleOfValue](./spec.pb.go#L283)
``` go
func (m *FieldOption) GetMultipleOfValue() float64
```

### <a name="FieldOption.GetRegexpValue">func</a> (\*FieldOption) [GetRegexpValue](./spec.pb.go#L318)
``` go
func (m *FieldOption) GetRegexpValue() string
```

### <a name="FieldOption.GetStructTag">func</a> (\*FieldOption) [GetStructTag](./spec.pb.go#L339)
``` go
func (m *FieldOption) GetStructTag() string
```

### <a name="FieldOption.ProtoMessage">func</a> (\*FieldOption) [ProtoMessage](./spec.pb.go#L266)
``` go
func (*FieldOption) ProtoMessage()
```

### <a name="FieldOption.Reset">func</a> (\*FieldOption) [Reset](./spec.pb.go#L264)
``` go
func (m *FieldOption) Reset()
```

### <a name="FieldOption.String">func</a> (\*FieldOption) [String](./spec.pb.go#L265)
``` go
func (m *FieldOption) String() string
```

## <a name="FileOption">type</a> [FileOption](./spec.pb.go#L157-L160)
``` go
type FileOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="FileOption.Descriptor">func</a> (\*FileOption) [Descriptor](./spec.pb.go#L165)
``` go
func (*FileOption) Descriptor() ([]byte, []int)
```

### <a name="FileOption.GetExternalDocs">func</a> (\*FileOption) [GetExternalDocs](./spec.pb.go#L167)
``` go
func (m *FileOption) GetExternalDocs() *ExternalDocumentation
```

### <a name="FileOption.ProtoMessage">func</a> (\*FileOption) [ProtoMessage](./spec.pb.go#L164)
``` go
func (*FileOption) ProtoMessage()
```

### <a name="FileOption.Reset">func</a> (\*FileOption) [Reset](./spec.pb.go#L162)
``` go
func (m *FileOption) Reset()
```

### <a name="FileOption.String">func</a> (\*FileOption) [String](./spec.pb.go#L163)
``` go
func (m *FileOption) String() string
```

## <a name="FileSpec">type</a> [FileSpec](./spec.pb.go#L40-L45)
``` go
type FileSpec struct {
    FileName         *string                `protobuf:"bytes,1,req,name=file_name,json=fileName" json:"file_name,omitempty"`
    PackageName      *string                `protobuf:"bytes,2,req,name=package_name,json=packageName" json:"package_name,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="FileSpec.Descriptor">func</a> (\*FileSpec) [Descriptor](./spec.pb.go#L50)
``` go
func (*FileSpec) Descriptor() ([]byte, []int)
```

### <a name="FileSpec.GetExternalDocs">func</a> (\*FileSpec) [GetExternalDocs](./spec.pb.go#L66)
``` go
func (m *FileSpec) GetExternalDocs() *ExternalDocumentation
```

### <a name="FileSpec.GetFileName">func</a> (\*FileSpec) [GetFileName](./spec.pb.go#L52)
``` go
func (m *FileSpec) GetFileName() string
```

### <a name="FileSpec.GetPackageName">func</a> (\*FileSpec) [GetPackageName](./spec.pb.go#L59)
``` go
func (m *FileSpec) GetPackageName() string
```

### <a name="FileSpec.ProtoMessage">func</a> (\*FileSpec) [ProtoMessage](./spec.pb.go#L49)
``` go
func (*FileSpec) ProtoMessage()
```

### <a name="FileSpec.Reset">func</a> (\*FileSpec) [Reset](./spec.pb.go#L47)
``` go
func (m *FileSpec) Reset()
```

### <a name="FileSpec.String">func</a> (\*FileSpec) [String](./spec.pb.go#L48)
``` go
func (m *FileSpec) String() string
```

## <a name="MessageOption">type</a> [MessageOption](./spec.pb.go#L232-L235)
``` go
type MessageOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="MessageOption.Descriptor">func</a> (\*MessageOption) [Descriptor](./spec.pb.go#L240)
``` go
func (*MessageOption) Descriptor() ([]byte, []int)
```

### <a name="MessageOption.GetExternalDocs">func</a> (\*MessageOption) [GetExternalDocs](./spec.pb.go#L242)
``` go
func (m *MessageOption) GetExternalDocs() *ExternalDocumentation
```

### <a name="MessageOption.ProtoMessage">func</a> (\*MessageOption) [ProtoMessage](./spec.pb.go#L239)
``` go
func (*MessageOption) ProtoMessage()
```

### <a name="MessageOption.Reset">func</a> (\*MessageOption) [Reset](./spec.pb.go#L237)
``` go
func (m *MessageOption) Reset()
```

### <a name="MessageOption.String">func</a> (\*MessageOption) [String](./spec.pb.go#L238)
``` go
func (m *MessageOption) String() string
```

## <a name="MethodOption">type</a> [MethodOption](./spec.pb.go#L199-L204)
``` go
type MethodOption struct {
    HttpMethod       *string                `protobuf:"bytes,1,opt,name=http_method,json=httpMethod" json:"http_method,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,2,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    CmdInfo          *CommandInfo           `protobuf:"bytes,3,opt,name=cmd_info,json=cmdInfo" json:"cmd_info,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="MethodOption.Descriptor">func</a> (\*MethodOption) [Descriptor](./spec.pb.go#L209)
``` go
func (*MethodOption) Descriptor() ([]byte, []int)
```

### <a name="MethodOption.GetCmdInfo">func</a> (\*MethodOption) [GetCmdInfo](./spec.pb.go#L225)
``` go
func (m *MethodOption) GetCmdInfo() *CommandInfo
```

### <a name="MethodOption.GetExternalDocs">func</a> (\*MethodOption) [GetExternalDocs](./spec.pb.go#L218)
``` go
func (m *MethodOption) GetExternalDocs() *ExternalDocumentation
```

### <a name="MethodOption.GetHttpMethod">func</a> (\*MethodOption) [GetHttpMethod](./spec.pb.go#L211)
``` go
func (m *MethodOption) GetHttpMethod() string
```

### <a name="MethodOption.ProtoMessage">func</a> (\*MethodOption) [ProtoMessage](./spec.pb.go#L208)
``` go
func (*MethodOption) ProtoMessage()
```

### <a name="MethodOption.Reset">func</a> (\*MethodOption) [Reset](./spec.pb.go#L206)
``` go
func (m *MethodOption) Reset()
```

### <a name="MethodOption.String">func</a> (\*MethodOption) [String](./spec.pb.go#L207)
``` go
func (m *MethodOption) String() string
```

## <a name="ServiceMethodSpec">type</a> [ServiceMethodSpec](./spec.pb.go#L106-L113)
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

### <a name="ServiceMethodSpec.Descriptor">func</a> (\*ServiceMethodSpec) [Descriptor](./spec.pb.go#L118)
``` go
func (*ServiceMethodSpec) Descriptor() ([]byte, []int)
```

### <a name="ServiceMethodSpec.GetExternalDocs">func</a> (\*ServiceMethodSpec) [GetExternalDocs](./spec.pb.go#L150)
``` go
func (m *ServiceMethodSpec) GetExternalDocs() *ExternalDocumentation
```

### <a name="ServiceMethodSpec.GetHttpMethod">func</a> (\*ServiceMethodSpec) [GetHttpMethod](./spec.pb.go#L143)
``` go
func (m *ServiceMethodSpec) GetHttpMethod() string
```

### <a name="ServiceMethodSpec.GetInputTypeName">func</a> (\*ServiceMethodSpec) [GetInputTypeName](./spec.pb.go#L129)
``` go
func (m *ServiceMethodSpec) GetInputTypeName() string
```

### <a name="ServiceMethodSpec.GetMethodName">func</a> (\*ServiceMethodSpec) [GetMethodName](./spec.pb.go#L122)
``` go
func (m *ServiceMethodSpec) GetMethodName() string
```

### <a name="ServiceMethodSpec.GetOutputTypeName">func</a> (\*ServiceMethodSpec) [GetOutputTypeName](./spec.pb.go#L136)
``` go
func (m *ServiceMethodSpec) GetOutputTypeName() string
```

### <a name="ServiceMethodSpec.ProtoMessage">func</a> (\*ServiceMethodSpec) [ProtoMessage](./spec.pb.go#L117)
``` go
func (*ServiceMethodSpec) ProtoMessage()
```

### <a name="ServiceMethodSpec.Reset">func</a> (\*ServiceMethodSpec) [Reset](./spec.pb.go#L115)
``` go
func (m *ServiceMethodSpec) Reset()
```

### <a name="ServiceMethodSpec.String">func</a> (\*ServiceMethodSpec) [String](./spec.pb.go#L116)
``` go
func (m *ServiceMethodSpec) String() string
```

## <a name="ServiceOption">type</a> [ServiceOption](./spec.pb.go#L174-L178)
``` go
type ServiceOption struct {
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,1,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    CmdInfo          *CommandInfo           `protobuf:"bytes,2,opt,name=cmd_info,json=cmdInfo" json:"cmd_info,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="ServiceOption.Descriptor">func</a> (\*ServiceOption) [Descriptor](./spec.pb.go#L183)
``` go
func (*ServiceOption) Descriptor() ([]byte, []int)
```

### <a name="ServiceOption.GetCmdInfo">func</a> (\*ServiceOption) [GetCmdInfo](./spec.pb.go#L192)
``` go
func (m *ServiceOption) GetCmdInfo() *CommandInfo
```

### <a name="ServiceOption.GetExternalDocs">func</a> (\*ServiceOption) [GetExternalDocs](./spec.pb.go#L185)
``` go
func (m *ServiceOption) GetExternalDocs() *ExternalDocumentation
```

### <a name="ServiceOption.ProtoMessage">func</a> (\*ServiceOption) [ProtoMessage](./spec.pb.go#L182)
``` go
func (*ServiceOption) ProtoMessage()
```

### <a name="ServiceOption.Reset">func</a> (\*ServiceOption) [Reset](./spec.pb.go#L180)
``` go
func (m *ServiceOption) Reset()
```

### <a name="ServiceOption.String">func</a> (\*ServiceOption) [String](./spec.pb.go#L181)
``` go
func (m *ServiceOption) String() string
```

## <a name="ServiceSpec">type</a> [ServiceSpec](./spec.pb.go#L73-L78)
``` go
type ServiceSpec struct {
    ServiceName      *string                `protobuf:"bytes,1,req,name=service_name,json=serviceName" json:"service_name,omitempty"`
    MethodList       []*ServiceMethodSpec   `protobuf:"bytes,2,rep,name=method_list,json=methodList" json:"method_list,omitempty"`
    ExternalDocs     *ExternalDocumentation `protobuf:"bytes,3,opt,name=external_docs,json=externalDocs" json:"external_docs,omitempty"`
    XXX_unrecognized []byte                 `json:"-"`
}
```

### <a name="ServiceSpec.Descriptor">func</a> (\*ServiceSpec) [Descriptor](./spec.pb.go#L83)
``` go
func (*ServiceSpec) Descriptor() ([]byte, []int)
```

### <a name="ServiceSpec.GetExternalDocs">func</a> (\*ServiceSpec) [GetExternalDocs](./spec.pb.go#L99)
``` go
func (m *ServiceSpec) GetExternalDocs() *ExternalDocumentation
```

### <a name="ServiceSpec.GetMethodList">func</a> (\*ServiceSpec) [GetMethodList](./spec.pb.go#L92)
``` go
func (m *ServiceSpec) GetMethodList() []*ServiceMethodSpec
```

### <a name="ServiceSpec.GetServiceName">func</a> (\*ServiceSpec) [GetServiceName](./spec.pb.go#L85)
``` go
func (m *ServiceSpec) GetServiceName() string
```

### <a name="ServiceSpec.ProtoMessage">func</a> (\*ServiceSpec) [ProtoMessage](./spec.pb.go#L82)
``` go
func (*ServiceSpec) ProtoMessage()
```

### <a name="ServiceSpec.Reset">func</a> (\*ServiceSpec) [Reset](./spec.pb.go#L80)
``` go
func (m *ServiceSpec) Reset()
```

### <a name="ServiceSpec.String">func</a> (\*ServiceSpec) [String](./spec.pb.go#L81)
``` go
func (m *ServiceSpec) String() string
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)