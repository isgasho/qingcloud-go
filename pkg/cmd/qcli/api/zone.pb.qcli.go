// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: zone.proto

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
	AllCommands = append(AllCommands, CmdZoneService)
}

var CmdZoneService = cli.Command{
	Name:    "zone",
	Aliases: []string{},
	Usage:   "manage ZoneService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeZones",
			Aliases: []string{},
			Usage:   "DescribeZones",
			Flags:   _flag_ZoneService_DescribeZones,
			Action:  _func_ZoneService_DescribeZones,
		},
	},
}

var _flag_ZoneService_DescribeZones = []cli.Flag{
	cli.StringFlag{
		Name:  "zones",
		Usage: "zones",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "", // json: slice/message/map/time
	},
}

func _func_ZoneService_DescribeZones(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewZoneService(conf, zone)

	in := new(pb.DescribeZonesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("zones") {
			if err := json.Unmarshal([]byte(c.String("zones")), &in.Zones); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DescribeZones(in)
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
