// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: key_pair.proto

package qcli_pb

import (
	"fmt"

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

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdKeyPairService)
}

var CmdKeyPairService = cli.Command{
	Name:    "keypair",
	Aliases: []string{},
	Usage:   "manage KeyPairService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeKeyPairs",
			Aliases: []string{},
			Usage:   "DescribeKeyPairs",
			Flags:   _flag_KeyPairService_DescribeKeyPairs,
			Action:  _cmd_KeyPairService_DescribeKeyPairs,
		},
		{
			Name:    "CreateKeyPair",
			Aliases: []string{},
			Usage:   "CreateKeyPair",
			Flags:   _flag_KeyPairService_CreateKeyPair,
			Action:  _cmd_KeyPairService_CreateKeyPair,
		},
		{
			Name:    "DeleteKeyPairs",
			Aliases: []string{},
			Usage:   "DeleteKeyPairs",
			Flags:   _flag_KeyPairService_DeleteKeyPairs,
			Action:  _cmd_KeyPairService_DeleteKeyPairs,
		},
		{
			Name:    "AttachKeyPairs",
			Aliases: []string{},
			Usage:   "AttachKeyPairs",
			Flags:   _flag_KeyPairService_AttachKeyPairs,
			Action:  _cmd_KeyPairService_AttachKeyPairs,
		},
		{
			Name:    "DetachKeyPairs",
			Aliases: []string{},
			Usage:   "DetachKeyPairs",
			Flags:   _flag_KeyPairService_DetachKeyPairs,
			Action:  _cmd_KeyPairService_DetachKeyPairs,
		},
		{
			Name:    "ModifyKeyPairAttributes",
			Aliases: []string{},
			Usage:   "ModifyKeyPairAttributes",
			Flags:   _flag_KeyPairService_ModifyKeyPairAttributes,
			Action:  _cmd_KeyPairService_ModifyKeyPairAttributes,
		},
	},
}

var _flag_KeyPairService_DescribeKeyPairs = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_DescribeKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DescribeKeyPairsInput)

	// TODO: fill field from flags

	out, err := qc.DescribeKeyPairs(in)
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

var _flag_KeyPairService_CreateKeyPair = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_CreateKeyPair(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.CreateKeyPairInput)

	// TODO: fill field from flags

	out, err := qc.CreateKeyPair(in)
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

var _flag_KeyPairService_DeleteKeyPairs = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_DeleteKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DeleteKeyPairsInput)

	// TODO: fill field from flags

	out, err := qc.DeleteKeyPairs(in)
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

var _flag_KeyPairService_AttachKeyPairs = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_AttachKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.AttachKeyPairsInput)

	// TODO: fill field from flags

	out, err := qc.AttachKeyPairs(in)
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

var _flag_KeyPairService_DetachKeyPairs = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_DetachKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DetachKeyPairsInput)

	// TODO: fill field from flags

	out, err := qc.DetachKeyPairs(in)
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

var _flag_KeyPairService_ModifyKeyPairAttributes = []cli.Flag{ /* fields */ }

func _cmd_KeyPairService_ModifyKeyPairAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.ModifyKeyPairAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyKeyPairAttributes(in)
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
