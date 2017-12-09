// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: volume.proto

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
	AllCommands = append(AllCommands, CmdVolumesService)
}

var CmdVolumesService = cli.Command{
	Name:    "volumes",
	Aliases: []string{},
	Usage:   "manage VolumesService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeVolumes",
			Aliases: []string{},
			Usage:   "DescribeVolumes",
			Flags:   _flag_VolumesService_DescribeVolumes,
			Action:  _func_VolumesService_DescribeVolumes,
		},
		{
			Name:    "CreateVolumes",
			Aliases: []string{},
			Usage:   "CreateVolumes",
			Flags:   _flag_VolumesService_CreateVolumes,
			Action:  _func_VolumesService_CreateVolumes,
		},
		{
			Name:    "DeleteVolumes",
			Aliases: []string{},
			Usage:   "DeleteVolumes",
			Flags:   _flag_VolumesService_DeleteVolumes,
			Action:  _func_VolumesService_DeleteVolumes,
		},
		{
			Name:    "AttachVolumes",
			Aliases: []string{},
			Usage:   "AttachVolumes",
			Flags:   _flag_VolumesService_AttachVolumes,
			Action:  _func_VolumesService_AttachVolumes,
		},
		{
			Name:    "DetachVolumes",
			Aliases: []string{},
			Usage:   "DetachVolumes",
			Flags:   _flag_VolumesService_DetachVolumes,
			Action:  _func_VolumesService_DetachVolumes,
		},
		{
			Name:    "ResizeVolumes",
			Aliases: []string{},
			Usage:   "ResizeVolumes",
			Flags:   _flag_VolumesService_ResizeVolumes,
			Action:  _func_VolumesService_ResizeVolumes,
		},
		{
			Name:    "ModifyVolumeAttributes",
			Aliases: []string{},
			Usage:   "ModifyVolumeAttributes",
			Flags:   _flag_VolumesService_ModifyVolumeAttributes,
			Action:  _func_VolumesService_ModifyVolumeAttributes,
		},
	},
}

var _flag_VolumesService_DescribeVolumes = []cli.Flag{
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "search_word",
		Usage: "search word",
		Value: "",
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
	cli.IntFlag{
		Name:  "volume_type",
		Usage: "volume type",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "volumes",
		Usage: "volumes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VolumesService_DescribeVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.DescribeVolumesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.String(c.String("search_word"))
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
		if c.IsSet("volume_type") {
			in.VolumeType = proto.Int32(int32(c.Int("volume_type")))
		}
		if c.IsSet("volumes") {
			if err := json.Unmarshal([]byte(c.String("volumes")), &in.Volumes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DescribeVolumes(in)
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

var _flag_VolumesService_CreateVolumes = []cli.Flag{
	cli.IntFlag{
		Name:  "count",
		Usage: "count",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "size",
		Usage: "size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "volume_name",
		Usage: "volume name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "volume_type",
		Usage: "volume type",
		Value: 0,
	},
}

func _func_VolumesService_CreateVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.CreateVolumesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("count") {
			in.Count = proto.Int32(int32(c.Int("count")))
		}
		if c.IsSet("size") {
			in.Size = proto.Int32(int32(c.Int("size")))
		}
		if c.IsSet("volume_name") {
			in.VolumeName = proto.String(c.String("volume_name"))
		}
		if c.IsSet("volume_type") {
			in.VolumeType = proto.Int32(int32(c.Int("volume_type")))
		}
	}

	out, err := qc.CreateVolumes(in)
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

var _flag_VolumesService_DeleteVolumes = []cli.Flag{
	cli.StringFlag{
		Name:  "volumes",
		Usage: "volumes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VolumesService_DeleteVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.DeleteVolumesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("volumes") {
			if err := json.Unmarshal([]byte(c.String("volumes")), &in.Volumes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteVolumes(in)
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

var _flag_VolumesService_AttachVolumes = []cli.Flag{
	cli.StringFlag{
		Name:  "instance",
		Usage: "instance",
		Value: "",
	},
	cli.StringFlag{
		Name:  "volumes",
		Usage: "volumes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VolumesService_AttachVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.AttachVolumesInput)

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
		if c.IsSet("volumes") {
			if err := json.Unmarshal([]byte(c.String("volumes")), &in.Volumes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AttachVolumes(in)
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

var _flag_VolumesService_DetachVolumes = []cli.Flag{
	cli.StringFlag{
		Name:  "instance",
		Usage: "instance",
		Value: "",
	},
	cli.StringFlag{
		Name:  "volumes",
		Usage: "volumes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VolumesService_DetachVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.DetachVolumesInput)

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
		if c.IsSet("volumes") {
			if err := json.Unmarshal([]byte(c.String("volumes")), &in.Volumes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DetachVolumes(in)
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

var _flag_VolumesService_ResizeVolumes = []cli.Flag{
	cli.IntFlag{
		Name:  "size",
		Usage: "size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "volumes",
		Usage: "volumes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VolumesService_ResizeVolumes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.ResizeVolumesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("size") {
			in.Size = proto.Int32(int32(c.Int("size")))
		}
		if c.IsSet("volumes") {
			if err := json.Unmarshal([]byte(c.String("volumes")), &in.Volumes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ResizeVolumes(in)
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

var _flag_VolumesService_ModifyVolumeAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
	cli.StringFlag{
		Name:  "volume",
		Usage: "volume",
		Value: "",
	},
	cli.StringFlag{
		Name:  "volume_name",
		Usage: "volume name",
		Value: "",
	},
}

func _func_VolumesService_ModifyVolumeAttributes(c *cli.Context) error {
	qc := pb.NewVolumesService(pkgGetServerInfo())
	in := new(pb.ModifyVolumeAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("volume") {
			in.Volume = proto.String(c.String("volume"))
		}
		if c.IsSet("volume_name") {
			in.VolumeName = proto.String(c.String("volume_name"))
		}
	}

	out, err := qc.ModifyVolumeAttributes(in)
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
