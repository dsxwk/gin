package command

import (
	"fmt"
	"gin/common/base"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
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
			"-a, --args",
			"示例参数, 如: arg1",
			true,
		},
	}
}

func (m *DemoCommand) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	arg := fs.StringP("args", "a", "", "示例参数, 如: arg1")
	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}

	fmt.Printf("执行命令: %s, 参数: %s\n", m.Name(), *arg)
}

func init() {
	cli.Register(&DemoCommand{})
}
