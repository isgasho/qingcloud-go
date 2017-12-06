// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: cache.proto

package qcli_pb

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/chai2010/qingcloud-go/pkg/logger"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
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
			Action:  _cmd_CacheService_DescribeCaches,
		},
		{
			Name:    "CreateCache",
			Aliases: []string{},
			Usage:   "CreateCache",
			Flags:   _flag_CacheService_CreateCache,
			Action:  _cmd_CacheService_CreateCache,
		},
		{
			Name:    "StopCaches",
			Aliases: []string{},
			Usage:   "StopCaches",
			Flags:   _flag_CacheService_StopCaches,
			Action:  _cmd_CacheService_StopCaches,
		},
		{
			Name:    "StartCaches",
			Aliases: []string{},
			Usage:   "StartCaches",
			Flags:   _flag_CacheService_StartCaches,
			Action:  _cmd_CacheService_StartCaches,
		},
		{
			Name:    "RestartCaches",
			Aliases: []string{},
			Usage:   "RestartCaches",
			Flags:   _flag_CacheService_RestartCaches,
			Action:  _cmd_CacheService_RestartCaches,
		},
		{
			Name:    "DeleteCaches",
			Aliases: []string{},
			Usage:   "DeleteCaches",
			Flags:   _flag_CacheService_DeleteCaches,
			Action:  _cmd_CacheService_DeleteCaches,
		},
		{
			Name:    "ResizeCaches",
			Aliases: []string{},
			Usage:   "ResizeCaches",
			Flags:   _flag_CacheService_ResizeCaches,
			Action:  _cmd_CacheService_ResizeCaches,
		},
		{
			Name:    "UpdateCache",
			Aliases: []string{},
			Usage:   "UpdateCache",
			Flags:   _flag_CacheService_UpdateCache,
			Action:  _cmd_CacheService_UpdateCache,
		},
		{
			Name:    "ChangeCacheVxnet",
			Aliases: []string{},
			Usage:   "ChangeCacheVxnet",
			Flags:   _flag_CacheService_ChangeCacheVxnet,
			Action:  _cmd_CacheService_ChangeCacheVxnet,
		},
		{
			Name:    "ModifyCacheAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheAttributes",
			Flags:   _flag_CacheService_ModifyCacheAttributes,
			Action:  _cmd_CacheService_ModifyCacheAttributes,
		},
		{
			Name:    "DescribeCacheNodes",
			Aliases: []string{},
			Usage:   "DescribeCacheNodes",
			Flags:   _flag_CacheService_DescribeCacheNodes,
			Action:  _cmd_CacheService_DescribeCacheNodes,
		},
		{
			Name:    "AddCacheNodes",
			Aliases: []string{},
			Usage:   "AddCacheNodes",
			Flags:   _flag_CacheService_AddCacheNodes,
			Action:  _cmd_CacheService_AddCacheNodes,
		},
		{
			Name:    "DeleteCacheNodes",
			Aliases: []string{},
			Usage:   "DeleteCacheNodes",
			Flags:   _flag_CacheService_DeleteCacheNodes,
			Action:  _cmd_CacheService_DeleteCacheNodes,
		},
		{
			Name:    "RestartCacheNodes",
			Aliases: []string{},
			Usage:   "RestartCacheNodes",
			Flags:   _flag_CacheService_RestartCacheNodes,
			Action:  _cmd_CacheService_RestartCacheNodes,
		},
		{
			Name:    "ModifyCacheNodeAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheNodeAttributes",
			Flags:   _flag_CacheService_ModifyCacheNodeAttributes,
			Action:  _cmd_CacheService_ModifyCacheNodeAttributes,
		},
		{
			Name:    "CreateCacheFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateCacheFromSnapshot",
			Flags:   _flag_CacheService_CreateCacheFromSnapshot,
			Action:  _cmd_CacheService_CreateCacheFromSnapshot,
		},
		{
			Name:    "DescribeCacheParameterGroups",
			Aliases: []string{},
			Usage:   "DescribeCacheParameterGroups",
			Flags:   _flag_CacheService_DescribeCacheParameterGroups,
			Action:  _cmd_CacheService_DescribeCacheParameterGroups,
		},
		{
			Name:    "CreateCacheParameterGroup",
			Aliases: []string{},
			Usage:   "CreateCacheParameterGroup",
			Flags:   _flag_CacheService_CreateCacheParameterGroup,
			Action:  _cmd_CacheService_CreateCacheParameterGroup,
		},
		{
			Name:    "ApplyCacheParameterGroup",
			Aliases: []string{},
			Usage:   "ApplyCacheParameterGroup",
			Flags:   _flag_CacheService_ApplyCacheParameterGroup,
			Action:  _cmd_CacheService_ApplyCacheParameterGroup,
		},
		{
			Name:    "DeleteCacheParameterGroups",
			Aliases: []string{},
			Usage:   "DeleteCacheParameterGroups",
			Flags:   _flag_CacheService_DeleteCacheParameterGroups,
			Action:  _cmd_CacheService_DeleteCacheParameterGroups,
		},
		{
			Name:    "ModifyCacheParameterGroupAttributes",
			Aliases: []string{},
			Usage:   "ModifyCacheParameterGroupAttributes",
			Flags:   _flag_CacheService_ModifyCacheParameterGroupAttributes,
			Action:  _cmd_CacheService_ModifyCacheParameterGroupAttributes,
		},
		{
			Name:    "DescribeCacheParameters",
			Aliases: []string{},
			Usage:   "DescribeCacheParameters",
			Flags:   _flag_CacheService_DescribeCacheParameters,
			Action:  _cmd_CacheService_DescribeCacheParameters,
		},
		{
			Name:    "UpdateCacheParameters",
			Aliases: []string{},
			Usage:   "UpdateCacheParameters",
			Flags:   _flag_CacheService_UpdateCacheParameters,
			Action:  _cmd_CacheService_UpdateCacheParameters,
		},
		{
			Name:    "ResetCacheParameters",
			Aliases: []string{},
			Usage:   "ResetCacheParameters",
			Flags:   _flag_CacheService_ResetCacheParameters,
			Action:  _cmd_CacheService_ResetCacheParameters,
		},
	},
}

var _flag_CacheService_DescribeCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DescribeCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCache = []cli.Flag{ /* fields */ }

func _cmd_CacheService_CreateCache(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateCache(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_StopCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_StopCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.StopCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StopCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_StartCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_StartCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.StartCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StartCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_RestartCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_RestartCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.RestartCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RestartCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DeleteCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ResizeCaches = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ResizeCaches(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ResizeCachesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ResizeCaches(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_UpdateCache = []cli.Flag{ /* fields */ }

func _cmd_CacheService_UpdateCache(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.UpdateCacheInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.UpdateCache(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ChangeCacheVxnet = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ChangeCacheVxnet(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ChangeCacheVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ChangeCacheVxnet(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheAttributes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ModifyCacheAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyCacheAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheNodes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DescribeCacheNodes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeCacheNodes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_AddCacheNodes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_AddCacheNodes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.AddCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddCacheNodes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCacheNodes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DeleteCacheNodes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteCacheNodes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_RestartCacheNodes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_RestartCacheNodes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.RestartCacheNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RestartCacheNodes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheNodeAttributes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ModifyCacheNodeAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheNodeAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyCacheNodeAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCacheFromSnapshot = []cli.Flag{ /* fields */ }

func _cmd_CacheService_CreateCacheFromSnapshot(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateCacheFromSnapshot(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheParameterGroups = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DescribeCacheParameterGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheParameterGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeCacheParameterGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_CreateCacheParameterGroup = []cli.Flag{ /* fields */ }

func _cmd_CacheService_CreateCacheParameterGroup(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.CreateCacheParameterGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateCacheParameterGroup(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ApplyCacheParameterGroup = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ApplyCacheParameterGroup(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ApplyCacheParameterGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ApplyCacheParameterGroup(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DeleteCacheParameterGroups = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DeleteCacheParameterGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DeleteCacheParameterGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteCacheParameterGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ModifyCacheParameterGroupAttributes = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ModifyCacheParameterGroupAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ModifyCacheParameterGroupAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyCacheParameterGroupAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_DescribeCacheParameters = []cli.Flag{ /* fields */ }

func _cmd_CacheService_DescribeCacheParameters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.DescribeCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeCacheParameters(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_UpdateCacheParameters = []cli.Flag{ /* fields */ }

func _cmd_CacheService_UpdateCacheParameters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.UpdateCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.UpdateCacheParameters(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_CacheService_ResetCacheParameters = []cli.Flag{ /* fields */ }

func _cmd_CacheService_ResetCacheParameters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewCacheService(conf, zone)

	in := new(pb.ResetCacheParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ResetCacheParameters(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}