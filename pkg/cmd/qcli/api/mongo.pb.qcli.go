// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: mongo.proto

package qcli_pb

import (
	"encoding/json"
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
	_ = json.Marshal
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
			Action:  _func_MongoService_DescribeMongoNodes,
		},
		{
			Name:    "DescribeMongoParameters",
			Aliases: []string{},
			Usage:   "DescribeMongoParameters",
			Flags:   _flag_MongoService_DescribeMongoParameters,
			Action:  _func_MongoService_DescribeMongoParameters,
		},
		{
			Name:    "ResizeMongos",
			Aliases: []string{},
			Usage:   "ResizeMongos",
			Flags:   _flag_MongoService_ResizeMongos,
			Action:  _func_MongoService_ResizeMongos,
		},
		{
			Name:    "CreateMongo",
			Aliases: []string{},
			Usage:   "CreateMongo",
			Flags:   _flag_MongoService_CreateMongo,
			Action:  _func_MongoService_CreateMongo,
		},
		{
			Name:    "StopMongos",
			Aliases: []string{},
			Usage:   "StopMongos",
			Flags:   _flag_MongoService_StopMongos,
			Action:  _func_MongoService_StopMongos,
		},
		{
			Name:    "StartMongos",
			Aliases: []string{},
			Usage:   "StartMongos",
			Flags:   _flag_MongoService_StartMongos,
			Action:  _func_MongoService_StartMongos,
		},
		{
			Name:    "DescribeMongos",
			Aliases: []string{},
			Usage:   "DescribeMongos",
			Flags:   _flag_MongoService_DescribeMongos,
			Action:  _func_MongoService_DescribeMongos,
		},
		{
			Name:    "DeleteMongos",
			Aliases: []string{},
			Usage:   "DeleteMongos",
			Flags:   _flag_MongoService_DeleteMongos,
			Action:  _func_MongoService_DeleteMongos,
		},
		{
			Name:    "CreateMongoFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateMongoFromSnapshot",
			Flags:   _flag_MongoService_CreateMongoFromSnapshot,
			Action:  _func_MongoService_CreateMongoFromSnapshot,
		},
		{
			Name:    "ChangeMongoVxnet",
			Aliases: []string{},
			Usage:   "ChangeMongoVxnet",
			Flags:   _flag_MongoService_ChangeMongoVxnet,
			Action:  _func_MongoService_ChangeMongoVxnet,
		},
		{
			Name:    "AddMongoInstances",
			Aliases: []string{},
			Usage:   "AddMongoInstances",
			Flags:   _flag_MongoService_AddMongoInstances,
			Action:  _func_MongoService_AddMongoInstances,
		},
		{
			Name:    "RemoveMongoInstances",
			Aliases: []string{},
			Usage:   "RemoveMongoInstances",
			Flags:   _flag_MongoService_RemoveMongoInstances,
			Action:  _func_MongoService_RemoveMongoInstances,
		},
		{
			Name:    "ModifyMongoAttributes",
			Aliases: []string{},
			Usage:   "ModifyMongoAttributes",
			Flags:   _flag_MongoService_ModifyMongoAttributes,
			Action:  _func_MongoService_ModifyMongoAttributes,
		},
		{
			Name:    "ModifyMongoInstances",
			Aliases: []string{},
			Usage:   "ModifyMongoInstances",
			Flags:   _flag_MongoService_ModifyMongoInstances,
			Action:  _func_MongoService_ModifyMongoInstances,
		},
	},
}

var _flag_MongoService_DescribeMongoNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: "",
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

func _func_MongoService_DescribeMongoNodes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
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

var _flag_MongoService_DescribeMongoParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: "",
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

func _func_MongoService_DescribeMongoParameters(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
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

var _flag_MongoService_ResizeMongos = []cli.Flag{
	cli.StringFlag{
		Name:  "mongos",
		Usage: "mongos",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "mongo_type",
		Usage: "mongo type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "storage_size",
		Usage: "storage size",
		Value: 0,
	},
}

func _func_MongoService_ResizeMongos(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongos") {
			if err := json.Unmarshal([]byte(c.String("mongos")), &in.Mongos); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("mongo_type") {
			in.MongoType = proto.Int32(int32(c.Int("mongo_type")))
		}
		if c.IsSet("storage_size") {
			in.StorageSize = proto.Int32(int32(c.Int("storage_size")))
		}
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

var _flag_MongoService_CreateMongo = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "mongo_version",
		Usage: "mongo version",
		Value: "",
	},
	cli.IntFlag{
		Name:  "mongo_type",
		Usage: "mongo type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "storage_size",
		Usage: "storage size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "mongo_name",
		Usage: "mongo name",
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
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MongoService_CreateMongo(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("mongo_version") {
			in.MongoVersion = proto.String(c.String("mongo_version"))
		}
		if c.IsSet("mongo_type") {
			in.MongoType = proto.Int32(int32(c.Int("mongo_type")))
		}
		if c.IsSet("storage_size") {
			in.StorageSize = proto.Int32(int32(c.Int("storage_size")))
		}
		if c.IsSet("mongo_name") {
			in.MongoName = proto.String(c.String("mongo_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_StopMongos = []cli.Flag{
	cli.StringFlag{
		Name:  "mongos",
		Usage: "mongos",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MongoService_StopMongos(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongos") {
			if err := json.Unmarshal([]byte(c.String("mongos")), &in.Mongos); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_StartMongos = []cli.Flag{
	cli.StringFlag{
		Name:  "mongos",
		Usage: "mongos",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MongoService_StartMongos(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongos") {
			if err := json.Unmarshal([]byte(c.String("mongos")), &in.Mongos); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_DescribeMongos = []cli.Flag{
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "mongo_name",
		Usage: "mongo name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "mongos",
		Usage: "mongos",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "", // json: slice/message/map/time
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
}

func _func_MongoService_DescribeMongos(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
		if c.IsSet("mongo_name") {
			in.MongoName = proto.String(c.String("mongo_name"))
		}
		if c.IsSet("mongos") {
			if err := json.Unmarshal([]byte(c.String("mongos")), &in.Mongos); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
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

var _flag_MongoService_DeleteMongos = []cli.Flag{
	cli.StringFlag{
		Name:  "mongos",
		Usage: "mongos",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MongoService_DeleteMongos(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongos") {
			if err := json.Unmarshal([]byte(c.String("mongos")), &in.Mongos); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_CreateMongoFromSnapshot = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.IntFlag{
		Name:  "mongo_type",
		Usage: "mongo type",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "mongo_name",
		Usage: "mongo name",
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

func _func_MongoService_CreateMongoFromSnapshot(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("mongo_type") {
			in.MongoType = proto.Int32(int32(c.Int("mongo_type")))
		}
		if c.IsSet("mongo_name") {
			in.MongoName = proto.String(c.String("mongo_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
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

var _flag_MongoService_ChangeMongoVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
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

func _func_MongoService_ChangeMongoVxnet(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_AddMongoInstances = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
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

func _func_MongoService_AddMongoInstances(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_RemoveMongoInstances = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: "",
	},
	cli.StringFlag{
		Name:  "mongo_instances",
		Usage: "mongo instances",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MongoService_RemoveMongoInstances(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("mongo_instances") {
			if err := json.Unmarshal([]byte(c.String("mongo_instances")), &in.MongoInstances); err != nil {
				logger.Fatal(err)
			}
		}
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

var _flag_MongoService_ModifyMongoAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: "",
	},
	cli.StringFlag{
		Name:  "mongo_name",
		Usage: "mongo name",
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

func _func_MongoService_ModifyMongoAttributes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
		if c.IsSet("mongo_name") {
			in.MongoName = proto.String(c.String("mongo_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
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

var _flag_MongoService_ModifyMongoInstances = []cli.Flag{
	cli.StringFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: "",
	},
}

func _func_MongoService_ModifyMongoInstances(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
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
		if c.IsSet("mongo") {
			in.Mongo = proto.String(c.String("mongo"))
		}
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
