// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: misc.proto

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
			Action:  _func_MiscService_GrantQuotaIndep,
		},
		{
			Name:    "RevokeQuotaIndep",
			Aliases: []string{},
			Usage:   "RevokeQuotaIndep",
			Flags:   _flag_MiscService_RevokeQuotaIndep,
			Action:  _func_MiscService_RevokeQuotaIndep,
		},
		{
			Name:    "GetQuotaLeft",
			Aliases: []string{},
			Usage:   "GetQuotaLeft",
			Flags:   _flag_MiscService_GetQuotaLeft,
			Action:  _func_MiscService_GetQuotaLeft,
		},
	},
}

var _flag_MiscService_GrantQuotaIndep = []cli.Flag{
	cli.StringFlag{
		Name:  "user",
		Usage: "user",
		Value: "",
	},
	cli.StringFlag{
		Name:  "zone",
		Usage: "zone",
		Value: "",
	},
	cli.IntFlag{
		Name:  "instance",
		Usage: "instance",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "cpu",
		Usage: "cpu",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "memory",
		Usage: "memory",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hp_instance",
		Usage: "hp instance",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hp_cpu",
		Usage: "hp cpu",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hp_memory",
		Usage: "hp memory",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "volume",
		Usage: "volume",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "volume_size",
		Usage: "volume size",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hc_volume",
		Usage: "hc volume",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hc_volume_size",
		Usage: "hc volume size",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hpp_volume",
		Usage: "hpp volume",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hpp_volume_size",
		Usage: "hpp volume size",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "image",
		Usage: "image",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "loadbalancer",
		Usage: "loadbalancer",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "loadbalancer_policy",
		Usage: "loadbalancer policy",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "router",
		Usage: "router",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "eip",
		Usage: "eip",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "eip_bandwidth",
		Usage: "eip bandwidth",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "rdb",
		Usage: "rdb",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hpp_rdb",
		Usage: "hpp rdb",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "cache",
		Usage: "cache",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hp_cache",
		Usage: "hp cache",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "mongo",
		Usage: "mongo",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "hp_mongo",
		Usage: "hp mongo",
		Value: 0,
	},
}

func _func_MiscService_GrantQuotaIndep(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService("", "", zone)

	in := new(pb.GrantQuotaIndepInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("user") {
			in.User = proto.String(c.String("user"))
		}
		if c.IsSet("zone") {
			in.Zone = proto.String(c.String("zone"))
		}
		if c.IsSet("instance") {
			in.Instance = proto.Int32(int32(c.Int("instance")))
		}
		if c.IsSet("cpu") {
			in.Cpu = proto.Int32(int32(c.Int("cpu")))
		}
		if c.IsSet("memory") {
			in.Memory = proto.Int32(int32(c.Int("memory")))
		}
		if c.IsSet("hp_instance") {
			in.HpInstance = proto.Int32(int32(c.Int("hp_instance")))
		}
		if c.IsSet("hp_cpu") {
			in.HpCpu = proto.Int32(int32(c.Int("hp_cpu")))
		}
		if c.IsSet("hp_memory") {
			in.HpMemory = proto.Int32(int32(c.Int("hp_memory")))
		}
		if c.IsSet("volume") {
			in.Volume = proto.Int32(int32(c.Int("volume")))
		}
		if c.IsSet("volume_size") {
			in.VolumeSize = proto.Int32(int32(c.Int("volume_size")))
		}
		if c.IsSet("hc_volume") {
			in.HcVolume = proto.Int32(int32(c.Int("hc_volume")))
		}
		if c.IsSet("hc_volume_size") {
			in.HcVolumeSize = proto.Int32(int32(c.Int("hc_volume_size")))
		}
		if c.IsSet("hpp_volume") {
			in.HppVolume = proto.Int32(int32(c.Int("hpp_volume")))
		}
		if c.IsSet("hpp_volume_size") {
			in.HppVolumeSize = proto.Int32(int32(c.Int("hpp_volume_size")))
		}
		if c.IsSet("image") {
			in.Image = proto.Int32(int32(c.Int("image")))
		}
		if c.IsSet("loadbalancer") {
			in.Loadbalancer = proto.Int32(int32(c.Int("loadbalancer")))
		}
		if c.IsSet("loadbalancer_policy") {
			in.LoadbalancerPolicy = proto.Int32(int32(c.Int("loadbalancer_policy")))
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.Int32(int32(c.Int("vxnet")))
		}
		if c.IsSet("router") {
			in.Router = proto.Int32(int32(c.Int("router")))
		}
		if c.IsSet("eip") {
			in.Eip = proto.Int32(int32(c.Int("eip")))
		}
		if c.IsSet("eip_bandwidth") {
			in.EipBandwidth = proto.Int32(int32(c.Int("eip_bandwidth")))
		}
		if c.IsSet("rdb") {
			in.Rdb = proto.Int32(int32(c.Int("rdb")))
		}
		if c.IsSet("hpp_rdb") {
			in.HppRdb = proto.Int32(int32(c.Int("hpp_rdb")))
		}
		if c.IsSet("cache") {
			in.Cache = proto.Int32(int32(c.Int("cache")))
		}
		if c.IsSet("hp_cache") {
			in.HpCache = proto.Int32(int32(c.Int("hp_cache")))
		}
		if c.IsSet("mongo") {
			in.Mongo = proto.Int32(int32(c.Int("mongo")))
		}
		if c.IsSet("hp_mongo") {
			in.HpMongo = proto.Int32(int32(c.Int("hp_mongo")))
		}
	}

	out, err := qc.GrantQuotaIndep(in)
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

var _flag_MiscService_RevokeQuotaIndep = []cli.Flag{
	cli.StringFlag{
		Name:  "user",
		Usage: "user",
		Value: "",
	},
	cli.StringFlag{
		Name:  "zone",
		Usage: "zone",
		Value: "",
	},
}

func _func_MiscService_RevokeQuotaIndep(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService("", "", zone)

	in := new(pb.RevokeQuotaIndepInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("user") {
			in.User = proto.String(c.String("user"))
		}
		if c.IsSet("zone") {
			in.Zone = proto.String(c.String("zone"))
		}
	}

	out, err := qc.RevokeQuotaIndep(in)
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

var _flag_MiscService_GetQuotaLeft = []cli.Flag{
	cli.StringFlag{
		Name:  "resource_types",
		Usage: "resource types",
		Value: "", // json: slice/message/map/time
	},
}

func _func_MiscService_GetQuotaLeft(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewMiscService("", "", zone)

	in := new(pb.GetQuotaLeftInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("resource_types") {
			if err := json.Unmarshal([]byte(c.String("resource_types")), &in.ResourceTypes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.GetQuotaLeft(in)
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
