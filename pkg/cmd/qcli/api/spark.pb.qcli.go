// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: spark.proto

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
	AllCommands = append(AllCommands, CmdSparkService)
}

var CmdSparkService = cli.Command{
	Name:    "spark",
	Aliases: []string{},
	Usage:   "manage SparkService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateSpark",
			Aliases: []string{},
			Usage:   "CreateSpark",
			Flags:   _flag_SparkService_CreateSpark,
			Action:  _func_SparkService_CreateSpark,
		},
		{
			Name:    "DescribeSparks",
			Aliases: []string{},
			Usage:   "DescribeSparks",
			Flags:   _flag_SparkService_DescribeSparks,
			Action:  _func_SparkService_DescribeSparks,
		},
		{
			Name:    "AddSparkNodes",
			Aliases: []string{},
			Usage:   "AddSparkNodes",
			Flags:   _flag_SparkService_AddSparkNodes,
			Action:  _func_SparkService_AddSparkNodes,
		},
		{
			Name:    "DeleteSparkNodes",
			Aliases: []string{},
			Usage:   "DeleteSparkNodes",
			Flags:   _flag_SparkService_DeleteSparkNodes,
			Action:  _func_SparkService_DeleteSparkNodes,
		},
		{
			Name:    "StartSparks",
			Aliases: []string{},
			Usage:   "StartSparks",
			Flags:   _flag_SparkService_StartSparks,
			Action:  _func_SparkService_StartSparks,
		},
		{
			Name:    "StopSparks",
			Aliases: []string{},
			Usage:   "StopSparks",
			Flags:   _flag_SparkService_StopSparks,
			Action:  _func_SparkService_StopSparks,
		},
		{
			Name:    "DeleteSparks",
			Aliases: []string{},
			Usage:   "DeleteSparks",
			Flags:   _flag_SparkService_DeleteSparks,
			Action:  _func_SparkService_DeleteSparks,
		},
	},
}

var _flag_SparkService_CreateSpark = []cli.Flag{
	cli.StringFlag{
		Name:  "spark_version",
		Usage: "spark version",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "enable_hdfs",
		Usage: "enable hdfs",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "spark_type",
		Usage: "spark type",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "storage_size",
		Usage: "storage size",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "vxnet",
		Usage: "vxnet",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
	cli.StringFlag{
		Name:  "spark_name",
		Usage: "spark name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "spark_class",
		Usage: "spark class",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_CreateSpark(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.CreateSparkInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("spark_version") {
			in.SparkVersion = proto.String(c.String("spark_version"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("enable_hdfs") {
			in.EnableHdfs = proto.Int32(int32(c.Int("enable_hdfs")))
		}
		if c.IsSet("spark_type") {
			in.SparkType = proto.Int32(int32(c.Int("spark_type")))
		}
		if c.IsSet("storage_size") {
			in.StorageSize = proto.Int32(int32(c.Int("storage_size")))
		}
		if c.IsSet("vxnet") {
			in.Vxnet = proto.String(c.String("vxnet"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
		if c.IsSet("spark_name") {
			in.SparkName = proto.String(c.String("spark_name"))
		}
		if c.IsSet("spark_class") {
			in.SparkClass = proto.Int32(int32(c.Int("spark_class")))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.CreateSpark(in)
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

var _flag_SparkService_DescribeSparks = []cli.Flag{
	cli.StringFlag{
		Name:  "sparks",
		Usage: "sparks",
		Value: "", // json: slice/message/map/time
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

func _func_SparkService_DescribeSparks(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.DescribeSparksInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("sparks") {
			if err := json.Unmarshal([]byte(c.String("sparks")), &in.Sparks); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("status") {
			if err := json.Unmarshal([]byte(c.String("status")), &in.Status); err != nil {
				log.Fatal(err)
			}
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

	out, err := qc.DescribeSparks(in)
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

var _flag_SparkService_AddSparkNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "spark",
		Usage: "spark",
		Value: "",
	},
	cli.IntFlag{
		Name:  "node_count",
		Usage: "node count",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "spark_node_name",
		Usage: "spark node name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "private_ips",
		Usage: "private ips",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_AddSparkNodes(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.AddSparkNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("spark") {
			in.Spark = proto.String(c.String("spark"))
		}
		if c.IsSet("node_count") {
			in.NodeCount = proto.Int32(int32(c.Int("node_count")))
		}
		if c.IsSet("spark_node_name") {
			in.SparkNodeName = proto.String(c.String("spark_node_name"))
		}
		if c.IsSet("private_ips") {
			if err := json.Unmarshal([]byte(c.String("private_ips")), &in.PrivateIps); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddSparkNodes(in)
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

var _flag_SparkService_DeleteSparkNodes = []cli.Flag{
	cli.StringFlag{
		Name:  "spark",
		Usage: "spark",
		Value: "",
	},
	cli.StringFlag{
		Name:  "spark_nodes",
		Usage: "spark nodes",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_DeleteSparkNodes(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.DeleteSparkNodesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("spark") {
			in.Spark = proto.String(c.String("spark"))
		}
		if c.IsSet("spark_nodes") {
			if err := json.Unmarshal([]byte(c.String("spark_nodes")), &in.SparkNodes); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteSparkNodes(in)
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

var _flag_SparkService_StartSparks = []cli.Flag{
	cli.StringFlag{
		Name:  "sparks",
		Usage: "sparks",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_StartSparks(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.StartSparksInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("sparks") {
			if err := json.Unmarshal([]byte(c.String("sparks")), &in.Sparks); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StartSparks(in)
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

var _flag_SparkService_StopSparks = []cli.Flag{
	cli.StringFlag{
		Name:  "sparks",
		Usage: "sparks",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_StopSparks(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.StopSparksInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("sparks") {
			if err := json.Unmarshal([]byte(c.String("sparks")), &in.Sparks); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.StopSparks(in)
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

var _flag_SparkService_DeleteSparks = []cli.Flag{
	cli.StringFlag{
		Name:  "sparks",
		Usage: "sparks",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SparkService_DeleteSparks(c *cli.Context) error {
	zone := c.GlobalString("zone")
	qc := pb.NewSparkService("", "", zone)

	in := new(pb.DeleteSparksInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("sparks") {
			if err := json.Unmarshal([]byte(c.String("sparks")), &in.Sparks); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteSparks(in)
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
