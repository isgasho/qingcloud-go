// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: eip.proto

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
	AllCommands = append(AllCommands, CmdEIPService)
}

var CmdEIPService = cli.Command{
	Name:     "eip",
	Aliases:  []string{},
	Usage:    "EIPService",
	Category: "SDK API Style Command",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeEips",
			Aliases: []string{},
			Usage:   "DescribeEips",
			Flags:   _flag_EIPService_DescribeEips,
			Action:  _func_EIPService_DescribeEips,
		},
		{
			Name:    "AllocateEips",
			Aliases: []string{},
			Usage:   "AllocateEips",
			Flags:   _flag_EIPService_AllocateEips,
			Action:  _func_EIPService_AllocateEips,
		},
		{
			Name:    "ReleaseEips",
			Aliases: []string{},
			Usage:   "ReleaseEips",
			Flags:   _flag_EIPService_ReleaseEips,
			Action:  _func_EIPService_ReleaseEips,
		},
		{
			Name:    "AssociateEip",
			Aliases: []string{},
			Usage:   "AssociateEip",
			Flags:   _flag_EIPService_AssociateEip,
			Action:  _func_EIPService_AssociateEip,
		},
		{
			Name:    "DissociateEips",
			Aliases: []string{},
			Usage:   "DissociateEips",
			Flags:   _flag_EIPService_DissociateEips,
			Action:  _func_EIPService_DissociateEips,
		},
		{
			Name:    "ChangeEipsBandwidth",
			Aliases: []string{},
			Usage:   "ChangeEipsBandwidth",
			Flags:   _flag_EIPService_ChangeEipsBandwidth,
			Action:  _func_EIPService_ChangeEipsBandwidth,
		},
		{
			Name:    "ChangeEipsBillingMode",
			Aliases: []string{},
			Usage:   "ChangeEipsBillingMode",
			Flags:   _flag_EIPService_ChangeEipsBillingMode,
			Action:  _func_EIPService_ChangeEipsBillingMode,
		},
		{
			Name:    "ModifyEipAttributes",
			Aliases: []string{},
			Usage:   "ModifyEipAttributes",
			Flags:   _flag_EIPService_ModifyEipAttributes,
			Action:  _func_EIPService_ModifyEipAttributes,
		},
	},
}

var _flag_EIPService_DescribeEips = []cli.Flag{
	cli.StringFlag{
		Name:  "eips",
		Usage: "eips",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instance_id",
		Usage: "instance id",
		Value: "",
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "",
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

func _func_EIPService_DescribeEips(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.DescribeEipsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eips") {
			if err := json.Unmarshal([]byte(c.String("eips")), &in.Eips); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("instance_id") {
			in.InstanceId = proto.String(c.String("instance_id"))
		}
		if c.IsSet("status") {
			in.Status = proto.String(c.String("status"))
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

	out, err := qc.DescribeEips(in)
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

var _flag_EIPService_AllocateEips = []cli.Flag{
	cli.IntFlag{
		Name:  "bandwidth",
		Usage: "bandwidth",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "billing_mode",
		Usage: "billing mode",
		Value: "",
	},
	cli.StringFlag{
		Name:  "eip_name",
		Usage: "eip name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "count",
		Usage: "count",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "need_icp",
		Usage: "need icp",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "target_user",
		Usage: "target user",
		Value: "",
	},
}

func _func_EIPService_AllocateEips(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.AllocateEipsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("bandwidth") {
			in.Bandwidth = proto.Int32(int32(c.Int("bandwidth")))
		}
		if c.IsSet("billing_mode") {
			in.BillingMode = proto.String(c.String("billing_mode"))
		}
		if c.IsSet("eip_name") {
			in.EipName = proto.String(c.String("eip_name"))
		}
		if c.IsSet("count") {
			in.Count = proto.Int32(int32(c.Int("count")))
		}
		if c.IsSet("need_icp") {
			in.NeedIcp = proto.Int32(int32(c.Int("need_icp")))
		}
		if c.IsSet("target_user") {
			in.TargetUser = proto.String(c.String("target_user"))
		}
	}

	out, err := qc.AllocateEips(in)
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

var _flag_EIPService_ReleaseEips = []cli.Flag{
	cli.StringFlag{
		Name:  "eips",
		Usage: "eips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_EIPService_ReleaseEips(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.ReleaseEipsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eips") {
			if err := json.Unmarshal([]byte(c.String("eips")), &in.Eips); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.ReleaseEips(in)
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

var _flag_EIPService_AssociateEip = []cli.Flag{
	cli.StringFlag{
		Name:  "eip",
		Usage: "eip",
		Value: "",
	},
	cli.StringFlag{
		Name:  "instance",
		Usage: "instance",
		Value: "",
	},
}

func _func_EIPService_AssociateEip(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.AssociateEipInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eip") {
			in.Eip = proto.String(c.String("eip"))
		}
		if c.IsSet("instance") {
			in.Instance = proto.String(c.String("instance"))
		}
	}

	out, err := qc.AssociateEip(in)
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

var _flag_EIPService_DissociateEips = []cli.Flag{
	cli.StringFlag{
		Name:  "eips",
		Usage: "eips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_EIPService_DissociateEips(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.DissociateEipsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eips") {
			if err := json.Unmarshal([]byte(c.String("eips")), &in.Eips); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DissociateEips(in)
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

var _flag_EIPService_ChangeEipsBandwidth = []cli.Flag{
	cli.StringFlag{
		Name:  "eips",
		Usage: "eips",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "bandwidth",
		Usage: "bandwidth",
		Value: 0,
	},
}

func _func_EIPService_ChangeEipsBandwidth(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.ChangeEipsBandwidthInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eips") {
			if err := json.Unmarshal([]byte(c.String("eips")), &in.Eips); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("bandwidth") {
			in.Bandwidth = proto.Int32(int32(c.Int("bandwidth")))
		}
	}

	out, err := qc.ChangeEipsBandwidth(in)
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

var _flag_EIPService_ChangeEipsBillingMode = []cli.Flag{
	cli.StringFlag{
		Name:  "eips",
		Usage: "eips",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "billing_mode",
		Usage: "billing mode",
		Value: "",
	},
}

func _func_EIPService_ChangeEipsBillingMode(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.ChangeEipsBillingModeInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eips") {
			if err := json.Unmarshal([]byte(c.String("eips")), &in.Eips); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("billing_mode") {
			in.BillingMode = proto.String(c.String("billing_mode"))
		}
	}

	out, err := qc.ChangeEipsBillingMode(in)
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

var _flag_EIPService_ModifyEipAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "eip",
		Usage: "eip",
		Value: "",
	},
	cli.StringFlag{
		Name:  "eip_name",
		Usage: "eip name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_EIPService_ModifyEipAttributes(c *cli.Context) error {
	qc := pb.NewEIPService(pkgGetServerInfo(c))
	in := new(pb.ModifyEipAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("eip") {
			in.Eip = proto.String(c.String("eip"))
		}
		if c.IsSet("eip_name") {
			in.EipName = proto.String(c.String("eip_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyEipAttributes(in)
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
