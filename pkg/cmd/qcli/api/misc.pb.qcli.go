// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: misc.proto

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
	AllCommands = append(AllCommands, CmdMiscService)
}

var CmdMiscService = cli.Command{
	Name:    "misc",
	Aliases: []string{},
	Usage:   "manage MiscService",
	Subcommands: []cli.Command{
		{
			Name:    "GrantQuotaIndep",
			Aliases: []string{},
			Usage:   "GrantQuotaIndep",
			Flags:   _flag_MiscService_GrantQuotaIndep,
			Action:  _cmd_MiscService_GrantQuotaIndep,
		},
		{
			Name:    "RevokeQuotaIndep",
			Aliases: []string{},
			Usage:   "RevokeQuotaIndep",
			Flags:   _flag_MiscService_RevokeQuotaIndep,
			Action:  _cmd_MiscService_RevokeQuotaIndep,
		},
		{
			Name:    "GetQuotaLeft",
			Aliases: []string{},
			Usage:   "GetQuotaLeft",
			Flags:   _flag_MiscService_GetQuotaLeft,
			Action:  _cmd_MiscService_GetQuotaLeft,
		},
	},
}

var _flag_MiscService_GrantQuotaIndep = []cli.Flag{ /* fields */ }

func _cmd_MiscService_GrantQuotaIndep(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService(conf, zone)

	in := new(pb.GrantQuotaIndepInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.GrantQuotaIndep(in)
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

var _flag_MiscService_RevokeQuotaIndep = []cli.Flag{ /* fields */ }

func _cmd_MiscService_RevokeQuotaIndep(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService(conf, zone)

	in := new(pb.RevokeQuotaIndepInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RevokeQuotaIndep(in)
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

var _flag_MiscService_GetQuotaLeft = []cli.Flag{ /* fields */ }

func _cmd_MiscService_GetQuotaLeft(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService(conf, zone)

	in := new(pb.GetQuotaLeftInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.GetQuotaLeft(in)
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