package make

import (
	"gin/common/base"
	"gin/utils"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type MakeCommand struct {
	base.BaseCommand
}

func (m *MakeCommand) Name() string {
	return "make:command"
}

func (m *MakeCommand) Description() string {
	return "服务创建"
}

func (m *MakeCommand) Help() []base.CommandOption {
	return []base.CommandOption{
		{"-f, --file", "文件路径, 如: cronjob/demo", true},
		{"-m, --name", "命令名称, 如: demo.command", false},
		{"-d, --desc", "描述, 如: demo-desc", false},
	}
}

func (m *MakeCommand) Execute(args []string) {
	var (
		fs    = pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
		_make = strings.TrimPrefix(m.Name(), "make:")
		file  = fs.StringP("file", "f", "", "文件路径, 如: cronjob/demo")
		name  = fs.StringP("name", "m", "", "demo.command")
		desc  = fs.StringP("desc", "d", "", "command-desc")
	)

	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	color.Green("✅  创建命令: %s (命令: %s 描述: %s)\n", *file, *name, *desc)
	f := m.GetMakeFile(*file, _make)
	m.generateFile(_make, f, *name, *desc)
}

func init() {
	cli.Register(&MakeCommand{})
}

func (m *MakeCommand) generateFile(_make, file, name, desc string) {
	templateFile := m.GetTemplate(_make)
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		color.Red("Error parsing template:", err.Error())
		os.Exit(1)
	}

	// 提取包名 (文件路径中的最后一个目录作为包名)
	packageName := filepath.Base(filepath.Dir(file))

	// 创建文件
	f := m.CheckDirAndFile(file)
	if f == nil {
		return
	}

	data := struct {
		Package     string // 提取的包名
		Name        string // 模块名称(首字母大写)
		Command     string // 如果为空,使用默认值
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Command:     name,
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		color.Red("Error executing template:", err.Error())
		os.Exit(1)
	}

	color.Green("✅ 命令行文件: " + file + " 生成成功!")
}
