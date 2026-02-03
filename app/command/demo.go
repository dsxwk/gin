package command

import (
	"gin/common/base"
	"gin/pkg/cli"
	"github.com/fatih/color"
)

type DemoCommand struct {
	base.BaseCommand
}

func (m *DemoCommand) Name() string {
	return "demo-command"
}

func (m *DemoCommand) Description() string {
	return "test-demo"
}

func (m *DemoCommand) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			base.Flag{
				Short: "a",
				Long:  "args",
			},
			"示例参数, 如: arg1",
			true,
		},
	}
}

func (m *DemoCommand) Execute(args []string) {
	values := m.ParseFlags(m.Name(), args, m.Help())
	color.Green("执行命令: %s %s", m.Name(), m.FormatArgs(values))
}

func init() {
	cli.Register(&DemoCommand{})
}
