// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: hadoop.proto

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
	AllCommands = append(AllCommands, CmdHadoopService)
}

var CmdHadoopService = cli.Command{
	Name:    "hadoop",
	Aliases: []string{},
	Usage:   "manage HadoopService",
	Subcommands: []cli.Command{
		{
			Name:    "AddHadoopNodes",
			Aliases: []string{},
			Usage:   "AddHadoopNodes",
			Flags:   _flag_HadoopService_AddHadoopNodes,
			Action:  _func_HadoopService_AddHadoopNodes,
		},
		{
			Name:    "DeleteHadoopNodes",
			Aliases: []string{},
			Usage:   "DeleteHadoopNodes",
			Flags:   _flag_HadoopService_DeleteHadoopNodes,
			Action:  _func_HadoopService_DeleteHadoopNodes,
		},
		{
			Name:    "StartHadoops",
			Aliases: []string{},
			Usage:   "StartHadoops",
			Flags:   _flag_HadoopService_StartHadoops,
			Action:  _func_HadoopService_StartHadoops,
		},
		{
			Name:    "StopHadoops",
			Aliases: []string{},
			Usage:   "StopHadoops",
			Flags:   _flag_HadoopService_StopHadoops,
			Action:  _func_HadoopService_StopHadoops,
		},
	},
}

var _flag_HadoopService_AddHadoopNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "hadoop",
		Usage: "hadoop",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "hadoop_node_name",
		Usage: "hadoop node name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_HadoopService_AddHadoopNodes(c *cli.Context) error {
	qc := pb.NewHadoopService(pkgGetServerInfo())
	in := new(pb.AddHadoopNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("hadoop") {
			in.Hadoop = proto.String(c.String("hadoop"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("hadoop_node_name") {
			in.HadoopNodeName = proto.String(c.String("hadoop_node_name"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddHadoopNodes(in)
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

var _flag_HadoopService_DeleteHadoopNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "hadoop",
		Usage: "hadoop",
		Value: "",
	},
	cli.StringFlag{
		Name:  "hadoop_nodes",
		Usage: "hadoop nodes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_HadoopService_DeleteHadoopNodes(c *cli.Context) error {
	qc := pb.NewHadoopService(pkgGetServerInfo())
	in := new(pb.DeleteHadoopNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("hadoop") {
			in.Hadoop = proto.String(c.String("hadoop"))
		}
		if c.IsSet("hadoop_nodes") {
			if err := json.Unmarshal([]byte(c.String("hadoop_nodes")), &in.HadoopNodes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteHadoopNodes(in)
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

var _flag_HadoopService_StartHadoops = []cli.Flag{
	cli.StringFlag{
		Name:  "hadoops",
		Usage: "hadoops",
		Value: "", // json: slice/message/map/time
	},
}

func _func_HadoopService_StartHadoops(c *cli.Context) error {
	qc := pb.NewHadoopService(pkgGetServerInfo())
	in := new(pb.StartHadoopsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("hadoops") {
			if err := json.Unmarshal([]byte(c.String("hadoops")), &in.Hadoops); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StartHadoops(in)
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

var _flag_HadoopService_StopHadoops = []cli.Flag{
	cli.StringFlag{
		Name:  "hadoops",
		Usage: "hadoops",
		Value: "", // json: slice/message/map/time
	},
}

func _func_HadoopService_StopHadoops(c *cli.Context) error {
	qc := pb.NewHadoopService(pkgGetServerInfo())
	in := new(pb.StopHadoopsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("hadoops") {
			if err := json.Unmarshal([]byte(c.String("hadoops")), &in.Hadoops); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StopHadoops(in)
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
