// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: nic.proto

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
	AllCommands = append(AllCommands, CmdNicService)
}

var CmdNicService = cli.Command{
	Name:    "nic",
	Aliases: []string{},
	Usage:   "manage NicService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateNics",
			Aliases: []string{},
			Usage:   "CreateNics",
			Flags:   _flag_NicService_CreateNics,
			Action:  _func_NicService_CreateNics,
		},
		{
			Name:    "DescribeNics",
			Aliases: []string{},
			Usage:   "DescribeNics",
			Flags:   _flag_NicService_DescribeNics,
			Action:  _func_NicService_DescribeNics,
		},
		{
			Name:    "AttachNics",
			Aliases: []string{},
			Usage:   "AttachNics",
			Flags:   _flag_NicService_AttachNics,
			Action:  _func_NicService_AttachNics,
		},
		{
			Name:    "DetachNics",
			Aliases: []string{},
			Usage:   "DetachNics",
			Flags:   _flag_NicService_DetachNics,
			Action:  _func_NicService_DetachNics,
		},
		{
			Name:    "ModifyNicAttributes",
			Aliases: []string{},
			Usage:   "ModifyNicAttributes",
			Flags:   _flag_NicService_ModifyNicAttributes,
			Action:  _func_NicService_ModifyNicAttributes,
		},
		{
			Name:    "DeleteNics",
			Aliases: []string{},
			Usage:   "DeleteNics",
			Flags:   _flag_NicService_DeleteNics,
			Action:  _func_NicService_DeleteNics,
		},
	},
}

var _flag_NicService_CreateNics = []cli.Flag{
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "nic_name",
		Usage: "nic name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "count",
		Usage: "count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_NicService_CreateNics(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.CreateNicsInput)

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
		if c.IsSet("nic_name") {
			in.NicName = proto.String(c.String("nic_name"))
		}
		if c.IsSet("count") {
			in.Count = proto.Int32(int32(c.Int("count")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.CreateNics(in)
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

var _flag_NicService_DescribeNics = []cli.Flag{
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "nic_name",
		Usage: "nic name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "nics",
		Usage: "nics",
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
		Value: "",
	},
	cli.IntFlag{
		Name:  "vxnet_type",
		Usage: "vxnet type",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "vxnets",
		Usage: "vxnets",
		Value: "", // json: slice/message/map/time
	},
}

func _func_NicService_DescribeNics(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.DescribeNicsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
		if c.IsSet("nic_name") {
			in.NicName = proto.String(c.String("nic_name"))
		}
		if c.IsSet("nics") {
			if err := json.Unmarshal([]byte(c.String("nics")), &in.Nics); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("status") {
			in.Status = proto.String(c.String("status"))
		}
		if c.IsSet("vxnet_type") {
			in.VxnetType = proto.Int32(int32(c.Int("vxnet_type")))
		}
		if c.IsSet("vxnets") {
			if err := json.Unmarshal([]byte(c.String("vxnets")), &in.Vxnets); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DescribeNics(in)
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

var _flag_NicService_AttachNics = []cli.Flag{
	cli.StringFlag{
		Name:  "nics",
		Usage: "nics",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instance",
		Usage: "instance",
		Value: "",
	},
}

func _func_NicService_AttachNics(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.AttachNicsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("nics") {
			if err := json.Unmarshal([]byte(c.String("nics")), &in.Nics); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("instance") {
			in.Instance = proto.String(c.String("instance"))
		}
	}

	out, err := qc.AttachNics(in)
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

var _flag_NicService_DetachNics = []cli.Flag{
	cli.StringFlag{
		Name:  "nics",
		Usage: "nics",
		Value: "", // json: slice/message/map/time
	},
}

func _func_NicService_DetachNics(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.DetachNicsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("nics") {
			if err := json.Unmarshal([]byte(c.String("nics")), &in.Nics); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DetachNics(in)
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

var _flag_NicService_ModifyNicAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "nic",
		Usage: "nic",
		Value: "",
	},
	cli.StringFlag{
		Name:  "nic_name",
		Usage: "nic name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ip",
		Usage: "private ip",
		Value: "",
	},
}

func _func_NicService_ModifyNicAttributes(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.ModifyNicAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("nic") {
			in.Nic = proto.String(c.String("nic"))
		}
		if c.IsSet("nic_name") {
			in.NicName = proto.String(c.String("nic_name"))
		}
		if c.IsSet("private_ip") {
			in.PrivateIp = proto.String(c.String("private_ip"))
		}
	}

	out, err := qc.ModifyNicAttributes(in)
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

var _flag_NicService_DeleteNics = []cli.Flag{
	cli.StringFlag{
		Name:  "nics",
		Usage: "nics",
		Value: "", // json: slice/message/map/time
	},
}

func _func_NicService_DeleteNics(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewNicService("", "", zone)

	in := new(pb.DeleteNicsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("nics") {
			if err := json.Unmarshal([]byte(c.String("nics")), &in.Nics); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteNics(in)
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
