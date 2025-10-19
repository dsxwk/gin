package command

import (
	"fmt"
	"gin/common/base"
	"gin/utils/cli"
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
		fmt.Println("解析参数失败:", err.Error())
		return
	}

	if *arg == "" {
		m.ExitError(`参数 --args 不能为空
Example: go run cli.go Demo:command --args=arg1 --desc=test-demo

Helper: go run cli.go Demo:command --help
`)
		return
	}

	fmt.Printf("执行命令: %s, 参数: %s\n", m.Name(), *arg)
}

func init() {
	cli.Register(&DemoCommand{})
}
