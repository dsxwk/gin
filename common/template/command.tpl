package {{.Package}}

import (
	"fmt"
	"gin/common/base"
	"gin/utils/cli"
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
		{"-a, --args", "示例参数, 如: arg1 (参数1必填)"},
	}
}

func (m *{{.Name}}Command) Execute(args []string) {
    fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
    arg := fs.StringP("args", "a", "", "示例参数, 如: arg1 (参数1必填)")

    if err := fs.Parse(args); err != nil {
        fmt.Println("解析参数失败:", err.Error())
        return
    }

    if *arg == "" {
        m.ExitError("参数 --args 不能为空\nExample: go run cli.go {{.Name}}:command --args=arg1 --desc=方法描述\nHelper: go run cli.go {{.Name}}:command --help")
        return
    }

    fmt.Printf("执行命令: %s, 参数: %s\n", m.Name(), *arg)
}

func init() {
	cli.AutoRegister(&{{.Name}}Command{})
}
