# qingcloud_plugin
`import "github.com/chai2010/qingcloud-go/pkg/cmd/protoc-gen-qingcloud"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
protoc-gen-go is a plugin for the Google protocol buffer compiler to generate
Go code.  Run it by building this program and putting it in your path with
the name

	protoc-gen-go

That word 'go' at the end becomes part of the option string set for the
protocol compiler, so once the protocol compiler (protoc) is installed
you can run

	protoc --go_out=output_directory input_directory/file.proto

to generate Go bindings for the protocol defined by file.proto.
With that input, the output will be written to

	output_directory/file.pb.go

The generated code is documented in the package comment for
the library.

See the README and documentation for protocol buffers to learn more:

	<a href="https://developers.google.com/protocol-buffers/">https://developers.google.com/protocol-buffers/</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/golang/protobuf/proto](https://godoc.org/github.com/golang/protobuf/proto)
- [github.com/golang/protobuf/protoc-gen-go/descriptor](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/descriptor)
- [github.com/golang/protobuf/protoc-gen-go/generator](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/generator)
- [github.com/golang/protobuf/protoc-gen-go/plugin](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/plugin)

## <a name="pkg-index">Index</a>
* [func Main()](#Main)
* [func RegisterServiceGenerator(g ServiceGenerator)](#RegisterServiceGenerator)
* [type ServiceGenerator](#ServiceGenerator)

#### <a name="pkg-files">Package files</a>
[main.go](./main.go) [plugin.go](./plugin.go) [qingcloud.go](./qingcloud.go) 

## <a name="Main">func</a> [Main](./main.go#L34)
``` go
func Main()
```

## <a name="RegisterServiceGenerator">func</a> [RegisterServiceGenerator](./plugin.go#L23)
``` go
func RegisterServiceGenerator(g ServiceGenerator)
```

## <a name="ServiceGenerator">type</a> [ServiceGenerator](./plugin.go#L14-L21)
``` go
type ServiceGenerator interface {
    Name() string
    FileNameExt() string

    HeaderCode(g *generator.Generator, file *generator.FileDescriptor) string
    ServiceCode(p *generator.Generator, file *generator.FileDescriptor, svc *descriptor.ServiceDescriptorProto) string
    MessageCode(p *generator.Generator, file *generator.FileDescriptor, msg *descriptor.DescriptorProto) string
}
```

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)