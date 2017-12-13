

# utils
`import "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func BuildFileSpec(p *generator.Generator, file *generator.FileDescriptor) *spec_metadata.FileSpec](#BuildFileSpec)
* [func BuildServiceSpec(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceSpec](#BuildServiceSpec)
* [func GetFileOption(file *descriptor.FileDescriptorProto) *spec_metadata.FileOption](#GetFileOption)
* [func GetMessageFieldOption(m *descriptor.FieldDescriptorProto) *spec_metadata.FieldOption](#GetMessageFieldOption)
* [func GetMessageOption(m *descriptor.DescriptorProto) *spec_metadata.MessageOption](#GetMessageOption)
* [func GetMethodInputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto](#GetMethodInputDescriptor)
* [func GetMethodOutputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto](#GetMethodOutputDescriptor)
* [func GetServiceMethodOption(m *descriptor.MethodDescriptorProto) *spec_metadata.MethodOption](#GetServiceMethodOption)
* [func GetServiceOption(svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceOption](#GetServiceOption)
* [func HasFieldOptions(message *descriptor.DescriptorProto) bool](#HasFieldOptions)
* [func IsSupportedBool(field *descriptor.FieldDescriptorProto) bool](#IsSupportedBool)
* [func IsSupportedFloat(field *descriptor.FieldDescriptorProto) bool](#IsSupportedFloat)
* [func IsSupportedInt(field *descriptor.FieldDescriptorProto) bool](#IsSupportedInt)
* [func IsSupportedRepeated(field *descriptor.FieldDescriptorProto) bool](#IsSupportedRepeated)
* [func IsSupportedString(field *descriptor.FieldDescriptorProto) bool](#IsSupportedString)


#### <a name="pkg-files">Package files</a>
[utils.go](/src/github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils/utils.go) 





## <a name="BuildFileSpec">func</a> [BuildFileSpec](/src/target/utils.go?s=684:782#L23)
``` go
func BuildFileSpec(p *generator.Generator, file *generator.FileDescriptor) *spec_metadata.FileSpec
```


## <a name="BuildServiceSpec">func</a> [BuildServiceSpec](/src/target/utils.go?s=918:1062#L30)
``` go
func BuildServiceSpec(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceSpec
```


## <a name="GetFileOption">func</a> [GetFileOption](/src/target/utils.go?s=1951:2033#L65)
``` go
func GetFileOption(file *descriptor.FileDescriptorProto) *spec_metadata.FileOption
```


## <a name="GetMessageFieldOption">func</a> [GetMessageFieldOption](/src/target/utils.go?s=3417:3506#L109)
``` go
func GetMessageFieldOption(m *descriptor.FieldDescriptorProto) *spec_metadata.FieldOption
```


## <a name="GetMessageOption">func</a> [GetMessageOption](/src/target/utils.go?s=3055:3136#L98)
``` go
func GetMessageOption(m *descriptor.DescriptorProto) *spec_metadata.MessageOption
```


## <a name="GetMethodInputDescriptor">func</a> [GetMethodInputDescriptor](/src/target/utils.go?s=3781:3899#L120)
``` go
func GetMethodInputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto
```


## <a name="GetMethodOutputDescriptor">func</a> [GetMethodOutputDescriptor](/src/target/utils.go?s=3984:4103#L123)
``` go
func GetMethodOutputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto
```


## <a name="GetServiceMethodOption">func</a> [GetServiceMethodOption](/src/target/utils.go?s=2685:2777#L87)
``` go
func GetServiceMethodOption(m *descriptor.MethodDescriptorProto) *spec_metadata.MethodOption
```


## <a name="GetServiceOption">func</a> [GetServiceOption](/src/target/utils.go?s=2314:2404#L76)
``` go
func GetServiceOption(svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceOption
```


## <a name="HasFieldOptions">func</a> [HasFieldOptions](/src/target/utils.go?s=1760:1822#L56)
``` go
func HasFieldOptions(message *descriptor.DescriptorProto) bool
```


## <a name="IsSupportedBool">func</a> [IsSupportedBool](/src/target/utils.go?s=4377:4442#L135)
``` go
func IsSupportedBool(field *descriptor.FieldDescriptorProto) bool
```


## <a name="IsSupportedFloat">func</a> [IsSupportedFloat](/src/target/utils.go?s=5097:5163#L161)
``` go
func IsSupportedFloat(field *descriptor.FieldDescriptorProto) bool
```


## <a name="IsSupportedInt">func</a> [IsSupportedInt](/src/target/utils.go?s=4605:4669#L146)
``` go
func IsSupportedInt(field *descriptor.FieldDescriptorProto) bool
```


## <a name="IsSupportedRepeated">func</a> [IsSupportedRepeated](/src/target/utils.go?s=4190:4259#L127)
``` go
func IsSupportedRepeated(field *descriptor.FieldDescriptorProto) bool
```


## <a name="IsSupportedString">func</a> [IsSupportedString](/src/target/utils.go?s=5372:5439#L172)
``` go
func IsSupportedString(field *descriptor.FieldDescriptorProto) bool
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
