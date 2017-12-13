// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: rdb.proto

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

	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdRDBService)
}

var CmdRDBService = cli.Command{
	Name:     "rdb",
	Aliases:  []string{},
	Usage:    "RDBService",
	Category: "SDK API Style Command",
	Subcommands: []cli.Command{
		{
			Name:    "CreateRDB",
			Aliases: []string{},
			Usage:   "CreateRDB",
			Flags:   _flag_RDBService_CreateRDB,
			Action:  _func_RDBService_CreateRDB,
		},
		{
			Name:    "DescribeRDBs",
			Aliases: []string{},
			Usage:   "DescribeRDBs",
			Flags:   _flag_RDBService_DescribeRDBs,
			Action:  _func_RDBService_DescribeRDBs,
		},
		{
			Name:    "DeleteRDBs",
			Aliases: []string{},
			Usage:   "DeleteRDBs",
			Flags:   _flag_RDBService_DeleteRDBs,
			Action:  _func_RDBService_DeleteRDBs,
		},
		{
			Name:    "StartRDBs",
			Aliases: []string{},
			Usage:   "StartRDBs",
			Flags:   _flag_RDBService_StartRDBs,
			Action:  _func_RDBService_StartRDBs,
		},
		{
			Name:    "StopRDBs",
			Aliases: []string{},
			Usage:   "StopRDBs",
			Flags:   _flag_RDBService_StopRDBs,
			Action:  _func_RDBService_StopRDBs,
		},
		{
			Name:    "ResizeRDBs",
			Aliases: []string{},
			Usage:   "ResizeRDBs",
			Flags:   _flag_RDBService_ResizeRDBs,
			Action:  _func_RDBService_ResizeRDBs,
		},
		{
			Name:    "RDBsLeaveVxnet",
			Aliases: []string{},
			Usage:   "RDBsLeaveVxnet",
			Flags:   _flag_RDBService_RDBsLeaveVxnet,
			Action:  _func_RDBService_RDBsLeaveVxnet,
		},
		{
			Name:    "RDBsJoinVxnet",
			Aliases: []string{},
			Usage:   "RDBsJoinVxnet",
			Flags:   _flag_RDBService_RDBsJoinVxnet,
			Action:  _func_RDBService_RDBsJoinVxnet,
		},
		{
			Name:    "CreateRDBFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateRDBFromSnapshot",
			Flags:   _flag_RDBService_CreateRDBFromSnapshot,
			Action:  _func_RDBService_CreateRDBFromSnapshot,
		},
		{
			Name:    "CreateTempRDBInstanceFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateTempRDBInstanceFromSnapshot",
			Flags:   _flag_RDBService_CreateTempRDBInstanceFromSnapshot,
			Action:  _func_RDBService_CreateTempRDBInstanceFromSnapshot,
		},
		{
			Name:    "GetRDBInstanceFiles",
			Aliases: []string{},
			Usage:   "GetRDBInstanceFiles",
			Flags:   _flag_RDBService_GetRDBInstanceFiles,
			Action:  _func_RDBService_GetRDBInstanceFiles,
		},
		{
			Name:    "CopyRDBInstanceFilesToFTP",
			Aliases: []string{},
			Usage:   "CopyRDBInstanceFilesToFTP",
			Flags:   _flag_RDBService_CopyRDBInstanceFilesToFTP,
			Action:  _func_RDBService_CopyRDBInstanceFilesToFTP,
		},
		{
			Name:    "PurgeRDBLogs",
			Aliases: []string{},
			Usage:   "PurgeRDBLogs",
			Flags:   _flag_RDBService_PurgeRDBLogs,
			Action:  _func_RDBService_PurgeRDBLogs,
		},
		{
			Name:    "CeaseRDBInstance",
			Aliases: []string{},
			Usage:   "CeaseRDBInstance",
			Flags:   _flag_RDBService_CeaseRDBInstance,
			Action:  _func_RDBService_CeaseRDBInstance,
		},
		{
			Name:    "ModifyRDBParameters",
			Aliases: []string{},
			Usage:   "ModifyRDBParameters",
			Flags:   _flag_RDBService_ModifyRDBParameters,
			Action:  _func_RDBService_ModifyRDBParameters,
		},
		{
			Name:    "ApplyRDBParameterGroup",
			Aliases: []string{},
			Usage:   "ApplyRDBParameterGroup",
			Flags:   _flag_RDBService_ApplyRDBParameterGroup,
			Action:  _func_RDBService_ApplyRDBParameterGroup,
		},
		{
			Name:    "DescribeRDBParameters",
			Aliases: []string{},
			Usage:   "DescribeRDBParameters",
			Flags:   _flag_RDBService_DescribeRDBParameters,
			Action:  _func_RDBService_DescribeRDBParameters,
		},
	},
}

var _flag_RDBService_CreateRDB = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_engine",
		Usage: "rdb engine",
		Value: "",
	},
	cli.StringFlag{
		Name:  "engine_version",
		Usage: "engine version",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_username",
		Usage: "rdb username",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_password",
		Usage: "rdb password",
		Value: "",
	},
	cli.IntFlag{
		Name:  "rdb_type",
		Usage: "rdb type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "storage_size",
		Usage: "storage size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "rdb_name",
		Usage: "rdb name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
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

func _func_RDBService_CreateRDB(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.CreateRDBInput)

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
		if c.IsSet("rdb_engine") {
			in.RdbEngine = proto.String(c.String("rdb_engine"))
		}
		if c.IsSet("engine_version") {
			in.EngineVersion = proto.String(c.String("engine_version"))
		}
		if c.IsSet("rdb_username") {
			in.RdbUsername = proto.String(c.String("rdb_username"))
		}
		if c.IsSet("rdb_password") {
			in.RdbPassword = proto.String(c.String("rdb_password"))
		}
		if c.IsSet("rdb_type") {
			in.RdbType = proto.Int32(int32(c.Int("rdb_type")))
		}
		if c.IsSet("storage_size") {
			in.StorageSize = proto.Int32(int32(c.Int("storage_size")))
		}
		if c.IsSet("rdb_name") {
			in.RdbName = proto.String(c.String("rdb_name"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
	}

	out, err := qc.CreateRDB(in)
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

var _flag_RDBService_DescribeRDBs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "rdb_engine",
		Usage: "rdb engine",
		Value: "",
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "rdb_name",
		Usage: "rdb name",
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

func _func_RDBService_DescribeRDBs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.DescribeRDBsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("rdb_engine") {
			in.RdbEngine = proto.String(c.String("rdb_engine"))
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("rdb_name") {
			in.RdbName = proto.String(c.String("rdb_name"))
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

	out, err := qc.DescribeRDBs(in)
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

var _flag_RDBService_DeleteRDBs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_RDBService_DeleteRDBs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.DeleteRDBsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteRDBs(in)
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

var _flag_RDBService_StartRDBs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_RDBService_StartRDBs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.StartRDBsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StartRDBs(in)
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

var _flag_RDBService_StopRDBs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_RDBService_StopRDBs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.StopRDBsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StopRDBs(in)
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

var _flag_RDBService_ResizeRDBs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "rdb_type",
		Usage: "rdb type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "storage_size",
		Usage: "storage size",
		Value: 0,
	},
}

func _func_RDBService_ResizeRDBs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.ResizeRDBsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("rdb_type") {
			in.RdbType = proto.Int32(int32(c.Int("rdb_type")))
		}
		if c.IsSet("storage_size") {
			in.StorageSize = proto.Int32(int32(c.Int("storage_size")))
		}
	}

	out, err := qc.ResizeRDBs(in)
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

var _flag_RDBService_RDBsLeaveVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
}

func _func_RDBService_RDBsLeaveVxnet(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.RDBsLeaveVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
	}

	out, err := qc.RDBsLeaveVxnet(in)
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

var _flag_RDBService_RDBsJoinVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "rdbs",
		Usage: "rdbs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
}

func _func_RDBService_RDBsJoinVxnet(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.RDBsJoinVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdbs") {
			if err := json.Unmarshal([]byte(c.String("rdbs")), &in.Rdbs); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
	}

	out, err := qc.RDBsJoinVxnet(in)
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

var _flag_RDBService_CreateRDBFromSnapshot = []cli.Flag{
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
		Name:  "rdb_type",
		Usage: "rdb type",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "rdb_username",
		Usage: "rdb username",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_password",
		Usage: "rdb password",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_name",
		Usage: "rdb name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
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

func _func_RDBService_CreateRDBFromSnapshot(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.CreateRDBFromSnapshotInput)

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
		if c.IsSet("rdb_type") {
			in.RdbType = proto.Int32(int32(c.Int("rdb_type")))
		}
		if c.IsSet("rdb_username") {
			in.RdbUsername = proto.String(c.String("rdb_username"))
		}
		if c.IsSet("rdb_password") {
			in.RdbPassword = proto.String(c.String("rdb_password"))
		}
		if c.IsSet("rdb_name") {
			in.RdbName = proto.String(c.String("rdb_name"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("auto_backup_time") {
			in.AutoBackupTime = proto.Int32(int32(c.Int("auto_backup_time")))
		}
	}

	out, err := qc.CreateRDBFromSnapshot(in)
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

var _flag_RDBService_CreateTempRDBInstanceFromSnapshot = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: "",
	},
	cli.StringFlag{
		Name:  "snapshot",
		Usage: "snapshot",
		Value: "",
	},
}

func _func_RDBService_CreateTempRDBInstanceFromSnapshot(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.CreateTempRDBInstanceFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
		if c.IsSet("snapshot") {
			in.Snapshot = proto.String(c.String("snapshot"))
		}
	}

	out, err := qc.CreateTempRDBInstanceFromSnapshot(in)
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

var _flag_RDBService_GetRDBInstanceFiles = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb_instance",
		Usage: "rdb instance",
		Value: "",
	},
}

func _func_RDBService_GetRDBInstanceFiles(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.GetRDBInstanceFilesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb_instance") {
			in.RdbInstance = proto.String(c.String("rdb_instance"))
		}
	}

	out, err := qc.GetRDBInstanceFiles(in)
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

var _flag_RDBService_CopyRDBInstanceFilesToFTP = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb_instance",
		Usage: "rdb instance",
		Value: "",
	},
	cli.StringFlag{
		Name:  "files",
		Usage: "files",
		Value: "", // json: slice/message/map/time
	},
}

func _func_RDBService_CopyRDBInstanceFilesToFTP(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.CopyRDBInstanceFilesToFTPInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb_instance") {
			in.RdbInstance = proto.String(c.String("rdb_instance"))
		}
		if c.IsSet("files") {
			if err := json.Unmarshal([]byte(c.String("files")), &in.Files); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.CopyRDBInstanceFilesToFTP(in)
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

var _flag_RDBService_PurgeRDBLogs = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_instance",
		Usage: "rdb instance",
		Value: "",
	},
	cli.StringFlag{
		Name:  "log_type",
		Usage: "log type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "before_file",
		Usage: "before file",
		Value: "",
	},
}

func _func_RDBService_PurgeRDBLogs(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.PurgeRDBLogsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
		if c.IsSet("rdb_instance") {
			in.RdbInstance = proto.String(c.String("rdb_instance"))
		}
		if c.IsSet("log_type") {
			in.LogType = proto.String(c.String("log_type"))
		}
		if c.IsSet("before_file") {
			in.BeforeFile = proto.String(c.String("before_file"))
		}
	}

	out, err := qc.PurgeRDBLogs(in)
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

var _flag_RDBService_CeaseRDBInstance = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rdb_instance",
		Usage: "rdb instance",
		Value: "",
	},
}

func _func_RDBService_CeaseRDBInstance(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.CeaseRDBInstanceInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
		if c.IsSet("rdb_instance") {
			in.RdbInstance = proto.String(c.String("rdb_instance"))
		}
	}

	out, err := qc.CeaseRDBInstance(in)
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

var _flag_RDBService_ModifyRDBParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: "",
	},
	cli.StringFlag{
		Name:  "parameters",
		Usage: "parameters",
		Value: "", // json: slice/message/map/time
	},
}

func _func_RDBService_ModifyRDBParameters(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.ModifyRDBParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
		if c.IsSet("parameters") {
			if err := json.Unmarshal([]byte(c.String("parameters")), &in.Parameters); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ModifyRDBParameters(in)
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

var _flag_RDBService_ApplyRDBParameterGroup = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: "",
	},
}

func _func_RDBService_ApplyRDBParameterGroup(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.ApplyRDBParameterGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
	}

	out, err := qc.ApplyRDBParameterGroup(in)
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

var _flag_RDBService_DescribeRDBParameters = []cli.Flag{
	cli.StringFlag{
		Name:  "rdb",
		Usage: "rdb",
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

func _func_RDBService_DescribeRDBParameters(c *cli.Context) error {
	qc := pb.NewRDBService(pkgGetServerInfo(c))
	in := new(pb.DescribeRDBParametersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("rdb") {
			in.Rdb = proto.String(c.String("rdb"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeRDBParameters(in)
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
