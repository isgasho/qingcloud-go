// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: cache.proto

package qcli_pb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = json.Marshal
	_ = log.Print
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdCacheService)
}

var CmdCacheService = cli.Command{
	Name:    "cache",
	Aliases: []string{},
	Usage:   "manage CacheService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeCaches",
			Aliases: []string{},
			Usage:   "DescribeCaches",
			Flags:   _flag_CacheService_DescribeCaches,
			Action:  _func_CacheService_DescribeCaches,
		},
		{
			Name:    "CreateCache",
			Aliases: []string{},
			Usage:   "CreateCache",
			Flags:   _flag_CacheService_CreateCache,
			Action:  _func_CacheService_CreateCache,
		},
		{
			Name:    "StopCaches",
			Aliases: []string{},
			Usage:   "StopCaches",
			Flags:   _flag_CacheService_StopCaches,
			Action:  _func_CacheService_StopCaches,
		},
		{
			Name:    "StartCaches",
			Aliases: []string{},
			Usage:   "StartCaches",
			Flags:   _flag_CacheService_StartCaches,
			Action:  _func_CacheService_StartCaches,
		},
		{
			Name:    "RestartCaches",
			Aliases: []string{},
			Usage:   "RestartCaches",
			Flags:   _flag_CacheService_RestartCaches,
			Action:  _func_CacheService_RestartCaches,
		},
		{
			Name:    "DeleteCaches",
			Aliases: []string{},
			Usage:   "DeleteCaches",
			Flags:   _flag_CacheService_DeleteCaches,
			Action:  _func_CacheService_DeleteCaches,
		},
		{
			Name:    "ResizeCaches",
			Aliases: []string{},
			Usage:   "ResizeCaches",
			Flags:   _flag_CacheService_ResizeCaches,
			Action:  _func_CacheService_ResizeCaches,
		},
		{
			Name:    "UpdateCache",
			Aliases: []string{},
			Usage:   "UpdateCache",
			Flags:   _flag_CacheService_UpdateCache,
			Action:  _func_CacheService_UpdateCache,
		},
		{
			Name:    "ChangeCacheVxnet",
			Aliases: []string{},
			Usage:   "ChangeCacheVxnet",
			Flags:   _flag_CacheService_ChangeCacheVxnet,
			Action:  _func_CacheService_ChangeCacheVxnet,
		},
		{
			Name:    "ModifyCacheAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheAttributes",
			Flags:   _flag_CacheService_ModifyCacheAttributes,
			Action:  _func_CacheService_ModifyCacheAttributes,
		},
		{
			Name:    "DescribeCacheNodes",
			Aliases: []string{},
			Usage:   "DescribeCacheNodes",
			Flags:   _flag_CacheService_DescribeCacheNodes,
			Action:  _func_CacheService_DescribeCacheNodes,
		},
		{
			Name:    "AddCacheNodes",
			Aliases: []string{},
			Usage:   "AddCacheNodes",
			Flags:   _flag_CacheService_AddCacheNodes,
			Action:  _func_CacheService_AddCacheNodes,
		},
		{
			Name:    "DeleteCacheNodes",
			Aliases: []string{},
			Usage:   "DeleteCacheNodes",
			Flags:   _flag_CacheService_DeleteCacheNodes,
			Action:  _func_CacheService_DeleteCacheNodes,
		},
		{
			Name:    "RestartCacheNodes",
			Aliases: []string{},
			Usage:   "RestartCacheNodes",
			Flags:   _flag_CacheService_RestartCacheNodes,
			Action:  _func_CacheService_RestartCacheNodes,
		},
		{
			Name:    "ModifyCacheNodeAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheNodeAttributes",
			Flags:   _flag_CacheService_ModifyCacheNodeAttributes,
			Action:  _func_CacheService_ModifyCacheNodeAttributes,
		},
		{
			Name:    "CreateCacheFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateCacheFromSnapshot",
			Flags:   _flag_CacheService_CreateCacheFromSnapshot,
			Action:  _func_CacheService_CreateCacheFromSnapshot,
		},
		{
			Name:    "DescribeCacheParameterGroups",
			Aliases: []string{},
			Usage:   "DescribeCacheParameterGroups",
			Flags:   _flag_CacheService_DescribeCacheParameterGroups,
			Action:  _func_CacheService_DescribeCacheParameterGroups,
		},
		{
			Name:    "CreateCacheParameterGroup",
			Aliases: []string{},
			Usage:   "CreateCacheParameterGroup",
			Flags:   _flag_CacheService_CreateCacheParameterGroup,
			Action:  _func_CacheService_CreateCacheParameterGroup,
		},
		{
			Name:    "ApplyCacheParameterGroup",
			Aliases: []string{},
			Usage:   "ApplyCacheParameterGroup",
			Flags:   _flag_CacheService_ApplyCacheParameterGroup,
			Action:  _func_CacheService_ApplyCacheParameterGroup,
		},
		{
			Name:    "DeleteCacheParameterGroups",
			Aliases: []string{},
			Usage:   "DeleteCacheParameterGroups",
			Flags:   _flag_CacheService_DeleteCacheParameterGroups,
			Action:  _func_CacheService_DeleteCacheParameterGroups,
		},
		{
			Name:    "ModifyCacheParameterGroupAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheParameterGroupAttributes",
			Flags:   _flag_CacheService_ModifyCacheParameterGroupAttributes,
			Action:  _func_CacheService_ModifyCacheParameterGroupAttributes,
		},
		{
			Name:    "DescribeCacheParameters",
			Aliases: []string{},
			Usage:   "DescribeCacheParameters",
			Flags:   _flag_CacheService_DescribeCacheParameters,
			Action:  _func_CacheService_DescribeCacheParameters,
		},
		{
			Name:    "UpdateCacheParameters",
			Aliases: []string{},
			Usage:   "UpdateCacheParameters",
			Flags:   _flag_CacheService_UpdateCacheParameters,
			Action:  _func_CacheService_UpdateCacheParameters,
		},
		{
			Name:    "ResetCacheParameters",
			Aliases: []string{},
			Usage:   "ResetCacheParameters",
			Flags:   _flag_CacheService_ResetCacheParameters,
			Action:  _func_CacheService_ResetCacheParameters,
		},
	},
}

var _flag_CacheService_DescribeCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "cache_type",
		Usage: "cache type",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "search_word",
		Usage: "search word",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tags",
		Usage: "tags",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "verbose",
		Usage: "verbose",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
}

func _func_CacheService_DescribeCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("cache_type") {
			if err := json.Unmarshal([]byte(c.String("cache_type")), &in.CacheType); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.String(c.String("search_word"))
		}
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCache = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.IntFlag{
		Name:  "cache_size",
		Usage: "cache size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "cache_type",
		Usage: "cache type",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "cache_name",
		Usage: "cache name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "auto_backup_time",
		Usage: "auto backup time",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "cache_class",
		Usage: "cache class",
		Value: 0,
	},
}

func _func_CacheService_CreateCache(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("cache_size") {
			in.CacheSize = proto.Int32(int32(c.Int("cache_size")))
		}
		if c.IsSet("cache_type") {
			in.CacheType = proto.String(c.String("cache_type"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("cache_name") {
			in.CacheName = proto.String(c.String("cache_name"))
		}
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
		if c.IsSet("cache_class") {
			in.CacheClass = proto.Int32(int32(c.Int("cache_class")))
		}
	}

	out, err := qc.CreateCache(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_StopCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_StopCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.StopCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StopCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_StartCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_StartCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.StartCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StartCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_RestartCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_RestartCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.RestartCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.RestartCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_DeleteCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ResizeCaches = []cli.Flag{
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "cache_size",
		Usage: "cache size",
		Value: 0,
	},
}

func _func_CacheService_ResizeCaches(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ResizeCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("cache_size") {
			in.CacheSize = proto.Int32(int32(c.Int("cache_size")))
		}
	}

	out, err := qc.ResizeCaches(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_UpdateCache = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_UpdateCache(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.UpdateCacheInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.UpdateCache(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ChangeCacheVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_ChangeCacheVxnet(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ChangeCacheVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ChangeCacheVxnet(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_name",
		Usage: "cache name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
	cli.IntFlag{
		Name:  "auto_backup_time",
		Usage: "auto backup time",
		Value: 0,
	},
}

func _func_CacheService_ModifyCacheAttributes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("cache_name") {
			in.CacheName = proto.String(c.String("cache_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
	}

	out, err := qc.ModifyCacheAttributes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_nodes",
		Usage: "cache nodes",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "search_word",
		Usage: "search word",
		Value: "",
	},
	cli.IntFlag{
		Name:  "verbose",
		Usage: "verbose",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
}

func _func_CacheService_DescribeCacheNodes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("cache_nodes") {
			if err := json.Unmarshal([]byte(c.String("cache_nodes")), &in.CacheNodes); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.String(c.String("search_word"))
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeCacheNodes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_AddCacheNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_AddCacheNodes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.AddCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddCacheNodes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCacheNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_nodes",
		Usage: "cache nodes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_DeleteCacheNodes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("cache_nodes") {
			if err := json.Unmarshal([]byte(c.String("cache_nodes")), &in.CacheNodes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteCacheNodes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_RestartCacheNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache",
		Usage: "cache",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_nodes",
		Usage: "cache nodes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_RestartCacheNodes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.RestartCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache") {
			in.Cache = proto.String(c.String("cache"))
		}
		if c.IsSet("cache_nodes") {
			if err := json.Unmarshal([]byte(c.String("cache_nodes")), &in.CacheNodes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.RestartCacheNodes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheNodeAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_node",
		Usage: "cache node",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_node_name",
		Usage: "cache node name",
		Value: "",
	},
}

func _func_CacheService_ModifyCacheNodeAttributes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheNodeAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_node") {
			in.CacheNode = proto.String(c.String("cache_node"))
		}
		if c.IsSet("cache_node_name") {
			in.CacheNodeName = proto.String(c.String("cache_node_name"))
		}
	}

	out, err := qc.ModifyCacheNodeAttributes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCacheFromSnapshot = []cli.Flag{
	cli.StringFlag{
		Name:  "snapshot",
		Usage: "snapshot",
		Value: "",
	},
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "cache_name",
		Usage: "cache name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.IntFlag{
		Name:  "auto_backup_time",
		Usage: "auto backup time",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "cache_class",
		Usage: "cache class",
		Value: 0,
	},
}

func _func_CacheService_CreateCacheFromSnapshot(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("snapshot") {
			in.Snapshot = proto.String(c.String("snapshot"))
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("cache_name") {
			in.CacheName = proto.String(c.String("cache_name"))
		}
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("cache_class") {
			in.CacheClass = proto.Int32(int32(c.Int("cache_class")))
		}
	}

	out, err := qc.CreateCacheFromSnapshot(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheParameterGroups = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_groups",
		Usage: "cache parameter groups",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "cache_type",
		Usage: "cache type",
		Value: "",
	},
	cli.IntFlag{
		Name:  "search_word",
		Usage: "search word",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "verbose",
		Usage: "verbose",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
}

func _func_CacheService_DescribeCacheParameterGroups(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheParameterGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_groups") {
			if err := json.Unmarshal([]byte(c.String("cache_parameter_groups")), &in.CacheParameterGroups); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("cache_type") {
			in.CacheType = proto.String(c.String("cache_type"))
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.Int32(int32(c.Int("search_word")))
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeCacheParameterGroups(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCacheParameterGroup = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_type",
		Usage: "cache type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_parameter_group_name",
		Usage: "cache parameter group name",
		Value: "",
	},
}

func _func_CacheService_CreateCacheParameterGroup(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheParameterGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_type") {
			in.CacheType = proto.String(c.String("cache_type"))
		}
		if c.IsSet("cache_parameter_group_name") {
			in.CacheParameterGroupName = proto.String(c.String("cache_parameter_group_name"))
		}
	}

	out, err := qc.CreateCacheParameterGroup(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ApplyCacheParameterGroup = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.StringFlag{
		Name:  "caches",
		Usage: "caches",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_ApplyCacheParameterGroup(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ApplyCacheParameterGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("caches") {
			if err := json.Unmarshal([]byte(c.String("caches")), &in.Caches); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ApplyCacheParameterGroup(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCacheParameterGroups = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_groups",
		Usage: "cache parameter groups",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_DeleteCacheParameterGroups(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCacheParameterGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_groups") {
			if err := json.Unmarshal([]byte(c.String("cache_parameter_groups")), &in.CacheParameterGroups); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteCacheParameterGroups(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheParameterGroupAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_parameter_group_name",
		Usage: "cache parameter group name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_CacheService_ModifyCacheParameterGroupAttributes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheParameterGroupAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("cache_parameter_group_name") {
			in.CacheParameterGroupName = proto.String(c.String("cache_parameter_group_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyCacheParameterGroupAttributes(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
}

func _func_CacheService_DescribeCacheParameters(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
	}

	out, err := qc.DescribeCacheParameters(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_UpdateCacheParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.StringFlag{
		Name:  "parameters",
		Usage: "parameters",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_UpdateCacheParameters(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.UpdateCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("parameters") {
			if err := json.Unmarshal([]byte(c.String("parameters")), &in.Parameters); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.UpdateCacheParameters(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ResetCacheParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "cache_parameter_group",
		Usage: "cache parameter group",
		Value: "",
	},
	cli.StringFlag{
		Name:  "cache_parameter_names",
		Usage: "cache parameter names",
		Value: "", // json: slice/message/map/time
	},
}

func _func_CacheService_ResetCacheParameters(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ResetCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("cache_parameter_group") {
			in.CacheParameterGroup = proto.String(c.String("cache_parameter_group"))
		}
		if c.IsSet("cache_parameter_names") {
			if err := json.Unmarshal([]byte(c.String("cache_parameter_names")), &in.CacheParameterNames); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ResetCacheParameters(in)
	if err != nil {
		log.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
	return nil
}
