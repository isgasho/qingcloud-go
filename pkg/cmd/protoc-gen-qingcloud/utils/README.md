# utils
`import "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud/utils"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/chai2010/qingcloud-go/pkg/api/spec_metadata](./../../../api/spec_metadata)
- [github.com/golang/protobuf/descriptor](https://godoc.org/github.com/golang/protobuf/descriptor)
- [github.com/golang/protobuf/proto](https://godoc.org/github.com/golang/protobuf/proto)
- [github.com/golang/protobuf/protoc-gen-go/descriptor](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/descriptor)
- [github.com/golang/protobuf/protoc-gen-go/generator](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/generator)

## <a name="pkg-index">Index</a>
* [func BuildFileSpec(p \*generator.Generator, file \*generator.FileDescriptor) \*spec\_metadata.FileSpec](#BuildFileSpec)
* [func BuildServiceSpec(p \*generator.Generator, file \*generator.FileDescriptor, svc \*descriptor.ServiceDescriptorProto) \*spec\_metadata.ServiceSpec](#BuildServiceSpec)
* [func GetFileOption(file \*descriptor.FileDescriptorProto) \*spec\_metadata.FileOption](#GetFileOption)
* [func GetMessageFieldOption(m \*descriptor.FieldDescriptorProto) \*spec\_metadata.FieldOption](#GetMessageFieldOption)
* [func GetMessageOption(m \*descriptor.DescriptorProto) \*spec\_metadata.MessageOption](#GetMessageOption)
* [func GetMethodInputDescriptor(p \*generator.Generator, m \*descriptor.MethodDescriptorProto) \*descriptor.DescriptorProto](#GetMethodInputDescriptor)
* [func GetMethodOutputDescriptor(p \*generator.Generator, m \*descriptor.MethodDescriptorProto) \*descriptor.DescriptorProto](#GetMethodOutputDescriptor)
* [func GetServiceMethodOption(m \*descriptor.MethodDescriptorProto) \*spec\_metadata.MethodOption](#GetServiceMethodOption)
* [func GetServiceOption(svc \*descriptor.ServiceDescriptorProto) \*spec\_metadata.ServiceOption](#GetServiceOption)
* [func HasFieldOptions(message \*descriptor.DescriptorProto) bool](#HasFieldOptions)
* [func IsSupportedBool(field \*descriptor.FieldDescriptorProto) bool](#IsSupportedBool)
* [func IsSupportedFloat(field \*descriptor.FieldDescriptorProto) bool](#IsSupportedFloat)
* [func IsSupportedInt(field \*descriptor.FieldDescriptorProto) bool](#IsSupportedInt)
* [func IsSupportedRepeated(field \*descriptor.FieldDescriptorProto) bool](#IsSupportedRepeated)
* [func IsSupportedString(field \*descriptor.FieldDescriptorProto) bool](#IsSupportedString)

#### <a name="pkg-files">Package files</a>
[utils.go](./utils.go) 

## <a name="BuildFileSpec">func</a> [BuildFileSpec](./utils.go#L23)
``` go
func BuildFileSpec(p *generator.Generator, file *generator.FileDescriptor) *spec_metadata.FileSpec
```

## <a name="BuildServiceSpec">func</a> [BuildServiceSpec](./utils.go#L30)
``` go
func BuildServiceSpec(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceSpec
```

## <a name="GetFileOption">func</a> [GetFileOption](./utils.go#L65)
``` go
func GetFileOption(file *descriptor.FileDescriptorProto) *spec_metadata.FileOption
```

## <a name="GetMessageFieldOption">func</a> [GetMessageFieldOption](./utils.go#L109)
``` go
func GetMessageFieldOption(m *descriptor.FieldDescriptorProto) *spec_metadata.FieldOption
```

## <a name="GetMessageOption">func</a> [GetMessageOption](./utils.go#L98)
``` go
func GetMessageOption(m *descriptor.DescriptorProto) *spec_metadata.MessageOption
```

## <a name="GetMethodInputDescriptor">func</a> [GetMethodInputDescriptor](./utils.go#L120)
``` go
func GetMethodInputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto
```

## <a name="GetMethodOutputDescriptor">func</a> [GetMethodOutputDescriptor](./utils.go#L123)
``` go
func GetMethodOutputDescriptor(p *generator.Generator, m *descriptor.MethodDescriptorProto) *descriptor.DescriptorProto
```

## <a name="GetServiceMethodOption">func</a> [GetServiceMethodOption](./utils.go#L87)
``` go
func GetServiceMethodOption(m *descriptor.MethodDescriptorProto) *spec_metadata.MethodOption
```

## <a name="GetServiceOption">func</a> [GetServiceOption](./utils.go#L76)
``` go
func GetServiceOption(svc *descriptor.ServiceDescriptorProto) *spec_metadata.ServiceOption
```

## <a name="HasFieldOptions">func</a> [HasFieldOptions](./utils.go#L56)
``` go
func HasFieldOptions(message *descriptor.DescriptorProto) bool
```

## <a name="IsSupportedBool">func</a> [IsSupportedBool](./utils.go#L135)
``` go
func IsSupportedBool(field *descriptor.FieldDescriptorProto) bool
```

## <a name="IsSupportedFloat">func</a> [IsSupportedFloat](./utils.go#L161)
``` go
func IsSupportedFloat(field *descriptor.FieldDescriptorProto) bool
```

## <a name="IsSupportedInt">func</a> [IsSupportedInt](./utils.go#L146)
``` go
func IsSupportedInt(field *descriptor.FieldDescriptorProto) bool
```

## <a name="IsSupportedRepeated">func</a> [IsSupportedRepeated](./utils.go#L127)
``` go
func IsSupportedRepeated(field *descriptor.FieldDescriptorProto) bool
```

## <a name="IsSupportedString">func</a> [IsSupportedString](./utils.go#L172)
``` go
func IsSupportedString(field *descriptor.FieldDescriptorProto) bool
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)