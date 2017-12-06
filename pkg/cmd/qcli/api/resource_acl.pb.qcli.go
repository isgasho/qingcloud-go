// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: resource_acl.proto

package qcli_pb

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/chai2010/qingcloud-go/pkg/logger"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdResourceACLService)
}

var CmdResourceACLService = cli.Command{
	Name:    "resourceacl",
	Aliases: []string{},
	Usage:   "manage ResourceACLService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeSharedResourceGroups",
			Aliases: []string{},
			Usage:   "DescribeSharedResourceGroups",
			Flags:   _flag_ResourceACLService_DescribeSharedResourceGroups,
			Action:  _cmd_ResourceACLService_DescribeSharedResourceGroups,
		},
		{
			Name:    "DescribeResourceGroups",
			Aliases: []string{},
			Usage:   "DescribeResourceGroups",
			Flags:   _flag_ResourceACLService_DescribeResourceGroups,
			Action:  _cmd_ResourceACLService_DescribeResourceGroups,
		},
		{
			Name:    "CreateResourceGroups",
			Aliases: []string{},
			Usage:   "CreateResourceGroups",
			Flags:   _flag_ResourceACLService_CreateResourceGroups,
			Action:  _cmd_ResourceACLService_CreateResourceGroups,
		},
		{
			Name:    "ModifyResourceGroupAttributes",
			Aliases: []string{},
			Usage:   "ModifyResourceGroupAttributes",
			Flags:   _flag_ResourceACLService_ModifyResourceGroupAttributes,
			Action:  _cmd_ResourceACLService_ModifyResourceGroupAttributes,
		},
		{
			Name:    "DeleteResourceGroups",
			Aliases: []string{},
			Usage:   "DeleteResourceGroups",
			Flags:   _flag_ResourceACLService_DeleteResourceGroups,
			Action:  _cmd_ResourceACLService_DeleteResourceGroups,
		},
		{
			Name:    "DescribeResourceGroupItems",
			Aliases: []string{},
			Usage:   "DescribeResourceGroupItems",
			Flags:   _flag_ResourceACLService_DescribeResourceGroupItems,
			Action:  _cmd_ResourceACLService_DescribeResourceGroupItems,
		},
		{
			Name:    "AddResourceGroupItems",
			Aliases: []string{},
			Usage:   "AddResourceGroupItems",
			Flags:   _flag_ResourceACLService_AddResourceGroupItems,
			Action:  _cmd_ResourceACLService_AddResourceGroupItems,
		},
		{
			Name:    "DeleteResourceGroupItems",
			Aliases: []string{},
			Usage:   "DeleteResourceGroupItems",
			Flags:   _flag_ResourceACLService_DeleteResourceGroupItems,
			Action:  _cmd_ResourceACLService_DeleteResourceGroupItems,
		},
		{
			Name:    "DescribeUserGroups",
			Aliases: []string{},
			Usage:   "DescribeUserGroups",
			Flags:   _flag_ResourceACLService_DescribeUserGroups,
			Action:  _cmd_ResourceACLService_DescribeUserGroups,
		},
		{
			Name:    "CreateUserGroups",
			Aliases: []string{},
			Usage:   "CreateUserGroups",
			Flags:   _flag_ResourceACLService_CreateUserGroups,
			Action:  _cmd_ResourceACLService_CreateUserGroups,
		},
		{
			Name:    "ModifyUserGroupAttributes",
			Aliases: []string{},
			Usage:   "ModifyUserGroupAttributes",
			Flags:   _flag_ResourceACLService_ModifyUserGroupAttributes,
			Action:  _cmd_ResourceACLService_ModifyUserGroupAttributes,
		},
		{
			Name:    "DeleteUserGroups",
			Aliases: []string{},
			Usage:   "DeleteUserGroups",
			Flags:   _flag_ResourceACLService_DeleteUserGroups,
			Action:  _cmd_ResourceACLService_DeleteUserGroups,
		},
		{
			Name:    "DescribeUserGroupMembers",
			Aliases: []string{},
			Usage:   "DescribeUserGroupMembers",
			Flags:   _flag_ResourceACLService_DescribeUserGroupMembers,
			Action:  _cmd_ResourceACLService_DescribeUserGroupMembers,
		},
		{
			Name:    "AddUserGroupMembers",
			Aliases: []string{},
			Usage:   "AddUserGroupMembers",
			Flags:   _flag_ResourceACLService_AddUserGroupMembers,
			Action:  _cmd_ResourceACLService_AddUserGroupMembers,
		},
		{
			Name:    "ModifyUserGroupMemberAttributes",
			Aliases: []string{},
			Usage:   "ModifyUserGroupMemberAttributes",
			Flags:   _flag_ResourceACLService_ModifyUserGroupMemberAttributes,
			Action:  _cmd_ResourceACLService_ModifyUserGroupMemberAttributes,
		},
		{
			Name:    "DeleteUserGroupMembers",
			Aliases: []string{},
			Usage:   "DeleteUserGroupMembers",
			Flags:   _flag_ResourceACLService_DeleteUserGroupMembers,
			Action:  _cmd_ResourceACLService_DeleteUserGroupMembers,
		},
		{
			Name:    "DescribeGroupRoles",
			Aliases: []string{},
			Usage:   "DescribeGroupRoles",
			Flags:   _flag_ResourceACLService_DescribeGroupRoles,
			Action:  _cmd_ResourceACLService_DescribeGroupRoles,
		},
		{
			Name:    "CreateGroupRoles",
			Aliases: []string{},
			Usage:   "CreateGroupRoles",
			Flags:   _flag_ResourceACLService_CreateGroupRoles,
			Action:  _cmd_ResourceACLService_CreateGroupRoles,
		},
		{
			Name:    "ModifyGroupRoleAttributes",
			Aliases: []string{},
			Usage:   "ModifyGroupRoleAttributes",
			Flags:   _flag_ResourceACLService_ModifyGroupRoleAttributes,
			Action:  _cmd_ResourceACLService_ModifyGroupRoleAttributes,
		},
		{
			Name:    "DeleteGroupRoles",
			Aliases: []string{},
			Usage:   "DeleteGroupRoles",
			Flags:   _flag_ResourceACLService_DeleteGroupRoles,
			Action:  _cmd_ResourceACLService_DeleteGroupRoles,
		},
		{
			Name:    "DescribeGroupRoleRules",
			Aliases: []string{},
			Usage:   "DescribeGroupRoleRules",
			Flags:   _flag_ResourceACLService_DescribeGroupRoleRules,
			Action:  _cmd_ResourceACLService_DescribeGroupRoleRules,
		},
		{
			Name:    "AddGroupRoleRules",
			Aliases: []string{},
			Usage:   "AddGroupRoleRules",
			Flags:   _flag_ResourceACLService_AddGroupRoleRules,
			Action:  _cmd_ResourceACLService_AddGroupRoleRules,
		},
		{
			Name:    "ModifyGroupRoleRuleAttributes",
			Aliases: []string{},
			Usage:   "ModifyGroupRoleRuleAttributes",
			Flags:   _flag_ResourceACLService_ModifyGroupRoleRuleAttributes,
			Action:  _cmd_ResourceACLService_ModifyGroupRoleRuleAttributes,
		},
		{
			Name:    "DeleteGroupRoleRules",
			Aliases: []string{},
			Usage:   "DeleteGroupRoleRules",
			Flags:   _flag_ResourceACLService_DeleteGroupRoleRules,
			Action:  _cmd_ResourceACLService_DeleteGroupRoleRules,
		},
		{
			Name:    "GrantResourceGroupsToUserGroups",
			Aliases: []string{},
			Usage:   "GrantResourceGroupsToUserGroups",
			Flags:   _flag_ResourceACLService_GrantResourceGroupsToUserGroups,
			Action:  _cmd_ResourceACLService_GrantResourceGroupsToUserGroups,
		},
		{
			Name:    "RevokeResourceGroupsFromUserGroups",
			Aliases: []string{},
			Usage:   "RevokeResourceGroupsFromUserGroups",
			Flags:   _flag_ResourceACLService_RevokeResourceGroupsFromUserGroups,
			Action:  _cmd_ResourceACLService_RevokeResourceGroupsFromUserGroups,
		},
		{
			Name:    "DescribeResourceUserGroups",
			Aliases: []string{},
			Usage:   "DescribeResourceUserGroups",
			Flags:   _flag_ResourceACLService_DescribeResourceUserGroups,
			Action:  _cmd_ResourceACLService_DescribeResourceUserGroups,
		},
	},
}

var _flag_ResourceACLService_DescribeSharedResourceGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeSharedResourceGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeSharedResourceGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSharedResourceGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeResourceGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeResourceGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeResourceGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeResourceGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_CreateResourceGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_CreateResourceGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.CreateResourceGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateResourceGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_ModifyResourceGroupAttributes = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_ModifyResourceGroupAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.ModifyResourceGroupAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyResourceGroupAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteResourceGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteResourceGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteResourceGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteResourceGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeResourceGroupItems = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeResourceGroupItems(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeResourceGroupItemsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeResourceGroupItems(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_AddResourceGroupItems = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_AddResourceGroupItems(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.AddResourceGroupItemsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddResourceGroupItems(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteResourceGroupItems = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteResourceGroupItems(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteResourceGroupItemsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteResourceGroupItems(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_CreateUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_CreateUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.CreateUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_ModifyUserGroupAttributes = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_ModifyUserGroupAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.ModifyUserGroupAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyUserGroupAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeUserGroupMembers = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeUserGroupMembers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeUserGroupMembersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeUserGroupMembers(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_AddUserGroupMembers = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_AddUserGroupMembers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.AddUserGroupMembersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddUserGroupMembers(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_ModifyUserGroupMemberAttributes = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_ModifyUserGroupMemberAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.ModifyUserGroupMemberAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyUserGroupMemberAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteUserGroupMembers = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteUserGroupMembers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteUserGroupMembersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteUserGroupMembers(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeGroupRoles = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeGroupRoles(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeGroupRolesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeGroupRoles(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_CreateGroupRoles = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_CreateGroupRoles(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.CreateGroupRolesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateGroupRoles(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_ModifyGroupRoleAttributes = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_ModifyGroupRoleAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.ModifyGroupRoleAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyGroupRoleAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteGroupRoles = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteGroupRoles(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteGroupRolesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteGroupRoles(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeGroupRoleRules = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeGroupRoleRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeGroupRoleRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeGroupRoleRules(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_AddGroupRoleRules = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_AddGroupRoleRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.AddGroupRoleRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddGroupRoleRules(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_ModifyGroupRoleRuleAttributes = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_ModifyGroupRoleRuleAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.ModifyGroupRoleRuleAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyGroupRoleRuleAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DeleteGroupRoleRules = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DeleteGroupRoleRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DeleteGroupRoleRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteGroupRoleRules(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_GrantResourceGroupsToUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_GrantResourceGroupsToUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.GrantResourceGroupsToUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.GrantResourceGroupsToUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_RevokeResourceGroupsFromUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_RevokeResourceGroupsFromUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.RevokeResourceGroupsFromUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RevokeResourceGroupsFromUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_ResourceACLService_DescribeResourceUserGroups = []cli.Flag{ /* fields */ }

func _cmd_ResourceACLService_DescribeResourceUserGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewResourceACLService(conf, zone)

	in := new(pb.DescribeResourceUserGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeResourceUserGroups(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}