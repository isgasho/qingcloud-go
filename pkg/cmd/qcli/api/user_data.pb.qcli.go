// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: user_data.proto

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
	AllCommands = append(AllCommands, CmdUserDataService)
}

var CmdUserDataService = cli.Command{
	Name:    "userdata",
	Aliases: []string{},
	Usage:   "manage UserDataService",
	Subcommands: []cli.Command{
		{
			Name:    "UploadUserDataAttachment",
			Aliases: []string{},
			Usage:   "UploadUserDataAttachment",
			Flags:   _flag_UserDataService_UploadUserDataAttachment,
			Action:  _func_UserDataService_UploadUserDataAttachment,
		},
	},
}

var _flag_UserDataService_UploadUserDataAttachment = []cli.Flag{
	cli.StringFlag{
		Name:  "attachment_name",
		Usage: "attachment name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "attachment_content",
		Usage: "attachment content",
		Value: "", // json: slice/message/map/time
	},
}

func _func_UserDataService_UploadUserDataAttachment(c *cli.Context) error {
	conf := config.MustLoad(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewUserDataService(conf, zone)

	in := new(pb.UploadUserDataAttachmentInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("attachment_name") {
			in.AttachmentName = proto.String(c.String("attachment_name"))
		}
		if c.IsSet("attachment_content") {
			if err := json.Unmarshal([]byte(c.String("attachment_content")), &in.AttachmentContent); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.UploadUserDataAttachment(in)
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
