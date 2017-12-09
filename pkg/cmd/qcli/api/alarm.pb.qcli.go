// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: alarm.proto

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
	AllCommands = append(AllCommands, CmdAlarmService)
}

var CmdAlarmService = cli.Command{
	Name:    "alarm",
	Aliases: []string{},
	Usage:   "manage AlarmService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeAlarmPolicies",
			Aliases: []string{},
			Usage:   "DescribeAlarmPolicies",
			Flags:   _flag_AlarmService_DescribeAlarmPolicies,
			Action:  _func_AlarmService_DescribeAlarmPolicies,
		},
		{
			Name:    "CreateAlarmPolicy",
			Aliases: []string{},
			Usage:   "CreateAlarmPolicy",
			Flags:   _flag_AlarmService_CreateAlarmPolicy,
			Action:  _func_AlarmService_CreateAlarmPolicy,
		},
		{
			Name:    "ModifyAlarmPolicyAttributes",
			Aliases: []string{},
			Usage:   "ModifyAlarmPolicyAttributes",
			Flags:   _flag_AlarmService_ModifyAlarmPolicyAttributes,
			Action:  _func_AlarmService_ModifyAlarmPolicyAttributes,
		},
		{
			Name:    "DeleteAlarmPolicies",
			Aliases: []string{},
			Usage:   "DeleteAlarmPolicies",
			Flags:   _flag_AlarmService_DeleteAlarmPolicies,
			Action:  _func_AlarmService_DeleteAlarmPolicies,
		},
		{
			Name:    "DescribeAlarmPolicyRules",
			Aliases: []string{},
			Usage:   "DescribeAlarmPolicyRules",
			Flags:   _flag_AlarmService_DescribeAlarmPolicyRules,
			Action:  _func_AlarmService_DescribeAlarmPolicyRules,
		},
		{
			Name:    "AddAlarmPolicyRules",
			Aliases: []string{},
			Usage:   "AddAlarmPolicyRules",
			Flags:   _flag_AlarmService_AddAlarmPolicyRules,
			Action:  _func_AlarmService_AddAlarmPolicyRules,
		},
		{
			Name:    "ModifyAlarmPolicyRuleAttributes",
			Aliases: []string{},
			Usage:   "ModifyAlarmPolicyRuleAttributes",
			Flags:   _flag_AlarmService_ModifyAlarmPolicyRuleAttributes,
			Action:  _func_AlarmService_ModifyAlarmPolicyRuleAttributes,
		},
		{
			Name:    "DeleteAlarmPolicyRules",
			Aliases: []string{},
			Usage:   "DeleteAlarmPolicyRules",
			Flags:   _flag_AlarmService_DeleteAlarmPolicyRules,
			Action:  _func_AlarmService_DeleteAlarmPolicyRules,
		},
		{
			Name:    "DescribeAlarmPolicyActions",
			Aliases: []string{},
			Usage:   "DescribeAlarmPolicyActions",
			Flags:   _flag_AlarmService_DescribeAlarmPolicyActions,
			Action:  _func_AlarmService_DescribeAlarmPolicyActions,
		},
		{
			Name:    "AddAlarmPolicyActions",
			Aliases: []string{},
			Usage:   "AddAlarmPolicyActions",
			Flags:   _flag_AlarmService_AddAlarmPolicyActions,
			Action:  _func_AlarmService_AddAlarmPolicyActions,
		},
		{
			Name:    "ModifyAlarmPolicyActionAttributes",
			Aliases: []string{},
			Usage:   "ModifyAlarmPolicyActionAttributes",
			Flags:   _flag_AlarmService_ModifyAlarmPolicyActionAttributes,
			Action:  _func_AlarmService_ModifyAlarmPolicyActionAttributes,
		},
		{
			Name:    "DeleteAlarmPolicyActions",
			Aliases: []string{},
			Usage:   "DeleteAlarmPolicyActions",
			Flags:   _flag_AlarmService_DeleteAlarmPolicyActions,
			Action:  _func_AlarmService_DeleteAlarmPolicyActions,
		},
		{
			Name:    "AssociateAlarmPolicy",
			Aliases: []string{},
			Usage:   "AssociateAlarmPolicy",
			Flags:   _flag_AlarmService_AssociateAlarmPolicy,
			Action:  _func_AlarmService_AssociateAlarmPolicy,
		},
		{
			Name:    "DissociateAlarmPolicy",
			Aliases: []string{},
			Usage:   "DissociateAlarmPolicy",
			Flags:   _flag_AlarmService_DissociateAlarmPolicy,
			Action:  _func_AlarmService_DissociateAlarmPolicy,
		},
		{
			Name:    "ApplyAlarmPolicy",
			Aliases: []string{},
			Usage:   "ApplyAlarmPolicy",
			Flags:   _flag_AlarmService_ApplyAlarmPolicy,
			Action:  _func_AlarmService_ApplyAlarmPolicy,
		},
		{
			Name:    "DescribeAlarms",
			Aliases: []string{},
			Usage:   "DescribeAlarms",
			Flags:   _flag_AlarmService_DescribeAlarms,
			Action:  _func_AlarmService_DescribeAlarms,
		},
		{
			Name:    "DescribeAlarmHistory",
			Aliases: []string{},
			Usage:   "DescribeAlarmHistory",
			Flags:   _flag_AlarmService_DescribeAlarmHistory,
			Action:  _func_AlarmService_DescribeAlarmHistory,
		},
	},
}

var _flag_AlarmService_DescribeAlarmPolicies = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policies",
		Usage: "alarm policies",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "alarm_policy_name",
		Usage: "alarm policy name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_type",
		Usage: "alarm policy type",
		Value: "",
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

func _func_AlarmService_DescribeAlarmPolicies(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DescribeAlarmPoliciesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policies") {
			if err := json.Unmarshal([]byte(c.String("alarm_policies")), &in.AlarmPolicies); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("alarm_policy_name") {
			in.AlarmPolicyName = proto.String(c.String("alarm_policy_name"))
		}
		if c.IsSet("alarm_policy_type") {
			in.AlarmPolicyType = proto.String(c.String("alarm_policy_type"))
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

	out, err := qc.DescribeAlarmPolicies(in)
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

var _flag_AlarmService_CreateAlarmPolicy = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy_type",
		Usage: "alarm policy type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "period",
		Usage: "period",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_name",
		Usage: "alarm policy name",
		Value: "",
	},
}

func _func_AlarmService_CreateAlarmPolicy(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.CreateAlarmPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy_type") {
			in.AlarmPolicyType = proto.String(c.String("alarm_policy_type"))
		}
		if c.IsSet("period") {
			in.Period = proto.String(c.String("period"))
		}
		if c.IsSet("alarm_policy_name") {
			in.AlarmPolicyName = proto.String(c.String("alarm_policy_name"))
		}
	}

	out, err := qc.CreateAlarmPolicy(in)
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

var _flag_AlarmService_ModifyAlarmPolicyAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_name",
		Usage: "alarm policy name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "period",
		Usage: "period",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_AlarmService_ModifyAlarmPolicyAttributes(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.ModifyAlarmPolicyAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("alarm_policy_name") {
			in.AlarmPolicyName = proto.String(c.String("alarm_policy_name"))
		}
		if c.IsSet("period") {
			in.Period = proto.String(c.String("period"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyAlarmPolicyAttributes(in)
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

var _flag_AlarmService_DeleteAlarmPolicies = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policies",
		Usage: "alarm policies",
		Value: "", // json: slice/message/map/time
	},
}

func _func_AlarmService_DeleteAlarmPolicies(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DeleteAlarmPoliciesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policies") {
			if err := json.Unmarshal([]byte(c.String("alarm_policies")), &in.AlarmPolicies); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteAlarmPolicies(in)
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

var _flag_AlarmService_DescribeAlarmPolicyRules = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_rules",
		Usage: "alarm policy rules",
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

func _func_AlarmService_DescribeAlarmPolicyRules(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DescribeAlarmPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("alarm_policy_rules") {
			if err := json.Unmarshal([]byte(c.String("alarm_policy_rules")), &in.AlarmPolicyRules); err != nil {
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

	out, err := qc.DescribeAlarmPolicyRules(in)
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

var _flag_AlarmService_AddAlarmPolicyRules = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "rules",
		Usage: "rules",
		Value: "", // json: slice/message/map/time
	},
}

func _func_AlarmService_AddAlarmPolicyRules(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.AddAlarmPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("rules") {
			if err := json.Unmarshal([]byte(c.String("rules")), &in.Rules); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddAlarmPolicyRules(in)
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

var _flag_AlarmService_ModifyAlarmPolicyRuleAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy_rule",
		Usage: "alarm policy rule",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_rule_name",
		Usage: "alarm policy rule name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "condition_type",
		Usage: "condition type",
		Value: "",
	},
	cli.StringFlag{
		Name:  "thresholds",
		Usage: "thresholds",
		Value: "",
	},
	cli.StringFlag{
		Name:  "data_processor",
		Usage: "data processor",
		Value: "",
	},
	cli.StringFlag{
		Name:  "consecutive_periods",
		Usage: "consecutive periods",
		Value: "",
	},
}

func _func_AlarmService_ModifyAlarmPolicyRuleAttributes(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.ModifyAlarmPolicyRuleAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy_rule") {
			in.AlarmPolicyRule = proto.String(c.String("alarm_policy_rule"))
		}
		if c.IsSet("alarm_policy_rule_name") {
			in.AlarmPolicyRuleName = proto.String(c.String("alarm_policy_rule_name"))
		}
		if c.IsSet("condition_type") {
			in.ConditionType = proto.String(c.String("condition_type"))
		}
		if c.IsSet("thresholds") {
			in.Thresholds = proto.String(c.String("thresholds"))
		}
		if c.IsSet("data_processor") {
			in.DataProcessor = proto.String(c.String("data_processor"))
		}
		if c.IsSet("consecutive_periods") {
			in.ConsecutivePeriods = proto.String(c.String("consecutive_periods"))
		}
	}

	out, err := qc.ModifyAlarmPolicyRuleAttributes(in)
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

var _flag_AlarmService_DeleteAlarmPolicyRules = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy_rules",
		Usage: "alarm policy rules",
		Value: "", // json: slice/message/map/time
	},
}

func _func_AlarmService_DeleteAlarmPolicyRules(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DeleteAlarmPolicyRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy_rules") {
			if err := json.Unmarshal([]byte(c.String("alarm_policy_rules")), &in.AlarmPolicyRules); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteAlarmPolicyRules(in)
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

var _flag_AlarmService_DescribeAlarmPolicyActions = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "alarm_policy_actions",
		Usage: "alarm policy actions",
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

func _func_AlarmService_DescribeAlarmPolicyActions(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DescribeAlarmPolicyActionsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("alarm_policy_actions") {
			if err := json.Unmarshal([]byte(c.String("alarm_policy_actions")), &in.AlarmPolicyActions); err != nil {
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

	out, err := qc.DescribeAlarmPolicyActions(in)
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

var _flag_AlarmService_AddAlarmPolicyActions = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "actions",
		Usage: "actions",
		Value: "", // json: slice/message/map/time
	},
}

func _func_AlarmService_AddAlarmPolicyActions(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.AddAlarmPolicyActionsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("actions") {
			if err := json.Unmarshal([]byte(c.String("actions")), &in.Actions); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.AddAlarmPolicyActions(in)
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

var _flag_AlarmService_ModifyAlarmPolicyActionAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy_action",
		Usage: "alarm policy action",
		Value: "",
	},
	cli.StringFlag{
		Name:  "trigger_action",
		Usage: "trigger action",
		Value: "",
	},
	cli.StringFlag{
		Name:  "trigger_status",
		Usage: "trigger status",
		Value: "",
	},
}

func _func_AlarmService_ModifyAlarmPolicyActionAttributes(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.ModifyAlarmPolicyActionAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy_action") {
			in.AlarmPolicyAction = proto.String(c.String("alarm_policy_action"))
		}
		if c.IsSet("trigger_action") {
			in.TriggerAction = proto.String(c.String("trigger_action"))
		}
		if c.IsSet("trigger_status") {
			in.TriggerStatus = proto.String(c.String("trigger_status"))
		}
	}

	out, err := qc.ModifyAlarmPolicyActionAttributes(in)
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

var _flag_AlarmService_DeleteAlarmPolicyActions = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy_actions",
		Usage: "alarm policy actions",
		Value: "", // json: slice/message/map/time
	},
}

func _func_AlarmService_DeleteAlarmPolicyActions(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DeleteAlarmPolicyActionsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy_actions") {
			if err := json.Unmarshal([]byte(c.String("alarm_policy_actions")), &in.AlarmPolicyActions); err != nil {
				log.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteAlarmPolicyActions(in)
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

var _flag_AlarmService_AssociateAlarmPolicy = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "resources",
		Usage: "resources",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "related_resource",
		Usage: "related resource",
		Value: "",
	},
}

func _func_AlarmService_AssociateAlarmPolicy(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.AssociateAlarmPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("resources") {
			if err := json.Unmarshal([]byte(c.String("resources")), &in.Resources); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("related_resource") {
			in.RelatedResource = proto.String(c.String("related_resource"))
		}
	}

	out, err := qc.AssociateAlarmPolicy(in)
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

var _flag_AlarmService_DissociateAlarmPolicy = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "resources",
		Usage: "resources",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "related_resource",
		Usage: "related resource",
		Value: "",
	},
}

func _func_AlarmService_DissociateAlarmPolicy(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DissociateAlarmPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
		if c.IsSet("resources") {
			if err := json.Unmarshal([]byte(c.String("resources")), &in.Resources); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("related_resource") {
			in.RelatedResource = proto.String(c.String("related_resource"))
		}
	}

	out, err := qc.DissociateAlarmPolicy(in)
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

var _flag_AlarmService_ApplyAlarmPolicy = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm_policy",
		Usage: "alarm policy",
		Value: "",
	},
}

func _func_AlarmService_ApplyAlarmPolicy(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.ApplyAlarmPolicyInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm_policy") {
			in.AlarmPolicy = proto.String(c.String("alarm_policy"))
		}
	}

	out, err := qc.ApplyAlarmPolicy(in)
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

var _flag_AlarmService_DescribeAlarms = []cli.Flag{
	cli.StringFlag{
		Name:  "alarms",
		Usage: "alarms",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "policy",
		Usage: "policy",
		Value: "",
	},
	cli.StringFlag{
		Name:  "status",
		Usage: "status",
		Value: "",
	},
	cli.StringFlag{
		Name:  "resource",
		Usage: "resource",
		Value: "",
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

func _func_AlarmService_DescribeAlarms(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DescribeAlarmsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarms") {
			if err := json.Unmarshal([]byte(c.String("alarms")), &in.Alarms); err != nil {
				log.Fatal(err)
			}
		}
		if c.IsSet("policy") {
			in.Policy = proto.String(c.String("policy"))
		}
		if c.IsSet("status") {
			in.Status = proto.String(c.String("status"))
		}
		if c.IsSet("resource") {
			in.Resource = proto.String(c.String("resource"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeAlarms(in)
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

var _flag_AlarmService_DescribeAlarmHistory = []cli.Flag{
	cli.StringFlag{
		Name:  "alarm",
		Usage: "alarm",
		Value: "",
	},
	cli.StringFlag{
		Name:  "history_type",
		Usage: "history type",
		Value: "",
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

func _func_AlarmService_DescribeAlarmHistory(c *cli.Context) error {
	qc := pb.NewAlarmService(nil, nil)
	in := new(pb.DescribeAlarmHistoryInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("alarm") {
			in.Alarm = proto.String(c.String("alarm"))
		}
		if c.IsSet("history_type") {
			in.HistoryType = proto.String(c.String("history_type"))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeAlarmHistory(in)
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
