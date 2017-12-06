// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: mongo.proto

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
	AllCommands = append(AllCommands, CmdMongoService)
}

var CmdMongoService = cli.Command{
	Name:    "mongo",
	Aliases: []string{},
	Usage:   "manage MongoService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeMongoNodes",
			Aliases: []string{},
			Usage:   "DescribeMongoNodes",
			Flags:   _flag_MongoService_DescribeMongoNodes,
			Action:  _cmd_MongoService_DescribeMongoNodes,
		},
		{
			Name:    "DescribeMongoParameters",
			Aliases: []string{},
			Usage:   "DescribeMongoParameters",
			Flags:   _flag_MongoService_DescribeMongoParameters,
			Action:  _cmd_MongoService_DescribeMongoParameters,
		},
		{
			Name:    "ResizeMongos",
			Aliases: []string{},
			Usage:   "ResizeMongos",
			Flags:   _flag_MongoService_ResizeMongos,
			Action:  _cmd_MongoService_ResizeMongos,
		},
		{
			Name:    "CreateMongo",
			Aliases: []string{},
			Usage:   "CreateMongo",
			Flags:   _flag_MongoService_CreateMongo,
			Action:  _cmd_MongoService_CreateMongo,
		},
		{
			Name:    "StopMongos",
			Aliases: []string{},
			Usage:   "StopMongos",
			Flags:   _flag_MongoService_StopMongos,
			Action:  _cmd_MongoService_StopMongos,
		},
		{
			Name:    "StartMongos",
			Aliases: []string{},
			Usage:   "StartMongos",
			Flags:   _flag_MongoService_StartMongos,
			Action:  _cmd_MongoService_StartMongos,
		},
		{
			Name:    "DescribeMongos",
			Aliases: []string{},
			Usage:   "DescribeMongos",
			Flags:   _flag_MongoService_DescribeMongos,
			Action:  _cmd_MongoService_DescribeMongos,
		},
		{
			Name:    "DeleteMongos",
			Aliases: []string{},
			Usage:   "DeleteMongos",
			Flags:   _flag_MongoService_DeleteMongos,
			Action:  _cmd_MongoService_DeleteMongos,
		},
		{
			Name:    "CreateMongoFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateMongoFromSnapshot",
			Flags:   _flag_MongoService_CreateMongoFromSnapshot,
			Action:  _cmd_MongoService_CreateMongoFromSnapshot,
		},
		{
			Name:    "ChangeMongoVxnet",
			Aliases: []string{},
			Usage:   "ChangeMongoVxnet",
			Flags:   _flag_MongoService_ChangeMongoVxnet,
			Action:  _cmd_MongoService_ChangeMongoVxnet,
		},
		{
			Name:    "AddMongoInstances",
			Aliases: []string{},
			Usage:   "AddMongoInstances",
			Flags:   _flag_MongoService_AddMongoInstances,
			Action:  _cmd_MongoService_AddMongoInstances,
		},
		{
			Name:    "RemoveMongoInstances",
			Aliases: []string{},
			Usage:   "RemoveMongoInstances",
			Flags:   _flag_MongoService_RemoveMongoInstances,
			Action:  _cmd_MongoService_RemoveMongoInstances,
		},
		{
			Name:    "ModifyMongoAttributes",
			Aliases: []string{},
			Usage:   "ModifyMongoAttributes",
			Flags:   _flag_MongoService_ModifyMongoAttributes,
			Action:  _cmd_MongoService_ModifyMongoAttributes,
		},
		{
			Name:    "ModifyMongoInstances",
			Aliases: []string{},
			Usage:   "ModifyMongoInstances",
			Flags:   _flag_MongoService_ModifyMongoInstances,
			Action:  _cmd_MongoService_ModifyMongoInstances,
		},
	},
}

var _flag_MongoService_DescribeMongoNodes = []cli.Flag{ /* fields */ }

func _cmd_MongoService_DescribeMongoNodes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.DescribeMongoNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeMongoNodes(in)
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

var _flag_MongoService_DescribeMongoParameters = []cli.Flag{ /* fields */ }

func _cmd_MongoService_DescribeMongoParameters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.DescribeMongoParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeMongoParameters(in)
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

var _flag_MongoService_ResizeMongos = []cli.Flag{ /* fields */ }

func _cmd_MongoService_ResizeMongos(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.ResizeMongosInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ResizeMongos(in)
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

var _flag_MongoService_CreateMongo = []cli.Flag{ /* fields */ }

func _cmd_MongoService_CreateMongo(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.CreateMongoInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateMongo(in)
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

var _flag_MongoService_StopMongos = []cli.Flag{ /* fields */ }

func _cmd_MongoService_StopMongos(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.StopMongosInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StopMongos(in)
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

var _flag_MongoService_StartMongos = []cli.Flag{ /* fields */ }

func _cmd_MongoService_StartMongos(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.StartMongosInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StartMongos(in)
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

var _flag_MongoService_DescribeMongos = []cli.Flag{ /* fields */ }

func _cmd_MongoService_DescribeMongos(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.DescribeMongosInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeMongos(in)
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

var _flag_MongoService_DeleteMongos = []cli.Flag{ /* fields */ }

func _cmd_MongoService_DeleteMongos(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.DeleteMongosInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteMongos(in)
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

var _flag_MongoService_CreateMongoFromSnapshot = []cli.Flag{ /* fields */ }

func _cmd_MongoService_CreateMongoFromSnapshot(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.CreateMongoFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateMongoFromSnapshot(in)
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

var _flag_MongoService_ChangeMongoVxnet = []cli.Flag{ /* fields */ }

func _cmd_MongoService_ChangeMongoVxnet(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.ChangeMongoVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ChangeMongoVxnet(in)
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

var _flag_MongoService_AddMongoInstances = []cli.Flag{ /* fields */ }

func _cmd_MongoService_AddMongoInstances(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.AddMongoInstancesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddMongoInstances(in)
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

var _flag_MongoService_RemoveMongoInstances = []cli.Flag{ /* fields */ }

func _cmd_MongoService_RemoveMongoInstances(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.RemoveMongoInstancesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RemoveMongoInstances(in)
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

var _flag_MongoService_ModifyMongoAttributes = []cli.Flag{ /* fields */ }

func _cmd_MongoService_ModifyMongoAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.ModifyMongoAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyMongoAttributes(in)
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

var _flag_MongoService_ModifyMongoInstances = []cli.Flag{ /* fields */ }

func _cmd_MongoService_ModifyMongoInstances(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMongoService(conf, zone)

	in := new(pb.ModifyMongoInstancesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyMongoInstances(in)
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