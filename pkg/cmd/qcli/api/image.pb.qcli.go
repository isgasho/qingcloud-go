// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: image.proto

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
			Action:  _func_ImageService_DescribeImages,
		},
		{
			Name:    "CaptureInstance",
			Aliases: []string{},
			Usage:   "CaptureInstance",
			Flags:   _flag_ImageService_CaptureInstance,
			Action:  _func_ImageService_CaptureInstance,
		},
		{
			Name:    "DeleteImages",
			Aliases: []string{},
			Usage:   "DeleteImages",
			Flags:   _flag_ImageService_DeleteImages,
			Action:  _func_ImageService_DeleteImages,
		},
		{
			Name:    "ModifyImageAttributes",
			Aliases: []string{},
			Usage:   "ModifyImageAttributes",
			Flags:   _flag_ImageService_ModifyImageAttributes,
			Action:  _func_ImageService_ModifyImageAttributes,
		},
		{
			Name:    "GrantImageToUsers",
			Aliases: []string{},
			Usage:   "GrantImageToUsers",
			Flags:   _flag_ImageService_GrantImageToUsers,
			Action:  _func_ImageService_GrantImageToUsers,
		},
		{
			Name:    "RevokeImageFromUsers",
			Aliases: []string{},
			Usage:   "RevokeImageFromUsers",
			Flags:   _flag_ImageService_RevokeImageFromUsers,
			Action:  _func_ImageService_RevokeImageFromUsers,
		},
		{
			Name:    "DescribeImageUsers",
			Aliases: []string{},
			Usage:   "DescribeImageUsers",
			Flags:   _flag_ImageService_DescribeImageUsers,
			Action:  _func_ImageService_DescribeImageUsers,
		},
		{
			Name:    "CloneImages",
			Aliases: []string{},
			Usage:   "CloneImages",
			Flags:   _flag_ImageService_CloneImages,
			Action:  _func_ImageService_CloneImages,
		},
	},
}

var _flag_ImageService_DescribeImages = []cli.Flag{
	cli.StringFlag{
		Name:  "images",
		Usage: "images",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "processor_type",
		Usage: "processor type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "os_family",
		Usage: "os family",
		Value: "",
	},
	cli.StringFlag{
		Name:  "visibility",
		Usage: "visibility",
		Value: "",
	},
	cli.StringFlag{
		Name:  "provider",
		Usage: "provider",
		Value: "",
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

func _func_ImageService_DescribeImages(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("images") {
			if err := json.Unmarshal([]byte(c.String("images")), &in.Images); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("processor_type") {
			in.ProcessorType = proto.String(c.String("processor_type"))
		}
		if c.IsSet("os_family") {
			in.OsFamily = proto.String(c.String("os_family"))
		}
		if c.IsSet("visibility") {
			in.Visibility = proto.String(c.String("visibility"))
		}
		if c.IsSet("provider") {
			in.Provider = proto.String(c.String("provider"))
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

	out, err := qc.DescribeImages(in)
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

var _flag_ImageService_CaptureInstance = []cli.Flag{
	cli.StringFlag{
		Name:  "instance",
		Usage: "instance",
		Value: "",
	},
	cli.StringFlag{
		Name:  "image_name",
		Usage: "image name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "target_user",
		Usage: "target user",
		Value: "",
	},
}

func _func_ImageService_CaptureInstance(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CaptureInstanceInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("instance") {
			in.Instance = proto.String(c.String("instance"))
		}
		if c.IsSet("image_name") {
			in.ImageName = proto.String(c.String("image_name"))
		}
		if c.IsSet("target_user") {
			in.TargetUser = proto.String(c.String("target_user"))
		}
	}

	out, err := qc.CaptureInstance(in)
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

var _flag_ImageService_DeleteImages = []cli.Flag{
	cli.StringFlag{
		Name:  "images",
		Usage: "images",
		Value: "", // json: slice/message/map/time
	},
}

func _func_ImageService_DeleteImages(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DeleteImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("images") {
			if err := json.Unmarshal([]byte(c.String("images")), &in.Images); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteImages(in)
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

var _flag_ImageService_ModifyImageAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "image",
		Usage: "image",
		Value: "",
	},
	cli.StringFlag{
		Name:  "image_name",
		Usage: "image name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_ImageService_ModifyImageAttributes(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.ModifyImageAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("image") {
			in.Image = proto.String(c.String("image"))
		}
		if c.IsSet("image_name") {
			in.ImageName = proto.String(c.String("image_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyImageAttributes(in)
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

var _flag_ImageService_GrantImageToUsers = []cli.Flag{
	cli.StringFlag{
		Name:  "image",
		Usage: "image",
		Value: "",
	},
	cli.StringFlag{
		Name:  "users",
		Usage: "users",
		Value: "", // json: slice/message/map/time
	},
}

func _func_ImageService_GrantImageToUsers(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.GrantImageToUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("image") {
			in.Image = proto.String(c.String("image"))
		}
		if c.IsSet("users") {
			if err := json.Unmarshal([]byte(c.String("users")), &in.Users); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.GrantImageToUsers(in)
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

var _flag_ImageService_RevokeImageFromUsers = []cli.Flag{
	cli.StringFlag{
		Name:  "image",
		Usage: "image",
		Value: "",
	},
	cli.StringFlag{
		Name:  "users",
		Usage: "users",
		Value: "", // json: slice/message/map/time
	},
}

func _func_ImageService_RevokeImageFromUsers(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.RevokeImageFromUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("image") {
			in.Image = proto.String(c.String("image"))
		}
		if c.IsSet("users") {
			if err := json.Unmarshal([]byte(c.String("users")), &in.Users); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.RevokeImageFromUsers(in)
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

var _flag_ImageService_DescribeImageUsers = []cli.Flag{
	cli.StringFlag{
		Name:  "image_id",
		Usage: "image id",
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

func _func_ImageService_DescribeImageUsers(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImageUsersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("image_id") {
			in.ImageId = proto.String(c.String("image_id"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeImageUsers(in)
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

var _flag_ImageService_CloneImages = []cli.Flag{
	cli.StringFlag{
		Name:  "image",
		Usage: "image",
		Value: "",
	},
	cli.IntFlag{
		Name:  "count",
		Usage: "count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "image_name",
		Usage: "image name",
		Value: "",
	},
}

func _func_ImageService_CloneImages(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CloneImagesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("image") {
			in.Image = proto.String(c.String("image"))
		}
		if c.IsSet("count") {
			in.Count = proto.Int32(int32(c.Int("count")))
		}
		if c.IsSet("image_name") {
			in.ImageName = proto.String(c.String("image_name"))
		}
	}

	out, err := qc.CloneImages(in)
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
