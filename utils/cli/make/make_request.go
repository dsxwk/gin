package make

import (
	"fmt"
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

type MakeRequest struct {
	base.BaseCommand
}

func (m *MakeRequest) Name() string {
	return "make:request"
}

func (m *MakeRequest) Description() string {
	return "验证请求创建"
}

func (m *MakeRequest) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			"-f, --file",
			"文件路径, 如: user",
			true,
		},
		{
			"-d, --desc",
			"描述, 如: 用户请求验证",
			false,
		},
	}
}

func (m *MakeRequest) Execute(args []string) {
	var (
		fs    = pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
		_make = strings.TrimPrefix(m.Name(), "make:")
		file  = fs.StringP("file", "f", "", "文件路径, 如: user")
		desc  = fs.StringP("desc", "d", "Validator", "描述")
	)

	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	fmt.Printf("✅  创建服务: %s (描述: %s)\n", *file, *desc)
	f := m.GetMakeFile(*file, _make)

	m.generateFile(_make, f, *desc)
}

func init() {
	cli.Register(&MakeRequest{})
}

func (m *MakeRequest) generateFile(_make, file, desc string) {
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
		os.Exit(1)
	}

	data := struct {
		Package     string // 提取的包名
		Name        string // 模块名称(首字母大写)
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		color.Red("Error executing template:", err.Error())
		os.Exit(1)
	}

	color.Green("✅  验证请求文件: " + file + " 生成成功!")
}
