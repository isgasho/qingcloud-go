// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: image.proto

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
	AllCommands = append(AllCommands, CmdImageService)
}

var CmdImageService = cli.Command{
	Name:    "image",
	Aliases: []string{},
	Usage:   "manage ImageService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeImages",
			Aliases: []string{},
			Usage:   "DescribeImages",
			Flags:   _flag_ImageService_DescribeImages,
			Action:  _cmd_ImageService_DescribeImages,
		},
		{
			Name:    "CaptureInstance",
			Aliases: []string{},
			Usage:   "CaptureInstance",
			Flags:   _flag_ImageService_CaptureInstance,
			Action:  _cmd_ImageService_CaptureInstance,
		},
		{
			Name:    "DeleteImages",
			Aliases: []string{},
			Usage:   "DeleteImages",
			Flags:   _flag_ImageService_DeleteImages,
			Action:  _cmd_ImageService_DeleteImages,
		},
		{
			Name:    "ModifyImageAttributes",
			Aliases: []string{},
			Usage:   "ModifyImageAttributes",
			Flags:   _flag_ImageService_ModifyImageAttributes,
			Action:  _cmd_ImageService_ModifyImageAttributes,
		},
		{
			Name:    "GrantImageToUsers",
			Aliases: []string{},
			Usage:   "GrantImageToUsers",
			Flags:   _flag_ImageService_GrantImageToUsers,
			Action:  _cmd_ImageService_GrantImageToUsers,
		},
		{
			Name:    "RevokeImageFromUsers",
			Aliases: []string{},
			Usage:   "RevokeImageFromUsers",
			Flags:   _flag_ImageService_RevokeImageFromUsers,
			Action:  _cmd_ImageService_RevokeImageFromUsers,
		},
		{
			Name:    "DescribeImageUsers",
			Aliases: []string{},
			Usage:   "DescribeImageUsers",
			Flags:   _flag_ImageService_DescribeImageUsers,
			Action:  _cmd_ImageService_DescribeImageUsers,
		},
		{
			Name:    "CloneImages",
			Aliases: []string{},
			Usage:   "CloneImages",
			Flags:   _flag_ImageService_CloneImages,
			Action:  _cmd_ImageService_CloneImages,
		},
	},
}

var _flag_ImageService_DescribeImages = []cli.Flag{ /* fields */ }

func _cmd_ImageService_DescribeImages(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeImages(in)
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

var _flag_ImageService_CaptureInstance = []cli.Flag{ /* fields */ }

func _cmd_ImageService_CaptureInstance(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CaptureInstanceInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CaptureInstance(in)
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

var _flag_ImageService_DeleteImages = []cli.Flag{ /* fields */ }

func _cmd_ImageService_DeleteImages(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DeleteImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteImages(in)
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

var _flag_ImageService_ModifyImageAttributes = []cli.Flag{ /* fields */ }

func _cmd_ImageService_ModifyImageAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.ModifyImageAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyImageAttributes(in)
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

var _flag_ImageService_GrantImageToUsers = []cli.Flag{ /* fields */ }

func _cmd_ImageService_GrantImageToUsers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.GrantImageToUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.GrantImageToUsers(in)
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

var _flag_ImageService_RevokeImageFromUsers = []cli.Flag{ /* fields */ }

func _cmd_ImageService_RevokeImageFromUsers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.RevokeImageFromUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RevokeImageFromUsers(in)
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

var _flag_ImageService_DescribeImageUsers = []cli.Flag{ /* fields */ }

func _cmd_ImageService_DescribeImageUsers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImageUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeImageUsers(in)
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

var _flag_ImageService_CloneImages = []cli.Flag{ /* fields */ }

func _cmd_ImageService_CloneImages(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CloneImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CloneImages(in)
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