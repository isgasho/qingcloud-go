// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: load_balancer.proto

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
	AllCommands = append(AllCommands, CmdLoadBalancerService)
}

var CmdLoadBalancerService = cli.Command{
	Name:    "loadbalancer",
	Aliases: []string{},
	Usage:   "manage LoadBalancerService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateLoadBalancer",
			Aliases: []string{},
			Usage:   "CreateLoadBalancer",
			Flags:   _flag_LoadBalancerService_CreateLoadBalancer,
			Action:  _cmd_LoadBalancerService_CreateLoadBalancer,
		},
		{
			Name:    "DescribeLoadBalancers",
			Aliases: []string{},
			Usage:   "DescribeLoadBalancers",
			Flags:   _flag_LoadBalancerService_DescribeLoadBalancers,
			Action:  _cmd_LoadBalancerService_DescribeLoadBalancers,
		},
		{
			Name:    "DeleteLoadBalancers",
			Aliases: []string{},
			Usage:   "DeleteLoadBalancers",
			Flags:   _flag_LoadBalancerService_DeleteLoadBalancers,
			Action:  _cmd_LoadBalancerService_DeleteLoadBalancers,
		},
		{
			Name:    "ModifyLoadBalancerAttributes",
			Aliases: []string{},
			Usage:   "ModifyLoadBalancerAttributes",
			Flags:   _flag_LoadBalancerService_ModifyLoadBalancerAttributes,
			Action:  _cmd_LoadBalancerService_ModifyLoadBalancerAttributes,
		},
		{
			Name:    "StartLoadBalancers",
			Aliases: []string{},
			Usage:   "StartLoadBalancers",
			Flags:   _flag_LoadBalancerService_StartLoadBalancers,
			Action:  _cmd_LoadBalancerService_StartLoadBalancers,
		},
		{
			Name:    "StopLoadBalancers",
			Aliases: []string{},
			Usage:   "StopLoadBalancers",
			Flags:   _flag_LoadBalancerService_StopLoadBalancers,
			Action:  _cmd_LoadBalancerService_StopLoadBalancers,
		},
		{
			Name:    "UpdateLoadBalancers",
			Aliases: []string{},
			Usage:   "UpdateLoadBalancers",
			Flags:   _flag_LoadBalancerService_UpdateLoadBalancers,
			Action:  _cmd_LoadBalancerService_UpdateLoadBalancers,
		},
		{
			Name:    "ResizeLoadBalancers",
			Aliases: []string{},
			Usage:   "ResizeLoadBalancers",
			Flags:   _flag_LoadBalancerService_ResizeLoadBalancers,
			Action:  _cmd_LoadBalancerService_ResizeLoadBalancers,
		},
		{
			Name:    "AssociateEipsToLoadBalancer",
			Aliases: []string{},
			Usage:   "AssociateEipsToLoadBalancer",
			Flags:   _flag_LoadBalancerService_AssociateEipsToLoadBalancer,
			Action:  _cmd_LoadBalancerService_AssociateEipsToLoadBalancer,
		},
		{
			Name:    "DissociateEipsFromLoadBalancer",
			Aliases: []string{},
			Usage:   "DissociateEipsFromLoadBalancer",
			Flags:   _flag_LoadBalancerService_DissociateEipsFromLoadBalancer,
			Action:  _cmd_LoadBalancerService_DissociateEipsFromLoadBalancer,
		},
		{
			Name:    "AddLoadBalancerListeners",
			Aliases: []string{},
			Usage:   "AddLoadBalancerListeners",
			Flags:   _flag_LoadBalancerService_AddLoadBalancerListeners,
			Action:  _cmd_LoadBalancerService_AddLoadBalancerListeners,
		},
		{
			Name:    "DescribeLoadBalancerListeners",
			Aliases: []string{},
			Usage:   "DescribeLoadBalancerListeners",
			Flags:   _flag_LoadBalancerService_DescribeLoadBalancerListeners,
			Action:  _cmd_LoadBalancerService_DescribeLoadBalancerListeners,
		},
		{
			Name:    "DeleteLoadBalancerListeners",
			Aliases: []string{},
			Usage:   "DeleteLoadBalancerListeners",
			Flags:   _flag_LoadBalancerService_DeleteLoadBalancerListeners,
			Action:  _cmd_LoadBalancerService_DeleteLoadBalancerListeners,
		},
		{
			Name:    "ModifyLoadBalancerListenerAttributes",
			Aliases: []string{},
			Usage:   "ModifyLoadBalancerListenerAttributes",
			Flags:   _flag_LoadBalancerService_ModifyLoadBalancerListenerAttributes,
			Action:  _cmd_LoadBalancerService_ModifyLoadBalancerListenerAttributes,
		},
		{
			Name:    "AddLoadBalancerBackends",
			Aliases: []string{},
			Usage:   "AddLoadBalancerBackends",
			Flags:   _flag_LoadBalancerService_AddLoadBalancerBackends,
			Action:  _cmd_LoadBalancerService_AddLoadBalancerBackends,
		},
		{
			Name:    "DescribeLoadBalancerBackends",
			Aliases: []string{},
			Usage:   "DescribeLoadBalancerBackends",
			Flags:   _flag_LoadBalancerService_DescribeLoadBalancerBackends,
			Action:  _cmd_LoadBalancerService_DescribeLoadBalancerBackends,
		},
		{
			Name:    "DeleteLoadBalancerBackends",
			Aliases: []string{},
			Usage:   "DeleteLoadBalancerBackends",
			Flags:   _flag_LoadBalancerService_DeleteLoadBalancerBackends,
			Action:  _cmd_LoadBalancerService_DeleteLoadBalancerBackends,
		},
		{
			Name:    "ModifyLoadBalancerBackendAttributes",
			Aliases: []string{},
			Usage:   "ModifyLoadBalancerBackendAttributes",
			Flags:   _flag_LoadBalancerService_ModifyLoadBalancerBackendAttributes,
			Action:  _cmd_LoadBalancerService_ModifyLoadBalancerBackendAttributes,
		},
		{
			Name:    "CreateLoadBalancerPolicy",
			Aliases: []string{},
			Usage:   "CreateLoadBalancerPolicy",
			Flags:   _flag_LoadBalancerService_CreateLoadBalancerPolicy,
			Action:  _cmd_LoadBalancerService_CreateLoadBalancerPolicy,
		},
		{
			Name:    "DescribeLoadBalancerPolicies",
			Aliases: []string{},
			Usage:   "DescribeLoadBalancerPolicies",
			Flags:   _flag_LoadBalancerService_DescribeLoadBalancerPolicies,
			Action:  _cmd_LoadBalancerService_DescribeLoadBalancerPolicies,
		},
		{
			Name:    "ModifyLoadBalancerPolicyAttributes",
			Aliases: []string{},
			Usage:   "ModifyLoadBalancerPolicyAttributes",
			Flags:   _flag_LoadBalancerService_ModifyLoadBalancerPolicyAttributes,
			Action:  _cmd_LoadBalancerService_ModifyLoadBalancerPolicyAttributes,
		},
		{
			Name:    "ApplyLoadBalancerPolicy",
			Aliases: []string{},
			Usage:   "ApplyLoadBalancerPolicy",
			Flags:   _flag_LoadBalancerService_ApplyLoadBalancerPolicy,
			Action:  _cmd_LoadBalancerService_ApplyLoadBalancerPolicy,
		},
		{
			Name:    "DeleteLoadBalancerPolicies",
			Aliases: []string{},
			Usage:   "DeleteLoadBalancerPolicies",
			Flags:   _flag_LoadBalancerService_DeleteLoadBalancerPolicies,
			Action:  _cmd_LoadBalancerService_DeleteLoadBalancerPolicies,
		},
		{
			Name:    "AddLoadBalancerPolicyRules",
			Aliases: []string{},
			Usage:   "AddLoadBalancerPolicyRules",
			Flags:   _flag_LoadBalancerService_AddLoadBalancerPolicyRules,
			Action:  _cmd_LoadBalancerService_AddLoadBalancerPolicyRules,
		},
		{
			Name:    "DescribeLoadBalancerPolicyRules",
			Aliases: []string{},
			Usage:   "DescribeLoadBalancerPolicyRules",
			Flags:   _flag_LoadBalancerService_DescribeLoadBalancerPolicyRules,
			Action:  _cmd_LoadBalancerService_DescribeLoadBalancerPolicyRules,
		},
		{
			Name:    "ModifyLoadBalancerPolicyRuleAttributes",
			Aliases: []string{},
			Usage:   "ModifyLoadBalancerPolicyRuleAttributes",
			Flags:   _flag_LoadBalancerService_ModifyLoadBalancerPolicyRuleAttributes,
			Action:  _cmd_LoadBalancerService_ModifyLoadBalancerPolicyRuleAttributes,
		},
		{
			Name:    "DeleteLoadBalancerPolicyRules",
			Aliases: []string{},
			Usage:   "DeleteLoadBalancerPolicyRules",
			Flags:   _flag_LoadBalancerService_DeleteLoadBalancerPolicyRules,
			Action:  _cmd_LoadBalancerService_DeleteLoadBalancerPolicyRules,
		},
		{
			Name:    "CreateServerCertificate",
			Aliases: []string{},
			Usage:   "CreateServerCertificate",
			Flags:   _flag_LoadBalancerService_CreateServerCertificate,
			Action:  _cmd_LoadBalancerService_CreateServerCertificate,
		},
		{
			Name:    "DescribeServerCertificates",
			Aliases: []string{},
			Usage:   "DescribeServerCertificates",
			Flags:   _flag_LoadBalancerService_DescribeServerCertificates,
			Action:  _cmd_LoadBalancerService_DescribeServerCertificates,
		},
		{
			Name:    "ModifyServerCertificateAttributes",
			Aliases: []string{},
			Usage:   "ModifyServerCertificateAttributes",
			Flags:   _flag_LoadBalancerService_ModifyServerCertificateAttributes,
			Action:  _cmd_LoadBalancerService_ModifyServerCertificateAttributes,
		},
		{
			Name:    "DeleteServerCertificates",
			Aliases: []string{},
			Usage:   "DeleteServerCertificates",
			Flags:   _flag_LoadBalancerService_DeleteServerCertificates,
			Action:  _cmd_LoadBalancerService_DeleteServerCertificates,
		},
	},
}

var _flag_LoadBalancerService_CreateLoadBalancer = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_CreateLoadBalancer(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.CreateLoadBalancerInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateLoadBalancer(in)
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

var _flag_LoadBalancerService_DescribeLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeLoadBalancers(in)
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

var _flag_LoadBalancerService_DeleteLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteLoadBalancers(in)
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

var _flag_LoadBalancerService_ModifyLoadBalancerAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyLoadBalancerAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyLoadBalancerAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyLoadBalancerAttributes(in)
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

var _flag_LoadBalancerService_StartLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_StartLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.StartLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StartLoadBalancers(in)
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

var _flag_LoadBalancerService_StopLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_StopLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.StopLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.StopLoadBalancers(in)
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

var _flag_LoadBalancerService_UpdateLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_UpdateLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.UpdateLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.UpdateLoadBalancers(in)
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

var _flag_LoadBalancerService_ResizeLoadBalancers = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ResizeLoadBalancers(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ResizeLoadBalancersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ResizeLoadBalancers(in)
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

var _flag_LoadBalancerService_AssociateEipsToLoadBalancer = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_AssociateEipsToLoadBalancer(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.AssociateEipsToLoadBalancerInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AssociateEipsToLoadBalancer(in)
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

var _flag_LoadBalancerService_DissociateEipsFromLoadBalancer = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DissociateEipsFromLoadBalancer(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DissociateEipsFromLoadBalancerInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DissociateEipsFromLoadBalancer(in)
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

var _flag_LoadBalancerService_AddLoadBalancerListeners = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_AddLoadBalancerListeners(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.AddLoadBalancerListenersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddLoadBalancerListeners(in)
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

var _flag_LoadBalancerService_DescribeLoadBalancerListeners = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeLoadBalancerListeners(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeLoadBalancerListenersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeLoadBalancerListeners(in)
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

var _flag_LoadBalancerService_DeleteLoadBalancerListeners = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteLoadBalancerListeners(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteLoadBalancerListenersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteLoadBalancerListeners(in)
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

var _flag_LoadBalancerService_ModifyLoadBalancerListenerAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyLoadBalancerListenerAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyLoadBalancerListenerAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyLoadBalancerListenerAttributes(in)
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

var _flag_LoadBalancerService_AddLoadBalancerBackends = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_AddLoadBalancerBackends(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.AddLoadBalancerBackendsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddLoadBalancerBackends(in)
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

var _flag_LoadBalancerService_DescribeLoadBalancerBackends = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeLoadBalancerBackends(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeLoadBalancerBackendsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeLoadBalancerBackends(in)
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

var _flag_LoadBalancerService_DeleteLoadBalancerBackends = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteLoadBalancerBackends(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteLoadBalancerBackendsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteLoadBalancerBackends(in)
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

var _flag_LoadBalancerService_ModifyLoadBalancerBackendAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyLoadBalancerBackendAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyLoadBalancerBackendAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyLoadBalancerBackendAttributes(in)
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

var _flag_LoadBalancerService_CreateLoadBalancerPolicy = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_CreateLoadBalancerPolicy(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.CreateLoadBalancerPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateLoadBalancerPolicy(in)
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

var _flag_LoadBalancerService_DescribeLoadBalancerPolicies = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeLoadBalancerPolicies(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeLoadBalancerPoliciesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeLoadBalancerPolicies(in)
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

var _flag_LoadBalancerService_ModifyLoadBalancerPolicyAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyLoadBalancerPolicyAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyLoadBalancerPolicyAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyLoadBalancerPolicyAttributes(in)
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

var _flag_LoadBalancerService_ApplyLoadBalancerPolicy = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ApplyLoadBalancerPolicy(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ApplyLoadBalancerPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ApplyLoadBalancerPolicy(in)
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

var _flag_LoadBalancerService_DeleteLoadBalancerPolicies = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteLoadBalancerPolicies(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteLoadBalancerPoliciesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteLoadBalancerPolicies(in)
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

var _flag_LoadBalancerService_AddLoadBalancerPolicyRules = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_AddLoadBalancerPolicyRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.AddLoadBalancerPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddLoadBalancerPolicyRules(in)
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

var _flag_LoadBalancerService_DescribeLoadBalancerPolicyRules = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeLoadBalancerPolicyRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeLoadBalancerPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeLoadBalancerPolicyRules(in)
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

var _flag_LoadBalancerService_ModifyLoadBalancerPolicyRuleAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyLoadBalancerPolicyRuleAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyLoadBalancerPolicyRuleAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyLoadBalancerPolicyRuleAttributes(in)
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

var _flag_LoadBalancerService_DeleteLoadBalancerPolicyRules = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteLoadBalancerPolicyRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteLoadBalancerPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteLoadBalancerPolicyRules(in)
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

var _flag_LoadBalancerService_CreateServerCertificate = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_CreateServerCertificate(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.CreateServerCertificateInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateServerCertificate(in)
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

var _flag_LoadBalancerService_DescribeServerCertificates = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DescribeServerCertificates(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DescribeServerCertificatesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeServerCertificates(in)
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

var _flag_LoadBalancerService_ModifyServerCertificateAttributes = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_ModifyServerCertificateAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.ModifyServerCertificateAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyServerCertificateAttributes(in)
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

var _flag_LoadBalancerService_DeleteServerCertificates = []cli.Flag{ /* fields */ }

func _cmd_LoadBalancerService_DeleteServerCertificates(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewLoadBalancerService(conf, zone)

	in := new(pb.DeleteServerCertificatesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteServerCertificates(in)
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