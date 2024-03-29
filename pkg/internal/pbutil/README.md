# pbutil
`import "github.com/chai2010/qingcloud-go/pkg/pbutil"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/golang/protobuf/descriptor](https://godoc.org/github.com/golang/protobuf/descriptor)
- [github.com/golang/protobuf/jsonpb](https://godoc.org/github.com/golang/protobuf/jsonpb)
- [github.com/golang/protobuf/proto](https://godoc.org/github.com/golang/protobuf/proto)
- [github.com/golang/protobuf/protoc-gen-go/descriptor](https://godoc.org/github.com/golang/protobuf/protoc-gen-go/descriptor)

## <a name="pkg-index">Index</a>
* [func DecodeFromMap(m map[string]string, x proto.Message) error](#DecodeFromMap)
* [func DecodeJson(data []byte, x proto.Message) error](#DecodeJson)
* [func EncodeJson(x proto.Message) ([]byte, error)](#EncodeJson)
* [func EncodeJsonIndent(m proto.Message) (string, error)](#EncodeJsonIndent)
* [func EncodeToMap(msg proto.Message) (m map[string]string, err error)](#EncodeToMap)
* [func GetMessageDescriptor(msg proto.Message) (fd \*protobuf.FileDescriptorProto, md \*protobuf.DescriptorProto)](#GetMessageDescriptor)
* [func JsonToMapString(s string) (map[string]string, error)](#JsonToMapString)
* [func ProtoMessageToMap(msg proto.Message) (m map[string]string, err error)](#ProtoMessageToMap)
* [func UnpackMapXToMapString(mapx map[string]interface{}) map[string]string](#UnpackMapXToMapString)

#### <a name="pkg-files">Package files</a>
[map.go](./map.go) [pbutil.go](./pbutil.go) 

## <a name="DecodeFromMap">func</a> [DecodeFromMap](./map.go#L15)
``` go
func DecodeFromMap(m map[string]string, x proto.Message) error
```

## <a name="DecodeJson">func</a> [DecodeJson](./pbutil.go#L47)
``` go
func DecodeJson(data []byte, x proto.Message) error
```

## <a name="EncodeJson">func</a> [EncodeJson](./pbutil.go#L19)
``` go
func EncodeJson(x proto.Message) ([]byte, error)
```

## <a name="EncodeJsonIndent">func</a> [EncodeJsonIndent](./pbutil.go#L33)
``` go
func EncodeJsonIndent(m proto.Message) (string, error)
```

## <a name="EncodeToMap">func</a> [EncodeToMap](./map.go#L11)
``` go
func EncodeToMap(msg proto.Message) (m map[string]string, err error)
```

## <a name="GetMessageDescriptor">func</a> [GetMessageDescriptor](./pbutil.go#L128)
``` go
func GetMessageDescriptor(msg proto.Message) (fd *protobuf.FileDescriptorProto, md *protobuf.DescriptorProto)
```

## <a name="JsonToMapString">func</a> [JsonToMapString](./pbutil.go#L79)
``` go
func JsonToMapString(s string) (map[string]string, error)
```

## <a name="ProtoMessageToMap">func</a> [ProtoMessageToMap](./pbutil.go#L58)
``` go
func ProtoMessageToMap(msg proto.Message) (m map[string]string, err error)
```

## <a name="UnpackMapXToMapString">func</a> [UnpackMapXToMapString](./pbutil.go#L88)
``` go
func UnpackMapXToMapString(mapx map[string]interface{}) map[string]string
```
X is oneof string/float64/[]interface/map[string]interface{}

- - -
Generated by [godoc2ghmd](https://github.com/GandalfUK/godoc2ghmd)