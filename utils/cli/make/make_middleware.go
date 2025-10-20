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

type MakeMiddleware struct {
	base.BaseCommand
}

func (m *MakeMiddleware) Name() string {
	return "make:middleware"
}

func (m *MakeMiddleware) Description() string {
	return "中间件创建"
}

func (m *MakeMiddleware) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			"-f, --file",
			"文件路径, 如: v1/user",
			true,
		},
		{
			"-d, --desc",
			"描述, 如: 列表",
			false,
		},
	}
}

func (m *MakeMiddleware) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: v1/user")
	desc := fs.StringP("desc", "d", "middleware-desc", "描述")
	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	fmt.Printf("✅ 创建中间件: %s (描述: %s)\n", *file, *desc)
	f := m.GetMakeFile(*file, _make)

	m.generateFile(_make, f, *desc)
}

func init() {
	cli.Register(&MakeMiddleware{})
}

func (m *MakeMiddleware) generateFile(_make, file, desc string) {
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
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Println("Error executing template:", err.Error())
		return
	} else {
		log.Println("Template executed and content written to file.")
	}

	color.Green("✅ 中间件文件: " + file + " 生成成功!")
}
