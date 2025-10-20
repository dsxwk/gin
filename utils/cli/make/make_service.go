package make

import (
	"fmt"
	"gin/common/base"
	"gin/utils"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"html/template"
	"log"
	"path/filepath"
	"strings"
)

type MakeService struct {
	base.BaseCommand
}

func (m *MakeService) Name() string {
	return "make:service"
}

func (m *MakeService) Description() string {
	return "服务创建"
}

func (m *MakeService) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			"-f, --file",
			"文件路径, 如: v1/user",
			true,
		},
		{
			"-F, --function",
			"方法名称, 如: list",
			false,
		},
		{
			"-d, --desc",
			"描述, 如: 列表",
			false,
		},
	}
}

func (m *MakeService) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: v1/user")
	function := fs.StringP("function", "F", "FuncName", "Func Name, 如: index")
	desc := fs.StringP("desc", "d", "", "描述")
	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	fmt.Printf("✅ 创建服务: %s (方法: %s 描述: %s)\n", *file, *function, *desc)
	f := m.GetMakeFile(*file, _make)

	m.generateFile(_make, f, *function, *desc)
}

func init() {
	cli.Register(&MakeService{})
}

func (m *MakeService) generateFile(_make, file, function, desc string) {
	templateFile := m.GetTemplate(_make)
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Error parsing template:", err.Error())
		return
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
		Function    string // 如果为空,使用默认值
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Function:    utils.UcFirst(function),
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Println("Error executing template:", err.Error())
		return
	} else {
		log.Println("Template executed and content written to file.")
	}

	color.Green("✅ 服务文件: " + file + " 生成成功!")
}
