// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: tag.proto

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
	AllCommands = append(AllCommands, CmdTagService)
}

var CmdTagService = cli.Command{
	Name:    "tag",
	Aliases: []string{},
	Usage:   "manage TagService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeTags",
			Aliases: []string{},
			Usage:   "DescribeTags",
			Flags:   _flag_TagService_DescribeTags,
			Action:  _func_TagService_DescribeTags,
		},
		{
			Name:    "CreateTag",
			Aliases: []string{},
			Usage:   "CreateTag",
			Flags:   _flag_TagService_CreateTag,
			Action:  _func_TagService_CreateTag,
		},
		{
			Name:    "DeleteTags",
			Aliases: []string{},
			Usage:   "DeleteTags",
			Flags:   _flag_TagService_DeleteTags,
			Action:  _func_TagService_DeleteTags,
		},
		{
			Name:    "ModifyTagAttributes",
			Aliases: []string{},
			Usage:   "ModifyTagAttributes",
			Flags:   _flag_TagService_ModifyTagAttributes,
			Action:  _func_TagService_ModifyTagAttributes,
		},
		{
			Name:    "AttachTags",
			Aliases: []string{},
			Usage:   "AttachTags",
			Flags:   _flag_TagService_AttachTags,
			Action:  _func_TagService_AttachTags,
		},
		{
			Name:    "DetachTags",
			Aliases: []string{},
			Usage:   "DetachTags",
			Flags:   _flag_TagService_DetachTags,
			Action:  _func_TagService_DetachTags,
		},
	},
}

var _flag_TagService_DescribeTags = []cli.Flag{
	cli.StringFlag{
		Name:  "tags",
		Usage: "tags",
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

func _func_TagService_DescribeTags(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.DescribeTagsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
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

	out, err := qc.DescribeTags(in)
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

var _flag_TagService_CreateTag = []cli.Flag{
	cli.StringFlag{
		Name:  "tag_name",
		Usage: "tag name",
		Value: "",
	},
}

func _func_TagService_CreateTag(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.CreateTagInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("tag_name") {
			in.TagName = proto.String(c.String("tag_name"))
		}
	}

	out, err := qc.CreateTag(in)
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

var _flag_TagService_DeleteTags = []cli.Flag{
	cli.StringFlag{
		Name:  "tags",
		Usage: "tags",
		Value: "", // json: slice/message/map/time
	},
}

func _func_TagService_DeleteTags(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.DeleteTagsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteTags(in)
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

var _flag_TagService_ModifyTagAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "tag",
		Usage: "tag",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tag_name",
		Usage: "tag name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_TagService_ModifyTagAttributes(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.ModifyTagAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("tag") {
			in.Tag = proto.String(c.String("tag"))
		}
		if c.IsSet("tag_name") {
			in.TagName = proto.String(c.String("tag_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyTagAttributes(in)
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

var _flag_TagService_AttachTags = []cli.Flag{
	cli.StringFlag{
		Name:  "resource_tag_pairs",
		Usage: "resource tag pairs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_TagService_AttachTags(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.AttachTagsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("resource_tag_pairs") {
			if err := json.Unmarshal([]byte(c.String("resource_tag_pairs")), &in.ResourceTagPairs); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AttachTags(in)
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

var _flag_TagService_DetachTags = []cli.Flag{
	cli.StringFlag{
		Name:  "resource_tag_pairs",
		Usage: "resource tag pairs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_TagService_DetachTags(c *cli.Context) error {
	qc := pb.NewTagService(nil, nil)
	in := new(pb.DetachTagsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("resource_tag_pairs") {
			if err := json.Unmarshal([]byte(c.String("resource_tag_pairs")), &in.ResourceTagPairs); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DetachTags(in)
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
