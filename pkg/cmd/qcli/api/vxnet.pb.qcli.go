// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: vxnet.proto

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
	AllCommands = append(AllCommands, CmdVxnetService)
}

var CmdVxnetService = cli.Command{
	Name:    "vxnet",
	Aliases: []string{},
	Usage:   "manage VxnetService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeVxnets",
			Aliases: []string{},
			Usage:   "DescribeVxnets",
			Flags:   _flag_VxnetService_DescribeVxnets,
			Action:  _func_VxnetService_DescribeVxnets,
		},
		{
			Name:    "CreateVxnets",
			Aliases: []string{},
			Usage:   "CreateVxnets",
			Flags:   _flag_VxnetService_CreateVxnets,
			Action:  _func_VxnetService_CreateVxnets,
		},
		{
			Name:    "DeleteVxnets",
			Aliases: []string{},
			Usage:   "DeleteVxnets",
			Flags:   _flag_VxnetService_DeleteVxnets,
			Action:  _func_VxnetService_DeleteVxnets,
		},
		{
			Name:    "JoinVxnet",
			Aliases: []string{},
			Usage:   "JoinVxnet",
			Flags:   _flag_VxnetService_JoinVxnet,
			Action:  _func_VxnetService_JoinVxnet,
		},
		{
			Name:    "LeaveVxnet",
			Aliases: []string{},
			Usage:   "LeaveVxnet",
			Flags:   _flag_VxnetService_LeaveVxnet,
			Action:  _func_VxnetService_LeaveVxnet,
		},
		{
			Name:    "ModifyVxnetAttributes",
			Aliases: []string{},
			Usage:   "ModifyVxnetAttributes",
			Flags:   _flag_VxnetService_ModifyVxnetAttributes,
			Action:  _func_VxnetService_ModifyVxnetAttributes,
		},
		{
			Name:    "DescribeVxnetInstances",
			Aliases: []string{},
			Usage:   "DescribeVxnetInstances",
			Flags:   _flag_VxnetService_DescribeVxnetInstances,
			Action:  _func_VxnetService_DescribeVxnetInstances,
		},
	},
}

var _flag_VxnetService_DescribeVxnets = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnets",
		Usage: "vxnets",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "vxnet_type",
		Usage: "vxnet type",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "search_word",
		Usage: "search word",
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

func _func_VxnetService_DescribeVxnets(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.DescribeVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("vxnets") {
			if err := json.Unmarshal([]byte(c.String("vxnets")), &in.Vxnets); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("vxnet_type") {
			in.VxnetType = proto.Int32(int32(c.Int("vxnet_type")))
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.String(c.String("search_word"))
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

	out, err := qc.DescribeVxnets(in)
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

var _flag_VxnetService_CreateVxnets = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet_name",
		Usage: "vxnet name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "vxnet_type",
		Usage: "vxnet type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "count",
		Usage: "count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "target_user",
		Usage: "target user",
		Value: "",
	},
}

func _func_VxnetService_CreateVxnets(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.CreateVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("vxnet_name") {
			in.VxnetName = proto.String(c.String("vxnet_name"))
		}
		if c.IsSet("vxnet_type") {
			in.VxnetType = proto.Int32(int32(c.Int("vxnet_type")))
		}
		if c.IsSet("count") {
			in.Count = proto.Int32(int32(c.Int("count")))
		}
		if c.IsSet("target_user") {
			in.TargetUser = proto.String(c.String("target_user"))
		}
	}

	out, err := qc.CreateVxnets(in)
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

var _flag_VxnetService_DeleteVxnets = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnets",
		Usage: "vxnets",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VxnetService_DeleteVxnets(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.DeleteVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("vxnets") {
			if err := json.Unmarshal([]byte(c.String("vxnets")), &in.Vxnets); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteVxnets(in)
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

var _flag_VxnetService_JoinVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VxnetService_JoinVxnet(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.JoinVxnetInput)

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
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.JoinVxnet(in)
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

var _flag_VxnetService_LeaveVxnet = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
}

func _func_VxnetService_LeaveVxnet(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.LeaveVxnetInput)

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
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.LeaveVxnet(in)
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

var _flag_VxnetService_ModifyVxnetAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "vxnet_name",
		Usage: "vxnet name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_VxnetService_ModifyVxnetAttributes(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.ModifyVxnetAttributesInput)

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
		if c.IsSet("vxnet_name") {
			in.VxnetName = proto.String(c.String("vxnet_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyVxnetAttributes(in)
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

var _flag_VxnetService_DescribeVxnetInstances = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instance_type",
		Usage: "instance type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "",
	},
	cli.StringFlag{
		Name:  "image",
		Usage: "image",
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

func _func_VxnetService_DescribeVxnetInstances(c *cli.Context) error {
	qc := pb.NewVxnetService(pkgGetServerInfo(c))
	in := new(pb.DescribeVxnetInstancesInput)

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
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("instance_type") {
			in.InstanceType = proto.String(c.String("instance_type"))
		}
		if c.IsSet("status") {
			in.Status = proto.String(c.String("status"))
		}
		if c.IsSet("image") {
			in.Image = proto.String(c.String("image"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeVxnetInstances(in)
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
