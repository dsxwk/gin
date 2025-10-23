package {{.Package}}

import (
	"fmt"
	"gin/common/base"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

type {{.Name}}Command struct {
	base.BaseCommand
}

func (m *{{.Name}}Command) Name() string {
    {{- if eq .Command "" }}
    return "{{.Name}}-command"
    {{- else }}
    return "{{.Command}}"
    {{- end }}
}

func (m *{{.Name}}Command) Description() string {
	return "{{.Description}}"
}

func (m *{{.Name}}Command) Help() []base.CommandOption {
	return []base.CommandOption{
		{
            "-a, --args",
            "示例参数, 如: arg1",
            true,
        },
	}
}

func (m *{{.Name}}Command) Execute(args []string) {
    fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
    arg := fs.StringP("args", "a", "", "示例参数, 如: arg1")
    if err := fs.Parse(args); err != nil {
        color.Red("参数解析失败: %s", err.Error())
    }

    fmt.Printf("执行命令: %s, 参数: %s\n", m.Name(), *arg)
}

func init() {
	cli.Register(&{{.Name}}Command{})
}
