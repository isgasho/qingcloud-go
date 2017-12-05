// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: monitor.proto

package qcli_pb

import (
	"fmt"

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

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdMonitorService)
}

var CmdMonitorService = cli.Command{
	Name:    "monitor",
	Aliases: []string{},
	Usage:   "manage MonitorService",
	Subcommands: []cli.Command{
		{
			Name:    "GetMonitor",
			Aliases: []string{},
			Usage:   "GetMonitor",
			Flags:   _flag_MonitorService_GetMonitor,
			Action:  _cmd_MonitorService_GetMonitor,
		},
		{
			Name:    "GetLoadBalancerMonitor",
			Aliases: []string{},
			Usage:   "GetLoadBalancerMonitor",
			Flags:   _flag_MonitorService_GetLoadBalancerMonitor,
			Action:  _cmd_MonitorService_GetLoadBalancerMonitor,
		},
		{
			Name:    "GetRDBMonitor",
			Aliases: []string{},
			Usage:   "GetRDBMonitor",
			Flags:   _flag_MonitorService_GetRDBMonitor,
			Action:  _cmd_MonitorService_GetRDBMonitor,
		},
		{
			Name:    "GetCacheMonitor",
			Aliases: []string{},
			Usage:   "GetCacheMonitor",
			Flags:   _flag_MonitorService_GetCacheMonitor,
			Action:  _cmd_MonitorService_GetCacheMonitor,
		},
		{
			Name:    "GetZooKeeperMonitor",
			Aliases: []string{},
			Usage:   "GetZooKeeperMonitor",
			Flags:   _flag_MonitorService_GetZooKeeperMonitor,
			Action:  _cmd_MonitorService_GetZooKeeperMonitor,
		},
		{
			Name:    "GetQueueMonitor",
			Aliases: []string{},
			Usage:   "GetQueueMonitor",
			Flags:   _flag_MonitorService_GetQueueMonitor,
			Action:  _cmd_MonitorService_GetQueueMonitor,
		},
	},
}

var _flag_MonitorService_GetMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetMonitor(in)
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

var _flag_MonitorService_GetLoadBalancerMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetLoadBalancerMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetLoadBalancerMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetLoadBalancerMonitor(in)
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

var _flag_MonitorService_GetRDBMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetRDBMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetRDBMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetRDBMonitor(in)
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

var _flag_MonitorService_GetCacheMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetCacheMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetCacheMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetCacheMonitor(in)
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

var _flag_MonitorService_GetZooKeeperMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetZooKeeperMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetZooKeeperMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetZooKeeperMonitor(in)
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

var _flag_MonitorService_GetQueueMonitor = []cli.Flag{ /* fields */ }

func _cmd_MonitorService_GetQueueMonitor(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewMonitorService(conf, zone)

	in := new(pb.GetQueueMonitorInput)

	// TODO: fill field from flags

	out, err := qc.GetQueueMonitor(in)
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
