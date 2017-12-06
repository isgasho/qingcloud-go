// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: snapshot.proto

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
	AllCommands = append(AllCommands, CmdSnapshotService)
}

var CmdSnapshotService = cli.Command{
	Name:    "snapshot",
	Aliases: []string{},
	Usage:   "manage SnapshotService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeSnapshots",
			Aliases: []string{},
			Usage:   "DescribeSnapshots",
			Flags:   _flag_SnapshotService_DescribeSnapshots,
			Action:  _cmd_SnapshotService_DescribeSnapshots,
		},
		{
			Name:    "CreateSnapshots",
			Aliases: []string{},
			Usage:   "CreateSnapshots",
			Flags:   _flag_SnapshotService_CreateSnapshots,
			Action:  _cmd_SnapshotService_CreateSnapshots,
		},
		{
			Name:    "DeleteSnapshots",
			Aliases: []string{},
			Usage:   "DeleteSnapshots",
			Flags:   _flag_SnapshotService_DeleteSnapshots,
			Action:  _cmd_SnapshotService_DeleteSnapshots,
		},
		{
			Name:    "ApplySnapshots",
			Aliases: []string{},
			Usage:   "ApplySnapshots",
			Flags:   _flag_SnapshotService_ApplySnapshots,
			Action:  _cmd_SnapshotService_ApplySnapshots,
		},
		{
			Name:    "ModifySnapshotAttributes",
			Aliases: []string{},
			Usage:   "ModifySnapshotAttributes",
			Flags:   _flag_SnapshotService_ModifySnapshotAttributes,
			Action:  _cmd_SnapshotService_ModifySnapshotAttributes,
		},
		{
			Name:    "CaptureInstanceFromSnapshot",
			Aliases: []string{},
			Usage:   "CaptureInstanceFromSnapshot",
			Flags:   _flag_SnapshotService_CaptureInstanceFromSnapshot,
			Action:  _cmd_SnapshotService_CaptureInstanceFromSnapshot,
		},
		{
			Name:    "CreateVolumeFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateVolumeFromSnapshot",
			Flags:   _flag_SnapshotService_CreateVolumeFromSnapshot,
			Action:  _cmd_SnapshotService_CreateVolumeFromSnapshot,
		},
	},
}

var _flag_SnapshotService_DescribeSnapshots = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_DescribeSnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.DescribeSnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSnapshots(in)
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

var _flag_SnapshotService_CreateSnapshots = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_CreateSnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.CreateSnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateSnapshots(in)
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

var _flag_SnapshotService_DeleteSnapshots = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_DeleteSnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.DeleteSnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteSnapshots(in)
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

var _flag_SnapshotService_ApplySnapshots = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_ApplySnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.ApplySnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ApplySnapshots(in)
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

var _flag_SnapshotService_ModifySnapshotAttributes = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_ModifySnapshotAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.ModifySnapshotAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifySnapshotAttributes(in)
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

var _flag_SnapshotService_CaptureInstanceFromSnapshot = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_CaptureInstanceFromSnapshot(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.CaptureInstanceFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CaptureInstanceFromSnapshot(in)
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

var _flag_SnapshotService_CreateVolumeFromSnapshot = []cli.Flag{ /* fields */ }

func _cmd_SnapshotService_CreateVolumeFromSnapshot(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSnapshotService(conf, zone)

	in := new(pb.CreateVolumeFromSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateVolumeFromSnapshot(in)
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
