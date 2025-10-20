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

type MakeController struct {
	base.BaseCommand
}

func (m *MakeController) Name() string {
	return "make:controller"
}

func (m *MakeController) Description() string {
	return "控制器创建"
}

func (m *MakeController) Help() []base.CommandOption {
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
			"-m, --method",
			"请求方式, 如: get",
			false,
		},
		{
			"-r, --router",
			"路由地址, 如: /user",
			false,
		},
		{
			"-d, --desc",
			"描述, 如: 列表",
			false,
		},
	}
}

func (m *MakeController) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)

	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: v1/user")
	function := fs.StringP("function", "F", "FuncName", "Func Name, 如: index")
	method := fs.StringP("method", "m", "get", "Request Method 如: get")
	router := fs.StringP("router", "r", "/your/router", "Router Address 如: /user")
	desc := fs.StringP("desc", "d", "", "描述")
	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	log.Printf("✅ 创建控制器: %s (方法: %s 请求方式: %s 路由: %s 描述: %s)\n", *file, *function, *method, *router, *desc)
	f := m.GetMakeFile(*file, _make)

	m.generateFile(_make, f, *function, *method, *router, *desc)
}

func init() {
	cli.Register(&MakeController{})
}

func (m *MakeController) generateFile(_make, file, function, method, router, desc string) {
	templateFile := m.GetTemplate(_make)
	color.Green(fmt.Sprintf("✅ 解析模板文件: %v\n", templateFile))
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
		Router      string // 如果为空,使用默认值
		Method      string // 如果为空,使用默认值
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Function:    utils.UcFirst(function),
		Router:      router,
		Method:      method,
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Println("Error executing template:", err.Error())
		return
	} else {
		log.Println("Template executed and content written to file.")
	}

	color.Green("✅ 控制器文件: " + file + " 生成成功!")
}
