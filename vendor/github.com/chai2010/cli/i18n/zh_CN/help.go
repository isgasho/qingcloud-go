package zh_CN

import (
	"github.com/chai2010/cli"
)

func init() {
	cli.VersionFlag = VersionFlag
	cli.HelpFlag = HelpFlag

	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate
	cli.SubcommandHelpTemplate = SubcommandHelpTemplate

	cli.HelpCommand = HelpCommand
	cli.HelpSubcommand = HelpSubcommand
}

// This flag prints the version for the application
var VersionFlag = cli.BoolFlag{
	Name:  "version, v",
	Usage: "显示版本",
}

// This flag prints the help for all commands and subcommands
// Set to the zero value (BoolFlag{}) to disable flag -- keeps subcommand
// unless HideHelp is set to true)
var HelpFlag = cli.BoolFlag{
	Name:  "help, h",
	Usage: "显示帮助",
}

// The text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var AppHelpTemplate = `名称:
   {{.Name}} - {{.Usage}}

用法:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[全局选项]{{end}}{{if .Commands}} 命令 [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}
   {{if .Version}}
版本:
   {{.Version}}
   {{end}}{{if len .Authors}}
作者:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
命令列表:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}
全局选项:
   {{range .Flags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
版权:
   {{.Copyright}}
   {{end}}
`

// The text template for the command help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var CommandHelpTemplate = `名称:
   {{.HelpName}} - {{.Usage}}

用法:
   {{.HelpName}}{{if .Flags}} [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{if .Description}}

描述:
   {{.Description}}{{end}}{{if .Flags}}

选项:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`

// The text template for the subcommand help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
var SubcommandHelpTemplate = `名称:
   {{.HelpName}} - {{.Usage}}

用法:
   {{.HelpName}} command{{if .Flags}} [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}

命令列表:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
选项:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

var HelpCommand = cli.Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "显示子命令列表或一个子命令的帮助信息",
	ArgsUsage: "[命令]",
	Action: func(c *cli.Context) {
		args := c.Args()
		if args.Present() {
			cli.ShowCommandHelp(c, args.First())
		} else {
			cli.ShowAppHelp(c)
		}
	},
}

var HelpSubcommand = cli.Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "显示子命令列表或一个子命令的帮助信息",
	ArgsUsage: "[命令]",
	Action: func(c *cli.Context) {
		args := c.Args()
		if args.Present() {
			cli.ShowCommandHelp(c, args.First())
		} else {
			cli.ShowSubcommandHelp(c)
		}
	},
}
