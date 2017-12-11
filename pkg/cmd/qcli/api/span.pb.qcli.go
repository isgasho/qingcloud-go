// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: span.proto

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
	AllCommands = append(AllCommands, CmdSpanService)
}

var CmdSpanService = cli.Command{
	Name:    "span",
	Aliases: []string{},
	Usage:   "manage SpanService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateSpan",
			Aliases: []string{},
			Usage:   "CreateSpan",
			Flags:   _flag_SpanService_CreateSpan,
			Action:  _func_SpanService_CreateSpan,
		},
		{
			Name:    "DescribeSpans",
			Aliases: []string{},
			Usage:   "DescribeSpans",
			Flags:   _flag_SpanService_DescribeSpans,
			Action:  _func_SpanService_DescribeSpans,
		},
		{
			Name:    "DeleteSpans",
			Aliases: []string{},
			Usage:   "DeleteSpans",
			Flags:   _flag_SpanService_DeleteSpans,
			Action:  _func_SpanService_DeleteSpans,
		},
		{
			Name:    "AddSpanMembers",
			Aliases: []string{},
			Usage:   "AddSpanMembers",
			Flags:   _flag_SpanService_AddSpanMembers,
			Action:  _func_SpanService_AddSpanMembers,
		},
		{
			Name:    "RemoveSpanMembers",
			Aliases: []string{},
			Usage:   "RemoveSpanMembers",
			Flags:   _flag_SpanService_RemoveSpanMembers,
			Action:  _func_SpanService_RemoveSpanMembers,
		},
		{
			Name:    "ModifySpanAttributes",
			Aliases: []string{},
			Usage:   "ModifySpanAttributes",
			Flags:   _flag_SpanService_ModifySpanAttributes,
			Action:  _func_SpanService_ModifySpanAttributes,
		},
		{
			Name:    "UpdateSpan",
			Aliases: []string{},
			Usage:   "UpdateSpan",
			Flags:   _flag_SpanService_UpdateSpan,
			Action:  _func_SpanService_UpdateSpan,
		},
	},
}

var _flag_SpanService_CreateSpan = []cli.Flag{
	cli.StringFlag{
		Name:  "span_name",
		Usage: "span name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "flag",
		Usage: "flag",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "ip_addr",
		Usage: "ip addr",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tunnel_type",
		Usage: "tunnel type",
		Value: "",
	},
	cli.IntFlag{
		Name:  "tunnel_key",
		Usage: "tunnel key",
		Value: 0,
	},
}

func _func_SpanService_CreateSpan(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.CreateSpanInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("span_name") {
			in.SpanName = proto.String(c.String("span_name"))
		}
		if c.IsSet("flag") {
			in.Flag = proto.Int32(int32(c.Int("flag")))
		}
		if c.IsSet("ip_addr") {
			in.IpAddr = proto.String(c.String("ip_addr"))
		}
		if c.IsSet("tunnel_type") {
			in.TunnelType = proto.String(c.String("tunnel_type"))
		}
		if c.IsSet("tunnel_key") {
			in.TunnelKey = proto.Int32(int32(c.Int("tunnel_key")))
		}
	}

	out, err := qc.CreateSpan(in)
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

var _flag_SpanService_DescribeSpans = []cli.Flag{
	cli.StringFlag{
		Name:  "spans",
		Usage: "spans",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "span_name",
		Usage: "span name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "ip_addr",
		Usage: "ip addr",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tags",
		Usage: "tags",
		Value: "", // json: slice/message/map/time
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

func _func_SpanService_DescribeSpans(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.DescribeSpansInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("spans") {
			if err := json.Unmarshal([]byte(c.String("spans")), &in.Spans); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("span_name") {
			in.SpanName = proto.String(c.String("span_name"))
		}
		if c.IsSet("ip_addr") {
			in.IpAddr = proto.String(c.String("ip_addr"))
		}
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeSpans(in)
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

var _flag_SpanService_DeleteSpans = []cli.Flag{
	cli.StringFlag{
		Name:  "spans",
		Usage: "spans",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SpanService_DeleteSpans(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.DeleteSpansInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("spans") {
			if err := json.Unmarshal([]byte(c.String("spans")), &in.Spans); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteSpans(in)
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

var _flag_SpanService_AddSpanMembers = []cli.Flag{
	cli.StringFlag{
		Name:  "span",
		Usage: "span",
		Value: "",
	},
	cli.StringFlag{
		Name:  "resources",
		Usage: "resources",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SpanService_AddSpanMembers(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.AddSpanMembersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("span") {
			in.Span = proto.String(c.String("span"))
		}
		if c.IsSet("resources") {
			if err := json.Unmarshal([]byte(c.String("resources")), &in.Resources); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddSpanMembers(in)
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

var _flag_SpanService_RemoveSpanMembers = []cli.Flag{
	cli.StringFlag{
		Name:  "span",
		Usage: "span",
		Value: "",
	},
	cli.StringFlag{
		Name:  "resources",
		Usage: "resources",
		Value: "", // json: slice/message/map/time
	},
}

func _func_SpanService_RemoveSpanMembers(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.RemoveSpanMembersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("span") {
			in.Span = proto.String(c.String("span"))
		}
		if c.IsSet("resources") {
			if err := json.Unmarshal([]byte(c.String("resources")), &in.Resources); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.RemoveSpanMembers(in)
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

var _flag_SpanService_ModifySpanAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "span_id",
		Usage: "span id",
		Value: "",
	},
	cli.StringFlag{
		Name:  "span_name",
		Usage: "span name",
		Value: "",
	},
	cli.IntFlag{
		Name:  "flag",
		Usage: "flag",
		Value: 0,
	},
	cli.StringFlag{
		Name:  "ip_addr",
		Usage: "ip addr",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tunnel_type",
		Usage: "tunnel type",
		Value: "",
	},
	cli.IntFlag{
		Name:  "tunnel_key",
		Usage: "tunnel key",
		Value: 0,
	},
}

func _func_SpanService_ModifySpanAttributes(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.ModifySpanAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("span_id") {
			in.SpanId = proto.String(c.String("span_id"))
		}
		if c.IsSet("span_name") {
			in.SpanName = proto.String(c.String("span_name"))
		}
		if c.IsSet("flag") {
			in.Flag = proto.Int32(int32(c.Int("flag")))
		}
		if c.IsSet("ip_addr") {
			in.IpAddr = proto.String(c.String("ip_addr"))
		}
		if c.IsSet("tunnel_type") {
			in.TunnelType = proto.String(c.String("tunnel_type"))
		}
		if c.IsSet("tunnel_key") {
			in.TunnelKey = proto.Int32(int32(c.Int("tunnel_key")))
		}
	}

	out, err := qc.ModifySpanAttributes(in)
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

var _flag_SpanService_UpdateSpan = []cli.Flag{
	cli.StringFlag{
		Name:  "span",
		Usage: "span",
		Value: "",
	},
}

func _func_SpanService_UpdateSpan(c *cli.Context) error {
	qc := pb.NewSpanService(pkgGetServerInfo(c))
	in := new(pb.UpdateSpanInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("span") {
			in.Span = proto.String(c.String("span"))
		}
	}

	out, err := qc.UpdateSpan(in)
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
